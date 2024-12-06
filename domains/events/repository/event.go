package repository

import (
	"context"
	"gohub/configs"
	"gohub/database"
	"gohub/domains/events/dto"
	"gohub/domains/events/model"
	"gohub/pkg/paging"
	"gorm.io/gorm"
)

type IEventRepository interface {
	GetEventById(ctx context.Context, id string, preload bool) (*model.Event, error)
	CreateEvent(ctx context.Context, event *model.Event, req *dto.CreateEventReq) error
	UpdateEvent(ctx context.Context, event *model.Event) error
	ListEvents(ctx context.Context, req *dto.ListEventReq) ([]*model.Event, *paging.Pagination, error)
	ListCreatedEvents(ctx context.Context, userId string, req *dto.ListEventReq) ([]*model.Event, *paging.Pagination, error)
	ListTrashedEvents(ctx context.Context, userId string, req *dto.ListEventReq) ([]*model.Event, *paging.Pagination, error)
	ListFavouriteEvents(ctx context.Context, userId string, req *dto.ListEventReq) ([]*model.Event, *paging.Pagination, error)
	Delete(ctx context.Context, id string) error
	DeleteByIds(ctx context.Context, ids []string) error
	RestoreByIds(ctx context.Context, ids []string) error
	CreateEventFavourite(ctx context.Context, eventFavourite *model.EventFavourite) error
	RemoveEventFavourite(ctx context.Context, eventFavourite *model.EventFavourite) error
	DeleteEventFavourite(ctx context.Context, ids []string) error
	RestoreEventFavourite(ctx context.Context, ids []string) error
	MakeEventPublicOrPrivate(ctx context.Context, req *dto.MakeEventPublicOrPrivateReq, isPrivate bool) error
	ApplyCoupons(ctx context.Context, eventId string, req *dto.ApplyCouponReq) error
	RemoveCoupons(ctx context.Context, eventId string, req *dto.RemoveCouponReq) error
}

type EventRepo struct {
	db database.IDatabase
}

func NewEventRepository(db database.IDatabase) *EventRepo {
	return &EventRepo{db: db}
}

func (e *EventRepo) CreateEvent(ctx context.Context, event *model.Event, req *dto.CreateEventReq) error {
	handler := func() error {
		if err := e.db.Create(ctx, event); err != nil {
			return err
		}

		var eventCategories []*model.EventCategory
		for _, categoryId := range req.CategoryIds {
			eventCategories = append(eventCategories, &model.EventCategory{EventId: event.ID, CategoryId: categoryId})
		}
		if err := e.db.CreateInBatches(ctx, &eventCategories, len(eventCategories)); err != nil {
			return err
		}

		var reasons []*model.Reason
		for _, reason := range req.ReasonItems {
			reasons = append(reasons, &model.Reason{EventId: event.ID, Content: reason})
		}
		if err := e.db.CreateInBatches(ctx, &reasons, len(reasons)); err != nil {
			return err
		}

		return nil
	}

	err := e.db.WithTransaction(handler)

	if err != nil {
		return err
	}

	return nil
}

func (e *EventRepo) UpdateEvent(ctx context.Context, event *model.Event) error {
	return e.db.Update(ctx, event)
}

func (e *EventRepo) GetEventById(ctx context.Context, id string, preload bool) (*model.Event, error) {
	var event model.Event

	opts := []database.FindOption{
		database.WithQuery(database.NewQuery("id = ?", id)),
	}

	if preload {
		opts = append(opts, database.WithPreload([]string{"User", "SubImages", "Categories", "Reasons", "Coupons", "TicketTypes", "Reviews"}))
	}

	if err := e.db.FindOne(ctx, &event, opts...); err != nil {
		return nil, err
	}

	return &event, nil
}

func (e *EventRepo) ListEvents(ctx context.Context, req *dto.ListEventReq) ([]*model.Event, *paging.Pagination, error) {
	ctx, cancel := context.WithTimeout(ctx, configs.DatabaseTimeout)
	defer cancel()

	query := make([]database.Query, 0)

	if req.IsPrivate {
		query = append(query, database.NewQuery("is_private = ?", true))
	} else {
		query = append(query, database.NewQuery("is_private = ?", false))
	}

	if req.Name != "" {
		query = append(query, database.NewQuery("name LIKE ?", "%"+req.Name+"%"))
	}

	order := "created_at"
	if req.OrderBy != "" {
		order = req.OrderBy
		if req.OrderDesc {
			order += " DESC"
		}
	}

	var total int64
	if err := e.db.Count(ctx, &model.Event{}, &total, database.WithQuery(query...)); err != nil {
		return nil, nil, err
	}

	pagination := paging.NewPagination(req.Page, req.Limit, total)

	if req.TakeAll {
		pagination.PageSize = total
	}

	var events []*model.Event
	if err := e.db.Find(
		ctx,
		&events,
		database.WithQuery(query...),
		database.WithLimit(int(pagination.PageSize)),
		database.WithOffset(int(pagination.Skip)),
		database.WithOrder(order),
		database.WithPreload([]string{"Reviews"}),
	); err != nil {
		return nil, nil, err
	}

	return events, pagination, nil
}

func (e *EventRepo) ListCreatedEvents(ctx context.Context, userId string, req *dto.ListEventReq) ([]*model.Event, *paging.Pagination, error) {
	ctx, cancel := context.WithTimeout(ctx, configs.DatabaseTimeout)
	defer cancel()

	query := make([]database.Query, 0)

	if req.IsPrivate {
		query = append(query, database.NewQuery("user_id = ? AND is_private = ?", userId, true))
	} else {
		query = append(query, database.NewQuery("user_id = ? AND is_private = ?", userId, false))
	}

	if req.Name != "" {
		query = append(query, database.NewQuery("name LIKE ?", "%"+req.Name+"%"))
	}

	order := "created_at"
	if req.OrderBy != "" {
		order = req.OrderBy
		if req.OrderDesc {
			order += " DESC"
		}
	}

	var total int64
	if err := e.db.Count(ctx, &model.Event{}, &total, database.WithQuery(query...)); err != nil {
		return nil, nil, err
	}

	pagination := paging.NewPagination(req.Page, req.Limit, total)

	if req.TakeAll {
		pagination.PageSize = total
	}

	var events []*model.Event
	if err := e.db.Find(
		ctx,
		&events,
		database.WithQuery(query...),
		database.WithLimit(int(pagination.PageSize)),
		database.WithOffset(int(pagination.Skip)),
		database.WithOrder(order),
		database.WithPreload([]string{"Reviews"}),
	); err != nil {
		return nil, nil, err
	}

	return events, pagination, nil
}

func (e *EventRepo) ListTrashedEvents(ctx context.Context, userId string, req *dto.ListEventReq) ([]*model.Event, *paging.Pagination, error) {
	ctx, cancel := context.WithTimeout(ctx, configs.DatabaseTimeout)
	defer cancel()

	query := make([]database.Query, 0)

	if req.Name != "" {
		query = append(query, database.NewQuery("user_id = ? AND name LIKE ?", userId, "%"+req.Name+"%"))
	} else {
		query = append(query, database.NewQuery("user_id = ?", userId))
	}

	order := "created_at"
	if req.OrderBy != "" {
		order = req.OrderBy
		if req.OrderDesc {
			order += " DESC"
		}
	}

	var total int64
	if err := e.db.CountUnscoped(ctx, &model.Event{}, &total, database.WithQuery(query...)); err != nil {
		return nil, nil, err
	}

	pagination := paging.NewPagination(req.Page, req.Limit, total)

	if req.TakeAll {
		pagination.PageSize = total
	}

	var events []*model.Event
	if err := e.db.FindUnscoped(
		ctx,
		&events,
		database.WithQuery(query...),
		database.WithLimit(int(pagination.PageSize)),
		database.WithOffset(int(pagination.Skip)),
		database.WithOrder(order),
		database.WithPreload([]string{"Reviews"}),
	); err != nil {
		return nil, nil, err
	}

	return events, pagination, nil
}

func (e *EventRepo) ListFavouriteEvents(ctx context.Context, userId string, req *dto.ListEventReq) ([]*model.Event, *paging.Pagination, error) {
	ctx, cancel := context.WithTimeout(ctx, configs.DatabaseTimeout)
	defer cancel()

	query := make([]database.Query, 0)

	query = append(query, database.NewQuery("user_id = ?", userId))

	order := "created_at"
	if req.OrderBy != "" {
		order = req.OrderBy
		if req.OrderDesc {
			order += " DESC"
		}
	}

	var total int64
	if err := e.db.Count(ctx, &model.EventFavourite{}, &total, database.WithQuery(query...)); err != nil {
		return nil, nil, err
	}

	pagination := paging.NewPagination(req.Page, req.Limit, total)

	if req.TakeAll {
		pagination.PageSize = total
	}

	var eventFavourites []*model.EventFavourite
	if err := e.db.Find(
		ctx,
		&eventFavourites,
		database.WithQuery(query...),
		database.WithLimit(int(pagination.PageSize)),
		database.WithOffset(int(pagination.Skip)),
		database.WithOrder(order),
		database.WithPreload([]string{"Event", "Event.Reviews"}),
	); err != nil {
		return nil, nil, err
	}

	var events []*model.Event
	for _, eventFavourite := range eventFavourites {
		events = append(events, eventFavourite.Event)
	}

	return events, pagination, nil
}

func (e *EventRepo) Delete(ctx context.Context, id string) error {
	event, err := e.GetEventById(ctx, id, false)
	if err != nil {
		return err
	}

	if err := e.db.Delete(ctx, event); err != nil {
		return err
	}

	return e.DeleteEventFavourite(ctx, []string{id})
}

func (e *EventRepo) DeleteByIds(ctx context.Context, ids []string) error {
	err := e.db.DeleteByIds(ctx, &model.Event{}, ids)
	if err != nil {
		return err
	}

	return e.DeleteEventFavourite(ctx, ids)
}

func (e *EventRepo) RestoreByIds(ctx context.Context, ids []string) error {
	err := e.db.RestoreByIds(ctx, &model.Event{}, ids)
	if err != nil {
		return err
	}

	return e.RestoreEventFavourite(ctx, ids)
}

func (e *EventRepo) CreateEventFavourite(ctx context.Context, eventFavourite *model.EventFavourite) error {
	return e.db.Create(ctx, eventFavourite)
}

func (e *EventRepo) RemoveEventFavourite(ctx context.Context, eventFavourite *model.EventFavourite) error {
	query := database.NewQuery("user_id = ? AND event_id = ?", eventFavourite.UserId, eventFavourite.EventId)

	return e.db.ForceDelete(ctx, eventFavourite, database.WithQuery(query))
}

func (e *EventRepo) DeleteEventFavourite(ctx context.Context, ids []string) error {
	ctx, cancel := context.WithTimeout(ctx, configs.DatabaseTimeout)
	defer cancel()

	return e.db.GetDB().Model(&model.EventFavourite{}).Unscoped().Where("event_id IN ?", ids).Update("deleted_at", gorm.Expr("NOW()")).Error
}

func (e *EventRepo) RestoreEventFavourite(ctx context.Context, ids []string) error {
	ctx, cancel := context.WithTimeout(ctx, configs.DatabaseTimeout)
	defer cancel()

	return e.db.GetDB().Model(&model.EventFavourite{}).Unscoped().Where("event_id IN ?", ids).Update("deleted_at", nil).Error
}

func (e *EventRepo) MakeEventPublicOrPrivate(ctx context.Context, req *dto.MakeEventPublicOrPrivateReq, isPrivate bool) error {
	ctx, cancel := context.WithTimeout(ctx, configs.DatabaseTimeout)
	defer cancel()

	return e.db.GetDB().Model(&model.Event{}).Where("id IN ? AND user_id = ?", req.Ids, req.UserId).Update("is_private", isPrivate).Error
}

func (e *EventRepo) ApplyCoupons(ctx context.Context, eventId string, req *dto.ApplyCouponReq) error {
	var eventCoupons []*model.EventCoupons
	for _, id := range req.Ids {
		eventCoupons = append(eventCoupons, &model.EventCoupons{EventId: eventId, CouponId: id})
	}

	return e.db.CreateInBatches(ctx, &eventCoupons, len(eventCoupons))
}

func (e *EventRepo) RemoveCoupons(ctx context.Context, eventId string, req *dto.RemoveCouponReq) error {
	handler := func() error {
		for _, id := range req.Ids {
			eventCoupon := model.EventCoupons{EventId: eventId, CouponId: id}
			query := database.NewQuery("event_id = ? AND coupon_id = ?", eventCoupon.EventId, eventCoupon.CouponId)

			err := e.db.ForceDelete(ctx, eventCoupon, database.WithQuery(query))
			if err != nil {
				return err
			}
		}
		return nil
	}
	err := e.db.WithTransaction(handler)

	if err != nil {
		return err
	}

	return nil
}
