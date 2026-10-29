package fedramp_test

import (
	"encoding/json"
	"fmt"

	fedramp "github.com/opensecurity-dev/fedramp-schemas-go"
)

// ExampleIncidentReport_marshal shows how to build and marshal an incident
// report document.
func ExampleIncidentReport_marshal() {
	report := fedramp.IncidentReport{
		CertificationPackageOverviewURI: "https://example.com/cpo.json",
		ReportType:                      fedramp.IncidentReportTypeInitial,
		ProviderTrackingID:              "INC-2026-0001",
	}
	out, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))
	// Output:
	// {
	//   "certificationPackageOverviewUri": "https://example.com/cpo.json",
	//   "reportType": "Initial",
	//   "providerTrackingId": "INC-2026-0001"
	// }
}

// ExampleIncidentReport_unmarshal shows how to decode an incident report from
// JSON.
func ExampleIncidentReport_unmarshal() {
	data := []byte(`{
	  "certificationPackageOverviewUri": "https://example.com/cpo.json",
	  "reportType": "Final",
	  "providerTrackingId": "INC-2026-0002",
	  "resolvedAt": "2026-06-02T09:00:00Z"
	}`)

	var report fedramp.IncidentReport
	if err := json.Unmarshal(data, &report); err != nil {
		panic(err)
	}
	fmt.Println(report.ReportType, report.ProviderTrackingID, report.ResolvedAt)
	// Output: Final INC-2026-0002 2026-06-02T09:00:00Z
}

// ExampleValidateBytes shows how to validate a document against its embedded
// FedRAMP schema. Here the incident report declares itself Final but omits the
// required resolvedAt field, so validation fails.
func ExampleValidateBytes() {
	data := []byte(`{
	  "certificationPackageOverviewUri": "https://example.com/cpo.json",
	  "reportType": "Final",
	  "providerTrackingId": "INC-2026-0003"
	}`)

	err := fedramp.ValidateBytes(fedramp.SchemaFileIncidentReport, data)
	fmt.Println(err != nil)
	// Output: true
}

// ExampleSchemas lists the embedded FedRAMP schemas and their versions.
func ExampleSchemas() {
	schemas, err := fedramp.Schemas()
	if err != nil {
		panic(err)
	}
	for _, s := range schemas {
		fmt.Printf("%s\t%s\n", s.Version, s.Name)
	}
	// Output:
	// 0.3.0	fedramp-common-definitions-schema-2026-06-24.json
	// 0.1.1	fedramp-accepted-vulnerability-info-schema-2026-06-24.json
	// 1.0.1	fedramp-advisor-information-schema-2026-06-24.json
	// 1.0.1	fedramp-assessor-information-schema-2026-06-24.json
	// 0.1.4	fedramp-certification-package-overview-schema-2026-06-24.json
	// 0.1.1	fedramp-historical-ver-activity-schema-2026-06-24.json
	// 0.2.0	fedramp-incident-report-schema-2026-06-24.json
	// 0.2.0	fedramp-ongoing-certification-report-schema-2026-06-24.json
	// 1.1.1	fedramp-security-decision-record-schema-2026-06-24.json
	// 0.1.2	fedramp-significant-change-notifications-schema-2026-06-24.json
	// 0.1.1	fedramp-vulnerability-detail-report-schema-2026-06-24.json
}

// ExampleCertificationType_IsValid shows enum validation.
func ExampleCertificationType_IsValid() {
	fmt.Println(fedramp.CertificationType20x.IsValid())
	fmt.Println(fedramp.CertificationType("HighBaseline").IsValid())
	// Output:
	// true
	// false
}
