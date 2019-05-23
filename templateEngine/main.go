package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/url"
	"strings"
)

func main() {
	rowURL := "http://localhost"
	purseFinishedURL, _ := url.Parse(rowURL)
	inputName := []string{"input", "text", "line"}
	templateURL := GetTemplateURL(purseFinishedURL, inputName)
	data := []string{"hoge", "fuga", "piyo"}
	ReplaceOnSimpleList(templateURL, inputName, data)

}
func ReplaceOnSimpleList(templateURL string, inputName []string, data []string) {
	replaceValues := map[string]string{}
	for _, name := range inputName {
		defaultValue := replaceValues[name]
		for _, v := range data {
			replaceValues[name] = v
			fmt.Println(ReplaceTemplateURL(templateURL, replaceValues))
		}
		replaceValues[name] = defaultValue
	}
}

func ReplaceOnAllSimpleList(templateURL string, inputName []string, data []string) {
	replaceValues := map[string]string{}
	for _, v := range data {
		for _, name := range inputName {
			replaceValues[name] = v
		}
		fmt.Println(ReplaceTemplateURL(templateURL, replaceValues))
	}
}

func GetTemplateURL(u *url.URL, name []string) string {
	var data []string
	for _, dataName := range name {
		data = append(data, dataName+"={{."+dataName+"}}")
	}
	q := strings.Join(data, "&")
	if q != "" {
		return u.String() + "?" + q
	}
	return u.String()
}

func ReplaceTemplateURL(urlTmplate string, data map[string]string) string {
	tmpl, err := template.New("test").Parse(urlTmplate)
	if err != nil {
		panic(err)
	}
	var doc bytes.Buffer
	if err := tmpl.Execute(&doc, data); err != nil {
		panic(err)
	}
	return doc.String()
}
