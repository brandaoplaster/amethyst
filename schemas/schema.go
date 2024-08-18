package schemas

import (
	"fmt"
	"reflect"
	"time"
)

type SchemaType struct {
	Name   string
	GoType reflect.Type
}

var (
	String  = SchemaType{Name: "string", GoType: reflect.TypeOf("")}
	Int     = SchemaType{Name: "int", GoType: reflect.TypeOf(0)}
	Float64 = SchemaType{Name: "float64", GoType: reflect.TypeOf(float64(0))}
	Bool    = SchemaType{Name: "bool", GoType: reflect.TypeOf(true)}
	Time    = SchemaType{Name: "time.Time", GoType: reflect.TypeOf(time.Time{})}
)

type Field struct {
	Name string
	Type SchemaType
}

type Schema struct {
	Name   string
	Fields []Field
}

func (s *Schema) AddField(name string, fieldType SchemaType) {
	s.Fields = append(s.Fields, Field{Name: name, Type: fieldType})
}

func (s *Schema) GenerateSchema() string {
	var result string
	result += fmt.Sprintf("type %s struct {\n", s.Name)
	for _, field := range s.Fields {
		result += fmt.Sprintf("\t%s %s\n", field.Name, field.Type.Name)
	}
	result += "}\n"
	return result
}
