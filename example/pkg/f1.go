package pkg

import "github.com/pkg/errors"

func f1() error {
	return errors.New("f1 error")
}
