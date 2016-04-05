package forms

import (
	"log"

	"github.com/trumae/valente"
	"github.com/trumae/valente/ws/action"
	"golang.org/x/net/websocket"
)

const htmlForm = `
<h3>Login</h3>
<form>	

   <label for="text-basic">User:</label>
   <input name="text-basic" id="user" value="" type="text">

   <label for="password">Password:</label>
   <input name="password" id="pass" value="" autocomplete="off" type="password"/>

	 <a href="javascript:sendEvent('loginOk')">Ok</a>
</form>
`

//FormLogin is a form for auth
type FormLogin struct {
	valente.FormImpl
}

func loginOk(ws *websocket.Conn, app *valente.App) {
	log.Println("Handling loginOk")

	user, err := action.Get(ws, "user")
	if err != nil {
		log.Println(err)
	}

	pass, err := action.Get(ws, "pass")
	if err != nil {
		log.Println(err)
	}

	if user == "admin" && pass == "admin" {
		action.Html(ws, "loginName", "admin")
		app.GoTo("home")
	}
}

//Initialize inits the login Form
func (form FormLogin) Initialize(ws *websocket.Conn) valente.Form {
	log.Println("FormLogin Initialize")

	f := form.AddEventHandler("loginOk", loginOk)

	action.Html(ws, "content", htmlForm)

	return f
}
