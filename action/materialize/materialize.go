package materialize

import (
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/trumae/valente/action"
)

//Toast show toast message in browser
func Toast(ws *websocket.Conn, message string, t uint) error {
	js := fmt.Sprintf("Materialize.toast('%s', %d);", message, t)
	err := ws.WriteMessage(websocket.TextMessage, []byte(js))
	if err != nil {
		return err
	}

	return nil
}

//AdjustInputLabels adjust sobreposition of labels into input texts
func AdjustInputLabels(ws *websocket.Conn) error {
	return action.Exec(ws, `
	$("input:text").filter(function() {
		return this.value != "";
	}).each(function() {
		$(this).addClass("active");
	});
	Materialize.updateTextFields();	
	`)
}

//SetupForm setup materialize
func SetupForm(ws *websocket.Conn) error {
	return action.Exec(ws, `
$('select').material_select();

$('.datepicker').pickadate({
    selectMonths: true, // Creates a dropdown to control month
    selectYears: 15, // Creates a dropdown of 15 years to control year,
    today: 'Hoje',
    clear: 'Limpar',
    close: 'Ok',
    closeOnSelect: false // Close upon selecting a date,
});

$('.timepicker').pickatime({
    default: 'now', // Set default time: 'now', '1:30AM', '16:30'
    fromnow: 0,       // set default time to * milliseconds from now (using with default = 'now')
    twelvehour: true, // Use AM/PM or 24-hour format
    donetext: 'OK', // text for done-button
    cleartext: 'Limpar', // text for clear-button
    canceltext: 'Cancelar', // Text for cancel-button
    autoclose: false, // automatic close timepicker
    ampmclickable: false, // make AM PM clickable
    aftershow: function(){} //Function for after opening timepicker
});
	`)
}
