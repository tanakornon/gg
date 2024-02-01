package gg

import (
	"gorm.io/gorm"
)

type Store struct {
	DB *gorm.DB
}

func NewStore(db *gorm.DB) *Store {
	return &Store{DB: db}
}

func (s *Store) PreparePreload(joins []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		for _, join := range joins {
			db = db.Preload(join)
		}
		return db
	}
}

func (s *Store) Count(model interface{}, search SearchParams) (int64, error) {
	var count int64

	base := s.DB

	if search.IsSearched() {
		base = base.Where("name ILIKE ?", search.GetKeyword())
	}

	err := base.Model(&model).Count(&count).Error

	return count, err
}

func (s *Store) GetAll(model interface{}, query QueryParams, preload func(db *gorm.DB) *gorm.DB) error {
	base := s.DB

	if preload != nil {
		base = preload(base)
	}

	if query.IsPaginated() {
		base = base.Offset(query.Offset()).Limit(query.Limit())
	}

	if query.IsSearched() {
		base = base.Where("name ILIKE ?", query.GetKeyword())
	}

	if query.IsSorted() {
		base = base.Order(query.GetOrderBy())
	}

	err := base.Find(model).Error

	return err
}

func (s *Store) Get(id uint, model interface{}) error {
	err := s.DB.First(model, id).Error
	return err
}

func (s *Store) Create(model interface{}) error {
	err := s.DB.Create(model).Error
	return err
}

func (s *Store) CreateMany(models interface{}) error {
	err := s.DB.Transaction(func(tx *gorm.DB) error {
		return s.DB.Create(models).Error
	})
	return err
}

func (s *Store) Upsert(id uint, model interface{}) error {
	err := s.DB.Save(model).Error
	return err
}

func (s *Store) Update(id uint, model interface{}) error {
	err := s.DB.Model(model).Where("id = ?", id).Updates(model).Error
	return err
}

func (s *Store) Delete(id uint, model interface{}) error {
	err := s.DB.Delete(model, id).Error
	return err
}
