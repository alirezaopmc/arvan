package main

import (
	"fmt"
	"os"

	"github.com/alirezaopmc/arvan/api"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	arvan := api.NewAPI(
		"https://napi.arvancloud.ir",
		os.Getenv("ARVAN_API_KEY"),
	)
	domains, err := arvan.CDN.GetDomains("", 99, 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(domains[0])
}
