package forms

import (
	"log"

	"github.com/gorilla/websocket"
	"github.com/trumae/valente"
	"github.com/trumae/valente/action"
)

const htmlForm = `
<center>
<ul data-role="listview" data-inset="true" style="width:280px;">
  <li data-role="list-divider">Login</li>
  <li>
User: admin<br/>
Password: admin
  </li>
  <li>
    <form>	

      <label for="text-basic">User:</label>
      <input name="text-basic" id="user" value="" type="text">

      <label for="password">Password:</label>
      <input name="password" id="pass" value="" autocomplete="off" type="password"/>

	    <a href="javascript:sendEvent('loginOk')" class="ui-shadow ui-btn ui-corner-all" >Ok</a>
    </form>
 </li>
</ul>
</center>
`

//FormLogin is a form for auth
type FormLogin struct {
	valente.FormImpl
}

func loginOk(ws *websocket.Conn, app *valente.App, params []string) {
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
func (form FormLogin) Initialize(ws *websocket.Conn) valente.Form {
	log.Println("FormLogin Initialize")

	f := form.AddEventHandler("loginOk", loginOk)

	action.HTML(ws, "content", htmlForm)
	applyJQM(ws)

	return f
}
