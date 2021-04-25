package yamlparser

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/go-openapi/inflect"
	"gopkg.in/yaml.v3"
)

type empty struct{}

type timestruct struct{ uint }

func typeEmpty() reflect.Type {
	return reflect.TypeOf(empty{})
}

func isTypeEmpty(in reflect.Type) bool {
	if in.Kind() == reflect.Ptr {
		return in.Elem() == typeEmpty()
	}
	return in == typeEmpty()
}

func ptrValue(in interface{}) reflect.Type {
	return reflect.PtrTo(reflect.TypeOf(in))
}

type typeMapping struct {
	IsPrefix bool
	IsSuffix bool
	AnyMatch bool
	IsType   reflect.Type
	WantType reflect.Type
}

func nodeWalk(node *yaml.Node, depth int, mappings map[string]typeMapping) reflect.Type {
	if mappings == nil {
		mappings = make(map[string]typeMapping)
	}
	// switch node.Tag {
	// case "!!map":
	// 	fmt.Printf("%s>>> %s(%d)\n", strings.Repeat("\t", depth), node.Tag, len(node.Content)/2)
	// case "!!seq":
	// 	fmt.Printf("%s>>> %s(%d)\n", strings.Repeat("\t", depth), node.Tag, len(node.Content))
	// default:
	// 	fmt.Printf("%s>>> %s\n", strings.Repeat("\t", depth), node.Tag)
	// }
	// defer fmt.Printf("%s<<< %s - %d\n", strings.Repeat("\t", depth), node.Tag, node.Kind)
	if node.Tag == "" {
		if len(node.Content) == 1 {
			return nodeWalk(node.Content[0], depth, mappings)
		}
	}
	switch node.Tag {
	case "!!timestamp":
		return reflect.TypeOf(time.Time{})
	case "!!str":
		return ptrValue(string(""))
	case "!!bool":
		return ptrValue(true)
	case "!!int":
		return ptrValue(int32(0))
	case "!!float":
		return ptrValue(float64(0))
	case "!!map":
		mapping := make(map[string]bool)
		for i := 0; i < len(node.Content); i = i + 2 {
			mapping[node.Content[i].Tag] = true
		}
		if len(mapping) > 1 {
			panic(mapping)
		}
		if len(mapping) == 0 {

		}
		for k := range mapping {
			switch k {
			case "!!str":
				var fields []reflect.StructField
				for i := 0; i < len(node.Content); i = i + 2 {
					// fmt.Printf("%smap[%s(%s)]%s (allstring)\n", strings.Repeat("\t", depth), node.Content[i].Tag, node.Content[i].Value, node.Content[i+1].Tag)

					fieldName := inflect.Capitalize(inflect.Camelize(node.Content[i].Value))
					nodeType := nodeWalk(node.Content[i+1], depth+1, mappings)
					nodeTypeStr := nodeType.String()
					_ = nodeTypeStr
					structTag := ""
					isPtr := false
					if nodeType.Kind() == reflect.Ptr {
						isPtr = true
						nodeType = nodeType.Elem()
						structTag = fmt.Sprintf(`yaml:"%s,omitempty"`, node.Content[i].Value)
					}

					switch nodeType.Kind() {
					case reflect.Struct:
						nodeType = reflect.PtrTo(nodeType)
						structTag = fmt.Sprintf(`yaml:"%s,omitempty"`, node.Content[i].Value)
					case reflect.Map, reflect.Slice:
						structTag = fmt.Sprintf(`yaml:"%s,omitempty"`, node.Content[i].Value)
					case reflect.Int32:
						switch fieldName {
						case "LocationID", "ItemID", "RecordID", "FleetID", "WingID", "SquadID", "StructureID", "ContextID", "RouteID", "PinID", "DestinationPinID", "SourcePinID", "Amount", "Waypoint":
							nodeType = reflect.TypeOf(int64(0))
						}
						switch inflect.Singularize(fieldName) {
						case "LocationID", "ItemID", "RecordID", "FleetID", "WingID", "SquadID", "StructureID", "ContextID", "RouteID", "PinID", "DestinationPinID", "SourcePinID", "Amount", "Waypoint":
							nodeType = reflect.TypeOf(int64(0))
						}
						if strings.Contains(fieldName, "Time") {
							nodeType = reflect.TypeOf(int64(0))
						}
					default:
						if structTag == "" {
							structTag = fmt.Sprintf(`yaml:"%s"`, node.Content[i].Value)
						}
					}
					for k, v := range mappings {
						matchName := false
						s, p := NameForType(fieldName)
						if v.AnyMatch && (strings.Contains(s, k) || strings.Contains(p, k)) {
							matchName = true
						}
						if !matchName && v.IsPrefix && (strings.HasPrefix(s, k) || strings.HasPrefix(p, k)) {
							matchName = true
						}
						if !matchName && v.IsSuffix && (strings.HasSuffix(s, k) || strings.HasSuffix(p, k)) {
							matchName = true
						}
						if !matchName && (s == k || p == k) {
							matchName = true
						}
						if matchName {
							if nodeType == v.IsType {
								nodeType = v.WantType
							}
						}
					}
					if isPtr {
						nodeType = reflect.PtrTo(nodeType)
					}
					fields = append(
						fields, reflect.StructField{
							Name: fieldName,
							Type: nodeType,
							Tag:  reflect.StructTag(structTag),
						},
					)

				}
				strRet := reflect.StructOf(fields)
				strRetName := strRet.String()
				_ = strRetName
				return strRet
			case "!!int":
				values := nodeWalk(node.Content[1], depth+1, mappings)
				for i := 2; i < len(node.Content); i = i + 2 {
					// fmt.Printf("%s%s(%s) - %s (allint)\n", strings.Repeat("\t", depth), node.Content[i].Tag, node.Content[i].Value, node.Content[i+1].Tag)
					values = deepMerge(values, nodeWalk(node.Content[i+1], depth+1, mappings))
				}
				mapRet := reflect.MapOf(reflect.TypeOf(int32(0)), values)
				mapType := mapRet.String()
				_ = mapType
				return mapRet
			case "!!float":
				values := nodeWalk(node.Content[1], depth+1, mappings)
				for i := 2; i < len(node.Content); i = i + 2 {
					// fmt.Printf("%s%s(%s) - %s (allfloat)\n", strings.Repeat("\t", depth), node.Content[i].Tag, node.Content[i].Value, node.Content[i+1].Tag)
					values = deepMerge(values, nodeWalk(node.Content[i+1], depth+1, mappings))
				}
				mapRet := reflect.MapOf(reflect.TypeOf(float64(0)), values)
				mapType := mapRet.String()
				_ = mapType
				return mapRet
			default:
				return typeEmpty()
			}
		}
		return typeEmpty()
	case "!!seq":
		mapping := make(map[string]bool)
		for i := 0; i < len(node.Content); i++ {
			// fmt.Printf("%s%s\n", strings.Repeat("\t", depth), node.Content[i].Tag)
			mapping[node.Content[i].Tag] = true
		}
		if len(mapping) > 1 {
			return reflect.SliceOf(reflect.TypeOf(string("")))
		}
		for k := range mapping {
			switch k {
			case "!!str":
				return reflect.SliceOf(reflect.TypeOf(string("")))
			case "!!int":
				return reflect.SliceOf(reflect.TypeOf(int32(0)))
			case "!!float":
				return reflect.SliceOf(reflect.TypeOf(float64(0)))
			case "!!seq":
				alltypes := make([]reflect.Type, 0)
				for _, elems := range node.Content {
					el := nodeWalk(elems, depth+1, mappings)
					if el.Kind() == reflect.Ptr {
						el = el.Elem()
					}
					alltypes = append(alltypes, el)
				}
				return reflect.SliceOf(DeepMergeList(alltypes))
			case "!!map":
				alltypes := make([]reflect.Type, 0)
				for _, elems := range node.Content {
					el := nodeWalk(elems, depth+1, mappings)
					if el.Kind() == reflect.Ptr {
						el = el.Elem()
					}
					alltypes = append(alltypes, el)
				}
				return reflect.SliceOf(DeepMergeList(alltypes))
			default:
				return reflect.SliceOf(reflect.TypeOf(string("")))
			}
		}
		return reflect.SliceOf(reflect.TypeOf(empty{}))
	default:
		return typeEmpty()
	}
}
