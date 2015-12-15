package autonomic

import (
	"time"

	log "github.com/elleFlorio/gru/Godeps/_workspace/src/github.com/Sirupsen/logrus"

	"github.com/elleFlorio/gru/autonomic/analyzer"
	"github.com/elleFlorio/gru/autonomic/executor"
	"github.com/elleFlorio/gru/autonomic/monitor"
	"github.com/elleFlorio/gru/autonomic/planner"
	"github.com/elleFlorio/gru/friends"
	"github.com/elleFlorio/gru/metric"
)

var timeInterval, maxFriends int

func Initialize(loopTimeInterval int, nFriends int) {
	timeInterval = loopTimeInterval
	maxFriends = nFriends
}

func RunLoop() {
	c_err := make(chan error)
	c_stop := make(chan struct{})

	monitor.Start(c_err, c_stop)
	planner.SetPlannerStrategy("probabilistic")

	// Set the ticker for the periodic execution
	ticker := time.NewTicker(time.Duration(timeInterval) * time.Second)

	log.Infoln("Running autonomic loop")
	for {
		select {
		case <-ticker.C:
			err := friends.UpdateFriendsData(maxFriends)
			if err != nil {
				log.WithField("err", err).Debugln("Cannot update friends data")
			}

			stats := monitor.Run()
			analytics := analyzer.Run(stats)
			plan := planner.Run(analytics)
			executor.Run(plan)

			collectMetrics()

			log.Infoln("-------------------------")

		case <-c_err:
			log.Errorln("Error running autonomic loop")
		case <-c_stop:
			ticker.Stop()
		}
	}
}

func collectMetrics() {
	log.Debugln("Collecting metrics")
	metric.UpdateMetrics()
	err := metric.StoreMetrics(metric.Metrics())
	if err != nil {
		log.WithField("errr", err).Errorln("Error collecting agent metrics")
	}
}
