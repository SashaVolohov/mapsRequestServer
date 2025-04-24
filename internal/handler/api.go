package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type BaseResponse struct {
	Status string
}

type GetResponse struct {
	Status string
	Data   string
}

func (h *Handler) createValueByKey(c *gin.Context) {

	currentTime := time.Now()

	key := c.Param("key")
	value := c.Param("value")
	lifeTime := c.Param("lifeTime")

	if key == "" || value == "" || lifeTime == "" {
		newErrorResponse(c, http.StatusBadRequest, viper.GetString("invalidRequestErrorText"))
		return
	}

	lifeTimeInteger, err := strconv.Atoi(lifeTime)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid request - lifeTime is not integer! - %s", err.Error()))
		return
	}

	currentTime = currentTime.Add(time.Second * time.Duration(lifeTimeInteger))
	err = h.services.API.CreateValueByKey(key, value, currentTime)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, BaseResponse{
		Status: "OK",
	})

}

func (h *Handler) getValueByKey(c *gin.Context) {

	key := c.Param("key")

	if key == "" {
		newErrorResponse(c, http.StatusBadRequest, viper.GetString("invalidRequestErrorText"))
		return
	}

	value, err := h.services.API.GetValueByKey(key)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, GetResponse{
		Data:   value,
		Status: "OK",
	})

}

func (h *Handler) deleteValueByKey(c *gin.Context) {

	key := c.Param("key")

	if key == "" {
		newErrorResponse(c, http.StatusBadRequest, viper.GetString("invalidRequestErrorText"))
		return
	}

	err := h.services.API.DeleteValueByKey(key)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, BaseResponse{
		Status: "OK",
	})

}
