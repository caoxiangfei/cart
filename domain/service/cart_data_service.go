package service

import (
	"git.imooc.com/cap1573/cart/domain/model"
	"git.imooc.com/cap1573/cart/domain/repository"
)

type ICartDataService interface {
	AddCart(*model.Cart) (int64, error)
	DeleteCart(int64) error
	UpdateCart(*model.Cart) error
	FindCartByID(int64) (*model.Cart, error)
	FindAllCart(int642 int64) ([]model.Cart, error)

	CleanCart(int642 int64) error
	IncrNum(int642 int64, int643 int64) error
	DecrNum(int642 int64, int643 int64) error
}

//创建
func NewCartDataService(cartRepository repository.ICartRepository) ICartDataService {
	return &CartDataService{cartRepository}
}

type CartDataService struct {
	CartRepository repository.ICartRepository
}

//插入
func (u *CartDataService) AddCart(cart *model.Cart) (int64, error) {
	return u.CartRepository.CreateCart(cart)
}

//删除
func (u *CartDataService) DeleteCart(cartID int64) error {
	return u.CartRepository.DeleteCartByID(cartID)
}

//更新
func (u *CartDataService) UpdateCart(cart *model.Cart) error {
	return u.CartRepository.UpdateCart(cart)
}

//查找
func (u *CartDataService) FindCartByID(cartID int64) (*model.Cart, error) {
	return u.CartRepository.FindCartByID(cartID)
}

//查找
func (u *CartDataService) FindAllCart(userID int64) ([]model.Cart, error) {
	return u.CartRepository.FindAll(userID)
}

func (u *CartDataService) CleanCart(userID int64) error {
	return u.CartRepository.CleanCart(userID)
}

func (u *CartDataService) IncrNum(userID int64, num int64) error {
	return u.CartRepository.IncrNum(userID, num)
}

func (u *CartDataService) DecrNum(userID int64, num int64) error {
	return u.CartRepository.IncrNum(userID, num)
}
