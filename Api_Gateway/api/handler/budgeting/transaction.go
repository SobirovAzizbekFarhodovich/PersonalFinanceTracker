package handler

import (
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
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	_, err := h.Transaction.CreateTransaction(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Transaction Create Successfully"})
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

	_, err := h.Transaction.UpdateTransaction(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Transaction Updated Successfully"})
}

// @Summary Delete Transaction
// @Description Delete Transaction
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
	req := &pb.DeleteTransactionRequest{Id: id}

	_, err := h.Transaction.DeleteTransaction(context.Background(), req)
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
