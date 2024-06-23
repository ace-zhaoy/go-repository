package contract

import "context"

type CrudRepository[ID comparable, T ENTITY[ID]] interface {
	Unscoped() CrudRepository[ID, T]
	IDField() string
	SoftDeleteField() string
	SoftDeleteEnabled() bool
	Create(ctx context.Context, entity T) (id ID, err error)
	FindByID(ctx context.Context, id ID) (entity T, err error)
	FindByIDs(ctx context.Context, ids []ID) (collection Collection[ID, T], err error)
	FindByPage(ctx context.Context, limit, offset int, orders ...Order) (collection Collection[ID, T], err error)
	FindByFilter(ctx context.Context, filter map[string]any) (collection Collection[ID, T], err error)
	FindByFilterWithPage(ctx context.Context, filter map[string]any, limit, offset int, orders ...Order) (collection Collection[ID, T], err error)
	FindAll(ctx context.Context) (collection Collection[ID, T], err error)
	Count(ctx context.Context) (count int, err error)
	CountByFilter(ctx context.Context, filter map[string]any) (count int, err error)
	Exists(ctx context.Context, filter map[string]any) (exists bool, err error)
	ExistsByID(ctx context.Context, id ID) (exists bool, err error)
	ExistsByIDs(ctx context.Context, ids []ID) (exists Dict[ID, bool], err error)
	Update(ctx context.Context, filter map[string]any, data map[string]any) (err error)
	UpdateByID(ctx context.Context, id ID, data map[string]any) (err error)
	UpdateNonZero(ctx context.Context, filter map[string]any, entity T) (err error)
	UpdateNonZeroByID(ctx context.Context, id ID, entity T) (err error)
	Delete(ctx context.Context, filter map[string]any) (err error)
	DeleteByID(ctx context.Context, id ID) (err error)
	DeleteByIDs(ctx context.Context, ids []ID) (err error)
	DeleteAll(ctx context.Context) (err error)
}
