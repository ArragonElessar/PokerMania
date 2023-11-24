package playercontroller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	Dealer "github.com/ArragonElessar/PokerMania/models/dealer"
	"github.com/ArragonElessar/PokerMania/models/player"
)

// handle the request to create a new player
func InitializePlayer(dealer *Dealer.Dealer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// defer close of the request
		defer r.Body.Close()

		// set the header content type
		w.Header().Set("Content-Type", "application/json")

		var values *player.Player
		err := json.NewDecoder(r.Body).Decode(&values)
		if err != nil {
			fmt.Println(err)
		}

		// firstName := fmt.Sprint(values["firstName"])
		// lastName := fmt.Sprint(values["lastName"])
		// money, _ := strconv.Atoi(fmt.Sprint(values["money"]))
		values.FirstName = fmt.Sprint(values.FirstName)
		values.LastName = fmt.Sprint(values.LastName)
		values.Money, _ = strconv.Atoi(fmt.Sprint(values.Money))

		// call function to create a player
		player := player.CreateNewPlayer(values.FirstName, values.LastName, values.Money)

		// now add this player to the dealer's table
		result := dealer.AddPlayer(player)
		if result == -1 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Cannot add player, max players reached"))
			return
		} else if result == -2 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Player already exists"))
			return
		} else if result == 1 {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Player added successfully"))
			return
		}

	}
}

// handle the request to view all players and their chips and monry
func ViewPlayers(dealer *Dealer.Dealer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// set the header content type
		w.Header().Set("Content-Type", "application/json")

		// if there are no active players
		if len(dealer.ActivePlayers) == 0 {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("No active players"))
			return
		}

		// return the encoding of all the active players
		responseJson, err := json.Marshal(dealer.ActivePlayers)
		if err != nil {
			log.Fatal(err)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(responseJson))
	}
}
