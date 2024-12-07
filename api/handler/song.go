package handler

import (
	"context"
	"net/http"
	"task/api/models"
	"task/config"
	"task/pkg/helpers"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CreateSong godoc
// @ID create_song
// @Router /song [POST]
// @Summary Create Song
// @Description Create Song
// @Tags Song
// @Accept json
// @Produce json
// @Param object body models.CreateSong true "CreateSongRequestBody"
// @Success 200 {object} Response{data=models.Song} "SongBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreateSong(c *gin.Context) {
	var createSong models.CreateSong

	err := c.ShouldBindJSON(&createSong)
	if err != nil {
		handleResponse(c, 400, "ShouldBindJSON err: "+err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()
	resp, err := h.strg.Song().Create(ctx, &createSong)
	if err != nil {
		h.logger.Error("Create Song error", zap.Error(err))
		return
	}

	handleResponse(c, http.StatusCreated, &resp)
}

// GetSong godoc
// @ID get_song
// @Router /song/{id} [GET]
// @Summary Get Song
// @Description Get Song
// @Tags Song
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "GetListSongResponseBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetSong(c *gin.Context) {
	var id = c.Param("id")
	if !helpers.IsValidUUID(id) {
		handleResponse(c, http.StatusBadRequest, "id is not uuid")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()

	resp, err := h.strg.Song().GetById(ctx, &models.SongPrimaryKey{Id: id})
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, http.StatusOK, resp)
}

// @Security ApiKeyAuth
// GetSongList godoc
// @ID get_song_list
// @Router /song [GET]
// @Summary Get Songs List
// @Description  Get Songs List
// @Tags Song
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Success 200 {object} models.GetSongResponse "GetAllSongResponseBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetSongList(c *gin.Context) {

	offset, err := h.getOffsetParam(c)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	limit, err := h.getLimitParam(c)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.strg.Song().GetAll(
		c.Request.Context(),
		&models.GetSongRequest{
			Limit:  int64(limit),
			Offset: int64(offset),
			Search: c.Query("search"),
		},
	)
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, http.StatusOK, resp)
}

// UpdateSong godoc
// @ID update_song
// @Router /song [PUT]
// @Summary Update Song
// @Description Update Song
// @Tags Song
// @Accept json
// @Produce json
// @Param object body models.UpdateSong true "UpdateSong"
// @Success 200 {object} Response{data=models.Song} "Song data"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdateSong(c *gin.Context) {

	var Song models.UpdateSong

	err := c.ShouldBindJSON(&Song)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.strg.Song().Update(
		c.Request.Context(),
		&Song,
	)
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, http.StatusOK, resp)
}

// DeleteSong godoc
// @ID delete_song
// @Router /song/{id} [DELETE]
// @Summary Delete Song
// @Description Delete Song
// @Tags Song
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=object{}} "Song data"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) DeleteSong(c *gin.Context) {

	SongId := c.Param("id")
	if !helpers.IsValidUUID(SongId) {
		handleResponse(c, http.StatusBadRequest, "id is not uuid")
		return
	}

	err := h.strg.Song().Delete(
		c.Request.Context(),
		&models.SongPrimaryKey{Id: SongId},
	)
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, http.StatusOK, "delete is success")
}
