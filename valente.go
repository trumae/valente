package valente

import (
	"errors"
	"log"
	"time"

	"github.com/trumae/valente/action"
	"github.com/trumae/valente/status"

	"golang.org/x/net/websocket"
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
type HandlerFunc func(*websocket.Conn, *App, []string)

//Form represents the unit of user interaction
type Form interface {
	AddEventHandler(evt string, f HandlerFunc) Form
	Run(ws *websocket.Conn, app *App) error
	Initialize(ws *websocket.Conn) Form
	Render(ws *websocket.Conn, app *App, params []string) error
}

//FormImpl its a simple Form
type FormImpl struct {
	trans map[string]HandlerFunc
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
func (form FormImpl) Run(ws *websocket.Conn, app *App) error {
	msgs := []string{}
	for {
		msg := ""
		err := websocket.Message.Receive(ws, &msg)
		if err != nil {
			log.Println("Error on WS Receive", err)
			return ErrEOFWS
		}
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
func (form FormImpl) Initialize(ws *websocket.Conn) Form {
	log.Println("FormImpl Initialize")
	return form
}

//Render form start
func (form FormImpl) Render(ws *websocket.Conn, app *App, params []string) error {
	log.Println("FormImpl Render")
	return nil
}

//App is a Web Application representation
type App struct {
	WS          *websocket.Conn
	Forms       map[string]Form
	Data        map[string]interface{}
	CurrentForm Form
	LastAccess  time.Time
}

//WebSocket set the WS value
func (app *App) WebSocket(ws *websocket.Conn) {
	app.WS = ws
	PutWS(ws)
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
	app.Data = map[string]interface{}{}
	status.Status.OpenSessions++
	go func() {
		c := time.Tick(10 * time.Second)
		for range c {
			err := action.Exec(app.WS, "1 == 1;")
			if err != nil {
				log.Println("Error in connection probe", err)
				status.Status.ClosedSessions++
				DropWS(app.WS)
				return
			}
			app.LastAccess = time.Now()
		}
	}()

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

func init() {
	status.Status.Started = time.Now()
	tablews = make([]*websocket.Conn, 0, 100)
	wschannel = make(chan wsmessage)
	go tableWSServer()
}
