package handler

import (
	pb "api/genprotos/budgeting"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Create Account
// @Description Create Account
// @Tags Account
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Create body pb.CreateAccountRequest true "Create"
// @Success 201 {object} string "Success"
// @Failure 400 {string} string "Error"
// @Router /account/create [post]
func (h *BudgetingHandler) CreateAccount(ctx *gin.Context) {
	req := &pb.CreateAccountRequest{}
	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	_, err := h.Account.CreateAccount(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Account Create Successfully"})
}

// @Summary Update Account
// @Description Update Account
// @Tags Account
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Update body pb.UpdateAccountRequest true "Update"
// @Success 200 {object} string "Success"
// @Failure 400 {string} string "Error"
// @Router /account/update [put]
func (h *BudgetingHandler) UpdateAccount(ctx *gin.Context) {
	req := &pb.UpdateAccountRequest{}
	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	_, err := h.Account.UpdateAccount(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Account Updated Successfully"})
}

// @Summary Delete Account
// @Description Delete Account
// @Tags Account
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Account ID"
// @Success 200 {object} string "Success"
// @Failure 400 {string} string "Error"
// @Router /account/delete/{id} [delete]
func (h *BudgetingHandler) DeleteAccount(ctx *gin.Context) {
	id := ctx.Param("id")
	req := &pb.DeleteAccountRequest{Id: id}

	_, err := h.Account.DeleteAccount(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Account Deleted Successfully"})
}

// @Summary Get Account
// @Description Get an existing Account record by ID
// @Tags Account
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Account ID"
// @Success 200 {object} pb.GetAccountResponse
// @Failure 400 {string} string "Error"
// @Router /account/get/{id} [get]
func (h *BudgetingHandler) GetAccount(ctx *gin.Context) {
	id := ctx.Param("id")
	req := &pb.GetAccountRequest{Id: id}

	res, err := h.Account.GetAccount(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// @Summary ListAccounts
// @Description ListAccounts
// @Tags Account
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param limit query int false "Limit"
// @Param page query int false "Page"
// @Success 200 {object} pb.ListAccountsResponse
// @Failure 400 {string} string "Bad Request"
// @Router /account/get [get]
func (h *BudgetingHandler) ListAccounts(ctx *gin.Context) {
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

	req := &pb.ListAccountsRequest{
		Limit: int32(limit),
		Page:  int32(page),
	}

	res, err := h.Account.ListAccounts(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, res)
}
