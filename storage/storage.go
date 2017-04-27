package storage

import (
	"github.com/AKovalevich/dapper/service"
)

// Storage interface
type Storable interface {
	// Clone the storage session.
	Close()

	SaveService(service service.Service) (*service.Service)

	LoadService()

	CreateSchemas() (error)
}