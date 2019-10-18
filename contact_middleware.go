package autopilot

type ContactCacheMiddleWare struct {
	redis IContactStorage
	autopilot IContactStorage
}

func NewContactMiddleWare() *ContactCacheMiddleWare {
	return &ContactCacheMiddleWare{
		redis:     NewRedisContact(RedisUrl),
		autopilot: NewApContact(),
	}
}

/**
	Load from redis
	if not in redis, Load from AutoPilot
	then cache into Redis
	and return object
*/
func (c *ContactCacheMiddleWare) Get(key string) (*Contact, error) {
	if contact, err := c.redis.Get(key); err !=nil {
		if contact, err := c.autopilot.Get(key); err !=nil {
			return nil, err
		} else {
			return c.redis.Create(contact)
		}
	} else {
		return contact, nil
	}

	return nil, nil
}

/**
	Store in autopilot, then store in redis
 */
func (c *ContactCacheMiddleWare) Create(contact *Contact) (*Contact, error) {
	if contact, err := c.autopilot.Create(contact); err !=nil {
		return contact, err
	} else {
		return c.redis.Create(contact)
	}
}

/**
	Store in autopilot, then store in redis
*/
func (c *ContactCacheMiddleWare)Delete(contact *Contact)  error {
	redis := NewRedisContact(RedisUrl)
	defer redis.Delete(contact)
	api := NewApContact()
	return api.Delete(contact)
}
