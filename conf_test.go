package configure

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestNewConf(t *testing.T) {
	c := NewConf()

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
	c := NewConf()
	c.Set("loc", "tokyo")
	if c.c["loc"] != "tokyo" {
		t.Fatal(c.c)
	}
}

func TestGet(t *testing.T) {
	c := NewConf()
	c.c["loc"] = "akiba"
	loc := c.Get("loc")

	s, ok := loc.(string)
	if !ok || s != "akiba" {
		t.Fatal(ok, "'"+s+"'")
	}
}

func TestWrite(t *testing.T) {
	JSON := "{\"loc\":\"kanda\"}"
	c := NewConf()
	c.c["loc"] = "kanda"
	err := c.Write("./test/tmp.json")
	if err != nil {
		t.Fatal(err)
	}
	bin, err := ioutil.ReadFile("./test/tmp.json")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(bin, []byte(JSON)) {
		t.Fatal(string(bin), JSON)
	}
	os.Remove("./test/tmp.json")
}

func TestRead(t *testing.T) {
	c := NewConf()
	c.Read("./test/conf.json")

	if c.c == nil {
		t.Fatal(c.c)
	}
}
