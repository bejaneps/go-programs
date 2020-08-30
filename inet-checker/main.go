package main

import (
	"log"
	"net/http"
	"os/exec"
	"time"
)

func main() {
	for {
		log.Printf("Checking internet connection...")
		_, err := http.Get("https://google.com")
		if err != nil {
			log.Printf("No internet connection yet =(")
			time.Sleep(2 * time.Second)
		} else {
			log.Printf("Found internet connection !")
			cmd := exec.Command("xdg-open", "song1.opus")
			if err := cmd.Run(); err != nil {
				log.Fatal(err)
			}
			break
		}
	}
}
