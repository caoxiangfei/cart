package repository

import (
	"errors"
	"github.com/caoxiangfei/cart/domain/model"
	"github.com/jinzhu/gorm"
)

type ICartRepository interface {
	InitTable() error
	FindCartByID(int64) (*model.Cart, error)
	CreateCart(*model.Cart) (int64, error)
	DeleteCartByID(int64) error
	UpdateCart(*model.Cart) error
	FindAll(int642 int64) ([]model.Cart, error)

	CleanCart(int642 int64) error
	IncrNum(int642 int64, int643 int64) error
	DecrNum(int642 int64, int643 int64) error
}

//创建cartRepository
func NewCartRepository(db *gorm.DB) ICartRepository {
	return &CartRepository{mysqlDb: db}
}

type CartRepository struct {
	mysqlDb *gorm.DB
}

//初始化表
func (u *CartRepository) InitTable() error {
	return u.mysqlDb.CreateTable(&model.Cart{}).Error
}

//根据ID查找Cart信息
func (u *CartRepository) FindCartByID(cartID int64) (cart *model.Cart, err error) {
	cart = &model.Cart{}
	return cart, u.mysqlDb.First(cart, cartID).Error
}

//创建Cart信息
func (u *CartRepository) CreateCart(cart *model.Cart) (int64, error) {
	db := u.mysqlDb.FirstOrCreate(cart, model.Cart{ProductID: cart.ProductID, SizeID: cart.SizeID, UserID: cart.UserID})
	if db.Error != nil {
		return 0, db.Error
	}
	if db.RowsAffected == 0 {
		return 0, errors.New("insert failed")
	}
	return cart.ID, nil
}

//根据ID删除Cart信息
func (u *CartRepository) DeleteCartByID(cartID int64) error {
	return u.mysqlDb.Where("id = ?", cartID).Delete(&model.Cart{}).Error
}

//更新Cart信息
func (u *CartRepository) UpdateCart(cart *model.Cart) error {
	return u.mysqlDb.Model(cart).Update(cart).Error
}

//获取结果集
func (u *CartRepository) FindAll(userID int64) (cartAll []model.Cart, err error) {
	return cartAll, u.mysqlDb.Where("user_id = ?", userID).Find(&cartAll).Error
}

//clean the cart by userid
func (u *CartRepository) CleanCart(userID int64) error {
	return u.mysqlDb.Where("user_id = ?", userID).Delete(&model.Cart{}).Error
}

//increase the number of product
func (u *CartRepository) IncrNum(cartID int64, num int64) error {
	cart := &model.Cart{ID: cartID}
	return u.mysqlDb.Model(cart).UpdateColumn("num", gorm.Expr("num + ?", num)).Error
}

//decrease
func (u *CartRepository) DecrNum(cartID int64, num int64) error {
	cart := &model.Cart{ID: cartID}
	db := u.mysqlDb.Model(cart).Where("num >= ?", num).UpdateColumn("num", gorm.Expr("num - ?", num))
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return errors.New("decrease failed")
	}
	return nil
}
