package infrastructure

import "gorm.io/gorm"

type SampleRepository struct {
	db *gorm.DB
}

func NewSampleRepository(db *gorm.DB) *SampleRepository {
	return &SampleRepository{
		db: db,
	}
}
