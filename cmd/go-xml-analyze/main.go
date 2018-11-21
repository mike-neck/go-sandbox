package main

import (
	"encoding/xml"
	"io"
	"log"
	"os"
	"reflect"
)

func main() {
	pom, err := os.Open("cmd/go-xml-analyze/testdata/junit-jupiter-api-5.0.0-M3.pom")
	if err != nil {
		log.Fatalln("failed to open pom file.", err)
	}
	defer pom.Close()

	err = run(pom)
	if err != nil {
		log.Println("error", err)
		return
	}
}

func run(pom io.Reader) error {
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
		log.Println(dep)
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
