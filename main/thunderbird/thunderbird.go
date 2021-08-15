package thunderbird

import (
	"fmt"
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
}
