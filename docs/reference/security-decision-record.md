# FedRAMP Security Decision Record (SDR-CSO-FRR)

JSON Schema for Cloud Service Provider (CSP) system submission for FedRAMP certification per SDR-CSO-FRR.

| Field | Value |
|---|---|
| Schema file | [`fedramp-security-decision-record-schema-2026-06-24.json`](../schemas/fedramp-security-decision-record-schema-2026-06-24.json) |
| `$schemaVersion` | `1.1.1` |
| `$id` | `https://fedramp.gov/schemas/fedramp-security-decision-record-schema-2026-06-24.json` |

## Properties

| Property | Type | Required | Description |
|---|---|---|---|
| `certificationPackageOverviewUri` | [certificationPackageOverviewUri](common-definitions.md#certificationpackageoverviewuri) | ✓ |  |
| `fedRampRequirements` | array of object | ✓ | Describe how any applicable FedRAMP required processes are met. |
| `keySecurityIndicators` | array of object |  | Describe how the system meets the FedRAMP Key Security Indicators (KSI). |
| `metadata` | object |  | Basic metadata about this Security Decision Record, per SDR-CSO-MTD. |
| `portsAndProtocols` | array of object |  | Ports and protocols used by the Cloud Service Offering |
| `securityControls` | array of object |  | Describe the security controls for the system. |

## Definitions

### `assessmentStatement`

**Assessment**

Description of how the control/requirement/ksi is assessed by an independent validator. May use Markdown for formatting.

Type: string

### `evidence`

**Evidence**

Results of the security test

| Property | Type | Required | Description |
|---|---|---|---|
| `evidenceDescription` | string |  | Detailed description of the evidence. |
| `evidenceLocation` | string (uri) |  | URI to the evidence document |
| `evidenceText` | string |  | Evidence included as plain text such as the output of a command or log. |
| `evidenceType` | enum: `Log`, `Report`, `Screenshot`, `Configuration`, `Policy`, `Procedure`, `Audit Record` |  | Type of evidence provided |
| `lastUpdated` | string (date) |  | Date when the evidence was last updated |

### `implementationStatement`

**Implementation**

Description of how the control/requirement/ksi is implemented. May use Markdown for formatting.

Type: string

### `validationStatement`

**Validation**

Description of the how the control/requirement/ksi is validated by the CSP internally. May use Markdown for formatting.

Type: string

---

_This page is generated from the JSON Schema by `go run ./gen/docs.go`. Do not edit by hand._
