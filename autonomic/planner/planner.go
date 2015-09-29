package planner

import (
	log "github.com/elleFlorio/gru/Godeps/_workspace/src/github.com/Sirupsen/logrus"

	"github.com/elleFlorio/gru/autonomic/analyzer"
	"github.com/elleFlorio/gru/autonomic/planner/strategy"
	"github.com/elleFlorio/gru/policy"
	"github.com/elleFlorio/gru/service"
)

var currentStrategy strategy.GruStrategy

func SetPlannerStrategy(strategyName string) {
	strtg, err := strategy.New(strategyName)
	if err != nil {
		log.WithFields(log.Fields{
			"status": "init",
			"error":  err,
		}).Errorln("Running Planner")

		// If error use default one
		strtg, err = strategy.New("dummy")
	}

	currentStrategy = strtg

	log.WithFields(log.Fields{
		"status":   "init",
		"strategy": strtg.Name(),
	}).Infoln("Running Planner")
}

func Run(analytics analyzer.GruAnalytics) strategy.GruPlan {
	log.WithField("status", "start").Debugln("Running planner")
	defer log.WithField("status", "done").Debugln("Running planner")

	plans := buildPlans(analytics)

	log.WithFields(log.Fields{
		"status": "plans builded",
		"plans":  len(plans),
	}).Debugln("Running Planner")

	thePlan, err := currentStrategy.MakeDecision(plans, analytics)
	if err != nil {
		log.WithFields(log.Fields{
			"status": "planning",
			"error":  err,
		}).Errorln("Running Planner")
	}

	log.WithFields(log.Fields{
		"status": "plan chosen",
		"plan":   thePlan,
	}).Debugln("Running Planner")

	return *thePlan
}

func buildPlans(analytics analyzer.GruAnalytics) []strategy.GruPlan {
	policies := policy.GetPolicies("proactive")
	plans := []strategy.GruPlan{}

	log.WithFields(log.Fields{
		"status":   "building plans",
		"policies": len(policies),
	}).Debugln("Running Planner")

	for _, name := range service.List() {
		for _, plc := range policies {
			if plc.Level() == "service" {

				if len(analytics.Service[name].Instances.Active) > 0 {
					plan := strategy.GruPlan{
						Service:      name,
						Weight:       plc.Weight(name, analytics),
						TargetType:   plc.Target(),
						TargetStatus: plc.TargetStatus(),
						Actions:      plc.Actions(),
					}

					plans = append(plans, plan)
				}
			}
		}
	}

	return plans
}
