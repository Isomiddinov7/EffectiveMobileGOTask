package handler

import (
	"log"
	"strconv"
	"task/config"
	"task/pkg/logger"
	"task/storage"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	cfg    *config.Config
	strg   storage.StorageI
	logger logger.LoggerI
}

// Response - Json model response
type Response struct {
	Status      int         `json:"status"`
	Description string      `json:"description"`
	Data        interface{} `json:"data"`
}

func NewHandler(cfg *config.Config, strg storage.StorageI) *Handler {
	return &Handler{
		cfg:  cfg,
		strg: strg,
	}
}

func getIntegerOrDefaultValue(value string, defaultValue int64) (int64, error) {

	if len(value) <= 0 {
		return defaultValue, nil
	}

	number, err := strconv.Atoi(value)
	return int64(number), err
}

func handleResponse(c *gin.Context, status int, data interface{}) {
	var description string
	switch code := status; {
	case code < 400:
		description = "success"
	default:
		description = "error"
		log.Println(config.Error, "error while:", Response{
			Status:      status,
			Description: description,
			Data:        data,
		})

		if code == 500 {
			data = "Internal Server Error"
		}
	}

	c.JSON(status, Response{
		Status:      status,
		Description: description,
		Data:        data,
	})
}

func (h *Handler) getOffsetParam(c *gin.Context) (offset int, err error) {
	if h.cfg.DefaultOffset != "" {
		offsetStr := c.DefaultQuery("offset", h.cfg.DefaultOffset)
		return strconv.Atoi(offsetStr)
	}

	offsetStr := c.DefaultQuery("offset", "0")
	return strconv.Atoi(offsetStr)
}

func (h *Handler) getLimitParam(c *gin.Context) (offset int, err error) {
	if h.cfg.DefaultLimit != "" {
		limitStr := c.DefaultQuery("limit", h.cfg.DefaultLimit)
		return strconv.Atoi(limitStr)
	}
	limitStr := c.DefaultQuery("limit", "10")
	return strconv.Atoi(limitStr)
}
