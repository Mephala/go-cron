package thunderbird

import (
	"fmt"
	"io/ioutil"
	"log"
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

	var filterFound = false
	for _, f := range files {
		if f.Name() == MsgFiltersFileName {
			filterFound = true
		}
	}
	if !filterFound {
		log.Fatal("No file named \"%s\" found in given thunderbird filter folder", MsgFiltersFileName)
	}
}
