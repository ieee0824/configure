package configure

import (
	"io/ioutil"
	"testing"
)

func TestNewConf(t *testing.T) {

	if NewConf("./test/conf.json") == nil {
		fileNames := []string{}
		files, _ := ioutil.ReadDir("./test")
		for _, f := range files {
			fileNames = append(fileNames, f.Name())
		}
		t.Log("can not allocated.")
		t.Fatal(fileNames)
	}
	if NewConf("") != nil {
		t.Log("file name validate test.")
		t.Fatal("invalid error determination.")
	}
	if NewConf("foo.json") != nil {
		t.Log("file read test.")
		t.Fatal("illegal read.")
	}
	if NewConf("./test/error.json") != nil {
		t.Log("json decode test")
		t.Fatal("illegal decode.")
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
