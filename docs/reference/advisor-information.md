# FedRAMP Advisor Information (MKT-CAS-WEB)

Public information an Advisory Service must supply in machine-readable format on their website per MKT-CAS-WEB.

| Field | Value |
|---|---|
| Schema file | [`fedramp-advisor-information-schema-2026-06-24.json`](../schemas/fedramp-advisor-information-schema-2026-06-24.json) |
| `$schemaVersion` | `1.0.1` |
| `$id` | `https://fedramp.gov/schemas/fedramp-advisor-information-schema-2026-06-24.json` |

## Properties

| Property | Type | Required | Description |
|---|---|---|---|
| `a2laId` | string |  | Identifier assigned to the organization by the American Association for Laboratory Accreditation. Optional for advisors; CR26 does not require advisory services to be A2LA accredited. |
| `advisorName` | string | ✓ | Name of the advisory organization as it should be displayed on the FedRAMP Marketplace (MKT-CAS-WEB). |
| `contactInformation` | array of string | ✓ | Contact information for the advisory organization (MKT-CAS-WEB). |
| `customerReferences` | array of string |  | Positive attestations from customers or customer references. Optional per MKT-CAS-WEB. |
| `logo` | [logoUri](common-definitions.md#logouri) | ✓ | URL of the advisor logo displayed on the FedRAMP Marketplace (MKT-CAS-WEB). |
| `serviceDescription` | string | ✓ | General description of the consulting or advisory service (MKT-CAS-WEB). |
| `servicesOffered` | array of [serviceOffering](#serviceoffering) | ✓ | Types of consulting or advisory services offered (MKT-CAS-WEB). |

## Definitions

### `serviceOffering`

**Service Offering**

Consulting or advisory service offered

| Property | Type | Required | Description |
|---|---|---|---|
| `serviceDescription` | string |  | Description of this service type and its scope. |
| `serviceName` | string | ✓ | Name of the consulting or advisory service type offered. |

---

_This page is generated from the JSON Schema by `go run ./gen/docs.go`. Do not edit by hand._
