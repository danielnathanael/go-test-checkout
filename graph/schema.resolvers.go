package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"go-test-checkout/graph/generated"
	"go-test-checkout/graph/model"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	hashedPw, err := bcrypt.GenerateFromPassword([]byte(input.Password), 14)

	if err != nil {
		return nil, err
	}

	user := model.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPw),
	}

	err = r.DB.Create(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *mutationResolver) CreateItem(ctx context.Context, input model.NewItem) (*model.Item, error) {
	item := model.Item{
		Sku:       input.Sku,
		Name:      input.Name,
		Price:     input.Price,
		Quantity:  input.Quantity,
		CreatedAt: time.Now().String(),
		UpdatedAt: time.Now().String(),
	}
	err := r.DB.Create(&item).Error

	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (r *mutationResolver) AddCart(ctx context.Context, userID string, itemSku string, quantity int) (bool, error) {

	var userCart model.Cart
	err := r.DB.Model(model.Cart{}).Where("user_id = ? AND item_id = ?", userID, itemSku).First(&userCart).Error
	if err != nil {
		return false, err
	}

	if &userCart == nil {
		//New Cart Item
		newCart := model.Cart{
			UserID:   userID,
			Quantity: quantity,
			ItemID:   itemSku,
		}
		err := r.DB.Create(&newCart).Error

		if err != nil {
			return false, err
		}

		return true, nil
	} else {
		//Existing Cart Item
		err := r.DB.Model(userCart).Update("quantity", userCart.Quantity+quantity).Error
		if err != nil {
			return false, err
		}

		return true, nil
	}
}

func (r *mutationResolver) Checkout(ctx context.Context, userID string) ([]*model.Order, error) {

	var userCart []*model.Cart

	err := r.DB.Model(&model.Cart{}).Where("user_id = ?", userID).Find(userCart).Error

	if err != nil {
		return nil, err
	}

	googleHomeCount := 0
	macBookProCount := 0
	alexaSpeakerCount := 0
	RaspberryPiBCount := 0

	for _, cart := range userCart {
		if cart.ID == "120P90" {
			googleHomeCount += 1
		} else if cart.ID == "43N23P" {
			macBookProCount += 1
		} else if cart.ID == "A304SD" {
			alexaSpeakerCount += 1
		} else if cart.ID == "234234" {
			RaspberryPiBCount += 1
		}
	}

	// Each sale of a MacBook Pro comes with a free Raspberry Pi B
	// INCOMPLETE : Assign $0 price for every new raspberry
	RaspberryPiBCount += macBookProCount

	//Buy 3 Google Homes for the price of 2
	// INCOMPLETE : Get real price of item
	googleHomePrice := float64(googleHomeCount-(googleHomeCount-(googleHomeCount%3))/3) * 49.99
	fmt.Println(googleHomePrice)

	//Buying more than 3 Alexa SPeakers will have 10% discount on all Alexa speakers
	// INCOMPLETE : Process price for every item on checkout, not final price.
	alexaSpeakerPrice := 30.00 * float64(alexaSpeakerCount)

	if alexaSpeakerCount > 3 {
		alexaSpeakerPrice = alexaSpeakerPrice * 0.9
		fmt.Println(alexaSpeakerPrice)
	}

	// INCOMPLETE : Limited time for the test
	return nil, nil
}

func (r *queryResolver) Login(ctx context.Context) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
	// Limited time for the test
}

func (r *queryResolver) Items(ctx context.Context) ([]*model.Item, error) {
	var items []*model.Item

	err := r.DB.Find(&items).Error

	if err != nil {
		return nil, err
	}

	return items, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) CheckoutItem(ctx context.Context, userID string, itemSku string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}
