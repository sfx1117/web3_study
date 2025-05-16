package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	t1 := template.New("t1")
	t1, err := t1.Parse("Value is {{.}}\n")
	if err != nil {
		log.Fatal(err)
	}
	t1 = template.Must(t1.Parse("Value is {{.}}\n"))
	t1.Execute(os.Stdout, "some text")
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string{
		"Go",
		"Rust",
		"C++",
		"C#",
	})

	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}
	t2 := Create("t2", "Name: {{.Name}}\n")
	t2.Execute(os.Stdout, struct {
		Name string
	}{"sfx"})
	t2.Execute(os.Stdout, map[string]string{
		"Name": "Mickey Mouse",
	})

	t3 := Create("t3", "{{if . -}} yes {{else -}} no {{end}}\n")
	//t3 := Create("t3", "{{if . -}} yes {{else -}} no {{end}}\n")
	t3.Execute(os.Stdout, "not empty")
	t3.Execute(os.Stdout, "")

	t4 := Create("t4", "Range: {{range .}}{{.}} {{end}}\n")
	t4.Execute(os.Stdout, []string{"Go",
		"Rust",
		"C++",
		"C#"})
}

//Value is some text
//Value is 5
//Value is [Go Rust C++ C#]
//Name: sfx
//Name: Mickey Mouse
//yes
//no
//Range: Go Rust C++ C#
