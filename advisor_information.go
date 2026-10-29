package fedramp

// AdvisorInformation is the top-level document for FedRAMP Advisor Information
// (MKT-CAS-WEB): the public information an Advisory Service must supply in
// machine-readable format on their website.
//
// Schema: fedramp-advisor-information-schema-2026-06-24.json
type AdvisorInformation struct {
	// AdvisorName is the name of the advisory organization as it should be
	// displayed on the FedRAMP Marketplace (MKT-CAS-WEB).
	AdvisorName string `json:"advisorName"`
	// Logo is the URL of the advisor logo displayed on the FedRAMP Marketplace
	// (MKT-CAS-WEB). It must point to a PNG, JPEG, GIF, SVG, WebP, ICO, BMP, or
	// TIFF file.
	Logo string `json:"logo"`
	// A2LAID is the identifier assigned to the organization by the American
	// Association for Laboratory Accreditation. Optional for advisors; CR26
	// does not require advisory services to be A2LA accredited.
	A2LAID string `json:"a2laId,omitempty"`
	// ServiceDescription is a general description of the consulting or advisory
	// service (MKT-CAS-WEB).
	ServiceDescription string `json:"serviceDescription"`
	// ContactInformation is contact information for the advisory organization
	// (MKT-CAS-WEB).
	ContactInformation []string `json:"contactInformation"`
	// ServicesOffered lists the types of consulting or advisory services
	// offered (MKT-CAS-WEB).
	ServicesOffered []ServiceOffering `json:"servicesOffered"`
	// CustomerReferences are positive attestations from customers or customer
	// references. Optional per MKT-CAS-WEB.
	CustomerReferences []string `json:"customerReferences,omitempty"`
}

// ServiceOffering is a consulting, advisory, or assessment service offered by
// an advisor or assessor.
type ServiceOffering struct {
	// ServiceName is the name of the service type offered.
	ServiceName string `json:"serviceName"`
	// ServiceDescription describes this service type and its scope.
	ServiceDescription string `json:"serviceDescription,omitempty"`
}
