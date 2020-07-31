package infrastructures

import (
	"database/sql"
	"github.com/gocraft/dbr/v2"
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
)

type ISQLConnection interface {
	Connect() *dbr.Session
}

type SQLConnection struct {
	DB *sql.DB
}

var dbConnection *dbr.Connection

func createConnection(dialect string, descriptors string, maxConn, maxIdle int) *dbr.Connection{
	conn, err := dbr.Open(dialect, descriptors, nil)

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	conn.SetMaxIdleConns(maxIdle)
	conn.SetMaxOpenConns(maxConn)

	return conn
}

func (s *SQLConnection) Connect() *dbr.Session{
	if dbConnection == nil {
		dbConnection = createConnection(
			viper.GetString("database.dialect"),
			viper.GetString("database.descriptors"),
			viper.GetInt("database.max_conn"),
			viper.GetInt("database.max_idle"))
	}

	session := dbConnection.NewSession(nil)
	session.Timeout = 5 * time.Second
	return session
}