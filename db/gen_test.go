package db

import (
	"log"
	"mule/hexagon"
	"mule/mybad"
	"mule/overpower"
	"testing"
)

func Test1(t *testing.T) {
	log.Println("TEST ONE")
}

func Test4(t *testing.T) {
	db, err := LoadDB()
	if my, bad := Check(err, "test2 loaddb fail"); bad {
		log.Println(my.MuleError())
		return
	}
	log.Println("DB loaded")
	// -------------------- //
	game, err := db.GetGame("owner", "AutoTest")
	if my, bad := Check(err, "test2 getgame fail", "owner", "AutoTest"); bad {
		if !my.BaseIs(ErrNoneFound) {
			log.Println(my.MuleError())
			return
		}
		log.Println("game not found")
		err = db.MakeGame("AutoTest", "Automade", "", 40)
		if my, bad := Check(err, "test2 automake game fail"); bad {
			log.Println(my.MuleError())
			return
		}
		log.Println("Made auto-test game")
		game, err = db.GetGame("owner", "AutoTest")
		if my, bad := Check(err, "test2 getgame fail two", "owner", "AutoTest"); bad {
			log.Println(my.MuleError())
			return
		}
	}
	log.Println("Looking at game:", game)
	gid := game.Gid()
	if 0 == 0 {
		err = db.DropGames("gid", gid)
		if my, bad := Check(err, "test2 failed dropgame", "gid", gid); bad {
			log.Println(my.MuleError())
			return
		}
		log.Println("Dropped game", gid)
		err = db.MakeGame("AutoTest", "Automade", "", 40)
		if my, bad := Check(err, "test2 automake game fail"); bad {
			log.Println(my.MuleError())
			return
		}
		log.Println("Made auto-test game")
		game, err = db.GetGame("owner", "AutoTest")
		if my, bad := Check(err, "test2 getgame fail two", "owner", "AutoTest"); bad {
			log.Println(my.MuleError())
			return
		}
		gid = game.Gid()
	}
	for _, name := range []string{
		"test1", "test2", "test3",
		"test4", "test5", "test6",
		"test7", "test8", "test9",
		"mule",
	} {
		fac, err := db.GetFaction("gid", gid, "owner", name)
		if my, bad := Check(err, "test2 getfaction failure"); bad {
			if !my.BaseIs(ErrNoneFound) {
				log.Println(my.MuleError())
				return
			}
			log.Println("fac", name, "not found")
			err = db.MakeFaction(gid, name, name)
			if my, bad := Check(err, "test2 faction make failure"); bad {
				log.Println(my.MuleError())
				return
			}
			log.Println("Made faction", name)
			fac, err = db.GetFaction("gid", gid, "owner", name)
			if my, bad := Check(err, "test2 getfaction failure"); bad {
				log.Println(my.MuleError())
				return
			}
		}
		_ = fac
		log.Println("Got faction:", fac)
	}
	planetViews, err := db.GetPlanetViews("gid", gid)
	if my, bad := Check(err, "test2 get planetviews failure"); bad {
		log.Println(my.MuleError())
		return
	}
	log.Println("Got planetviews:", len(planetViews))
	if len(planetViews) == 0 {

		f := func(d DB) error {
			source := d.NewSource(gid)
			err := overpower.MakeGalaxy(source)
			if my, bad := Check(err, "make galaxy failed", "gid", gid); bad {
				return my
			}
			err = source.Commit()
			if my, bad := Check(err, "make galaxy turn commit failed", "gid", gid); bad {
				return my
			}
			return nil
		}

		err = db.Transact(f)
		if my, bad := Check(err, "transaction error in test3"); bad {
			log.Println(my.MuleError())
			return
		}
		log.Println("MADE GALAXY!!")
	} else {

		f := func(d DB) error {
			source := d.NewSource(gid)
			breakE, logE := overpower.RunGameTurn(source, true)
			if logE != nil {
				log.Println(logE.(*mybad.MuleError).MuleError())
			}
			if my, bad := Check(breakE, "rungameturn failed", "gid", gid); bad {
				return my
			}
			err := source.Commit()
			if my, bad := Check(err, "rungame turn commit failed", "gid", gid); bad {
				return my
			}
			return nil
		}
		err = db.Transact(f)
		if my, bad := Check(err, "transaction error in test3"); bad {
			log.Println(my.MuleError())
			return
		}
		log.Println("Ran turn")
	}

	log.Println(" // -------------------------------- //")
	log.Println("PLANETVIEWS")
	for i, plv := range planetViews {
		if plv.Turn() > 0 {
			log.Print(i, ": ", plv, "\n")
		}
	}
	// -------------------------------- //
	return
	err = db.DropGames("gid", gid)
	if my, bad := Check(err, "test2 failed dropgame", "gid", gid); bad {
		log.Println(my.MuleError())
		return
	}
	log.Println("Dropped game", gid)
}

func XTest2(t *testing.T) {
	db, err := LoadDB()
	if my, bad := Check(err, "test2 loaddb fail"); bad {
		log.Println(my.MuleError())
		return
	}
	log.Println("DB loaded")
	// -------------------- //
	game, err := db.GetGame("owner", "AutoTest")
	if my, bad := Check(err, "test2 getgame fail", "owner", "AutoTest"); bad {
		if !my.BaseIs(ErrNoneFound) {
			log.Println(my.MuleError())
			return
		}
		err = db.MakeGame("AutoTest", "Automade", "", 40)
		if my, bad := Check(err, "test2 automake game fail"); bad {
			log.Println(my.MuleError())
			return
		}
		log.Println("Made auto-test game")
		game, err = db.GetGame("owner", "AutoTest")
		if my, bad := Check(err, "test2 getgame fail two", "owner", "AutoTest"); bad {
			log.Println(my.MuleError())
			return
		}
	}
	log.Println("Looking at game:", game)
	gid := game.Gid()
	if 0 == 1 {
		err = db.DropGames("gid", gid)
		if my, bad := Check(err, "test2 failed dropgame", "gid", gid); bad {
			log.Println(my.MuleError())
			return
		}
		log.Println("Dropped game", gid)
		return
	}
	err = db.MakeFaction(gid, "AutoTest", "AutoFaction")
	if my, bad := Check(err, "test2 faction make failure"); bad {
		log.Println(my.LogError())
	} else {
		log.Println("Made faction!")
	}
	fac, err := db.GetFaction("gid", gid, "owner", "AutoTest")
	if my, bad := Check(err, "test2 getfaction failure"); bad {
		log.Println(my.MuleError())
		return
	}
	log.Println("Got faction:", fac)
	err = db.MakeMapView(gid, fac.Fid(), hexagon.Coord{10, 10})
	if my, bad := Check(err, "test2 mapview make failure", "gid", gid, "fid", fac.Fid()); bad {
		log.Println(my.LogError())
	} else {
		log.Println("Made mapview!")
	}
	mapV, err := db.GetMapView("gid", gid, "fid", fac.Fid())
	if my, bad := Check(err, "test2 getmapview failure"); bad {
		log.Println(my.MuleError())
		return
	}
	log.Println("Got mapView:", mapV)
	err = db.MakePlanet(gid, 100, 0, 10, 20, 30, "Planet Test", hexagon.Coord{10, 10})
	if my, bad := Check(err, "test2 make planet failure"); bad {
		log.Println(my.LogError())
	} else {
		log.Println("Made planet!")
	}
	pl, err := db.GetPlanet("gid", gid, "pid", 100)
	if my, bad := Check(err, "test2 get planet failure"); bad {
		log.Println(my.MuleError())
		return
	}
	log.Println("Got planet", pl)
	err = db.MakePlanet(gid, 101, 0, 10, 20, 30, "Planet Test2", hexagon.Coord{20, 20})
	if my, bad := Check(err, "test2 make planet2 failure"); bad {
		log.Println(my.LogError())
	} else {
		log.Println("Made planet2!")
	}
	pl2, err := db.GetPlanet("gid", gid, "pid", 101)
	if my, bad := Check(err, "test2 get planet2 failure"); bad {
		log.Println(my.MuleError())
		return
	}
	log.Println("Got planet2", pl2)
	plList, err := db.GetPlanetsByPlid(gid, 100, 101)
	if my, bad := Check(err, "test2 multi-planet get failed"); bad {
		log.Println(my.MuleError())
		return
	}
	log.Println("Got multi planets!", plList[0], plList[1])

	err = db.MakePlanetView(gid, pl.Pid(), fac.Fid(), 0, 0, 10, 10, 10, "Planet Test", hexagon.Coord{10, 10})
	if my, bad := Check(err, "test2 make planetview failure"); bad {
		log.Println(my.LogError())
	} else {
		log.Println("Made planetview!")
	}
	plV, err := db.GetPlanetView("gid", gid, "fid", fac.Fid(), "pid", 100)
	if my, bad := Check(err, "test2 get planetview failure"); bad {
		log.Println(my.MuleError())
		return
	}
	log.Println("Got planetview", plV)
	err = db.MakeOrder(gid, fac.Fid(), pl.Pid(), pl.Pid(), 5)
	if my, bad := Check(err, "test2 make order failure"); bad {
		log.Println(my.LogError())
	} else {
		log.Println("Made order!")
	}
	ord, err := db.GetOrder("gid", gid)
	if my, bad := Check(err, "test2 get order failure"); bad {
		log.Println(my.MuleError())
		return
	}
	log.Println("Got order", ord)
	err = db.MakeShip(gid, fac.Fid(), 5, 1, []hexagon.Coord{hexagon.Coord{10, 11}})
	if my, bad := Check(err, "test2 make ship failure"); bad {
		log.Println(my.LogError())
	} else {
		log.Println("Made ship!")
	}
	sh, err := db.GetShip("gid", gid)
	if my, bad := Check(err, "test2 get ship failure"); bad {
		log.Println(my.MuleError())
		return
	}
	log.Println("Got ship:", sh)

	err = db.MakeShipView(gid, fac.Fid(), sh.Sid(), 1, sh.Size(), sh.Fid(), hexagon.NullCoord{Valid: true, Coord: hexagon.Coord{15, 15}}, hexagon.NullCoord{hexagon.Coord{0, 100}, false}, nil)
	if my, bad := Check(err, "test2 make shipview failure"); bad {
		log.Println(my.LogError())
	} else {
		log.Println("Made shipview!")
	}
	shV, err := db.GetShipView("gid", gid)
	if my, bad := Check(err, "test2 get shipview failure"); bad {
		log.Println(my.MuleError())
		return
	}
	log.Println("Got shipview:", shV)

	err = db.MakeReport(gid, fac.Fid(), 1, []string{"HEL,LO"})
	if my, bad := Check(err, "test2 make report failure"); bad {
		if my.BaseIs(ErrNotUnique) {
			log.Println("Report not unique!")
		} else {
			log.Println(my.LogError())
		}
	} else {
		log.Println("Made report!")
	}
	rpt, err := db.GetReport("gid", gid)
	if my, bad := Check(err, "test2 get report failure"); bad {
		log.Println(my.MuleError())
		return
	}
	log.Println("Got report:", rpt, len(rpt.Contents()))

	// ---------------------- //
	f := func(d2 DB) error {
		game.IncTurn()
		log.Println("Incing Turn to:", game.Turn())
		err = d2.UpdateGames(game)
		if my, bad := Check(err, "test2 update failure"); bad {
			log.Println(my.MuleError())
			return my
		}
		log.Println("Updated game", gid)
		fac.SetDone(true)
		err = d2.UpdateFactions(fac)
		if my, bad := Check(err, "test2 update faction failure"); bad {
			log.Println(my.MuleError())
			return my
		}
		log.Println("Updated faction", fac)
		mapV.SetZoom(16)
		err = d2.UpdateMapViews(mapV)
		if my, bad := Check(err, "test2 update mapV failure"); bad {
			log.Println(my.MuleError())
			return my
		}
		log.Println("Updated mapV", mapV)
		ord.SetSize(-1)
		err = d2.UpdateOrders(ord)
		if my, bad := Check(err, "test2 update ord failure"); bad {
			log.Println(my.MuleError())
			return my
		}
		log.Println("Updated Ord", ord)
		return nil
	}
	err = db.Transact(f)
	if my, bad := Check(err, "test2 failed transaction"); bad {
		log.Println(my.MuleError())
		return
	}
	log.Println("Transaction success!")
	// -------------------- //
	ord, err = db.GetOrder("gid", gid)
	if my, bad := Check(err, "test2 second get order failure"); bad {
		if my.BaseIs(ErrNoneFound) {
			log.Println("2nd time order not found!")
		} else {
			log.Println(my.MuleError())
			return
		}
	} else {
		log.Println("Got order (2nd time)", ord)
	}
	err = db.dropTheseShips([]overpower.Ship{sh})
	if my, bad := Check(err, "dropship failure", "ship", sh); bad {
		log.Println(my.MuleError())
		return
	}
	log.Println("dropped ships!")
	shList, err := db.GetShips("gid", gid)
	if my, bad := Check(err, "test2 get shipList failure"); bad {
		log.Println(my.MuleError())
		return
	}
	log.Println("Got shipList:", shList)
	// -------------------- //
	err = db.DropGames("gid", gid)
	if my, bad := Check(err, "test2 failed dropgame", "gid", gid); bad {
		log.Println(my.MuleError())
		return
	}
	log.Println("Dropped game", gid)
}

func XTest3(t *testing.T) {
	db, err := LoadDB()
	if my, bad := Check(err, "test2 loaddb fail"); bad {
		log.Println(my.MuleError())
		return
	}
	log.Println("DB loaded")
	// -------------------- //
	game, err := db.GetGame("owner", "AutoTest")
	if my, bad := Check(err, "test2 getgame fail", "owner", "AutoTest"); bad {
		if !my.BaseIs(ErrNoneFound) {
			log.Println(my.MuleError())
			return
		}
		log.Println("game not found")
		err = db.MakeGame("AutoTest", "Automade", "", 40)
		if my, bad := Check(err, "test2 automake game fail"); bad {
			log.Println(my.MuleError())
			return
		}
		log.Println("Made auto-test game")
		game, err = db.GetGame("owner", "AutoTest")
		if my, bad := Check(err, "test2 getgame fail two", "owner", "AutoTest"); bad {
			log.Println(my.MuleError())
			return
		}
	}
	log.Println("Looking at game:", game)
	gid := game.Gid()
	fac, err := db.GetFaction("gid", gid, "owner", "AutoTest")
	if my, bad := Check(err, "test2 getfaction failure"); bad {
		if !my.BaseIs(ErrNoneFound) {
			log.Println(my.MuleError())
			return
		}
		log.Println("fac not found")
		err = db.MakeFaction(gid, "AutoTest", "AutoFaction")
		if my, bad := Check(err, "test2 faction make failure"); bad {
			log.Println(my.MuleError())
			return
		}
		log.Println("Made faction!")
		fac, err = db.GetFaction("gid", gid, "owner", "AutoTest")
		if my, bad := Check(err, "test2 getfaction failure"); bad {
			log.Println(my.MuleError())
			return
		}
	}
	log.Println("Got faction:", fac)

	rpts, err := db.GetReports("gid", gid)
	if my, bad := Check(err, "test2 get reports failure"); bad {
		log.Println(my.MuleError())
		return
	}
	if len(rpts) > 0 {
		err = db.DropReports("gid", gid)
		if my, bad := Check(err, "test2 drop reports failure"); bad {
			log.Println(my.MuleError())
			return
		}
		log.Println("Dropped", len(rpts), "reports!")
	}
	log.Println("Making report")
	err = db.MakeReport(gid, fac.Fid(), 1, []string{"HEL<,\"'LO", "D"})
	if my, bad := Check(err, "test2 make report failure"); bad {
		log.Println(my.LogError())
		return
	}
	log.Println("Made report!")
	rpts, err = db.GetReports("gid", gid)
	if my, bad := Check(err, "test2 get reports failure"); bad {
		log.Println(my.MuleError())
		return
	}
	if len(rpts) == 0 {
		log.Println("No reports found!")
		return
	}
	log.Println("Got reports:", rpts, len(rpts))
	for i, rpt := range rpts {
		log.Println("REPORT", i, ":", rpt)
		for j, item := range rpt.Contents() {
			log.Println("   CONTENT ITEM", j, ":", item)
		}
	}
	// -------------------------------- //
	return
	err = db.DropGames("gid", gid)
	if my, bad := Check(err, "test2 failed dropgame", "gid", gid); bad {
		log.Println(my.MuleError())
		return
	}
	log.Println("Dropped game", gid)
}