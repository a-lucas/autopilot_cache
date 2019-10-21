package autopilot

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

type APContacts struct {
	autopilotClient *AutopilotHttp
}

func NewApContact() *APContacts {
	http := &AutopilotHttp{token:ApiKey}
	return &APContacts{autopilotClient:http}
}

func (ap *APContacts) Create(contact *Contact) (*Contact, error) {

	type contactPostRequest struct {
		Contact *Contact `json:"contact"`
	}

	if marshalled, err := json.Marshal(contactPostRequest{Contact:contact}); err !=nil {
		return contact, err
	} else {
		if resp, err := ap.autopilotClient.Post("https://api2.autopilothq.com/v1/contact", marshalled); err !=nil {
			return contact, err
		} else {
			return contact, contact.ParseId(resp)
		}
	}
}


func (ap *APContacts) Delete(contact *Contact)  error {
	_, err := ap.autopilotClient.Delete("https://api2.autopilothq.com/v1/contact/" + contact.ContactId)
	return err
}

func (ap *APContacts) Get(id string) (*Contact, error) {
	if b, err := ap.autopilotClient.Get("https://api2.autopilothq.com/v1/contact/" + id); err !=nil {
		return nil, err
	} else {
		c := &Contact{}
		return c, json.Unmarshal(b, &c)
	}
}


type AutopilotHttp struct {
	token string
}

func (a *AutopilotHttp) Get(url string) ([]byte, error) {
	var netClient = &http.Client{
		Timeout: time.Second * 10,

	}

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("autopilotapikey", a.token)

	resp, err := netClient.Do(req)
	if err !=nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	return ioutil.ReadAll(resp.Body)
}
func (a *AutopilotHttp) Put(url string, body []byte) ([]byte, error) {
	var netClient = &http.Client{
		Timeout: time.Second * 10,

	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(body))
	req.Header.Set("autopilotapikey", a.token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := netClient.Do(req)
	if err !=nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	return ioutil.ReadAll(resp.Body)
}
func (a *AutopilotHttp) Post(url string, body []byte) ([]byte, error) {
	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err !=nil {
		return nil, err
	}
	req.Header.Set("autopilotapikey", a.token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := netClient.Do(req)
	if err !=nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	return ioutil.ReadAll(resp.Body)
}
func (a *AutopilotHttp) Delete(url string) ([]byte, error) {
	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("DELETE", url, nil)
	req.Header.Set("autopilotapikey", a.token)
	req.Header.Set("Content-Type", "application/json")


	resp, err := netClient.Do(req)
	if err !=nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	return ioutil.ReadAll(resp.Body)
}


