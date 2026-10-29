package fedramp

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
)

// SchemaSetDate is the FedRAMP schema-set publication date this package binds
// to. The embedded schema files are named with this date suffix.
const SchemaSetDate = "2026-06-24"

//go:embed docs/schemas/*.json
var schemaFS embed.FS

// schemaDir is the directory within schemaFS holding the embedded schemas.
const schemaDir = "docs/schemas"

// Schema identifies a single embedded FedRAMP JSON Schema document.
type Schema struct {
	// Name is the schema file's base name (e.g.
	// "fedramp-incident-report-schema-2026-06-24.json").
	Name string
	// Title is the human-readable schema title.
	Title string
	// Version is the value of the schema's "$schemaVersion" field.
	Version string
}

// Schema file base names for each embedded FedRAMP JSON Schema document.
const (
	SchemaFileCommonDefinitions              = "fedramp-common-definitions-schema-2026-06-24.json"
	SchemaFileAcceptedVulnerabilityInfo      = "fedramp-accepted-vulnerability-info-schema-2026-06-24.json"
	SchemaFileAdvisorInformation             = "fedramp-advisor-information-schema-2026-06-24.json"
	SchemaFileAssessorInformation            = "fedramp-assessor-information-schema-2026-06-24.json"
	SchemaFileCertificationPackageOverview   = "fedramp-certification-package-overview-schema-2026-06-24.json"
	SchemaFileHistoricalVERActivity          = "fedramp-historical-ver-activity-schema-2026-06-24.json"
	SchemaFileIncidentReport                 = "fedramp-incident-report-schema-2026-06-24.json"
	SchemaFileOngoingCertificationReport     = "fedramp-ongoing-certification-report-schema-2026-06-24.json"
	SchemaFileSecurityDecisionRecord         = "fedramp-security-decision-record-schema-2026-06-24.json"
	SchemaFileSignificantChangeNotifications = "fedramp-significant-change-notifications-schema-2026-06-24.json"
	SchemaFileVulnerabilityDetailReport      = "fedramp-vulnerability-detail-report-schema-2026-06-24.json"
)

// SchemaFiles lists the base names of all embedded FedRAMP JSON Schema
// documents, in a stable order.
var SchemaFiles = []string{
	SchemaFileCommonDefinitions,
	SchemaFileAcceptedVulnerabilityInfo,
	SchemaFileAdvisorInformation,
	SchemaFileAssessorInformation,
	SchemaFileCertificationPackageOverview,
	SchemaFileHistoricalVERActivity,
	SchemaFileIncidentReport,
	SchemaFileOngoingCertificationReport,
	SchemaFileSecurityDecisionRecord,
	SchemaFileSignificantChangeNotifications,
	SchemaFileVulnerabilityDetailReport,
}

// SchemaBytes returns the raw JSON Schema document for the given base file
// name (see the SchemaFile* constants). It returns an error if no such schema
// is embedded.
func SchemaBytes(name string) ([]byte, error) {
	b, err := schemaFS.ReadFile(schemaDir + "/" + name)
	if err != nil {
		return nil, fmt.Errorf("fedramp: schema %q not found: %w", name, err)
	}
	return b, nil
}

// SchemaFS returns a read-only filesystem rooted at the embedded schema
// directory, so callers can enumerate or read schemas by name directly.
func SchemaFS() (fs.FS, error) {
	return fs.Sub(schemaFS, schemaDir)
}

// schemaMeta captures the metadata fields common to every FedRAMP schema.
type schemaMeta struct {
	ID            string `json:"$id"`
	SchemaVersion string `json:"$schemaVersion"`
	Title         string `json:"title"`
}

// SchemaMeta parses and returns the title and version metadata for the schema
// with the given base file name (see the SchemaFile* constants).
func SchemaMeta(name string) (Schema, error) {
	b, err := SchemaBytes(name)
	if err != nil {
		return Schema{}, err
	}
	var m schemaMeta
	if err := json.Unmarshal(b, &m); err != nil {
		return Schema{}, fmt.Errorf("fedramp: parsing schema %q: %w", name, err)
	}
	return Schema{Name: name, Title: m.Title, Version: m.SchemaVersion}, nil
}

// Schemas returns metadata for every embedded FedRAMP schema, in the order of
// SchemaFiles.
func Schemas() ([]Schema, error) {
	out := make([]Schema, 0, len(SchemaFiles))
	for _, name := range SchemaFiles {
		s, err := SchemaMeta(name)
		if err != nil {
			return nil, err
		}
		out = append(out, s)
	}
	return out, nil
}
