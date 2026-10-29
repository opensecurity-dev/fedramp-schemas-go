package fedramp

// OngoingCertificationReport is the top-level document for a FedRAMP Ongoing
// Certification Report (CCM-OCR-AVL): a quarterly report covering the entire
// period since the previous report.
//
// Schema: fedramp-ongoing-certification-report-schema-2026-06-24.json
type OngoingCertificationReport struct {
	// CertificationPackageOverviewURI is the full URI of the provider's
	// Certification Package Overview document.
	CertificationPackageOverviewURI string `json:"certificationPackageOverviewUri"`
	// ReportPeriod is the 3-month period covered by this report.
	ReportPeriod ReportPeriodDate `json:"reportPeriod"`
	// CertificationDataChanges are high-level summaries of changes to FedRAMP
	// Certification Data since the previous report.
	CertificationDataChanges []string `json:"certificationDataChanges"`
	// PlannedCertificationDataChanges are high-level summaries of planned
	// changes to FedRAMP Certification Data for at least the next 3 months.
	PlannedCertificationDataChanges PlannedCertificationDataChanges `json:"plannedCertificationDataChanges"`
	// AcceptedVulnerabilities summarizes accepted vulnerabilities. Full records
	// are reported per VER-RPT-AVI.
	AcceptedVulnerabilities string `json:"acceptedVulnerabilities"`
	// TransformativeChanges are high-level summaries of transformative changes
	// during this period.
	TransformativeChanges []string `json:"transformativeChanges"`
	// UpdatedRecommendations are updated recommendations or best practices for
	// security, configuration, usage, or similar aspects of the cloud service
	// offering.
	UpdatedRecommendations []string `json:"updatedRecommendations"`
	// ActiveAgencies lists all agencies directly using the product.
	ActiveAgencies []string `json:"activeAgencies"`
	// ReportableIncidents are the FedRAMP Reportable Incidents during this
	// period. An empty Incidents slice attests that none occurred.
	ReportableIncidents ReportableIncidents `json:"reportableIncidents"`
}

// PlannedCertificationDataChanges holds high-level summaries of planned changes
// to FedRAMP Certification Data.
type PlannedCertificationDataChanges struct {
	// PlanningHorizonThrough is the date through which planned changes are
	// covered, at least 3 months from report end (YYYY-MM-DD).
	PlanningHorizonThrough string `json:"planningHorizonThrough"`
	// Changes summarizes planned changes to FedRAMP Certification Data.
	Changes []string `json:"changes"`
}

// ReportableIncidents holds the FedRAMP Reportable Incidents for a reporting
// period. An empty Incidents slice attests that none occurred.
type ReportableIncidents struct {
	// Incidents summarizes each incident. Provide an empty array to attest that
	// no FedRAMP Reportable Incidents occurred during this period.
	Incidents []IncidentSummary `json:"incidents"`
}

// IncidentSummary is a summary of a FedRAMP Reportable Incident.
type IncidentSummary struct {
	// Summary is a brief description of the incident.
	Summary string `json:"summary,omitempty"`
	// OccurredAt is the date and time the incident occurred (RFC 3339
	// date-time).
	OccurredAt string `json:"occurredAt,omitempty"`
	// ResolvedAt is the date and time the incident was resolved (RFC 3339
	// date-time).
	ResolvedAt string `json:"resolvedAt,omitempty"`
	// IncidentLessonsLearned are lessons learned and changes planned or made as
	// a result of FedRAMP Reportable Incidents.
	IncidentLessonsLearned []string `json:"incidentLessonsLearned,omitempty"`
}
