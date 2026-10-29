package fedramp

import (
	"encoding/json"
	"reflect"
	"sort"
	"strings"
	"testing"
)

// docStructForSchema maps each document schema file to the Go struct that
// represents its root object. The common-definitions schema is excluded: it
// declares only $defs (no top-level "properties") and its types are verified
// through the documents that reference them.
var docStructForSchema = map[string]any{
	SchemaFileAcceptedVulnerabilityInfo:      AcceptedVulnerabilityInfoReport{},
	SchemaFileAdvisorInformation:             AdvisorInformation{},
	SchemaFileAssessorInformation:            AssessorInformation{},
	SchemaFileCertificationPackageOverview:   CertificationPackageOverview{},
	SchemaFileHistoricalVERActivity:          HistoricalVERActivity{},
	SchemaFileIncidentReport:                 IncidentReport{},
	SchemaFileOngoingCertificationReport:     OngoingCertificationReport{},
	SchemaFileSecurityDecisionRecord:         SecurityDecisionRecord{},
	SchemaFileSignificantChangeNotifications: SignificantChangeNotification{},
	SchemaFileVulnerabilityDetailReport:      VulnerabilityDetailReport{},
}

// schemaTopLevelProps returns the top-level "properties" keys of a schema.
func schemaTopLevelProps(t *testing.T, name string) []string {
	t.Helper()
	b, err := SchemaBytes(name)
	if err != nil {
		t.Fatalf("SchemaBytes(%q): %v", name, err)
	}
	var doc struct {
		Properties map[string]json.RawMessage `json:"properties"`
	}
	if err := json.Unmarshal(b, &doc); err != nil {
		t.Fatalf("unmarshal %q: %v", name, err)
	}
	keys := make([]string, 0, len(doc.Properties))
	for k := range doc.Properties {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// structJSONNames returns the JSON property names declared by a struct's tags.
func structJSONNames(t *testing.T, v any) []string {
	t.Helper()
	rt := reflect.TypeOf(v)
	var names []string
	for f := range rt.Fields() {
		tag := f.Tag.Get("json")
		if tag == "" || tag == "-" {
			continue
		}
		name, _, _ := strings.Cut(tag, ",")
		if name != "" {
			names = append(names, name)
		}
	}
	sort.Strings(names)
	return names
}

// TestSchemaStructParity ensures each document struct's top-level JSON tags
// exactly match its schema's top-level property names, so tags cannot silently
// drift from the authoritative schema.
func TestSchemaStructParity(t *testing.T) {
	for name, v := range docStructForSchema {
		t.Run(name, func(t *testing.T) {
			want := schemaTopLevelProps(t, name)
			got := structJSONNames(t, v)
			if !reflect.DeepEqual(want, got) {
				t.Errorf("top-level property mismatch for %q:\n schema: %v\n struct: %v", name, want, got)
			}
		})
	}
}
