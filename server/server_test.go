package main

import (
	"autopilot"
	"encoding/json"
	"errors"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"math"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {

	t.Run("Test encodingJson", func(t *testing.T) {
		g := NewGomegaWithT(t)

		response := httptest.NewRecorder()

		someData := &autopilot.Contact{
			ContactId:  "A",
			Email:      "b",
			FirstName:  "c",
			LastName:   "d",
			Salutation: "e",
			Title:      "f",
			Phone:      "g",
		}

		encoded, err := json.Marshal(someData)
		g.Expect(err).To(BeNil())

		encodeJson(response, someData)

		g.Expect(response.Code).To(Equal(200))
		g.Expect(response.Body.Bytes()).To(Equal(encoded))
		g.Expect(response.Header().Get("content-type")).To(Equal("application/json"))

		nanNumber := math.NaN()
		response = httptest.NewRecorder()
		encodeJson(response, nanNumber)
		g.Expect(response.Code).To(Equal(500))
		g.Expect(response.Header().Get("content-type")).To(Equal("application/json"))
	})

	t.Run("encodeError", func(t *testing.T) {
		customErr := errors.New("some Error")
		errorMessage := "somme message"

		g := NewGomegaWithT(t)

		response := httptest.NewRecorder()

		encodeError(response, customErr, errorMessage)
		g.Expect(response.Code).To(Equal(500))

		body, err := ioutil.ReadAll(response.Body)
		g.Expect(err).To(BeNil())

		var em ErrorMessage
		err = json.Unmarshal(body, &em)

		g.Expect(err).To(BeNil())
		g.Expect(em.Error).To(Equal(errorMessage))
		g.Expect(em.Message).To(Equal(customErr.Error()))

	})
}
