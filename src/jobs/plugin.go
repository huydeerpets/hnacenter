package jobs

import (
	"hnnaserver/src/cron"
)

const DEFAULT_JOB_POOL_SIZE = 10

var (
	// Singleton instance of the underlying job scheduler.
	MainCron *cron.Cron

	// This limits the number of jobs allowed to run concurrently.
	workPermits chan struct{}

	// Is a single job allowed to run concurrently with itself?
	selfConcurrent bool
)

func init() {
	MainCron = cron.New()
	workPermits = make(chan struct{}, DEFAULT_JOB_POOL_SIZE)
	selfConcurrent = false
	MainCron.Start()
}
