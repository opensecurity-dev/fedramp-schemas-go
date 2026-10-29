package fedramp

import "testing"

// These tests exercise fully-populated, schema-valid instances of the document
// types not covered by roundtrip_test.go, checking both JSON round-tripping and
// validation against the embedded schema.

func TestAdvisorInformationRoundTripValidate(t *testing.T) {
	doc := AdvisorInformation{
		AdvisorName:        "Acme Advisors",
		Logo:               "https://acme.example/logo.png",
		A2LAID:             "A-100",
		ServiceDescription: "Advisory services for FedRAMP.",
		ContactInformation: []string{"advisory@acme.example"},
		ServicesOffered: []ServiceOffering{
			{ServiceName: "Readiness Assessment", ServiceDescription: "Pre-assessment gap analysis."},
		},
		CustomerReferences: []string{"Reference A"},
	}
	data := assertRoundTrip(t, doc)
	mustContainKeys(t, data, "advisorName", "logo", "serviceDescription", "contactInformation", "servicesOffered")
	if err := Validate(doc); err != nil {
		t.Errorf("valid AdvisorInformation rejected: %v", err)
	}
}

func TestAssessorInformationRoundTripValidate(t *testing.T) {
	doc := AssessorInformation{
		AssessorName:          "Acme Assessors",
		Logo:                  "https://acme.example/logo.svg",
		A2LAID:                "123456",
		A2LAAccreditationDate: "2025-01-15",
		ServiceDescription:    "Independent assessment services.",
		ContactInformation:    []string{"assess@acme.example"},
		ServicesOffered: []ServiceOffering{
			{ServiceName: "3PAO Assessment"},
		},
	}
	data := assertRoundTrip(t, doc)
	mustContainKeys(t, data, "assessorName", "logo", "a2laId", "a2laAccreditationDate", "servicesOffered")
	if err := Validate(doc); err != nil {
		t.Errorf("valid AssessorInformation rejected: %v", err)
	}
}

func TestSignificantChangeNotificationRoundTripValidate(t *testing.T) {
	doc := SignificantChangeNotification{
		CertificationPackageOverviewURI: "https://example.com/cpo.json",
		ChangeType:                      ChangeTypeTransformative,
		ChangeTypeExplanation:           "New region introduces new boundary.",
		ChangeDescription:               "Add EU region.",
		Reason:                          "Customer demand.",
		CustomerImpact:                  "New data residency options.",
		PlanAndTimeline: &ChangePlanAndTimeline{
			Summary:           "Phased rollout over Q3.",
			PlannedStart:      "2026-07-01",
			PlannedCompletion: "2026-09-30",
			Milestones: []ChangeMilestone{
				{MilestoneDescription: "Region live", TargetDate: "2026-08-15"},
			},
		},
		ImpactedControls: []string{"KSI-CNA-01"},
		ImpactAnalysis:   "See attached BIA.",
	}
	data := assertRoundTrip(t, doc)
	mustContainKeys(t, data, "certificationPackageOverviewUri", "changeType", "changeDescription", "planAndTimeline")
	if err := Validate(doc); err != nil {
		t.Errorf("valid SignificantChangeNotification rejected: %v", err)
	}
}

func TestOngoingCertificationReportRoundTripValidate(t *testing.T) {
	doc := OngoingCertificationReport{
		CertificationPackageOverviewURI: "https://example.com/cpo.json",
		ReportPeriod:                    ReportPeriodDate{From: "2026-04-01", To: "2026-06-30"},
		CertificationDataChanges:        []string{"Updated contact information."},
		PlannedCertificationDataChanges: PlannedCertificationDataChanges{
			PlanningHorizonThrough: "2026-09-30",
			Changes:                []string{"Planned SCG update."},
		},
		AcceptedVulnerabilities: "No new accepted vulnerabilities this period.",
		TransformativeChanges:   []string{"None."},
		UpdatedRecommendations:  []string{"Enable MFA for admin console."},
		ActiveAgencies:          []string{"Agency A", "Agency B"},
		ReportableIncidents: ReportableIncidents{
			Incidents: []IncidentSummary{
				{
					Summary:                "Brief outage.",
					OccurredAt:             "2026-05-01T10:00:00Z",
					ResolvedAt:             "2026-05-01T11:00:00Z",
					IncidentLessonsLearned: []string{"Improved alerting."},
				},
			},
		},
	}
	data := assertRoundTrip(t, doc)
	mustContainKeys(t, data, "certificationPackageOverviewUri", "reportPeriod", "plannedCertificationDataChanges", "reportableIncidents")
	if err := Validate(doc); err != nil {
		t.Errorf("valid OngoingCertificationReport rejected: %v", err)
	}
}

func TestOngoingCertificationReportEmptyIncidentsValid(t *testing.T) {
	// An empty incidents array attests that no reportable incidents occurred.
	doc := OngoingCertificationReport{
		CertificationPackageOverviewURI: "https://example.com/cpo.json",
		ReportPeriod:                    ReportPeriodDate{From: "2026-04-01", To: "2026-06-30"},
		CertificationDataChanges:        []string{},
		PlannedCertificationDataChanges: PlannedCertificationDataChanges{
			PlanningHorizonThrough: "2026-09-30",
			Changes:                []string{},
		},
		AcceptedVulnerabilities: "None.",
		TransformativeChanges:   []string{},
		UpdatedRecommendations:  []string{},
		ActiveAgencies:          []string{},
		ReportableIncidents:     ReportableIncidents{Incidents: []IncidentSummary{}},
	}
	if err := Validate(doc); err != nil {
		t.Errorf("report with empty incidents rejected: %v", err)
	}
}

func TestHistoricalVERActivityRoundTripValidate(t *testing.T) {
	doc := HistoricalVERActivity{
		CertificationPackageOverviewURI: "https://example.com/cpo.json",
		GeneratedAt:                     "2026-06-30T23:59:59Z",
		ActiveVulnerabilities: []VulnerabilityDetail{
			{
				ProviderTrackingID:       "V-100",
				Detection:                Detection{DetectedAt: "2026-06-01T00:00:00Z", DetectionSource: "scanner"},
				VulnerabilityDescription: "Active vuln.",
				CurrentRating:            new(NRating4),
			},
		},
		AcceptedVulnerabilities: []AcceptedVulnerabilityInfo{
			{
				VulnerabilityDetail: VulnerabilityDetail{
					ProviderTrackingID:       "V-200",
					Detection:                Detection{DetectedAt: "2026-05-01T00:00:00Z", DetectionSource: "pentest"},
					VulnerabilityDescription: "Accepted vuln.",
					FinalDisposition:         FinalDispositionPartiallyMitigated,
				},
				AcceptanceRationale: "Compensating controls in place.",
			},
		},
	}
	data := assertRoundTrip(t, doc)
	mustContainKeys(t, data, "certificationPackageOverviewUri", "generatedAt", "activeVulnerabilities", "acceptedVulnerabilities")
	if err := Validate(doc); err != nil {
		t.Errorf("valid HistoricalVERActivity rejected: %v", err)
	}
}
