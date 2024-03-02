package main

import (
	"os"
	"wl-optimizer/core"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		core.PrintHelp()
		os.Exit(0)
	}

	parsedArgs, inputFilepath, outputFilePath := core.ParseArgs(args)

	var optimizer = core.WlOptimizer{}

	optimizer.Init(inputFilepath, outputFilePath, parsedArgs)

	optimizer.OptimizeWordlist()

}
