package mtg

import (
	"io/ioutil"
	"log"
	"net/http"
)

func CardSearch(w http.ResponseWriter, r *http.Request) {
	// Do card search stuff here
	body_bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(body_bytes)
	/*
		card, err := sdk.NewQuery().Random(1)
		if err != nil {
			log.Fatal(err)
		}
		w.Write([]byte(card[0].Name))
	*/
}
