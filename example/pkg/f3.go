package pkg

import "github.com/HannahMarsh/PrettyLogger"

func F3() error {
	err3 := f2()
	return PrettyLogger.WrapError(err3, "f3 message")
}
