package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dimiro1/banner"
	"github.com/mattn/go-colorable"
	"github.com/trumae/valente"
)

const textBanner = `
  ____  _        _             
 / ___|| |_ __ _| |_ _   _ ___ 
 \___ \| __/ _\ | __| | | / __|
  ___) | || (_| | |_| |_| \__ \
 |____/ \__\__,_|\__|\__,_|___/

GoVersion: {{ .GoVersion }}
GOOS: {{ .GOOS }}

`

func printStatus(status valente.StatusInfo) {
	fmt.Println("Started at:", status.Started)
	fmt.Println("Open Sessions:", status.OpenSessions)
	fmt.Println("Closed Sessions:", status.ClosedSessions)
	fmt.Println("Goto between forms:", status.Gotos)
}

func main() {
	isEnabled := true
	isColorEnabled := true
	banner.Init(colorable.NewColorableStdout(), isEnabled, isColorEnabled, bytes.NewBufferString(textBanner))

	resp, err := http.Get("http://localhost:8000/status")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	status := valente.StatusInfo{}
	json.Unmarshal(body, &status)

	printStatus(status)
}
