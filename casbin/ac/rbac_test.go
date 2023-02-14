package ac

import (
	"github.com/casbin/casbin/v2"
	"testing"
)

func TestRbacAuthorize(t *testing.T) {
	e, err := casbin.NewEnforcer()
	if err != nil {
		t.Fatal(err)
	}

}
