package api

import (
	"task/api/handler"
	"task/config"
	"task/storage"

	_ "task/api/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetUpApi(r *gin.Engine, cfg *config.Config, strg storage.StorageI) {
	handler := handler.NewHandler(cfg, strg)

	r.POST("/song", handler.CreateSong)
	r.GET("/song/:id", handler.GetSong)
	r.GET("/song", handler.GetSongList)
	r.PUT("/song", handler.UpdateSong)
	r.DELETE("/song/:id", handler.DeleteSong)

	url := ginSwagger.URL("swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
