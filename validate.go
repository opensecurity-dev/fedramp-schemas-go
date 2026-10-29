package fedramp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/santhosh-tekuri/jsonschema/v6"
)

// Validator validates documents against the embedded FedRAMP JSON Schemas.
//
// A Validator registers every embedded schema by its "$id", so cross-file
// "$ref"s (which use fully-qualified https://fedramp.gov/schemas/... URIs)
// resolve against the embedded copies rather than being fetched over the
// network. Compiled schemas are cached, so a Validator is safe and efficient
// for concurrent reuse.
type Validator struct {
	compiler *jsonschema.Compiler
	// idByFile maps a schema base file name to its "$id".
	idByFile map[string]string

	mu       sync.Mutex
	compiled map[string]*jsonschema.Schema
}

// NewValidator builds a Validator with every embedded FedRAMP schema
// registered as a resource.
func NewValidator() (*Validator, error) {
	c := jsonschema.NewCompiler()
	idByFile := make(map[string]string, len(SchemaFiles))
	for _, name := range SchemaFiles {
		b, err := SchemaBytes(name)
		if err != nil {
			return nil, err
		}
		doc, err := jsonschema.UnmarshalJSON(bytes.NewReader(b))
		if err != nil {
			return nil, fmt.Errorf("fedramp: decoding schema %q: %w", name, err)
		}
		var meta schemaMeta
		if err := json.Unmarshal(b, &meta); err != nil {
			return nil, fmt.Errorf("fedramp: reading $id of schema %q: %w", name, err)
		}
		if meta.ID == "" {
			return nil, fmt.Errorf("fedramp: schema %q has no $id", name)
		}
		if err := c.AddResource(meta.ID, doc); err != nil {
			return nil, fmt.Errorf("fedramp: registering schema %q: %w", name, err)
		}
		idByFile[name] = meta.ID
	}
	return &Validator{
		compiler: c,
		idByFile: idByFile,
		compiled: make(map[string]*jsonschema.Schema, len(SchemaFiles)),
	}, nil
}

// schema returns the compiled schema for the given base file name (see the
// SchemaFile* constants), compiling and caching it on first use.
func (v *Validator) schema(schemaFile string) (*jsonschema.Schema, error) {
	v.mu.Lock()
	defer v.mu.Unlock()
	if s, ok := v.compiled[schemaFile]; ok {
		return s, nil
	}
	id, ok := v.idByFile[schemaFile]
	if !ok {
		return nil, fmt.Errorf("fedramp: unknown schema %q", schemaFile)
	}
	s, err := v.compiler.Compile(id)
	if err != nil {
		return nil, fmt.Errorf("fedramp: compiling schema %q: %w", schemaFile, err)
	}
	v.compiled[schemaFile] = s
	return s, nil
}

// ValidateBytes validates raw JSON document bytes against the embedded schema
// identified by schemaFile (see the SchemaFile* constants). It returns nil if
// the document is valid; otherwise the returned error is a
// *jsonschema.ValidationError describing every violation.
func (v *Validator) ValidateBytes(schemaFile string, data []byte) error {
	s, err := v.schema(schemaFile)
	if err != nil {
		return err
	}
	inst, err := jsonschema.UnmarshalJSON(bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("fedramp: decoding document: %w", err)
	}
	return s.Validate(inst)
}

// ValidateValue marshals v to JSON and validates it against the embedded schema
// identified by schemaFile (see the SchemaFile* constants). It is a convenience
// wrapper over ValidateBytes for validating a populated document struct.
func (v *Validator) ValidateValue(schemaFile string, doc any) error {
	data, err := json.Marshal(doc)
	if err != nil {
		return fmt.Errorf("fedramp: marshaling document: %w", err)
	}
	return v.ValidateBytes(schemaFile, data)
}

var (
	defaultValidator     *Validator
	defaultValidatorErr  error
	defaultValidatorOnce sync.Once
)

func sharedValidator() (*Validator, error) {
	defaultValidatorOnce.Do(func() {
		defaultValidator, defaultValidatorErr = NewValidator()
	})
	return defaultValidator, defaultValidatorErr
}

// ValidateBytes validates raw JSON document bytes against the embedded schema
// identified by schemaFile (see the SchemaFile* constants), using a shared
// process-wide Validator. It returns nil if the document is valid; otherwise
// the returned error is a *jsonschema.ValidationError.
func ValidateBytes(schemaFile string, data []byte) error {
	v, err := sharedValidator()
	if err != nil {
		return err
	}
	return v.ValidateBytes(schemaFile, data)
}

// ValidateValue marshals doc to JSON and validates it against the embedded
// schema identified by schemaFile (see the SchemaFile* constants), using a
// shared process-wide Validator.
func ValidateValue(schemaFile string, doc any) error {
	v, err := sharedValidator()
	if err != nil {
		return err
	}
	return v.ValidateValue(schemaFile, doc)
}
