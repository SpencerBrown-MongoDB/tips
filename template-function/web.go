package main

import (
	"bytes"
	"fmt"
	"log"
	"text/template"
)

func main() {
	tbare := template.New("t")
	ttxt := "some text in a template: {{Version}}"
	tdata := "SUBMEIN"
	fm := template.FuncMap{
		"Version": Version,
	}
	t := tbare.Funcs(fm)
	tp, err := t.Parse(ttxt)
	if err != nil {
		log.Fatalf("Error parsing template '%s': %v\n", ttxt, err)
	}
	var bb bytes.Buffer
	err = tp.Execute(&bb, tdata)
	if err != nil {
		log.Fatalf("Error executing template '%s' with data '%v': %v", ttxt, tdata, err)
	}
	fmt.Printf("Template: %s\nResult  : %s\n", ttxt, bb.Bytes())
}

func Version() string {
	return "0.5.5"
}
