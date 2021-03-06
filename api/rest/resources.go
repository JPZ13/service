package rest

import (
	"log"

	"github.com/JPZ13/service/core"
)

type baseResource struct {
	logger  log.Logger
	service core.Service
}

// Resources holds api resources
type Resources struct {
	authorsResource
	postsResource
}

// Config holds api resource config
type Config struct {
	Core   core.Service
	Logger log.Logger
}

// NewResources inits api resources
func NewResources(config *Config) Resources {
	br := baseResource{
		logger:  config.Logger,
		service: config.Core,
	}

	return Resources{
		authorsResource{br},
		postsResource{br},
	}
}
