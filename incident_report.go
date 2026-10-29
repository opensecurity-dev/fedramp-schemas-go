package fedramp

// IncidentReportType identifies which incident report type in the IEC-CSO
// lifecycle a report represents.
type IncidentReportType string

// Valid IncidentReportType values.
const (
	// IncidentReportTypeInitial is an Initial Incident Report (IEC-CSO-IIR).
	IncidentReportTypeInitial IncidentReportType = "Initial"
	// IncidentReportTypeOngoing is an Ongoing Incident Report (IEC-CSO-OIR).
	IncidentReportTypeOngoing IncidentReportType = "Ongoing"
	// IncidentReportTypeFinal is a Final Incident Report (IEC-CSO-FIR).
	IncidentReportTypeFinal IncidentReportType = "Final"
)

// IncidentReport is the top-level document for a FedRAMP Incident Report. It is
// a unified schema for the three incident report types in the IEC-CSO
// lifecycle: Initial (IEC-CSO-IIR), Ongoing (IEC-CSO-OIR), and Final
// (IEC-CSO-FIR).
//
// Schema: fedramp-incident-report-schema-2026-06-24.json
type IncidentReport struct {
	// CertificationPackageOverviewURI is the full URI of the provider's
	// Certification Package Overview document.
	CertificationPackageOverviewURI string `json:"certificationPackageOverviewUri"`
	// ReportType is Initial (IEC-CSO-IIR), Ongoing (IEC-CSO-OIR), or Final
	// (IEC-CSO-FIR).
	ReportType IncidentReportType `json:"reportType"`
	// ProviderTrackingID is the provider's internally assigned identifier for
	// this incident. It must be consistent across IIR, OIR, and FIR reports for
	// the same incident.
	ProviderTrackingID string `json:"providerTrackingId"`
	// FederalIncidentCoordinator is contact information for the provider's
	// federal incident response coordinator.
	FederalIncidentCoordinator string `json:"federalIncidentCoordinator,omitempty"`
	// IncidentDescription describes the incident.
	IncidentDescription string `json:"incidentDescription,omitempty"`
	// Timeline is the timeline of the incident per IEC-CSO-IIR field 4.
	Timeline *IncidentTimeline `json:"timeline,omitempty"`
	// PotentialImpact describes the current and historical potential impact to
	// the system.
	PotentialImpact *IncidentPotentialImpact `json:"potentialImpact,omitempty"`
	// FunctionalImpact describes the actual or potential impact to federal
	// agency customers.
	FunctionalImpact string `json:"functionalImpact,omitempty"`
	// RecoveryPlan is the estimated recovery plan, milestones, and timelines.
	RecoveryPlan string `json:"recoveryPlan,omitempty"`
	// AffectedAgencies lists customer agencies that may be affected by the
	// incident.
	AffectedAgencies []string `json:"affectedAgencies,omitempty"`
	// ObservedActivity is the observed incident activity per IEC-CSO-OIR.
	ObservedActivity string `json:"observedActivity,omitempty"`
	// IndicatorsOfCompromise per IEC-CSO-OIR.
	IndicatorsOfCompromise []string `json:"indicatorsOfCompromise,omitempty"`
	// RelatedCveIDs per IEC-CSO-OIR. Include if applicable. Each value matches
	// the pattern CVE-YYYY-NNNN(...).
	RelatedCveIDs []string `json:"relatedCveIds,omitempty"`
	// RootCause per IEC-CSO-OIR.
	RootCause string `json:"rootCause,omitempty"`
	// ResponseAndRecoveryActivities per IEC-CSO-OIR.
	ResponseAndRecoveryActivities string `json:"responseAndRecoveryActivities,omitempty"`
	// ResolvedAt is when the incident was resolved and recovery completed
	// (RFC 3339 date-time). Required for Final reports (IEC-CSO-FIR).
	ResolvedAt string `json:"resolvedAt,omitempty"`
}

// IncidentTimeline is the timeline of an incident per IEC-CSO-IIR field 4.
type IncidentTimeline struct {
	// StartedAt is when the incident started (RFC 3339 date-time).
	StartedAt string `json:"startedAt,omitempty"`
	// DetectedAt is when the incident was detected (RFC 3339 date-time).
	DetectedAt string `json:"detectedAt,omitempty"`
	// DetectionSource is the source of the incident detection.
	DetectionSource string `json:"detectionSource,omitempty"`
	// EvaluationCompletedAt is when the FedRAMP reportable incident evaluation
	// was completed (RFC 3339 date-time).
	EvaluationCompletedAt string `json:"evaluationCompletedAt,omitempty"`
	// Milestones are major events in the incident lifecycle.
	Milestones []IncidentMilestone `json:"milestones,omitempty"`
}

// IncidentMilestone is a major event in the incident lifecycle.
type IncidentMilestone struct {
	// OccurredAt is when the milestone occurred (RFC 3339 date-time).
	OccurredAt string `json:"occurredAt"`
	// MilestoneDescription describes the milestone.
	MilestoneDescription string `json:"milestoneDescription"`
}

// IncidentPotentialImpact describes the current and historical potential impact
// of an incident to the system.
type IncidentPotentialImpact struct {
	// HistoricalRating is the historical rating of the incident.
	HistoricalRating []NRating `json:"historicalRating,omitempty"`
	// CurrentRating is the current rating of the incident.
	CurrentRating *NRating `json:"currentRating,omitempty"`
	// EvaluationNotes explains the PAIN evaluation per IEC-CSO-EFI. Required if
	// applicable.
	EvaluationNotes string `json:"evaluationNotes,omitempty"`
}
