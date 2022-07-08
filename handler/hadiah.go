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

type hadiahHandler struct {
	service service.HadiahService
}

func NewHadiahHandler(service service.HadiahService) *hadiahHandler {
	return &hadiahHandler{service}
}
func (h *hadiahHandler) GetHadiah(c *gin.Context) {
	var input input.InputIDHadiah
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.ApiResponse("Failed to get Hadiah", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	hadiahDetail, err := h.service.HadiahServiceGetByID(input)
	if err != nil {
		response := helper.ApiResponse("Failed to get Hadiah", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Detail Hadiah", http.StatusOK, "success", formatter.FormatHadiah(hadiahDetail))
	c.JSON(http.StatusOK, response)
}

func (h *hadiahHandler) GetHadiahs(c *gin.Context) {
	hadiahs, err := h.service.HadiahServiceGetAll()
	if err != nil {
		response := helper.ApiResponse("Failed to get Hadiahs", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("List of Hadiahs", http.StatusOK, "success", formatter.FormatHadiahs(hadiahs))
	c.JSON(http.StatusOK, response)
}

func (h *hadiahHandler) CreateHadiah(c *gin.Context) {
	var input input.HadiahInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.ApiResponse("Create Hadiah failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	newHadiah, err := h.service.HadiahServiceCreate(input)
	if err != nil {
		response := helper.ApiResponse("Create Hadiah failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Successfully Create Hadiah", http.StatusOK, "success", formatter.FormatHadiah(newHadiah))
	c.JSON(http.StatusOK, response)
}
func (h *hadiahHandler) UpdateHadiah(c *gin.Context) {
	var inputID input.InputIDHadiah
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.ApiResponse("Failed to get Hadiahs", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	var inputData input.HadiahInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.ApiResponse("Update Hadiah failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	updatedHadiah, err := h.service.HadiahServiceUpdate(inputID, inputData)
	if err != nil {
		response := helper.ApiResponse("Failed to get Hadiahs", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Successfully Update Hadiah", http.StatusOK, "success", formatter.FormatHadiah(updatedHadiah))
	c.JSON(http.StatusOK, response)
}
func (h *hadiahHandler) DeleteHadiah(c *gin.Context) {
	param := c.Param("id")
	id, _ := strconv.Atoi(param)
	var inputID input.InputIDHadiah
	inputID.ID = id
	_, err := h.service.HadiahServiceGetByID(inputID)
	if err != nil {
		response := helper.ApiResponse("Failed to get Hadiahs", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	_, err = h.service.HadiahServiceDeleteByID(inputID)
	if err != nil {
		response := helper.ApiResponse("Failed to get Hadiahs", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Successfully Delete Hadiah", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}
