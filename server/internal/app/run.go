package app

import "github.com/bouhartsev/infinity_realty/internal/server"

func (a *application) Run() error {
	s, err := server.New(a.logger, a.cfg)

	if err != nil {
		return err
	}

	return s.Run()
}
