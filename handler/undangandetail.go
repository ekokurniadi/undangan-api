package handler

import (
	"net/http"
	"strconv"
	"undangan_online_api/entity"
	"undangan_online_api/formatter"
	"undangan_online_api/helper"
	"undangan_online_api/input"
	"undangan_online_api/service"

	"github.com/gin-gonic/gin"
)

type undangandetailHandler struct {
	service service.UndanganDetailService
}

func NewUndanganDetailHandler(service service.UndanganDetailService) *undangandetailHandler {
	return &undangandetailHandler{service}
}
func (h *undangandetailHandler) GetUndanganDetail(c *gin.Context) {
	var input input.InputIDUndanganDetail
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.ApiResponse("Failed to get UndanganDetail", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	undangandetailDetail, err := h.service.UndanganDetailServiceGetByID(input)
	if err != nil {
		response := helper.ApiResponse("Failed to get UndanganDetail", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Detail UndanganDetail", http.StatusOK, "success", formatter.FormatUndanganDetail(undangandetailDetail))
	c.JSON(http.StatusOK, response)
}
func (h *undangandetailHandler) GetUndanganDetailByKode(c *gin.Context) {
	kode := c.Param("id")
	undangandetailDetail, err := h.service.UndanganDetailServiceGetByKode(kode)
	if err != nil {
		response := helper.ApiResponse("Failed to get UndanganDetail", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Detail UndanganDetail", http.StatusOK, "success", formatter.FormatUndanganDetails(undangandetailDetail))
	c.JSON(http.StatusOK, response)
}

func (h *undangandetailHandler) GetUndanganDetails(c *gin.Context) {
	undangandetails, err := h.service.UndanganDetailServiceGetAll()
	if err != nil {
		response := helper.ApiResponse("Failed to get UndanganDetails", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("List of UndanganDetails", http.StatusOK, "success", formatter.FormatUndanganDetails(undangandetails))
	c.JSON(http.StatusOK, response)
}

func (h *undangandetailHandler) CreateUndanganDetail(c *gin.Context) {
	var input []input.UndanganDetailInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.ApiResponse("Create UndanganDetail failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	var newUndanganDetail []entity.UndanganDetail
	for detail := range input {
		undangan, err := h.service.UndanganDetailServiceCreate(input[detail])
		if err != nil {
			response := helper.ApiResponse("Create UndanganDetail failed", http.StatusBadRequest, "error", nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}
		newUndanganDetail = append(newUndanganDetail, undangan)
	}

	response := helper.ApiResponse("Successfully Create UndanganDetail", http.StatusOK, "success", formatter.FormatUndanganDetails(newUndanganDetail))
	c.JSON(http.StatusOK, response)
}

func (h *undangandetailHandler) UpdateUndanganDetail(c *gin.Context) {
	var inputID input.InputIDUndanganDetail
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.ApiResponse("Failed to get UndanganDetails", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	var inputData input.UndanganDetailInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.ApiResponse("Update UndanganDetail failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	updatedUndanganDetail, err := h.service.UndanganDetailServiceUpdate(inputID, inputData)
	if err != nil {
		response := helper.ApiResponse("Failed to get UndanganDetails", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Successfully Update UndanganDetail", http.StatusOK, "success", formatter.FormatUndanganDetail(updatedUndanganDetail))
	c.JSON(http.StatusOK, response)
}
func (h *undangandetailHandler) DeleteUndanganDetail(c *gin.Context) {
	param := c.Param("id")
	id, _ := strconv.Atoi(param)
	var inputID input.InputIDUndanganDetail
	inputID.ID = id
	_, err := h.service.UndanganDetailServiceGetByID(inputID)
	if err != nil {
		response := helper.ApiResponse("Failed to get UndanganDetails", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	_, err = h.service.UndanganDetailServiceDeleteByID(inputID)
	if err != nil {
		response := helper.ApiResponse("Failed to get UndanganDetails", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Successfully Delete UndanganDetail", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}
