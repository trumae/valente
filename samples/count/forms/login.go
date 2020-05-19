package forms

import (
	"log"

	"github.com/trumae/valente"
	"github.com/trumae/valente/action"
)

const htmlForm = `
<h3>Login</h3>
<p>
User: admin<br/>
Password: admin
</p>
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

func loginOk(ws *action.WebSocket, app *valente.App, params []string) {
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
		action.HTML(ws, "loginName", "admin")
		app.GoTo("home", nil)
	}
}

//Initialize inits the login Form
func (form FormLogin) Initialize(ws *action.WebSocket) valente.Form {
	log.Println("FormLogin Initialize")

	f := form.AddEventHandler("loginOk", loginOk)

	action.HTML(ws, "content", htmlForm)

	return f
}
