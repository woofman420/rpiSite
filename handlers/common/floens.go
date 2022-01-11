package common

import (
	"log"
	"math/rand"
	"os"
	"strings"
	"sync"

	"github.com/gofiber/fiber/v2"
)

var (
	catFiles     []os.DirEntry
	getFilesOnce sync.Once
)

func getFiles() []os.DirEntry {
	getFilesOnce.Do(func() {
		var err error
		catFiles, err = os.ReadDir("./static/cats")
		if err != nil {
			panic(err)
		}
	})
	return catFiles
}

func getFileExtension(path string) string {
	n := strings.LastIndexByte(path, '.')
	if n < 0 {
		return ""
	}
	return path[n:]
}

func FloensGet(c *fiber.Ctx) error {
	files := getFiles()
	randNum := rand.Intn(len(files))
	catFile := files[randNum].Name()
	f, err := os.Open("./static/cats/" + catFile)
	if err != nil {
		c.SendString("Error!")
		log.Println(err)
		return nil
	}
	stat, err := f.Stat()
	if err != nil {
		c.SendString("Error!")
		log.Println(err)
		return nil
	}
	c.Type(getFileExtension(stat.Name()))
	c.Response().Header.Add("Refresh", "2;url=/floens")
	c.Response().Header.Add("Cache-Control", "no-cache")
	c.Response().SetBodyStream(f, int(stat.Size()))
	return nil
}
