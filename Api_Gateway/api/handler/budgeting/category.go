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

// @Summary Create Category
// @Description Create Category
// @Tags Category
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Create body pb.CreateCategoryRequest true "Create"
// @Success 201 {object} string "Success"
// @Failure 400 {string} string "Error"
// @Router /category/create [post]
func (h *BudgetingHandler) CreateCategory(ctx *gin.Context) {
	req := &pb.CreateCategoryRequest{}
	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c := config.Load()
	id, _ := token.GetIdFromToken(ctx.Request, &c)
	req.Category.UserId = id

	_, err := h.Category.CreateCategory(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Category Create Successfully"})
}

// @Summary Update Category
// @Description Update Category
// @Tags Category
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Update body pb.UpdateCategoryRequest true "Update"
// @Success 200 {object} string "Success"
// @Failure 400 {string} string "Error"
// @Router /category/update [put]
func (h *BudgetingHandler) UpdateCategory(ctx *gin.Context) {
	req := &pb.UpdateCategoryRequest{}
	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	_, err := h.Category.UpdateCategory(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c := config.Load()
	id, _ := token.GetIdFromToken(ctx.Request, &c)
	req.Category.UserId = id

	ctx.JSON(http.StatusOK, gin.H{"message": "Category Updated Successfully"})
}

// @Summary Delete Category
// @Description Delete Category
// @Tags Category
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Category ID"
// @Success 200 {object} string "Success"
// @Failure 400 {string} string "Error"
// @Router /category/delete/{id} [delete]
func (h *BudgetingHandler) DeleteCategory(ctx *gin.Context) {
	id := ctx.Param("id")
	req := &pb.DeleteCategoryRequest{Id: id}

	_, err := h.Category.DeleteCategory(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Category Deleted Successfully"})
}

// @Summary Get Category
// @Description Get an existing Category record by ID
// @Tags Category
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Category ID"
// @Success 200 {object} pb.GetCategoryResponse
// @Failure 400 {string} string "Error"
// @Router /category/get/{id} [get]
func (h *BudgetingHandler) GetCategory(ctx *gin.Context) {
	id := ctx.Param("id")
	req := &pb.GetCategoryRequest{Id: id}

	res, err := h.Category.GetCategory(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// @Summary ListCategories
// @Description ListCategories
// @Tags Category
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param limit query int false "Limit"
// @Param page query int false "Page"
// @Success 200 {object} pb.ListCategoriesResponse
// @Failure 400 {string} string "Bad Request"
// @Router /category/get [get]
func (h *BudgetingHandler) ListCategories(ctx *gin.Context) {
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

	req := &pb.ListCategoriesRequest{
		Limit: int32(limit),
		Page:  int32(page),
	}

	res, err := h.Category.ListCategories(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, res)
}
