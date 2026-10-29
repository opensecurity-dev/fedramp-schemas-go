# FedRAMP Ongoing Certification Report (CCM-OCR-AVL)

Quarterly Ongoing Certification Report (OCR) per CCM-OCR-AVL, covering the entire period since the previous report.

| Field | Value |
|---|---|
| Schema file | [`fedramp-ongoing-certification-report-schema-2026-06-24.json`](../schemas/fedramp-ongoing-certification-report-schema-2026-06-24.json) |
| `$schemaVersion` | `0.2.0` |
| `$id` | `https://fedramp.gov/schemas/fedramp-ongoing-certification-report-schema-2026-06-24.json` |

## Properties

| Property | Type | Required | Description |
|---|---|---|---|
| `acceptedVulnerabilities` | string | ✓ | Summary of accepted vulnerabilities. Full records are reported per VER-RPT-AVI. |
| `activeAgencies` | array of string | ✓ | List of all agencies directly using the product. |
| `certificationDataChanges` | array of string | ✓ | High-level summaries of changes to FedRAMP Certification Data since the previous report. |
| `certificationPackageOverviewUri` | [certificationPackageOverviewUri](common-definitions.md#certificationpackageoverviewuri) | ✓ |  |
| `plannedCertificationDataChanges` | object | ✓ | High-level summaries of planned changes to FedRAMP Certification Data for at least the next 3 months. |
| `reportPeriod` | [reportPeriodDate](common-definitions.md#reportperioddate) | ✓ | The 3-month period covered by this report. |
| `reportableIncidents` | object | ✓ | FedRAMP Reportable Incidents during this period. An empty `incidents` array attests that none occurred. |
| `transformativeChanges` | array of string | ✓ | High-level summaries of transformative changes during this period. |
| `updatedRecommendations` | array of string | ✓ | Updated recommendations or best practices for security, configuration, usage, or similar aspects of the cloud service offering. |

## Definitions

### `incidentSummary`

**Incident Summary**

Summary of a FedRAMP Reportable Incident.

| Property | Type | Required | Description |
|---|---|---|---|
| `incidentLessonsLearned` | array of string |  | Lessons learned and changes planned or made as a result of FedRAMP Reportable Incidents. |
| `occurredAt` | string (date-time) |  | Date and time the incident occurred. |
| `resolvedAt` | string (date-time) |  | Date and time the incident was resolved. |
| `summary` | string |  | Brief description of the incident. |

---

_This page is generated from the JSON Schema by `go run ./gen/docs.go`. Do not edit by hand._
