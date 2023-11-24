package main

import (
	"fmt"
	"log"
	"net/http"

	Dealer "github.com/ArragonElessar/PokerMania/models/dealer"
	"github.com/ArragonElessar/PokerMania/routes"
)

const port = 4000

func main() {

	/*
	   1. Create a router to host the game
	   2. Initate a dealer with the required parameters
	   3. Decide the pot etc.
	   4. Start accepting clients
	   5. For an interruptable infinite loop, do the following:
	   	5.1 Wait for 2 players
	   	5.2 Play the game
	   	5.3 Check if any player has exited and return back to step 1


	*/

	// 1. Create the router
	dealer := Dealer.CreateNewDealer("Sweety", 1)

	// 2. Get the router
	router := routes.CreateRouter(dealer)

	// 2.1 Start the server
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
