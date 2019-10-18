package autopilot

import (
	"encoding/json"
	. "github.com/onsi/gomega"
	"testing"
)

func TestContact(t *testing.T) {

	g := NewGomegaWithT(t)
	type contactTmp struct {
		Id string `json:"contact_id"`
	}

	testId := &contactTmp{Id:"test  Id"}

	b, err := json.Marshal(testId)
	g.Expect(err).To(BeNil())

	contact := &Contact{Email:"My@email.com"}

	err  = contact.ParseId(b)
	g.Expect(err).To(BeNil())

	g.Expect(contact.ContactId).To(Equal(testId.Id))
}
