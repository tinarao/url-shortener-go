package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/tinarao/url-shortener-go/db"
	"github.com/tinarao/url-shortener-go/helpers"
	"github.com/tinarao/url-shortener-go/models"
	"gorm.io/gorm"
	"net/http"
	"os"
)

type Link struct {
	Link string
}

func Shortener(w http.ResponseWriter, r *http.Request) {
	link := &Link{}
	err := json.NewDecoder(r.Body).Decode(&link)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(link.Link) == 0 {
		w.WriteHeader(400)
		helpers.JSON("message", "Поле ссылки пусто или не предоставлено совсем", w)
		return
	}

	doc := &models.Link{
		OriginalLink: link.Link,
	}

	db.DB.Db.Create(&doc)

	helpers.JSON("link", fmt.Sprintf("%v/m/%v", os.Getenv("SELF_URL"), doc.ID), w)
	return
}

func GetAllLinks(w http.ResponseWriter, r *http.Request) {
	var links []models.Link

	db.DB.Db.Find(&links)

	bytes, _ := json.Marshal(links)
	fmt.Fprintf(w, string(bytes))
	return
}

func RedirectToShortened(w http.ResponseWriter, r *http.Request) {
	linkID := mux.Vars(r)["linkID"]
	link := &models.Link{}

	res := db.DB.Db.Where("id = ?", linkID).First(&link)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		helpers.JSON("link", os.Getenv("SELF_URL"), w)
		return
	}

	helpers.JSON("link", link.OriginalLink, w)
	return
}
