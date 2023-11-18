package domain

import "errors"

type PetRequest struct {
	Name          string `json:"name_pet"`
	Birth_Date    string `json:"birth_date_pet"`
	Visit_Counter int    `json:"visit_counter_pet"`
}

type Pet struct {
	ID            string `dynamodbav:"id_pet"`
	Name          string `dynamodbav:"name_pet"`
	Birth_Date    string `dynamodbav:"birth_date_pet"`
	Visit_Counter int    `dynamodbav:"visit_counter_pet"`
}

type PetResponse struct {
	ID            string `json:"id_pet"`
	Name          string `json:"name_pet"`
	Birth_Date    string `json:"birth_date_pet"`
	Visit_Counter int    `json:"visit_counter_pet"`
	Message       string `json:"message"`
}

func (p *PetRequest) ToPet() (Pet, error) {
    if p.Name == "" {
        return Pet{}, errors.New("pet name is required")
    }

    if p.Birth_Date == "" {
        return Pet{}, errors.New("pet birth date is required")
    }

    return Pet{
        Name:          p.Name,
        Birth_Date:    p.Birth_Date,
        Visit_Counter: p.Visit_Counter,
    }, nil
}

func (p *Pet) ToPetResponse() (PetResponse, error) {
	return PetResponse{
		ID:            p.ID,
		Name:          p.Name,
		Birth_Date:    p.Birth_Date,
		Visit_Counter: p.Visit_Counter,
	}, nil
}
