package autopilot

import "encoding/json"

type Contact struct {
	ContactId  string `json:"contact_id"`
	Email      string
	twitter    string
	FirstName  string
	LastName   string
	Salutation string
	Title      string
	Phone      string
}


func (c *Contact) ParseId(b []byte) error {
	return json.Unmarshal(b, &c)
}
