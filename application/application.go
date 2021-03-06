package application

// Encapsulates application details
type Tags struct {
	Application string
	Service     string
	Cluster     string
	Shard       string
	CustomTags  map[string]string
}

// Creates a new Tags with application and service name
func New(application, service string) Tags {
	return Tags{
		Application: application,
		Service:     service,
		CustomTags:  make(map[string]string, 0),
	}
}

// Map with all values
func (a Tags) Map() map[string]string {
	allTags := make(map[string]string)

	allTags["application"] = a.Application
	allTags["service"] = a.Service

	if len(a.Cluster) > 0 {
		allTags["cluster"] = a.Cluster
	}

	if len(a.Shard) > 0 {
		allTags["shard"] = a.Shard
	}

	for k, v := range a.CustomTags {
		allTags[k] = v
	}
	return allTags
}
