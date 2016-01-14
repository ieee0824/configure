package configure

import (
	"encoding/json"
	"io/ioutil"
)

type Conf struct {
	c map[string]interface{}
}

func (c *Conf) Set(k string, v interface{}) {
	c.c[k] = v
}

func (c *Conf) Get(k string) interface{} {
	return c.c[k]
}

func NewConf(path string) *Conf {
	var c Conf
	bin, err := ioutil.ReadFile(path)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(bin, &c)
	if err != nil {
		return nil
	}
	return &c
}
