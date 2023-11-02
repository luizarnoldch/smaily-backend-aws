package pets_test

import (
	"testing"

	"main/src/domain"
	"main/src/infrastructure"
	"main/tests/mocks"

	"github.com/stretchr/testify/suite"
)

type PetsUnitTestSuite struct {
	suite.Suite
	petsMock *mocks.PetRepository
}

func TestUnitTestsSuite(t *testing.T) {
	suite.Run(t, new(PetsUnitTestSuite))
}

func (suite *PetsUnitTestSuite) SetupTest() {
    suite.petsMock = &mocks.PetRepository{}
}


func (suite *PetsUnitTestSuite) TestPetRequest_NoError() {
	request := domain.PetRequest{
		Name:          "Pet_Name",
		Birth_Date:    "1698880480955",
		Visit_Counter: 2,
	}

	suite.Equal("Pet_Name", request.Name)

	suite.Equal(13, len(request.Birth_Date))
	suite.Equal("1698880480955", request.Birth_Date)

	suite.Equal(2, request.Visit_Counter)
}

func (suite *PetsUnitTestSuite) TestDynamoPetRepositoryInheritPetRepository() {
	var repo infrastructure.PetRepository = &infrastructure.DynamoDBPetRepository{}
	_, ok := interface{}(repo).(*infrastructure.DynamoDBPetRepository)
	if !ok {
		suite.Fail("Not a DynamoDBPetRepository")
	}
}

func (suite *PetsUnitTestSuite) TestGetAllPets() {
	mockOutput := []domain.PetResponse{
		{
			ID:            "Pet_Name",
			Name:          "Pet_Name",
			Birth_Date:    "1698880480955",
			Visit_Counter: 2,
		},
		{
			ID:            "Pet_Name2",
			Name:          "Pet_Name2",
			Birth_Date:    "1698880480956",
			Visit_Counter: 3,
		},
	}

	suite.petsMock.On("GetAllPets").Return(mockOutput, nil)

	response, err := suite.petsMock.GetAllPets()
	suite.Nil(err)
	suite.Equal(len(mockOutput), len(response))

	suite.petsMock.AssertCalled(suite.T(), "GetAllPets")
}
