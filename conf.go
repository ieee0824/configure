package configure

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
	return &Conf{map[string]interface{}{}}
}
