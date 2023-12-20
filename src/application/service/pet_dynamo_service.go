package service

import (
	"main/src/domain/model"
	"main/src/domain/repository"
)

type PetDynamoService struct {
	repo repository.PetRepository
}

func NewPetDynamoService(repo repository.PetRepository) *PetDynamoService {
	return &PetDynamoService{
		repo: repo,
	}
}

func (service *PetDynamoService) GetAllPets() ([]model.PetResponse, error) {
	response, err := service.repo.GetAllPets()
	if err != nil {
		return []model.PetResponse{}, err
	}
	return response, nil
}

func (service *PetDynamoService) CreatePet(req *model.PetRequest) (*model.PetResponse, error) {
	pet, err := req.ToPet()
	if err != nil {
		return nil, err
	}
	response, err := service.repo.CreatePet(&pet)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (service *PetDynamoService) DeletePet(petId string) (*model.PetResponse, error) {
	return service.repo.DeletePet(petId)
}
