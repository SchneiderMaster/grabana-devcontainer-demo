package timeseriesx

import (
	"github.com/K-Phoen/grabana/timeseries"
	"github.com/K-Phoen/sdk"
	"github.com/fastleansmart/grabana-devcontainer/pkg/extensions/grafanax"
)

func WithGrafanaTarget(queryType grafanax.QueryType, options ...grafanax.Option) timeseries.Option {
	grafana := grafanax.New(queryType, options...)
	fields := []string{}
	for _, field := range grafana.MeasurementFields {
		fields = append(fields, string(field))
	}
	return func(graph *timeseries.TimeSeries) error {
		graph.Builder.AddTarget(&sdk.Target{
			RefID:     grafana.Ref,
			Hide:      grafana.Hidden,
			QueryType: string(grafana.QueryType),
			Channel:   grafana.Channel,
			Filter: struct {
				Fields []string "json:\"fields,omitempty\""
			}{fields},
			Path: grafana.Path,
		})

		return nil
	}
}
