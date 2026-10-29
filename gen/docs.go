//go:build ignore

// Command docs generates MkDocs reference pages from the embedded FedRAMP JSON
// Schema files. Run it from the repository root:
//
//	go run ./gen/docs.go
//
// It reads docs/schemas/*.json and writes one Markdown page per schema into
// docs/reference/, plus docs/reference/index.md. The generated pages are
// checked into the repository so the MkDocs site can be built without a Go
// toolchain.
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

const (
	schemaDir = "docs/schemas"
	outDir    = "docs/reference"
)

// stringOrArray decodes a JSON value that may be a single string or an array
// of strings (the JSON Schema "type" keyword allows both).
type stringOrArray []string

func (s *stringOrArray) UnmarshalJSON(b []byte) error {
	b = []byte(strings.TrimSpace(string(b)))
	if len(b) > 0 && b[0] == '[' {
		return json.Unmarshal(b, (*[]string)(s))
	}
	var str string
	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}
	*s = stringOrArray{str}
	return nil
}

// node is the subset of JSON Schema we render.
type node struct {
	Ref         string           `json:"$ref"`
	SchemaID    string           `json:"$id"`
	Version     string           `json:"$schemaVersion"`
	Type        stringOrArray    `json:"type"`
	Title       string           `json:"title"`
	Description string           `json:"description"`
	Format      string           `json:"format"`
	Enum        []any            `json:"enum"`
	Pattern     string           `json:"pattern"`
	Items       *node            `json:"items"`
	Properties  map[string]*node `json:"properties"`
	Required    []string         `json:"required"`
	Defs        map[string]*node `json:"$defs"`
}

func main() {
	entries, err := os.ReadDir(schemaDir)
	if err != nil {
		fatalf("read %s: %v", schemaDir, err)
	}
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		fatalf("mkdir %s: %v", outDir, err)
	}

	var files []string
	for _, e := range entries {
		if !e.IsDir() && strings.HasSuffix(e.Name(), ".json") {
			files = append(files, e.Name())
		}
	}
	sort.Strings(files)

	type indexRow struct {
		title, page, file, version, desc string
	}
	var rows []indexRow

	for _, file := range files {
		b, err := os.ReadFile(filepath.Join(schemaDir, file))
		if err != nil {
			fatalf("read %s: %v", file, err)
		}
		var root node
		if err := json.Unmarshal(b, &root); err != nil {
			fatalf("parse %s: %v", file, err)
		}
		slug := slugFor(file)
		page := slug + ".md"
		if err := os.WriteFile(filepath.Join(outDir, page), []byte(renderPage(file, slug, &root)), 0o644); err != nil {
			fatalf("write %s: %v", page, err)
		}
		rows = append(rows, indexRow{root.Title, page, file, root.Version, firstLine(root.Description)})
		fmt.Printf("generated %s\n", filepath.Join(outDir, page))
	}

	// Index page.
	var sb strings.Builder
	sb.WriteString("# Schema Reference\n\n")
	sb.WriteString("Reference documentation generated from the FedRAMP JSON Schema files ")
	sb.WriteString("vendored in this repository. The authoritative JSON Schema documents are ")
	sb.WriteString("available for download from each page.\n\n")
	sb.WriteString("| Schema | Version | Description |\n|---|---|---|\n")
	for _, r := range rows {
		fmt.Fprintf(&sb, "| [%s](%s) | `%s` | %s |\n", r.title, r.page, r.version, r.desc)
	}
	if err := os.WriteFile(filepath.Join(outDir, "index.md"), []byte(sb.String()), 0o644); err != nil {
		fatalf("write index.md: %v", err)
	}
	fmt.Printf("generated %s\n", filepath.Join(outDir, "index.md"))

	// Download page (docs/schemas.md), linking the raw JSON Schema files.
	var dl strings.Builder
	dl.WriteString("# Download Schemas\n\n")
	dl.WriteString("The authoritative FedRAMP JSON Schema documents vendored in this repository ")
	dl.WriteString("(schema set `2026-06-24`). These files are served as-is by the documentation site.\n\n")
	dl.WriteString("| Schema | Version | File |\n|---|---|---|\n")
	for _, r := range rows {
		fmt.Fprintf(&dl, "| %s | `%s` | [`%s`](schemas/%s) |\n", r.title, r.version, r.file, r.file)
	}
	if err := os.WriteFile("docs/schemas.md", []byte(dl.String()), 0o644); err != nil {
		fatalf("write docs/schemas.md: %v", err)
	}
	fmt.Println("generated docs/schemas.md")
}

func renderPage(file, slug string, root *node) string {
	var sb strings.Builder
	title := root.Title
	if title == "" {
		title = file
	}
	fmt.Fprintf(&sb, "# %s\n\n", title)
	if root.Description != "" {
		fmt.Fprintf(&sb, "%s\n\n", root.Description)
	}
	sb.WriteString("| Field | Value |\n|---|---|\n")
	fmt.Fprintf(&sb, "| Schema file | [`%s`](../schemas/%s) |\n", file, file)
	if root.Version != "" {
		fmt.Fprintf(&sb, "| `$schemaVersion` | `%s` |\n", root.Version)
	}
	if root.SchemaID != "" {
		fmt.Fprintf(&sb, "| `$id` | `%s` |\n", root.SchemaID)
	}
	sb.WriteString("\n")

	if len(root.Properties) > 0 {
		sb.WriteString("## Properties\n\n")
		writePropertyTable(&sb, slug, root.Properties, root.Required)
	}

	if len(root.Defs) > 0 {
		sb.WriteString("## Definitions\n\n")
		names := sortedKeys(root.Defs)
		for _, name := range names {
			def := root.Defs[name]
			fmt.Fprintf(&sb, "### `%s`\n\n", name)
			if def.Title != "" {
				fmt.Fprintf(&sb, "**%s**\n\n", def.Title)
			}
			if def.Description != "" {
				fmt.Fprintf(&sb, "%s\n\n", def.Description)
			}
			if len(def.Properties) > 0 {
				writePropertyTable(&sb, slug, def.Properties, def.Required)
			} else {
				fmt.Fprintf(&sb, "Type: %s\n\n", typeString(slug, def))
			}
		}
	}

	sb.WriteString("---\n\n")
	sb.WriteString("_This page is generated from the JSON Schema by `go run ./gen/docs.go`. Do not edit by hand._\n")
	return sb.String()
}

func writePropertyTable(sb *strings.Builder, slug string, props map[string]*node, required []string) {
	req := make(map[string]bool, len(required))
	for _, r := range required {
		req[r] = true
	}
	sb.WriteString("| Property | Type | Required | Description |\n|---|---|---|---|\n")
	for _, name := range sortedKeys(props) {
		p := props[name]
		yes := ""
		if req[name] {
			yes = "✓"
		}
		desc := firstLine(p.Description)
		if desc == "" {
			desc = firstLine(p.Title)
		}
		fmt.Fprintf(sb, "| `%s` | %s | %s | %s |\n", name, typeString(slug, p), yes, escapePipes(desc))
	}
	sb.WriteString("\n")
}

func typeString(slug string, n *node) string {
	if n == nil {
		return ""
	}
	if n.Ref != "" {
		return "[" + refName(n.Ref) + "](" + refLink(slug, n.Ref) + ")"
	}
	if len(n.Enum) > 0 {
		vals := make([]string, 0, len(n.Enum))
		for _, v := range n.Enum {
			vals = append(vals, fmt.Sprintf("`%v`", v))
		}
		return "enum: " + strings.Join(vals, ", ")
	}
	base := strings.Join([]string(n.Type), " \\| ")
	if base == "array" && n.Items != nil {
		return "array of " + typeString(slug, n.Items)
	}
	if n.Format != "" {
		base += " (" + n.Format + ")"
	}
	if base == "" {
		return "object"
	}
	return base
}

// refLink resolves a $ref to a Markdown link target. A same-file ref becomes a
// same-page anchor; a cross-file ref links to that schema's reference page and
// anchor. currentSlug is the slug of the page being rendered.
func refLink(currentSlug, ref string) string {
	anchor := "#" + strings.ToLower(refName(ref))
	before, _, ok := strings.Cut(ref, "#")
	if !ok || before == "" {
		return anchor // same-document ref
	}
	// Cross-file ref: before is a URL or path ending in the schema file name.
	fileName := before
	if i := strings.LastIndex(before, "/"); i >= 0 {
		fileName = before[i+1:]
	}
	targetSlug := slugFor(fileName)
	if targetSlug == currentSlug {
		return anchor
	}
	return targetSlug + ".md" + anchor
}

// refName returns a short, readable label for a $ref target.
func refName(ref string) string {
	if _, after, ok := strings.Cut(ref, "#/$defs/"); ok {
		return after
	}
	if i := strings.LastIndex(ref, "/"); i >= 0 {
		return ref[i+1:]
	}
	return ref
}

func slugFor(file string) string {
	s := strings.TrimSuffix(file, ".json")
	s = strings.TrimPrefix(s, "fedramp-")
	s = strings.TrimSuffix(s, "-schema-2026-06-24")
	return s
}

func firstLine(s string) string {
	s = strings.TrimSpace(s)
	if i := strings.IndexByte(s, '\n'); i >= 0 {
		s = s[:i]
	}
	return s
}

func escapePipes(s string) string { return strings.ReplaceAll(s, "|", "\\|") }

func sortedKeys(m map[string]*node) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func fatalf(format string, args ...any) {
	fmt.Fprintf(os.Stderr, "gen/docs: "+format+"\n", args...)
	os.Exit(1)
}
