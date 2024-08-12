package handler

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"regexp"
	"time"

	token "auth/api/token"
	"auth/config"
	pb "auth/genprotos"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type changePass struct {
	CurrentPassword string
	NewPassword     string
}

type resetPass struct {
	ResetToken  string
	NewPassword string
}

// RegisterUser handles the creation of a new user
// @Summary Register User
// @Description Register a new user
// @Tags User
// @Accept json
// @Produce json
// @Param Create body pb.RegisterUserRequest true "Create"
// @Success 201 {object} string "Create Successfully"
// @Failure 400 {string} string "Error while creating user"
// @Router /user/register [post]
func (h *Handler) RegisterUser(ctx *gin.Context) {
	user := pb.RegisterUserRequest{}
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	req, err := json.Marshal(&user)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	h.producer.ProduceMessages("user", req)

	ctx.JSON(http.StatusCreated, gin.H{"message": "User Create Successfully"})
}

// UpdateUser handles updating an existing user
// @Summary Update User
// @Description Update an existing user
// @Tags User
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param Update body pb.UpdateUserRequest true "Update"
// @Success 200 {object} string "Update Successful"
// @Failure 400 {string} string "Error while updating user"
// @Router /user/update/{id} [put]
func (h *Handler) UpdateUser(ctx *gin.Context) {
	user := pb.UpdateUserRequest{}
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	res, err := h.UserStorage.User().UpdateUser(&user)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, res)
}

// DeleteUser handles the deletion of a user
// @Summary Delete User
// @Description Delete an existing user
// @Tags User
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} string "Delete Successfully"
// @Failure 400 {string} string "Error while deleting user"
// @Router /user/delete/{id} [delete]
func (h *Handler) DeleteUser(ctx *gin.Context) {
	id := pb.DeleteUserRequest{Id: ctx.Param("id")}

	_, err := h.UserStorage.User().DeleteUser(&id)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, gin.H{"message": "User deleted successfully"})
}

// GetByIdUser handles retrieving a user by ID
// @Summary Get User By ID
// @Description Get a user by ID
// @Tags User
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "User ID"
// @Success 200 {object} pb.GetByIdUserResponse "Get By ID Successful"
// @Failure 400 {string} string "Error while retrieving user"
// @Failure 404 {string} string "User Not Found"
// @Router /user/get-by-id/{id} [get]
func (h *Handler) GetbyIdUser(ctx *gin.Context) {
	id := pb.GetByIdUserRequest{Id: ctx.Param("id")}

	res, err := h.UserStorage.User().GetByIdUser(&id)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			ctx.JSON(404, "User Not Found")
			return
		}
		ctx.JSON(400, "Error while retrieving user")
		return
	}

	ctx.JSON(200, res)
}

// LoginUser handles user login
// @Summary Login User
// @Description Login a user
// @Tags User
// @Accept json
// @Produce json
// @Param Create body pb.LoginUserRequest true "Create"
// @Success 200 {object} string "Login Successfully"
// @Failure 400 {string} string "Error while logging in"
// @Failure 404 {string} string "User Not Found"
// @Router /user/login [post]
func (h *Handler) LoginUser(ctx *gin.Context) {
	user := pb.LoginUserRequest{}
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	res, err := h.UserStorage.User().LoginUser(&user)
	if err != nil {
		if err.Error() == "invalid email or password" {
			ctx.JSON(404, "User Not Found or Invalid Password")
			return
		}
		ctx.JSON(400, err.Error())
		return
	}

	t := token.GenereteJWTToken(res)
	ctx.JSON(200, t)
}

// ChangePassword handles changing user password
// @Summary Change Password
// @Description Change user password
// @Tags User
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param ChangePass body changePass true "Change Password"
// @Success 200 {body} string "Password Changed Successfully"
// @Failure 400 {string} string "Error while changing password"
// @Router /user/change-password [post]
func (h *Handler) ChangePassword(ctx *gin.Context) {
	changePas := changePass{}
	err := ctx.BindJSON(&changePas)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	changePass := pb.ChangePasswordRequest{CurrentPassword: changePas.CurrentPassword, NewPassword: changePas.NewPassword}
	cnf := config.Load()
	id, _ := token.GetIdFromToken(ctx.Request, &cnf)
	changePass.Id = id

	_, err = h.UserStorage.User().ChangePassword(&changePass)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, gin.H{"message": "Password Changed Successfully"})
}

// ForgotPassword handles initiating the forgot password process
// @Summary Forgot Password
// @Description Initiate forgot password process
// @Tags User
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param ForgotPass body pb.ForgotPasswordRequest true "Forgot Password"
// @Success 200 {body} string "Forgot Password Initiated Successfully"
// @Failure 400 {string} string "Error while initiating forgot password"
// @Router /user/forgot-password [post]
func (h *Handler) ForgotPassword(ctx *gin.Context) {
	forgotPass := pb.ForgotPasswordRequest{}
	err := ctx.BindJSON(&forgotPass)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	// Validate email format
	if !isValidEmail(forgotPass.Email) {
		ctx.JSON(400, "Invalid email address")
		return
	}

	f := rand.Intn(899999) + 100000
	err = h.redis.SaveToken(forgotPass.Email, fmt.Sprintf("%d", f), time.Minute*2)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, gin.H{"message": "Email message has been sent"})
}

// isValidEmail validates the email format using a regex
func isValidEmail(email string) bool {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(regex)
	return re.MatchString(email)
}

// ResetPassword handles resetting the user password
// @Summary Reset Password
// @Description Reset user password
// @Tags User
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param ResetPass body resetPass true "Reset Password"
// @Success 200 {string} string "Password Reset Successfully"
// @Failure 400 {string} string "Error while resetting password"
// @Router /user/reset-password [post]
func (h *Handler) ResetPassword(ctx *gin.Context) {
	resetPas := resetPass{}
	err := ctx.BindJSON(&resetPas)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	resetPass := pb.ResetPasswordRequest{ResetToken: resetPas.ResetToken, PasswordHash: resetPas.NewPassword}
	cnf := config.Load()
	id, _ := token.GetIdFromToken(ctx.Request, &cnf)
	resetPass.Id = id

	email, _ := token.GetEmailFromToken(ctx.Request, &cnf)
	e, err := h.redis.Get(email)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	if e != resetPass.ResetToken {
		ctx.JSON(400, "Invalid reset-password")
		return
	}

	_, err = h.UserStorage.User().ResetPassword(&resetPass)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, "Password Reset Successfully")
}

// GetProfil handles retrieving a user Profil
// @Summary Get User Profil
// @Description Get a user Profil
// @Tags User
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} pb.GetByIdUserResponse "Get Profil Successful"
// @Failure 400 {string} string "Error while retrieving user"
// @Router /user/get-profil [get]
func (h *Handler) GetProfil(ctx *gin.Context) {
	cnf := config.Load()
	id, _ := token.GetIdFromToken(ctx.Request, &cnf)

	res, err := h.UserStorage.User().GetByIdUser(&pb.GetByIdUserRequest{Id: id})
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, res)
}

// UpdateUser handles updating an existing user
// @Summary Update Profil
// @Description Update an existing user
// @Tags User
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param Update body pb.UpdateUserRequest true "Update"
// @Success 200 {object} pb.UpdateUserResponse "Update Successful"
// @Failure 400 {string} string "Error while updating user"
// @Router /user/update-profil [put]
func (h *Handler) UpdateProfil(ctx *gin.Context) {
	user := pb.UpdateUserRequest{}
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	cnf := config.Load()
	user.Id, _ = token.GetIdFromToken(ctx.Request, &cnf)
	res, err := h.UserStorage.User().UpdateUser(&user)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, res)
}

// DeleteProfil handles the deletion of a Profil
// @Summary Delete Profil
// @Description Delete an existing Profil
// @Tags User
// @Accept json
// @Security BearerAuth
// @Produce json
// @Success 200 {string} string "Delete Successful"
// @Failure 400 {string} string "Error while deleting user"
// @Router /user/delete-profil [delete]
func (h *Handler) DeleteProfil(ctx *gin.Context) {
	cnf := config.Load()
	id, _ := token.GetIdFromToken(ctx.Request, &cnf)

	_, err := h.UserStorage.User().DeleteUser(&pb.DeleteUserRequest{Id: id})
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, "Success!!!")
}
