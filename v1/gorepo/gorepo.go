package gorepo

import "context"

type BatchOperationResponse struct {
	RowsAffected int64
}

// Crudable : curdable interface
type Crudable[t CrudModel] interface {
	ReadOnly[t]
	WriteOnly[t]
}

// ReadOnly : read only operations for a crud store
type ReadOnly[t CrudModel] interface {
	FindAll() ([]*t, error)
	FindAllC(ctx context.Context) ([]*t, error)
	FindAllPaginated(p *PaginationInput) (PaginationResponse[t], error)
	FindAllPaginatedC(ctx context.Context, p *PaginationInput) (PaginationResponse[t], error)
	FindByIndex(i *FieldInput) (*t, error)
	FindByIndexC(ctx context.Context, i *FieldInput) (*t, error)
	FindByFields(i *FieldInput) ([]*t, error)
	FindByFieldsC(ctx context.Context, i *FieldInput) ([]*t, error)
}

type WriteOnly[t CrudModel] interface {
	DeleteByIndex(i *FieldInput) error
	DeleteByIndexC(ctx context.Context, i *FieldInput) error
	DeleteManyByFields(i *FieldInput) (BatchOperationResponse, error)
	DeleteManyByFieldsC(ctx context.Context, i *FieldInput) (BatchOperationResponse, error)
	Create(m *t) error
	CreateC(ctx context.Context, m *t) error
	Update(m *t) error
	UpdateC(ctx context.Context, m *t) error
	Upsert(m *t) error
	UpsertC(ctx context.Context, m *t) error
	CreateMany(ms []*t) (*BatchOperationResponse, error)
	CreateManyC(ctx context.Context, ms []*t) (*BatchOperationResponse, error)
	UpdateMany(ms []*t) (*BatchOperationResponse, error)
	UpdateManyC(ctx context.Context, ms []*t) (*BatchOperationResponse, error)
	UpsertMany(ms []*t) (*BatchOperationResponse, error)
	UpsertManyC(ctx context.Context, ms []*t) (*BatchOperationResponse, error)
}
