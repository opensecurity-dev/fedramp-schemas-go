package fedramp

import (
	"encoding/json"
	"strings"
	"testing"
)

// wantVersions maps each embedded schema file to its expected "$schemaVersion".
// Update this table when re-vendoring a new FedRAMP schema set.
var wantVersions = map[string]string{
	SchemaFileCommonDefinitions:              "0.3.0",
	SchemaFileAcceptedVulnerabilityInfo:      "0.1.1",
	SchemaFileAdvisorInformation:             "1.0.1",
	SchemaFileAssessorInformation:            "1.0.1",
	SchemaFileCertificationPackageOverview:   "0.1.4",
	SchemaFileHistoricalVERActivity:          "0.1.1",
	SchemaFileIncidentReport:                 "0.2.0",
	SchemaFileOngoingCertificationReport:     "0.2.0",
	SchemaFileSecurityDecisionRecord:         "1.1.1",
	SchemaFileSignificantChangeNotifications: "0.1.2",
	SchemaFileVulnerabilityDetailReport:      "0.1.1",
}

func TestSchemaFilesEmbedded(t *testing.T) {
	if len(SchemaFiles) != len(wantVersions) {
		t.Fatalf("SchemaFiles has %d entries, wantVersions has %d", len(SchemaFiles), len(wantVersions))
	}
	for _, name := range SchemaFiles {
		b, err := SchemaBytes(name)
		if err != nil {
			t.Fatalf("SchemaBytes(%q): %v", name, err)
		}
		if len(b) == 0 {
			t.Fatalf("SchemaBytes(%q): empty", name)
		}
		if !json.Valid(b) {
			t.Fatalf("SchemaBytes(%q): not valid JSON", name)
		}
		if !strings.HasSuffix(name, SchemaSetDate+".json") {
			t.Errorf("schema %q does not carry the schema-set date %q", name, SchemaSetDate)
		}
	}
}

func TestSchemaBytesUnknown(t *testing.T) {
	if _, err := SchemaBytes("does-not-exist.json"); err == nil {
		t.Fatal("expected error for unknown schema, got nil")
	}
}

func TestSchemaMetaVersions(t *testing.T) {
	for name, want := range wantVersions {
		s, err := SchemaMeta(name)
		if err != nil {
			t.Fatalf("SchemaMeta(%q): %v", name, err)
		}
		if s.Version != want {
			t.Errorf("%q: version = %q, want %q", name, s.Version, want)
		}
		if s.Title == "" {
			t.Errorf("%q: empty title", name)
		}
		if s.Name != name {
			t.Errorf("%q: Name = %q", name, s.Name)
		}
	}
}

func TestSchemas(t *testing.T) {
	schemas, err := Schemas()
	if err != nil {
		t.Fatalf("Schemas(): %v", err)
	}
	if len(schemas) != len(SchemaFiles) {
		t.Fatalf("Schemas() returned %d, want %d", len(schemas), len(SchemaFiles))
	}
	for i, s := range schemas {
		if s.Name != SchemaFiles[i] {
			t.Errorf("Schemas()[%d].Name = %q, want %q", i, s.Name, SchemaFiles[i])
		}
	}
}

func TestSchemaFS(t *testing.T) {
	fsys, err := SchemaFS()
	if err != nil {
		t.Fatalf("SchemaFS(): %v", err)
	}
	b, err := fsys.Open(SchemaFileIncidentReport)
	if err != nil {
		t.Fatalf("open %q via SchemaFS: %v", SchemaFileIncidentReport, err)
	}
	_ = b.Close()
}
