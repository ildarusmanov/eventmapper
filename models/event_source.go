package models

type EventSource struct {
	SourceType string `validate:"nonzero,min=1,max=255" json:"source_type,omitempty"`
	SourceId string `validate:"nonzero,min=1,max=255" json:"source_id,omitempty"`
	Origin string `validate:"max=500" json:"origin,omitempty"`
	Params map[string]string `validate:"max=50" json:"params,omitempty"`
}

func CreateNewEventSource(sourceType, sourceId, origin string, params map[string]string) *EventSource {
	return &EventSource{sourceType, sourceId, origin, params}
}