# FedRAMP Certification Package Overview (FRC-CSO-PKG)

Identification, service properties, contacts, and supporting documentation for a Cloud Service Offering per FRC-CSO-PKG.

| Field | Value |
|---|---|
| Schema file | [`fedramp-certification-package-overview-schema-2026-06-24.json`](../schemas/fedramp-certification-package-overview-schema-2026-06-24.json) |
| `$schemaVersion` | `0.1.4` |
| `$id` | `https://fedramp.gov/schemas/fedramp-certification-package-overview-schema-2026-06-24.json` |

## Properties

| Property | Type | Required | Description |
|---|---|---|---|
| `assessor` | object |  | The Independent Assessment Service that performs verification and validation for the Cloud Service Offering (CDS-CSO-PUB). |
| `certifiedServices` | array of object |  | List of services offered by the provider that are included in this certification (CDS-CSO-SVC). |
| `contactInformation` | array of [contactInfo](#contactinfo) | ✓ | Key contacts for the Cloud Service Provider. Must include at least one Security Contact and one Sales Contact (CDS-CSO-PUB). |
| `serviceIdentification` | object | ✓ | Basic identification details for the Cloud Service Offering. |
| `serviceProperties` | object | ✓ | Technical and operational properties of the Cloud Service Offering. |
| `thirdPartyInformationResources` | object |  | Third-party information resources within the Minimum Assessment Scope for the Cloud Service Offering (MAS-CSO-TPR). |

## Definitions

### `contactInfo`

**Contact**

A contact for the Cloud Service Provider.

| Property | Type | Required | Description |
|---|---|---|---|
| `contactEmail` | string (email) |  | Email address of the contact. |
| `contactName` | string |  | Name of the contact. May be a specific person or the role of the contact. |
| `contactPhone` | string |  | Phone number in ###-###-#### format (e.g. 202-555-0123). |
| `contactType` | string | ✓ | The role or category of this contact. |

### `repository`

**Repository**

A repository location for system authorization documentation, policies, procedures, assessment reports, etc.

| Property | Type | Required | Description |
|---|---|---|---|
| `accessRequestInstructions` | string |  | Instructions for requesting access to the repository. May use Markdown for formatting. |
| `authenticationRequired` | boolean | ✓ | Whether authentication is required to access the repository. |
| `repositoryDescription` | string | ✓ | Description of the repository contents. |
| `repositoryType` | array of string | ✓ | The type(s) or categories of this repository. A single repository may satisfy multiple type requirements. |
| `url` | string (uri) | ✓ | URL of the repository. |

---

_This page is generated from the JSON Schema by `go run ./gen/docs.go`. Do not edit by hand._
