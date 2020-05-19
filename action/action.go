package action

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/trumae/valente/status"
)

var (
	errNotImplemented = errors.New("Not Implemented!")
)

type WebSocket struct {
	WS    *websocket.Conn
	Mutex sync.Mutex
}

// Exec execute the js code on WebBrowser
func Exec(ws *WebSocket, js string) error {
	ws.Mutex.Lock()
	defer ws.Mutex.Unlock()

	err := ws.WS.WriteMessage(websocket.TextMessage, []byte(js))
	if err != nil {
		return err
	}
	status.Status.SendedBytes += len(js)
	return err
}

//Enable the target form field or button.
func Enable(ws *WebSocket, target string) error {
	return errNotImplemented
}

//Disable  the target form field or button.
func Disable(ws *WebSocket, target string) error {
	return errNotImplemented
}

//Replace target with new content
func Replace(ws *WebSocket, target, content string) error {
	ws.Mutex.Lock()
	defer ws.Mutex.Unlock()

	c := strings.Replace(content, "\n", "\\n", -1)
	c = strings.Replace(c, "\"", "\\\"", -1)
	js := fmt.Sprintf("$( \"#%s\" ).replaceWith(\"%s\");", target, c)

	err := ws.WS.WriteMessage(websocket.TextMessage, []byte(js))
	if err != nil {
		return err
	}
	status.Status.SendedBytes += len(js)
	return err
}

//HTML replace target with new content
func HTML(ws *WebSocket, target, content string) error {
	ws.Mutex.Lock()
	defer ws.Mutex.Unlock()

	c := strings.Replace(content, "\n", "\\n", -1)
	c = strings.Replace(c, "\"", "\\\"", -1)
	js := fmt.Sprintf("$( \"#%s\" ).html(\"%s\");", target, c)

	err := ws.WS.WriteMessage(websocket.TextMessage, []byte(js))
	if err != nil {
		return err
	}
	status.Status.SendedBytes += len(js)
	return err
}

//Hide the target
func Hide(ws *WebSocket, target string, duration string) error {
	ws.Mutex.Lock()
	defer ws.Mutex.Unlock()

	js := fmt.Sprintf("$( \"#%s\" ).hide(\"%s\");", target, duration)

	err := ws.WS.WriteMessage(websocket.TextMessage, []byte(js))
	if err != nil {
		return err
	}
	status.Status.SendedBytes += len(js)
	return err
}

//Show the target
func Show(ws *WebSocket, target string, duration string) error {
	ws.Mutex.Lock()
	defer ws.Mutex.Unlock()

	js := fmt.Sprintf("$( \"#%s\" ).show(\"%s\");", target, duration)

	err := ws.WS.WriteMessage(websocket.TextMessage, []byte(js))
	if err != nil {
		return err
	}
	status.Status.SendedBytes += len(js)
	return err
}

//Remove target from the DOM
func Remove(ws *WebSocket, target string) error {
	return errNotImplemented
}

//Append concate content at target
func Append(ws *WebSocket, target, content string) error {
	ws.Mutex.Lock()
	defer ws.Mutex.Unlock()

	c := strings.Replace(content, "\n", "\\n", -1)
	c = strings.Replace(c, "\"", "\\\"", -1)
	js := fmt.Sprintf("$( \"#%s\" ).append(\"%s\");", target, c)

	err := ws.WS.WriteMessage(websocket.TextMessage, []byte(js))
	if err != nil {
		return err
	}
	status.Status.SendedBytes += len(js)
	return err
}

//Prepend concate content at the begin of target
func Prepend(ws *WebSocket, target, content string) error {
	ws.Mutex.Lock()
	defer ws.Mutex.Unlock()

	c := strings.Replace(content, "\n", "\\n", -1)
	c = strings.Replace(c, "\"", "\\\"", -1)
	js := fmt.Sprintf("$( \"#%s\" ).prepend(\"%s\");", target, c)

	err := ws.WS.WriteMessage(websocket.TextMessage, []byte(js))
	if err != nil {
		return err
	}
	status.Status.SendedBytes += len(js)
	return err
}

//Redirect to url
func Redirect(ws *WebSocket, url string) error {
	return errNotImplemented
}

//AddClass add a class for an element
func AddClass(ws *WebSocket, target, class string) error {
	ws.Mutex.Lock()
	defer ws.Mutex.Unlock()

	js := fmt.Sprintf("$('#%s').addClass('%s');", target, class)
	err := ws.WS.WriteMessage(websocket.TextMessage, []byte(js))
	if err != nil {
		return err
	}
	status.Status.SendedBytes += len(js)
	return nil
}

//RemoveClass add a class for an element
func RemoveClass(ws *WebSocket, target, class string) error {
	ws.Mutex.Lock()
	defer ws.Mutex.Unlock()

	js := fmt.Sprintf("$('#%s').removeClass('%s');", target, class)
	err := ws.WS.WriteMessage(websocket.TextMessage, []byte(js))
	if err != nil {
		return err
	}
	status.Status.SendedBytes += len(js)
	return nil
}

//Set update a form element (textbox, dropdown, checkbox, etc) to set text value of TargetID.
func Set(ws *WebSocket, target, value string) error {
	ws.Mutex.Lock()
	defer ws.Mutex.Unlock()

	c := strings.Replace(value, "\n", "\\n", -1)
	c = strings.Replace(c, "\"", "\\\"", -1)
	js := fmt.Sprintf("$('#%s').val('%s');", target, c)
	err := ws.WS.WriteMessage(websocket.TextMessage, []byte(js))
	if err != nil {
		return err
	}
	status.Status.SendedBytes += len(js)
	return nil
}

//Get content of form element
func Get(ws *WebSocket, target string) (string, error) {
	ws.Mutex.Lock()
	defer ws.Mutex.Unlock()

	ret := ""
	js := fmt.Sprintf("ws.send($('#%s').val());", target)
	status.Status.SendedBytes += len(js)
	err := ws.WS.WriteMessage(websocket.TextMessage, []byte(js))
	if err != nil {
		return "", err
	}

	_, bret, err := ws.WS.ReadMessage()
	if err != nil {
		return "", err
	}
	ret = string(bret)
	status.Status.ReceivedBytes += len(ret)

	return ret, nil
}

//Wire bind an action to an event on target
func Wire(ws *WebSocket, target, event, act string) error {
	return errNotImplemented
}

//SendEvent send an event to server
func SendEvent(ws *WebSocket, event string) error {
	ws.Mutex.Lock()
	defer ws.Mutex.Unlock()

	js := fmt.Sprintf("sendEvent('%s');", event)
	err := ws.WS.WriteMessage(websocket.TextMessage, []byte(js))
	if err != nil {
		return err
	}
	status.Status.SendedBytes += len(js)

	return nil
}

//Alert show alert message in browser
func Alert(ws *WebSocket, message string) error {
	ws.Mutex.Lock()
	defer ws.Mutex.Unlock()

	js := fmt.Sprintf("alert('%s');", message)
	err := ws.WS.WriteMessage(websocket.TextMessage, []byte(js))
	if err != nil {
		return err
	}
	status.Status.SendedBytes += len(js)

	return nil
}

var (
	//BlockMessage is the message showed in waiting time
	BlockMessage = "<h2>Please, wait...</h2>"
)

//BlockUI block page interaction
func BlockUI(ws *WebSocket) {
	Exec(ws, fmt.Sprintf("$.blockUI({ message: '%s' });", BlockMessage))
}

//UnblockUI block page interaction
func UnblockUI(ws *WebSocket) {
	Exec(ws, "$.unblockUI();")
}
