package model

import (
	relation "gohub/domains/shares/model"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	ID 						string 							`json:"id" gorm:"unique;not null;index;primary_key"`
	UserId 					string 							`json:"userId" gorm:"not null"`
	Name 					string 							`json:"name" gorm:"not null"`
	Description 			string 							`json:"description"`
	CoverImageUrl 			string 							`json:"coverImageUrl"`
	CoverImageFileName 		string 							`json:"coverImageFileName"`
	EventSubImages          []*EventSubImage 				`json:"eventSubImages"`
	Categories              []*relation.EventCategory		`json:"categories" gorm:"many2many:event_categories;"`
	Reasons 				[]*Reason 						`json:"reasons"`
	TicketTypes 			[]*TicketType 					`json:"ticketTypes"`
	StartTime 				*time.Time						`json:"startDate" gorm:"not null"`
	EndTime					*time.Time						`json:"endDate" gorm:"not null"`
	Location 				string 							`json:"location" gorm:"not null"`
	Promotion 				float64 						`json:"promotion"`
	NumberOfFavourites 		int 							`json:"numberOfFavourites" gorm:"default:0"`
	NumberOfShares 			int 							`json:"numberOfShares" gorm:"default:0"`
	NumberOfSoldTickets		int 							`json:"numberOfSoldTickets" gorm:"default:0"`
	Status                  int		 						`json:"status" gorm:"default:0"`
	EventCycleType			int								`json:"eventCycleType" gorm:"default:0"`
	EventPaymentType		int								`json:"eventPaymentType" gorm:"default:0"`
	IsPrivate 				bool 							`json:"isPrivate" gorm:"default:0"`
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