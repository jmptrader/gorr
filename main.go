package main

import (
	"net/http"
	"strings"
	"encoding/json"
	"io/ioutil"
	"html/template"
	"bytes"
	"fmt"
	"log"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/views/")
	jsonParam := r.FormValue("params")
	params := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonParam), &params)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	content, err := Render(path, params)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(content)
}

func Render(path string, params interface{}) ([]byte, error) {
	content, err := ioutil.ReadFile("views/" + path)
	if err != nil {
		return nil, err
	}
	compiledTemplate, err := template.New("template").Parse(string(content))
	if err != nil {
		return nil, err
	}
	out := bytes.Buffer{}
	err = compiledTemplate.Execute(&out, params)
	if err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}

func main() {
	http.HandleFunc("/views/", viewHandler)
	log.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
