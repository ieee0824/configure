package configure

import "testing"

func TestNewConf(t *testing.T) {
	c := NewConf("./test/conf.json")

	if c == nil {
		t.Fatal("can not read configure")
	}
}
