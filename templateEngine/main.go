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
	inputName := []string{"input"}
	tmplateURL := GetTemplateURL(purseFinishedURL, inputName)
	replaceValues := map[string]string{}
	for _, v := range []string{"hoge", "fuga", "piyo"} {
		replaceValues["input"] = v
		fmt.Println(ReplaceTemplateURL(tmplateURL, replaceValues))
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
