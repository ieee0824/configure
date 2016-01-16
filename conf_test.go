package configure

import (
	"io/ioutil"
	"testing"
)

func TestNewConf(t *testing.T) {
	c := NewConf("./test/conf.json")

	if c == nil {
		fileNames := []string{}
		files, _ := ioutil.ReadDir("./test")
		for _, f := range files {
			fileNames = append(fileNames, f.Name())
		}
		t.Fatal(fileNames)
	}
}

func TestSet(t *testing.T) {
	c := NewConf("./test/conf.json")
	c.Set("loc", "tokyo")
	if c.c["loc"] != "tokyo" {
		t.Fatal(c.c)
	}
}

func TestGet(t *testing.T) {
	c := NewConf("./test/conf.json")
	c.c["loc"] = "akiba"
	loc := c.Get("loc")

	s, ok := loc.(string)
	if !ok || s != "akiba" {
		t.Fatal(ok, "'"+s+"'")
	}
}
