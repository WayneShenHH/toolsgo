// Package nosql is interface for nssql db
package nosql

// Handler for access to datastore
type Handler interface {
	// Get retrieves a entity by its ID, entity should be a pointer of entity, ex: &entity{}
	Get(entityName string, id int64, entity interface{}) error
	// Put saves a given entity, assigning it with ID, entity should be a pointer of entity, ex: &entity{}
	Put(entityName string, id int64, entity interface{}) error
	// PutMulti saves a given slice of entity, input slice should be a slice of pointer, ex: []*entity{}
	PutMulti(entityName string, ids []int64, slice interface{}) error
	// List returns a list of entities, input slice should be a pointer of slice, ex: &[]entity{}
	List(entityName string, query map[string]interface{}, slice interface{}) error
	// Delete removes a given entity by its ID.
	Delete(entityName string, id int64) error
	// Delete removes a given entities by its IDs.
	DeleteMulti(entityName string, ids []int64) error
}
