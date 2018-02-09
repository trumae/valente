package action

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/trumae/valente/status"
)

var (
	errNotImplemented = errors.New("Not Implemented!")
)

// Exec execute the js code on WebBrowser
func Exec(ws *websocket.Conn, js string) error {
	err := ws.WriteMessage(websocket.TextMessage, []byte(js))
	if err != nil {
		return err
	}
	status.Status.SendedBytes += len(js)
	return err
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

	err := ws.WriteMessage(websocket.TextMessage, []byte(js))
	if err != nil {
		return err
	}
	status.Status.SendedBytes += len(js)
	return err
}

//HTML replace target with new content
func HTML(ws *websocket.Conn, target, content string) error {
	c := strings.Replace(content, "\n", "\\n", -1)
	c = strings.Replace(c, "\"", "\\\"", -1)
	js := fmt.Sprintf("$( \"#%s\" ).html(\"%s\");", target, c)

	err := ws.WriteMessage(websocket.TextMessage, []byte(js))
	if err != nil {
		return err
	}
	status.Status.SendedBytes += len(js)
	return err
}

//Hide the target
func Hide(ws *websocket.Conn, target string, duration string) error {
	js := fmt.Sprintf("$( \"#%s\" ).hide(\"%s\");", target, duration)

	err := ws.WriteMessage(websocket.TextMessage, []byte(js))
	if err != nil {
		return err
	}
	status.Status.SendedBytes += len(js)
	return err
}

//Show the target
func Show(ws *websocket.Conn, target string, duration string) error {
	js := fmt.Sprintf("$( \"#%s\" ).show(\"%s\");", target, duration)

	err := ws.WriteMessage(websocket.TextMessage, []byte(js))
	if err != nil {
		return err
	}
	status.Status.SendedBytes += len(js)
	return err
}

//Remove target from the DOM
func Remove(ws *websocket.Conn, target string) error {
	return errNotImplemented
}

//Append concate content at target
func Append(ws *websocket.Conn, target, content string) error {
	c := strings.Replace(content, "\n", "\\n", -1)
	c = strings.Replace(c, "\"", "\\\"", -1)
	js := fmt.Sprintf("$( \"#%s\" ).append(\"%s\");", target, c)

	err := ws.WriteMessage(websocket.TextMessage, []byte(js))
	if err != nil {
		return err
	}
	status.Status.SendedBytes += len(js)
	return err
}

//Prepend concate content at the begin of target
func Prepend(ws *websocket.Conn, target, content string) error {
	c := strings.Replace(content, "\n", "\\n", -1)
	c = strings.Replace(c, "\"", "\\\"", -1)
	js := fmt.Sprintf("$( \"#%s\" ).prepend(\"%s\");", target, c)

	err := ws.WriteMessage(websocket.TextMessage, []byte(js))
	if err != nil {
		return err
	}
	status.Status.SendedBytes += len(js)
	return err
}

//Redirect to url
func Redirect(ws *websocket.Conn, url string) error {
	return errNotImplemented
}

//AddClass add a class for an element
func AddClass(ws *websocket.Conn, target, class string) error {
	js := fmt.Sprintf("$('#%s').addClass('%s');", target, class)
	err := ws.WriteMessage(websocket.TextMessage, []byte(js))
	if err != nil {
		return err
	}
	status.Status.SendedBytes += len(js)
	return nil
}

//RemoveClass add a class for an element
func RemoveClass(ws *websocket.Conn, target, class string) error {
	js := fmt.Sprintf("$('#%s').removeClass('%s');", target, class)
	err := ws.WriteMessage(websocket.TextMessage, []byte(js))
	if err != nil {
		return err
	}
	status.Status.SendedBytes += len(js)
	return nil
}

//Set update a form element (textbox, dropdown, checkbox, etc) to set text value of TargetID.
func Set(ws *websocket.Conn, target, value string) error {
	c := strings.Replace(value, "\n", "\\n", -1)
	c = strings.Replace(c, "\"", "\\\"", -1)
	js := fmt.Sprintf("$('#%s').val('%s');", target, c)
	err := ws.WriteMessage(websocket.TextMessage, []byte(js))
	if err != nil {
		return err
	}
	status.Status.SendedBytes += len(js)
	return nil
}

//Get content of form element
func Get(ws *websocket.Conn, target string) (string, error) {
	ret := ""
	js := fmt.Sprintf("ws.send($('#%s').val());", target)
	status.Status.SendedBytes += len(js)
	err := ws.WriteMessage(websocket.TextMessage, []byte(js))
	if err != nil {
		return "", err
	}

	_, bret, err := ws.ReadMessage()
	if err != nil {
		return "", err
	}
	ret = string(bret)
	status.Status.ReceivedBytes += len(ret)

	return ret, nil
}

//Wire bind an action to an event on target
func Wire(ws *websocket.Conn, target, event, act string) error {
	return errNotImplemented
}

//SendEvent send an event to server
func SendEvent(ws *websocket.Conn, event string) error {
	js := fmt.Sprintf("sendEvent('%s');", event)
	err := ws.WriteMessage(websocket.TextMessage, []byte(js))
	if err != nil {
		return err
	}
	status.Status.SendedBytes += len(js)

	return nil
}

//Alert show alert message in browser
func Alert(ws *websocket.Conn, message string) error {
	js := fmt.Sprintf("alert('%s');", message)
	err := ws.WriteMessage(websocket.TextMessage, []byte(js))
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
func BlockUI(ws *websocket.Conn) {
	Exec(ws, fmt.Sprintf("$.blockUI({ message: '%s' });", BlockMessage))
}

//UnblockUI block page interaction
func UnblockUI(ws *websocket.Conn) {
	Exec(ws, "$.unblockUI();")
}
