package autopilot


type IContactStorage interface {
	Create(contact *Contact) (*Contact, error)
	Delete(contact *Contact)  error
	Get(id string) (*Contact, error)
}
