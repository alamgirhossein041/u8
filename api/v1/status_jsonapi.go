package v1

import (
	"github.com/uvite/u8/core"
)

// StatusJSONAPI is JSON API envelop for metrics
type StatusJSONAPI struct {
	Data statusData `json:"data"`
}

// NewStatusJSONAPI creates the JSON API status envelop
func NewStatusJSONAPI(s Status) StatusJSONAPI {
	return StatusJSONAPI{
		Data: statusData{
			ID:         "default",
			Type:       "status",
			Attributes: s,
		},
	}
}

// Status extract the v1.Status from the JSON API envelop
func (s StatusJSONAPI) Status() Status {
	return s.Data.Attributes
}

type statusData struct {
	Type       string `json:"type"`
	ID         string `json:"id"`
	Attributes Status `json:"attributes"`
}

func newStatusJSONAPIFromEngine(engine *core.Engine) StatusJSONAPI {
	return NewStatusJSONAPI(NewStatus(engine))
}
