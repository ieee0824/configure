package configure

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"path/filepath"
	"strings"
)

// Conf to hold the configure that you read.
type Conf struct {
	loopDetector map[string]interface{}
	c            map[string]interface{}
}

// Set additional configuration items.
func (c *Conf) Set(k string, v interface{}) {
	c.c[k] = v
}

func (c *Conf) SetDefaultVal(k string, v interface{}) {
	if c.c[k] == nil {
		c.c[k] = v
	}
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
		var err error
		c.loopDetector, err = loopDetection(includeFileName, c.loopDetector)
		if err != nil {
			return
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

func loopDetection(path string, loopDetector map[string]interface{}) (map[string]interface{}, error) {
	path, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	if loopDetector[path] != nil {
		return loopDetector, errors.New("loop")
	}
	loopDetector[path] = true
	return loopDetector, nil
}

// NewConf to get the configuration item from json.
func NewConf(filePath string) *Conf {
	loop := map[string]interface{}{}
	var conf map[string]interface{}
	loop, err := loopDetection(filePath, loop)
	if err != nil {
		return nil
	}
	bin, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(bin, &conf)
	if err != nil {
		return nil
	}
	c := Conf{
		c:            conf,
		loopDetector: loop,
	}

	c.include(filePath)

	delete(c.c, "include")
	return &c
}
