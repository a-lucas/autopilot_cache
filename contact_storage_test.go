package autopilot

import (
	. "github.com/onsi/gomega"
	"testing"
)

func testContactStorage(Api IContactStorage,  t *testing.T) {

	t.Run("Testing storage ", func(t *testing.T) {
		g := NewGomegaWithT(t)

		contact := &Contact{Email:"james@test.com", ContactId: "DummyId"}

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
	})
}


func TestContactStorages(t *testing.T) {

	t.Run("AutoPilot API", func(t *testing.T) {
		testContactStorage(NewApContact(), t)
	})

	t.Run("Redis API", func(t *testing.T) {
		testContactStorage(NewRedisContact(RedisUrl), t)
	})

	t.Run("MiddleWare API", func(t *testing.T) {
		testContactStorage(NewContactMiddleWare(), t)
	})
}