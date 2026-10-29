package fedramp

// ChangeType is the significant change type per SCN-CSO-EVA. Routine recurring
// changes (SCN-RTR) do not require a Significant Change Notification.
type ChangeType string

// Valid ChangeType values.
const (
	ChangeTypeAdaptive       ChangeType = "Adaptive"
	ChangeTypeTransformative ChangeType = "Transformative"
)

// SignificantChangeNotification is the top-level document for a FedRAMP
// Significant Change Notification (SCN-CSO-INF). Per the rule, the structure of
// the information may vary depending on how the provider tracks this
// internally.
//
// Schema: fedramp-significant-change-notifications-schema-2026-06-24.json
type SignificantChangeNotification struct {
	// CertificationPackageOverviewURI is the full URI of the provider's
	// Certification Package Overview document.
	CertificationPackageOverviewURI string `json:"certificationPackageOverviewUri"`
	// AssessorName is the name of the assessor involved in this change, if
	// applicable.
	AssessorName string `json:"assessorName,omitempty"`
	// RelatedVulnerability is an identifier or description of any related
	// vulnerability driving this change, if applicable.
	RelatedVulnerability string `json:"relatedVulnerability,omitempty"`
	// ChangeType is the significant change type per SCN-CSO-EVA.
	ChangeType ChangeType `json:"changeType"`
	// ChangeTypeExplanation explains why this change was categorized as the
	// selected type.
	ChangeTypeExplanation string `json:"changeTypeExplanation,omitempty"`
	// ChangeDescription is a short description of the change.
	ChangeDescription string `json:"changeDescription"`
	// Reason is the reason for the change.
	Reason string `json:"reason,omitempty"`
	// CustomerImpact summarizes customer impact, including changes to services
	// and customer configuration responsibilities.
	CustomerImpact string `json:"customerImpact,omitempty"`
	// PlanAndTimeline is the plan and timeline for the change, including
	// verification, assessment, and/or validation of impacted Key Security
	// Indicators (KSIs) or Rev5 Controls.
	PlanAndTimeline *ChangePlanAndTimeline `json:"planAndTimeline,omitempty"`
	// ImpactedControls lists KSI or control identifiers that will be verified,
	// assessed, or validated as part of this change.
	ImpactedControls []string `json:"impactedControls,omitempty"`
	// ImpactAnalysis is a copy or summary of the business or security impact
	// analysis. It may be the full text or a URI referencing the document.
	ImpactAnalysis string `json:"impactAnalysis,omitempty"`
}

// ChangePlanAndTimeline describes the plan and timeline for a significant
// change.
type ChangePlanAndTimeline struct {
	// Summary is a human-readable summary of the timeline and plan for the
	// change. May use Markdown for formatting.
	Summary string `json:"summary"`
	// PlannedStart is the first date on which changes will be made to the
	// system (YYYY-MM-DD).
	PlannedStart string `json:"plannedStart,omitempty"`
	// PlannedCompletion is the date on which the changes will be considered
	// complete, including any planned testing and monitoring activities
	// (YYYY-MM-DD).
	PlannedCompletion string `json:"plannedCompletion,omitempty"`
	// Milestones are the milestones for the change, including dates and
	// descriptions.
	Milestones []ChangeMilestone `json:"milestones,omitempty"`
}

// ChangeMilestone is a milestone for a significant change.
type ChangeMilestone struct {
	// MilestoneDescription is a human-readable description of the milestone.
	MilestoneDescription string `json:"milestoneDescription"`
	// TargetDate is the date on which the milestone is expected to be completed
	// (YYYY-MM-DD).
	TargetDate string `json:"targetDate,omitempty"`
}
