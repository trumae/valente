package main

import (
	"bytes"
	"html/template"
	"log"
	"os"
	path "path/filepath"
)

func createForm(form string) {
	curpath, _ := os.Getwd()
	log.Println("current path:", curpath)

	formspath := path.Join(curpath, "forms")

	if !isExist(formspath) {
		log.Printf("[ERRO] Path (%s) not exists\n", formspath)
		os.Exit(2)
	}

	tmpl, err := template.New("form").Parse(tplForm)
	if err != nil {
		panic(err)
	}

	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, form)
	if err != nil {
		panic(err)
	}

	log.Println("Writing Form", form)
	writetofile(path.Join(formspath, form+".go"), buf.String())
}

const tplForm = `package forms

import (
  "log"
  "github.com/trumae/valente"
  "github.com/trumae/valente/action"
  "golang.org/x/net/websocket"
)

//{{.}}Form struct
type {{.}}Form struct {
    valente.FormImpl
}

//Render the initial html form to Form{{.}}
func (form {{.}}Form) Render(ws *websocket.Conn, app *valente.App, params []string) error {
	action.Alert(ws, "Render of {{.}}Form not implemented")
	return nil
}

//Initialize inits the {{.}} Form
func (form {{.}}Form) Initialize(ws *websocket.Conn) valente.Form {
  log.Println("{{.}}Form Initialize")
  return form
}

`
