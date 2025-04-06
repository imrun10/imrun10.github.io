package main

import (
	"fmt"
	"net/http"
	"os"                 // add this import
	"portfolio/handlers" // add this import

	"github.com/a-h/templ"
)

func main() {
	// Render the header component

	// Set up routes
	var cards CardsData
	cards = GetCards()

	for _, card := range cards.Cards {
		fmt.Println(card.ProjectName)
		fmt.Println(card.ProjectDescription)
		fmt.Println(card.ImgDir)
		fmt.Println(card.ReadMore)
		fmt.Println(card.ReadMoreLink)
		fmt.Println(card.Github)
		fmt.Println(card.GithubLink)
	}
	Body := Body("My Portfolio", cards.Cards) // Note the package prefix
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	http.Handle("/", templ.Handler(Body))
	http.HandleFunc("/mouse_entered", handlers.Test) // now this is recognized

	// Start server
	fmt.Println("Listening on :3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Printf("Server failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Server started successfully")
}
