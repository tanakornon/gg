package gg

type Entities interface {
	interface{}
}

type IService[T Entities] interface {
	GetAll(query QueryParams) (GetAllResponse[T], error)
	Get(id uint) (GetResponse[T], error)
	Create(req T) (GetResponse[T], error)
	CreateMany(req []T) (GetResponse[[]T], error)
	Upsert(id uint, entity T) (GetResponse[T], error)
	Update(id uint, entity T) (GetResponse[T], error)
	Delete(id uint) error
}

type IRepository[T Entities] interface {
	Count(search SearchParams) (int64, error)
	GetAll(query QueryParams) ([]T, error)
	Get(id uint) (T, error)
	Create(entity T) (T, error)
	CreateMany(entities []T) ([]T, error)
	Upsert(id uint, entity T) (T, error)
	Update(id uint, entity T) (T, error)
	Delete(id uint) error
}
