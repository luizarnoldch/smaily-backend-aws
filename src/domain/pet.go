package domain

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

func (p *PetRequest) ToPet() Pet {
	return Pet{
		ID:            p.Name,
		Name:          p.Name,
		Birth_Date:    p.Birth_Date,
		Visit_Counter: p.Visit_Counter,
	}
}

func (p *Pet) ToPetResponse() PetResponse {
	return PetResponse{
		ID:            p.ID,
		Name:          p.Name,
		Birth_Date:    p.Birth_Date,
		Visit_Counter: p.Visit_Counter,
	}
}
