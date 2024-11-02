package main

import (
	"crypto/tls"
	"errors"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"time"

	"github.com/pelletier/go-toml/v2"
)

var (
	EndpointAPI    = "https://disboard.org/"
	EndpointServer = EndpointAPI + "server/"

	EndpointServerEdit = func(ID string) string { return EndpointServer + "edit/" + ID }
	EndpointServerBump = func(ID string) string { return EndpointServer + "bump/" + ID }
)

type Account struct {
	Servers []string
	Cookies string
}

var Config = map[string]*Account{}

func init() {
	for _, fn := range []string{"disboard.toml", "bumper.toml", "disboard-bumper.toml"} {
		_, err := os.Stat(fn)
		if errors.Is(err, os.ErrNotExist) {
			continue
		}

		buffer, err := os.ReadFile(fn)
		if err != nil {
			log.Println("Failed to read", fn)
			continue
		}

		toml.Unmarshal(buffer, &Config)
		break
	}
}

var Client = &http.Client{
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	},
}

func init() {
	Client.Jar, _ = cookiejar.New(nil)
}

var Retries = 2

func worker(name string, cfg *Account) {
	rwd := 120 / len(cfg.Servers)
	if rwd < 30 {
		rwd = 30
	}

	delay := time.Minute * time.Duration(rwd)
	timer := time.NewTicker(delay)

	log.Printf("[%s] Delay = %s", name, delay)

	for {
		for _, ID := range cfg.Servers {
			for range Retries {
				resp, err := Client.Get(EndpointServerBump(ID))

				if err == nil {
					log.Printf("[%s] >> [%s] %s", name, ID, resp.Status)
					break
				} else {
					log.Printf("[%s] >> [%s] %s", name, ID, err.Error())
					time.Sleep(time.Minute)
				}
			}

			<-timer.C
		}
	}
}

func main() {
	for key, cfg := range Config {
		if len(cfg.Servers) == 0 {
			continue
		}

		cookies, err := http.ParseCookie(cfg.Cookies)
		if err != nil {
			log.Println("[", key, "] Err parsing cookies: ", err)
			return
		}

		for _, ID := range cfg.Servers {
			endpoint, err := url.Parse(EndpointServerBump(ID))
			if err != nil {
				panic(err)
			}

			Client.Jar.SetCookies(endpoint, cookies)
		}

		go worker(key, cfg)
	}

	<-make(chan struct{})
}
