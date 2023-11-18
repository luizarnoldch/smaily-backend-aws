package pets_test

import (
	"main/src/application"
	"main/src/domain"
	"main/tests/mocks"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type DynamoPetRepositorySuite struct {
    suite.Suite
}

// func (suite *DynamoPetRepositorySuite) SetupTest() {
// }

func (suite *DynamoPetRepositorySuite) TestCreatePetSuccessful() {
    mockPetRepository := new(mocks.PetRepository)
    testPetRequest := &domain.PetRequest{
        Name: "Test Pet",
        Birth_Date: "2021-01-01",
        Visit_Counter: 1,
    }

    mockedPetResponse := &domain.PetResponse{
        ID: "test-uuid",
        Name: testPetRequest.Name,
        Birth_Date: testPetRequest.Birth_Date,
        Visit_Counter: testPetRequest.Visit_Counter,
    }
    mockPetRepository.On("CreatePet", mock.AnythingOfType("*domain.Pet")).Return(mockedPetResponse, nil)

    petService := application.NewPetDynamoService(mockPetRepository)

    petResponse, err := petService.CreatePet(testPetRequest)
    suite.NoError(err)
    suite.NotEmpty(petResponse.ID)
    suite.Equal(testPetRequest.Name, petResponse.Name)
    suite.Equal(testPetRequest.Birth_Date, petResponse.Birth_Date)
    suite.Equal(testPetRequest.Visit_Counter, petResponse.Visit_Counter)

    mockPetRepository.AssertExpectations(suite.T())
}

func (suite *DynamoPetRepositorySuite) TestCreatePetWithoutName() {
    mockPetRepository := new(mocks.PetRepository)
    testPetRequest := &domain.PetRequest{
        Name: "",
        Birth_Date: "2021-01-01",
        Visit_Counter: 1,
    }

    petService := application.NewPetDynamoService(mockPetRepository)

    petResponse, err := petService.CreatePet(testPetRequest)

    suite.Nil(petResponse)
    suite.Error(err)
    suite.Equal("pet name is required", err.Error())
}

func (suite *DynamoPetRepositorySuite) TestCreatePetWithoutBirthDate() {
    mockPetRepository := new(mocks.PetRepository)
    testPetRequest := &domain.PetRequest{
        Name: "Test Pet",
        Birth_Date: "",
        Visit_Counter: 1,
    }

    petService := application.NewPetDynamoService(mockPetRepository)

    petResponse, err := petService.CreatePet(testPetRequest)

    suite.Nil(petResponse)
    suite.Error(err)
    suite.Equal("pet birth date is required", err.Error())
}

func TestDynamoPetRepositorySuite(t *testing.T) {
    suite.Run(t, new(DynamoPetRepositorySuite))
}