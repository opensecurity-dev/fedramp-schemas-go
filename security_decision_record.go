package fedramp

// ImplementationStatus is the status of a control, requirement, or KSI
// implementation.
type ImplementationStatus string

// Valid ImplementationStatus values.
const (
	ImplementationStatusImplemented          ImplementationStatus = "Implemented"
	ImplementationStatusNotImplemented       ImplementationStatus = "Not Implemented"
	ImplementationStatusPartiallyImplemented ImplementationStatus = "Partially Implemented"
)

// EvidenceType is the type of evidence provided for a security test.
type EvidenceType string

// Valid EvidenceType values.
const (
	EvidenceTypeLog           EvidenceType = "Log"
	EvidenceTypeReport        EvidenceType = "Report"
	EvidenceTypeScreenshot    EvidenceType = "Screenshot"
	EvidenceTypeConfiguration EvidenceType = "Configuration"
	EvidenceTypePolicy        EvidenceType = "Policy"
	EvidenceTypeProcedure     EvidenceType = "Procedure"
	EvidenceTypeAuditRecord   EvidenceType = "Audit Record"
)

// SecurityDecisionRecord is the top-level document for a FedRAMP Security
// Decision Record (SDR-CSO-FRR): the Cloud Service Provider (CSP) system
// submission for FedRAMP certification.
//
// Schema: fedramp-security-decision-record-schema-2026-06-24.json
type SecurityDecisionRecord struct {
	// CertificationPackageOverviewURI is the full URI of the provider's
	// Certification Package Overview document.
	CertificationPackageOverviewURI string `json:"certificationPackageOverviewUri"`
	// Metadata is basic metadata about this Security Decision Record, per
	// SDR-CSO-MTD.
	Metadata *SecurityDecisionRecordMetadata `json:"metadata,omitempty"`
	// PortsAndProtocols are the ports and protocols used by the Cloud Service
	// Offering.
	PortsAndProtocols []PortAndProtocol `json:"portsAndProtocols,omitempty"`
	// SecurityControls describe the NIST 800-53 security controls for the
	// system.
	SecurityControls []SecurityControl `json:"securityControls,omitempty"`
	// FedRAMPRequirements describe how any applicable FedRAMP required
	// processes are met.
	FedRAMPRequirements []FedRAMPRequirement `json:"fedRampRequirements"`
	// KeySecurityIndicators describe how the system meets the FedRAMP Key
	// Security Indicators (KSI).
	KeySecurityIndicators []KeySecurityIndicator `json:"keySecurityIndicators,omitempty"`
}

// SecurityDecisionRecordMetadata is basic metadata about a Security Decision
// Record, per SDR-CSO-MTD.
type SecurityDecisionRecordMetadata struct {
	// Version is the version of this Security Decision Record.
	Version string `json:"version"`
	// LastUpdated is the date and time this Security Decision Record was last
	// updated (RFC 3339 date-time).
	LastUpdated string `json:"lastUpdated"`
	// UpdateSource is the source of the update (e.g., person, team, or
	// automated process).
	UpdateSource string `json:"updateSource"`
}

// PortAndProtocol is a port and protocol exposed by the Cloud Service Offering.
type PortAndProtocol struct {
	// ServiceName is the name of the service exposed on this port.
	ServiceName string `json:"serviceName,omitempty"`
	// PortNumber is the port number used by the service. For services that
	// operate on a range of ports, provide the range in the format
	// 'start-end'.
	PortNumber string `json:"portNumber,omitempty"`
	// TransportProtocol is the transport protocol used by the service.
	TransportProtocol string `json:"transportProtocol,omitempty"`
	// Encryption is the encryption method used by the service (e.g., TLS, SSL,
	// None).
	Encryption string `json:"encryption,omitempty"`
	// Purpose is a general description of how the port and protocol are used in
	// the Cloud Service Offering (e.g., Web access, database connection).
	Purpose string `json:"purpose,omitempty"`
}

// SecurityControl is a NIST 800-53 security control implementation.
type SecurityControl struct {
	// ControlID is the NIST 800-53 Control ID (e.g., AC-1, AC-2, AC-3).
	ControlID string `json:"controlId,omitempty"`
	// ParameterValues are the parameter values for the control.
	ParameterValues []ControlParameterValue `json:"parameterValues,omitempty"`
	// ControlImplementationStatus is the status of the control implementation.
	ControlImplementationStatus ImplementationStatus `json:"controlImplementationStatus,omitempty"`
	// ControlImplementationDescription describes the control implementation.
	// May use Markdown for formatting.
	ControlImplementationDescription string `json:"controlImplementationDescription,omitempty"`
}

// ControlParameterValue is a parameter value for a NIST 800-53 control.
type ControlParameterValue struct {
	// ParameterID is the parameter ID as defined in the NIST 800-53 Control
	// Implementation Guide.
	ParameterID string `json:"parameterId"`
	// ParameterValue is the parameter value selected by the Cloud Service
	// Provider. If the parameter is a boolean, provide 'true' or 'false'. If
	// the parameter is a list, provide a comma-separated list of values.
	ParameterValue string `json:"parameterValue"`
}

// FedRAMPRequirement describes how an applicable FedRAMP required process is
// met.
type FedRAMPRequirement struct {
	// FRRID is the FedRAMP requirement ID.
	FRRID string `json:"frrID"`
	// FRRImplementationStatus is the status of the requirement implementation.
	FRRImplementationStatus ImplementationStatus `json:"frrImplementationStatus,omitempty"`
	// FRRImplementation describes the requirement implementation. May use
	// Markdown for formatting.
	FRRImplementation []string `json:"frrImplementation"`
	// FRRValidation describes the requirement validation. May use Markdown for
	// formatting.
	FRRValidation []string `json:"frrValidation,omitempty"`
	// FRRAssessment describes the requirement assessment. May use Markdown for
	// formatting.
	FRRAssessment []string `json:"frrAssessment,omitempty"`
}

// KeySecurityIndicator describes how a FedRAMP Key Security Indicator (KSI)
// objective is achieved and monitored.
type KeySecurityIndicator struct {
	// KSIID is the FedRAMP KSI ID.
	KSIID string `json:"ksiId"`
	// KSIImplementationStatus is the status of the KSI implementation.
	KSIImplementationStatus ImplementationStatus `json:"ksiImplementationStatus,omitempty"`
	// KSIImplementation describes the Key Security Indicator implementation.
	// May use Markdown for formatting.
	KSIImplementation []string `json:"ksiImplementation"`
	// KSIValidation describes the Key Security Indicator validation. May use
	// Markdown for formatting.
	KSIValidation []string `json:"ksiValidation"`
	// KSIAssessment describes the Key Security Indicator assessment. May use
	// Markdown for formatting.
	KSIAssessment []string `json:"ksiAssessment"`
	// KSITests lists tests used to validate the KSI implementation.
	KSITests []string `json:"ksiTests"`
	// KSIEvidence are the results of the security test used to validate the KSI
	// implementation.
	KSIEvidence []Evidence `json:"ksiEvidence"`
}

// Evidence holds the results of a security test.
type Evidence struct {
	// EvidenceType is the type of evidence provided.
	EvidenceType EvidenceType `json:"evidenceType,omitempty"`
	// EvidenceDescription is a detailed description of the evidence.
	EvidenceDescription string `json:"evidenceDescription,omitempty"`
	// EvidenceLocation is a URI to the evidence document.
	EvidenceLocation string `json:"evidenceLocation,omitempty"`
	// EvidenceText is evidence included as plain text such as the output of a
	// command or log.
	EvidenceText string `json:"evidenceText,omitempty"`
	// LastUpdated is the date when the evidence was last updated (YYYY-MM-DD).
	LastUpdated string `json:"lastUpdated,omitempty"`
}
