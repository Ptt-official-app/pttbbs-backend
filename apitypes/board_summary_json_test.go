package apitypes

import (
	"encoding/json"
	"testing"
)

func TestBoardSummaryJSONOptionalFields(t *testing.T) {
	for _, enabled := range []bool{false, true} {
		summary := BoardSummary{IsPopular: enabled}
		if enabled {
			summary.StatAttr = 1
		}
		data, err := json.Marshal(summary)
		if err != nil {
			t.Fatal(err)
		}
		var fields map[string]json.RawMessage
		if err := json.Unmarshal(data, &fields); err != nil {
			t.Fatal(err)
		}
		for _, key := range []string{"is_popular", "stat_attr"} {
			if _, present := fields[key]; present != enabled {
				t.Errorf("%s presence = %v, want %v: %s", key, present, enabled, data)
			}
		}
		if string(fields["is_over_18"]) != "false" {
			t.Errorf("required false flag missing: %s", data)
		}
		if enabled && (string(fields["is_popular"]) != "true" || string(fields["stat_attr"]) != "1") {
			t.Errorf("nonzero values changed: %s", data)
		}
	}
}
