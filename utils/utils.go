package utils

import (
	"github.com/agoblet/pants-go-cyclic-dependency/logic"
)

func Foo() bool {
	return logic.CoolLogic()
}
