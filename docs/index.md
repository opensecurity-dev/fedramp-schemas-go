# FedRAMP Schemas (Go)

Go types and vendored JSON Schemas for the
[FedRAMP machine-readable schema set](https://www.fedramp.gov/schemas/).

This project provides:

- **Go structs** mapping one-to-one to each FedRAMP JSON Schema document
  (draft 2020-12), for the `2026-06-24` schema set.
- **Embedded JSON Schemas** so applications can validate documents against the
  authoritative schema at runtime — no network round-trip required.
- **Reference documentation** generated directly from the schemas.

## Install

```bash
go get github.com/opensecurity-dev/fedramp-schemas-go
```

```go
import fedramp "github.com/opensecurity-dev/fedramp-schemas-go"
```

## Schema documents

| Document | Go type | FedRAMP requirement |
|---|---|---|
| [Certification Package Overview](reference/certification-package-overview.md) | `CertificationPackageOverview` | FRC-CSO-PKG |
| [Security Decision Record](reference/security-decision-record.md) | `SecurityDecisionRecord` | SDR-CSO-FRR |
| [Significant Change Notification](reference/significant-change-notifications.md) | `SignificantChangeNotification` | SCN-CSO-INF |
| [Incident Report](reference/incident-report.md) | `IncidentReport` | IEC-CSO-IIR / OIR / FIR |
| [Ongoing Certification Report](reference/ongoing-certification-report.md) | `OngoingCertificationReport` | CCM-OCR-AVL |
| [Vulnerability Detail Report](reference/vulnerability-detail-report.md) | `VulnerabilityDetailReport` | VER-RPT-VDT |
| [Accepted Vulnerability Info](reference/accepted-vulnerability-info.md) | `AcceptedVulnerabilityInfoReport` | VER-RPT-AVI |
| [Historical VER Activity](reference/historical-ver-activity.md) | `HistoricalVERActivity` | VER-TFR-MRH |
| [Advisor Information](reference/advisor-information.md) | `AdvisorInformation` | MKT-CAS-WEB |
| [Assessor Information](reference/assessor-information.md) | `AssessorInformation` | MKT-IAS-WEB |
| [Common Definitions](reference/common-definitions.md) | shared types | — |

## Next steps

- [Go Usage](usage.md) — marshalling, unmarshalling, and accessing embedded schemas.
- [Schema Reference](reference/index.md) — field-level documentation.
- [Download Schemas](schemas.md) — the raw JSON Schema files.
