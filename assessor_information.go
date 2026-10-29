package fedramp

// AssessorInformation is the top-level document for FedRAMP Assessor Information
// (MKT-IAS-WEB): the public information an Independent Assessment Service must
// supply in machine-readable format on their website.
//
// Schema: fedramp-assessor-information-schema-2026-06-24.json
type AssessorInformation struct {
	// AssessorName is the name of the independent assessment organization as it
	// should be displayed on the FedRAMP Marketplace (MKT-IAS-WEB).
	AssessorName string `json:"assessorName"`
	// Logo is the URL of the assessor logo displayed on the FedRAMP Marketplace
	// (MKT-IAS-WEB). It must point to a PNG, JPEG, GIF, SVG, WebP, ICO, BMP, or
	// TIFF file.
	Logo string `json:"logo"`
	// A2LAID is the identifier assigned to the assessor by the American
	// Association for Laboratory Accreditation. Required for FedRAMP
	// Recognition per REC-IAS-ACC.
	A2LAID string `json:"a2laId"`
	// A2LAAccreditationDate is the date the American Association for Laboratory
	// Accreditation accredited this assessor, in YYYY-MM-DD format
	// (REC-IAS-ACC).
	A2LAAccreditationDate string `json:"a2laAccreditationDate"`
	// ServiceDescription is a general description of the independent assessment
	// service (MKT-IAS-WEB).
	ServiceDescription string `json:"serviceDescription"`
	// ContactInformation is contact information for the assessor organization
	// (MKT-IAS-WEB).
	ContactInformation []string `json:"contactInformation"`
	// ServicesOffered lists the types of independent assessment services the
	// assessor offers (MKT-IAS-WEB).
	ServicesOffered []ServiceOffering `json:"servicesOffered"`
	// CustomerReferences are positive attestations from customers or customer
	// references. Optional per MKT-IAS-WEB.
	CustomerReferences []string `json:"customerReferences,omitempty"`
}
