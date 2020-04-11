package boot

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// GetSQL reads SQL files in a given location and returns a map with .sql file
// names and their content
func GetSQL(root string) map[string]string {
	sqlFiles := make(map[string]string)
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".sql" {
			sql, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			sqlFiles[info.Name()] = string(sql)
			return nil
		}
		return nil
	})
	if err != nil {
		log.Fatalf("Reading SQL files failed:\n%s", err)
	}
	return sqlFiles
}
