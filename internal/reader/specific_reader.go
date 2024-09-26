package reader

import (
	"Practics_with_templates/internal/random"
	"bytes"
	"math/rand"
	"text/template"
	"time"
)

func Calculator(number1, number2 int) int {
	return number1 + number2
}

func FuncReader(filename string) (string, error) {
	var buf bytes.Buffer

	funcMap := template.FuncMap{
		"Now": time.Now,
		"f1":  func(param1, param2 string) int { return Calculator(rand.Intn(799), rand.Intn(799)) },
	}
	tpl, err := template.New("another_test_template").Funcs(funcMap).ParseFiles(filename)

	if err != nil {
		return "", err
	}
	data := make(map[string]string)
	data["client"] = random.RandomName()
	err = tpl.Execute(&buf, data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
