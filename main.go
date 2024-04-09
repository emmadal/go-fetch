package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"github.com/emmadal/go-fetch/pkg"
)

func main() {
	startTime := time.Now()
	repl := os.Args

	if len(repl) == 1 {
		log.Fatalln("missing URL for downloading.")
	}
	if len(repl) > 2 {
		log.Fatalln("Too much URL. please keep only one URL")
	}

	pkg.Fetcher(os.Args[1])
	fmt.Printf("🍺 Executed in %f sec\n", time.Since(startTime).Seconds())
}
