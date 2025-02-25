package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type File struct {
	Path   string
	Auth   AuthStorage
	Source Source `yaml:"source"`
	Sink   Sink   `yaml:"sink"`
	// TODO: filters for source events
	Transformations   []Transformer `yaml:"transformations,omitempty"`
	Sync              Sync          `yaml:"sync"`
	UpdateConcurrency int           `yaml:"updateConcurrency,omitempty"`
}

func NewFromFile(path string) (*File, error) {
	yamlFile, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("cannot read config file: %w", err)
	}

	config := File{
		Path: path,
		Auth: AuthStorage{
			StorageMode: "yaml",
			Config: CustomMap{
				"path": "./auth-storage.yaml",
			},
		},
	}

	if err := yaml.Unmarshal(yamlFile, &config); err != nil {
		return nil, fmt.Errorf("cannot unmarshal config file: %w", err)
	}

	return &config, nil
}

type AuthStorage struct {
	StorageMode string `yaml:"storage_mode"`
	// Any kind of parameter which can be passed to the StorageMode
	Config CustomMap `yaml:"config"`
}

type Source struct {
	Adapter Adapter `yaml:"adapter"`
}

type Sink struct {
	Adapter Adapter `yaml:"adapter"`
}

type Adapter struct {
	// Type of the adapter (client) to use for the source calendar
	Type string `yaml:"type"`
	// ID of the calendar in which the adapter will work.
	Calendar string `yaml:"calendar"`
	// CustomMap is an adapter-specific map to configure it.
	Config CustomMap `yaml:"config"`
	// OAuth values for the adapter
	OAuth OAuth `yaml:"oAuth"`
}

type OAuth struct {
	ClientID  string `yaml:"clientId,omitempty"`
	ClientKey string `yaml:"clientKey,omitempty"`
	TenantID  string `yaml:"tenantId,omitempty"`
}

// CustomMap is meant to provide custom parameters to different adapters/transformers.
type CustomMap map[string]interface{}

// Transformer configures the name
type Transformer struct {
	// Name of the transformer to run
	Name string `yaml:"name"`
	// Any kind of parameter which can be passed to a transformer.
	Config CustomMap `yaml:"config"`
}

// Sync configuration
type Sync struct {
	StartTime SyncTime `yaml:"start"`
	EndTime   SyncTime `yaml:"end"`
}

type SyncTime struct {
	Identifier string `yaml:"identifier"`
	Offset     int    `yaml:"offset,omitempty"`
}
