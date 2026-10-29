package fedramp

// CertificationType is the FedRAMP certification type for a Cloud Service
// Offering.
type CertificationType string

// Valid CertificationType values.
const (
	CertificationType20x  CertificationType = "20x"
	CertificationTypeRev5 CertificationType = "Rev5"
)

// ServiceModel is a cloud service model.
type ServiceModel string

// Valid ServiceModel values.
const (
	ServiceModelSaaS ServiceModel = "SaaS"
	ServiceModelPaaS ServiceModel = "PaaS"
	ServiceModelIaaS ServiceModel = "IaaS"
)

// DeploymentModel is the deployment model of a system.
type DeploymentModel string

// Valid DeploymentModel values.
const (
	DeploymentModelPublicCloud              DeploymentModel = "Public Cloud"
	DeploymentModelGovernmentOnlyCloud      DeploymentModel = "Government-Only Cloud"
	DeploymentModelHybridCloud              DeploymentModel = "Hybrid Cloud"
	DeploymentModelCommunityCloud           DeploymentModel = "Community Cloud"
	DeploymentModelGovernmentCommunityCloud DeploymentModel = "Government Community Cloud"
)

// BusinessCategory is a business category supported by a service. See the
// FedRAMP schema for the full enumerated list of permitted values.
type BusinessCategory string

// CertificationPackageOverview is the top-level document for a FedRAMP
// Certification Package Overview (FRC-CSO-PKG): identification, service
// properties, contacts, and supporting documentation for a Cloud Service
// Offering.
//
// Schema: fedramp-certification-package-overview-schema-2026-06-24.json
type CertificationPackageOverview struct {
	// ServiceIdentification holds basic identification details for the Cloud
	// Service Offering.
	ServiceIdentification ServiceIdentification `json:"serviceIdentification"`
	// ServiceProperties holds technical and operational properties of the
	// Cloud Service Offering.
	ServiceProperties ServiceProperties `json:"serviceProperties"`
	// ContactInformation lists key contacts for the Cloud Service Provider. It
	// must include at least one Security contact and one Sales contact
	// (CDS-CSO-PUB).
	ContactInformation []ContactInfo `json:"contactInformation"`
	// Assessor is the Independent Assessment Service that performs verification
	// and validation for the Cloud Service Offering (CDS-CSO-PUB).
	Assessor *PackageAssessor `json:"assessor,omitempty"`
	// CertifiedServices lists services offered by the provider that are
	// included in this certification (CDS-CSO-SVC).
	CertifiedServices []CertifiedService `json:"certifiedServices,omitempty"`
	// ThirdPartyInformationResources are third-party information resources
	// within the Minimum Assessment Scope for the Cloud Service Offering
	// (MAS-CSO-TPR).
	ThirdPartyInformationResources *ThirdPartyInformationResources `json:"thirdPartyInformationResources,omitempty"`
}

// ServiceIdentification holds basic identification details for a Cloud Service
// Offering.
type ServiceIdentification struct {
	// FedRAMPPackageID is the unique identifier assigned to the CSP by FedRAMP.
	// If no FedRAMP ID is available, use the CSP's name and acronym.
	FedRAMPPackageID string `json:"fedRampPackageId"`
	// UEINumber is the Unique Entity Identifier (UEI) assigned to the CSP by
	// SAM.gov. If no UEI is available leave this blank (CDS-CSO-PUB).
	UEINumber string `json:"ueiNumber,omitempty"`
	// ProviderName is the name of the Cloud Service Provider.
	ProviderName string `json:"providerName"`
	// ServiceName is the full name of the Cloud Service Offering.
	ServiceName string `json:"serviceName"`
	// ServiceAcronym is the short acronym or abbreviation for the Cloud Service
	// Offering.
	ServiceAcronym string `json:"serviceAcronym"`
	// ServiceDescription is a detailed description of the Cloud Service
	// Offering. May use Markdown for formatting (CDS-CSO-PUB).
	ServiceDescription string `json:"serviceDescription"`
	// CertificationType is the FedRAMP certification type for the Cloud Service
	// Offering.
	CertificationType CertificationType `json:"certificationType"`
	// Website is the official website for the Cloud Service Offering
	// (CDS-CSO-PUB).
	Website string `json:"website"`
	// Logo is the URL of the CSP logo. It must point to a PNG, JPEG, GIF, SVG,
	// WebP, ICO, BMP, or TIFF file (CDS-CSO-PUB).
	Logo string `json:"logo"`
}

// ServiceProperties holds technical and operational properties of a Cloud
// Service Offering.
type ServiceProperties struct {
	// ServiceType is the cloud service model(s). Select all that apply
	// (CDS-CSO-PUB).
	ServiceType []ServiceModel `json:"serviceType"`
	// DeploymentModel is the deployment model of the system (CDS-CSO-PUB).
	DeploymentModel DeploymentModel `json:"deploymentModel"`
	// BusinessCategory lists the business category(s) supported by the service
	// (CDS-CSO-PUB).
	BusinessCategory []BusinessCategory `json:"businessCategory,omitempty"`
	// TrustCenter is the URL and any access instructions needed for the Cloud
	// Service Provider's FedRAMP Trust Center (CDS-CSO-UTC).
	TrustCenter *Repository `json:"trustCenter,omitempty"`
	// SecureConfigurationGuidance is the URL and any access instructions needed
	// for the Cloud Service Provider's Secure Configuration Guide
	// (SCG-CSO-RSC).
	SecureConfigurationGuidance *Repository `json:"secureConfigurationGuidance,omitempty"`
	// AdditionalRepositories are repository locations for system authorization
	// documentation, policies, procedures, and assessment reports.
	AdditionalRepositories []Repository `json:"additionalRepositories,omitempty"`
	// NextOngoingCertificationReportDate is the date of the next ongoing
	// certification report for the Cloud Service Provider, in YYYY-MM-DD format
	// (CDS-CSO-PUB).
	NextOngoingCertificationReportDate string `json:"nextOngoingCertificationReportDate,omitempty"`
}

// Repository is a repository location for system authorization documentation,
// policies, procedures, assessment reports, etc.
type Repository struct {
	// RepositoryType is the type(s) or categories of this repository. A single
	// repository may satisfy multiple type requirements.
	RepositoryType []string `json:"repositoryType"`
	// URL is the URL of the repository.
	URL string `json:"url"`
	// RepositoryDescription describes the repository contents.
	RepositoryDescription string `json:"repositoryDescription"`
	// AuthenticationRequired indicates whether authentication is required to
	// access the repository.
	AuthenticationRequired bool `json:"authenticationRequired"`
	// AccessRequestInstructions are instructions for requesting access to the
	// repository. Required when AuthenticationRequired is true. May use
	// Markdown for formatting.
	AccessRequestInstructions string `json:"accessRequestInstructions,omitempty"`
}

// ContactInfo is a contact for the Cloud Service Provider.
type ContactInfo struct {
	// ContactType is the role or category of this contact (e.g., Primary,
	// Security, Sales).
	ContactType string `json:"contactType"`
	// ContactName is the name of the contact. It may be a specific person or
	// the role of the contact.
	ContactName string `json:"contactName,omitempty"`
	// ContactEmail is the email address of the contact.
	ContactEmail string `json:"contactEmail,omitempty"`
	// ContactPhone is the phone number in ###-###-#### format (e.g.
	// 202-555-0123).
	ContactPhone string `json:"contactPhone,omitempty"`
}

// PackageAssessor is the Independent Assessment Service that performs
// verification and validation for a Cloud Service Offering.
type PackageAssessor struct {
	// Name is the name of the independent assessor.
	Name string `json:"name"`
	// AssessorID is the unique identifier assigned to the assessor by FedRAMP.
	// It matches the pattern ^\d{6}$.
	AssessorID string `json:"assessorID"`
}

// CertifiedService is a service offered by the provider that is included in a
// certification.
type CertifiedService struct {
	// ServiceName is the name of the service as it appears on the provider's
	// website and in marketing materials.
	ServiceName string `json:"serviceName"`
	// ServiceDescription describes the service. May use Markdown for
	// formatting.
	ServiceDescription string `json:"serviceDescription"`
	// DateAvailable is the date the service became available or was certified
	// (YYYY-MM-DD).
	DateAvailable string `json:"dateAvailable,omitempty"`
}

// ThirdPartyInformationResources are third-party information resources within
// the Minimum Assessment Scope for a Cloud Service Offering (MAS-CSO-TPR).
type ThirdPartyInformationResources struct {
	// Certified are FedRAMP Certified third-party information resources.
	Certified []CertifiedThirdPartyResource `json:"certified,omitempty"`
	// NonCertified are non-FedRAMP Certified third-party information resources.
	NonCertified []NonCertifiedThirdPartyResource `json:"nonCertified,omitempty"`
}

// CertifiedThirdPartyResource is a FedRAMP Certified third-party information
// resource.
type CertifiedThirdPartyResource struct {
	// FedRAMPCertifiedThirdPartyInformationResource is the unique FedRAMP ID of
	// the third-party information resource.
	FedRAMPCertifiedThirdPartyInformationResource string `json:"fedRampCertifiedThirdPartyInformationResource"`
	// UseCase describes how this resource is used in the context of this
	// offering.
	UseCase string `json:"useCase"`
}

// NonCertifiedThirdPartyResource is a non-FedRAMP Certified third-party
// information resource.
type NonCertifiedThirdPartyResource struct {
	// Name is the name of the third-party information resource.
	Name string `json:"name"`
	// Provider is the name of the provider of the third-party information
	// resource.
	Provider string `json:"provider"`
	// Website is the website of the third-party information resource.
	Website string `json:"website,omitempty"`
	// UseCase describes how this resource is used in the context of this
	// offering.
	UseCase string `json:"useCase"`
}
