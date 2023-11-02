package application

import "main/src/domain"

type PetService interface {
	GetAllPets()([]domain.PetResponse, error)
	CreatePet(*domain.PetRequest) (*domain.PetResponse, error)
}