//go:build unit
// +build unit

package user

import (
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"neural_storage/cube/core/entities/user"
	"neural_storage/cube/core/ports/interactors"
)

type FindSuite struct {
	TestSuite
}

func (s *FindSuite) SetupTest() {
	s.TestSuite.SetupTest()
}

func (s *FindSuite) TearDownTest() {
	s.TestSuite.TearDownTest()
}

func (s *FindSuite) TestFind() {
	filter := interactors.UserInfoFilter{}
	expected := []user.Info{
		*user.NewInfo("", "", "", "", "", 0, time.Time{}),
	}
	expectedTotal := int64(10)

	s.mockedRepo.On("Find", mock.Anything).Return(expected, expectedTotal, nil)
	info, total, err := s.interactor.Find(s.ctx, filter)

	require.NoError(s.T(), err)
	require.Equal(s.T(), expected, info)
	require.Equal(s.T(), expectedTotal, total)

	require.True(s.T(), s.mockedRepo.AssertExpectations(s.T()))
}

func TestFindSuite(t *testing.T) {
	suite.Run(t, new(FindSuite))
}
