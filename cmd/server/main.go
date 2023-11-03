package main

import (
	"fmt"
	"log"

	"github.com/Krynegal/numeral-system-translator.git/config"
	"github.com/Krynegal/numeral-system-translator.git/internal/converter"
	"github.com/Krynegal/numeral-system-translator.git/internal/handlers"
)

func main() {
	cfg := config.New()

	conv := converter.New()

	r := handlers.NewRouter(conv)

	err := r.Run(fmt.Sprintf("%s:%s", cfg.ServerAddr, cfg.ServerPort))

	if err != nil {
		log.Println(err)
	}
}
