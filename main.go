package main

import (
	"github.com/hpcloud/tail"
	"log"
	"strings"
	"os"
	"github.com/cactus/go-statsd-client/statsd"
)

func main() {

	statsdAddress := getenv("STATSD_ADDRESS", "127.0.0.1:8125")
	statsdPrefix := getenv("STATSD_PREFIX", "fail2ban")

	client, err := statsd.NewClient(statsdAddress, statsdPrefix)

	if err != nil{
		log.Fatal("Error creating statsd statsd", err)
		return
	}

	t, err := tail.TailFile("/logs/fail2ban.log", tail.Config{Follow: true})

	if err != nil {
		log.Fatal("Error reading file.", err)
		return
	}

	for line := range t.Lines {
		if strings.Contains(line.Text, "Ban") {
			client.Inc("Ban", 1, 1.0)
		} else if strings.Contains(line.Text, "Unban") {
			client.Inc("Unban", 1, 1.0)
		}
	}

}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}