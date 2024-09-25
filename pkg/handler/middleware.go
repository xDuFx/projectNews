package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "id"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	fmt.Println(header)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	if len(headerParts[1]) == 0 {
		newErrorResponse(c, http.StatusUnauthorized, "token is empty")
		return
	}

	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userId)
}

func (h *Handler) checkAccess(c *gin.Context) {
	id, ok := c.Get(userCtx)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	idInt, ok := id.(int)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	check, err := h.services.Authorization.CheckAccess(idInt)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	if check {
		c.Next()
	}
	c.AbortWithStatus(http.StatusUnauthorized)
}
