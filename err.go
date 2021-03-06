package hcutil

import "fmt"

type Err struct {
	Pkg  string
	Info string
	Err  error
}

func (e *Err) Error() string {
	return fmt.Sprintf("%s: %s\n%v", e.Pkg, e.Info, e.Err)
}

func makeErr(err error, info string) *Err {
	return &Err{
		Pkg:  `hcutil`,
		Info: info,
		Err:  err,
	}
}
