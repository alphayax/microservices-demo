package reporitory

import (
	"cart-service/model"
	"context"
	"github.com/redis/go-redis/v9"
)

func GetCart(ctx context.Context, cartId string) (cart *model.Cart, err error) {
	cartRaw, err := client.Get(ctx, cartId).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, err
	}

	cart, err = model.DecodeCart(cartRaw)
	if err != nil {
		return nil, err
	}

	return cart, nil
}

func UpdateCart(ctx context.Context, cartId string, items []string) error {
	cart, err := GetCart(ctx, cartId)
	if err != nil {
		return err
	}

	if cart == nil {
		cart = &model.Cart{
			Id:    cartId,
			Items: []string{},
		}
	}

	cart.Items = items

	cartRaw, err := model.EncodeCart(cart)
	if err != nil {
		return err
	}

	err = client.Set(ctx, cartId, cartRaw, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func DeleteCart(ctx context.Context, cartId string) error {
	err := client.Del(ctx, cartId).Err()
	if err != nil {
		return err
	}

	return nil
}

/*
// Future implementation
func AddToCart(ctx context.Context, cartId string, itemId string) error {
	cart, err := GetCart(ctx, cartId)
	if err != nil {
		return err
	}

	if cart == nil {
		cart = &model.Cart{
			Id:    cartId,
			Items: []string{},
		}
	}

	cart.Items = append(cart.Items, itemId)

	cartRaw, err := model.EncodeCart(cart)
	if err != nil {
		return err
	}

	err = client.Set(ctx, cartId, cartRaw, 0).Err()
	if err != nil {
		return err
	}

	return nil
}
*/
