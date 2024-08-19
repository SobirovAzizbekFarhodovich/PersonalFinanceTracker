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

// @Summary Create Goal
// @Description Create Goal
// @Tags Goal
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Create body pb.CreateGoalRequest true "Create"
// @Success 201 {object} string "Success"
// @Failure 400 {string} string "Error"
// @Router /goal/create [post]
func (h *BudgetingHandler) CreateGoal(ctx *gin.Context) {

	req := &pb.CreateGoalRequest{}
	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c := config.Load()
	id, _ := token.GetIdFromToken(ctx.Request, &c)
	req.Goal.UserId = id

	_, err := h.Goal.CreateGoal(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Goal Create Successfully"})
}

// @Summary Update Goal
// @Description Update Goal
// @Tags Goal
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Update body pb.UpdateGoalRequest true "Update"
// @Success 200 {object} string "Success"
// @Failure 400 {string} string "Error"
// @Router /goal/update [put]
func (h *BudgetingHandler) UpdateGoal(ctx *gin.Context) {
	req := &pb.UpdateGoalRequest{}
	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
    c := config.Load()
	id, _ := token.GetIdFromToken(ctx.Request, &c)
	req.Goal.UserId = id
	_, err := h.Goal.UpdateGoal(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Goal Updated Successfully"})
}

// @Summary Delete Goal
// @Description Delete Goal
// @Tags Goal
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Goal ID"
// @Success 200 {object} string "Success"
// @Failure 400 {string} string "Error"
// @Router /goal/delete/{id} [delete]
func (h *BudgetingHandler) DeleteGoal(ctx *gin.Context) {
	id := ctx.Param("id")
	req := &pb.DeleteGoalRequest{Id: id}

	_, err := h.Goal.DeleteGoal(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Goal Deleted Successfully"})
}

// @Summary Get Goal
// @Description Get an existing Goal record by ID
// @Tags Goal
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Goal ID"
// @Success 200 {object} pb.GetGoalResponse
// @Failure 400 {string} string "Error"
// @Router /goal/get/{id} [get]
func (h *BudgetingHandler) GetGoal(ctx *gin.Context) {
	id := ctx.Param("id")

	req := &pb.GetGoalRequest{Id: id}
	res, err := h.Goal.GetGoal(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "goal not found"})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// @Summary List Goals
// @Description List all goals with pagination
// @Tags Goal
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param limit query int false "Limit"
// @Param page query int false "Page"
// @Success 200 {object} pb.ListGoalsResponse
// @Failure 400 {string} string "Bad Request"
// @Router /goal/get [get]
func (h *BudgetingHandler) ListGoals(ctx *gin.Context) {
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

	req := &pb.ListGoalsRequest{
		Limit: int32(limit),
		Page:  int32(page),
	}

	res, err := h.Goal.ListGoals(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to list goals"})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// @Summary Generate Goal Progress Report
// @Description Generate a progress report for a specific goal by ID
// @Tags Goal
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Goal ID"
// @Success 200 {object} pb.GenerateGoalProgressReportResponse
// @Failure 400 {string} string "Error"
// @Router /goal/getprogress/{id} [get]
func (h *BudgetingHandler) GenerateGoalProgressReport(ctx *gin.Context) {
	id := ctx.Param("id")
	req := &pb.GenerateGoalProgressReportRequest{Id: id}

	res, err := h.Goal.GenerateGoalProgressReport(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to generate goal progress report"})
		return
	}
	ctx.JSON(http.StatusOK, res)
}
