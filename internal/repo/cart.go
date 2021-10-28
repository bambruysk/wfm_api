package repo

import (
	"context"
	"wfm_api/internal/model"
)

//GetCartForUser get cart for user
func (r *Repo) GetCartForUser(ctx context.Context, user *model.User) (model.Cart, error) {
	var cart model.Cart
	res := r.db.WithContext(ctx).Where("user = ?", user).First(&cart)
	return cart, res.Error
}

func (r *Repo) GetAllCarts(ctx context.Context, user *model.User, limit, offset int) ([]model.Cart, error) {
	var carts []model.Cart
	res := r.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&carts)
	return carts, res.Error
}

func (r *Repo) CreateCart(ctx context.Context, user model.User) (*model.Cart, error) {
	var cart model.Cart
	cart.User = user
	return &cart, r.db.Save(&cart).Error
}

func (r *Repo) DeleteCart(ctx context.Context, cart *model.Cart) error {
	return r.db.WithContext(ctx).Delete(cart).Error
}

func (r *Repo) AddGoodToCart(ctx context.Context, cart model.Cart, good model.Good, count uint) error {
	var cartgood model.CartGood
	r.db.WithContext(ctx).Where(&model.CartGood{
		Cart: cart,
		Good: good,
	}).First(&cartgood)
	cartgood.Count += count
	return r.db.WithContext(ctx).Save(&cartgood).Error
}

func (r *Repo) RemoveNGoodFromCart(ctx context.Context, cart *model.Cart, good *model.Good, count uint) error {
	var cartgood model.CartGood

	res := r.db.WithContext(ctx).Where(&model.CartGood{
		Cart: *cart,
		Good: *good,
	}).First(&cartgood)

	if res.Error != nil {
		return res.Error
	}

	if cartgood.Count < count {
		return r.db.WithContext(ctx).Where(&cartgood).Delete(&cartgood).Error
	}
	cartgood.Count -= count
	return r.db.Save(&cartgood).Error
}

func (r *Repo) RemoveGoodFromCart(ctx context.Context, cart *model.Cart, good *model.Good) error {
	cartgood := model.CartGood{
		Cart: *cart,
		Good: *good,
	}
	return r.db.WithContext(ctx).Where(&cartgood).Delete(&cartgood).Error
}
