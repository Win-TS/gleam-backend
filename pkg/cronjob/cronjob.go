// cronjob.go

package cronjob

import (
	"context"
	"log"
	"time"

	groupdb "github.com/Win-TS/gleam-backend.git/pkg/database/postgres/groupdb/sqlc"
	"github.com/go-co-op/gocron"
)

type CronjobService interface {
	Start() error
}

type cronjob struct {
	store groupdb.Store
}

func NewCronjobService(store groupdb.Store) CronjobService {
	return &cronjob{store}
}

func (c *cronjob) Start() error {
	cron := gocron.NewScheduler(time.UTC)

	cron.Every(2).Day().At("00:00").Do(func() {
		c.checkIncompleteTasks()
	})

	// cron.Every(1).Minutes().Do(func() {
	// 	c.checkIncompleteTasks()
	// })

	cron.StartAsync()

	log.Println("Cronjob started successfully")

	return nil
}

func (c *cronjob) checkIncompleteTasks() {
	c.store.ResetTotalStreak(context.Background())
	c.store.ResetWeeklyStreak(context.Background())
}
