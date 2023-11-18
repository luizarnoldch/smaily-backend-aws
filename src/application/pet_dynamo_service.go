package application

import (
    "main/src/domain"
    "main/src/infrastructure"
)

type PetDynamoService struct {
    repo infrastructure.PetRepository
}

func NewPetDynamoService(repo infrastructure.PetRepository) *PetDynamoService {
    return &PetDynamoService{
        repo: repo,
    }
}

func (service *PetDynamoService) GetAllPets() ([]domain.PetResponse, error) {
    response, err := service.repo.GetAllPets()
    if err != nil {
        return []domain.PetResponse{}, err
    }
    return response, nil
}

func (service *PetDynamoService) CreatePet(req *domain.PetRequest) (*domain.PetResponse, error) {
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