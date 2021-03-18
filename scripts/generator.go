/*
 *    Copyright 2021 FerociousBite and Contributors
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

// +build ignore
package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/davecgh/go-spew/spew"
)

type Parsed struct {
	Structs []string
	Missing []string
	Times   []struct {
		Struct string
		Field  string
	}
}

func main() {
	tmpl := template.Must(template.New("cases").Parse(`
// Code generated by github.com/ferocious-space/eveapi/scripts/generator.go ; DO NOT EDIT.
package notification

import (
	"fmt"
	"time"
	"strings"
	"errors"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-openapi/swag"
	"gopkg.in/yaml.v2"

	"github.com/ferocious-space/eveapi/esi/character"
)

func ParseNotification(n *character.GetCharactersCharacterIDNotificationsOKBodyItems0) (interface{},error) {
	switch strings.TrimSpace(swag.StringValue(n.Type)) {
	{{range $decl := .Structs }}
	case "{{$decl}}":
		value := new({{$decl}})
		err := yaml.Unmarshal([]byte(n.Text),&value)
		if err != nil {
			spew.Dump(n)
			return nil,err
		}
		return value,nil
	{{end}}
	// Missing implementations
	{{range $decl := .Missing }}
	case "{{$decl}}":
		spew.Dump(n)
		bytes, _ := yaml.Marshal(n)
		fmt.Println(string(bytes))
		value := new({{$decl}})
		err := yaml.Unmarshal([]byte(n.Text),&value)
		if err != nil {
			spew.Dump(n)
			return nil,err
		}
		return value,nil
	{{end}}
	}
	return nil,errors.New("not implemented")
}
{{range $arg := .Missing}}
type {{$arg}} interface{}
{{end}}
{{range $agr := .Times }}
func (s *{{$agr.Struct}}) Get{{$agr.Field}}() time.Time {
	return TimeFromCCPTimestamp(s.{{$agr.Field}})
}
{{end}}
`))
	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, "notifications/notification.go", nil, parser.AllErrors)
	if err != nil {
		spew.Dump(err)
		return
	}
	var p Parsed
	for _, d := range f.Decls {
		switch d.(type) {
		case *ast.GenDecl:
			typeDecl := d.(*ast.GenDecl)
			for _, s := range typeDecl.Specs {
				switch s.(type) {
				case *ast.TypeSpec:
					structDecl := s.(*ast.TypeSpec).Name.Name
					p.Structs = append(p.Structs, structDecl)
					switch s.(*ast.TypeSpec).Type.(type) {
					case *ast.StructType:
						fields := s.(*ast.TypeSpec).Type.(*ast.StructType).Fields.List
						for _, f := range fields {
							if strings.Contains(f.Names[0].Name, "Time") {
								p.Times = append(p.Times, struct {
									Struct string
									Field  string
								}{Struct: structDecl, Field: f.Names[0].Name})
							}
						}
					}
				}
			}
		}
	}
	fi, err := os.Create("notifications/parser.go")
	if err != nil {
		log.Fatal(err)
	}
	defer fi.Close()

	fs2 := token.NewFileSet()
	f2, err := parser.ParseFile(fs2, "esi/character/get_characters_character_id_notifications_responses.go", nil, parser.AllErrors)
	if err != nil {
		spew.Dump(err)
		return
	}

	for _, k := range f2.Decls {
		switch k.(type) {
		case *ast.GenDecl:
			typeDecl := k.(*ast.GenDecl)
			for _, s := range typeDecl.Specs {
				switch s.(type) {
				case *ast.ValueSpec:
					structDecl := s.(*ast.ValueSpec).Names[0].Name
					if strings.Contains(structDecl, "GetCharactersCharacterIDNotificationsOKBodyItems0Type") {
						enu := strings.TrimPrefix(structDecl, "GetCharactersCharacterIDNotificationsOKBodyItems0Type")
						valid := false
						for e := range p.Structs {
							if p.Structs[e] == enu {
								valid = true
								break
							}
						}
						if !valid {
							p.Missing = append(p.Missing, enu)
							println("missing ", enu)
						}
					}
				}
			}
		}
	}

	_ = tmpl.Execute(fi, p)
}
