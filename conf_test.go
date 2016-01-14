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
