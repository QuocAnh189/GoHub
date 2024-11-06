package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Event struct {
	ID 						string 		`json:"id" gorm:"unique;not null;index;primary_key"`
	UserId 					string 		`json:"userId"`
	Name 					string 		`json:"name" gorm:"not null"`
	Description 			string 		`json:"description"`
	CoverImageUrl 			string 		`json:"coverImageUrl"`
	CoverImageFileName 		string 		`json:"coverImageFileName"`
	StartDate 				string 		`json:"startDate" gorm:"not null"`
	EndDate 				string 		`json:"endDate" gorm:"not null"`
	Location 				string 		`json:"location" gorm:"not null"`
	Promotion 				float64 	`json:"promotion"`
	NumberOfFavourites 		int 		`json:"numberOfFavourites" gorm:"default:0"`
	NumberOfShares 			int 		`json:"numberOfShares" gorm:"default:0"`
	NumberOfSoldTickets		int 		`json:"numberOfSoldTickets" gorm:"default:0"`
	Status                  int		 	`json:"status" gorm:"default:0"`
	EventCycleType			int			`json:"eventCycleType" gorm:"default:0"`
	EventPaymentType		int			`json:"eventPaymentType" gorm:"default:0"`
	IsPrivate 				bool 		`json:"isPrivate" gorm:"default:0"`
	IsDeleted 				bool       	`json:"isDeleted" gorm:"default:0"`
	DeletedAt 				*time.Time 	`json:"deletedAt" gorm:"index"`
	CreatedAt 				time.Time  	`json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt 				time.Time  	`json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`
}

func (e *Event) BeforeCreate(db *gorm.DB) (err error) {
	e.ID = uuid.New().String()
	return nil
}