package infrastructure

import "main/src/domain"

type PetRepository interface {
	GetAllPets() ([]domain.PetResponse, error)
	CreatePet(*domain.Pet) (*domain.PetResponse, error)
}