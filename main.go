package main

import (
	"os"
	"time"
	"fmt"

	"github.com/azer/logger"
	"gopkg.in/alecthomas/kingpin.v2"
	"github.com/itsTurnip/dishooks"
)

var (
	app      = kingpin.New("Cryson Ban Bot", "Discord Ban Bot to post Cryson Ban Status").Author("Todd")
	discordaddr = app.Flag("discord-webhook", "Webhookaddress").Envar("DISCORD_WEBHOOK").Required().String()
	csvhour = app.Flag("csvhour", "Set the hour in which the script will run").Envar("CSVHOUR").Default("7").Int()
	csvminute = app.Flag("csvminute", "Set the minute in which the script will run").Envar("CSVMINUTE").Default("0").Int()
	deadline = app.Flag("deadline", "deadline till unban").Envar("DEADLINE").Default("2020-08-18T20:23:00+01:00").String()
	log = logger.New("DiscordLog")
)

type countdown struct {
	t int
	d int
	h int
	m int
	s int
}

func getTimeRemaining(t time.Time) countdown {
	currentTime := time.Now()
	difference := t.Sub(currentTime)

	total := int(difference.Seconds())
	days := int(total / (60 * 60 * 24))
	hours := int(total / (60 * 60) % 24)
	minutes := int(total/60) % 60
	seconds := int(total % 60)

	return countdown{
		t: total,
		d: days,
		h: hours,
		m: minutes,
		s: seconds,
	}
}

func initWait() {
	t := time.Now()
	n := time.Date(t.Year(), t.Month(), t.Day(), *csvhour, *csvminute, 0, 0, t.Location())
	d := n.Sub(t)
	if d < 0 {
		log.Info("Waiting 24 hours to run again")
		log.Info("Will run again at hour %d:%.2d", *csvhour, *csvminute)
		n = n.Add(24 * time.Hour)
		d = n.Sub(t)
	}
	for {
		time.Sleep(d)
		d = 24 * time.Hour
		log.Info("Executing Discord Webhook call")
		v, err := time.Parse(time.RFC3339, *deadline)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		timeRemaining := getTimeRemaining(v)
		if timeRemaining.t <= 0 {
			webhook, _ := dishooks.WebhookFromURL(*discordaddr)
			_ = webhook.Get()
			_, _ = webhook.SendMessage(&dishooks.WebhookMessage{
				Content: "Cryson is unbanned!" ,
			})
		}

		unbantime := fmt.Sprintf("%d Days %d Hours %d Minutes %d Seconds", timeRemaining.d, timeRemaining.h, timeRemaining.m, timeRemaining.s)
		webhook, _ := dishooks.WebhookFromURL(*discordaddr)
		_ = webhook.Get()
		_, _ = webhook.SendMessage(&dishooks.WebhookMessage{
			Content: "Cryson will be unbanned in:\n" + unbantime ,
		})
	}
}


func main() {
	// Parse kingpin flags
	kingpin.MustParse(app.Parse(os.Args[1:]))

	// Initialize the wait routine and webhooks
	log.Info("Starting DiscordBanBot")
	initWait()
}
