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

func (c *Conf) Write(path string) error {
	j, err := json.Marshal(c.c)

	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, j, 0644)
}

func (c *Conf) Read(path string) {
	bin, err := ioutil.ReadFile(path)
	if err != nil {
		c.c = nil
		return
	}
	err = json.Unmarshal(bin, &c.c)
	if err != nil {
		c.c = nil
		return
	}
}

func NewConf() *Conf {
	return &Conf{map[string]interface{}{}}
}
