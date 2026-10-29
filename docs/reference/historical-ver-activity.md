# FedRAMP Historical Vulnerability Evaluation and Reporting Activity (VER-TFR-MRH)

Machine-readable historical Vulnerability Evaluation and Reporting activity for automated retrieval per VER-TFR-MRH.

| Field | Value |
|---|---|
| Schema file | [`fedramp-historical-ver-activity-schema-2026-06-24.json`](../schemas/fedramp-historical-ver-activity-schema-2026-06-24.json) |
| `$schemaVersion` | `0.1.1` |
| `$id` | `https://fedramp.gov/schemas/fedramp-historical-ver-activity-schema-2026-06-24.json` |

## Properties

| Property | Type | Required | Description |
|---|---|---|---|
| `acceptedVulnerabilities` | array of [acceptedVulnerabilityInfo](common-definitions.md#acceptedvulnerabilityinfo) | ✓ | All accepted vulnerabilities. Each record follows the fields required by VER-RPT-AVI. |
| `activeVulnerabilities` | array of [vulnerabilityDetail](common-definitions.md#vulnerabilitydetail) | ✓ | All non-accepted vulnerabilities. Each record follows the fields required by VER-RPT-VDT. |
| `certificationPackageOverviewUri` | [certificationPackageOverviewUri](common-definitions.md#certificationpackageoverviewuri) | ✓ |  |
| `generatedAt` | string (date-time) | ✓ | When this snapshot was generated. |

---

_This page is generated from the JSON Schema by `go run ./gen/docs.go`. Do not edit by hand._
