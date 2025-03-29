package functions

import (
	"errors"
	"io/fs"
	"log"
	"os"
	"strings"
)

func BackupInfo() string {
	if _, err := os.Stat("./backup"); errors.Is(err, fs.ErrNotExist) {
		err = os.Mkdir("backup", 0755)
		if err != nil && !errors.Is(err, fs.ErrExist) {
			log.Println(err)
			return "cannot create folder './backup'"
		}

		return "latest backup: never"
	}

	d, err := os.Open("./backup")
	if err != nil {
		log.Println(err)
		return "cannot open folder './backup'"
	}

	files, err := d.ReadDir(0)
	if err != nil {
		log.Println(err)
		return "cannot list files from folder './backup'"
	}

	var backupList []string
	backupList = append(backupList, "backup list:")

	for _, f := range files {
		if !f.IsDir() {
			backupList = append(backupList, f.Name())
		}
	}

	if len(backupList) == 1 {
		return "latest backup: never"
	}

	return strings.Join(backupList, "\n")
}
