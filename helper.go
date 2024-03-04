package gocalc

import "fmt"

func catch(err *error) {
	if r := recover(); r != nil {
		*err = fmt.Errorf("%v", r)
	}
}
