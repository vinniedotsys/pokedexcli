package main


func main() {
	cfg := &config{
		commands: commandsList(),
	}

	startRepl(cfg)
}
