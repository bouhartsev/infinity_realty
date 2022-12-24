package main

import (
	"github.com/bouhartsev/infinity_realty/internal/app"
)

// @title Esoft API
// @version 1.0
// @accept json
// @produce json

func main() {
	apl, err := app.New()
	if err != nil {
		panic(err)
	}

	if err = apl.Run(); err != nil {
		panic(err)
	}
}
