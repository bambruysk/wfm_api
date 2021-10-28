package repo

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
	"testing"
	"wfm_api/internal/model"
)

func (s *TestRepoSuite) AddTestsGoods(N int) {
	var goods []model.Good
	for i := 0; i < N; i++ {
		goods = append(goods, model.Good{
			Name:        fmt.Sprintf("Good %d", i),
			Description: fmt.Sprintf("Good %d", i),
			Count:       uint(i),
		})
	}
	s.DB.Create(&goods)
	s.goods = goods
}

func CompareGood(a, b model.Good) bool {
	return a.Name == b.Name && a.Description == b.Description && a.Count == b.Count
}

func CompareGoods(a, b []model.Good) bool {
	for i, v := range a {
		if !CompareGood(v, b[i]) {
			return false
		}
	}
	return true
}

func TestRepo_ChangeGoodCount(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx       context.Context
		id        uint
		diffCount int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repo{
				db: tt.fields.db,
			}
			if err := r.ChangeGoodCount(tt.args.ctx, tt.args.id, tt.args.diffCount); (err != nil) != tt.wantErr {
				t.Errorf("ChangeGoodCount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func (s *TestRepoSuite) TestRepo_CreateGood() {

	type args struct {
		ctx         context.Context
		name        string
		description string
		count       uint
	}
	tests := []struct {
		name    string
		args    args
		want    model.Good
		wantErr bool
	}{
		{
			name: "Positive 1",
			args: args{
				ctx:         context.Background(),
				name:        "Test Good 1",
				description: "Test Good Description",
				count:       100,
			},
			want: model.Good{
				Name:        "Test Good 1",
				Description: "Test Good Description",
				Count:       100,
			},
			wantErr: false,
		},
		{
			name: "Exist 1",
			args: args{
				ctx:         context.Background(),
				name:        s.goods[0].Name,
				description: s.goods[0].Description,
				count:       s.goods[0].Count,
			},
			want:    s.goods[0],
			wantErr: true,
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			got, err := s.repo.CreateGood(tt.args.ctx, tt.args.name, tt.args.description, tt.args.count)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateGood() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !CompareGood(got, tt.want) {
				t.Errorf("CreateGood() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func (s *TestRepoSuite) TestRepo_GetAllGoods() {
	type args struct {
		ctx    context.Context
		limit  int
		offset int
	}
	tests := []struct {
		name    string
		args    args
		want    []model.Good
		wantErr bool
	}{
		{
			name: "Positive test 1",
			args: args{
				ctx:    context.Background(),
				limit:  3,
				offset: 3,
			},
			want:    s.goods[3:6],
			wantErr: false,
		},

		{
			name: "Negativeive test 2",
			args: args{
				ctx:    context.Background(),
				limit:  -10,
				offset: -10,
			},
			want:    s.goods,
			wantErr: true,
		},
		{
			name: "Negativeive test 2",
			args: args{
				ctx:    context.Background(),
				limit:  0,
				offset: 0,
			},
			want:    s.goods,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			got, err := s.repo.GetAllGoods(tt.args.ctx, tt.args.limit, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllGoods() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !CompareGoods(got, tt.want) {
				t.Errorf("GetAllGoods() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func (s *TestRepoSuite) TestRepo_GetByName() {

	type args struct {
		ctx  context.Context
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    model.Good
		wantErr bool
	}{
		{
			name: "Positive case 1",
			args: args{
				context.Background(),
				s.goods[0].Name,
			},
			want:    s.goods[0],
			wantErr: false,
		},
		{
			name: "N case 2",
			args: args{
				context.Background(),
				"some name",
			},
			want:    model.Good{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			got, err := s.repo.GetByName(tt.args.ctx, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !CompareGood(got, tt.want) {
				t.Errorf("GetByName() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func (s *TestRepoSuite) TestRepo_UpdateGoodCount() {
	type args struct {
		ctx   context.Context
		id    uint
		count uint
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Positive test",
			args: args{
				ctx:   context.Background(),
				id:    1,
				count: 1000,
			},
			wantErr: false,
		},
		{
			name: "Negative test",
			args: args{
				ctx:   context.Background(),
				id:    1000,
				count: 1000,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if err := s.repo.UpdateGoodCount(tt.args.ctx, tt.args.id, tt.args.count); (err != nil) != tt.wantErr {
				t.Errorf("UpdateGoodCount() error = %v, wantErr %v", err, tt.wantErr)
			}
			// check
			if !tt.wantErr {
				var good model.Good
				res := s.DB.First(&good, tt.args.id)
				require.NoError(t, res.Error, " not found with id")
				require.Equal(t, tt.args.count, good.Count)
			}
		})
	}
}
