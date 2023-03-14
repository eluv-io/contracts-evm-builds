package main

import (
	"fmt"
	"log"
	"os"

	"github.com/eluv-io/contracts-evm-builds/cmd/eventpkggen/gen"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		log.Fatalf("config file not provided\n")
	}
	cfg, err := gen.NewEventPkgGenConfig(args[1])
	if err != nil {
		log.Fatalf("err:%v\n", err)
	}
	outputDirToTmplStructMap := cfg.BuildTemplateStruct()
	for _, tmplStruct := range outputDirToTmplStructMap {
		err := tmplStruct.GenerateEventPkgFiles()
		if err != nil {
			log.Fatalf("err:%v\n", err)
		}
		fmt.Printf("event package files generated in %v directory...\n", tmplStruct.OutputPath)
	}
}
