package configure

import (
	"encoding/json"
	"io/ioutil"
)

// Conf to hold the configure that you read.
type Conf struct {
	c map[string]interface{}
}

// Set additional configuration items.
func (c *Conf) Set(k string, v interface{}) {
	c.c[k] = v
}

// Get is to get the configuration item.
func (c *Conf) Get(k string) interface{} {
	return c.c[k]
}

// NewConf to get the configuration item from json.
func NewConf(filePath string) *Conf {
	if filePath == "" {
		return nil
	}
	c := Conf{map[string]interface{}{}}
	bin, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(bin, &c)
	if err != nil {
		return nil
	}
	return &c
}
