package main

import (
	"fmt"
	"github.com/Krynegal/numeral-system-translator.git/config"
	"github.com/Krynegal/numeral-system-translator.git/internal/handlers"
)

func main() {
	cfg := config.New()
	r := handlers.NewRouter()

	err := r.Run(fmt.Sprintf("%s:%s", cfg.ServerAddr, cfg.ServerPort))

	if err != nil {
		return
	}
}
