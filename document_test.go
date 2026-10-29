package fedramp

import "testing"

func TestSchemaFileMappingsAreKnown(t *testing.T) {
	docs := []Document{
		AcceptedVulnerabilityInfoReport{},
		AdvisorInformation{},
		AssessorInformation{},
		CertificationPackageOverview{},
		HistoricalVERActivity{},
		IncidentReport{},
		OngoingCertificationReport{},
		SecurityDecisionRecord{},
		SignificantChangeNotification{},
		VulnerabilityDetailReport{},
	}
	if len(docs) != len(SchemaFiles)-1 { // -1: common-definitions has no root document
		t.Fatalf("have %d document types, expected %d", len(docs), len(SchemaFiles)-1)
	}
	seen := map[string]bool{}
	for _, d := range docs {
		name := d.SchemaFile()
		if _, err := SchemaBytes(name); err != nil {
			t.Errorf("%T.SchemaFile() = %q, not an embedded schema: %v", d, name, err)
		}
		if seen[name] {
			t.Errorf("%T.SchemaFile() = %q is duplicated", d, name)
		}
		seen[name] = true
	}
}

func TestValidateDocument(t *testing.T) {
	// Valid Initial incident report passes.
	ok := IncidentReport{
		CertificationPackageOverviewURI: "https://example.com/cpo.json",
		ReportType:                      IncidentReportTypeInitial,
		ProviderTrackingID:              "INC-1",
	}
	if err := Validate(ok); err != nil {
		t.Errorf("valid document rejected: %v", err)
	}

	// Final report missing resolvedAt fails via the document's own schema.
	bad := IncidentReport{
		CertificationPackageOverviewURI: "https://example.com/cpo.json",
		ReportType:                      IncidentReportTypeFinal,
		ProviderTrackingID:              "INC-2",
	}
	if err := Validate(bad); err == nil {
		t.Error("expected error for Final report without resolvedAt")
	}
}

func TestMarshalValidated(t *testing.T) {
	ok := IncidentReport{
		CertificationPackageOverviewURI: "https://example.com/cpo.json",
		ReportType:                      IncidentReportTypeInitial,
		ProviderTrackingID:              "INC-3",
	}
	data, err := MarshalValidated(ok)
	if err != nil {
		t.Fatalf("MarshalValidated: %v", err)
	}
	if len(data) == 0 {
		t.Fatal("expected non-empty JSON")
	}

	bad := IncidentReport{
		CertificationPackageOverviewURI: "https://example.com/cpo.json",
		ReportType:                      IncidentReportTypeFinal,
		ProviderTrackingID:              "INC-4",
	}
	if _, err := MarshalValidated(bad); err == nil {
		t.Error("expected MarshalValidated to fail for invalid document")
	}
}

func TestValidatorValidateMethod(t *testing.T) {
	v, err := NewValidator()
	if err != nil {
		t.Fatalf("NewValidator: %v", err)
	}
	doc := VulnerabilityDetailReport{
		CertificationPackageOverviewURI: "https://example.com/cpo.json",
		ReportPeriod:                    ReportPeriodDateTime{From: "2026-01-01T00:00:00Z", To: "2026-03-31T23:59:59Z"},
		Vulnerabilities: []VulnerabilityDetail{
			{
				ProviderTrackingID:       "V-1",
				Detection:                Detection{DetectedAt: "2026-01-15T00:00:00Z", DetectionSource: "scanner"},
				VulnerabilityDescription: "example",
			},
		},
	}
	if err := v.Validate(doc); err != nil {
		t.Errorf("valid document rejected: %v", err)
	}
}
