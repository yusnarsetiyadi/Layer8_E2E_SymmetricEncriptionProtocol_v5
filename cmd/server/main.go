package main

import (
	"log"

	"github.com/valyala/fasthttp"
)

func main() {
	log.Fatal(fasthttp.ListenAndServe(":8080", fasthttp.FSHandler("../../assets", 0)))
}
