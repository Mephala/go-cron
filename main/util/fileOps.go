package fileOps

import (
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

func RenameIfNecessary(absolutePath string) {
	in, err := os.Open(absolutePath)
	if err != nil {
		log.Println("No existing backup file detected, no need to rename.")
		return
	}
	defer in.Close()

	var newPath = absolutePath + "." + strconv.FormatInt(makeTimestamp(), 16) + ".old"
	log.Println("Existing backup file detected, renaming it to:", newPath)

	e := os.Rename(absolutePath, newPath)
	if e != nil {
		log.Fatal(e)
	}
}

func Copy(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}

func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
