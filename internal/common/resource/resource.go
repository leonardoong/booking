package resource

import (
	"fmt"
	"log"
	"sync"

	"github.com/jmoiron/sqlx"
	"github.com/leonardoong/booking/internal/config"

	_ "github.com/lib/pq"
)

var (
	resources *Resources
	once      sync.Once
)

type Resources struct {
	// database
	DB sqlx.DB
}

func GetResources(cfg *config.Config) (*Resources, error) {
	var err error
	once.Do(func() {
		resources, err = initResources(cfg)
	})
	return resources, err
}

func initResources(cfg *config.Config) (*Resources, error) {
	// init DB
	psqlInfo := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.DBName,
	)

	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("error connect to db: %v \n", err)
		return nil, err
	}

	return &Resources{
		DB: *db,
	}, err

}
