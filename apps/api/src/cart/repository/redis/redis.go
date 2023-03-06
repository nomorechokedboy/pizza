package redis

import (
	cartDomain "api/src/cart/domain"
	"api/src/cartItem/domain"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type CartRedisRepo struct {
	Conn *redis.Client
}

func reverse[S ~[]E, E any](s S) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func (r *CartRedisRepo) CartToRedis(key string, ctx *context.Context, cart *cartDomain.Cart) error {
	carItemsString, jsonErr := json.Marshal(cart.CartItems)
	if jsonErr != nil {
		return errors.New("serialize error")
	}

	cartRedis := cartDomain.CartRedis{Total: cart.Total, UserId: cart.UserId, CartItems: string(carItemsString)}
	setResult := r.Conn.HSet(*ctx, key, cartRedis)
	if setResult.Err() != nil {
		return errors.New("unknown error")
	}

	return nil
}

func (r *CartRedisRepo) AddFirstCartItem(cart *cartDomain.Cart, req *domain.WriteCartItemRequest, key string, ctx *context.Context) (*cartDomain.Cart, error) {
	cartItem := &domain.CartItem{ProductID: req.ProductID, Quantity: req.Quantity, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	cartItems := []*domain.CartItem{cartItem}
	cart.CartItems = cartItems
	cart.Total = 1

	err := r.CartToRedis(key, ctx, cart)
	if err != nil {
		return nil, err
	}

	return cart, nil
}

func (r *CartRedisRepo) Insert(uid uint, req domain.WriteCartItemRequest) (*cartDomain.Cart, error) {
	cart, err := r.Find(uid)
	if err != nil {
		return cart, err
	}

	key := fmt.Sprintf("cart::%d", uid)
	ctx := context.Background()

	if len(cart.CartItems) == 0 && cart.Total == 0 {
		return r.AddFirstCartItem(cart, &req, key, &ctx)
	}

	var found *domain.CartItem
	for _, item := range cart.CartItems {
		if item.ProductID == req.ProductID {
			item.Quantity += req.Quantity
			found = item
		}
	}

	if found == nil {
		found = &domain.CartItem{ProductID: req.ProductID, Quantity: req.Quantity}
		cart.CartItems = append(cart.CartItems, found)
		cart.Total++
	}

	err = r.CartToRedis(key, &ctx, cart)
	if err != nil {
		return nil, err
	}

	return cart, nil
}

func (r *CartRedisRepo) Find(uid uint) (*cartDomain.Cart, error) {
	key := fmt.Sprintf("cart::%d", uid)
	ctx := context.Background()
	result := r.Conn.HGetAll(ctx, key)
	val, getErr := result.Result()
	if getErr != nil {
		return nil, errors.New("unknown error")
	}

	if len(val) == 0 {
		return &cartDomain.Cart{Total: 0, UserId: uid, CartItems: make([]*domain.CartItem, 0)}, nil
	}

	cartRedis := cartDomain.CartRedis{}
	if err := result.Scan(&cartRedis); err != nil {
		return nil, errors.New("redis scan error")
	}
	cartItems := []*domain.CartItem{}
	err := json.Unmarshal([]byte(cartRedis.CartItems), &cartItems)
	if err != nil {
		return nil, errors.New("deserialize error")
	}
	reverse(cartItems)

	return &cartDomain.Cart{Total: cartRedis.Total, UserId: cartRedis.UserId, CartItems: cartItems}, nil
}

func (r *CartRedisRepo) Delete(uid uint, ProductID uint) (*cartDomain.Cart, error) {
	cart, err := r.Find(uid)
	if err != nil {
		return cart, err
	}

	filterCartItems := []*domain.CartItem{}
	var found *domain.CartItem
	for _, cartItem := range cart.CartItems {
		if cartItem.ProductID == ProductID {
			found = cartItem
			continue
		}

		filterCartItems = append(filterCartItems, cartItem)
	}

	if found == nil {
		return nil, nil
	}
	cart.CartItems = filterCartItems
	cart.Total = uint(len(filterCartItems))

	key := fmt.Sprintf("cart::%d", uid)
	ctx := context.Background()
	carItemsString, jsonErr := json.Marshal(filterCartItems)
	if jsonErr != nil {
		return nil, errors.New("serialize error")
	}

	cartRedis := cartDomain.CartRedis{Total: cart.Total, UserId: uid, CartItems: string(carItemsString)}
	setResult := r.Conn.HSet(ctx, key, cartRedis)
	if setResult.Err() != nil {
		return nil, errors.New("unknown error")
	}

	return cart, nil
}

func (r *CartRedisRepo) Update(uid uint, req domain.WriteCartItemRequest) (*cartDomain.Cart, error) {
	cart, err := r.Find(uid)
	if err != nil {
		return cart, err
	}

	var updated *domain.CartItem
	for _, cartItem := range cart.CartItems {
		if cartItem.ProductID == req.ProductID {
			cartItem.Quantity = req.Quantity
			cartItem.UpdatedAt = time.Now()
			updated = cartItem
		}
	}

	if updated == nil {
		return nil, nil
	}

	key := fmt.Sprintf("cart::%d", uid)
	ctx := context.Background()
	err = r.CartToRedis(key, &ctx, cart)
	if err != nil {
		return nil, err
	}

	return cart, nil
}
