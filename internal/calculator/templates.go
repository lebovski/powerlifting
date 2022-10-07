package calculator

import (
	"fmt"
)

type training []*unit

func (t training) String() string {
	res := ""

	for _, u := range t {
		res += fmt.Sprintf(" %s", u)
	}

	return res
}

type unit struct {
	weight float64
	repeat int
	count  int
}

func (u *unit) String() string {
	if u.count == 1 {
		return fmt.Sprintf("%v/%v", u.weight, u.repeat)
	}

	return fmt.Sprintf("%v/%v*%v", u.weight, u.repeat, u.count)
}
