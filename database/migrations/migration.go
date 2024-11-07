package migrations

import (
	"gohub/database"

	"github.com/QuocAnh189/GoBin/logger"

	categoryModel "gohub/domains/categories/model"
	conversationModel "gohub/domains/conversations/model"
	eventModel "gohub/domains/events/model"
	labelsModel "gohub/domains/labels/model"
	paymentModel "gohub/domains/payments/model"
	permissionModel "gohub/domains/permissions/model"
	reviewModel "gohub/domains/reviews/model"
	ticketModel "gohub/domains/tickets/model"
	userModel "gohub/domains/users/model"
)


func AutoMigrate(db *database.Database) error {
	err := db.AutoMigrate(
		&permissionModel.Permission{},
		&permissionModel.Command{},
		&permissionModel.Function{},
		&permissionModel.CommandInFunction{},
		&categoryModel.Category{},
		&userModel.User{},
		&userModel.Role{},
        &userModel.UserRole{},
		&userModel.UserFollower{},
		&userModel.Invitation{},
		&userModel.UserPayment{},
		&conversationModel.Conversation{},
		&conversationModel.Message{},
		&eventModel.Event{},
		&eventModel.EventCategory{},
		&eventModel.EventFavourite{},
		&eventModel.EventSubImage{},
		&eventModel.Reason{},
		&eventModel.TicketType{},
		&reviewModel.Review{},
		&ticketModel.Ticket{},
		&labelsModel.Label{},
		&labelsModel.LabelInEvent{},
		&labelsModel.LabelInUser{},
		&paymentModel.Payment{},
		&paymentModel.PaymentItem{},
		&paymentModel.PaymentMethod{},
	)

	if err != nil {
        return err
    }

	logger.Info("Migration database successfully")
	return nil
}