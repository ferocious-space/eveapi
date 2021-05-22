package yamlparser

import (
	"reflect"
	"strings"
	"time"

	. "github.com/dave/jennifer/jen"
	"github.com/go-openapi/inflect"
)

func jenGenerator(in reflect.Type, typeName string, depth int) (*Statement, map[string]*Statement) {
	extra := make(map[string]*Statement)
	ref := in
	t := Empty()
	switch ref.Kind() {
	case reflect.Ptr:
		elem, x := jenGenerator(ref.Elem(), typeName, depth+1)
		for k, v := range x {
			extra[k] = v
		}
		t.Add(Op("*").Add(elem))
	case reflect.Map:
		field := ref.Elem()
		isEmpty := false
		isStruct := false
		isPtr := false
		isTimeStruct := false
		if field.Kind() == reflect.Ptr {
			isPtr = true
			field = field.Elem()
		}
		if field.Kind() == reflect.Struct {
			isStruct = true
			if field.NumField() == 0 {
				isEmpty = true
			}
		}

		if field == reflect.TypeOf(time.Time{}) {
			isTimeStruct = true
		}

		key, _ := jenGenerator(ref.Key(), typeName, depth+1)
		if isStruct {
			if isEmpty {
				t.Add(Map(key).Interface())
			} else {
				if isTimeStruct {
					t.Add(Map(key).Qual("time", "Time"))
				} else {
					singularName := inflect.Singularize(typeName)
					if depth > 0 {
						singularName = typeName
					}
					elem, x := jenGenerator(field, singularName, depth+1)
					for k, v := range x {
						extra[k] = v
					}
					if isPtr {
						t.Add(Map(key).Op("*").Id(singularName))
					} else {
						t.Add(Map(key).Id(singularName))
					}
					extra[singularName] = elem
				}
			}
		} else {
			elem, x := jenGenerator(field, typeName, depth+1)
			for k, v := range x {
				extra[k] = v
			}
			t.Add(Map(key).Add(elem))
		}

	case reflect.Slice:
		field := ref.Elem()
		isStruct := false
		isPtr := false
		isEmpty := false
		isTimeStruct := false
		if field.Kind() == reflect.Ptr {
			isPtr = true
			field = field.Elem()
		}
		if field.Kind() == reflect.Struct {
			isStruct = true
			if field.NumField() == 0 {
				isEmpty = true
			}
		}

		if field == reflect.TypeOf(time.Time{}) {
			isTimeStruct = true
		}

		singularName := inflect.Singularize(typeName)
		if depth > 0 {
			singularName = typeName
		}
		elem, x := jenGenerator(field, singularName, depth+1)
		for k, v := range x {
			extra[k] = v
		}
		if isStruct {
			if isEmpty {
				t.Index().Interface()
			} else {
				if isTimeStruct {
					t.Qual("time", "Time")
				} else {
					if isPtr {
						t.Index().Op("*").Id(singularName)
					} else {
						t.Index().Id(singularName)
					}
					extra[singularName] = elem
				}
			}
		} else {
			t.Index().Add(elem)
		}
	case reflect.String:
		t.Add(String())
	case reflect.Bool:
		t.Add(Bool())
	case reflect.Float64:
		t.Add(Float64())
	case reflect.Int32:
		t.Add(Int32())
	case reflect.Int64:
		t.Add(Int64())
	case reflect.Int:
		t.Add(Int())
	case reflect.Uint8:
		t.Add(Uint8())
	case reflect.Uint64:
		t.Add(Uint64())
	case reflect.Uint:
		t.Add(Uint())
	case reflect.Struct:
		fields := make([]Code, 0)
		if ref.NumField() == 0 {
			t.Add(Interface())
		} else {
			for i := 0; i < ref.NumField(); i++ {

				field := ref.Field(i)
				fieldName := field.Name
				singularName := inflect.Singularize(typeName) + fieldName

				if strings.HasSuffix(fieldName, "ID") {
					singularName = inflect.Singularize(typeName) + inflect.Singularize(fieldName)
				}

				if depth > 1 {
					singularName = typeName + fieldName
					if strings.HasSuffix(fieldName, "ID") {
						singularName = typeName + inflect.Singularize(fieldName)
					}
				}

				isEmpty := false
				isStruct := false
				isPtr := false
				isTimeStruct := false

				if field.Type.Kind() == reflect.Ptr {
					isPtr = true
					field.Type = field.Type.Elem()
				}

				if field.Type == reflect.TypeOf(time.Time{}) {
					isTimeStruct = true
				}

				if field.Type.Kind() == reflect.Struct {
					isStruct = true
					if field.Type.NumField() == 0 {
						isEmpty = true
					}
				}

				content := field.Tag.Get("yaml")

				tags := map[string]string{
					"yaml": content,
					"json": content,
					"bson": content,
					"db":   content,
				}

				elem, x := jenGenerator(field.Type, singularName, depth+1)
				for k, v := range x {
					extra[k] = v
				}

				if isStruct {
					if isEmpty {
						fields = append(fields, Id(fieldName).Add(Interface()).Tag(tags))
					} else {
						if isTimeStruct {
							fields = append(fields, Id(fieldName).Add(Qual("time", "Time")).Tag(tags))
							extra = map[string]*Statement{}
						} else {
							if isPtr {
								fields = append(fields, Id(fieldName).Op("*").Id(singularName).Tag(tags))
							} else {
								fields = append(fields, Id(fieldName).Id(singularName).Tag(tags))
							}
							extra[singularName] = Empty().Add(elem)
						}
					}
				} else {
					if isPtr {
						fields = append(fields, Id(fieldName).Op("*").Add(elem).Tag(tags))
					} else {
						fields = append(fields, Id(fieldName).Add(elem).Tag(tags))
					}
				}
			}
			t.Add(Struct(fields...))
		}
	}
	return t, extra
}
