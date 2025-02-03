package service

import (
	"context"
	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/checkout/session"
	"gohub/domains/payments/dto"
	"gohub/domains/payments/model"
	"gohub/domains/payments/repository"
	"gohub/internal/libs/validation"
	"gohub/pkg/paging"
)

type IPaymentService interface {
	GetTransactions(ctx context.Context, userId string, req *dto.ListTransactionReq) ([]*model.Payment, *paging.Pagination, error)
	GetOrders(ctx context.Context, userId string, req *dto.ListOrderReq) ([]*model.Payment, *paging.Pagination, error)
	CreateSession(ctx context.Context, req *dto.TicketCheckoutRequest, stripeKey string) (string, string, string, error)
	Checkout(ctx context.Context, req *dto.TicketCheckoutRequest) error
}

type PaymentService struct {
	validator   validation.Validation
	repoPayment repository.IPaymentRepository
}

func NewPaymentService(validator validation.Validation, repoPayment repository.IPaymentRepository) *PaymentService {
	return &PaymentService{
		validator:   validator,
		repoPayment: repoPayment,
	}
}

func (s *PaymentService) GetTransactions(ctx context.Context, userId string, req *dto.ListTransactionReq) ([]*model.Payment, *paging.Pagination, error) {
	transactions, pagination, err := s.repoPayment.GetTransactions(ctx, userId, req)
	if err != nil {
		return nil, nil, err
	}

	return transactions, pagination, nil
}

func (s *PaymentService) GetOrders(ctx context.Context, userId string, req *dto.ListOrderReq) ([]*model.Payment, *paging.Pagination, error) {
	orders, pagination, err := s.repoPayment.GetOrders(ctx, userId, req)
	if err != nil {
		return nil, nil, err
	}

	return orders, pagination, nil
}

func (s *PaymentService) CreateSession(ctx context.Context, req *dto.TicketCheckoutRequest, stripeKey string) (string, string, string, error) {
	stripe.Key = stripeKey

	var lineItems []*stripe.CheckoutSessionLineItemParams
	for _, item := range req.TicketItems {
		lineItems = append(lineItems, &stripe.CheckoutSessionLineItemParams{
			PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
				Currency: stripe.String("vnd"),
				ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
					Name: stripe.String(item.Name),
				},
				UnitAmount: stripe.Int64(int64(item.Price)),
			},
			Quantity: stripe.Int64(int64(item.Quantity)),
		})
	}

	params := &stripe.CheckoutSessionParams{
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
		LineItems:          lineItems,
		Mode:               stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL:         stripe.String("http://localhost:3000/payment/successfully"),
		CancelURL:          stripe.String("http://localhost:3000/payment/failure"),
		CustomerEmail:      stripe.String(req.CustomerEmail),
	}

	sessionCheckout, err := session.New(params)

	if err != nil {
		return "", "", "", err
	}

	var paymentIntentID string
	if sessionCheckout.PaymentIntent != nil {
		paymentIntentID = sessionCheckout.PaymentIntent.ID
	}

	return sessionCheckout.ID, sessionCheckout.URL, paymentIntentID, nil
}

func (s *PaymentService) Checkout(ctx context.Context, req *dto.TicketCheckoutRequest) error {
	if err := s.repoPayment.Checkout(ctx, req); err != nil {
		return err
	}

	return nil
}
