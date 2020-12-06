package main

import (
	"fmt"
	"github.com/rehmus/webscreen"
	"log"
)

var ConfigPath = "config.yml"

func main() {
	conf, err := webscreen.ConfigFromFile(ConfigPath)
	if err != nil {
		log.Fatalln("Could not load config:", err)
	}
	engine, err := webscreen.NewEngineFromConfig(conf)
	if err != nil {
		log.Fatalln("Could not create engine:", err)
	}
	defer engine.Stop()
	r, err := engine.NewRunner()
	if err != nil {
		log.Fatalln("Could not create runner:", err)
	}
	if err := r.SetSize(1296, 852); err != nil {
		log.Fatalln("Could not set size of runner:", err)
	}
	urls := []string{"https://www.google.com", "https://www.github.com/rehmus", "https://golang.org"}
	for i, url := range urls {
		if err := r.Get(url); err != nil {
			log.Printf("error loading url %s: %s\n", url, err)
			continue
		}
		if err := r.Screenshot(fmt.Sprintf("out/image_%d.png", i)); err != nil {
			log.Printf("error capturing screenshot: %s\n", err)
		}
	}
}
