package dblayer

import (
	"Week02/internal/models"
	"fmt"

	"github.com/pkg/errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DBORM gorm.DB
type DBORM struct {
	*gorm.DB
}

// NewORM new grom.DB
func NewORM(dbpath string) (*DBORM, error) {
	db, err := gorm.Open(sqlite.Open(dbpath), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	return &DBORM{
		DB: db,
	}, err
}

// GetNodes get all nodes
func (db *DBORM) GetNodes() (nodes []models.Node, err error) {
	// db.AutoMigrate(&models.Node{})
	err = db.Find(&nodes).Error
	if err != nil {
		return []models.Node{}, errors.Wrapf(err, fmt.Sprintf("dblayer: get all nodes err"))
	}
	return nodes, nil
}

// AddNode get all nodes
func (db *DBORM) AddNode(node models.Node) (err error) {
	return db.Create(&node).Error
}
