package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	path "path/filepath"
	"strings"
)

var (
	cmd     string
	appname string
)

func writetofile(filename, content string) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString(content)
}

func isExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func containsString(slice []string, element string) bool {
	for _, elem := range slice {
		if elem == element {
			return true
		}
	}
	return false
}

func askForConfirmation() bool {
	var response string
	_, err := fmt.Scanln(&response)
	if err != nil {
		log.Fatal(err)
	}
	okayResponses := []string{"y", "Y", "yes", "Yes", "YES"}
	nokayResponses := []string{"n", "N", "no", "No", "NO"}
	if containsString(okayResponses, response) {
		return true
	} else if containsString(nokayResponses, response) {
		return false
	} else {
		fmt.Println("Please type yes or no and then press enter:")
		return askForConfirmation()
	}
}

func createApp(name string) {
	curpath, _ := os.Getwd()
	log.Println("current path:", curpath)
	gopath := os.Getenv("GOPATH")
	log.Println("gopath:", gopath)
	if gopath == "" {
		log.Println("[ERRO] $GOPATH not found\n")
		log.Println("[HINT] Set $GOPATH in your environment vairables\n")
		os.Exit(2)
	}

	haspath := false
	appsrcpath := ""

	wgopath := path.SplitList(gopath)
	for _, wg := range wgopath {

		wg = path.Join(wg, "src")

		if strings.HasPrefix(strings.ToLower(curpath), strings.ToLower(wg)) {
			haspath = true
			appsrcpath = wg
			break
		}

		wg, _ = path.EvalSymlinks(wg)

		if strings.HasPrefix(strings.ToLower(curpath), strings.ToLower(wg)) {
			haspath = true
			appsrcpath = wg
			break
		}

	}

	if !haspath {
		log.Printf("[ERRO] Unable to create an application outside of $GOPATH%ssrc(%s%ssrc)\n", string(path.Separator), gopath, string(path.Separator))
		log.Printf("[HINT] Change your work directory by `cd ($GOPATH%ssrc)`\n", string(path.Separator))
		os.Exit(2)
	}

	apppath := path.Join(curpath, name)

	if isExist(apppath) {
		log.Printf("[ERRO] Path (%s) already exists\n", apppath)
		log.Printf("[WARN] Do you want to overwrite it? [yes|no]]")
		if !askForConfirmation() {
			os.Exit(2)
		}
	}
	log.Println("SRC PATH ", appsrcpath)

	log.Println("Creating application ...")

	os.MkdirAll(apppath, 0755)
	fmt.Println(apppath + string(path.Separator))
	os.Mkdir(path.Join(apppath, "forms"), 0755)
	fmt.Println(path.Join(apppath, "forms") + string(path.Separator))

}

func main() {

	flag.StringVar(&cmd, "cmd", "", "task: (create)")
	flag.StringVar(&appname, "app", "app", "name of app")

	flag.Parse()
	log.SetFlags(0)

	switch cmd {
	case "create":
		log.Println("Creating app ", appname)
		createApp(appname)
	default:
		flag.Usage()
	}
}
