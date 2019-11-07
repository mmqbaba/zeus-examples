package main

import (
	"log"
	"os"
	"runtime"
	"strconv"

	"gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/plugin/container"
	"gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/service"
)

func main() {
	num := runtime.NumCPU()
	log.Printf("[NumCPU] %v\n", num)
	gmp := os.Getenv("GOMAXPROCS")
	if gmp != "" {
		r, e := strconv.Atoi(gmp)
		if e == nil && r < num && r > 0 {
			num = r
		}
	}
	log.Printf("[GOMAXPROCS] %v\n", num)
	curr := runtime.GOMAXPROCS(num)
	log.Printf("[CURRENT GOMAXPROCS] %v\n", curr)
	log.Println("service run ...")
	if err := service.Run(container.GetContainer()); err != nil {
		log.Printf("Service exited with error: %s\n", err)
		os.Exit(255)
	} else {
		log.Println("Service exited gracefully")
	}
}
