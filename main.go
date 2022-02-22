package main

import (
	"io/ioutil"
	"log"
	"os"
)

var (
	minimumSize   int64 = 204800
	windowsAssets       = "/Packages/Microsoft.Windows.ContentDeliveryManager_cw5n1h2txyewy/LocalState/Assets"
	destination         = "/Desktop/Wallpapers"
)

func main() {
	userCacheDir, err := os.UserCacheDir()
	if err != nil {
		log.Fatal(err)
	}

	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	assetsPath := userCacheDir + windowsAssets
	destinationPath := userHomeDir + destination + "/"

	createDirectory(destinationPath)
	copyWallpapers(assetsPath, destinationPath)
}

func createDirectory(path string) {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		log.Println(err)
	}
}

func copyWallpapers(src string, dest string) {
	files, err := ioutil.ReadDir(src)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if f.Size() > minimumSize {
			currentFile := src + "/" + f.Name()
			destFile := dest + f.Name() + ".jpg"
			copyFile(currentFile, destFile)
		}
	}
}

func copyFile(src string, dest string) {
	input, err := ioutil.ReadFile(src)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(dest, input, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
