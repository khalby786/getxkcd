package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func LatestComic(w http.ResponseWriter, r *http.Request) {
	// contrary to what you might find online,
	// do not send a status code
	// because it sends the data as text/plain
	// idk why
	// w.WriteHeader(http.StatusCreated)

	// set Content-Type so that we send JSON
	w.Header().Set("Content-Type", "application/json")

	// get the latest xkcd comic
	resp, err := http.Get("https://xkcd.com/info.0.json")
	if err != nil {
		log.Fatalln(err)
	}

	// make it readable or something i really don't know
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	if err != nil {
		fmt.Println("Error happened somewhere along the way. Err:", err)
	} else {
		// 200 send it lads
		w.Write(body)
	}
}
