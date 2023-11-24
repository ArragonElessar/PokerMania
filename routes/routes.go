package routes

import (
	playercontroller "github.com/ArragonElessar/PokerMania/controllers/player_controller"
	Dealer "github.com/ArragonElessar/PokerMania/models/dealer"
	"github.com/gorilla/mux"
)

func CreateRouter(dealer *Dealer.Dealer) *mux.Router {

	r := *mux.NewRouter()
	r.HandleFunc("/player/create", playercontroller.InitializePlayer(dealer)).Methods("POST")
	r.HandleFunc("/player/viewall", playercontroller.ViewPlayers(dealer)).Methods("GET")
	return &r
}
