package configure

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"strings"
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

// GetKeys is to get configuration keys
func (c *Conf) GetKeys() []string {
	if c.c == nil {
		return nil
	} else if len(c.c) == 0 {
		return nil
	}
	keys := make([]string, 0, len(c.c))
	for k := range c.c {
		keys = append(keys, k)
	}
	return keys
}

func (c *Conf) getIncludeFileNames() []string {
	if fileNames, ok := c.Get("include").([]interface{}); ok {
		return func() []string {
			ret := make([]string, 0, len(fileNames))
			for _, fileName := range fileNames {
				if name, ok := fileName.(string); ok {
					ret = append(ret, name)
				}
			}
			return ret
		}()
	}
	return nil
}

func (c *Conf) include(parentConfPath string) {
	includeFileNames := c.getIncludeFileNames()
	for _, includeFileName := range includeFileNames {
		if !filepath.IsAbs(includeFileName) {
			confPath := func() string {
				parentConfPath, err := filepath.Abs(parentConfPath)
				if err != nil {
					return ""
				}
				splitPath := strings.Split(parentConfPath, "/")
				return "/" + strings.Join(splitPath[:len(splitPath)-1], "/")
			}()
			var err error
			includeFileName, err = filepath.Abs(confPath + "/" + includeFileName)
			if err != nil {
				continue
			}
		}
		childConf := NewConf(includeFileName)
		if childConf == nil {
			continue
		}
		childKeys := childConf.GetKeys()
		for _, childKey := range childKeys {
			c.Set(childKey, childConf.Get(childKey))
		}
	}
}

// NewConf to get the configuration item from json.
func NewConf(filePath string) *Conf {
	var conf map[string]interface{}
	bin, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(bin, &conf)
	if err != nil {
		return nil
	}

	c := Conf{
		c: conf,
	}

	c.include(filePath)

	delete(c.c, "include")
	return &c
}
