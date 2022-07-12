package handler

import (
	"net/http"
	"strconv"
	"undangan_online_api/formatter"
	"undangan_online_api/helper"
	"undangan_online_api/input"
	"undangan_online_api/service"

	"github.com/gin-gonic/gin"
)

type turutmengundangHandler struct {
	service service.TurutMengundangService
}

func NewTurutMengundangHandler(service service.TurutMengundangService) *turutmengundangHandler {
	return &turutmengundangHandler{service}
}
func (h *turutmengundangHandler) GetTurutMengundang(c *gin.Context) {
	var input input.InputIDTurutMengundang
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.ApiResponse("Failed to get TurutMengundang", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	turutmengundangDetail, err := h.service.TurutMengundangServiceGetByID(input)
	if err != nil {
		response := helper.ApiResponse("Failed to get TurutMengundang", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Detail TurutMengundang", http.StatusOK, "success", formatter.FormatTurutMengundang(turutmengundangDetail))
	c.JSON(http.StatusOK, response)
}

func (h *turutmengundangHandler) GetTurutMengundangs(c *gin.Context) {
	turutmengundangs, err := h.service.TurutMengundangServiceGetAll()
	if err != nil {
		response := helper.ApiResponse("Failed to get TurutMengundangs", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("List of TurutMengundangs", http.StatusOK, "success", formatter.FormatTurutMengundangs(turutmengundangs))
	c.JSON(http.StatusOK, response)
}

func (h *turutmengundangHandler) CreateTurutMengundang(c *gin.Context) {
	var input input.TurutMengundangInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.ApiResponse("Create TurutMengundang failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	newTurutMengundang, err := h.service.TurutMengundangServiceCreate(input)
	if err != nil {
		response := helper.ApiResponse("Create TurutMengundang failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Successfully Create TurutMengundang", http.StatusOK, "success", formatter.FormatTurutMengundang(newTurutMengundang))
	c.JSON(http.StatusOK, response)
}
func (h *turutmengundangHandler) UpdateTurutMengundang(c *gin.Context) {
	var inputID input.InputIDTurutMengundang
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.ApiResponse("Failed to get TurutMengundangs", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	var inputData input.TurutMengundangInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.ApiResponse("Update TurutMengundang failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	updatedTurutMengundang, err := h.service.TurutMengundangServiceUpdate(inputID, inputData)
	if err != nil {
		response := helper.ApiResponse("Failed to get TurutMengundangs", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Successfully Update TurutMengundang", http.StatusOK, "success", formatter.FormatTurutMengundang(updatedTurutMengundang))
	c.JSON(http.StatusOK, response)
}
func (h *turutmengundangHandler) DeleteTurutMengundang(c *gin.Context) {
	param := c.Param("id")
	id, _ := strconv.Atoi(param)
	var inputID input.InputIDTurutMengundang
	inputID.ID = id
	_, err := h.service.TurutMengundangServiceGetByID(inputID)
	if err != nil {
		response := helper.ApiResponse("Failed to get TurutMengundangs", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	_, err = h.service.TurutMengundangServiceDeleteByID(inputID)
	if err != nil {
		response := helper.ApiResponse("Failed to get TurutMengundangs", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Successfully Delete TurutMengundang", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}
