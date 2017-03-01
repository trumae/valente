package status

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

//StatusInfo represents info of app status
type StatusInfo struct {
	Started        time.Time
	OpenSessions   int
	ClosedSessions int
	Gotos          int
	SendedBytes    int
	ReceivedBytes  int
}

var (
	//Status is app status
	Status StatusInfo
)

//ValenteStatusHandler handle a status request sending an json
func ValenteStatusHandler(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(Status)
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(b))
}
