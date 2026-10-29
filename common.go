package fedramp

// This file contains the shared type definitions from the FedRAMP Common
// Definitions schema (fedramp-common-definitions-schema-2026-06-24.json),
// which other FedRAMP schemas reference via $ref.

// NRating is a Potential Agency Impact N-Rating (N1–N5). See VER-EVA-EPA or
// IEC-CSO-EFI for the N1–N5 definitions.
type NRating int

// Valid N-Rating values.
const (
	NRating1 NRating = 1
	NRating2 NRating = 2
	NRating3 NRating = 3
	NRating4 NRating = 4
	NRating5 NRating = 5
)

// FinalDisposition is the final disposition of a vulnerability. It is omitted
// while the vulnerability is still active.
type FinalDisposition string

// Valid FinalDisposition values.
const (
	FinalDispositionFullyMitigated     FinalDisposition = "Fully Mitigated"
	FinalDispositionPartiallyMitigated FinalDisposition = "Partially Mitigated"
	FinalDispositionRemediated         FinalDisposition = "Remediated"
	FinalDispositionFalsePositive      FinalDisposition = "False Positive"
)

// Detection holds the required information for a single detected vulnerability
// per VER-RPT-VDT.
type Detection struct {
	// DetectedAt is when the vulnerability was detected (RFC 3339 date-time).
	DetectedAt string `json:"detectedAt"`
	// DetectionSource is the tool or activity that identified the vulnerability
	// (e.g., scanner, penetration test, bug bounty).
	DetectionSource string `json:"detectionSource"`
}

// PainReductionEvent holds the required information for a single PAIN reduction
// event per VER-RPT-VDT.
type PainReductionEvent struct {
	// ReducedAt is when the PAIN reduction was completed (RFC 3339 date-time).
	ReducedAt string `json:"reducedAt"`
	// Rating is the N-rating after the PAIN reduction.
	Rating NRating `json:"rating"`
}

// ProjectedNextReduction is the next planned mitigation step. It is omitted if
// the vulnerability is fully mitigated or no reduction is planned.
type ProjectedNextReduction struct {
	// EstimatedAt is when the next reduction is expected (RFC 3339 date-time).
	EstimatedAt string `json:"estimatedAt"`
	// TargetRating is the next Potential Agency Impact N-rating.
	TargetRating NRating `json:"targetRating"`
}

// OverdueStatus describes whether a vulnerability is overdue. See VER-TFR-EVU
// and VER-TFR-MAV for the overdue definitions.
type OverdueStatus struct {
	// IsOverdue is true if the vulnerability is overdue.
	IsOverdue bool `json:"isOverdue"`
	// Explanation is the reason for the overdue status. It is required when
	// IsOverdue is true.
	Explanation string `json:"explanation,omitempty"`
}

// VulnerabilityDetail holds the required information for a single detected
// vulnerability per VER-RPT-VDT.
type VulnerabilityDetail struct {
	// ProviderTrackingID is the provider's internally assigned identifier for
	// this vulnerability.
	ProviderTrackingID string `json:"providerTrackingId"`
	// Detection describes when and how the vulnerability was detected.
	Detection Detection `json:"detection"`
	// VulnerabilityDescription describes the vulnerability.
	VulnerabilityDescription string `json:"vulnerabilityDescription"`
	// PotentialAgencyImpact describes the potential impact.
	PotentialAgencyImpact string `json:"potentialAgencyImpact,omitempty"`
	// EvaluationCompletedAt is when evaluation was completed (RFC 3339
	// date-time). The clock for VER-TFR-MAV runs from this date.
	EvaluationCompletedAt string `json:"evaluationCompletedAt,omitempty"`
	// IsInternetReachable is the IRV classification per VER-EVA-EIR.
	IsInternetReachable *bool `json:"isInternetReachable,omitempty"`
	// IsLikelyExploitable is the LEV classification per VER-EVA-ELX and
	// VER-EVA-AIA.
	IsLikelyExploitable *bool `json:"isLikelyExploitable,omitempty"`
	// CurrentRating is the currently estimated Potential Agency Impact
	// N-rating.
	CurrentRating *NRating `json:"currentRating,omitempty"`
	// ProjectedNextReduction is the next planned mitigation step.
	ProjectedNextReduction *ProjectedNextReduction `json:"projectedNextReduction,omitempty"`
	// OverdueStatus indicates whether the vulnerability is overdue.
	OverdueStatus *OverdueStatus `json:"overdueStatus,omitempty"`
	// SupplementaryRiskInformation per VER-RPT-VDT field 11 and VER-RPT-NID.
	SupplementaryRiskInformation string `json:"supplementaryRiskInformation,omitempty"`
	// PainReductionEvents is the history of completed PAIN-rating reductions
	// for this vulnerability per VER-RPT-VDT.
	PainReductionEvents []PainReductionEvent `json:"painReductionEvents,omitempty"`
	// FinalDisposition is omitted if the vulnerability is still active.
	FinalDisposition FinalDisposition `json:"finalDisposition,omitempty"`
}

// AcceptedVulnerabilityInfo holds the required information for a single accepted
// vulnerability per VER-RPT-AVI.
type AcceptedVulnerabilityInfo struct {
	// VulnerabilityDetail is the detail record for the accepted vulnerability.
	VulnerabilityDetail VulnerabilityDetail `json:"vulnerabilityDetail"`
	// AcceptanceRationale explains why this is an accepted vulnerability.
	AcceptanceRationale string `json:"acceptanceRationale"`
}

// ReportPeriodDateTime is the From and To bounds of a report period, using
// RFC 3339 date-time values.
type ReportPeriodDateTime struct {
	// From is the start date-time for the report period (RFC 3339 date-time).
	From string `json:"from"`
	// To is the end date-time for the report period (RFC 3339 date-time).
	To string `json:"to"`
}

// ReportPeriodDate is the From and To bounds of a report period, using
// date-only (YYYY-MM-DD) values.
type ReportPeriodDate struct {
	// From is the start date for the report period (YYYY-MM-DD).
	From string `json:"from"`
	// To is the end date for the report period (YYYY-MM-DD).
	To string `json:"to"`
}
