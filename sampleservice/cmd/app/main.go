package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"

	"github.com/mmqbaba/zeus/plugin/container"
	"github.com/mmqbaba/zeus/service"

	"github.com/mmqbaba/zeus-examples/sampleservice/global"
)

var (
	BuildDate = ""
	Version   = ""
	GoVersion = ""
)

func main() {
	args := os.Args
	if len(args) == 2 && (args[1] == "--version" || args[1] == "-version" || args[1] == "-v") {
		log.Println("-------------------------")
		log.Printf("Git Commit Hash: %s\n", Version)
		log.Printf("UTC Build Date : %s\n", BuildDate)
		log.Printf("Golang Version : %s\n", GoVersion)
		return
	}

	log.Println("----------version info----------")
	log.Printf("Git Commit Hash: %s\n", Version)
	log.Printf("UTC Build Date : %s\n", BuildDate)
	log.Printf("Golang Version : %s\n", GoVersion)
	log.Println("--------------------------------")
	fmt.Print("\n")

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
	runtime.GOMAXPROCS(num)

	log.Println("service run ...")
	if err := service.Run(container.GetContainer(), nil, global.ServiceOpts...); err != nil {
		log.Printf("Service exited with error: %s\n", err)
		os.Exit(255)
	} else {
		log.Println("Service exited gracefully")
	}
}
