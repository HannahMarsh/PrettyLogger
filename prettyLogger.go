package PrettyLogger

import "fmt"

func WrapError(err error, msg string) error {
	return err
}

func PrintError(err error, ms string) {
	fmt.Printf("%s: %+v", ms, err)
}
