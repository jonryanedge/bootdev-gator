package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/jonryanedge/bootdev-gator/internal/config"
)

var user = "jonryanedge"

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}

	if err := cfg.SetUser(user); err != nil {
		log.Fatalf("Failed to write config: %v", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("Failed to read updated config: %v", err)
	}

	fmt.Println("Config successfully updated and reloaded!")
	fmt.Println("Current config contents:")

	prettyJSON, _ := json.MarshalIndent(cfg, "", "    ")
	fmt.Println(string(prettyJSON))

	fmt.Printf("Current user: %s\n", cfg.CurrentUserName)
}
