package service

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5/pgconn"
	"gohub/domains/events/dto"
	"gohub/domains/events/model"
	"gohub/domains/events/repository"
	"gohub/internal/libs/logger"
	"gohub/internal/libs/validation"
	"gohub/pkg/messages"
	"gohub/pkg/paging"
	"gohub/pkg/utils"
	"gorm.io/gorm"
)

type IEventService interface {
	GetEvents(ctx context.Context, req *dto.ListEventReq) ([]*model.Event, *paging.Pagination, error)
	CreateEvent(ctx context.Context, req *dto.CreateEventReq) (*model.Event, error)
	GetEventById(ctx context.Context, id string) (*model.Event, error)
	UpdateEvent(ctx context.Context, id string, req *dto.UpdateEventReq) (*model.Event, error)
	DeleteEvent(ctx context.Context, id string) error
	DeleteEvents(ctx context.Context, ids *dto.DeleteRequest) error
	GetCreatedEvent(ctx context.Context, userId string, req *dto.ListEventReq, statistic *dto.StatisticMyEvent) ([]*model.Event, *paging.Pagination, error)
	GetCreatedEventAnalysis(ctx context.Context, userId string, req *dto.ListEventReq) ([]*model.Event, *paging.Pagination, error)
	RestoreEvents(ctx context.Context, ids *dto.RestoreRequest) error
	GetTrashedEvent(ctx context.Context, userId string, req *dto.ListEventReq) ([]*model.Event, *paging.Pagination, error)
	FavouriteEvent(ctx context.Context, req *dto.CreateEventFavouriteReq) error
	UnFavouriteEvent(ctx context.Context, req *dto.CreateEventFavouriteReq) error
	GetFavouriteEvent(ctx context.Context, userId string, req *dto.ListEventReq) ([]*model.Event, *paging.Pagination, error)
	MakeEventPrivate(ctx context.Context, req *dto.MakeEventPublicOrPrivateReq) error
	MakeEventPublic(ctx context.Context, req *dto.MakeEventPublicOrPrivateReq) error
	ApplyCoupons(ctx context.Context, eventId string, req *dto.ApplyCouponReq) error
	CheckFavourite(ctx context.Context, req *dto.UserFavouriteEvent) (bool, error)
}

type EventService struct {
	validator validation.Validation
	eventRepo repository.IEventRepository
}

func NewEventService(
	validator validation.Validation,
	eventRepo repository.IEventRepository,
) *EventService {
	return &EventService{
		validator: validator,
		eventRepo: eventRepo,
	}
}

func (e *EventService) GetEvents(ctx context.Context, req *dto.ListEventReq) ([]*model.Event, *paging.Pagination, error) {
	events, pagination, err := e.eventRepo.ListEvents(ctx, req)
	if err != nil {
		return nil, nil, err
	}

	return events, pagination, nil
}

func (e *EventService) CreateEvent(ctx context.Context, req *dto.CreateEventReq) (*model.Event, error) {
	var event model.Event
	utils.MapStruct(&event, req)

	err := e.eventRepo.CreateEvent(ctx, &event, req)
	if err != nil {
		logger.Errorf("Create fail, error: %s", err)
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" {
			return nil, errors.New(messages.TitleExpenseAlreadyExists)
		}
		return nil, err
	}

	return &event, nil
}

func (e *EventService) GetEventById(ctx context.Context, id string) (*model.Event, error) {
	event, err := e.eventRepo.GetEventById(ctx, id, true)
	if err != nil {
		return nil, err
	}

	var totalRate float32
	for _, review := range event.Reviews {
		totalRate += review.Rate
	}
	if len(event.Reviews) > 0 {
		event.AverageRate = totalRate / float32(len(event.Reviews))
	} else {
		event.AverageRate = 0
	}

	return event, nil
}

func (e *EventService) UpdateEvent(ctx context.Context, id string, req *dto.UpdateEventReq) (*model.Event, error) {
	if err := e.validator.ValidateStruct(req); err != nil {
		return nil, err
	}

	event, err := e.eventRepo.GetEventById(ctx, id, false)
	if err != nil {
		logger.Errorf("Update.GetEventByID fail, id: %s, error: %s", id, err)
		return nil, errors.New(messages.CategoryNotFound)
	}

	utils.MapStruct(event, req)
	err = e.eventRepo.UpdateEvent(ctx, event, req)
	if err != nil {
		logger.Errorf("Update fail, id: %s, error: %s", id, err)
		return nil, err
	}

	return event, nil
}

func (e *EventService) DeleteEvent(ctx context.Context, id string) error {
	err := e.eventRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (e *EventService) DeleteEvents(ctx context.Context, ids *dto.DeleteRequest) error {
	var err error
	if len(ids.Ids) == 1 {
		err = e.eventRepo.Delete(ctx, ids.Ids[0])
	} else {
		err = e.eventRepo.DeleteByIds(ctx, ids.Ids)
	}

	if err != nil {
		return err
	}

	return nil
}

func (e *EventService) GetCreatedEvent(ctx context.Context, userId string, req *dto.ListEventReq, statistic *dto.StatisticMyEvent) ([]*model.Event, *paging.Pagination, error) {
	events, pagination, err := e.eventRepo.ListCreatedEvents(ctx, userId, req, statistic)
	if err != nil {
		return nil, nil, err
	}

	return events, pagination, nil
}

func (e *EventService) GetCreatedEventAnalysis(ctx context.Context, userId string, req *dto.ListEventReq) ([]*model.Event, *paging.Pagination, error) {
	events, pagination, err := e.eventRepo.ListCreatedEventsAnalysis(ctx, userId, req)
	if err != nil {
		return nil, nil, err
	}

	return events, pagination, nil
}

func (e *EventService) RestoreEvents(ctx context.Context, ids *dto.RestoreRequest) error {
	err := e.eventRepo.RestoreByIds(ctx, ids.Ids)

	if err != nil {
		return err
	}

	return nil
}

func (e *EventService) GetTrashedEvent(ctx context.Context, userId string, req *dto.ListEventReq) ([]*model.Event, *paging.Pagination, error) {
	events, pagination, err := e.eventRepo.ListTrashedEvents(ctx, userId, req)
	if err != nil {
		return nil, nil, err
	}

	return events, pagination, nil
}

func (e *EventService) FavouriteEvent(ctx context.Context, req *dto.CreateEventFavouriteReq) error {
	if err := e.validator.ValidateStruct(req); err != nil {
		return err
	}

	var eventFavourite model.EventFavourite
	utils.MapStruct(&eventFavourite, req)

	err := e.eventRepo.CreateEventFavourite(ctx, &eventFavourite)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" {
			return errors.New(messages.EventFavouriteAlreadyExists)
		}
		return err
	}

	return nil
}

func (e *EventService) UnFavouriteEvent(ctx context.Context, req *dto.CreateEventFavouriteReq) error {
	if err := e.validator.ValidateStruct(req); err != nil {
		return err
	}

	var eventFavourite *model.EventFavourite
	utils.MapStruct(&eventFavourite, &req)

	if err := e.eventRepo.RemoveEventFavourite(ctx, eventFavourite); err != nil {
		return err
	}
	return nil
}

func (e *EventService) GetFavouriteEvent(ctx context.Context, userId string, req *dto.ListEventReq) ([]*model.Event, *paging.Pagination, error) {
	events, pagination, err := e.eventRepo.ListFavouriteEvents(ctx, userId, req)
	if err != nil {
		return nil, nil, err
	}

	return events, pagination, nil
}

func (e *EventService) MakeEventPrivate(ctx context.Context, req *dto.MakeEventPublicOrPrivateReq) error {
	if err := e.validator.ValidateStruct(req); err != nil {
		return err
	}

	err := e.eventRepo.MakeEventPublicOrPrivate(ctx, req, true)
	if err != nil {
		return err
	}

	return nil
}

func (e *EventService) MakeEventPublic(ctx context.Context, req *dto.MakeEventPublicOrPrivateReq) error {
	if err := e.validator.ValidateStruct(req); err != nil {
		return err
	}

	err := e.eventRepo.MakeEventPublicOrPrivate(ctx, req, false)
	if err != nil {
		return err
	}

	return nil
}

func (e *EventService) ApplyCoupons(ctx context.Context, eventId string, req *dto.ApplyCouponReq) error {
	if err := e.validator.ValidateStruct(req); err != nil {
		return err
	}

	if err := e.eventRepo.ApplyCoupons(ctx, eventId, req); err != nil {
		logger.Errorf("Create fail, error: %s", err)
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" {
			return errors.New(messages.ApplyCouponAlreadyExists)
		}

		return err
	}

	return nil
}

func (e *EventService) CheckFavourite(ctx context.Context, req *dto.UserFavouriteEvent) (bool, error) {
	result, err := e.eventRepo.CheckFavourite(ctx, req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return result, nil
}
