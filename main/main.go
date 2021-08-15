package main

import (
	"flag"
	"fmt"
	"go-cron/main/thunderbird"
	"os"
)

const (
	InfoColor         = "\033[1;34m%s\033[0m"
	NoticeColor       = "\033[1;36m%s\033[0m"
	WarningColor      = "\033[1;33m%s\033[0m"
	ErrorColor        = "\033[1;31m%s\033[0m"
	DebugColor        = "\033[0;36m%s\033[0m"
	ThunderbirdEnvVar = "THUNDERBIRD_FILTER_FOLDER_LOCATION"
	BackupLocEnvVar   = "THUNDERBIRD_FILTER_BACKUP_FOLDER_LOCATION"
)

func main() {
	filterLocPtr := flag.String("fl", os.Getenv(ThunderbirdEnvVar),
		"Absolute path to folder containing thunderbird filters to store")

	backupLocationPtr := flag.String("bl", os.Getenv(BackupLocEnvVar),
		"Absolute path of folder to backup thunderbird filters")

	flag.Parse()

	if len(*filterLocPtr) == 0 {
		fmt.Printf(ErrorColor, "error: Please provide thunderbird filter location via -fl flag or set env variable:"+ThunderbirdEnvVar+"\n")
		os.Exit(1)
	}
	if len(*backupLocationPtr) == 0 {
		fmt.Printf(ErrorColor, "error: Please provide location of backup folder -bl flag or set env variable:"+BackupLocEnvVar+"\n")
		os.Exit(2)
	}

	thunderbirdInstance := thunderbird.New(*filterLocPtr, *backupLocationPtr)

	thunderbirdInstance.Backup()
}
