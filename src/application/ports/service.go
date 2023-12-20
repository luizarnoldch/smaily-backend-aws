package ports

import "main/src/domain/model"

type PetService interface {
	GetAllPets()([]model.PetResponse, error)
	CreatePet(*model.PetRequest) (*model.PetResponse, error)
	DeletePet(petId string) (*model.PetResponse, error)
}