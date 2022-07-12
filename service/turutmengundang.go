package service

import (
	"undangan_online_api/entity"
	"undangan_online_api/input"
	"undangan_online_api/repository"
)

type TurutMengundangService interface {
	TurutMengundangServiceGetAll() ([]entity.TurutMengundang, error)
	TurutMengundangServiceGetByID(inputID input.InputIDTurutMengundang) (entity.TurutMengundang, error)
	TurutMengundangServiceCreate(input input.TurutMengundangInput) (entity.TurutMengundang, error)
	TurutMengundangServiceUpdate(inputID input.InputIDTurutMengundang, inputData input.TurutMengundangInput) (entity.TurutMengundang, error)
	TurutMengundangServiceDeleteByID(inputID input.InputIDTurutMengundang) (bool, error)
}
type turutmengundangService struct {
	repository repository.TurutMengundangRepository
}

func NewTurutMengundangService(repository repository.TurutMengundangRepository) *turutmengundangService {
	return &turutmengundangService{repository}
}
func (s *turutmengundangService) TurutMengundangServiceCreate(input input.TurutMengundangInput) (entity.TurutMengundang, error) {
	turutmengundang := entity.TurutMengundang{}
	turutmengundang.IdUndangan = input.IdUndangan
	turutmengundang.Pihak = input.Pihak
	turutmengundang.Nama = input.Nama
	turutmengundang.Hubungan = input.Hubungan
	newTurutMengundang, err := s.repository.SaveTurutMengundang(turutmengundang)
	if err != nil {
		return newTurutMengundang, err
	}
	return newTurutMengundang, nil
}
func (s *turutmengundangService) TurutMengundangServiceUpdate(inputID input.InputIDTurutMengundang, inputData input.TurutMengundangInput) (entity.TurutMengundang, error) {
	turutmengundang, err := s.repository.FindByIDTurutMengundang(inputID.ID)
	if err != nil {
		return turutmengundang, err
	}
	turutmengundang.IdUndangan = inputData.IdUndangan
	turutmengundang.Pihak = inputData.Pihak
	turutmengundang.Nama = inputData.Nama
	turutmengundang.Hubungan = inputData.Hubungan

	updatedTurutMengundang, err := s.repository.UpdateTurutMengundang(turutmengundang)

	if err != nil {
		return updatedTurutMengundang, err
	}
	return updatedTurutMengundang, nil
}
func (s *turutmengundangService) TurutMengundangServiceGetByID(inputID input.InputIDTurutMengundang) (entity.TurutMengundang, error) {
	turutmengundang, err := s.repository.FindByIDTurutMengundang(inputID.ID)
	if err != nil {
		return turutmengundang, err
	}
	return turutmengundang, nil
}
func (s *turutmengundangService) TurutMengundangServiceGetAll() ([]entity.TurutMengundang, error) {
	turutmengundangs, err := s.repository.FindAllTurutMengundang()
	if err != nil {
		return turutmengundangs, err
	}
	return turutmengundangs, nil
}
func (s *turutmengundangService) TurutMengundangServiceDeleteByID(inputID input.InputIDTurutMengundang) (bool, error) {
	_, err := s.repository.FindByIDTurutMengundang(inputID.ID)
	if err != nil {
		return false, err
	}
	_, err = s.repository.DeleteByIDTurutMengundang(inputID.ID)
	if err != nil {
		return false, err
	}
	return true, nil
}

//Generated by Micagen at 12 Juli 2022
