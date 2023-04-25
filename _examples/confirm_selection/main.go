package main

import (
	"fmt"
	inf "github.com/gozelle/infinite"
	"github.com/gozelle/infinite/components/selection/confirm"
)

func main() {
	
	val, _ := inf.NewConfirmWithSelection(
		//confirm.WithDisableOutputResult(),
		//confirm.WithDisableShowHelp(),
		confirm.WithDefaultYes(),
	).Display()
	
	if val {
		fmt.Println("yes, you are.")
	} else {
		fmt.Println("no,you are not.")
	}
}
