//cmd/webserver/main.go
package main

import (
	poker "github.com/dexfs/go-fundamentals/http-server"
	"log"
	"net/http"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}

	defer close()

	server := poker.NewPlayerServer(store)
	if err := http.ListenAndServe(":5005", server); err != nil {
		log.Fatalf("could not listen on port 5005 %v", err)
	}
}
