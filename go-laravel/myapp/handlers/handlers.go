package handlers

import (
	"fmt"
	"github.com/CloudyKit/jet/v6"
	"github.com/johnwr-response/celeritas"
	"myapp/data"
	"net/http"
	"time"
)

type Handlers struct {
	App    *celeritas.Celeritas
	Models data.Models
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	defer h.App.LoadTime(time.Now())
	err := h.render(w, r, "home", nil, nil)
	if err != nil {
		h.App.ErrorLog.Println("Error rendering:", err)
	}
}

func (h *Handlers) GoPage(w http.ResponseWriter, r *http.Request) {
	err := h.App.Render.GoPage(w, r, "home", nil)
	if err != nil {
		h.App.ErrorLog.Println("Error rendering:", err)
	}
}

func (h *Handlers) JetPage(w http.ResponseWriter, r *http.Request) {
	err := h.App.Render.JetPage(w, r, "jet-template", nil, nil)
	if err != nil {
		h.App.ErrorLog.Println("Error rendering:", err)
	}
}

func (h *Handlers) SessionTest(w http.ResponseWriter, r *http.Request) {
	myData := "bar"
	h.App.Session.Put(r.Context(), "foo", myData)
	myValue := h.App.Session.GetString(r.Context(), "foo")
	vars := make(jet.VarMap)
	vars.Set("foo", myValue)

	err := h.App.Render.JetPage(w, r, "sessions", vars, nil)
	if err != nil {
		h.App.ErrorLog.Println("Error rendering:", err)
	}
}

func (h *Handlers) JSON(w http.ResponseWriter, _ *http.Request) {
	var payload struct {
		ID      int64    `json:"id"`
		Name    string   `json:"content"`
		Hobbies []string `json:"hobbies"`
	}

	payload.ID = 10
	payload.Name = "Jack Jones"
	payload.Hobbies = []string{"Karate", "Tennis", "Programming"}

	err := h.App.WriteJSON(w, http.StatusOK, payload)
	if err != nil {
		h.App.ErrorLog.Println("Error rendering:", err)
	}
}

func (h *Handlers) XML(w http.ResponseWriter, _ *http.Request) {
	type Payload struct {
		ID      int64    `xml:"id"`
		Name    string   `xml:"name"`
		Hobbies []string `xml:"hobbies>hobby"`
	}

	var payload Payload

	payload.ID = 10
	payload.Name = "John Smith"
	payload.Hobbies = []string{"Karate", "Tennis", "Programming"}

	err := h.App.WriteXML(w, http.StatusOK, payload)
	if err != nil {
		h.App.ErrorLog.Println("Error rendering:", err)
	}
}

func (h *Handlers) DownloadFile(w http.ResponseWriter, r *http.Request) {
	_ = h.App.DownloadFile(w, r, "./public/images", "celeritas.jpg")
}

func (h *Handlers) TestCrypto(w http.ResponseWriter, _ *http.Request) {
	plainText := "Hello World!"
	_, _ = fmt.Fprint(w, "Unencrypted: ", plainText+"\n")
	encrypted, err := h.encrypt(plainText)
	if err != nil {
		h.App.ErrorLog.Println("Error encrypting data:", err)
		h.App.Error500(w)
		return
	}
	_, _ = fmt.Fprint(w, "Encrypted: ", encrypted+"\n")
	decrypted, err := h.decrypt(encrypted)
	if err != nil {
		h.App.ErrorLog.Println("Error decrypting data:", err)
		h.App.Error500(w)
		return
	}
	_, _ = fmt.Fprint(w, "Decrypted: ", decrypted+"\n")
}
