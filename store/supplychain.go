package store

import (
	"context"
	"errors"
	"supplychain/graph/model"
	"supplychain/pkg/logs"

	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SupplyChainStore struct {
	db *qmgo.Database
}

func NewSupplyChainStore(db *qmgo.Database) *SupplyChainStore {
	return &SupplyChainStore{
		db: db,
	}
}

// This function is used to create the invetory items with specific fields as params and returns the created id
func (sc *SupplyChainStore) AddInventoryItem(name string, sku string, quantity int, warehouse string) (primitive.ObjectID, error) {
	ctx := context.Background()
	result, err := sc.db.Collection("inventory").InsertOne(ctx, bson.M{"name": name, "sku": sku, "quantity": quantity, "warehouse": warehouse})
	if err != nil {
		logs.ErrorLogger.Println("error while adding inventory", err)
		return primitive.NilObjectID, errors.New("error while adding inventory")
	}
	logs.InfoLogger.Println("Created Inventory", result)
	return result.InsertedID.(primitive.ObjectID), nil
}

// This function is used to update specific inventory item with specific fields as params and return the updated id
func (sc *SupplyChainStore) UpdateInventoryItem(id primitive.ObjectID, name *string, sku *string, quantity *int, warehouse *string) (primitive.ObjectID, error) {
	ctx := context.Background()
	err := sc.db.Collection("inventory").UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"name": name, "sku": sku, "quantity": quantity, "warehouse": warehouse}})
	if err != nil {
		logs.ErrorLogger.Println("error while updating inventory", err)
		return primitive.NilObjectID, errors.New("error while updating inventoru")
	}
	logs.InfoLogger.Println("updated inventory")
	return id, nil
}

// This function is used to create supplier details with specific fields in the params and returns the created id
func (sc *SupplyChainStore) AddSupplier(name string, contactPerson string, phone string, email string) (primitive.ObjectID, error) {
	ctx := context.Background()
	result, err := sc.db.Collection("supplier").InsertOne(ctx, bson.M{"name": name, "contactPerson": contactPerson, "phone": phone, "email": email})
	if err != nil {
		logs.InfoLogger.Println("error while adding supplier", err)
		return primitive.NilObjectID, errors.New("error while adding supplier")
	}
	logs.InfoLogger.Println("Created Supplier", result)
	return result.InsertedID.(primitive.ObjectID), nil
}

// This function is used to update supplier detials for a specific supplier only with fields in params and returns the updated id
func (sc *SupplyChainStore) UpdateSupplier(id primitive.ObjectID, name *string, contactPerson *string, phone *string, email *string) (primitive.ObjectID, error) {
	ctx := context.Background()
	err := sc.db.Collection("supplier").UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"name": name, "contactPerson": contactPerson, "phone": phone, "email": email}})
	if err != nil {
		logs.ErrorLogger.Println("error while updating supplier", err)
		return primitive.NilObjectID, errors.New("error while updating supplier")
	}
	logs.InfoLogger.Println("updated supplier")
	return id, nil
}

// This function is used to update shipment status and returns the updated shipment id
func (sc *SupplyChainStore) UpdateShipmentStatus(id primitive.ObjectID, status string) (primitive.ObjectID, error) {
	ctx := context.Background()
	err := sc.db.Collection("shipment").UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"status": status}})
	if err != nil {
		return primitive.NilObjectID, errors.New("error while updating shipment status")
	}
	logs.InfoLogger.Println("updated shipment status")
	return id, nil
}

// This function is used to delete paricular invenorty only and returns boolean value
func (sc *SupplyChainStore) DeleteInventoryItem(id primitive.ObjectID) bool {
	ctx := context.Background()
	err := sc.db.Collection("inventory").Remove(ctx, bson.M{"_id": id})
	if err != nil {
		return false
	} else {
		logs.InfoLogger.Println("deleted one inventory")
		return true
	}
}

// This function is used to delete particular supplier only and returns boolean value
func (sc *SupplyChainStore) DeleteSupplier(id primitive.ObjectID) bool {
	ctx := context.Background()
	err := sc.db.Collection("supplier").Remove(ctx, bson.M{"_id": id})
	if err != nil {
		return false
	} else {
		logs.InfoLogger.Println("deleted one supplier")
		return true
	}
}

// This funciton is used to read all the inventory items with limit and skip passed through params and returns number of inventory items
func (sc *SupplyChainStore) GetInventoryItems(limit int64, offset int64) ([]*model.InventoryItem, error) {
	var m []*model.InventoryItem
	ctx := context.Background()
	if err := sc.db.Collection("inventory").Find(ctx, bson.M{}).Limit(limit).Skip(offset).All(&m); err != nil {
		logs.ErrorLogger.Println("error while reading inventory items", err)
		return nil, errors.New("error while reading iventory items")
	}
	return m, nil
}

// This function returns single shipment for a particular shipment id
func (sc *SupplyChainStore) GetShipment(id primitive.ObjectID) (*model.Shipment, error) {
	var m model.Shipment
	ctx := context.Background()
	if err := sc.db.Collection("shipment").Find(ctx, bson.M{"_id": id}).One(&m); err != nil {
		logs.ErrorLogger.Println("error while reading shipment", err)
		return nil, errors.New("no data found or error while reading shipment")
	}
	return &m, nil
}

// This function returns all supplier details
func (sc *SupplyChainStore) GetSuppliers() ([]*model.Supplier, error) {
	var m []*model.Supplier
	ctx := context.Background()
	if err := sc.db.Collection("supplier").Find(ctx, bson.M{}).All(&m); err != nil {
		logs.ErrorLogger.Println("error while reading supplier", err)
		return nil, errors.New("no data found or error while reading supplier")
	}
	return m, nil
}

// This funtion returns single inventory for a particular inventory id
func (sc *SupplyChainStore) GetSingleInventory(id primitive.ObjectID) (*model.InventoryItem, error) {
	var m model.InventoryItem
	ctx := context.Background()
	if err := sc.db.Collection("inventory").Find(ctx, bson.M{"_id": id}).One(&m); err != nil {
		logs.ErrorLogger.Println("error while reading single inventory", err)
		return nil, errors.New("mo data found or error in inventory data")
	}
	return &m, nil
}

// This function returns single supplier for a particular supplier id
func (sc *SupplyChainStore) GetSingleSupplier(id primitive.ObjectID) (*model.Supplier, error) {
	var m model.Supplier
	ctx := context.Background()
	if err := sc.db.Collection("supplier").Find(ctx, bson.M{"_id": id}).One(&m); err != nil {
		logs.ErrorLogger.Println("error while reading single supplier", err)
		return nil, errors.New("no data found or error in supplier data")
	}
	return &m, nil
}
