package main

import (
	"fmt"
	"github.com/HannahMarsh/PrettyLogger"
	"github.com/HannahMarsh/PrettyLogger/example/pkg"
	"golang.org/x/exp/slog"
)

type Example struct {
	Name string
}

func (e *Example) FuncA(param int, strParam string) {
	fmt.Println(PrettyLogger.GetFuncName(param, strParam))
	fmt.Println(PrettyLogger.GetFuncName())
	//fmt.Println(PrettyLogger.GetFunctionNameAndParameters2(e.FuncA, param, strParam))
}

func FuncB(param int, strParam string) {
	fmt.Println(PrettyLogger.GetFuncName(param, strParam))
	fmt.Println(PrettyLogger.GetFuncName())
	//fmt.Println(PrettyLogger.GetFunctionNameAndParameters2(e.FuncA, param, strParam))
}

func main() {

	PrettyLogger.SetUpLogrusAndSlog("debug")

	err2 := pkg.F3()

	slog.Error("main error", err2)
	slog.Info("main info", "var", "info")
	slog.Warn("main warn", "warn", "warn")
	slog.Debug("main debug", "debug", "debug")

	slog.Info("this is a test", "var", "value")

	e := &Example{Name: "example"}
	e.FuncA(4, "examplestring")

	FuncB(5, "examplestring")

}
