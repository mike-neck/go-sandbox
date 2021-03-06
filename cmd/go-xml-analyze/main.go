package main

import (
	"encoding/xml"
	"log"
	"os"
	"reflect"
)

func main() {
	err := run("cmd/go-xml-analyze/testdata/junit-jupiter-api-5.0.0-M3.pom")
	if err != nil {
		log.Fatalln("error", err)
	}
}

func run(file string) error {
	log.Println("parse file", file)
	pom, err := os.Open(file)
	if err != nil {
		log.Println("failed to open pom file.", err)
		return err
	}
	defer pom.Close()

	decoder := xml.NewDecoder(pom)

	for {
		elem, err := read(decoder)
		if err != nil {
			return err
		}
		if elem.show() {
			return nil
		}
	}
}

type Elem interface {
	show() bool
}

func read(decoder *xml.Decoder) (Elem, error) {
	token, err := decoder.Token()
	if err != nil {
		log.Println("failed to get token", err)
		return nil, err
	}
	switch tokenType := token.(type) {
	case xml.StartElement:
		var project Model
		err := project.UnmarshalXML(decoder, tokenType)
		if err != nil {
			log.Println("failed to parse pom", err)
			return nil, err
		}
		return project, nil
	default:
		return &Another{token: tokenType}, nil
	}
}

func (m Model) show() bool {
	log.Println("parsing pom is succeeded.")
	log.Println(m)
	for _, dep := range m.Dependencies.Dependency {
		log.Println("group:", dep.GroupId, "artifact:", dep.ArtifactId, "version:", dep.Version, "scope:", dep.Scope)
	}
	return true
}

type Another struct {
	token xml.Token
}

func (a *Another) show() bool {
	log.Println("type of token", reflect.TypeOf(a.token), a.token)
	return false
}
