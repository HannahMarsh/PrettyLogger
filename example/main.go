package main

import (
	"encoding/base64"
	"fmt"
	"github.com/HannahMarsh/PrettyLogger"
	"github.com/HannahMarsh/PrettyLogger/example/pkg"
	"log/slog"
)

type Example struct {
	Name string
}

func (e *Example) FuncA(param int, strParam string) {
	fmt.Println(PrettyLogger.GetFuncName(param, strParam))
	fmt.Println(PrettyLogger.GetFuncName())
	fmt.Println(PrettyLogger.GetFuncName(param, strParam))
}

func FuncB(param int, strParam string) {
	fmt.Println(PrettyLogger.GetFuncName(param, strParam))
	fmt.Println(PrettyLogger.GetFuncName())
	slog.Info("FuncB", "param", param, "strParam", strParam)
	//fmt.Println(PrettyLogger.GetFunctionNameAndParameters2(e.FuncA, param, strParam))
}

func main() {

	fmt.Printf(PrettyLogger.ColorRed("Red:\t\t") + PrettyLogger.Bold(PrettyLogger.ColorRed("Bold\t")) + PrettyLogger.ColorRed(PrettyLogger.Italic("Italic\t")) + PrettyLogger.ColorRed(PrettyLogger.Underline("Underline\n")))

	fmt.Printf(PrettyLogger.ColorBrightRed("Bright Red:\t") + PrettyLogger.Bold(PrettyLogger.ColorBrightRed("Bold\t")) + PrettyLogger.ColorBrightRed(PrettyLogger.Italic("Italic\t")) + PrettyLogger.ColorBrightRed(PrettyLogger.Underline("Underline\n")))

	fmt.Printf(PrettyLogger.ColorBrightYellow("Bright Yellow:\t") + PrettyLogger.Bold(PrettyLogger.ColorBrightYellow("Bold\t")) + PrettyLogger.ColorBrightYellow(PrettyLogger.Italic("Italic\t")) + PrettyLogger.ColorBrightYellow(PrettyLogger.Underline("Underline\n")))

	fmt.Printf(PrettyLogger.ColorYellow("Yellow:\t\t") + PrettyLogger.Bold(PrettyLogger.ColorYellow("Bold\t")) + PrettyLogger.ColorYellow(PrettyLogger.Italic("Italic\t")) + PrettyLogger.ColorYellow(PrettyLogger.Underline("Underline\n")))

	fmt.Printf(PrettyLogger.ColorGreen("Green:\t\t") + PrettyLogger.Bold(PrettyLogger.ColorGreen("Bold\t")) + PrettyLogger.ColorGreen(PrettyLogger.Italic("Italic\t")) + PrettyLogger.ColorGreen(PrettyLogger.Underline("Underline\n")))

	fmt.Printf(PrettyLogger.ColorBrightGreen("Bright Green:\t") + PrettyLogger.Bold(PrettyLogger.ColorBrightGreen("Bold\t")) + PrettyLogger.ColorBrightGreen(PrettyLogger.Italic("Italic\t")) + PrettyLogger.ColorBrightGreen(PrettyLogger.Underline("Underline\n")))

	fmt.Printf(PrettyLogger.ColorCyan("Cyan:\t\t") + PrettyLogger.Bold(PrettyLogger.ColorCyan("Bold\t")) + PrettyLogger.ColorCyan(PrettyLogger.Italic("Italic\t")) + PrettyLogger.ColorCyan(PrettyLogger.Underline("Underline\n")))

	fmt.Printf(PrettyLogger.ColorBrightCyan("Bright Cyan:\t") + PrettyLogger.Bold(PrettyLogger.ColorBrightCyan("Bold\t")) + PrettyLogger.ColorBrightCyan(PrettyLogger.Italic("Italic\t")) + PrettyLogger.ColorBrightCyan(PrettyLogger.Underline("Underline\n")))

	fmt.Printf(PrettyLogger.ColorBlue("Blue:\t\t") + PrettyLogger.Bold(PrettyLogger.ColorBlue("Bold\t")) + PrettyLogger.ColorBlue(PrettyLogger.Italic("Italic\t")) + PrettyLogger.ColorBlue(PrettyLogger.Underline("Underline\n")))

	fmt.Printf(PrettyLogger.ColorBrightBlue("Bright Blue:\t") + PrettyLogger.Bold(PrettyLogger.ColorBrightBlue("Bold\t")) + PrettyLogger.ColorBrightBlue(PrettyLogger.Italic("Italic\t")) + PrettyLogger.ColorBrightBlue(PrettyLogger.Underline("Underline\n")))

	fmt.Printf(PrettyLogger.ColorPurple("Purple:\t\t") + PrettyLogger.Bold(PrettyLogger.ColorPurple("Bold\t")) + PrettyLogger.ColorPurple(PrettyLogger.Italic("Italic\t")) + PrettyLogger.ColorPurple(PrettyLogger.Underline("Underline\n")))

	fmt.Printf(PrettyLogger.ColorPink("Pink:\t\t") + PrettyLogger.Bold(PrettyLogger.ColorPink("Bold\t")) + PrettyLogger.ColorPink(PrettyLogger.Italic("Italic\t")) + PrettyLogger.ColorPink(PrettyLogger.Underline("Underline\n")))

	fmt.Printf(PrettyLogger.ColorBrightWhite("Bright White:\t") + PrettyLogger.Bold(PrettyLogger.ColorBrightWhite("Bold\t")) + PrettyLogger.ColorBrightWhite(PrettyLogger.Italic("Italic\t")) + PrettyLogger.ColorBrightWhite(PrettyLogger.Underline("Underline\n")))
	fmt.Printf("White:\t\t" + PrettyLogger.Bold("Bold\t") + PrettyLogger.Italic("Italic\t") + PrettyLogger.Underline("Underline\n"))

	fmt.Printf(PrettyLogger.ColorGrey("Grey:\t\t") + PrettyLogger.Bold(PrettyLogger.ColorGrey("Bold\t")) + PrettyLogger.ColorGrey(PrettyLogger.Italic("Italic\t")) + PrettyLogger.ColorGrey(PrettyLogger.Underline("Underline\n")))

	fmt.Printf(PrettyLogger.ColorBlack("Black:\t\t") + PrettyLogger.Bold(PrettyLogger.ColorBlack("Bold\t")) + PrettyLogger.ColorBlack(PrettyLogger.Italic("Italic\t")) + PrettyLogger.ColorBlack(PrettyLogger.Underline("Underline\n")))

	PrettyLogger.SetUpLogrusAndSlog("debug")

	err2 := pkg.F3()

	slog.Error("main error", err2)
	slog.Info("main info", "var", "info")
	slog.Warn("main warn", "warn", "warn")
	slog.Debug("main debug", "debug", "debug")

	slog.Info("this is a test", "var", "value")

	type mystr struct {
		Value int
		Name  string
	}
	var m mystr = mystr{Value: 5, Name: "name"}
	slog.Info("Write a little message", "latency", 0.345, "status", "ok", "value", 123, "bool", true, "mystr", m, "null", nil)

	e := &Example{Name: "example"}
	e.FuncA(4, "examplestring")

	FuncB(5, "examplestring")

	_, err := base64.StdEncoding.DecodeString("test string")
	if err != nil {
		slog.Error("error decoding string", err)
	}

}
