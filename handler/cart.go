package handler

import (
	"context"
	"git.imooc.com/cap1573/cart/common"
	"git.imooc.com/cap1573/cart/domain/model"
	"git.imooc.com/cap1573/cart/domain/service"
	cart "git.imooc.com/cap1573/cart/proto/cart"
	"strconv"
)

type Cart struct {
	CartDataService service.ICartDataService
}

func (h *Cart) AddCart(ctx context.Context, request *cart.CartInfo, response *cart.ResponseAdd) (err error) {
	cart := &model.Cart{}
	if err := common.SwapTo(request, cart); err != nil {
		return err
	}
	response.CartId, err = h.CartDataService.AddCart(cart)
	return err
}

func (h *Cart) CleanCart(ctx context.Context, request *cart.Clean, response *cart.Response) (err error) {
	if err := h.CartDataService.CleanCart(request.UserId); err != nil {
		return err
	}
	response.Msg = "cart cleaned"
	return nil
}

func (h *Cart) Incr(ctx context.Context, request *cart.Item, response *cart.Response) (err error) {
	if err := h.CartDataService.IncrNum(request.Id, request.ChangeNum); err != nil {
		return err
	}
	response.Msg = "cart increased " + strconv.FormatInt(request.ChangeNum, 10) + " product"
	return nil
}

func (h *Cart) Decr(ctx context.Context, request *cart.Item, response *cart.Response) (err error) {
	if err := h.CartDataService.DecrNum(request.Id, request.ChangeNum); err != nil {
		return err
	}
	response.Msg = "cart decreased " + strconv.FormatInt(request.ChangeNum, 10) + " product"
	return nil
}

func (h *Cart) DeleteItemByID(ctx context.Context, request *cart.CartID, response *cart.Response) (err error) {
	if err := h.CartDataService.DeleteCart(request.Id); err != nil {
		return err
	}
	response.Msg = "product deleted"
	return nil
}

func (h *Cart) GetAll(ctx context.Context, request *cart.CartFindAll, response *cart.CartAll) (err error) {
	cartAll, err := h.CartDataService.FindAllCart(request.UserId)
	if err != nil {
		return err
	}

	for _, v := range cartAll {
		cart := &cart.CartInfo{}
		if err := common.SwapTo(v, cart); err != nil {
			return err
		}
		response.CartInfo = append(response.CartInfo, cart)
	}
	return nil
}
