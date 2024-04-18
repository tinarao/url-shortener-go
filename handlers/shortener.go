package handlers

import (
	"encoding/json"
	"github.com/tinarao/url-shortener-go/helpers"
	"net/http"
)

type ShortenerDTO struct {
	Link string
}

// generate 8-char uuid
// save to pg

func Shortener(w http.ResponseWriter, r *http.Request) {
	var reqBody ShortenerDTO
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	helpers.JSON("link", reqBody.Link, w)

}
