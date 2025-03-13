package grafanax

type QueryType string

const (
	RandomWalk       QueryType = "randomWalk"
	ListPublicFiles  QueryType = "list"
	LiveMeasurements QueryType = "measurements"
)

type MeasurementField string

const (
	Time  MeasurementField = "time"
	value MeasurementField = "value"
	Min   MeasurementField = "min"
	Max   MeasurementField = "max"
)

type Option func(target *Grafana)

type Grafana struct {
	Ref               string
	Hidden            bool
	QueryType         QueryType
	Path              string
	Channel           string
	MeasurementFields []MeasurementField
}

func New(queryType QueryType, options ...Option) *Grafana {
	grafana := &Grafana{
		QueryType: queryType,
	}

	for _, opt := range options {
		opt(grafana)
	}

	return grafana
}

func Ref(ref string) Option {
	return func(grafana *Grafana) {
		grafana.Ref = ref
	}
}

func Hide() Option {
	return func(grafana *Grafana) {
		grafana.Hidden = true
	}
}

func FilePath(path string) Option {
	return func(grafana *Grafana) {
		grafana.Path = path
	}
}

func MeasurementChannel(channel string) Option {
	return func(grafana *Grafana) {
		grafana.Channel = channel
	}
}

func MeasurementFields(fields []MeasurementField) Option {
	return func(grafana *Grafana) {
		grafana.MeasurementFields = fields
	}
}
