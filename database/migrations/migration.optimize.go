package migrations

import (
	"fmt"
	"gohub/database"

	"github.com/QuocAnh189/GoBin/logger"

	categoryModel "gohub/domains/categories/model"
	commandModel "gohub/domains/commands/model"
	conversationModel "gohub/domains/conversations/model"
	couponModel "gohub/domains/coupons/model"
	eventModel "gohub/domains/events/model"
	functionModel "gohub/domains/functions/model"
	paymentModel "gohub/domains/payments/model"
	permissionModel "gohub/domains/permissions/model"
	reviewModel "gohub/domains/reviews/model"
	roleModel "gohub/domains/roles/model"
	relationModel "gohub/domains/shares/model"
	ticketModel "gohub/domains/tickets/model"
	userModel "gohub/domains/users/model"
)


func AutoMigrateOptimize(db *database.Database) error {
	tables := []interface{}{
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
	}

	for _, table := range tables {
		tableName, tableExists := db.HasTable(table)
		if !tableExists {
			err := db.AutoMigrate(table)
			if err != nil {
				return err
			}
			logger.Info(fmt.Sprintf("Table %s migrated successfully", tableName))
		} else {
			logger.Info(fmt.Sprintf("Table %s already exists, skipping migration", tableName))
		}
	}

	return nil
}
