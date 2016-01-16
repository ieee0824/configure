package configure

import (
	"io/ioutil"
	"testing"
)

var includeList = []string{
	"conf3.json",
	"./conf3.json",
	"../test/conf3.json",
}

func isExist(list []string, s string) bool {
	for _, v := range list {
		if v == s {
			return true
		}
	}
	return false
}

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
	if s, ok := NewConf("./test/conf.json").c["hoge"].(string); ok {
		if s != "hage" {
			t.Log(ok, "\n")
			t.Log(NewConf("./test/conf.json"), "\n")
			t.Log(NewConf("./test/conf.json").c["hoge"], "\n")
			t.Log(s, "\n")
			t.Fail()
		}
	}
	if s, ok := NewConf("./test/conf.json").c["hoge"].(string); !ok {
		if s == "hage" {
			t.Log(ok, "\n")
			t.Log(NewConf("./test/conf.json"), "\n")
			t.Log(NewConf("./test/conf.json").c["hoge"], "\n")
			t.Log(s, "\n")
			t.Fail()
		}
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

func TestGetKeys(t *testing.T) {
	t.Log("TestGetKeys\n")
	normal := NewConf("./test/conf.json")
	if len(normal.GetKeys()) == 0 {
		t.Log("can not get keys.\n")
		t.Log("length is zero.")
		t.Fail()
	}
	if func() bool {
		for _, key := range normal.GetKeys() {
			if key == "hoge" {
				return false
			}
		}
		return true
	}() {
		t.Log(normal.GetKeys(), "\n")
		t.Fail()
	}
	abnormal := NewConf("./test/conf2.json")
	if len(abnormal.GetKeys()) != 0 {
		t.Log("can not get keys.\n")
		t.Log("length is not zero.")
		t.Fail()
	}
}

func TestGetIncludeFileNames(t *testing.T) {
	t.Log("TestGetIncludeFileNames\n")
	abnormal := NewConf("./test/conf2.json")
	if len(abnormal.getIncludeFileNames()) != 0 || abnormal.getIncludeFileNames() != nil {
		t.Log(abnormal)
		t.Fail()
	}
}

func TestInclude(t *testing.T) {
	normal := NewConf("./test/conf.json")
	normal.include("./test/conf3.json")
	if normal.Get("loc").(string) != "tokyo" {
		t.Log(normal.Get("loc"))
		t.Fail()
	}
	abnormal := NewConf("./test/conf2.json")
	abnormal.include("./test/conf3.json")
	if _, ok := abnormal.Get("loc").(string); ok {
		t.Log(abnormal.Get("loc"))
		t.Fail()
	}
	t.Fatal(normal)
}
