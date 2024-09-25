package handler

import (
	"gin_news/pkg/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	dst = "./static/image/"
)

func (h *Handler) createNews(c *gin.Context) {
	switch c.Request.Method {
	case http.MethodGet:
		c.HTML(http.StatusOK, "photo.html", gin.H{})
	case http.MethodPost:

		file, err := c.FormFile("file")

		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		// Upload the file to specific dst.
		var input models.News

		input.Title = c.PostForm("title")
		input.Body = c.PostForm("body")
		input.Mark = c.PostForm("mark")
		input.Reliz = c.PostForm("reliz")
		input.Image = dst + file.Filename
		err = c.SaveUploadedFile(file, input.Image)
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		id, err := h.services.Newslist.Create(input)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, map[string]interface{}{
			"id": id,
		})
	}
}

func (h *Handler) getAllNews(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
	items, err := h.services.Newslist.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *Handler) getNewsById(c *gin.Context) {
	newsId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	item, err := h.services.Newslist.GetByIdNews(newsId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *Handler) delete_news(c *gin.Context) {
	newsId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	err = h.services.Newslist.DeleteNews(newsId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) update_news(c *gin.Context) {
	var input models.UpdateNews
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.Newslist.UpdateNews(id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}
