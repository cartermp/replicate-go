package replicate

import (
	"context"
	"encoding/json"
	"fmt"
)

type Hardware struct {
	SKU  string `json:"sku"`
	Name string `json:"name"`

	rawJSON json.RawMessage `json:"-"`
}

func (h Hardware) MarshalJSON() ([]byte, error) {
	if h.rawJSON != nil {
		return h.rawJSON, nil
	} else {
		type Alias Hardware
		return json.Marshal(&struct{ *Alias }{Alias: (*Alias)(&h)})
	}
}

func (h *Hardware) UnmarshalJSON(data []byte) error {
	h.rawJSON = data
	type Alias Hardware
	alias := &struct{ *Alias }{Alias: (*Alias)(h)}
	return json.Unmarshal(data, alias)
}

// ListHardware returns a list of available hardware.
func (r *Client) ListHardware(ctx context.Context) (*[]Hardware, error) {
	response := &[]Hardware{}
	err := r.fetch(ctx, "GET", "/hardware", nil, response)
	if err != nil {
		return nil, fmt.Errorf("failed to list collections: %w", err)
	}
	return response, nil
}
