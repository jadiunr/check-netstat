package main

import (
	"fmt"
	"sort"

	"github.com/iancoleman/strcase"
	"github.com/sensu-community/sensu-plugin-sdk/sensu"
	"github.com/sensu/sensu-go/types"
	"github.com/shirou/gopsutil/v3/net"
)

// Config represents the check plugin config.
type Config struct {
	sensu.PluginConfig
}

type MetricGroup struct {
	Comment string
	Type    string
	Name    string
	Value   int64
}

func (g *MetricGroup) Output() {
	fmt.Printf("# HELP %s %s\n", g.Name, g.Comment)
	fmt.Printf("# TYPE %s %s\n", g.Name, g.Type)
	fmt.Printf("%s %v\n", g.Name, g.Value)
}

var (
	plugin = Config{
		PluginConfig: sensu.PluginConfig{
			Name:     "check-netstat",
			Short:    "Simple cross-platform network statistics checks",
			Keyspace: "sensu.io/plugins/check-netstat/config",
		},
	}
)

func main() {
	check := sensu.NewGoCheck(&plugin.PluginConfig, nil, checkArgs, executeCheck, false)
	check.Execute()
}

func checkArgs(event *types.Event) (int, error) {
	return sensu.CheckStateOK, nil
}

func executeCheck(event *types.Event) (int, error) {
	protoCounters, err := net.ProtoCounters(nil)
	if err != nil {
		return sensu.CheckStateCritical, fmt.Errorf("Failed to get network statistics, error: %v", err)
	}

	var metricGroups = map[string]*MetricGroup{}

	for _, protoCounter := range protoCounters {
		for k, v := range protoCounter.Stats {
			measurement := fmt.Sprintf("%s_%s", protoCounter.Protocol, strcase.ToSnake(k))
			metricGroups[measurement] = &MetricGroup{
				Name:    measurement,
				Type:    "untyped",
				Comment: fmt.Sprintf("Statistic %s %s", protoCounter.Protocol, k),
				Value:   v,
			}
		}
	}

	sortedMeasurements := make([]string, 0, len(metricGroups))
	for measurement := range metricGroups {
		sortedMeasurements = append(sortedMeasurements, measurement)
	}
	sort.Strings(sortedMeasurements)
	for _, sortedMeasurement := range sortedMeasurements {
		metricGroups[sortedMeasurement].Output()
	}

	return sensu.CheckStateOK, nil
}
