package cache

import "time"

type Item struct {
	v string
	isInf bool
	deadline time.Time
}

type Cache struct {
	m map[string]Item
}

func NewCache() Cache {
	return Cache{
		m : make(map[string]Item),
	}
}

func (c Cache) Get(key string) (string, bool) {
	now := time.Now()
	data, ok := c.m[key]
	if !ok {
		return "", ok
	}
	if !data.isInf {
		notExp := now.Before(data.deadline)
		if notExp {
			return data.v, true
		} else {
			delete(c.m, key)
			return "", false
		}
	}
	return data.v, ok
}

func (c Cache) Put(key, value string) {
	 c.m[key] = Item{
		 v : value,
		 isInf: true,
	 }
}

func (c Cache) Keys() (ks []string) {
	now := time.Now()
	ks = make([]string, 0, len(c.m))
	for k, v := range c.m {
		if now.Before(v.deadline) {
			ks = append(ks, k)
		}
	}
	return
}

func (c Cache) PutTill(key, value string, deadline time.Time) {
	c.m[key] = Item {
		v : value,
		deadline: deadline,
	}
}
