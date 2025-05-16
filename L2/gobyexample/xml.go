package main

import (
	"encoding/xml"
	"fmt"
)

type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"Origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v",
		p.Id, p.Name, p.Origin)
}

func main() {

	coff := &Plant{Id: 25, Name: "sfx"}
	coff.Origin = []string{"aaa", "bbb", "ccc"}

	out, _ := xml.MarshalIndent(coff, "", "")
	fmt.Println(string(out))
	fmt.Println(xml.Header + string(out))

	var p Plant
	err := xml.Unmarshal(out, &p)
	if err != nil {
		panic(err)
	}
	fmt.Println(p)

	tomato := &Plant{Id: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}
	type nesting struct {
		XMLNme xml.Name `xml:"nesting"`
		Plants []*Plant `xml:"parent>child>plant"`
	}
	nestings := &nesting{}
	nestings.Plants = []*Plant{coff, tomato}
	indent, _ := xml.MarshalIndent(nestings, "", "")
	fmt.Println(string(indent))
}

//<plant id="25"><name>sfx</name><Origin>aaa</Origin><Origin>bbb</Origin><Origin>ccc</Origin></plant>
//<?xml version="1.0" encoding="UTF-8"?>
//<plant id="25"><name>sfx</name><Origin>aaa</Origin><Origin>bbb</Origin><Origin>ccc</Origin></plant>
//Plant id=25, name=sfx, origin=[aaa bbb ccc]
//<nesting><nesting></nesting><parent><child><plant id="25"><name>sfx</name><Origin>aaa</Origin><Origin>bbb</Origin><Origin>ccc</Origin></plant><plant id="81"><name>Tomato</name><Origin>Mexico</Origin><Origin>California</Origin></plant></child></parent></nesting>
