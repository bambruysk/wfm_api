package repo

import (
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"reflect"
	"testing"
	"wfm_api/internal/model"
)


type TestRepoSuite struct {
	suite.Suite
	DB    *gorm.DB
	repo  *Repo
	goods []model.Good
}

func (s *TestRepoSuite) AfterTest(_, _ string) {
	//	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInit(t *testing.T) {
	suite.Run(t, new(TestRepoSuite))
}


func (s *TestRepoSuite) SetupTest() {
	var err error
	configPath := "./../../config/config_test.yaml"
	dsn := GetDBConfig(configPath)
	s.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	require.NoError(s.T(), err)
	s.repo = New(s.DB)
	// clean table
	err = s.repo.db.Migrator().DropTable(&model.Good{})
	require.NoError(s.T(), err)

	s.repo.db.AutoMigrate(&model.Good{})
	require.NoError(s.T(), err)

	s.AddTestsGoods(100)
}




func TestNew(t *testing.T) {
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want *Repo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
