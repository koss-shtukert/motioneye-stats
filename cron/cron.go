package cron

import (
	"github.com/koss-shtukert/motioneye-stats/bot"
	"github.com/koss-shtukert/motioneye-stats/config"
	"github.com/koss-shtukert/motioneye-stats/cron/job"
	"github.com/robfig/cron/v3"
	"github.com/rs/zerolog"
)

type Cron struct {
	cron   *cron.Cron
	tgBot  *bot.Bot
	logger *zerolog.Logger
	config *config.Config
}

func NewCron(l *zerolog.Logger, cfg *config.Config, b *bot.Bot) *Cron {
	logger := l.With().Str("type", "cron").Logger()

	c := &Cron{
		cron:   cron.New(),
		tgBot:  b,
		logger: &logger,
		config: cfg,
	}

	return c
}

func (c *Cron) AddDiskUsageJob() {
	if _, err := c.cron.AddFunc(c.config.CronDiskUsageJobInterval, job.DiskUsageJob(c.logger, c.config, c.tgBot)); err != nil {
		c.logger.Err(err).Msg("Failed to schedule DiskUsage job")
	}
}

func (c *Cron) Start() {
	c.cron.Start()
}
