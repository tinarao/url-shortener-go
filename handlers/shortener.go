package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
	"github.com/tinarao/url-shortener-go/db"
	"github.com/tinarao/url-shortener-go/helpers"
	"github.com/tinarao/url-shortener-go/models"
	"gorm.io/gorm"
	"log/slog"
	"net/http"
	"os"
	"strings"
)

type Link struct {
	Link  string
	Alias string
}

func Shortener(w http.ResponseWriter, r *http.Request) {
	link := &Link{}
	err := json.NewDecoder(r.Body).Decode(&link)
	if err != nil {
		slog.Error("Error while decoding body", err)
		helpers.JSON("message", "Произошла ошибка, попробуйте ещё раз", w)
		return
	}

	if len(link.Link) == 0 {
		w.WriteHeader(400)
		helpers.JSON("message", "Поле ссылки пусто или не предоставлено совсем", w)
		return
	}

	validURL := govalidator.IsURL(link.Link)
	if !validURL {
		w.WriteHeader(400)
		helpers.JSON("message", "Provided URL is not valid", w)
		return
	}

	if len(link.Alias) == 0 {
		w.WriteHeader(400)
		helpers.JSON("message", "Alias is not provided", w)
		return

		// TODO: lib.GenerateAlias(len: 6)
	}

	if !strings.Contains(link.Link, "https://") {
		link.Link = fmt.Sprintf("https://%v", link.Link)
	}

	doc := &models.Link{
		OriginalLink: link.Link,
		Alias:        link.Alias,
	}

	duplicate := &models.Link{}
	if err := db.DB.Db.Where("alias = ?", link.Alias).First(&duplicate); err != nil {
		if errors.Is(err.Error, gorm.ErrRecordNotFound) {

			// ok scenario!

			db.DB.Db.Create(&doc)
			helpers.JSON("link", fmt.Sprintf("%v/l/%v", os.Getenv("SELF_URL"), doc.Alias), w)
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
	_, err := fmt.Fprintf(w, string(bytes))
	if err != nil {
		w.WriteHeader(500)
		helpers.JSON("message", "Произошла ошибка, попробуйте ещё раз", w)
		return
	}

	return
}

func RedirectToShortened(w http.ResponseWriter, r *http.Request) {
	alias := mux.Vars(r)["alias"]
	link := &models.Link{}

	res := db.DB.Db.Where("alias = ?", alias).First(&link)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		helpers.JSON("link", os.Getenv("SELF_URL"), w)
		return
	}

	http.Redirect(w, r, link.OriginalLink, 302)
	return
}

func DeleteAllLinks(w http.ResponseWriter, r *http.Request) {

	if true {
		// blocked by default to prevent missclicks and stuff
		fmt.Fprintf(w, "Forbidden")
		return
	}

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
	return
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
		slog.Error("Error while marshaling", err)
		return
	}

	fmt.Fprintf(w, string(resp))
	return
}
