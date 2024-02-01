package gg

import "gorm.io/gorm"

type PGStore[T Entities] struct {
	db *gorm.DB
}

func NewPGStore[T Entities](db *gorm.DB) IRepository[T] {
	return &PGStore[T]{
		db: db,
	}
}

func (s *PGStore[T]) Count(search SearchParams) (int64, error) {
	var entity T
	var count int64

	base := s.db

	if search.IsSearched() {
		keyword := search.GetKeyword()
		base = base.Where("name ILIKE ?", keyword)
	}

	err := base.Model(&entity).Count(&count).Error
	return count, err
}

func (s *PGStore[T]) GetAll(query QueryParams) ([]T, error) {
	var entities []T

	base := s.db

	if query.IsPaginated() {
		offset := query.Offset()
		limit := query.Limit()
		base = base.Offset(offset).Limit(limit)
	}

	if query.IsSearched() {
		keyword := query.GetKeyword()
		base = base.Where("name ILIKE ?", keyword)
	}

	if query.IsSorted() {
		orderBy := query.GetOrderBy()
		base = base.Order(orderBy)
	}

	err := base.Find(&entities).Error
	return entities, err
}

func (s *PGStore[T]) Get(id uint) (T, error) {
	var entity T
	err := s.db.First(&entity, id).Error
	return entity, HandleRetrievalError(err)
}

func (s *PGStore[T]) Create(entity T) (T, error) {
	err := s.db.Create(&entity).Error
	return entity, HandleCreationError(err)
}

func (s *PGStore[T]) CreateMany(entities []T) ([]T, error) {
	err := s.db.Transaction(func(tx *gorm.DB) error {
		return s.db.Create(&entities).Error
	})
	return entities, HandleCreationError(err)
}

func (s *PGStore[T]) Upsert(id uint, entity T) (T, error) {
	err := s.db.Save(&entity).Error
	return entity, HandleCreationError(err)
}

func (s *PGStore[T]) Update(id uint, entity T) (T, error) {
	err := s.db.Model(&entity).Where("id = ?", id).Updates(entity).Error
	if err != nil {
		return entity, HandleCreationError(err)
	}

	return s.Get(id)
}

func (s *PGStore[T]) Delete(id uint) error {
	var entity T
	result := s.db.Delete(&entity, id)
	return HandleDeletionError(result.Error)
}
