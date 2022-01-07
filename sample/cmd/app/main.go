package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"

	"github.com/mmqbaba/zeus/plugin/container"
	"github.com/mmqbaba/zeus/service"

	"zeus-examples/sample/global"
	_ "zeus-examples/sample/proto"
)

var (
	BuildDate = ""
	Version   = ""
	GoVersion = ""
)

func main() {

	log.Println("----------version info----------")
	log.Printf("Git Commit Hash: %s\n", Version)
	log.Printf("Build Date     : %s\n", BuildDate)
	log.Printf("Golang Version : %s\n", GoVersion)
	log.Println("--------------------------------")
	fmt.Print("\n")

	args := os.Args
	if len(args) == 2 && (args[1] == "--version" || args[1] == "-version" || args[1] == "-v") {
		return
	}

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

	// loadEngineFnOpt := service.WithLoadEngineFnOption(func(ng engine.Engine) {
	// 	log.Println("WithLoadEngineFnOption: SetNG success.")
	// 	global.SetNG(ng)
	// })
	// global.ServiceOpts = append(global.ServiceOpts, loadEngineFnOpt)

	// initServiceFnOpt := service.WithInitServiceCompleteFnOption(func(ng engine.Engine) {
	// 	log.Println("WithInitServiceCompleteFnOption")
	// })
	// global.ServiceOpts = append(global.ServiceOpts, initServiceFnOpt)

	log.Println("service run ...")
	if err := service.Run(container.GetContainer(), nil, global.ServiceOpts...); err != nil {
		log.Printf("Service exited with error: %s\n", err)
		os.Exit(255)
	} else {
		log.Println("Service exited gracefully")
	}
}
