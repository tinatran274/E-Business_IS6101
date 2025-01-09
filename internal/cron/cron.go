package crons

import (
	"github.com/robfig/cron/v3"
)

type Cron struct {
	cron *cron.Cron
}

func NewCron() *Cron {
	return &Cron{
		cron: cron.New(),
	}
}

func (c *Cron) Start() {
	if _, err := c.cron.AddFunc("0 0 0 1 * *", func() {
	}); err != nil {
		panic(err)
	}

	c.cron.Start()
}

func (c *Cron) Stop() <-chan struct{} {
	return c.cron.Stop().Done()
}
