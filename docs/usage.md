# Go Usage

## Import

```go
import fedramp "github.com/opensecurity-dev/fedramp-schemas-go"
```

## Unmarshalling a document

```go
data, err := os.ReadFile("incident-report.json")
if err != nil {
    return err
}

var report fedramp.IncidentReport
if err := json.Unmarshal(data, &report); err != nil {
    return err
}

fmt.Println(report.ReportType, report.ProviderTrackingID)
```

## Marshalling a document

```go
report := fedramp.IncidentReport{
    CertificationPackageOverviewURI: "https://example.com/cpo.json",
    ReportType:                      fedramp.IncidentReportTypeInitial,
    ProviderTrackingID:              "INC-2026-0001",
}

out, err := json.MarshalIndent(report, "", "  ")
if err != nil {
    return err
}
fmt.Println(string(out))
```

## Enumerated values

Enumerations are represented as named string (or integer) types with exported
constants:

```go
_ = fedramp.CertificationType20x           // "20x"
_ = fedramp.ImplementationStatusImplemented // "Implemented"
_ = fedramp.NRating3                        // 3
```

## Optional fields

Optional booleans and enumerated integers are pointers so that an absent value
is distinguishable from a zero value:

```go
overdue := true
vd := fedramp.VulnerabilityDetail{
    ProviderTrackingID:       "V-1",
    Detection:                fedramp.Detection{DetectedAt: "2026-01-15T00:00:00Z", DetectionSource: "scanner"},
    VulnerabilityDescription: "Example",
    IsInternetReachable:      &overdue,
    CurrentRating:            new(fedramp.NRating3),
}
```

## Dates and date-times

Date and date-time fields are represented as `string` to losslessly round-trip
both the JSON Schema `date` (`YYYY-MM-DD`) and `date-time` (RFC 3339) formats.
Each field's doc comment states which format applies.

## Accessing the embedded schemas

The authoritative JSON Schema files are embedded, so you can validate documents
at runtime without fetching them over the network:

```go
// Raw bytes of a single schema.
raw, err := fedramp.SchemaBytes(fedramp.SchemaFileIncidentReport)

// Metadata (title, $schemaVersion) for every embedded schema.
schemas, err := fedramp.Schemas()
for _, s := range schemas {
    fmt.Printf("%s: %s (%s)\n", s.Name, s.Title, s.Version)
}

// A read-only fs.FS rooted at the embedded schema directory.
fsys, err := fedramp.SchemaFS()
```

## Validating documents

The package validates documents against the embedded schemas, so you don't need
to wire up your own JSON Schema validator. Cross-file `$ref`s resolve against
the embedded copies — no network access is required.

```go
// Validate raw JSON bytes.
if err := fedramp.ValidateBytes(fedramp.SchemaFileIncidentReport, data); err != nil {
    // err is a *jsonschema.ValidationError describing every violation.
    return err
}

// Or validate a populated struct.
if err := fedramp.ValidateValue(fedramp.SchemaFileIncidentReport, report); err != nil {
    return err
}
```

Validation enforces the conditional rules that the Go structs alone cannot
express, including:

- A Final incident report (`reportType: "Final"`) requires `resolvedAt`.
- `overdueStatus.isOverdue == true` requires an `explanation`.
- A repository with `authenticationRequired == true` requires
  `accessRequestInstructions`.
- `contactInformation` must contain both a Security and a Sales contact.

For repeated validation, construct a reusable `Validator` once (it caches
compiled schemas and is safe for concurrent use):

```go
v, err := fedramp.NewValidator()
if err != nil {
    return err
}
err = v.ValidateBytes(fedramp.SchemaFileSecurityDecisionRecord, data)
```

## Enumerated value helpers

Each enumerated type exposes `Values()` and `IsValid()`:

```go
fedramp.CertificationType20x.IsValid()               // true
fedramp.CertificationType("HighBaseline").IsValid()  // false

for _, m := range fedramp.DeploymentModel("").Values() {
    fmt.Println(m)
}
```
