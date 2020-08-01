package infrastructures

import (
	"github.com/gocraft/dbr/v2"
	_ "github.com/lib/pq"

	"log"
	"os"
)

type ISQLConnection interface {
	Connect() *dbr.Session
}

type SQLConnection struct {
	Connection *dbr.Connection
}

func CreateConnection(dialect string, descriptors string, maxConn, maxIdle int) *dbr.Connection{
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
	session := s.Connection.NewSession(nil)
	return session
}