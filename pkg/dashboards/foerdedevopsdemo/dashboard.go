package foerdedevopsdemo

import (
	"github.com/K-Phoen/grabana/dashboard"
	"github.com/K-Phoen/grabana/timeseries"
	"github.com/K-Phoen/grabana/variable/custom"
	"github.com/fastleansmart/grabana-devcontainer/pkg/extensions/grafanax"
	"github.com/fastleansmart/grabana-devcontainer/pkg/extensions/textx"
	"github.com/fastleansmart/grabana-devcontainer/pkg/extensions/timeseriesx"
)

func Dashboard() (dashboard.Builder, error) {
	return dashboard.New(
		"Förde DevOps Demo",
		dashboard.UID("foerde-devops-demo"),
		dashboard.VariableAsCustom("variable",
			custom.Label("Custom variable"),
			custom.Values(custom.ValuesMap{"some value": "some value", "another value": "another value", "yet again, a value": "yet again\\, a value"}),
		),
		dashboard.WithText(
			"Cool demo text",
			textx.Markdown(`# Hello World
This is some markdown text that has been generated using Grabana!

Your currently selected option is `+"`"+`${variable}`+"`"+`.
			`),
			textx.GridPos(12, 12, 0, 0),
		),
		dashboard.WithTimeSeries("Some timeseries",
			timeseries.DataSource("grafana"),
			timeseriesx.WithGrafanaTarget(grafanax.RandomWalk),
			timeseriesx.GridPos(12, 12, 12, 0),
		),
		dashboard.WithTimeSeries("Another interesting timeseries",
			timeseries.DataSource("grafana"),
			timeseriesx.WithGrafanaTarget(grafanax.RandomWalk),
			timeseriesx.GridPos(12, 24, 0, 12),
		),
	)
}
