package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	path "path/filepath"
	"strings"
)

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

    app.GoTo("login", nil)
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

		log.Println("Server running")
    e.Run(standard.New(":8000"))
}

`

func createApp(name string) {
	curpath, _ := os.Getwd()
	log.Println("current path:", curpath)
	gopath := os.Getenv("GOPATH")
	log.Println("gopath:", gopath)
	if gopath == "" {
		log.Println("[ERRO] $GOPATH not found")
		log.Println("[HINT] Set $GOPATH in your environment vairables")
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
