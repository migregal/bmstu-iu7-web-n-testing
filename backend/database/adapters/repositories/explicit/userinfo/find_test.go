//go:build unit
// +build unit

package userinfo

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"neural_storage/cube/core/entities/user"
	"neural_storage/cube/core/ports/repositories"
	"neural_storage/database/core/entities/user_info"
	"neural_storage/database/test/mock/utils"
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
	id := "test"
	expected := []user.Info{*user.NewInfo(id, "", "", "", "", 0, time.Time{})}
	expectedTotal := int64(10)
	res := []user_info.UserInfo{{ID: id}}

	filter := repositories.UserInfoFilter{UserIds: make([]string, 1), Limit: 10}

	s.SqlMock.
		ExpectQuery(`^SELECT \* FROM "users_info" WHERE id in .* LIMIT 10`).
		WillReturnRows(utils.MockRows(res))

	s.SqlMock.
		ExpectQuery(`^SELECT count\(\*\) FROM "users_info"`).
		WillReturnRows(sqlmock.NewRows([]string{"count(*)"}).AddRow(10))

	info, total, err := s.repo.Find(filter)

	require.NoError(s.T(), err)
	require.Equal(s.T(), expectedTotal, total)
	require.Equal(s.T(), expected, info)
}

func TestFindSuite(t *testing.T) {
	suite.Run(t, new(FindSuite))
}
