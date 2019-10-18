package autopilot

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestContactApi(t *testing.T) {
	g := NewGomegaWithT(t)

	Api := NewApContact()

	contact := &Contact{Email:"james@test.com"}

	contact, err := Api.Create(contact)
	g.Expect(err).To(BeNil())
	g.Expect(len(contact.ContactId)>0).To(BeTrue())

	contact2, err := Api.Get(contact.ContactId)
	g.Expect(err).To(BeNil())
	g.Expect(contact2.Email).To(Equal(contact.Email))

	err = Api.Delete(contact)
	g.Expect(err).To(BeNil())

	_, err = Api.Get(contact.ContactId)
	g.Expect(err).ToNot(BeNil())

}
