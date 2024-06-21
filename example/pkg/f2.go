package pkg

import "github.com/HannahMarsh/PrettyLogger"

func f2() error {
	err1 := f1()
	return PrettyLogger.WrapError(err1, "f2 message")
}
