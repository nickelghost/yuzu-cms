package boot

import (
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"path/filepath"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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

func ForwardWebpack(e *echo.Echo, targetURL string) {
	g := e.Group("/admin")
	webpackURL, err := url.Parse(targetURL)
	if err != nil {
		log.Fatal(err)
	}
	g.Use(middleware.Proxy(middleware.NewRandomBalancer(
		[]*middleware.ProxyTarget{{URL: webpackURL}},
	)))
}
