package controllers

import (
	"errors"
	"net/http"
	"strconv"
	dtos "smaash-web/internal/DTOs"
	"smaash-web/internal/models"
	"smaash-web/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LevelsController struct {
	levelService services.LevelService
}

type LevelCreateDTO struct {
	Name   string `json:"name" binding:"required,max=20"`
	ImgUri string `json:"img_uri" binding:"required"`
}

type LevelUpdateDTO struct {
	Name   string `json:"name" binding:"required,max=20"`
	ImgUri string `json:"img_uri" binding:"required"`
}

func NewLevelsController(levelService services.LevelService) *LevelsController {
	return &LevelsController{levelService: levelService}
}

func (lc *LevelsController) CreateLevel(c *gin.Context) {
	var body LevelCreateDTO
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, dtos.NewErrResp(err.Error(), c.Request.URL.Path))
		return
	}

	level := models.Level{
		Name:   body.Name,
		ImgUri: body.ImgUri,
	}

	if err := lc.levelService.Create(c.Request.Context(), &level); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			c.JSON(http.StatusBadRequest, dtos.NewErrResp("Level already exists", c.Request.URL.Path))
			return
		}
		c.JSON(http.StatusInternalServerError, dtos.NewErrResp(err.Error(), c.Request.URL.Path))
		return
	}

	c.JSON(http.StatusCreated, level)
}

func (lc *LevelsController) ReadAllLevels(c *gin.Context) {
	levels, err := lc.levelService.ReadAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.NewErrResp(err.Error(), c.Request.URL.Path))
		return
	}
	c.JSON(http.StatusOK, levels)
}

func (lc *LevelsController) ReadLevelByID(c *gin.Context) {
	id, err := parseUintParam(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, dtos.NewErrResp("Invalid id", c.Request.URL.Path))
		return
	}

	level, err := lc.levelService.ReadByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, dtos.NewErrResp("Level not found", c.Request.URL.Path))
			return
		}
		c.JSON(http.StatusInternalServerError, dtos.NewErrResp(err.Error(), c.Request.URL.Path))
		return
	}
	c.JSON(http.StatusOK, level)
}

func (lc *LevelsController) UpdateLevel(c *gin.Context) {
	id, err := parseUintParam(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, dtos.NewErrResp("Invalid id", c.Request.URL.Path))
		return
	}

	var body LevelUpdateDTO
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, dtos.NewErrResp(err.Error(), c.Request.URL.Path))
		return
	}

	level := models.Level{
		Model: gorm.Model{ID: id},
		Name:  body.Name,
		ImgUri: body.ImgUri,
	}

	if err := lc.levelService.Update(c.Request.Context(), &level); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, dtos.NewErrResp("Level not found", c.Request.URL.Path))
			return
		}
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			c.JSON(http.StatusBadRequest, dtos.NewErrResp("Level already exists", c.Request.URL.Path))
			return
		}
		c.JSON(http.StatusInternalServerError, dtos.NewErrResp(err.Error(), c.Request.URL.Path))
		return
	}

	c.Status(http.StatusNoContent)
}

func (lc *LevelsController) DeleteLevel(c *gin.Context) {
	id, err := parseUintParam(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, dtos.NewErrResp("Invalid id", c.Request.URL.Path))
		return
	}

	if err := lc.levelService.Delete(c.Request.Context(), id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, dtos.NewErrResp("Level not found", c.Request.URL.Path))
			return
		}
		c.JSON(http.StatusInternalServerError, dtos.NewErrResp(err.Error(), c.Request.URL.Path))
		return
	}

	c.Status(http.StatusNoContent)
}

func parseUintParam(c *gin.Context, name string) (uint, error) {
	idStr := c.Param(name)
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}
