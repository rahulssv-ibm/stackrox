package service

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	storeMocks "github.com/stackrox/rox/central/group/store/mocks"
	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stretchr/testify/suite"
)

func TestUserService(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(UserServiceTestSuite))
}

type UserServiceTestSuite struct {
	suite.Suite

	mockCtrl *gomock.Controller

	mockStore *storeMocks.MockStore
	ser       Service
}

func (suite *UserServiceTestSuite) SetupSuite() {
	suite.mockCtrl = gomock.NewController(suite.T())

	suite.mockStore = storeMocks.NewMockStore(suite.mockCtrl)
	suite.ser = New(suite.mockStore)
}

func (suite *UserServiceTestSuite) TestBatchUpdate() {
	update := &v1.GroupBatchUpdateRequest{
		PreviousGroups: []*storage.Group{
			{
				Props: &storage.GroupProperties{ // should be removed since the props are not in required
					AuthProviderId: "ap1",
					Key:            "k1",
					Value:          "v1",
				},
				RoleName: "r1",
			},
			{
				Props: &storage.GroupProperties{ // should be ignored since the props have the same role name in required
					AuthProviderId: "ap2",
					Key:            "k1",
					Value:          "v1",
				},
				RoleName: "r2",
			},
			{
				Props: &storage.GroupProperties{ // should get updated since the props have a new role in required
					AuthProviderId: "ap2",
					Key:            "k1",
					Value:          "v2",
				},
				RoleName: "r2",
			},
		},
		RequiredGroups: []*storage.Group{
			{
				Props: &storage.GroupProperties{ // repeat of the second group above
					AuthProviderId: "ap2",
					Key:            "k1",
					Value:          "v1",
				},
				RoleName: "r2",
			},
			{
				Props: &storage.GroupProperties{ // update to the third group above
					AuthProviderId: "ap2",
					Key:            "k1",
					Value:          "v2",
				},
				RoleName: "r3",
			},
			{
				Props: &storage.GroupProperties{ // newly added group since the props do not appear in previous.
					AuthProviderId: "ap2",
					Key:            "k2",
					Value:          "v1",
				},
				RoleName: "r3",
			},
		},
	}

	suite.mockStore.EXPECT().
		Mutate(
			[]*storage.Group{update.GetPreviousGroups()[0]},
			[]*storage.Group{update.GetRequiredGroups()[1]},
			[]*storage.Group{update.GetRequiredGroups()[2]}).
		Return(nil)

	_, err := suite.ser.BatchUpdate(context.Context(nil), update)
	suite.NoError(err, "request should not fail with valid user data")
}
