package handler

import (
	"fmt"
	"gin_news/pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	switch c.Request.Method {
	case http.MethodGet:
		c.HTML(http.StatusOK, "signUp.html", nil)
	case http.MethodPost:
		var input models.User
		if err := c.BindJSON(&input); err != nil {
			newErrorResponse(c, http.StatusBadRequest, "invalid input body")
			return
		}

		id, err := h.services.Authorization.CreateUser(input)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, map[string]interface{}{
			"id": id,
		})
		c.Redirect(http.StatusOK, "/auth/sign-in")
	}
}

func (h *Handler) signIn(c *gin.Context) {
	switch c.Request.Method {
	case http.MethodGet:
		c.HTML(http.StatusOK, "signIn.html", nil)
	case http.MethodPost:
		var input models.Login
		if err := c.BindJSON(&input); err != nil {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
		}

		token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		fmt.Println(token, "****")

		c.JSON(http.StatusOK, map[string]interface{}{
			"token": token,
		})
		c.Redirect(http.StatusSeeOther, "/api/news")
	}

}
