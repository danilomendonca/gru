package policy

import (
	"math"

	"github.com/elleFlorio/gru/autonomic/analyzer"
	cfg "github.com/elleFlorio/gru/configuration"
	"github.com/elleFlorio/gru/enum"
	"github.com/elleFlorio/gru/service"
)

const c_SWAP_MAX_DIST = 0.6

type swapCreator struct{}

func (p *swapCreator) getPolicyName() string {
	return "swap"
}

func (p *swapCreator) listActions() []string {
	return []string{"stop", "remove", "start"}
}

func (p *swapCreator) createPolicies(srvList []string, analytics analyzer.GruAnalytics) []Policy {
	swapPolicies := []Policy{}

	swapPairs := p.createSwapPairs(srvList)
	for running, inactives := range swapPairs {
		for _, inactive := range inactives {
			policyName := p.getPolicyName()
			policyWeight := p.computeWeight(running, inactive, analytics)
			policyTargets := []string{running, inactive}
			policyActions := map[string][]enum.Action{
				running:  []enum.Action{enum.STOP, enum.REMOVE},
				inactive: []enum.Action{enum.START},
			}

			swapPolicy := Policy{
				Name:    policyName,
				Weight:  policyWeight,
				Targets: policyTargets,
				Actions: policyActions,
			}

			swapPolicies = append(swapPolicies, swapPolicy)
		}
	}

	return swapPolicies
}

func (p *swapCreator) createSwapPairs(srvList []string) map[string][]string {
	pairs := map[string][]string{}

	running := []string{}
	inactive := []string{}

	for _, name := range srvList {
		srv, _ := service.GetServiceByName(name)
		if len(srv.Instances.Running) > 0 {
			running = append(running, name)
		} else {
			inactive = append(inactive, name)
		}
	}

	for _, name := range running {
		pairs[name] = inactive
	}

	return pairs
}

func (p *swapCreator) computeWeight(running string, candidate string, analytics analyzer.GruAnalytics) float64 {
	srv_run, _ := service.GetServiceByName(running)
	srv_cand, _ := service.GetServiceByName(candidate)
	nRun := len(srv_run.Instances.Running)
	baseServices := cfg.GetNodeConstraints().BaseServices

	if p.contains(baseServices, running) && nRun < 2 {
		return 0.0
	}

	runAnalytics := analytics.Service[running]
	candAnalytics := analytics.Service[candidate]

	// If the service has the resources to start without stopping the other
	// there is no reason to swap them
	if candAnalytics.Resources.Available > 0 {
		return 0.0
	}

	// TODO now this works only with homogeneous containers
	// and taking into account only the CPUs. This is not a
	// a good thing, so in the feuture the swap policy should
	// be able to compare the resources needed by each containers
	// and evaulte if it is possible to swap a container with
	// more than one that is active, in order to obtain
	// the requested amount of resources.
	if srv_run.Docker.CPUnumber != srv_cand.Docker.CPUnumber {
		return 0.0
	}

	cpuDist := candAnalytics.Resources.Cpu - runAnalytics.Resources.Cpu
	loadDist := candAnalytics.Load - runAnalytics.Load

	cpuValue := math.Min(1.0, cpuDist/c_SWAP_MAX_DIST)
	loadValue := math.Min(1.0, loadDist/c_SWAP_MAX_DIST)

	weight := math.Max(0.0, (cpuValue+loadValue)/2)

	return weight
}

// FIX duplicated from scalein...
func (p *swapCreator) contains(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}

	return false
}
