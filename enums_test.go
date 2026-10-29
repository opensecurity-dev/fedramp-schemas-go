package fedramp

import "testing"

func TestEnumIsValid(t *testing.T) {
	tests := []struct {
		name  string
		valid bool
		got   bool
	}{
		{"NRating3", true, NRating3.IsValid()},
		{"NRating0", false, NRating(0).IsValid()},
		{"NRating6", false, NRating(6).IsValid()},
		{"CertificationType20x", true, CertificationType20x.IsValid()},
		{"CertificationType bad", false, CertificationType("HighBaseline").IsValid()},
		{"ServiceModelSaaS", true, ServiceModelSaaS.IsValid()},
		{"ServiceModel bad", false, ServiceModel("FaaS").IsValid()},
		{"DeploymentModelPublicCloud", true, DeploymentModelPublicCloud.IsValid()},
		{"DeploymentModel bad", false, DeploymentModel("Private").IsValid()},
		{"ImplementationStatusImplemented", true, ImplementationStatusImplemented.IsValid()},
		{"ImplementationStatus bad", false, ImplementationStatus("Maybe").IsValid()},
		{"EvidenceTypeLog", true, EvidenceTypeLog.IsValid()},
		{"EvidenceType bad", false, EvidenceType("Video").IsValid()},
		{"ChangeTypeAdaptive", true, ChangeTypeAdaptive.IsValid()},
		{"ChangeType bad", false, ChangeType("Minor").IsValid()},
		{"IncidentReportTypeFinal", true, IncidentReportTypeFinal.IsValid()},
		{"IncidentReportType bad", false, IncidentReportType("Interim").IsValid()},
		{"FinalDispositionRemediated", true, FinalDispositionRemediated.IsValid()},
		{"FinalDisposition bad", false, FinalDisposition("Ignored").IsValid()},
	}
	for _, tt := range tests {
		if tt.got != tt.valid {
			t.Errorf("%s: IsValid() = %v, want %v", tt.name, tt.got, tt.valid)
		}
	}
}

func TestEnumValuesAllValid(t *testing.T) {
	for _, v := range NRating(0).Values() {
		if !v.IsValid() {
			t.Errorf("NRating value %d reported invalid", v)
		}
	}
	for _, v := range (DeploymentModel("")).Values() {
		if !v.IsValid() {
			t.Errorf("DeploymentModel value %q reported invalid", v)
		}
	}
	if got := len((EvidenceType("")).Values()); got != 7 {
		t.Errorf("EvidenceType Values() len = %d, want 7", got)
	}
}
