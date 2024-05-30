package supplychain

import (
	"supplychain/graph/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Store interface {
	AddInventoryItem(string, string, int, string) (primitive.ObjectID, error)
	UpdateInventoryItem(primitive.ObjectID, *string, *string, *int, *string) (primitive.ObjectID, error)
	DeleteInventoryItem(primitive.ObjectID) bool
	UpdateShipmentStatus(primitive.ObjectID, string) (primitive.ObjectID, error)
	AddSupplier(string, string, string, string) (primitive.ObjectID, error)
	UpdateSupplier(primitive.ObjectID, *string, *string, *string, *string) (primitive.ObjectID, error)
	DeleteSupplier(primitive.ObjectID) bool

	GetInventoryItems(int64, int64) ([]*model.InventoryItem, error)
	GetShipment(primitive.ObjectID) (*model.Shipment, error)
	GetSuppliers() ([]*model.Supplier, error)
	GetSingleSupplier(primitive.ObjectID) (*model.Supplier, error)
	GetSingleInventory(primitive.ObjectID) (*model.InventoryItem, error)
}
