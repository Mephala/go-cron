package thunderbird

import (
	"fmt"
	fileOps "go-cron/main/util"
	"io/ioutil"
	"log"
	"os"
	"path"
	//"os"
)

const (
	MsgFiltersFileName = "msgFilterRules.dat"
)

type Thunderbird struct {
	folderLoc string
	backupLoc string
}

func New(folderLoc string, backupLoc string) Thunderbird {
	instance := Thunderbird{folderLoc, backupLoc}
	return instance
}

func (thunderbird Thunderbird) Backup() {
	fmt.Printf("Starting to backup filters from %s to %s \n", thunderbird.folderLoc, thunderbird.backupLoc)

	files, err := ioutil.ReadDir(thunderbird.folderLoc)
	if err != nil {
		log.Fatal(err)
	}

	var filterFileInfo os.FileInfo
	for _, f := range files {
		if f.Name() == MsgFiltersFileName {
			filterFileInfo = f
		}
	}
	if filterFileInfo == nil {
		log.Fatal("No file named \"%s\" found in given thunderbird filter folder", MsgFiltersFileName)
	}

	var filterFileAbsolutePath = path.Join(thunderbird.folderLoc, MsgFiltersFileName)
	var backupFileAbsolutePath = path.Join(thunderbird.backupLoc, MsgFiltersFileName)

	fileOps.RenameIfNecessary(backupFileAbsolutePath)
	err = fileOps.Copy(filterFileAbsolutePath, backupFileAbsolutePath)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Thunderbird backup operation completed without any errors.")
}
