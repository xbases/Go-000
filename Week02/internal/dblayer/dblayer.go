package dblayer

import (
	"Week02/internal/models"
)

// DBLayer db interface
type DBLayer interface {
	GetNodes() ([]models.Node, error)
	AddNode(models.Node) error
}
