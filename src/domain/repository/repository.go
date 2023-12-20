package repository

import "main/src/domain/model"

type PetRepository interface {
	GetAllPets() ([]model.PetResponse, error)
	CreatePet(*model.Pet) (*model.PetResponse, error)
	DeletePet(petId string) (*model.PetResponse, error)
}