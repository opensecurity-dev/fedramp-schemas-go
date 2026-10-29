# FedRAMP Incident Report (IEC-CSO-IIR / IEC-CSO-OIR / IEC-CSO-FIR)

Unified schema for the three incident report types in the IEC-CSO lifecycle: Initial (IEC-CSO-IIR), Ongoing (IEC-CSO-OIR), and Final (IEC-CSO-FIR).

| Field | Value |
|---|---|
| Schema file | [`fedramp-incident-report-schema-2026-06-24.json`](../schemas/fedramp-incident-report-schema-2026-06-24.json) |
| `$schemaVersion` | `0.2.0` |
| `$id` | `https://fedramp.gov/schemas/fedramp-incident-report-schema-2026-06-24.json` |

## Properties

| Property | Type | Required | Description |
|---|---|---|---|
| `affectedAgencies` | array of string |  | List of customer agencies that may be affected by the incident. |
| `certificationPackageOverviewUri` | [certificationPackageOverviewUri](common-definitions.md#certificationpackageoverviewuri) | ✓ |  |
| `federalIncidentCoordinator` | string |  | Contact information for the provider's federal incident response coordinator. |
| `functionalImpact` | string |  | Describe the actual or potential impact to federal agency customers. |
| `incidentDescription` | string |  | Description of the incident. |
| `indicatorsOfCompromise` | array of string |  | Per IEC-CSO-OIR. |
| `observedActivity` | string |  | Per IEC-CSO-OIR. |
| `potentialImpact` | object |  | Describe the current and historical potential impact to your system. |
| `providerTrackingId` | string | ✓ | Provider's internally assigned identifier for this incident. Must be consistent across IIR, OIR, and FIR reports for the same incident. |
| `recoveryPlan` | string |  | Estimated recovery plan, milestones, and timelines. |
| `relatedCveIds` | array of string |  | Per IEC-CSO-OIR. Include if applicable. |
| `reportType` | enum: `Initial`, `Ongoing`, `Final` | ✓ | Initial = IEC-CSO-IIR; Ongoing = IEC-CSO-OIR; Final = IEC-CSO-FIR. |
| `resolvedAt` | string (date-time) |  | When the incident was resolved and recovery completed. Required for Final reports (IEC-CSO-FIR). |
| `responseAndRecoveryActivities` | string |  | Per IEC-CSO-OIR. |
| `rootCause` | string |  | Per IEC-CSO-OIR. |
| `timeline` | object |  | Timeline of the incident per IEC-CSO-IIR field 4. |

## Definitions

### `milestone`

**Incident Milestone**

Major events in the incident lifecycle.

| Property | Type | Required | Description |
|---|---|---|---|
| `milestoneDescription` | string | ✓ | Description of the milestone. |
| `occurredAt` | string (date-time) | ✓ | When the milestone occurred. |

---

_This page is generated from the JSON Schema by `go run ./gen/docs.go`. Do not edit by hand._
