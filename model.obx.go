// Code generated by ObjectBox; DO NOT EDIT.
// Learn more about defining entities and generating this file - visit https://golang.objectbox.io/entity-annotations

package main

import (
	"errors"
	"github.com/google/flatbuffers/go"
	"github.com/objectbox/objectbox-go/objectbox"
	"github.com/objectbox/objectbox-go/objectbox/fbutils"
)

type resource_EntityInfo struct {
	objectbox.Entity
	Uid uint64
}

var ResourceBinding = resource_EntityInfo{
	Entity: objectbox.Entity{
		Id: 1,
	},
	Uid: 8864280590638099451,
}

// Resource_ contains type-based Property helpers to facilitate some common operations such as Queries.
var Resource_ = struct {
	Id      *objectbox.PropertyUint64
	Type    *objectbox.PropertyString
	Name    *objectbox.PropertyString
	Amount  *objectbox.PropertyFloat64
	EBC     *objectbox.PropertyFloat64
	MinTemp *objectbox.PropertyFloat64
	MaxTemp *objectbox.PropertyFloat64
	OberG   *objectbox.PropertyBool
	ISO     *objectbox.PropertyFloat64
	Opened  *objectbox.PropertyInt64
	Exp     *objectbox.PropertyInt64
}{
	Id: &objectbox.PropertyUint64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     1,
			Entity: &ResourceBinding.Entity,
		},
	},
	Type: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     2,
			Entity: &ResourceBinding.Entity,
		},
	},
	Name: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     3,
			Entity: &ResourceBinding.Entity,
		},
	},
	Amount: &objectbox.PropertyFloat64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     4,
			Entity: &ResourceBinding.Entity,
		},
	},
	EBC: &objectbox.PropertyFloat64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     5,
			Entity: &ResourceBinding.Entity,
		},
	},
	MinTemp: &objectbox.PropertyFloat64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     6,
			Entity: &ResourceBinding.Entity,
		},
	},
	MaxTemp: &objectbox.PropertyFloat64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     7,
			Entity: &ResourceBinding.Entity,
		},
	},
	OberG: &objectbox.PropertyBool{
		BaseProperty: &objectbox.BaseProperty{
			Id:     8,
			Entity: &ResourceBinding.Entity,
		},
	},
	ISO: &objectbox.PropertyFloat64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     9,
			Entity: &ResourceBinding.Entity,
		},
	},
	Opened: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     10,
			Entity: &ResourceBinding.Entity,
		},
	},
	Exp: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     11,
			Entity: &ResourceBinding.Entity,
		},
	},
}

// GeneratorVersion is called by ObjectBox to verify the compatibility of the generator used to generate this code
func (resource_EntityInfo) GeneratorVersion() int {
	return 6
}

// AddToModel is called by ObjectBox during model build
func (resource_EntityInfo) AddToModel(model *objectbox.Model) {
	model.Entity("Resource", 1, 8864280590638099451)
	model.Property("Id", 6, 1, 5493231812372595979)
	model.PropertyFlags(1)
	model.Property("Type", 9, 2, 887142186992516674)
	model.Property("Name", 9, 3, 7704535535632126273)
	model.Property("Amount", 8, 4, 3015859585488509340)
	model.Property("EBC", 8, 5, 8027933678563871299)
	model.Property("MinTemp", 8, 6, 6978596190461251037)
	model.Property("MaxTemp", 8, 7, 1466987121535296644)
	model.Property("OberG", 1, 8, 6482886941541151457)
	model.Property("ISO", 8, 9, 6879353326633125653)
	model.Property("Opened", 10, 10, 2589939524641695745)
	model.Property("Exp", 10, 11, 1865559351558115529)
	model.EntityLastPropertyId(11, 1865559351558115529)
}

// GetId is called by ObjectBox during Put operations to check for existing ID on an object
func (resource_EntityInfo) GetId(object interface{}) (uint64, error) {
	return object.(*Resource).Id, nil
}

// SetId is called by ObjectBox during Put to update an ID on an object that has just been inserted
func (resource_EntityInfo) SetId(object interface{}, id uint64) error {
	object.(*Resource).Id = id
	return nil
}

// PutRelated is called by ObjectBox to put related entities before the object itself is flattened and put
func (resource_EntityInfo) PutRelated(ob *objectbox.ObjectBox, object interface{}, id uint64) error {
	return nil
}

// Flatten is called by ObjectBox to transform an object to a FlatBuffer
func (resource_EntityInfo) Flatten(object interface{}, fbb *flatbuffers.Builder, id uint64) error {
	obj := object.(*Resource)
	var propOpened int64
	{
		var err error
		propOpened, err = objectbox.TimeInt64ConvertToDatabaseValue(obj.Opened)
		if err != nil {
			return errors.New("converter objectbox.TimeInt64ConvertToDatabaseValue() failed on Resource.Opened: " + err.Error())
		}
	}

	var propExp int64
	{
		var err error
		propExp, err = objectbox.TimeInt64ConvertToDatabaseValue(obj.Exp)
		if err != nil {
			return errors.New("converter objectbox.TimeInt64ConvertToDatabaseValue() failed on Resource.Exp: " + err.Error())
		}
	}

	var offsetType = fbutils.CreateStringOffset(fbb, obj.Type)
	var offsetName = fbutils.CreateStringOffset(fbb, obj.Name)

	// build the FlatBuffers object
	fbb.StartObject(11)
	fbutils.SetUint64Slot(fbb, 0, id)
	fbutils.SetUOffsetTSlot(fbb, 1, offsetType)
	fbutils.SetUOffsetTSlot(fbb, 2, offsetName)
	fbutils.SetFloat64Slot(fbb, 3, obj.Amount)
	fbutils.SetFloat64Slot(fbb, 4, obj.EBC)
	fbutils.SetFloat64Slot(fbb, 5, obj.MinTemp)
	fbutils.SetFloat64Slot(fbb, 6, obj.MaxTemp)
	fbutils.SetBoolSlot(fbb, 7, obj.OberG)
	fbutils.SetFloat64Slot(fbb, 8, obj.ISO)
	fbutils.SetInt64Slot(fbb, 9, propOpened)
	fbutils.SetInt64Slot(fbb, 10, propExp)
	return nil
}

// Load is called by ObjectBox to load an object from a FlatBuffer
func (resource_EntityInfo) Load(ob *objectbox.ObjectBox, bytes []byte) (interface{}, error) {
	if len(bytes) == 0 { // sanity check, should "never" happen
		return nil, errors.New("can't deserialize an object of type 'Resource' - no data received")
	}

	var table = &flatbuffers.Table{
		Bytes: bytes,
		Pos:   flatbuffers.GetUOffsetT(bytes),
	}

	var propId = table.GetUint64Slot(4, 0)

	propOpened, err := objectbox.TimeInt64ConvertToEntityProperty(fbutils.GetInt64Slot(table, 22))
	if err != nil {
		return nil, errors.New("converter objectbox.TimeInt64ConvertToEntityProperty() failed on Resource.Opened: " + err.Error())
	}

	propExp, err := objectbox.TimeInt64ConvertToEntityProperty(fbutils.GetInt64Slot(table, 24))
	if err != nil {
		return nil, errors.New("converter objectbox.TimeInt64ConvertToEntityProperty() failed on Resource.Exp: " + err.Error())
	}

	return &Resource{
		Id:      propId,
		Type:    fbutils.GetStringSlot(table, 6),
		Name:    fbutils.GetStringSlot(table, 8),
		Amount:  fbutils.GetFloat64Slot(table, 10),
		EBC:     fbutils.GetFloat64Slot(table, 12),
		MinTemp: fbutils.GetFloat64Slot(table, 14),
		MaxTemp: fbutils.GetFloat64Slot(table, 16),
		OberG:   fbutils.GetBoolSlot(table, 18),
		ISO:     fbutils.GetFloat64Slot(table, 20),
		Opened:  propOpened,
		Exp:     propExp,
	}, nil
}

// MakeSlice is called by ObjectBox to construct a new slice to hold the read objects
func (resource_EntityInfo) MakeSlice(capacity int) interface{} {
	return make([]*Resource, 0, capacity)
}

// AppendToSlice is called by ObjectBox to fill the slice of the read objects
func (resource_EntityInfo) AppendToSlice(slice interface{}, object interface{}) interface{} {
	if object == nil {
		return append(slice.([]*Resource), nil)
	}
	return append(slice.([]*Resource), object.(*Resource))
}

// Box provides CRUD access to Resource objects
type ResourceBox struct {
	*objectbox.Box
}

// BoxForResource opens a box of Resource objects
func BoxForResource(ob *objectbox.ObjectBox) *ResourceBox {
	return &ResourceBox{
		Box: ob.InternalBox(1),
	}
}

// Put synchronously inserts/updates a single object.
// In case the Id is not specified, it would be assigned automatically (auto-increment).
// When inserting, the Resource.Id property on the passed object will be assigned the new ID as well.
func (box *ResourceBox) Put(object *Resource) (uint64, error) {
	return box.Box.Put(object)
}

// Insert synchronously inserts a single object. As opposed to Put, Insert will fail if given an ID that already exists.
// In case the Id is not specified, it would be assigned automatically (auto-increment).
// When inserting, the Resource.Id property on the passed object will be assigned the new ID as well.
func (box *ResourceBox) Insert(object *Resource) (uint64, error) {
	return box.Box.Insert(object)
}

// Update synchronously updates a single object.
// As opposed to Put, Update will fail if an object with the same ID is not found in the database.
func (box *ResourceBox) Update(object *Resource) error {
	return box.Box.Update(object)
}

// PutAsync asynchronously inserts/updates a single object.
// Deprecated: use box.Async().Put() instead
func (box *ResourceBox) PutAsync(object *Resource) (uint64, error) {
	return box.Box.PutAsync(object)
}

// PutMany inserts multiple objects in single transaction.
// In case Ids are not set on the objects, they would be assigned automatically (auto-increment).
//
// Returns: IDs of the put objects (in the same order).
// When inserting, the Resource.Id property on the objects in the slice will be assigned the new IDs as well.
//
// Note: In case an error occurs during the transaction, some of the objects may already have the Resource.Id assigned
// even though the transaction has been rolled back and the objects are not stored under those IDs.
//
// Note: The slice may be empty or even nil; in both cases, an empty IDs slice and no error is returned.
func (box *ResourceBox) PutMany(objects []*Resource) ([]uint64, error) {
	return box.Box.PutMany(objects)
}

// Get reads a single object.
//
// Returns nil (and no error) in case the object with the given ID doesn't exist.
func (box *ResourceBox) Get(id uint64) (*Resource, error) {
	object, err := box.Box.Get(id)
	if err != nil {
		return nil, err
	} else if object == nil {
		return nil, nil
	}
	return object.(*Resource), nil
}

// GetMany reads multiple objects at once.
// If any of the objects doesn't exist, its position in the return slice is nil
func (box *ResourceBox) GetMany(ids ...uint64) ([]*Resource, error) {
	objects, err := box.Box.GetMany(ids...)
	if err != nil {
		return nil, err
	}
	return objects.([]*Resource), nil
}

// GetManyExisting reads multiple objects at once, skipping those that do not exist.
func (box *ResourceBox) GetManyExisting(ids ...uint64) ([]*Resource, error) {
	objects, err := box.Box.GetManyExisting(ids...)
	if err != nil {
		return nil, err
	}
	return objects.([]*Resource), nil
}

// GetAll reads all stored objects
func (box *ResourceBox) GetAll() ([]*Resource, error) {
	objects, err := box.Box.GetAll()
	if err != nil {
		return nil, err
	}
	return objects.([]*Resource), nil
}

// Remove deletes a single object
func (box *ResourceBox) Remove(object *Resource) error {
	return box.Box.Remove(object)
}

// RemoveMany deletes multiple objects at once.
// Returns the number of deleted object or error on failure.
// Note that this method will not fail if an object is not found (e.g. already removed).
// In case you need to strictly check whether all of the objects exist before removing them,
// you can execute multiple box.Contains() and box.Remove() inside a single write transaction.
func (box *ResourceBox) RemoveMany(objects ...*Resource) (uint64, error) {
	var ids = make([]uint64, len(objects))
	for k, object := range objects {
		ids[k] = object.Id
	}
	return box.Box.RemoveIds(ids...)
}

// Creates a query with the given conditions. Use the fields of the Resource_ struct to create conditions.
// Keep the *ResourceQuery if you intend to execute the query multiple times.
// Note: this function panics if you try to create illegal queries; e.g. use properties of an alien type.
// This is typically a programming error. Use QueryOrError instead if you want the explicit error check.
func (box *ResourceBox) Query(conditions ...objectbox.Condition) *ResourceQuery {
	return &ResourceQuery{
		box.Box.Query(conditions...),
	}
}

// Creates a query with the given conditions. Use the fields of the Resource_ struct to create conditions.
// Keep the *ResourceQuery if you intend to execute the query multiple times.
func (box *ResourceBox) QueryOrError(conditions ...objectbox.Condition) (*ResourceQuery, error) {
	if query, err := box.Box.QueryOrError(conditions...); err != nil {
		return nil, err
	} else {
		return &ResourceQuery{query}, nil
	}
}

// Async provides access to the default Async Box for asynchronous operations. See ResourceAsyncBox for more information.
func (box *ResourceBox) Async() *ResourceAsyncBox {
	return &ResourceAsyncBox{AsyncBox: box.Box.Async()}
}

// ResourceAsyncBox provides asynchronous operations on Resource objects.
//
// Asynchronous operations are executed on a separate internal thread for better performance.
//
// There are two main use cases:
//
// 1) "execute & forget:" you gain faster put/remove operations as you don't have to wait for the transaction to finish.
//
// 2) Many small transactions: if your write load is typically a lot of individual puts that happen in parallel,
// this will merge small transactions into bigger ones. This results in a significant gain in overall throughput.
//
// In situations with (extremely) high async load, an async method may be throttled (~1ms) or delayed up to 1 second.
// In the unlikely event that the object could still not be enqueued (full queue), an error will be returned.
//
// Note that async methods do not give you hard durability guarantees like the synchronous Box provides.
// There is a small time window in which the data may not have been committed durably yet.
type ResourceAsyncBox struct {
	*objectbox.AsyncBox
}

// AsyncBoxForResource creates a new async box with the given operation timeout in case an async queue is full.
// The returned struct must be freed explicitly using the Close() method.
// It's usually preferable to use ResourceBox::Async() which takes care of resource management and doesn't require closing.
func AsyncBoxForResource(ob *objectbox.ObjectBox, timeoutMs uint64) *ResourceAsyncBox {
	var async, err = objectbox.NewAsyncBox(ob, 1, timeoutMs)
	if err != nil {
		panic("Could not create async box for entity ID 1: %s" + err.Error())
	}
	return &ResourceAsyncBox{AsyncBox: async}
}

// Put inserts/updates a single object asynchronously.
// When inserting a new object, the Id property on the passed object will be assigned the new ID the entity would hold
// if the insert is ultimately successful. The newly assigned ID may not become valid if the insert fails.
func (asyncBox *ResourceAsyncBox) Put(object *Resource) (uint64, error) {
	return asyncBox.AsyncBox.Put(object)
}

// Insert a single object asynchronously.
// The Id property on the passed object will be assigned the new ID the entity would hold if the insert is ultimately
// successful. The newly assigned ID may not become valid if the insert fails.
// Fails silently if an object with the same ID already exists (this error is not returned).
func (asyncBox *ResourceAsyncBox) Insert(object *Resource) (id uint64, err error) {
	return asyncBox.AsyncBox.Insert(object)
}

// Update a single object asynchronously.
// The object must already exists or the update fails silently (without an error returned).
func (asyncBox *ResourceAsyncBox) Update(object *Resource) error {
	return asyncBox.AsyncBox.Update(object)
}

// Remove deletes a single object asynchronously.
func (asyncBox *ResourceAsyncBox) Remove(object *Resource) error {
	return asyncBox.AsyncBox.Remove(object)
}

// Query provides a way to search stored objects
//
// For example, you can find all Resource which Id is either 42 or 47:
// 		box.Query(Resource_.Id.In(42, 47)).Find()
type ResourceQuery struct {
	*objectbox.Query
}

// Find returns all objects matching the query
func (query *ResourceQuery) Find() ([]*Resource, error) {
	objects, err := query.Query.Find()
	if err != nil {
		return nil, err
	}
	return objects.([]*Resource), nil
}

// Offset defines the index of the first object to process (how many objects to skip)
func (query *ResourceQuery) Offset(offset uint64) *ResourceQuery {
	query.Query.Offset(offset)
	return query
}

// Limit sets the number of elements to process by the query
func (query *ResourceQuery) Limit(limit uint64) *ResourceQuery {
	query.Query.Limit(limit)
	return query
}

type recipe_EntityInfo struct {
	objectbox.Entity
	Uid uint64
}

var RecipeBinding = recipe_EntityInfo{
	Entity: objectbox.Entity{
		Id: 2,
	},
	Uid: 7524476822026071533,
}

// Recipe_ contains type-based Property helpers to facilitate some common operations such as Queries.
var Recipe_ = struct {
	Id          *objectbox.PropertyUint64
	Name        *objectbox.PropertyString
	Description *objectbox.PropertyString
	IsoAlpha    *objectbox.PropertyFloat64
	HopSugg     *objectbox.PropertyStringVector
	DryHop      *objectbox.PropertyStringVector
	SHA         *objectbox.PropertyFloat64
	AlcTarget   *objectbox.PropertyFloat64
	OGTarget    *objectbox.PropertyFloat64
}{
	Id: &objectbox.PropertyUint64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     1,
			Entity: &RecipeBinding.Entity,
		},
	},
	Name: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     2,
			Entity: &RecipeBinding.Entity,
		},
	},
	Description: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     3,
			Entity: &RecipeBinding.Entity,
		},
	},
	IsoAlpha: &objectbox.PropertyFloat64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     4,
			Entity: &RecipeBinding.Entity,
		},
	},
	HopSugg: &objectbox.PropertyStringVector{
		BaseProperty: &objectbox.BaseProperty{
			Id:     5,
			Entity: &RecipeBinding.Entity,
		},
	},
	DryHop: &objectbox.PropertyStringVector{
		BaseProperty: &objectbox.BaseProperty{
			Id:     6,
			Entity: &RecipeBinding.Entity,
		},
	},
	SHA: &objectbox.PropertyFloat64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     7,
			Entity: &RecipeBinding.Entity,
		},
	},
	AlcTarget: &objectbox.PropertyFloat64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     8,
			Entity: &RecipeBinding.Entity,
		},
	},
	OGTarget: &objectbox.PropertyFloat64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     9,
			Entity: &RecipeBinding.Entity,
		},
	},
}

// GeneratorVersion is called by ObjectBox to verify the compatibility of the generator used to generate this code
func (recipe_EntityInfo) GeneratorVersion() int {
	return 6
}

// AddToModel is called by ObjectBox during model build
func (recipe_EntityInfo) AddToModel(model *objectbox.Model) {
	model.Entity("Recipe", 2, 7524476822026071533)
	model.Property("Id", 6, 1, 385067527585461866)
	model.PropertyFlags(1)
	model.Property("Name", 9, 2, 6590878796668904790)
	model.Property("Description", 9, 3, 6111302923799653287)
	model.Property("IsoAlpha", 8, 4, 6675220912961035469)
	model.Property("HopSugg", 30, 5, 6003672669915185166)
	model.Property("DryHop", 30, 6, 1837696198939072290)
	model.Property("SHA", 8, 7, 7365373180582720820)
	model.Property("AlcTarget", 8, 8, 77502409174959323)
	model.Property("OGTarget", 8, 9, 7723634411214686482)
	model.EntityLastPropertyId(9, 7723634411214686482)
}

// GetId is called by ObjectBox during Put operations to check for existing ID on an object
func (recipe_EntityInfo) GetId(object interface{}) (uint64, error) {
	return object.(*Recipe).Id, nil
}

// SetId is called by ObjectBox during Put to update an ID on an object that has just been inserted
func (recipe_EntityInfo) SetId(object interface{}, id uint64) error {
	object.(*Recipe).Id = id
	return nil
}

// PutRelated is called by ObjectBox to put related entities before the object itself is flattened and put
func (recipe_EntityInfo) PutRelated(ob *objectbox.ObjectBox, object interface{}, id uint64) error {
	return nil
}

// Flatten is called by ObjectBox to transform an object to a FlatBuffer
func (recipe_EntityInfo) Flatten(object interface{}, fbb *flatbuffers.Builder, id uint64) error {
	obj := object.(*Recipe)
	var offsetName = fbutils.CreateStringOffset(fbb, obj.Name)
	var offsetDescription = fbutils.CreateStringOffset(fbb, obj.Description)
	var offsetHopSugg = fbutils.CreateStringVectorOffset(fbb, obj.HopSugg)
	var offsetDryHop = fbutils.CreateStringVectorOffset(fbb, obj.DryHop)

	// build the FlatBuffers object
	fbb.StartObject(9)
	fbutils.SetUint64Slot(fbb, 0, id)
	fbutils.SetUOffsetTSlot(fbb, 1, offsetName)
	fbutils.SetUOffsetTSlot(fbb, 2, offsetDescription)
	fbutils.SetFloat64Slot(fbb, 3, obj.IsoAlpha)
	fbutils.SetUOffsetTSlot(fbb, 4, offsetHopSugg)
	fbutils.SetUOffsetTSlot(fbb, 5, offsetDryHop)
	fbutils.SetFloat64Slot(fbb, 6, obj.SHA)
	fbutils.SetFloat64Slot(fbb, 7, obj.AlcTarget)
	fbutils.SetFloat64Slot(fbb, 8, obj.OGTarget)
	return nil
}

// Load is called by ObjectBox to load an object from a FlatBuffer
func (recipe_EntityInfo) Load(ob *objectbox.ObjectBox, bytes []byte) (interface{}, error) {
	if len(bytes) == 0 { // sanity check, should "never" happen
		return nil, errors.New("can't deserialize an object of type 'Recipe' - no data received")
	}

	var table = &flatbuffers.Table{
		Bytes: bytes,
		Pos:   flatbuffers.GetUOffsetT(bytes),
	}

	var propId = table.GetUint64Slot(4, 0)

	return &Recipe{
		Id:          propId,
		Name:        fbutils.GetStringSlot(table, 6),
		Description: fbutils.GetStringSlot(table, 8),
		IsoAlpha:    fbutils.GetFloat64Slot(table, 10),
		HopSugg:     fbutils.GetStringVectorSlot(table, 12),
		DryHop:      fbutils.GetStringVectorSlot(table, 14),
		SHA:         fbutils.GetFloat64Slot(table, 16),
		AlcTarget:   fbutils.GetFloat64Slot(table, 18),
		OGTarget:    fbutils.GetFloat64Slot(table, 20),
	}, nil
}

// MakeSlice is called by ObjectBox to construct a new slice to hold the read objects
func (recipe_EntityInfo) MakeSlice(capacity int) interface{} {
	return make([]*Recipe, 0, capacity)
}

// AppendToSlice is called by ObjectBox to fill the slice of the read objects
func (recipe_EntityInfo) AppendToSlice(slice interface{}, object interface{}) interface{} {
	if object == nil {
		return append(slice.([]*Recipe), nil)
	}
	return append(slice.([]*Recipe), object.(*Recipe))
}

// Box provides CRUD access to Recipe objects
type RecipeBox struct {
	*objectbox.Box
}

// BoxForRecipe opens a box of Recipe objects
func BoxForRecipe(ob *objectbox.ObjectBox) *RecipeBox {
	return &RecipeBox{
		Box: ob.InternalBox(2),
	}
}

// Put synchronously inserts/updates a single object.
// In case the Id is not specified, it would be assigned automatically (auto-increment).
// When inserting, the Recipe.Id property on the passed object will be assigned the new ID as well.
func (box *RecipeBox) Put(object *Recipe) (uint64, error) {
	return box.Box.Put(object)
}

// Insert synchronously inserts a single object. As opposed to Put, Insert will fail if given an ID that already exists.
// In case the Id is not specified, it would be assigned automatically (auto-increment).
// When inserting, the Recipe.Id property on the passed object will be assigned the new ID as well.
func (box *RecipeBox) Insert(object *Recipe) (uint64, error) {
	return box.Box.Insert(object)
}

// Update synchronously updates a single object.
// As opposed to Put, Update will fail if an object with the same ID is not found in the database.
func (box *RecipeBox) Update(object *Recipe) error {
	return box.Box.Update(object)
}

// PutAsync asynchronously inserts/updates a single object.
// Deprecated: use box.Async().Put() instead
func (box *RecipeBox) PutAsync(object *Recipe) (uint64, error) {
	return box.Box.PutAsync(object)
}

// PutMany inserts multiple objects in single transaction.
// In case Ids are not set on the objects, they would be assigned automatically (auto-increment).
//
// Returns: IDs of the put objects (in the same order).
// When inserting, the Recipe.Id property on the objects in the slice will be assigned the new IDs as well.
//
// Note: In case an error occurs during the transaction, some of the objects may already have the Recipe.Id assigned
// even though the transaction has been rolled back and the objects are not stored under those IDs.
//
// Note: The slice may be empty or even nil; in both cases, an empty IDs slice and no error is returned.
func (box *RecipeBox) PutMany(objects []*Recipe) ([]uint64, error) {
	return box.Box.PutMany(objects)
}

// Get reads a single object.
//
// Returns nil (and no error) in case the object with the given ID doesn't exist.
func (box *RecipeBox) Get(id uint64) (*Recipe, error) {
	object, err := box.Box.Get(id)
	if err != nil {
		return nil, err
	} else if object == nil {
		return nil, nil
	}
	return object.(*Recipe), nil
}

// GetMany reads multiple objects at once.
// If any of the objects doesn't exist, its position in the return slice is nil
func (box *RecipeBox) GetMany(ids ...uint64) ([]*Recipe, error) {
	objects, err := box.Box.GetMany(ids...)
	if err != nil {
		return nil, err
	}
	return objects.([]*Recipe), nil
}

// GetManyExisting reads multiple objects at once, skipping those that do not exist.
func (box *RecipeBox) GetManyExisting(ids ...uint64) ([]*Recipe, error) {
	objects, err := box.Box.GetManyExisting(ids...)
	if err != nil {
		return nil, err
	}
	return objects.([]*Recipe), nil
}

// GetAll reads all stored objects
func (box *RecipeBox) GetAll() ([]*Recipe, error) {
	objects, err := box.Box.GetAll()
	if err != nil {
		return nil, err
	}
	return objects.([]*Recipe), nil
}

// Remove deletes a single object
func (box *RecipeBox) Remove(object *Recipe) error {
	return box.Box.Remove(object)
}

// RemoveMany deletes multiple objects at once.
// Returns the number of deleted object or error on failure.
// Note that this method will not fail if an object is not found (e.g. already removed).
// In case you need to strictly check whether all of the objects exist before removing them,
// you can execute multiple box.Contains() and box.Remove() inside a single write transaction.
func (box *RecipeBox) RemoveMany(objects ...*Recipe) (uint64, error) {
	var ids = make([]uint64, len(objects))
	for k, object := range objects {
		ids[k] = object.Id
	}
	return box.Box.RemoveIds(ids...)
}

// Creates a query with the given conditions. Use the fields of the Recipe_ struct to create conditions.
// Keep the *RecipeQuery if you intend to execute the query multiple times.
// Note: this function panics if you try to create illegal queries; e.g. use properties of an alien type.
// This is typically a programming error. Use QueryOrError instead if you want the explicit error check.
func (box *RecipeBox) Query(conditions ...objectbox.Condition) *RecipeQuery {
	return &RecipeQuery{
		box.Box.Query(conditions...),
	}
}

// Creates a query with the given conditions. Use the fields of the Recipe_ struct to create conditions.
// Keep the *RecipeQuery if you intend to execute the query multiple times.
func (box *RecipeBox) QueryOrError(conditions ...objectbox.Condition) (*RecipeQuery, error) {
	if query, err := box.Box.QueryOrError(conditions...); err != nil {
		return nil, err
	} else {
		return &RecipeQuery{query}, nil
	}
}

// Async provides access to the default Async Box for asynchronous operations. See RecipeAsyncBox for more information.
func (box *RecipeBox) Async() *RecipeAsyncBox {
	return &RecipeAsyncBox{AsyncBox: box.Box.Async()}
}

// RecipeAsyncBox provides asynchronous operations on Recipe objects.
//
// Asynchronous operations are executed on a separate internal thread for better performance.
//
// There are two main use cases:
//
// 1) "execute & forget:" you gain faster put/remove operations as you don't have to wait for the transaction to finish.
//
// 2) Many small transactions: if your write load is typically a lot of individual puts that happen in parallel,
// this will merge small transactions into bigger ones. This results in a significant gain in overall throughput.
//
// In situations with (extremely) high async load, an async method may be throttled (~1ms) or delayed up to 1 second.
// In the unlikely event that the object could still not be enqueued (full queue), an error will be returned.
//
// Note that async methods do not give you hard durability guarantees like the synchronous Box provides.
// There is a small time window in which the data may not have been committed durably yet.
type RecipeAsyncBox struct {
	*objectbox.AsyncBox
}

// AsyncBoxForRecipe creates a new async box with the given operation timeout in case an async queue is full.
// The returned struct must be freed explicitly using the Close() method.
// It's usually preferable to use RecipeBox::Async() which takes care of resource management and doesn't require closing.
func AsyncBoxForRecipe(ob *objectbox.ObjectBox, timeoutMs uint64) *RecipeAsyncBox {
	var async, err = objectbox.NewAsyncBox(ob, 2, timeoutMs)
	if err != nil {
		panic("Could not create async box for entity ID 2: %s" + err.Error())
	}
	return &RecipeAsyncBox{AsyncBox: async}
}

// Put inserts/updates a single object asynchronously.
// When inserting a new object, the Id property on the passed object will be assigned the new ID the entity would hold
// if the insert is ultimately successful. The newly assigned ID may not become valid if the insert fails.
func (asyncBox *RecipeAsyncBox) Put(object *Recipe) (uint64, error) {
	return asyncBox.AsyncBox.Put(object)
}

// Insert a single object asynchronously.
// The Id property on the passed object will be assigned the new ID the entity would hold if the insert is ultimately
// successful. The newly assigned ID may not become valid if the insert fails.
// Fails silently if an object with the same ID already exists (this error is not returned).
func (asyncBox *RecipeAsyncBox) Insert(object *Recipe) (id uint64, err error) {
	return asyncBox.AsyncBox.Insert(object)
}

// Update a single object asynchronously.
// The object must already exists or the update fails silently (without an error returned).
func (asyncBox *RecipeAsyncBox) Update(object *Recipe) error {
	return asyncBox.AsyncBox.Update(object)
}

// Remove deletes a single object asynchronously.
func (asyncBox *RecipeAsyncBox) Remove(object *Recipe) error {
	return asyncBox.AsyncBox.Remove(object)
}

// Query provides a way to search stored objects
//
// For example, you can find all Recipe which Id is either 42 or 47:
// 		box.Query(Recipe_.Id.In(42, 47)).Find()
type RecipeQuery struct {
	*objectbox.Query
}

// Find returns all objects matching the query
func (query *RecipeQuery) Find() ([]*Recipe, error) {
	objects, err := query.Query.Find()
	if err != nil {
		return nil, err
	}
	return objects.([]*Recipe), nil
}

// Offset defines the index of the first object to process (how many objects to skip)
func (query *RecipeQuery) Offset(offset uint64) *RecipeQuery {
	query.Query.Offset(offset)
	return query
}

// Limit sets the number of elements to process by the query
func (query *RecipeQuery) Limit(limit uint64) *RecipeQuery {
	query.Query.Limit(limit)
	return query
}
