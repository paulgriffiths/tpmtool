package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	log.SetPrefix(fmt.Sprintf("%s: ", appName))
	log.SetFlags(0)

	if len(os.Args) < 2 {
		usageError()
	}

	for _, cmd := range commands {
		if os.Args[1] == cmd.name {
			if cmd.flagSet != nil {
				cmd.flagSet.Parse(os.Args[2:])
			}

			if isFlagPassed(cmd.flagSet, "help") {
				cmd.usageFunc()
			} else {
				if err := cmd.cmdFunc(); err != nil {
					log.Fatalf("%v", err)
				}
			}

			return
		}
	}

	usageError()
}
