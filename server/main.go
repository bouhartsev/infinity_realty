package main

import (
	"github.com/bouhartsev/infinity_realty/internal/app"
)

func main() {
	apl, err := app.New()
	if err != nil {
		panic(err)
	}

	if err = apl.Run(); err != nil {
		panic(err)
	}
}
