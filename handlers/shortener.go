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
	Link  string
	Alias string
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

	if len(link.Alias) == 0 {
		w.WriteHeader(400)
		helpers.JSON("message", "Alias is not provided", w)
		return

		//lib.GenerateAlias(len: 6)
	}

	doc := &models.Link{
		OriginalLink: link.Link,
		Alias:        link.Alias,
	}

	duplicate := &models.Link{}
	if err := db.DB.Db.Where("alias = ?", link.Alias).First(&duplicate); err != nil {
		if errors.Is(err.Error, gorm.ErrRecordNotFound) {
			db.DB.Db.Create(&doc)
			helpers.JSON("link", fmt.Sprintf("%v/m/%v", os.Getenv("SELF_URL"), doc.Alias), w)
			return
		} else {
			w.WriteHeader(400)
			helpers.JSON("message", "Alias duplicate", w)
			return
		}
	}
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

func DeleteAllLings(w http.ResponseWriter, r *http.Request) {

	//if true {
	//	fmt.Fprintf(w, "Forbidden")
	//	return
	//}

	var links []models.Link
	deletedAccounts := 0

	db.DB.Db.Find(&links)

	if len(links) == 0 {
		helpers.JSON("message", "No rows to delete", w)
		return
	}

	for _, x := range links {
		db.DB.Db.Unscoped().Delete(&x)
		deletedAccounts += 1
	}

	helpers.JSON("message", fmt.Sprintf("Удалено %v аккаунтов", deletedAccounts), w)
}

func GetByAlias(w http.ResponseWriter, r *http.Request) {
	alias := mux.Vars(r)["alias"]
	link := &models.Link{}

	res := db.DB.Db.Where("alias = ?", alias).First(&link)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		w.WriteHeader(404)
		helpers.JSON("message", "Record not found", w)
		return
	}

	resp, err := json.Marshal(link)
	if err != nil {
		fmt.Println("Error while marshaling", err)
		return
	}

	fmt.Fprintf(w, string(resp))
}
