package handler

import (
	"api/api/token"
	"api/config"
	pb "api/genprotos/budgeting"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Create Transaction
// @Description Create Transaction
// @Tags Transaction
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Create body pb.CreateTransactionRequest true "Create"
// @Success 201 {object} string "Success"
// @Failure 400 {string} string "Error"
// @Router /transaction/create [post]
func (h *BudgetingHandler) CreateTransaction(ctx *gin.Context) {
	req := &pb.CreateTransactionRequest{}
	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c := config.Load()
	id, _ := token.GetIdFromToken(ctx.Request, &c)
	req.Transaction.UserId = id
	_, err := h.Transaction.CreateTransaction(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Transaction not created", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Transaction Created Successfully"})
}

// @Summary Update Transaction
// @Description Update Transaction
// @Tags Transaction
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Update body pb.UpdateTransactionRequest true "Update"
// @Success 200 {object} string "Success"
// @Failure 400 {string} string "Error"
// @Router /transaction/update [put]
func (h *BudgetingHandler) UpdateTransaction(ctx *gin.Context) {
	req := &pb.UpdateTransactionRequest{}
	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c := config.Load()
	id, _ := token.GetIdFromToken(ctx.Request, &c)
	req.Transaction.UserId = id
	_, err := h.Transaction.UpdateTransaction(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Transaction Updated Successfully"})
}

// @Summary Delete Transaction
// @Description Delete a transaction and revert the user's balance accordingly
// @Tags Transaction
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Transaction ID"
// @Success 200 {object} string "Success"
// @Failure 400 {string} string "Error"
// @Router /transaction/delete/{id} [delete]
func (h *BudgetingHandler) DeleteTransaction(ctx *gin.Context) {
	id := ctx.Param("id")
	_, err := h.Transaction.DeleteTransaction(context.Background(), &pb.DeleteTransactionRequest{Id: id})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Transaction Deleted Successfully"})
}

// @Summary Get Transaction
// @Description Get an existing Transaction record by ID
// @Tags Transaction
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Transaction ID"
// @Success 200 {object} pb.GetTransactionResponse
// @Failure 400 {string} string "Error"
// @Router /transaction/get/{id} [get]
func (h *BudgetingHandler) GetTransaction(ctx *gin.Context) {
	id := ctx.Param("id")
	req := &pb.GetTransactionRequest{Id: id}

	res, err := h.Transaction.GetTransaction(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// @Summary ListTransactions
// @Description ListTransactions
// @Tags Transaction
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param limit query int false "Limit"
// @Param page query int false "Page"
// @Success 200 {object} pb.ListTransactionsResponse
// @Failure 400 {string} string "Bad Request"
// @Router /transaction/get [get]
func (h *BudgetingHandler) ListTransactions(ctx *gin.Context) {
	defaultLimit := 10
	defaultPage := 1

	limitStr := ctx.Query("limit")
	pageStr := ctx.Query("page")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = defaultLimit
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = defaultPage
	}

	req := &pb.ListTransactionsRequest{
		Limit: int32(limit),
		Page:  int32(page),
	}

	res, err := h.Transaction.ListTransactions(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// Spending godoc
// @Summary Get spending details
// @Description Get the count of spending transactions and the total amount spent.
// @Tags Transaction
// @Accept  json
// @Produce  json
// @Param user_id query string true "User ID"
// @Security BearerAuth
// @Success 200 {object} pb.SpendingResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "User Not Found"
// @Router /transaction/spending [get]
func (h *BudgetingHandler) Spending(ctx *gin.Context) {
	userId := ctx.Query("user_id")
	if userId == "" {
		ctx.JSON(http.StatusBadRequest, "User ID is required")
		return
	}

	req := &pb.SpendingRequest{
		UserId: userId,
	}

	res, err := h.Transaction.Spending(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// Income godoc
// @Summary Get income details
// @Description Get the count of income transactions and the total amount received.
// @Tags Transaction
// @Accept  json
// @Produce  json
// @Param user_id query string true "User ID"
// @Success 200 {object} pb.IncomeResponse
// @Security BearerAuth
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "User Not Found"
// @Router /transaction/income [get]
func (h *BudgetingHandler) Income(ctx *gin.Context) {
	userId := ctx.Query("user_id")
	if userId == "" {
		ctx.JSON(http.StatusBadRequest, "User ID is required")
		return
	}

	req := &pb.IncomeRequest{
		UserId: userId,
	}

	res, err := h.Transaction.Income(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, res)
}
