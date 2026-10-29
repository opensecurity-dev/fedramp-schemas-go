// Package fedramp provides Go types for the FedRAMP machine-readable schemas
// published at https://www.fedramp.gov/schemas/.
//
// The types map one-to-one to the FedRAMP JSON Schema documents (draft
// 2020-12) for the 2026-06-24 schema set. Each document schema has a
// corresponding top-level Go struct, and shared definitions from the
// FedRAMP Common Definitions schema live in this package's common types
// (see common.go).
//
// The original JSON Schema files are embedded (see schemas.go) so that
// applications can validate documents against the authoritative schema at
// runtime without a network round-trip.
//
// Date and date-time fields are represented as strings to losslessly
// round-trip both the JSON Schema "date" (YYYY-MM-DD) and "date-time"
// (RFC 3339) formats. Field-level doc comments note which format applies.
//
// Optional scalar fields that carry meaning in their zero value (booleans
// and enumerated integers) are represented as pointers so that "absent" is
// distinguishable from "false" or "zero".
package fedramp
