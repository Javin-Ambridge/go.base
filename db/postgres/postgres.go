package postgres

import (
	"fmt"
	"github.com/jinzhu/gorm"

	"github.com/Javin-Ambridge/go.base/go.base/entity"
	"github.com/Javin-Ambridge/go.base/go.base/utils/goutils"

	_ "github.com/lib/pq"
)

// For better testing
var (
	openDBConnection = gorm.Open
)

// Postgres is the interface to the PostgreSQL DB
type Postgres interface {
}

type postgres struct {
	secrets entity.Secrets
	db      *gorm.DB
}

// New provides a new Postgres Interface to Fx
func New(secrets entity.Secrets) (Postgres, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		secrets.PostgresInfo.Host,
		secrets.PostgresInfo.Port,
		secrets.PostgresInfo.User,
		secrets.PostgresInfo.Name,
	)

	// Create the connection to postgresql server
	db, err := openDBConnection("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	db.LogMode(false)

	return &postgres{
		secrets: secrets,
		db:      db,
	}, nil
}
