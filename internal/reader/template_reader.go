package reader

import (
	"Practics_with_templates/internal/random"
	"bytes"
	"math/rand"
	"strconv"
	"text/template"
)

type GenericInterfaceStruct interface {
	Strings() string
}

type TestTemplateStruct struct {
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	Salary     int    `json:"salary"`
}

func (t TestTemplateStruct) Strings() string {
	return t.Name + " " + t.Occupation + " " + strconv.Itoa(t.Salary)
}

func DefaultTestTemplate() TestTemplateStruct {
	return TestTemplateStruct{
		Name:       random.RandomName(),
		Occupation: random.RandomOccupation(),
		Salary:     rand.Intn(1500),
	}
}

func Read(filename string, data GenericInterfaceStruct) (string, error) {
	var buf bytes.Buffer
	tpl, err := template.ParseFiles(filename)
	if err != nil {
		return "", err
	}

	err = tpl.Execute(&buf, data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
