package main

import (
	"fmt"
	"url-shortener/internal/config"
)

func main() {

	cfg := config.MustLoad()

	fmt.Print(cfg)

	//todo: init logger: slog  import log/slog

	//todo: init storage: sqlite

	//todo: init router: chi "chi render"

	//todo: run server

}
