# fedramp-schemas-go

[![Go Reference](https://pkg.go.dev/badge/github.com/opensecurity-dev/fedramp-schemas-go.svg)](https://pkg.go.dev/github.com/opensecurity-dev/fedramp-schemas-go)
[![Docs](https://img.shields.io/badge/docs-mkdocs-blue)](https://opensecurity-dev.github.io/fedramp-schemas-go/)

Go types and vendored JSON Schemas for the
[FedRAMP machine-readable schema set](https://www.fedramp.gov/schemas/).

This library provides:

- **Go structs** mapping one-to-one to each FedRAMP JSON Schema document
  (JSON Schema draft 2020-12), for the `2026-06-24` schema set.
- **Embedded JSON Schemas**, so applications can validate documents against the
  authoritative schema at runtime without a network round-trip.
- A generated **MkDocs documentation site**
  ([opensecurity-dev.github.io/fedramp-schemas-go](https://opensecurity-dev.github.io/fedramp-schemas-go/)).

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
| Certification Package Overview | `CertificationPackageOverview` | FRC-CSO-PKG |
| Security Decision Record | `SecurityDecisionRecord` | SDR-CSO-FRR |
| Significant Change Notification | `SignificantChangeNotification` | SCN-CSO-INF |
| Incident Report | `IncidentReport` | IEC-CSO-IIR / OIR / FIR |
| Ongoing Certification Report | `OngoingCertificationReport` | CCM-OCR-AVL |
| Vulnerability Detail Report | `VulnerabilityDetailReport` | VER-RPT-VDT |
| Accepted Vulnerability Info | `AcceptedVulnerabilityInfoReport` | VER-RPT-AVI |
| Historical VER Activity | `HistoricalVERActivity` | VER-TFR-MRH |
| Advisor Information | `AdvisorInformation` | MKT-CAS-WEB |
| Assessor Information | `AssessorInformation` | MKT-IAS-WEB |
| Common Definitions | shared types (`VulnerabilityDetail`, `NRating`, …) | — |

## Usage

Unmarshal a document:

```go
var report fedramp.IncidentReport
if err := json.Unmarshal(data, &report); err != nil {
    return err
}
fmt.Println(report.ReportType, report.ProviderTrackingID)
```

Access the embedded authoritative JSON Schemas:

```go
raw, err := fedramp.SchemaBytes(fedramp.SchemaFileIncidentReport)

schemas, err := fedramp.Schemas()
for _, s := range schemas {
    fmt.Printf("%s: %s (%s)\n", s.Name, s.Title, s.Version)
}
```

Validate a document against its embedded schema — this enforces conditional
rules the structs can't express (e.g. a Final incident report requires
`resolvedAt`, and contacts must include both a Security and a Sales entry):

```go
// Raw JSON bytes:
err := fedramp.ValidateBytes(fedramp.SchemaFileIncidentReport, data)

// Or a populated struct:
err = fedramp.ValidateValue(fedramp.SchemaFileIncidentReport, report)
```

Check enumerated values:

```go
fedramp.CertificationType20x.IsValid()      // true
fedramp.CertificationType("HighBaseline").IsValid() // false
fedramp.DeploymentModel("").Values()        // all permitted values
```

See the [Go Usage guide](https://opensecurity-dev.github.io/fedramp-schemas-go/usage/)
for more, including enums, optional fields, and date handling.

## Conventions

- **Dates/date-times** are `string` (RFC 3339 for `date-time`, `YYYY-MM-DD` for
  `date`) to losslessly round-trip both JSON Schema formats. Field doc comments
  note which applies.
- **Optional booleans and enumerated integers** are pointers, so an absent
  value is distinguishable from a zero value.
- **Enumerations** are named string/integer types with exported constants.

## Repository layout

```
.                              # package fedramp (root)
├── common.go                  # shared Common Definitions types
├── *.go                       # one file per document schema
├── schemas.go                 # //go:embed of the JSON Schemas + accessors
├── validate.go                # runtime validation against embedded schemas
├── enums.go                   # IsValid()/Values() for enumerated types
├── gen/docs.go                # //go:build ignore doc-page generator
└── docs/
    ├── schemas/*.json         # vendored FedRAMP JSON Schemas (source of truth)
    ├── reference/*.md         # generated schema reference (do not edit)
    ├── index.md, usage.md     # hand-written pages
    └── requirements.txt       # MkDocs build dependencies
```

## Development

```bash
go test ./...          # run tests (includes schema/struct parity checks)
go vet ./...
golangci-lint run
go run ./gen/docs.go   # regenerate docs/reference/*.md and docs/schemas.md
mkdocs serve           # preview the docs site locally
```

The schema files under `docs/schemas/` are the source of truth. When
re-vendoring a new FedRAMP schema set, update the Go types, the version table
in `schemas_test.go`, and regenerate the reference pages.

## License and Source

The Go code is licensed under the terms in [LICENSE](LICENSE). The JSON Schemas
under `docs/schemas/` are the FedRAMP machine-readable schema set published by
the U.S. federal government, and are treated as public domain; see
[NOTICE](NOTICE) for their source.
