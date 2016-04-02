package main

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"io"
	"log"
	"os"
	path "path/filepath"
	"strings"
)

var (
	cmd     string
	appname string
)

func copyFile(source string, dest string) (err error) {
	sourcefile, err := os.Open(source)
	if err != nil {
		return err
	}

	defer sourcefile.Close()

	destfile, err := os.Create(dest)
	if err != nil {
		return err
	}

	defer destfile.Close()

	_, err = io.Copy(destfile, sourcefile)
	if err == nil {
		sourceinfo, err := os.Stat(source)
		if err != nil {
			err = os.Chmod(dest, sourceinfo.Mode())
		}

	}

	return
}

func copyDir(source string, dest string) (err error) {

	// get properties of source dir
	sourceinfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	// create dest dir

	err = os.MkdirAll(dest, sourceinfo.Mode())
	if err != nil {
		return err
	}

	directory, _ := os.Open(source)

	objects, err := directory.Readdir(-1)

	for _, obj := range objects {

		sourcefilepointer := source + "/" + obj.Name()

		destinationfilepointer := dest + "/" + obj.Name()

		if obj.IsDir() {
			// create sub-directories - recursively
			err = copyDir(sourcefilepointer, destinationfilepointer)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			// perform copy
			err = copyFile(sourcefilepointer, destinationfilepointer)
			if err != nil {
				fmt.Println(err)
			}
		}

	}
	return
}

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
	pathToValente := path.Join(appsrcpath, "github.com", "trumae", "valente")
	pathToValenteData := path.Join(pathToValente, "data")
	pathToValenteDataPublic := path.Join(pathToValenteData, "public")
	pathToValenteDataForms := path.Join(pathToValenteData, "forms")

	log.Println("Creating application ...")

	os.MkdirAll(apppath, 0755)
	fmt.Println(apppath + string(path.Separator))

	copyDir(pathToValenteDataPublic, path.Join(apppath, "public"))
	copyDir(pathToValenteDataForms, path.Join(apppath, "forms"))

	packageForms := path.Join(apppath[len(appsrcpath)+1:], "forms")

	tmpl, err := template.New("forms").Parse(tplMainGo)
	if err != nil {
		panic(err)
	}

	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, packageForms)
	if err != nil {
		panic(err)
	}
	writetofile(path.Join(apppath, "main.go"), buf.String())

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

const tplMainGo = `package main

import (
   "log"
   "runtime"

   "github.com/labstack/echo"
   "github.com/labstack/echo/engine/standard"
   "github.com/labstack/echo/middleware"
   "github.com/trumae/valente"
	 "{{ . }}"

   "golang.org/x/net/websocket"
)

//App is a Web Application representation
type App struct {
    valente.App
}

//Initialize inits the App
func (app *App) Initialize() {
    log.Println("App Initialize")

    app.AddForm("login", forms.FormLogin{})
    app.AddForm("home", forms.FormHome{})

    app.GoTo("login")
}

func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())

    e := echo.New()

    e.Use(middleware.Logger())
    e.Use(middleware.Recover())
    e.Use(middleware.Static("public"))

    e.Get("/ws", standard.WrapHandler(websocket.Handler(func(ws *websocket.Conn) {
       app := App{}
       app.WS = ws
       app.Initialize()
       app.Run()
    })))

    e.Run(standard.New(":8000"))
}

`
