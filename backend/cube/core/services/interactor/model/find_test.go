//go:build unit
// +build unit

package model

import (
	"neural_storage/cube/core/entities/model"
	"neural_storage/cube/core/ports/interactors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
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
	filter := interactors.ModelInfoFilter{}
	expected := []*model.Info{
		model.NewInfo("", "", "", nil),
	}
	expectedTotal := int64(11)

	s.mockedModelInfo.On("Find", mock.Anything).Return(expected, expectedTotal, nil)
	info, total, err := s.interactor.Find(s.ctx, filter)

	require.NoError(s.T(), err)
	require.Equal(s.T(), expected, info)
	require.Equal(s.T(), expectedTotal, total)

	require.True(s.T(), s.mockedModelInfo.AssertExpectations(s.T()))
}

func TestFindSuite(t *testing.T) {
	suite.Run(t, new(FindSuite))
}
