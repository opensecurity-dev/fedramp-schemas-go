# FedRAMP Significant Change Notification (SCN-CSO-INF)

Required information for a Significant Change Notification per SCN-CSO-INF. Note per the rule: structure of the information may vary depending on how the provider tracks this internally.

| Field | Value |
|---|---|
| Schema file | [`fedramp-significant-change-notifications-schema-2026-06-24.json`](../schemas/fedramp-significant-change-notifications-schema-2026-06-24.json) |
| `$schemaVersion` | `0.1.2` |
| `$id` | `https://fedramp.gov/schemas/fedramp-significant-change-notifications-schema-2026-06-24.json` |

## Properties

| Property | Type | Required | Description |
|---|---|---|---|
| `assessorName` | string |  | Name of the assessor involved in this change, if applicable. |
| `certificationPackageOverviewUri` | [certificationPackageOverviewUri](common-definitions.md#certificationpackageoverviewuri) | ✓ |  |
| `changeDescription` | string | ✓ | Short description of the change. |
| `changeType` | enum: `Adaptive`, `Transformative` | ✓ | Significant change type per SCN-CSO-EVA. Routine recurring changes (SCN-RTR) do not require a Significant Change Notification. |
| `changeTypeExplanation` | string |  | Explanation of why this change was categorized as the selected type. |
| `customerImpact` | string |  | Summary of customer impact, including changes to services and customer configuration responsibilities. |
| `impactAnalysis` | string |  | Copy or summary of the business or security impact analysis. May be the full text or a URI referencing the document. |
| `impactedControls` | array of string |  | KSI or control identifiers that will be verified, assessed, or validated as part of this change. |
| `planAndTimeline` | object |  | Plan and timeline for the change, including verification, assessment, and/or validation of impacted Key Security Indicators (KSIs) or Rev5 Controls. |
| `reason` | string |  | Reason for the change. |
| `relatedVulnerability` | string |  | Identifier or description of any related vulnerability driving this change, if applicable. |

---

_This page is generated from the JSON Schema by `go run ./gen/docs.go`. Do not edit by hand._
