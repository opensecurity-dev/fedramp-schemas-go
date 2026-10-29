# FedRAMP Common Definitions

Shared type definitions referenced by other FedRAMP schemas.

| Field | Value |
|---|---|
| Schema file | [`fedramp-common-definitions-schema-2026-06-24.json`](../schemas/fedramp-common-definitions-schema-2026-06-24.json) |
| `$schemaVersion` | `0.3.0` |
| `$id` | `https://fedramp.gov/schemas/fedramp-common-definitions-schema-2026-06-24.json` |

## Definitions

### `acceptedVulnerabilityInfo`

**Accepted Vulnerability Info**

Required information for a single accepted vulnerability per VER-RPT-AVI.

| Property | Type | Required | Description |
|---|---|---|---|
| `acceptanceRationale` | string | ✓ | Explanation of why this is an accepted vulnerability. |
| `vulnerabilityDetail` | [vulnerabilityDetail](#vulnerabilitydetail) | ✓ |  |

### `certificationPackageOverviewUri`

**Certification Package Overview**

Full URI of the provider's Certification Package Overview document.

Type: string (uri)

### `detection`

**Detection**

Required information for a single detected vulnerability per VER-RPT-VDT.

| Property | Type | Required | Description |
|---|---|---|---|
| `detectedAt` | string (date-time) | ✓ | When the vulnerability was detected. |
| `detectionSource` | string | ✓ | Tool or activity that identified the vulnerability (e.g., scanner, penetration test, bug bounty). |

### `logoUri`

**Logo**

URL of an organization logo. Must point to a PNG, JPEG, GIF, SVG, WebP, ICO, BMP, or TIFF file.

Type: string (uri)

### `nRating`

**Potential Agency Impact N-Rating**

See VER-EVA-EPA or IEC-CSO-EFI for N1–N5 definitions.

Type: enum: `1`, `2`, `3`, `4`, `5`

### `painReductionEvent`

**PAIN Reduction Event**

Required information for a single PAIN reduction event per VER-RPT-VDT.

| Property | Type | Required | Description |
|---|---|---|---|
| `rating` | [nRating](#nrating) | ✓ | N-rating after the PAIN reduction. |
| `reducedAt` | string (date-time) | ✓ | When the PAIN reduction was completed. |

### `reportPeriodDate`

**Report Period**

From and To dates for the report period.

| Property | Type | Required | Description |
|---|---|---|---|
| `from` | string (date) | ✓ | Start date for the report period. |
| `to` | string (date) | ✓ | End date for the report period. |

### `reportPeriodDateTime`

**Report Period**

From and To dates for the report period.

| Property | Type | Required | Description |
|---|---|---|---|
| `from` | string (date-time) | ✓ | Start date for the report period. |
| `to` | string (date-time) | ✓ | End date for the report period. |

### `vulnerabilityDetail`

**Vulnerability Detail**

Required information for a single detected vulnerability per VER-RPT-VDT.

| Property | Type | Required | Description |
|---|---|---|---|
| `currentRating` | [nRating](#nrating) |  | Currently estimated Potential Agency Impact N-rating. |
| `detection` | [detection](#detection) | ✓ |  |
| `evaluationCompletedAt` | string (date-time) |  | When evaluation was completed. The clock for VER-TFR-MAV runs from this date. |
| `finalDisposition` | enum: `Fully Mitigated`, `Partially Mitigated`, `Remediated`, `False Positive` |  | Omit if still active. |
| `isInternetReachable` | boolean |  | IRV classification per VER-EVA-EIR. |
| `isLikelyExploitable` | boolean |  | LEV classification per VER-EVA-ELX and VER-EVA-AIA. |
| `overdueStatus` | object |  | See VER-TFR-EVU and VER-TFR-MAV for overdue definitions. |
| `painReductionEvents` | array of [painReductionEvent](#painreductionevent) |  | History of completed PAIN-rating reductions for this vulnerability per VER-RPT-VDT. |
| `potentialAgencyImpact` | string |  | Description of the potential impact. |
| `projectedNextReduction` | object |  | Next planned mitigation step. Omit if fully mitigated or none planned. |
| `providerTrackingId` | string | ✓ | Provider's internally assigned identifier for this vulnerability. |
| `supplementaryRiskInformation` | string |  | Per VER-RPT-VDT field 11 and VER-RPT-NID. |
| `vulnerabilityDescription` | string | ✓ | Description of the vulnerability. |

---

_This page is generated from the JSON Schema by `go run ./gen/docs.go`. Do not edit by hand._
