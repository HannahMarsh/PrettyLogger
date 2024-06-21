package pkg

import (
	"github.com/HannahMarsh/PrettyLogger"
)

func f1() error {
	return PrettyLogger.NewError("f1 %s", "message")
}
