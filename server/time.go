package main

import (
	"mule/overpower"
	"time"
)

func AutoTimer() {
	for {
		now := time.Now()
		then := now.Round(time.Hour)
		if hour := then.Hour(); hour > 22 {
			then = then.Add(time.Hour * time.Duration(23+(24-hour)))
		} else {
			then = then.Add(time.Hour * time.Duration(23-hour))
		}
		then = then.Add(time.Minute * time.Duration(2))
		Announce("Starting autotimer:", now, "\n SLEEP TILL:", then)
		dur := then.Sub(now)
		time.Sleep(dur)
		now = time.Now()
		if now.Hour() != 23 {
			ErrLogger.Println("AUTO RUN ERROR: SLEEP DID NOT REACH HOUR 23")
			continue
		} else {
			Announce("Autotimer woke:", now)
		}
		DBLOCK = true
		time.Sleep(5 * time.Minute)
		games, err := OPDB.GetGames()
		if my, bad := Check(err, "resource failure in autotimer", "resource", "games"); bad {
			Log(my)
			continue
		}
		var count int
		countChan := make(chan byte)
		wkDay := int(now.Weekday())
		toUpdate := make([]overpower.Game, 0)
		for _, g := range games {
			if g.Turn() < 1 {
				continue
			}
			days := g.AutoDays()
			if days[wkDay] {
				if free := g.FreeAutos(); free > 0 {
					g.SetFreeAutos(free - 1)
					toUpdate = append(toUpdate, g)
				} else {
					count++
					go func(g overpower.Game, done chan byte) {
						Announce("AUTO RUNNING GAME", g.Gid())
						err := RunGameTurn(g.Gid())
						if my, bad := Check(err, "autorun game turn failure", "game", g); bad {
							Log(my)
						}
						done <- 0
					}(g, countChan)
				}
			}
		}
		if len(toUpdate) > 0 {
			err = OPDB.UpdateGames(toUpdate...)
			if my, bad := Check(err, "autorun game freeturn inc failure", "games", toUpdate); bad {
				Log(my)
			}
		}
		for count > 0 {
			<-countChan
			count -= 1
		}
		DBLOCK = false
	}
}
