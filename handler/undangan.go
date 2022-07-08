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

type undanganHandler struct {
	service service.UndanganService
}

func NewUndanganHandler(service service.UndanganService) *undanganHandler {
	return &undanganHandler{service}
}

func (h *undanganHandler) GetUndangan(c *gin.Context) {
	var input input.InputIDUndangan
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.ApiResponse("Failed to get Undangan", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	undanganDetail, err := h.service.UndanganServiceGetByID(input)
	if err != nil {
		response := helper.ApiResponse("Failed to get Undangan", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Detail Undangan", http.StatusOK, "success", formatter.FormatUndangan(undanganDetail))
	c.JSON(http.StatusOK, response)
}

func (h *undanganHandler) GetUndangans(c *gin.Context) {
	undangans, err := h.service.UndanganServiceGetAll()
	if err != nil {
		response := helper.ApiResponse("Failed to get Undangans", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("List of Undangans", http.StatusOK, "success", formatter.FormatUndangans(undangans))
	c.JSON(http.StatusOK, response)
}

func (h *undanganHandler) CreateUndangan(c *gin.Context) {
	var input input.UndanganInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.ApiResponse("Create Undangan failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	newUndangan, err := h.service.UndanganServiceCreate(input)
	if err != nil {
		response := helper.ApiResponse("Create Undangan failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Successfully Create Undangan", http.StatusOK, "success", formatter.FormatUndangan(newUndangan))
	c.JSON(http.StatusOK, response)
}
func (h *undanganHandler) UpdateUndangan(c *gin.Context) {
	var inputID input.InputIDUndangan
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.ApiResponse("Failed to get Undangans", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	var inputData input.UndanganInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.ApiResponse("Update Undangan failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	updatedUndangan, err := h.service.UndanganServiceUpdate(inputID, inputData)
	if err != nil {
		response := helper.ApiResponse("Failed to get Undangans", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Successfully Update Undangan", http.StatusOK, "success", formatter.FormatUndangan(updatedUndangan))
	c.JSON(http.StatusOK, response)
}
func (h *undanganHandler) DeleteUndangan(c *gin.Context) {
	param := c.Param("id")
	id, _ := strconv.Atoi(param)
	var inputID input.InputIDUndangan
	inputID.ID = id
	_, err := h.service.UndanganServiceGetByID(inputID)
	if err != nil {
		response := helper.ApiResponse("Failed to get Undangans", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	_, err = h.service.UndanganServiceDeleteByID(inputID)
	if err != nil {
		response := helper.ApiResponse("Failed to get Undangans", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Successfully Delete Undangan", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}
