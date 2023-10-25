package handlers

import (
	"log"
	"net/http"

	"github.com/CloudyKit/jet/v6"
	"github.com/gorilla/websocket"
)

var views = jet.NewSet(

	jet.NewOSFileSystemLoader("./html"),
	jet.InDevelopmentMode(),
)

// Upgrade the connection to weboscket connection
var upgradeConnection = websocket.Upgrader{

	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// Security check
	CheckOrigin: func(r *http.Request) bool { return true },
}

func Home(w http.ResponseWriter, r *http.Request) {

	err := renderPage(w, "home.jet", nil)
	if err != nil {
		log.Println(err)

	}

}

// WebSocketJSONresponse defines the response sent back from websocket
type WebSocketJSONresponse struct {
	Action      string `json:"action"`
	Message     string `json:"message"`
	MessageType string `json:"message_type"`
}

// WebsocketEndPoint upgrades connection to websocket
func WebsocketEndPoint(w http.ResponseWriter, r *http.Request) {

	ws, err := upgradeConnection.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Client connected to endpoint 😎🤘")

	var response WebSocketJSONresponse
	response.Message = `<em><small>Connected to server 🤓🤘</small></em>`

	err = ws.WriteJSON(response)
	if err != nil {
		log.Println(err)
	}
}

func renderPage(w http.ResponseWriter, tmpl string, data jet.VarMap) error {
	// view = html template
	view, err := views.GetTemplate(tmpl)
	if err != nil {
		log.Println(err)
		return err
	}

	err = view.Execute(w, data, nil)
	if err != nil {
		if err != nil {
			log.Println(err)
			return err
		}
	}

	return nil
}
