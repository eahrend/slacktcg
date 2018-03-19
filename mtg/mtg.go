package mtg

import (
	"encoding/json"
	"net/http"

	"github.com/nlopes/slack"
)

func CardSearch(w http.ResponseWriter, r *http.Request) {
	// Do card search stuff here
	cmd, _ := slack.SlashCommandParse(r)
	b, _ := json.Marshal(&cmd)
	w.Write(b)
	/*
		card, err := sdk.NewQuery().Random(1)
		if err != nil {
			log.Fatal(err)
		}
		w.Write([]byte(card[0].Name))
	*/
}
