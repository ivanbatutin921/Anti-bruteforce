package main

import(
	"github.com/ivanbatutin921/Anti-bruteforce/mk/core/app"
)

func main() {
	a := app.NewApp()
	err := a.Run()
	if err != nil {
		panic(err)
	}
}