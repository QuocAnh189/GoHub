package migrations

import (
	"gohub/database"

	"github.com/QuocAnh189/GoBin/logger"

	categoryModel "gohub/domains/categories/model"
	commandModel "gohub/domains/commands/model"
	conversationModel "gohub/domains/conversations/model"
	couponModel "gohub/domains/coupons/model"
	eventModel "gohub/domains/events/model"
	functionModel "gohub/domains/functions/model"
	labelsModel "gohub/domains/labels/model"
	paymentModel "gohub/domains/payments/model"
	permissionModel "gohub/domains/permissions/model"
	reviewModel "gohub/domains/reviews/model"
	roleModel "gohub/domains/roles/model"
	relationModel "gohub/domains/shares/model"
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
		&reviewModel.Review{},
		&couponModel.Coupon{},	
		&ticketModel.Ticket{},
		&labelsModel.Label{},
		&labelsModel.LabelInEvent{},
		&labelsModel.LabelInUser{},
		&paymentModel.Payment{},
		&paymentModel.PaymentLine{},
		&paymentModel.PaymentMethod{},
		&relationModel.CommandInFunction{},
		&relationModel.EventCategory{},
		&relationModel.EventCoupons{},
		&relationModel.EventFavourite{},
		&relationModel.Invitation{},
		&relationModel.UserPayment{},
		&relationModel.UserRole{},
	)

	if err != nil {
        return err
    }

	logger.Info("Migration database successfully")
	return nil
}