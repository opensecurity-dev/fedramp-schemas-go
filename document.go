package fedramp

import (
	"encoding/json"
	"fmt"
)

// Document is a top-level FedRAMP document type. Each implementation reports the
// base file name of the JSON Schema it conforms to (one of the SchemaFile*
// constants), which lets Validate and MarshalValidated select the correct
// schema without the caller naming it — removing the risk of pairing a document
// with the wrong schema.
type Document interface {
	// SchemaFile returns the base file name of the document's JSON Schema.
	SchemaFile() string
}

// Compile-time assertions that every root document type implements Document.
var (
	_ Document = AcceptedVulnerabilityInfoReport{}
	_ Document = AdvisorInformation{}
	_ Document = AssessorInformation{}
	_ Document = CertificationPackageOverview{}
	_ Document = HistoricalVERActivity{}
	_ Document = IncidentReport{}
	_ Document = OngoingCertificationReport{}
	_ Document = SecurityDecisionRecord{}
	_ Document = SignificantChangeNotification{}
	_ Document = VulnerabilityDetailReport{}
)

// SchemaFile returns the schema file name for the document.
func (AcceptedVulnerabilityInfoReport) SchemaFile() string {
	return SchemaFileAcceptedVulnerabilityInfo
}

// SchemaFile returns the schema file name for the document.
func (AdvisorInformation) SchemaFile() string { return SchemaFileAdvisorInformation }

// SchemaFile returns the schema file name for the document.
func (AssessorInformation) SchemaFile() string { return SchemaFileAssessorInformation }

// SchemaFile returns the schema file name for the document.
func (CertificationPackageOverview) SchemaFile() string {
	return SchemaFileCertificationPackageOverview
}

// SchemaFile returns the schema file name for the document.
func (HistoricalVERActivity) SchemaFile() string { return SchemaFileHistoricalVERActivity }

// SchemaFile returns the schema file name for the document.
func (IncidentReport) SchemaFile() string { return SchemaFileIncidentReport }

// SchemaFile returns the schema file name for the document.
func (OngoingCertificationReport) SchemaFile() string { return SchemaFileOngoingCertificationReport }

// SchemaFile returns the schema file name for the document.
func (SecurityDecisionRecord) SchemaFile() string { return SchemaFileSecurityDecisionRecord }

// SchemaFile returns the schema file name for the document.
func (SignificantChangeNotification) SchemaFile() string {
	return SchemaFileSignificantChangeNotifications
}

// SchemaFile returns the schema file name for the document.
func (VulnerabilityDetailReport) SchemaFile() string { return SchemaFileVulnerabilityDetailReport }

// Validate validates a document against its own embedded schema, using the
// shared process-wide Validator. It returns nil if the document is valid;
// otherwise the returned error is a *jsonschema.ValidationError.
func Validate(doc Document) error {
	return ValidateValue(doc.SchemaFile(), doc)
}

// Validate validates a document against its own embedded schema using this
// Validator.
func (v *Validator) Validate(doc Document) error {
	return v.ValidateValue(doc.SchemaFile(), doc)
}

// MarshalValidated marshals a document to JSON, returning the bytes only if the
// document is valid against its own embedded schema. It is a convenience for
// producing a document that is guaranteed to conform to its FedRAMP schema.
func MarshalValidated(doc Document) ([]byte, error) {
	data, err := json.Marshal(doc)
	if err != nil {
		return nil, fmt.Errorf("fedramp: marshaling document: %w", err)
	}
	if err := ValidateBytes(doc.SchemaFile(), data); err != nil {
		return nil, err
	}
	return data, nil
}
