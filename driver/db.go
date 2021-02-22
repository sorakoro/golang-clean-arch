package driver

import (
	"fmt"
	"net/url"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

// NewDBConn DBに接続する
func NewDBConn() (*sqlx.DB, func(), error) {
	dbHost := viper.GetString("database.host")
	dbPort := viper.GetString("database.port")
	dbUser := viper.GetString("database.user")
	dbPassword := viper.GetString("database.password")
	dbName := viper.GetString("database.name")
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	uv := url.Values{}
	uv.Add("parseTime", "1")
	uv.Add("loc", "Asia/Tokyo")
	dsn := fmt.Sprintf("%s?%s", connection, uv.Encode())
	db, err := sqlx.Connect("mysql", dsn)

	if err != nil {
		fmt.Fprint(os.Stderr, err)
		return nil, nil, err
	}

	cleanup := func() {
		err := db.Close()
		if err != nil {
			fmt.Fprint(os.Stderr, err)
		}
	}
	return db, cleanup, nil
}
