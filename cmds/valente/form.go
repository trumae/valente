package main

import (
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

}
