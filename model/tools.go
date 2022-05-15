package model

import (
	"bytes"
	"io/ioutil"
	"log"
	"os/exec"
	"reflect"
	"strings"
	"text/template"

	"github.com/pkg/errors"
)

type ObjectProto struct {
	ObjectName string
	InfoProxy  interface{}
	OutputFile string
	Simple     bool
}

func CreateObject(proto ObjectProto) error {
	t := reflect.TypeOf(proto.InfoProxy)
	type field struct {
		Name, Type string
	}
	type data struct {
		ObjectName string
		ObjectVar  string
		Fields     []field
	}
	d := data{
		ObjectName: proto.ObjectName,
		ObjectVar:  strings.ToLower(string(proto.ObjectName[0])),
	}
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		name, typ := f.Name, f.Type.String()
		d.Fields = append(d.Fields, field{
			Name: name,
			Type: typ,
		})
	}
	var tmpl string
	if proto.Simple {
		tmpl = `
	package model

	{{$objectName := .ObjectName}}
	{{$objectVar := .ObjectVar}}

	{{range .Fields}}
	// {{.Name}} returns the {{.Type}} {{.Name}} of this {{$objectName}}
	func ({{$objectVar}} *{{$objectName}}) {{.Name}}() {{.Type}} {
		return {{$objectVar}}.info.{{.Name}}
	}
	{{end}}

	func ({{$objectVar}} *{{$objectName}}) String() string {
		var buf bytes.Buffer
		buf.WriteString("{{$objectName}}(")
		{{- range $i, $v := .Fields }}
			{{ if eq $i 0}}buf.WriteString("{{$v.Name}}="){{ else }}buf.WriteString(", {{$v.Name}}="){{ end }}
			{{ if eq .Type "string" }}buf.WriteString({{$objectVar}}.{{$v.Name}}()){{ else }}buf.WriteString(fmt.Sprintf("%v", {{$objectVar}}.{{$v.Name}}())){{ end }}
		{{- end }}
		buf.WriteString(" )")
		return buf.String()
	}	
	`
	} else {
		tmpl = `
	package model

	{{$objectName := .ObjectName}}
	{{$objectVar := .ObjectVar}}

	{{range .Fields}}
	// {{.Name}} returns the {{.Type}} {{.Name}} of this {{$objectName}}
	func ({{$objectVar}} *{{$objectName}}) {{.Name}}() ({{.Type}}, error) {
		info, err := {{$objectVar}}.getInfo()
		if err != nil {
			return "", err
		}
		return info.{{.Name}}, nil
	}

	// {{.Name}} returns {{.Name}}() or panics if there is an error
	func ({{$objectVar}} *{{$objectName}}) {{.Name}}OrPanic() {{.Type}} {
		res, err := {{$objectVar}}.{{.Name}}()
		if err != nil {
			panic(fmt.Sprintf("error calling {{.Name}}(): %v", err))
		}
		return res
	}
	{{end}}

	func ({{$objectVar}} *{{$objectName}}) String() string {
		var buf bytes.Buffer
		buf.WriteString("{{$objectName}}(")
		{{- range $i, $v := .Fields }}
			{{ if eq $i 0}}buf.WriteString("{{$v.Name}}="){{ else }}buf.WriteString(", {{$v.Name}}="){{ end }}
			{{ if eq .Type "string" }}buf.WriteString({{$objectVar}}.{{$v.Name}}OrPanic()){{ else }}buf.WriteString(fmt.Sprintf("%v", {{$objectVar}}.{{$v.Name}}OrPanic())){{ end }}
		{{- end }}
		buf.WriteString(")")
		return buf.String()
	}	
	`
	}

	output, err := renderTemplate(tmpl, proto.ObjectName, d)
	if err != nil {
		return err
	}
	outfile := proto.OutputFile
	if err := ioutil.WriteFile(outfile, output, 0755); err != nil {
		return err
	}
	if err := exec.Command("go", "fmt", outfile).Run(); err != nil {
		return errors.Errorf("formatting main code to %s: %v", outfile, err)
	}
	if err := exec.Command("goimports", "-w", outfile).Run(); err != nil {
		return errors.Errorf("fixing imports for main code in %s: %v", outfile, err)
	}
	log.Printf("wrote to %s", outfile)
	return nil
}

func renderTemplate(t string, name string, data interface{}) ([]byte, error) {
	tmpl, err := template.New(name).Parse(strings.TrimSpace(t))
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
