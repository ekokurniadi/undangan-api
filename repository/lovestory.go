package repository

import (
	"undangan_online_api/entity"

	"gorm.io/gorm"
)

type LoveStoryRepository interface {
	SaveLoveStory(lovestory entity.LoveStory) (entity.LoveStory, error)
	UpdateLoveStory(lovestory entity.LoveStory) (entity.LoveStory, error)
	FindByIDLoveStory(ID int) (entity.LoveStory, error)
	FindAllLoveStory(params string) ([]entity.LoveStory, error)
	DeleteByIDLoveStory(ID int) (entity.LoveStory, error)
}

type lovestoryRepository struct {
	db *gorm.DB
}

func NewLoveStoryRepository(db *gorm.DB) *lovestoryRepository {
	return &lovestoryRepository{db}
}

func (r *lovestoryRepository) SaveLoveStory(lovestory entity.LoveStory) (entity.LoveStory, error) {
	err := r.db.Create(&lovestory).Error
	if err != nil {
		return lovestory, err
	}
	return lovestory, nil

}
func (r *lovestoryRepository) FindByIDLoveStory(ID int) (entity.LoveStory, error) {
	var lovestory entity.LoveStory
	err := r.db.Where("id = ? ", ID).Find(&lovestory).Error
	if err != nil {
		return lovestory, err
	}
	return lovestory, nil

}
func (r *lovestoryRepository) UpdateLoveStory(lovestory entity.LoveStory) (entity.LoveStory, error) {
	err := r.db.Save(&lovestory).Error
	if err != nil {
		return lovestory, err
	}
	return lovestory, nil

}
func (r *lovestoryRepository) FindAllLoveStory(params string) ([]entity.LoveStory, error) {
	var lovestorys []entity.LoveStory
	err := r.db.Where("id_undangan = ? ", params).Find(&lovestorys).Error
	if err != nil {
		return lovestorys, err
	}
	return lovestorys, nil

}
func (r *lovestoryRepository) DeleteByIDLoveStory(ID int) (entity.LoveStory, error) {
	var lovestory entity.LoveStory
	err := r.db.Where("id = ? ", ID).Delete(&lovestory).Error
	if err != nil {
		return lovestory, err
	}
	return lovestory, nil

}

//Generated by Micagen at 05 Juli 2022