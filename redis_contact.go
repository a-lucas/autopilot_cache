package autopilot

import (
	"encoding/json"
	"errors"
	"github.com/gomodule/redigo/redis"
	"time"
)

type RedisContact struct {
	pool *redis.Pool


}

//"127.0.0.1:6379
func NewRedisContact(url string) *RedisContact {
	return &RedisContact{
		pool: &redis.Pool{
			MaxIdle:     3,
			IdleTimeout: 240 * time.Second,
			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial("tcp", url)
				if err != nil {
					return nil, err
				}
				//if _, err := c.Do("AUTH", password); err != nil {
				//	c.Close()
				//	return nil, err
				//}
				//if _, err := c.Do("SELECT", db); err != nil {
				//	c.Close()
				//	return nil, err
				//}
				return c, nil

			},
		},
	}
}

func (ap *RedisContact) Create(contact *Contact) (*Contact, error) {

	if len(contact.ContactId) == 0 {
		return nil, errors.New("cannot create a cached version of contact which has not contact_id")
	}
	if marshalled, err := json.Marshal(contact); err != nil {
		return contact, err
	} else {
		conn := ap.pool.Get()
		defer conn.Close()
		_, err := conn.Do("SET", contact.ContactId, marshalled)
		return contact, err
	}
}

func (ap *RedisContact) Delete(contact *Contact)  error {
	conn := ap.pool.Get()
	defer conn.Close()
	_, err := conn.Do("DEL", contact.ContactId)
	return err
}

func (ap *RedisContact) Get(id string) (*Contact, error) {
	conn := ap.pool.Get()
	defer conn.Close()
	if b, err := redis.Bytes(conn.Do("GET", id)); err != nil {
		return nil, err
	} else {
		c := &Contact{}
		return c, json.Unmarshal(b, &c)
	}
}

