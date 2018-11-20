package main

import (
	"encoding/xml"
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

	decoder := xml.NewDecoder(pom)
	token, err := decoder.Token()
	if err != nil {
		log.Println("failed to get token", err)
		return
	}
	switch tokenType := token.(type) {
	case xml.StartElement:
		var project Model
		err := project.UnmarshalXML(decoder, tokenType)
		if err != nil {
			log.Fatalln("failed to parse pom", err)
		}
		log.Println(project)
		for _, dep := range project.Dependencies.Dependency {
			log.Println(dep)
		}
	default:
		log.Println("type of token", reflect.TypeOf(tokenType), tokenType)
	}

	token, err = decoder.Token()
	if err != nil {
		log.Println("failed to get token", err)
		return
	}
	switch tokenType := token.(type) {
	case xml.StartElement:
		var project Model
		err := project.UnmarshalXML(decoder, tokenType)
		if err != nil {
			log.Fatalln("failed to parse pom", err)
		}
		log.Println(project)
		for _, dep := range project.Dependencies.Dependency {
			log.Println(dep)
		}
	default:
		log.Println("type of token", reflect.TypeOf(tokenType), tokenType)
	}

	token, err = decoder.Token()
	if err != nil {
		log.Println("failed to get token", err)
		return
	}
	switch tokenType := token.(type) {
	case xml.StartElement:
		var project Model
		err := project.UnmarshalXML(decoder, tokenType)
		if err != nil {
			log.Fatalln("failed to parse pom", err)
		}
		log.Println(project)
		for _, dep := range project.Dependencies.Dependency {
			log.Println(dep)
		}
	default:
		log.Println("type of token", reflect.TypeOf(tokenType), tokenType)
	}
}
