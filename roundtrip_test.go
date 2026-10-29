package fedramp

import (
	"bytes"
	"encoding/json"
	"testing"
)

// assertRoundTrip marshals v, unmarshals the JSON into a fresh value of the
// same type, re-marshals that, and fails if the two JSON encodings differ. It
// returns the JSON produced from v so callers can make additional assertions.
func assertRoundTrip[T any](t *testing.T, v T) []byte {
	t.Helper()
	first, err := json.Marshal(v)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var back T
	if err := json.Unmarshal(first, &back); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	second, err := json.Marshal(back)
	if err != nil {
		t.Fatalf("re-marshal: %v", err)
	}
	if !bytes.Equal(first, second) {
		t.Fatalf("round-trip mismatch:\n first: %s\nsecond: %s", first, second)
	}
	return first
}

func mustContainKeys(t *testing.T, data []byte, keys ...string) {
	t.Helper()
	var m map[string]json.RawMessage
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatalf("unmarshal to map: %v", err)
	}
	for _, k := range keys {
		if _, ok := m[k]; !ok {
			t.Errorf("expected top-level key %q in %s", k, data)
		}
	}
}

func TestIncidentReportRoundTrip(t *testing.T) {
	ir := IncidentReport{
		CertificationPackageOverviewURI: "https://example.com/cpo.json",
		ReportType:                      IncidentReportTypeFinal,
		ProviderTrackingID:              "INC-2026-0001",
		Timeline: &IncidentTimeline{
			DetectedAt: "2026-06-01T12:00:00Z",
			Milestones: []IncidentMilestone{
				{OccurredAt: "2026-06-01T13:00:00Z", MilestoneDescription: "Contained"},
			},
		},
		PotentialImpact: &IncidentPotentialImpact{
			HistoricalRating: []NRating{NRating2, NRating3},
			CurrentRating:    new(NRating1),
		},
		RelatedCveIDs: []string{"CVE-2026-1234"},
		ResolvedAt:    "2026-06-02T09:00:00Z",
	}
	data := assertRoundTrip(t, ir)
	mustContainKeys(t, data, "certificationPackageOverviewUri", "reportType", "providerTrackingId", "resolvedAt")
}

func TestCertificationPackageOverviewRoundTrip(t *testing.T) {
	cpo := CertificationPackageOverview{
		ServiceIdentification: ServiceIdentification{
			FedRAMPPackageID:   "F1234567890",
			ProviderName:       "Acme Cloud",
			ServiceName:        "Acme Widgets",
			ServiceAcronym:     "AW",
			ServiceDescription: "A widget service.",
			CertificationType:  CertificationType20x,
			Website:            "https://acme.example",
			Logo:               "https://acme.example/logo.png",
		},
		ServiceProperties: ServiceProperties{
			ServiceType:     []ServiceModel{ServiceModelSaaS, ServiceModelPaaS},
			DeploymentModel: DeploymentModelPublicCloud,
			BusinessCategory: []BusinessCategory{
				"Analytics", "Development Tools",
			},
			TrustCenter: &Repository{
				RepositoryType:         []string{"Trust Center"},
				URL:                    "https://acme.example/trust",
				RepositoryDescription:  "Trust center",
				AuthenticationRequired: false,
			},
		},
		ContactInformation: []ContactInfo{
			{ContactType: "Security", ContactEmail: "security@acme.example"},
			{ContactType: "Sales", ContactPhone: "202-555-0123"},
		},
		Assessor: &PackageAssessor{Name: "Assessor Co", AssessorID: "123456"},
	}
	data := assertRoundTrip(t, cpo)
	mustContainKeys(t, data, "serviceIdentification", "serviceProperties", "contactInformation", "assessor")
}

func TestVulnerabilityDetailReportRoundTrip(t *testing.T) {
	rep := VulnerabilityDetailReport{
		CertificationPackageOverviewURI: "https://example.com/cpo.json",
		ReportPeriod: ReportPeriodDateTime{
			From: "2026-01-01T00:00:00Z",
			To:   "2026-03-31T23:59:59Z",
		},
		Vulnerabilities: []VulnerabilityDetail{
			{
				ProviderTrackingID:       "V-1",
				Detection:                Detection{DetectedAt: "2026-01-15T00:00:00Z", DetectionSource: "scanner"},
				VulnerabilityDescription: "Example vuln",
				IsInternetReachable:      new(true),
				IsLikelyExploitable:      new(false),
				CurrentRating:            new(NRating3),
				OverdueStatus:            &OverdueStatus{IsOverdue: true, Explanation: "pending fix"},
				FinalDisposition:         FinalDispositionPartiallyMitigated,
			},
		},
	}
	data := assertRoundTrip(t, rep)
	mustContainKeys(t, data, "certificationPackageOverviewUri", "reportPeriod", "vulnerabilities")
}

func TestSecurityDecisionRecordRoundTrip(t *testing.T) {
	sdr := SecurityDecisionRecord{
		CertificationPackageOverviewURI: "https://example.com/cpo.json",
		Metadata: &SecurityDecisionRecordMetadata{
			Version:      "1.0",
			LastUpdated:  "2026-06-24T00:00:00Z",
			UpdateSource: "automation",
		},
		FedRAMPRequirements: []FedRAMPRequirement{
			{
				FRRID:                   "FRR-1",
				FRRImplementationStatus: ImplementationStatusImplemented,
				FRRImplementation:       []string{"Implemented via X."},
			},
		},
		KeySecurityIndicators: []KeySecurityIndicator{
			{
				KSIID:             "KSI-1",
				KSIImplementation: []string{"impl"},
				KSIValidation:     []string{"val"},
				KSIAssessment:     []string{"assess"},
				KSITests:          []string{"test"},
				KSIEvidence:       []Evidence{{EvidenceType: EvidenceTypeLog, EvidenceText: "log line"}},
			},
		},
	}
	data := assertRoundTrip(t, sdr)
	mustContainKeys(t, data, "certificationPackageOverviewUri", "fedRampRequirements")
}

func TestOptionalPointersOmitted(t *testing.T) {
	// A VulnerabilityDetail with no optional pointers set should not emit the
	// pointer-valued keys.
	vd := VulnerabilityDetail{
		ProviderTrackingID:       "V-2",
		Detection:                Detection{DetectedAt: "2026-01-15T00:00:00Z", DetectionSource: "scanner"},
		VulnerabilityDescription: "Example",
	}
	data, err := json.Marshal(vd)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var m map[string]json.RawMessage
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	for _, k := range []string{"isInternetReachable", "isLikelyExploitable", "currentRating", "overdueStatus"} {
		if _, ok := m[k]; ok {
			t.Errorf("expected optional key %q to be omitted, got %s", k, data)
		}
	}
}
