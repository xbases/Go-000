package handler

import (
	"Week02/internal/dblayer"
	"Week02/internal/models"
)

// Interface handler interface
type Interface interface {
	GetNodes() ([]models.Node, error)
	AddNode(models.Node)
}

// Handler db
type Handler struct {
	db dblayer.DBLayer
}

// NewHandler handler
func NewHandler() (Interface, error) {
	db, err := dblayer.NewORM("data.db")
	if err != nil {
		return nil, err
	}
	return &Handler{
		db: db,
	}, nil
}

// AddNode add node
func (h *Handler) AddNode(node models.Node) {
	h.db.AddNode(node)
}

// GetNodes get all nodes
func (h *Handler) GetNodes() ([]models.Node, error) {
	return h.db.GetNodes()
}
