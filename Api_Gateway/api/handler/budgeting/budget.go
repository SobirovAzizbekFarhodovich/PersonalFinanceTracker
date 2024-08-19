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

// @Summary Create Budget
// @Description Create Budget
// @Tags Budget
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Create body pb.CreateBudgetRequest true "Create"
// @Success 201 {object} string "Success"
// @Failure 400 {string} string "Error"
// @Router /budget/create [post]
func (h *BudgetingHandler) CreateBudget(ctx *gin.Context) {
	req := &pb.CreateBudgetRequest{}
	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c := config.Load()
	id, _ := token.GetIdFromToken(ctx.Request, &c)
	req.Budget.UserId = id
	_, err := h.Budget.CreateBudget(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Budget Create Successfully"})
}

// @Summary Update Budget
// @Description Update Budget
// @Tags Budget
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Update body pb.UpdateBudgetRequest true "Update"
// @Success 200 {object} string "Success"
// @Failure 400 {string} string "Error"
// @Router /budget/update [put]
func (h *BudgetingHandler) UpdateBudget(ctx *gin.Context) {
	req := &pb.UpdateBudgetRequest{}
	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c := config.Load()
	id, _ := token.GetIdFromToken(ctx.Request, &c)
	req.Budget.UserId = id

	_, err := h.Budget.UpdateBudget(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Budget Updated Successfully"})
}

// @Summary Delete Budget
// @Description Delete Budget
// @Tags Budget
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Budget ID"
// @Success 200 {object} string "Success"
// @Failure 400 {string} string "Error"
// @Router /budget/delete/{id} [delete]
func (h *BudgetingHandler) DeleteBudget(ctx *gin.Context) {
	id := ctx.Param("id")
	req := &pb.DeleteBudgetRequest{Id: id}

	_, err := h.Budget.DeleteBudget(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Budget Deleted Successfully"})
}

// @Summary Get Budget
// @Description Get an existing Budget record by ID
// @Tags Budget
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Budget ID"
// @Success 200 {object} pb.GetBudgetResponse
// @Failure 400 {string} string "Error"
// @Router /budget/get/{id} [get]
func (h *BudgetingHandler) GetBudget(ctx *gin.Context) {
	id := ctx.Param("id")
	req := &pb.GetBudgetRequest{Id: id}

	res, err := h.Budget.GetBudget(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// @Summary ListBudgets
// @Description ListBudgets
// @Tags Budget
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param limit query int false "Limit"
// @Param page query int false "Page"
// @Success 200 {object} pb.ListBudgetsResponse
// @Failure 400 {string} string "Bad Request"
// @Router /budget/get [get]
func (h *BudgetingHandler) ListBudgets(ctx *gin.Context) {
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

	req := &pb.ListBudgetsRequest{
		Limit: int32(limit),
		Page:  int32(page),
	}

	res, err := h.Budget.ListBudgets(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// @Summary Generate Budget Performance Report
// @Description Generate a performance report for a specific budget. If the spent amount exceeds the budget amount, a notification will be created.
// @Tags Budget
// @Accept json
// @Produce json
// @Param id path string true "Budget ID"
// @Security BearerAuth
// @Success 200 {object} pb.GenerateBudgetPerformanceReportResponse
// @Failure 400 {string} string "Failed to generate budget performance report"
// @Failure 500 {string} string "Notification not created"
// @Router /budget/{id}/performance-report [get]
func (h *BudgetingHandler) GenerateBudgetPerformanceReport(ctx *gin.Context) {
	id := ctx.Param("id")
	req := &pb.GenerateBudgetPerformanceReportRequest{Id: id}

	res, err := h.Budget.GenerateBudgetPerformanceReport(context.Background(), req)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Failed to generate budget performance report"})
		return
	}
	ctx.JSON(http.StatusOK, res)
}
