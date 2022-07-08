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

type galeriHandler struct {
	service service.GaleriService
}

func NewGaleriHandler(service service.GaleriService) *galeriHandler {
	return &galeriHandler{service}
}
func (h *galeriHandler) GetGaleri(c *gin.Context) {
	var input input.InputIDGaleri
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.ApiResponse("Failed to get Galeri", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	galeriDetail, err := h.service.GaleriServiceGetByID(input)
	if err != nil {
		response := helper.ApiResponse("Failed to get Galeri", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Detail Galeri", http.StatusOK, "success", formatter.FormatGaleri(galeriDetail))
	c.JSON(http.StatusOK, response)
}

func (h *galeriHandler) GetGaleris(c *gin.Context) {
	params := c.Param("id")
	galeris, err := h.service.GaleriServiceGetAll(params)
	if err != nil {
		response := helper.ApiResponse("Failed to get Galeris", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("List of Galeris", http.StatusOK, "success", formatter.FormatGaleris(galeris))
	c.JSON(http.StatusOK, response)
}

func (h *galeriHandler) CreateGaleri(c *gin.Context) {
	var input input.GaleriInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.ApiResponse("Create Galeri failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	newGaleri, err := h.service.GaleriServiceCreate(input)
	if err != nil {
		response := helper.ApiResponse("Create Galeri failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Successfully Create Galeri", http.StatusOK, "success", formatter.FormatGaleri(newGaleri))
	c.JSON(http.StatusOK, response)
}
func (h *galeriHandler) UpdateGaleri(c *gin.Context) {
	var inputID input.InputIDGaleri
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.ApiResponse("Failed to get Galeris", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	var inputData input.GaleriInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.ApiResponse("Update Galeri failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	updatedGaleri, err := h.service.GaleriServiceUpdate(inputID, inputData)
	if err != nil {
		response := helper.ApiResponse("Failed to get Galeris", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Successfully Update Galeri", http.StatusOK, "success", formatter.FormatGaleri(updatedGaleri))
	c.JSON(http.StatusOK, response)
}
func (h *galeriHandler) DeleteGaleri(c *gin.Context) {
	param := c.Param("id")
	id, _ := strconv.Atoi(param)
	var inputID input.InputIDGaleri
	inputID.ID = id
	_, err := h.service.GaleriServiceGetByID(inputID)
	if err != nil {
		response := helper.ApiResponse("Failed to get Galeris", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	_, err = h.service.GaleriServiceDeleteByID(inputID)
	if err != nil {
		response := helper.ApiResponse("Failed to get Galeris", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Successfully Delete Galeri", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}
