package ac

import (
	"github.com/casbin/casbin/v2"
	"path"
	"runtime"
	"testing"
)

var (
	modelConf = "/conf/acl_model.conf"
	policyCsv = "/conf/acl_policy.csv"
)

func getWd() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}

func init() {
	wd := path.Dir(getWd())
	modelConf = wd + modelConf
	policyCsv = wd + policyCsv
}

func TestAclAuthorize(t *testing.T) {
	e, err := casbin.NewEnforcer(modelConf, policyCsv)
	if err != nil {
		t.Fatalf("create enforcer failed: %s", err)
	}

	sub := "alice"
	obj := "data1"
	//act := "read"
	act := "write"
	ok, err := e.Enforce(sub, obj, act)
	if err != nil {
		t.Fatalf("enforce failed: %s", err)
	}
	if ok == true {
		t.Log("Approved!")
	} else {
		t.Log("Rejected!")
	}

	results, _ := e.BatchEnforce([][]interface{}{{"alice", "data1", "read"}, {"bob", "data2", "write"}, {"jack", "data3", "read"}})
	t.Log("results:", results)
}
