package action

import (
	"errors"
	"fmt"
	"strings"

	"golang.org/x/net/websocket"
)

var (
	errNotImplemented = errors.New("Not Implemented!")
)

// Exec execute the js code on WebBrowser
func Exec(ws *websocket.Conn, js string) error {
	return websocket.Message.Send(ws, js)
}

//Enable the target form field or button.
func Enable(ws *websocket.Conn, target string) error {
	return errNotImplemented
}

//Disable  the target form field or button.
func Disable(ws *websocket.Conn, target string) error {
	return errNotImplemented
}

//Replace target with new content
func Replace(ws *websocket.Conn, target, content string) error {
	c := strings.Replace(content, "\n", "\\n", -1)
	c = strings.Replace(c, "\"", "\\\"", -1)
	js := fmt.Sprintf("$( \"#%s\" ).replaceWith(\"%s\");", target, c)
	return websocket.Message.Send(ws, js)
}

//HTML replace target with new content
func HTML(ws *websocket.Conn, target, content string) error {
	c := strings.Replace(content, "\n", "\\n", -1)
	c = strings.Replace(c, "\"", "\\\"", -1)
	js := fmt.Sprintf("$( \"#%s\" ).html(\"%s\");", target, c)
	return websocket.Message.Send(ws, js)
}

//Remove target from the DOM
func Remove(ws *websocket.Conn, target string) error {
	return errNotImplemented
}

//InsertTop Insert content at the top of target
func InsertTop(ws *websocket.Conn, target, content string) error {
	return errNotImplemented
}

//InsertBottom Insert content at the bottom of target
func InsertBottom(ws *websocket.Conn, target, content string) error {
	return errNotImplemented
}

//InsertBefore Insert content at the before of target
func InsertBefore(ws *websocket.Conn, target, content string) error {
	return errNotImplemented
}

//InsertAfter Insert content at the after of target
func InsertAfter(ws *websocket.Conn, target, content string) error {
	return errNotImplemented
}

//Redirect to url
func Redirect(ws *websocket.Conn, url string) error {
	return errNotImplemented
}

//Set update a form element (textbox, dropdown, checkbox, etc) to set text value of TargetID.
func Set(ws *websocket.Conn, target, value string) error {
	js := fmt.Sprintf("$('#%s').val('%s');", target, value)
	err := websocket.Message.Send(ws, js)
	if err != nil {
		return err
	}
	return nil
}

//Get content of form element
func Get(ws *websocket.Conn, target string) (string, error) {
	ret := ""
	js := fmt.Sprintf("ws.send($('#%s').val());", target)
	err := websocket.Message.Send(ws, js)
	if err != nil {
		return "", err
	}

	err = websocket.Message.Receive(ws, &ret)
	if err != nil {
		return "", err
	}

	return ret, nil
}

//Wire bind an action to an event on target
func Wire(ws *websocket.Conn, target, event, act string) error {
	return errNotImplemented
}

//SendEvent send an event to server
func SendEvent(ws *websocket.Conn, event string) error {
	js := fmt.Sprintf("sendEvent('%s');", event)
	err := websocket.Message.Send(ws, js)
	if err != nil {
		return err
	}

	return nil
}

//Alert show alert message in browser
func Alert(ws *websocket.Conn, message string) error {
	js := fmt.Sprintf("alert('%s');", message)
	err := websocket.Message.Send(ws, js)
	if err != nil {
		return err
	}

	return nil
}

//BlockUI block page interaction
func BlockUI(ws *websocket.Conn) {
	Exec(ws, "$.blockUI();")
}

//UnblockUI block page interaction
func UnblockUI(ws *websocket.Conn) {
	Exec(ws, "$.unblockUI();")
}
