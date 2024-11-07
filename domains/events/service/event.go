package service

import (
	"context"
	"gohub/domains/events/model"
	"gohub/domains/events/repository"

	"github.com/QuocAnh189/GoBin/validation"
)

type IEventService interface {
	GetEvents(ctx context.Context) 
	CreateEvent(ctx context.Context, event *model.Event)
	GetEvent(ctx context.Context, id string)
	UpdateEvent(ctx context.Context, event *model.Event) 
	DeleteEvent(ctx context.Context, id string)
	GetCreatedEvent(ctx context.Context, id string) 
	DeletePermanentlyEvent(ctx context.Context, id string) 
	RestoreEvent(ctx context.Context, id string) 
	GetTrashedEvent(ctx context.Context, id string) 
	FavouriteEvent(ctx context.Context, id string) 
	UnfavouriteEvent(ctx context.Context, id string) 
	GetFavouriteEvent(ctx context.Context, id string) 
	MakeEventPrivate(ctx context.Context, id string) 
	MakeEventPublic(ctx context.Context, id string) 
}

type EventService struct {
	validator validation.Validation
	repo      repository.IEventRepository
}

func NewEventService(validator validation.Validation, repo repository.IEventRepository) *EventService {
	return &EventService{
		validator: validator,
		repo:      repo,
	}
}

func (e *EventService) GetEvents(ctx context.Context) {
	panic("unimplemented")
}

func (e *EventService) CreateEvent(ctx context.Context, event *model.Event) {
	panic("unimplemented")
}

func (e *EventService) GetEvent(ctx context.Context, id string) {
    panic("unimplemented")
}

func (e *EventService) UpdateEvent(ctx context.Context, event *model.Event) {
	panic("unimplemented")
}

func (e *EventService) DeleteEvent(ctx context.Context, id string) {	
	panic("unimplemented")
}

func (e *EventService) GetCreatedEvent(ctx context.Context, id string) {
	panic("unimplemented")
}

func (e *EventService) DeletePermanentlyEvent(ctx context.Context, id string) {
	panic("unimplemented")
}

func (e *EventService) RestoreEvent(ctx context.Context, id string) {
    panic("unimplemented")
}

func (e *EventService) GetTrashedEvent(ctx context.Context, id string) {
    panic("unimplemented")
}

func (e *EventService) FavouriteEvent(ctx context.Context, id string) {	
	panic("unimplemented")
}

func (e *EventService) UnfavouriteEvent(ctx context.Context, id string) {
	panic("unimplemented")
}

func (e *EventService) GetFavouriteEvent(ctx context.Context, id string) {
	panic("unimplemented")
}

func (e *EventService) MakeEventPrivate(ctx context.Context, id string) {
    panic("unimplemented")
}

func (e *EventService) MakeEventPublic(ctx context.Context, id string) {
	panic("unimplemented")
}

