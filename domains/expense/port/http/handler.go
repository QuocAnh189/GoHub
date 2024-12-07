package http

import (
	"github.com/QuocAnh189/GoBin/logger"
	"github.com/gin-gonic/gin"
	"gohub/domains/expense/dto"
	"gohub/domains/expense/service"
	"gohub/pkg/messages"
	"gohub/pkg/response"
	"gohub/pkg/utils"
	"net/http"
)

type ExpenseHandler struct {
	service service.IExpenseService
}

func NewExpenseHandler(service service.IExpenseService) *ExpenseHandler {
	return &ExpenseHandler{
		service: service,
	}
}

//		@Summary	 Retrieve an expense by its ID
//	 @Description Fetches the details of a specific expense based on the provided expense ID.
//		@Tags		 Expenses
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"Category created successfully"
//		@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 404	{object}	response.Response	"Not Found - Category with the specified ID not found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/expenses/{expenseId} [get]
func (h *ExpenseHandler) GetExpenseById(c *gin.Context) {
	var res dto.Expense

	expenseId := c.Param("id")
	expense, err := h.service.GetExpenseById(c, expenseId)
	if err != nil {
		logger.Error("Failed to get expense detail: ", err)
		response.Error(c, http.StatusNotFound, err, "Not found")
		return
	}

	utils.MapStruct(&res, &expense)
	response.JSON(c, http.StatusOK, res)
}

//		@Summary	 Retrieve list expenses
//	 @Description Fetches the details of a specific expense based on the provided expense ID.
//		@Tags		 Expenses
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"Category created successfully"
//		@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 404	{object}	response.Response	"Not Found - Category with the specified ID not found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/expenses/get-by-event/{eventId} [get]
func (h *ExpenseHandler) GetExpensesByEvent(c *gin.Context) {
	var req dto.ListExpenseReq
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("Failed to parse request query: ", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	eventId := c.Param("eventId")
	expenses, pagination, err := h.service.GetExpenseByEvent(c, eventId, &req)
	if err != nil {
		logger.Error("Failed to get list products: ", err)
		response.Error(c, http.StatusInternalServerError, err, "Something went wrong")
		return
	}

	var res dto.ListExpenseRes
	utils.MapStruct(&res.Expense, &expenses)
	res.Pagination = pagination
	response.JSON(c, http.StatusOK, res)
}

//		@Summary	 Create a new expense
//	 @Description Creates a new expense based on the provided details.
//		@Tags		 Expenses
//		@Produce	 json
//		@Success	 201	{object}	response.Response	"Expense created successfully"
//		@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/expenses [post]
func (h *ExpenseHandler) CreateExpense(c *gin.Context) {
	var req dto.CreatedExpenseReq
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Failed to get body", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	expense, err := h.service.CreateExpense(c, &req)
	if err != nil {
		logger.Error("Failed to create expense ", err.Error())
		switch err.Error() {
		case messages.TitleExpenseAlreadyExists:
			response.Error(c, http.StatusConflict, err, messages.TitleExpenseAlreadyExists)
		default:
			response.Error(c, http.StatusInternalServerError, err, "Failed to create expense")
		}
		return
	}

	var res dto.Expense
	utils.MapStruct(&res, &expense)
	response.JSON(c, http.StatusOK, res)
}

//		@Summary	 Update an exists expense
//	 @Description Creates a new expense based on the provided details. The request must include multipart form data.
//		@Tags		 Expenses
//		@Produce	 json
//		@Success	 201	{object}	response.Response	"Expense updated successfully"
//		@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/expenses/{expenseId} [post]
func (h *ExpenseHandler) UpdateExpense(c *gin.Context) {
	var req dto.UpdatedExpenseReq
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Failed to get body", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	expenseId := c.Param("id")
	expense, err := h.service.UpdateExpense(c, expenseId, &req)
	if err != nil {
		logger.Error("Failed to update expense ", err.Error())
		switch err.Error() {
		default:
			response.Error(c, http.StatusInternalServerError, err, "Failed to update expense")
		}
		return
	}

	var res dto.Expense
	utils.MapStruct(&res, &expense)
	response.JSON(c, http.StatusOK, res)
}

//		@Summary	 Delete an expense
//	 @Description Creates a new expense based on the provided details
//		@Tags		 Expenses
//		@Produce	 json
//		@Success	 201	{object}	response.Response	"Expense created successfully"
//		@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/expenses/{expenseId} [post]
func (h *ExpenseHandler) DeleteExpense(c *gin.Context) {
	expenseId := c.Param("id")

	err := h.service.DeleteExpense(c, expenseId)

	if err != nil {
		logger.Error("Failed to delete expense: ", err)
		response.Error(c, http.StatusNotFound, err, "Not found")
		return
	}

	response.JSON(c, http.StatusOK, "Delete expense successfully")
}

//		@Summary	 Create a new subExpense
//	 @Description Creates a new expense based on the provided detail.
//		@Tags		 Expenses
//		@Produce	 json
//		@Success	 201	{object}	response.Response	"SubExpense created successfully"
//		@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/expenses/{expensesId}/sub-expense [post]
func (h *ExpenseHandler) CreateSubExpense(c *gin.Context) {
	var req dto.CreatedSubExpenseReq
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Failed to get body", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	subExpense, err := h.service.CreateSubExpense(c, &req)
	if err != nil {
		logger.Error("Failed to create expense ", err.Error())
		switch err.Error() {
		case messages.NameSubExpenseAlreadyExists:
			response.Error(c, http.StatusConflict, err, messages.NameSubExpenseAlreadyExists)
		default:
			response.Error(c, http.StatusInternalServerError, err, "Failed to create expense")
		}
		return
	}

	var res dto.SubExpense
	utils.MapStruct(&res, &subExpense)
	response.JSON(c, http.StatusOK, res)
}

//		@Summary	 Update an exists subExpense
//	 @Description Creates a new expense based on the provided details.
//		@Tags		 Expenses
//		@Produce	 json
//		@Success	 201	{object}	response.Response	"SubExpense created successfully"
//		@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/expenses/{expensesId}/sub-expense/{subExpenseId} [put]
func (h *ExpenseHandler) UpdateSubExpense(c *gin.Context) {
	var req dto.UpdateSubExpenseReq
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Failed to get body", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	subExpenseId := c.Param("subExpenseId")
	err := h.service.UpdateSubExpense(c, subExpenseId, &req)
	if err != nil {
		logger.Error("Failed to update sub_expense ", err.Error())
		switch err.Error() {
		default:
			response.Error(c, http.StatusInternalServerError, err, "Failed to update expense")
		}
		return
	}

	response.JSON(c, http.StatusOK, "Update Sub Expense successfully")
}

//		@Summary	 Create a new expense
//	 @Description Creates a new expense based on the provided details.
//		@Tags		 Expenses
//		@Produce	 json
//		@Success	 201	{object}	response.Response	"SubExpense created successfully"
//		@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/expenses/{expensesId}/sub-expense/{subExpenseId} [delete]
func (h *ExpenseHandler) DeleteSubExpense(c *gin.Context) {
	subExpenseId := c.Param("subExpenseId")

	err := h.service.DeleteSubExpense(c, subExpenseId)

	if err != nil {
		logger.Error("Failed to delete subExpense: ", err)
		response.Error(c, http.StatusNotFound, err, "Not found")
		return
	}

	response.JSON(c, http.StatusOK, "Delete subExpense successfully")
}
