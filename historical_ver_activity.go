package fedramp

// HistoricalVERActivity is the top-level document for FedRAMP Historical
// Vulnerability Evaluation and Reporting Activity (VER-TFR-MRH): a
// machine-readable snapshot of historical VER activity for automated retrieval.
//
// Schema: fedramp-historical-ver-activity-schema-2026-06-24.json
type HistoricalVERActivity struct {
	// CertificationPackageOverviewURI is the full URI of the provider's
	// Certification Package Overview document.
	CertificationPackageOverviewURI string `json:"certificationPackageOverviewUri"`
	// GeneratedAt is when this snapshot was generated (RFC 3339 date-time).
	GeneratedAt string `json:"generatedAt"`
	// ActiveVulnerabilities lists all non-accepted vulnerabilities. Each record
	// follows the fields required by VER-RPT-VDT.
	ActiveVulnerabilities []VulnerabilityDetail `json:"activeVulnerabilities"`
	// AcceptedVulnerabilities lists all accepted vulnerabilities. Each record
	// follows the fields required by VER-RPT-AVI.
	AcceptedVulnerabilities []AcceptedVulnerabilityInfo `json:"acceptedVulnerabilities"`
}
