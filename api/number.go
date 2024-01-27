package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func SpecificComic(w http.ResponseWriter, r *http.Request) {
	// contrary to what you might find online,
	// do not send a status code
	// because it sends the data as text/plain
	// idk why
	// w.WriteHeader(http.StatusCreated)

	// set Content-Type so that we send JSON
	w.Header().Set("Content-Type", "application/json")

	// extract comic number from path
	comic := r.URL.Path[1:]

	// reverse compatibility with old API
	// /api/comic?num=327
	queryParams := r.URL.Query().Get("num")
	if queryParams != "" {
		comic = queryParams
	}

	// get the specific xkcd comic
	resp, err := http.Get(fmt.Sprintf("https://xkcd.com/%s/info.0.json", comic))
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
