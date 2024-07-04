package logic_test

import (
	"github.com/agoblet/pants-go-cyclic-dependency/logic"
	"github.com/agoblet/pants-go-cyclic-dependency/utils"
	"testing"
)

func TestLogic(t *testing.T) {
	if utils.Foo() != logic.CoolLogic() {
		t.Error("oops")
	}
}
