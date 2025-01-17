package main

import (
	inf "github.com/gozelle/infinite"
	"github.com/gozelle/infinite/components"
	"github.com/gozelle/infinite/components/spinner"
	"time"
)

func main() {
	_ = inf.NewSpinner(
		spinner.WithShape(components.Dot),
		//spinner.WithDisableOutputResult(),
	).Display(func(spinner *spinner.Spinner) {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Millisecond * 100)
			spinner.Refreshf("hello world %d", i)
		}
		
		spinner.Finish("finish")
		
		spinner.Refresh("is finish?")
	})
}
