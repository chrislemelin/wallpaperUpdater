package main

import (
	"log"
	"os"
	"path/filepath"
	"wallpaperUpdater/util"

	_ "github.com/chrislemelin/wallpaperUpdater/util"
	"github.com/reujab/wallpaper"
)

func main() {
	setWallpaper()
}

func setWallpaper() {
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
