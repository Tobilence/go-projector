package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/tobilence/go-aoc/pkg/projector"
)

func main() {

	opts, err := projector.GetOpts()
	if err != nil {
		log.Fatalf("unable to get options %v", err)
	}
	conf, err := projector.NewConfig(*opts)
	if err != nil {
		log.Fatalf("unable to get the config %v", err)
	}

	proj := projector.FromConfig(conf)

	if conf.Operation == projector.Print {
		if len(conf.Args) == 0 {
			data := proj.GetValueAll()
			jsonString, err := json.Marshal(data)
			if err != nil {
				log.Fatalf("Unexpected Error %v", err)
			}

			fmt.Printf("%v", string(jsonString))
		} else if value, ok := proj.GetValue(conf.Args[0]); ok {
			fmt.Printf("%v", value)
		}
	}

	if conf.Operation == projector.Add {
		proj.SetValue(conf.Args[0], conf.Args[1])
		proj.Save()
	}

	if conf.Operation == projector.Add {
		proj.RemoveValue(conf.Args[0])
		proj.Save()
	}
}
