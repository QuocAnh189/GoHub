package migrations

import (
	"gohub/database"

	"gohub/internal/libs/logger"

	categoryModel "gohub/domains/categories/model"
	commandModel "gohub/domains/commands/model"
	conversationModel "gohub/domains/conversations/model"
	couponModel "gohub/domains/coupons/model"
	eventModel "gohub/domains/events/model"
	expenseModel "gohub/domains/expense/model"
	functionModel "gohub/domains/functions/model"
	paymentModel "gohub/domains/payments/model"
	permissionModel "gohub/domains/permissions/model"
	reviewModel "gohub/domains/reviews/model"
	roleModel "gohub/domains/roles/model"
	ticketModel "gohub/domains/tickets/model"
	userModel "gohub/domains/users/model"
)

func AutoMigrate(db *database.Database) error {
	err := db.AutoMigrate(
		&permissionModel.Permission{},
		&commandModel.Command{},
		&functionModel.Function{},
		&categoryModel.Category{},
		&roleModel.Role{},
		&userModel.User{},
		&userModel.UserFollower{},
		&conversationModel.Conversation{},
		&conversationModel.Message{},
		&conversationModel.MessageAttachment{},
		&eventModel.Event{},
		&eventModel.EventSubImage{},
		&eventModel.Reason{},
		&eventModel.TicketType{},
		&expenseModel.Expense{},
		&expenseModel.SubExpense{},
		&reviewModel.Review{},
		&couponModel.Coupon{},
		&ticketModel.Ticket{},
		&paymentModel.Payment{},
		&paymentModel.PaymentLine{},
		&paymentModel.PaymentMethod{},
		&commandModel.CommandInFunction{},
		&eventModel.EventCategory{},
		&eventModel.EventCoupons{},
		&eventModel.EventFavourite{},
		&eventModel.Invitation{},
		&userModel.UserPayment{},
		&userModel.UserRole{},
	)

	if err != nil {
		return err
	}

	logger.Info("Migration database successfully")
	return nil
}
