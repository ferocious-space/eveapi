package yamlparser

import (
	"bytes"
	"reflect"
	"sort"

	"github.com/dave/jennifer/jen"
)

func GenerateType(p reflect.Type, typeName string, packagePath string, packageName string) ([]byte, error) {
	singularName, _ := NameForType(typeName)
	buffer := new(bytes.Buffer)
	pkg := jen.NewFilePathName(packagePath, packageName)
	pkg.HeaderComment("// Code generated by YAML2GO. DO NOT EDIT.")
	main, extra := jenGenerator(p, typeName, 0)
	switch p.Kind() {
	case reflect.Slice:
		pkg.Add(jen.Type().Id(singularName + "List").Add(main))
		pkg.Add(GenerateLoader(singularName + "List"))
	case reflect.Map:
		pkg.Add(jen.Type().Id(singularName + "Map").Add(main))
		pkg.Add(GenerateLoader(singularName + "Map"))
		pkg.Add(GenerateMapGetter(singularName+"Map", singularName))
	default:
		pkg.Add(jen.Type().Id(singularName).Add(main))
	}
	tNames := []string{}
	for n := range extra {
		tNames = append(tNames, n)
	}
	sort.Strings(tNames)
	for _, xType := range tNames {
		pkg.Add(jen.Type().Id(xType).Add(extra[xType]))
	}
	if err := pkg.Render(buffer); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func GenerateLoader(typeName string) *jen.Statement {
	return jen.Func().Params(jen.Id("x").Op("*").Id(typeName)).Id("Load").Params(jen.Id("path").String()).Error().Block(
		jen.List(jen.Id("f"), jen.Err()).Op(":=").Qual("os", "Open").Call(jen.Id("path")),
		jen.If(jen.Err().Op("!=").Nil()).Block(
			jen.Return(jen.Err()),
		),
		jen.Defer().Id("f").Dot("Close").Call(),
		jen.Return(jen.Qual("gopkg.in/yaml.v3", "NewDecoder").Call(jen.Id("f"))).Dot("Decode").Call(jen.Id("x")),
	)
}

func GenerateMapGetter(typeName, returnName string) *jen.Statement {
	return jen.Func().Params(jen.Id("x").Id(typeName)).Id("Get").Params(jen.Id("ID").Int32()).Op("*").Id(returnName).Block(
		jen.If(
			jen.List(jen.Id("a"), jen.Id("ok")).Op(":=").Id("x").Index(jen.Id("ID")),
			jen.Id("ok"),
		).Block(
			jen.Return(jen.Op("&").Id("a")),
		),
		jen.Return(jen.Nil()),
	)
}
