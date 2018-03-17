package mtg

import (
	"log"
	"net/http"

	sdk "github.com/MagicTheGathering/mtg-sdk-go"
)

func CardSearch(w http.ResponseWriter, r *http.Request) {
	// Do card search stuff here
	card, err := sdk.NewQuery().Random(1)
	if err != nil {
		log.Fatal(err)
	}
	w.Write([]byte(card[0].Name))
}
