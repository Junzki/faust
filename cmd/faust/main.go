package main

import (
	"faust/internal/faust"
	"faust/internal/pkg/asset"
	"faust/internal/pkg/config"
	"faust/internal/timed"
	"flag"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"os"
)


var (
	configFile	string
	g errgroup.Group
)


func loadFlags() {
	flag.StringVar(&configFile, "c", "", "Path to config file.")
	flag.Parse()
}

func main() {
	loadFlags()
	if "" == configFile {
		log.Fatal("Config file now specified.")
	}

	configFile = asset.ExpandUserDir(configFile)
	f, err := os.Open(configFile)
	if nil != err {
		log.Fatal(err)
	}

	log.Infof("Load config from %s", configFile)
	err = config.GetConfig().UpdateFromFile(f)
	if nil != err {
		log.Fatal(err)
	}

	svc := faust.GetSvc()
	g.Go(svc.Serve)

	timeSvc, _ := timed.NewService(0)
	g.Go(timeSvc.Serve)

	if err := g.Wait(); nil != err {
		log.Fatal(err)
	}
}
