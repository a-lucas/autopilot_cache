package main

import (
	"autopilot"
	"encoding/json"
	"errors"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

var contact *autopilot.ContactCacheMiddleWare
var contactOnce sync.Once

func contactMiddleWare() *autopilot.ContactCacheMiddleWare {
	contactOnce.Do(func() {
		contact = autopilot.NewContactMiddleWare()
	})
	return contact
}

func encodeJson(w http.ResponseWriter, data interface{}) {
	w.Header().Set("content-type", "application/json")
	if b, err := json.Marshal(data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(200)
		w.Write(b)
	}
}

type ErrorMessage struct {
	Error   string `json:'error'`
	Message string `json:"message"`
}

func encodeError(w http.ResponseWriter, err error, errorType string) {
	em := ErrorMessage{errorType, err.Error()}
	b, _ := json.Marshal(em)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(b)
}

func NewUpdateContact(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if body, err := ioutil.ReadAll(r.Body); err != nil {
		encodeError(w, err, "Invalid contact payload in body")
	} else {
		c := &autopilot.Contact{}
		if err := json.Unmarshal(body, &c); err != nil {
			encodeError(w, err, "Invalid json contact payload in body")
		} else {
			if contact, err := contactMiddleWare().Create(c); err != nil {
				encodeError(w, err, "Error creating new contact")
			} else {
				encodeJson(w, contact)
			}
		}
	}
}

func GetContact(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if contact, err := contactMiddleWare().Get(ps.ByName("id")); err != nil {
		encodeError(w, errors.New("contact not found: "+ps.ByName("id")), "Logic Error")
	} else {
		encodeJson(w, contact)
	}
}

func main() {
	router := httprouter.New()
	router.GET("/contact/:id", GetContact)
	router.POST("/contact", NewUpdateContact)
	router.PUT("/contact", NewUpdateContact)

	log.Fatal(http.ListenAndServe(":8080", router))
}
