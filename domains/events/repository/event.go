package repository

import (
	"context"
	"gohub/configs"
	"gohub/database"
	"gohub/domains/events/dto"
	"gohub/domains/events/model"
	"gohub/pkg/paging"
	"gohub/pkg/utils"
	"gorm.io/gorm"
)

type IEventRepository interface {
	GetEventById(ctx context.Context, id string, preload bool) (*model.Event, error)
	CreateEvent(ctx context.Context, event *model.Event, req *dto.CreateEventReq) error
	UpdateEvent(ctx context.Context, event *model.Event, req *dto.UpdateEventReq) error
	ListEvents(ctx context.Context, req *dto.ListEventReq) ([]*model.Event, *paging.Pagination, error)
	ListCreatedEvents(ctx context.Context, userId string, req *dto.ListEventReq, statistic *dto.StatisticMyEvent) ([]*model.Event, *paging.Pagination, error)
	ListCreatedEventsAnalysis(ctx context.Context, userId string, req *dto.ListEventReq) ([]*model.Event, *paging.Pagination, error)
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
	CheckFavourite(ctx context.Context, req *dto.UserFavouriteEvent) (bool, error)
}

type EventRepo struct {
	db database.IDatabase
}

func NewEventRepository(db database.IDatabase) *EventRepo {
	return &EventRepo{db: db}
}

func (e *EventRepo) CreateEvent(ctx context.Context, event *model.Event, req *dto.CreateEventReq) error {
	handler := func() error {
		if req.CoverImage.Header != nil && req.CoverImage.Filename != "" {
			uploadUrl, err := utils.ImageUpload(req.CoverImage, "/eventhub/events")
			if err != nil {
				return err
			}

			event.CoverImageFileName = req.CoverImage.Filename
			event.CoverImageUrl = uploadUrl
		}

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

		var subImages []*model.EventSubImage
		for _, subImage := range req.SubImageItems {
			if subImage.Header != nil && subImage.Filename != "" {
				uploadUrl, err := utils.ImageUpload(subImage, "/eventhub/events")
				if err != nil {
					return err
				}

				subImages = append(subImages, &model.EventSubImage{EventId: event.ID, ImageUrl: uploadUrl, ImageFileName: subImage.Filename})
			}
		}
		if err := e.db.CreateInBatches(ctx, &subImages, len(subImages)); err != nil {
			return err
		}

		var reasons []*model.Reason
		for _, reason := range req.ReasonItems {
			reasons = append(reasons, &model.Reason{EventId: event.ID, Content: reason})
		}
		if err := e.db.CreateInBatches(ctx, &reasons, len(reasons)); err != nil {
			return err
		}

		var ticketTypes []*model.TicketType
		for _, ticketItem := range req.TicketTypeItems {
			ticketTypes = append(ticketTypes,
				&model.TicketType{EventId: event.ID, Name: ticketItem.Name, Quantity: ticketItem.Quantity, Price: ticketItem.Price},
			)
		}
		if err := e.db.CreateInBatches(ctx, &ticketTypes, len(ticketTypes)); err != nil {
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

func (e *EventRepo) UpdateEvent(ctx context.Context, event *model.Event, req *dto.UpdateEventReq) error {
	handler := func() error {
		if req.CoverImage.Header != nil && req.CoverImage.Filename != "" {
			uploadUrl, err := utils.ImageUpload(req.CoverImage, "/eventhub/events")
			if err != nil {
				return err
			}

			event.CoverImageFileName = req.CoverImage.Filename
			event.CoverImageUrl = uploadUrl
		}

		if err := e.db.Update(ctx, event); err != nil {
			return err
		}

		if err := e.db.ForceDelete(ctx, model.EventCategory{}, database.WithQuery(database.NewQuery("event_id = ?", req.ID))); err != nil {
			return err
		}
		var eventCategories []*model.EventCategory
		for _, categoryId := range req.CategoryIds {
			eventCategories = append(eventCategories, &model.EventCategory{EventId: event.ID, CategoryId: categoryId})
		}
		if err := e.db.CreateInBatches(ctx, &eventCategories, len(eventCategories)); err != nil {
			return err
		}

		//if err := e.db.ForceDelete(ctx, model.EventSubImage{}, database.WithQuery(database.NewQuery("event_id = ?", req.ID))); err != nil {
		//	return err
		//}
		//var subImages []*model.EventSubImage
		//for _, subImage := range req.SubImageItems {
		//	if subImage.Header != nil && subImage.Filename != "" {
		//		uploadUrl, err := utils.ImageUpload(subImage, "/eventhub/events")
		//		if err != nil {
		//			return err
		//		}
		//
		//		subImages = append(subImages, &model.EventSubImage{EventId: event.ID, ImageUrl: uploadUrl, ImageFileName: subImage.Filename})
		//	}
		//}
		//if err := e.db.CreateInBatches(ctx, &subImages, len(subImages)); err != nil {
		//	return err
		//}

		if err := e.db.ForceDelete(ctx, model.Reason{}, database.WithQuery(database.NewQuery("event_id = ?", req.ID))); err != nil {
			return err
		}
		var reasons []*model.Reason
		for _, reason := range req.ReasonItems {
			reasons = append(reasons, &model.Reason{EventId: event.ID, Content: reason})
		}
		if err := e.db.CreateInBatches(ctx, &reasons, len(reasons)); err != nil {
			return err
		}

		if err := e.db.ForceDelete(ctx, model.TicketType{}, database.WithQuery(database.NewQuery("event_id = ?", req.ID))); err != nil {
			return err
		}
		var ticketTypes []*model.TicketType
		for _, ticketItem := range req.TicketTypeItems {
			ticketTypes = append(ticketTypes,
				&model.TicketType{EventId: event.ID, Name: ticketItem.Name, Quantity: ticketItem.Quantity, Price: ticketItem.Price},
			)
		}
		if err := e.db.CreateInBatches(ctx, &ticketTypes, len(ticketTypes)); err != nil {
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
	args := make([]interface{}, 0)

	queryString := "is_private = ?"
	args = append(args, false)

	if req.Search != "" {
		queryString += " AND name LIKE ?"
		args = append(args, "%"+req.Search+"%")
	}

	if len(req.CategoryIds) > 0 {
		queryString += " AND event_categories.category_id IN (?)"
		args = append(args, req.CategoryIds)
	}

	if req.Status != "All" {
		switch req.Status {
		case "Upcoming":
			queryString += " AND start_time::date > NOW()::date"
		case "Opening":
			queryString += " AND NOW()::date BETWEEN start_time::date AND end_time::date"
		case "Close":
			queryString += " AND end_time::date < NOW()::date"
		default:
			queryString += ""
		}
	}

	if req.StartTimeRange != "" && req.EndTimeRange != "" {
		queryString += " AND start_time BETWEEN ? AND ?"
		args = append(args, req.StartTimeRange, req.EndTimeRange)
	}

	havingClause := ""
	if req.MinRate > 0 && req.MinRate <= 5 {
		havingClause = "COALESCE(AVG(reviews.rate), 0) >= ?"
	}

	query = append(query, database.NewQuery(queryString, args...))

	order := "created_at DESC"
	if req.OrderBy != "" {
		order = req.OrderBy
		if req.OrderDesc {
			order += " DESC"
		}
	}

	var total int64
	if err := e.db.Count(
		ctx,
		&model.Event{},
		&total,
		database.WithQuery(query...),
		database.WithJoin(`
			INNER JOIN event_categories ON event_categories.event_id = events.id
			LEFT JOIN reviews ON reviews.event_id = events.id
    	`),
		database.WithSelect(`
			events.*,
			COALESCE(AVG(reviews.rate), 0) AS average_rate
		`),
		database.WithGroupBy("events.id"),
		database.WithOrder("events.id"),
		database.WithHaving(havingClause, req.MinRate),
	); err != nil {
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
		database.WithJoin(`
			INNER JOIN event_categories ON event_categories.event_id = events.id
			LEFT JOIN reviews ON reviews.event_id = events.id
    	`),
		database.WithSelect(`
			events.*, 
			COALESCE(AVG(reviews.rate), 0) AS average_rate
    	`),
		database.WithGroupBy("events.id"),
		database.WithHaving(havingClause, req.MinRate),
		database.WithPreload([]string{"Reviews", "Categories"}),
	); err != nil {
		return nil, nil, err
	}

	return events, pagination, nil
}

func (e *EventRepo) ListCreatedEvents(ctx context.Context, userId string, req *dto.ListEventReq, statistic *dto.StatisticMyEvent) ([]*model.Event, *paging.Pagination, error) {
	ctx, cancel := context.WithTimeout(ctx, configs.DatabaseTimeout)
	defer cancel()

	query := make([]database.Query, 0)
	args := make([]interface{}, 0)

	queryString := "user_id = ?"
	args = append(args, userId)

	if len(req.CategoryIds) > 0 {
		valid := true
		for _, id := range req.CategoryIds {
			if id == "" {
				valid = false
			}
		}

		if valid {
			queryString += " AND event_categories.category_id IN (?)"
			args = append(args, req.CategoryIds)
		}
	}

	switch req.Visibility {
	case "Public":
		queryString += " AND is_private = ?"
		args = append(args, false)
	case "Private":
		queryString += " AND is_private = ?"
		args = append(args, true)
	default:
		break
	}

	switch req.Status {
	case "Upcoming":
		queryString += " AND start_time::date > NOW()::date"
	case "Opening":
		queryString += " AND NOW()::date BETWEEN start_time::date AND end_time::date"
	case "Close":
		queryString += " AND end_time::date < NOW()::date"
	default:
		queryString += ""
	}

	switch req.PaymentType {
	case "Free":
		queryString += " AND event_payment_type = ?"
		args = append(args, "Free")
	case "Paid":
		queryString += " AND event_payment_type = ?"
		args = append(args, "Paid")
	default:
		queryString += ""
	}

	if req.Search != "" {
		queryString += " AND name ILIKE ?"
		args = append(args, "%"+req.Search+"%")
	}

	query = append(query, database.NewQuery(queryString, args...))

	var total int64
	if err := e.db.Count(
		ctx,
		&model.Event{},
		&total,
		database.WithQuery(query...),
		database.WithJoin(`
			INNER JOIN event_categories ON event_categories.event_id = events.id
		`)); err != nil {
		return nil, nil, err
	}

	pagination := paging.NewPagination(req.Page, req.Limit, total)

	if req.TakeAll {
		pagination.PageSize = total
	}

	order := "events.created_at DESC"
	if req.OrderBy != "" {
		order = req.OrderBy
		if req.OrderDesc {
			order += " DESC"
		}
	}

	var events []*model.Event
	if err := e.db.Find(
		ctx,
		&events,
		database.WithQuery(query...),
		database.WithLimit(int(pagination.PageSize)),
		database.WithOffset(int(pagination.Skip)),
		database.WithOrder(order),
		database.WithSelect("events.id, events.name, events.cover_image_url, events.start_time, events.location, events.is_private, events.deleted_at"),
		database.WithJoin(`
			INNER JOIN event_categories ON event_categories.event_id = events.id
		`),
		database.WithPreload([]string{"Coupons"}),
	); err != nil {
		return nil, nil, err
	}

	var totalAll int64
	totalAllQuery := database.NewQuery("user_id = ?", userId)
	if err := e.db.Count(ctx, &model.Event{}, &totalAll, database.WithQuery(totalAllQuery)); err != nil {
		return nil, nil, err
	}
	statistic.TotalAll = totalAll

	var totalPublic int64
	totalPublicQuery := database.NewQuery("user_id = ? AND is_private = ?", userId, false)
	if err := e.db.Count(ctx, &model.Event{}, &totalPublic, database.WithQuery(totalPublicQuery)); err != nil {
		return nil, nil, err
	}
	statistic.TotalPublic = totalPublic
	statistic.TotalPrivate = totalAll - totalPublic

	return events, pagination, nil
}

func (e *EventRepo) ListCreatedEventsAnalysis(ctx context.Context, userId string, req *dto.ListEventReq) ([]*model.Event, *paging.Pagination, error) {
	ctx, cancel := context.WithTimeout(ctx, configs.DatabaseTimeout)
	defer cancel()

	query := make([]database.Query, 0)
	args := make([]interface{}, 0)

	queryString := "events.user_id = ?"
	args = append(args, userId)

	if req.Search != "" {
		queryString += " AND name ILIKE ?"
		args = append(args, "%"+req.Search+"%")
	}

	query = append(query, database.NewQuery(queryString, args...))

	var total int64
	if err := e.db.Count(
		ctx,
		&model.Event{},
		&total,
		database.WithQuery(query...),
	); err != nil {
		return nil, nil, err
	}

	pagination := paging.NewPagination(req.Page, req.Limit, total)

	if req.TakeAll {
		pagination.PageSize = total
	}

	order := "created_at DESC"
	if req.OrderBy != "" {
		order = req.OrderBy
		if req.OrderDesc {
			order += " DESC"
		}
	}

	var events []*model.Event
	if err := e.db.Find(
		ctx,
		&events,
		database.WithQuery(query...),
		database.WithLimit(int(pagination.PageSize)),
		database.WithOffset(int(pagination.Skip)),
		database.WithOrder(order),
		database.WithJoin(`
			LEFT JOIN event_favourites ON event_favourites.event_id = events.id
			LEFT JOIN reviews ON reviews.event_id = events.id
    	`),
		database.WithSelect(`
			events.*, 
			COUNT(DISTINCT event_favourites.id) AS total_favourite,
			COALESCE(AVG(reviews.rate), 0) AS average_rate
    	`),
		database.WithGroupBy("events.id"),
		database.WithPreload([]string{"Expenses"}),
	); err != nil {
		return nil, nil, err
	}

	return events, pagination, nil
}

func (e *EventRepo) ListTrashedEvents(ctx context.Context, userId string, req *dto.ListEventReq) ([]*model.Event, *paging.Pagination, error) {
	ctx, cancel := context.WithTimeout(ctx, configs.DatabaseTimeout)
	defer cancel()

	query := make([]database.Query, 0)
	args := make([]interface{}, 0)

	queryString := "user_id = ?"
	args = append(args, userId)

	if req.Search != "" {
		queryString += " AND name LIKE ?"
		args = append(args, "%"+req.Search+"%")
	}

	query = append(query, database.NewQuery(queryString, args...))

	order := "created_at DESC"
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
	args := make([]interface{}, 0)

	queryString := "event_favourites.user_id = ?"
	args = append(args, userId)

	if req.Search != "" {
		queryString += " AND name LIKE ?"
		args = append(args, "%"+req.Search+"%")
	}

	query = append(query, database.NewQuery(queryString, args...))

	order := "created_at DESC"
	if req.OrderBy != "" {
		order = req.OrderBy
		if req.OrderDesc {
			order += " DESC"
		}
	}

	var total int64
	if err := e.db.Count(
		ctx,
		&model.Event{},
		&total,
		database.WithQuery(query...),
		database.WithJoin(`
			INNER JOIN event_favourites ON event_favourites.event_id = events.id
			LEFT JOIN reviews ON reviews.event_id = events.id
		`),
		database.WithSelect(`
			events.*,
			COALESCE(AVG(reviews.rate), 0) AS average_rate
		`),
		database.WithGroupBy("events.id"),
		database.WithOrder("events.id"),
	); err != nil {
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
		database.WithJoin(`
			INNER JOIN event_favourites ON event_favourites.event_id = events.id
			LEFT JOIN reviews ON reviews.event_id = events.id
		`),
		database.WithSelect(`
			events.*,
			COALESCE(AVG(reviews.rate), 0) AS average_rate
		`),
		database.WithGroupBy("events.id"),
		database.WithPreload([]string{"Categories"}),
	); err != nil {
		return nil, nil, err
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
	query := database.NewQuery("event_id = ?", eventId)

	err := e.db.ForceDelete(ctx, model.EventCoupons{}, database.WithQuery(query))
	if err != nil {
		return err
	}

	var eventCoupons []*model.EventCoupons
	for _, id := range req.Ids {
		eventCoupons = append(eventCoupons, &model.EventCoupons{EventId: eventId, CouponId: id})
	}

	return e.db.CreateInBatches(ctx, &eventCoupons, len(eventCoupons))
}

func (e *EventRepo) CheckFavourite(ctx context.Context, req *dto.UserFavouriteEvent) (bool, error) {
	query := database.NewQuery("user_id = ? AND event_id = ?", req.UserId, req.EventId)
	if err := e.db.FindOne(ctx, &model.EventFavourite{}, database.WithQuery(query)); err != nil {
		return false, err
	}
	return true, nil
}
