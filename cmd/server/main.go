package main

import (
	"fmt"
	"log"

	"github.com/valyala/fasthttp"
)

func main() {
	fmt.Println("WebAssembly Run!")
	log.Fatal(fasthttp.ListenAndServe(":8080", fasthttp.FSHandler("../../assets", 0)))
}
