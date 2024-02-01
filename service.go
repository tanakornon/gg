package gg

import "reflect"

type Service[T Entities] struct {
	repo IRepository[T]
}

func NewService[T Entities](repo IRepository[T]) IService[T] {
	return &Service[T]{
		repo: repo,
	}
}

func (s *Service[T]) GetAll(query QueryParams) (res GetAllResponse[T], err error) {
	res.FirstIndex = query.Offset() + 1

	if res.Total, err = s.repo.Count(query.SearchParams); err != nil {
		return res, err
	}

	if res.Data, err = s.repo.GetAll(query); err != nil {
		return res, err
	}

	return res, nil
}

func (s *Service[T]) Get(id uint) (res GetResponse[T], err error) {
	res.Data, err = s.repo.Get(id)
	return res, err
}

func (s *Service[T]) Create(req T) (res GetResponse[T], err error) {
	res.Data, err = s.repo.Create(req)
	return res, err
}

func (s *Service[T]) CreateMany(req []T) (res GetResponse[[]T], err error) {
	res.Data, err = s.repo.CreateMany(req)
	return res, err
}

func (s *Service[T]) Upsert(id uint, req T) (res GetResponse[T], err error) {
	elements := reflect.ValueOf(&req).Elem()
	elements.FieldByName("ID").SetUint(uint64(id))

	res.Data, err = s.repo.Upsert(id, req)
	return res, err
}

func (s *Service[T]) Update(id uint, req T) (res GetResponse[T], err error) {
	res.Data, err = s.repo.Update(id, req)
	return res, err
}

func (s *Service[T]) Delete(id uint) error {
	return s.repo.Delete(id)
}
