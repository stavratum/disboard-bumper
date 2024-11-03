package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pelletier/go-toml/v2"
	"github.com/stavratum/discordgo-self"
)

var (
	Command = &discordgo.ApplicationCommand{
		ID:            "947088344167366698",
		ApplicationID: "302050872383242240",
		Version:       "1051151064008769576",

		Name: "bump",

		Type: discordgo.ChatApplicationCommand,
	}

	Config = map[string]struct {
		Channels []string
		Token    string
	}{}
)

func main() {
	buf, err := os.ReadFile("disboard-bumper.toml")
	if err != nil {
		log.Panic("couldn't find disboard-bumper.toml in this directory.")
	}

	if err = toml.Unmarshal(buf, &Config); err != nil {
		log.Panic(err)
	}

	suc := false

	for k, ac := range Config {
		if len(ac.Channels) == 0 {
			continue
		}

		session, _ := discordgo.New(ac.Token)
		if err = session.Open(); err != nil {
			log.Printf("[%s] %s", k, err)
			return
		}
		defer session.Close()

		suc = true

		go func() {
			rwd := 120 / len(ac.Channels)
			if rwd < 30 {
				rwd = 30
			}

			delay := time.Minute * time.Duration(rwd+1)
			timer := time.NewTicker(delay)

			log.Printf("[%s] Delay = %s", k, delay)

			for {
				for _, cID := range ac.Channels {
					for range 2 {
						if err = session.ApplicationCommandSend(cID, Command); err == nil {
							log.Printf("[%s > %s]: OK", k, cID)
							break
						}

						log.Printf("[%s > %s]: %s", k, cID, err)
					}

					<-timer.C
				}
			}
		}()
	}

	if !suc {
		return
	}

	stop := make(chan os.Signal, 2)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	<-stop
}
