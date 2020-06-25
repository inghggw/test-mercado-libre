package main

import (
	"os"
	"test_mercado_libre/httpd/handler"
)

func main() {
	os.Setenv("api_mercado_libre", "https://api.mercadolibre.com/items/")

	handler.Routes().Run(":8800")
}
