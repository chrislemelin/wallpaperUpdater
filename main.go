package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/chrislemelin/wallpaperUpdater/util"
	"github.com/jasonlvhit/gocron"
	"github.com/reujab/wallpaper"
)

func main() {
	s := gocron.NewScheduler()
	s.Every(1).Day().At("23:59").Do(setWallpaper)
	<-s.Start()
}

func setWallpaper() {
	fmt.Println("setting wallpaper")
	success, url := util.GetLink()
	if !success {
		log.Fatal("can't retrieve url from reddit")
	}
	file := util.GetFilePath()
	util.Download(file, url)

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	wallpaper.SetFromFile(dir + "/" + file)

}
