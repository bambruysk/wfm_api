package repo

import (
	"context"
	"gorm.io/gorm"
	"reflect"
	"testing"
	"wfm_api/internal/model"
)

func TestRepo_AddGoodToCart(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx   context.Context
		cart  model.Cart
		good  model.Good
		count uint
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
			if err := r.AddGoodToCart(tt.args.ctx, tt.args.cart, tt.args.good, tt.args.count); (err != nil) != tt.wantErr {
				t.Errorf("AddGoodToCart() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepo_CreateCart(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx  context.Context
		user model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Cart
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repo{
				db: tt.fields.db,
			}
			got, err := r.CreateCart(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateCart() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateCart() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepo_DeleteCart(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx  context.Context
		cart *model.Cart
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
			if err := r.DeleteCart(tt.args.ctx, tt.args.cart); (err != nil) != tt.wantErr {
				t.Errorf("DeleteCart() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepo_GetAllCarts(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx    context.Context
		user   *model.User
		limit  int
		offset int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []model.Cart
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repo{
				db: tt.fields.db,
			}
			got, err := r.GetAllCarts(tt.args.ctx, tt.args.user, tt.args.limit, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllCarts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllCarts() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepo_GetCartForUser(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx  context.Context
		user *model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.Cart
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repo{
				db: tt.fields.db,
			}
			got, err := r.GetCartForUser(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCartForUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCartForUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepo_RemoveGoodFromCart(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx  context.Context
		cart *model.Cart
		good *model.Good
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
			if err := r.RemoveGoodFromCart(tt.args.ctx, tt.args.cart, tt.args.good); (err != nil) != tt.wantErr {
				t.Errorf("RemoveGoodFromCart() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepo_RemoveNGoodFromCart(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx   context.Context
		cart  *model.Cart
		good  *model.Good
		count uint
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
			if err := r.RemoveNGoodFromCart(tt.args.ctx, tt.args.cart, tt.args.good, tt.args.count); (err != nil) != tt.wantErr {
				t.Errorf("RemoveNGoodFromCart() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
