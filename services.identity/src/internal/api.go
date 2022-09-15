package identity

import (
	"net/http"

	"github.com/faozimipa/micro/services.identity/src/entity"
	"github.com/faozimipa/micro/services.identity/src/jwt"
	"github.com/gin-gonic/gin"
)

func (h *Handler) health(context *gin.Context) {
	context.JSON(http.StatusOK, "Identity service is up!")
}

func (h *Handler) signUp(context *gin.Context) {

	var input *entity.User
	context.BindJSON(&input)

	user, err := h.service.SignUp(input)
	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}
	context.JSON(http.StatusOK, user)
}

func (h *Handler) signIn(context *gin.Context) {

	var input *entity.User
	context.BindJSON(&input)

	user, err := h.service.ValidateUser(input.Email, input.Password)
	if err != nil {
		context.JSON(http.StatusBadRequest, "Invalid Credentials.")
		return
	}

	token := jwt.GenerateToken(user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	context.JSON(http.StatusOK, jwt.TokenResponse{AccessToken: token})
}

func (h *Handler) login(context *gin.Context) {
	var input *entity.LoginCredential
	context.BindJSON(&input)
	token, err := h.service.Login(input.Username, input.Password)
	if err != nil {
		context.JSON(http.StatusBadRequest, "Invalid Credentials.")
		return
	}
	context.JSON(http.StatusOK, token)
}
