package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/jonryanedge/bootdev-gator/internal/config"
	"github.com/jonryanedge/bootdev-gator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DbUrl)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	dbQueries := database.New(db)

	s := &state{
		db:  dbQueries,
		cfg: &cfg,
	}

	cmds := &commands{
		handlers: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerResetUsers)
	cmds.register("users", handlerGetUsers)

	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatalf("Not enough arguments. Usage: gator <command> [args...]")
	}

	cmdName := args[0]
	cmdArgs := args[1:]

	cmd := command{
		name: cmdName,
		args: cmdArgs,
	}

	if err := cmds.run(s, cmd); err != nil {
		log.Fatalf("Error: %v", err)
	}

	// if err := cfg.SetUser(user); err != nil {
	// 	log.Fatalf("Failed to write config: %v", err)
	// }
	//
	// cfg, err = config.Read()
	// if err != nil {
	// 	log.Fatalf("Failed to read updated config: %v", err)
	// }
	//
	// fmt.Println("Config successfully updated and reloaded!")
	// fmt.Println("Current config contents:")
	//
	// prettyJSON, _ := json.MarshalIndent(cfg, "", "    ")
	// fmt.Println(string(prettyJSON))
	//
	// fmt.Printf("Current user: %s\n", cfg.CurrentUserName)
}
