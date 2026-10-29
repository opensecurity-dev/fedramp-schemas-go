package fedramp

import (
	"strings"
	"testing"
)

func TestValidatorAllSchemasCompile(t *testing.T) {
	v, err := NewValidator()
	if err != nil {
		t.Fatalf("NewValidator: %v", err)
	}
	for _, name := range SchemaFiles {
		if _, err := v.schema(name); err != nil {
			t.Errorf("compile %q: %v", name, err)
		}
	}
}

func TestValidateValueValidDocuments(t *testing.T) {
	v, err := NewValidator()
	if err != nil {
		t.Fatalf("NewValidator: %v", err)
	}

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
			ServiceType:     []ServiceModel{ServiceModelSaaS},
			DeploymentModel: DeploymentModelPublicCloud,
		},
		ContactInformation: []ContactInfo{
			{ContactType: "Security", ContactEmail: "security@acme.example"},
			{ContactType: "Sales", ContactPhone: "202-555-0123"},
		},
	}
	if err := v.ValidateValue(SchemaFileCertificationPackageOverview, cpo); err != nil {
		t.Errorf("valid certification package overview rejected: %v", err)
	}
}

func TestValidateConditionalRules(t *testing.T) {
	v, err := NewValidator()
	if err != nil {
		t.Fatalf("NewValidator: %v", err)
	}

	t.Run("final incident requires resolvedAt", func(t *testing.T) {
		// A Final report without resolvedAt must fail the if/then rule.
		bad := IncidentReport{
			CertificationPackageOverviewURI: "https://example.com/cpo.json",
			ReportType:                      IncidentReportTypeFinal,
			ProviderTrackingID:              "INC-1",
		}
		if err := v.ValidateValue(SchemaFileIncidentReport, bad); err == nil {
			t.Error("expected error for Final incident report without resolvedAt")
		}

		// Adding resolvedAt satisfies it.
		bad.ResolvedAt = "2026-06-02T09:00:00Z"
		if err := v.ValidateValue(SchemaFileIncidentReport, bad); err != nil {
			t.Errorf("Final incident report with resolvedAt rejected: %v", err)
		}

		// An Initial report without resolvedAt is fine.
		initial := IncidentReport{
			CertificationPackageOverviewURI: "https://example.com/cpo.json",
			ReportType:                      IncidentReportTypeInitial,
			ProviderTrackingID:              "INC-2",
		}
		if err := v.ValidateValue(SchemaFileIncidentReport, initial); err != nil {
			t.Errorf("Initial incident report rejected: %v", err)
		}
	})

	t.Run("contacts must include security and sales", func(t *testing.T) {
		cpo := CertificationPackageOverview{
			ServiceIdentification: ServiceIdentification{
				FedRAMPPackageID:   "F1",
				ProviderName:       "Acme",
				ServiceName:        "Widgets",
				ServiceAcronym:     "AW",
				ServiceDescription: "desc",
				CertificationType:  CertificationTypeRev5,
				Website:            "https://acme.example",
				Logo:               "https://acme.example/logo.png",
			},
			ServiceProperties: ServiceProperties{
				ServiceType:     []ServiceModel{ServiceModelIaaS},
				DeploymentModel: DeploymentModelHybridCloud,
			},
			// Only a Security contact; missing Sales -> must fail.
			ContactInformation: []ContactInfo{
				{ContactType: "Security", ContactEmail: "sec@acme.example"},
			},
		}
		if err := v.ValidateValue(SchemaFileCertificationPackageOverview, cpo); err == nil {
			t.Error("expected error when Sales contact is missing")
		}
	})
}

func TestValidateBytesInvalidEnum(t *testing.T) {
	// certificationType outside the enum must be rejected.
	doc := `{
	  "serviceIdentification": {
	    "fedRampPackageId": "F1", "providerName": "Acme", "serviceName": "W",
	    "serviceAcronym": "AW", "serviceDescription": "d",
	    "certificationType": "NotAType",
	    "website": "https://acme.example", "logo": "https://acme.example/logo.png"
	  },
	  "serviceProperties": { "serviceType": ["SaaS"], "deploymentModel": "Public Cloud" },
	  "contactInformation": [
	    {"contactType": "Security"}, {"contactType": "Sales"}
	  ]
	}`
	err := ValidateBytes(SchemaFileCertificationPackageOverview, []byte(doc))
	if err == nil {
		t.Fatal("expected error for invalid certificationType enum value")
	}
	if !strings.Contains(err.Error(), "certificationType") {
		t.Errorf("error should mention certificationType, got: %v", err)
	}
}

func TestValidateBytesUnknownSchema(t *testing.T) {
	if err := ValidateBytes("nope.json", []byte(`{}`)); err == nil {
		t.Fatal("expected error for unknown schema file")
	}
}
