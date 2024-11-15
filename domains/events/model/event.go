package model

import (
	relation "gohub/domains/shares/model"
	modelUser "gohub/domains/users/model"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	ID 						string 							`json:"id" gorm:"unique;not null;index;primary_key"`
	UserId 					string 							`json:"userId" gorm:"not null"`
	User                    *modelUser.User                 `json:"user"`
	Name 					string 							`json:"name" gorm:"not null"`
	Description 			string 							`json:"description"`
	CoverImageUrl 			string 							`json:"coverImageUrl"`
	CoverImageFileName 		string 							`json:"coverImageFileName"`
	StartTime 				*time.Time						`json:"startDate" gorm:"not null"`
	EndTime					*time.Time						`json:"endDate" gorm:"not null"`
	Location 				string 							`json:"location" gorm:"not null"`
	LocationPath 			string 							`json:"locationPath" gorm:"not null"`
	EventCycleType			string							`json:"eventCycleType"`
	EventPaymentType		string							`json:"eventPaymentType"`
	IsPrivate 				bool 							`json:"isPrivate" gorm:"default:0"`
	Status                  string		 					`json:"status"`
	EventSubImages          []*EventSubImage 				`json:"eventSubImages"`
	Categories              []*relation.EventCategory		`json:"categories" gorm:"many2many:event_categories;"`
	Reasons 				[]*Reason 						`json:"reasons"`
	Coupons                 []*relation.EventCoupons		`json:"coupons" gorm:"many2many:event_coupons;"` 
	TicketTypes 			[]*TicketType 					`json:"ticketTypes"`
	UserFavourite           []*relation.EventFavourite		`json:"userFavourite" gorm:"many2many:event_favourites;"`
	UserInviter			 	[]*relation.Invitation			`json:"userInviter" gorm:"many2many:invitations;"`
}

func (e *Event) BeforeCreate(db *gorm.DB) (err error) {
	e.ID = uuid.New().String()
	return nil
}

func (Event) TableName() string {
	return "events"
}