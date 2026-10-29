package fedramp

import "slices"

// This file adds Values() and IsValid() helpers for the enumerated types
// defined across the package. Values() returns the permitted values in schema
// order; IsValid() reports whether a value is one of them.

// Values returns the permitted NRating values.
func (NRating) Values() []NRating {
	return []NRating{NRating1, NRating2, NRating3, NRating4, NRating5}
}

// IsValid reports whether n is a permitted N-Rating (N1–N5).
func (n NRating) IsValid() bool { return slices.Contains(NRating(0).Values(), n) }

// Values returns the permitted FinalDisposition values.
func (FinalDisposition) Values() []FinalDisposition {
	return []FinalDisposition{
		FinalDispositionFullyMitigated,
		FinalDispositionPartiallyMitigated,
		FinalDispositionRemediated,
		FinalDispositionFalsePositive,
	}
}

// IsValid reports whether d is a permitted FinalDisposition.
func (d FinalDisposition) IsValid() bool { return slices.Contains(FinalDisposition("").Values(), d) }

// Values returns the permitted CertificationType values.
func (CertificationType) Values() []CertificationType {
	return []CertificationType{CertificationType20x, CertificationTypeRev5}
}

// IsValid reports whether c is a permitted CertificationType.
func (c CertificationType) IsValid() bool { return slices.Contains(CertificationType("").Values(), c) }

// Values returns the permitted ServiceModel values.
func (ServiceModel) Values() []ServiceModel {
	return []ServiceModel{ServiceModelSaaS, ServiceModelPaaS, ServiceModelIaaS}
}

// IsValid reports whether m is a permitted ServiceModel.
func (m ServiceModel) IsValid() bool { return slices.Contains(ServiceModel("").Values(), m) }

// Values returns the permitted DeploymentModel values.
func (DeploymentModel) Values() []DeploymentModel {
	return []DeploymentModel{
		DeploymentModelPublicCloud,
		DeploymentModelGovernmentOnlyCloud,
		DeploymentModelHybridCloud,
		DeploymentModelCommunityCloud,
		DeploymentModelGovernmentCommunityCloud,
	}
}

// IsValid reports whether m is a permitted DeploymentModel.
func (m DeploymentModel) IsValid() bool { return slices.Contains(DeploymentModel("").Values(), m) }

// Values returns the permitted ImplementationStatus values.
func (ImplementationStatus) Values() []ImplementationStatus {
	return []ImplementationStatus{
		ImplementationStatusImplemented,
		ImplementationStatusNotImplemented,
		ImplementationStatusPartiallyImplemented,
	}
}

// IsValid reports whether s is a permitted ImplementationStatus.
func (s ImplementationStatus) IsValid() bool {
	return slices.Contains(ImplementationStatus("").Values(), s)
}

// Values returns the permitted EvidenceType values.
func (EvidenceType) Values() []EvidenceType {
	return []EvidenceType{
		EvidenceTypeLog,
		EvidenceTypeReport,
		EvidenceTypeScreenshot,
		EvidenceTypeConfiguration,
		EvidenceTypePolicy,
		EvidenceTypeProcedure,
		EvidenceTypeAuditRecord,
	}
}

// IsValid reports whether e is a permitted EvidenceType.
func (e EvidenceType) IsValid() bool { return slices.Contains(EvidenceType("").Values(), e) }

// Values returns the permitted ChangeType values.
func (ChangeType) Values() []ChangeType {
	return []ChangeType{ChangeTypeAdaptive, ChangeTypeTransformative}
}

// IsValid reports whether c is a permitted ChangeType.
func (c ChangeType) IsValid() bool { return slices.Contains(ChangeType("").Values(), c) }

// Values returns the permitted IncidentReportType values.
func (IncidentReportType) Values() []IncidentReportType {
	return []IncidentReportType{
		IncidentReportTypeInitial,
		IncidentReportTypeOngoing,
		IncidentReportTypeFinal,
	}
}

// IsValid reports whether t is a permitted IncidentReportType.
func (t IncidentReportType) IsValid() bool {
	return slices.Contains(IncidentReportType("").Values(), t)
}
