# FedRAMP Assessor Information (MKT-IAS-WEB)

Public information an Independent Assessment Service must supply in machine-readable format on their website per MKT-IAS-WEB.

| Field | Value |
|---|---|
| Schema file | [`fedramp-assessor-information-schema-2026-06-24.json`](../schemas/fedramp-assessor-information-schema-2026-06-24.json) |
| `$schemaVersion` | `1.0.1` |
| `$id` | `https://fedramp.gov/schemas/fedramp-assessor-information-schema-2026-06-24.json` |

## Properties

| Property | Type | Required | Description |
|---|---|---|---|
| `a2laAccreditationDate` | string (date) | ✓ | Date the American Association for Laboratory Accreditation accredited this assessor, in YYYY-MM-DD format (REC-IAS-ACC). |
| `a2laId` | string | ✓ | Identifier assigned to the assessor by the American Association for Laboratory Accreditation. Required for FedRAMP Recognition per REC-IAS-ACC. |
| `assessorName` | string | ✓ | Name of the independent assessment organization as it should be displayed on the FedRAMP Marketplace (MKT-IAS-WEB). |
| `contactInformation` | array of string | ✓ | Contact information for the assessor organization (MKT-IAS-WEB). |
| `customerReferences` | array of string |  | Positive attestations from customers or customer references. Optional per MKT-IAS-WEB. |
| `logo` | [logoUri](common-definitions.md#logouri) | ✓ | URL of the assessor logo displayed on the FedRAMP Marketplace (MKT-IAS-WEB). |
| `serviceDescription` | string | ✓ | General description of the independent assessment service (MKT-IAS-WEB). |
| `servicesOffered` | array of [serviceOffering](#serviceoffering) | ✓ | Types of independent assessment services the assessor offers (MKT-IAS-WEB). |

## Definitions

### `serviceOffering`

**Service Offering**

A specific type of assessment service offered by the assessor.

| Property | Type | Required | Description |
|---|---|---|---|
| `serviceDescription` | string |  | Description of this service type and its scope. |
| `serviceName` | string | ✓ | Name of the assessment service type offered. |

---

_This page is generated from the JSON Schema by `go run ./gen/docs.go`. Do not edit by hand._
