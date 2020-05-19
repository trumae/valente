package valente

import (
	"errors"
	"log"
	"time"

	"github.com/trumae/valente/action"
	"github.com/trumae/valente/status"
)

const (
	endofmessage = "___ENDOFMESSAGE___"
)

var (
	//ErrProtocol is an error of WS Valente protocol
	ErrProtocol = errors.New("Protocol Error")

	//ErrEOFWS end of file of WS
	ErrEOFWS = errors.New("EOF Error")
)

//HandlerFunc is a function of handle an event received into websocket.Conn
type HandlerFunc func(*action.WebSocket, *App, []string)

//Form represents the unit of user interaction
type Form interface {
	Name() string
	AddEventHandler(evt string, f HandlerFunc) Form
	Run(ws *action.WebSocket, app *App) error
	Initialize(ws *action.WebSocket) Form
	Render(ws *action.WebSocket, app *App, params []string) error
}

//FormImpl its a simple Form
type FormImpl struct {
	FName string
	trans map[string]HandlerFunc
}

//Name return the name string
func (form FormImpl) Name() string {
	return form.FName
}

//AddEventHandler add an f function to handle evt event
func (form FormImpl) AddEventHandler(evt string, f HandlerFunc) Form {
	if form.trans == nil {
		form.trans = map[string]HandlerFunc{}
	}
	form.trans[evt] = f
	return form
}

//Run execs the form
func (form FormImpl) Run(ws *action.WebSocket, app *App) error {
	msgs := []string{}
	for {
		_, bmsg, err := ws.WS.ReadMessage()
		if err != nil {
			log.Println("Error on WS Receive", err)
			return ErrEOFWS
		}
		msg := string(bmsg)
		status.Status.ReceivedBytes += len(msg)
		if msg == endofmessage {
			break
		} else {
			msgs = append(msgs, msg)
		}
	}
	log.Printf("currentForm = %v msgs = %v\n", app.CurrentForm, msgs)
	if len(msgs) < 1 {
		log.Println("Protocol error len(msgs) < 1", msgs)
		return ErrProtocol
	}

	app.LastAccess = time.Now()
	f, present := form.trans[msgs[0]]
	if present {
		f(ws, app, msgs)
	} else {
		log.Println("Evt not found", msgs[0])
	}
	return nil
}

//Initialize inits the form
func (form FormImpl) Initialize(ws *action.WebSocket) Form {
	log.Println("FormImpl Initialize")
	return form
}

//Render form start
func (form FormImpl) Render(ws *action.WebSocket, app *App, params []string) error {
	log.Println("FormImpl Render")
	return nil
}

//App is a Web Application representation
type App struct {
	WS          *action.WebSocket
	Forms       map[string]Form
	Data        map[string]interface{}
	CurrentForm Form
	LastAccess  time.Time
}

//WebSocket set the WS value
func (app *App) WebSocket(ws *action.WebSocket) {
	app.WS = ws
}

//GoTo replace the current form into app
func (app *App) GoTo(formName string, params []string) error {
	log.Println("App goto", formName)
	form, present := app.Forms[formName]
	if present {
		app.CurrentForm = form.Initialize(app.WS)
		err := form.Render(app.WS, app, params)
		if err != nil {
			log.Println("Error on goto", err)
		}
		status.Status.Gotos++
		action.UnblockUI(app.WS)
	} else {
		log.Println("[ERROR] Form not registred", formName)
	}
	return nil
}

//Run handle events
func (app *App) Run() {
	if app.Data == nil {
		app.Data = map[string]interface{}{}
	}
	status.Status.OpenSessions++

	for {
		err := app.CurrentForm.Run(app.WS, app)
		if err != nil {
			log.Println(err)
			switch {
			case err == ErrProtocol:
				continue
			case err == ErrEOFWS:
				return
			default:
				continue
			}
		}
	}
}

//Initialize inits the App
func (app *App) Initialize() {
	log.Println("App Initialize")
}

//AddForm add a new form to App
func (app *App) AddForm(name string, f Form) {
	log.Println("AddForm", name, f)
	if app.Forms == nil {
		app.Forms = map[string]Form{}
	}

	app.Forms[name] = f
}
