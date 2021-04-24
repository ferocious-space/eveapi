package yamlparser

import (
	"bufio"
	"fmt"
	"os"
	"reflect"

	"github.com/go-openapi/inflect"
	jsoniter "github.com/json-iterator/go"
	"gopkg.in/yaml.v3"
)

func dumpObj(i interface{}) {
	bin, _ := jsoniter.MarshalIndent(i, "", " ")
	fmt.Println(string(bin))
}

func RawUnmarshaler(path string) (yaml.Node, error) {
	var out yaml.Node
	dataFile, err := os.Open(path)
	if err != nil {
		return yaml.Node{}, err
	}
	decoder := yaml.NewDecoder(bufio.NewReader(dataFile))
	if err := decoder.Decode(&out); err != nil {
		return yaml.Node{}, err
	}
	return out, nil
}

func NameForType(typeName string) (singularName string, pluralName string) {
	singularName = inflect.Singularize(typeName)
	pluralName = inflect.Pluralize(singularName)
	return
}

func ParseString(yamlString string) (reflect.Type, error) {
	ntf := yaml.Node{}
	if err := yaml.Unmarshal([]byte(yamlString), &ntf); err != nil {
		return nil, err
	}
	return ParseYAML(&ntf)
}

func ParseYAML(node *yaml.Node) (reflect.Type, error) {
	return nodeWalk(node, 0, nil), nil
}

func ParseFile(yamlPath string) (reflect.Type, error) {
	data, err := RawUnmarshaler(yamlPath)
	if err != nil {
		return nil, err
	}
	return nodeWalk(&data, 0, nil), nil
}

func DeepMergeList(types []reflect.Type) reflect.Type {
	if len(types) > 0 {
		if len(types) == 1 {
			return types[0]
		}
		first := types[0]
		rest := types[1:]
		for _, t := range rest {
			first = deepMerge(first, t)
		}
		return first
	}
	return nil
}

func deepMerge(src reflect.Type, dst reflect.Type) reflect.Type {
	if src == nil {
		if dst == nil {
			return typeEmpty()
		}
		return dst
	}
	if dst == nil {
		return src
	}
	if src == dst {
		return src
	}

	// handle empty
	if isTypeEmpty(src) {
		if isTypeEmpty(dst) {
			return typeEmpty()
		}
		return dst
	} else {
		if isTypeEmpty(dst) {
			return src
		}
	}

	isPtr := false
	stype := src.String()
	dtype := dst.String()
	_, _ = stype, dtype
	var source, dest reflect.Type
	source = src
	if src.Kind() == reflect.Ptr {
		isPtr = true
		source = src.Elem()
	}
	dest = dst
	if dst.Kind() == reflect.Ptr {
		isPtr = true
		dest = dst.Elem()
	}

	// they are the same type
	if source == dest {
		return src
	}
	// they are same kind
	if source.Kind() == dest.Kind() {
		switch source.Kind() {
		case reflect.Struct:
			var mergeFields []reflect.StructField
			for i := 0; i < source.NumField(); i++ {
				mergeFields = append(mergeFields, source.Field(i))
			}
			for i := 0; i < dest.NumField(); i++ {
				sField := dest.Field(i)
				if isTypeEmpty(sField.Type) {
					continue
				}
				found := false
				index := 0
				for k, mField := range mergeFields {
					if mField.Name == sField.Name {
						found = true
						index = k
						break
					}
				}
				// not found , just add
				if !found {
					mergeFields = append(mergeFields, sField)
					continue
				}
				mergeFields[index].Type = deepMerge(mergeFields[index].Type, sField.Type)
			}
			resolved := reflect.StructOf(mergeFields)
			if isPtr {
				resolved = reflect.PtrTo(resolved)
			}
			resolvedText := resolved.String()
			_ = resolvedText
			return resolved
		case reflect.Map:
			resolved := reflect.MapOf(source.Key(), deepMerge(source.Elem(), dest.Elem()))
			resolvedText := resolved.String()
			_ = resolvedText
			return resolved
		case reflect.Slice:
			resolved := reflect.SliceOf(deepMerge(source.Elem(), dest.Elem()))
			resolvedText := resolved.String()
			_ = resolvedText
			return resolved
		}
	}
	// they are different kind
	if source.Kind() == reflect.Int32 && dest.Kind() == reflect.Float64 {
		return dst
	}
	if source.Kind() == reflect.Float64 && dest.Kind() == reflect.Int32 {
		return src
	}
	if source.Kind() == reflect.String {
		return src
	}
	return src
}
