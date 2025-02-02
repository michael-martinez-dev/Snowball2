package config

import (
	"sync"
)

// Config holds the configuration for the application
type Config struct {
	DBType    string
	DBPath    string
	NotesPath string
}

// ConfigService is a service that holds the configuration for the application
type ConfigService struct {
	mu     sync.RWMutex
	config *Config
}

// NewConfigService creates a new ConfigService with the given initial configuration
func NewConfigService(initial *Config) *ConfigService {
	return &ConfigService{
		config: initial,
	}
}

// GetConfig returns a copy of the current configuration
func (cs *ConfigService) GetConfig() Config {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	return *cs.config
}

// UpdateConfig updates the configuration
func (cs *ConfigService) UpdateConfig(newConfig Config) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.config = &newConfig
}
