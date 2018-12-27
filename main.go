package main

import "fmt"

func main() {
	console = new(cli)
	console.prompt = "5G-sim > "
	err := console.initMemory()
	if err != nil {
		fmt.Println("Unable to init CLI module")
	} else {
		console.addCommand(helpCmd)
		console.addCommand(shCmd)
		console.addCommand(quitCmd)
		console.addCommand(nfCmd)
		console.addCommand(loadCfgCmd)
		console.runner()
	}
}
