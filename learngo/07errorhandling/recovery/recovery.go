package recovery

import (
	// "errors"
	"fmt"
)

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("error occured: ", err)
		} else {
			panic(fmt.Sprintf("I dont know what to do", r))
		}
	}()
	// a, b := 5, 0
	// fmt.Println(a / b)
	// panic(errors.New("this is an error"))
	panic(123)

}
func TryRecover() {
	tryRecover()
}
