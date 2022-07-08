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

type lovestoryHandler struct {
	service service.LoveStoryService
}

func NewLoveStoryHandler(service service.LoveStoryService) *lovestoryHandler {
	return &lovestoryHandler{service}
}
func (h *lovestoryHandler) GetLoveStory(c *gin.Context) {
	var input input.InputIDLoveStory
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.ApiResponse("Failed to get LoveStory", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	lovestoryDetail, err := h.service.LoveStoryServiceGetByID(input)
	if err != nil {
		response := helper.ApiResponse("Failed to get LoveStory", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Detail LoveStory", http.StatusOK, "success", formatter.FormatLoveStory(lovestoryDetail))
	c.JSON(http.StatusOK, response)
}

func (h *lovestoryHandler) GetLoveStorys(c *gin.Context) {
	params := c.Param("id")
	lovestorys, err := h.service.LoveStoryServiceGetAll(params)
	if err != nil {
		response := helper.ApiResponse("Failed to get LoveStorys", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("List of LoveStorys", http.StatusOK, "success", formatter.FormatLoveStorys(lovestorys))
	c.JSON(http.StatusOK, response)
}

func (h *lovestoryHandler) CreateLoveStory(c *gin.Context) {
	var input input.LoveStoryInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.ApiResponse("Create LoveStory failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	newLoveStory, err := h.service.LoveStoryServiceCreate(input)
	if err != nil {
		response := helper.ApiResponse("Create LoveStory failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Successfully Create LoveStory", http.StatusOK, "success", formatter.FormatLoveStory(newLoveStory))
	c.JSON(http.StatusOK, response)
}
func (h *lovestoryHandler) UpdateLoveStory(c *gin.Context) {
	var inputID input.InputIDLoveStory
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.ApiResponse("Failed to get LoveStorys", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	var inputData input.LoveStoryInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.ApiResponse("Update LoveStory failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	updatedLoveStory, err := h.service.LoveStoryServiceUpdate(inputID, inputData)
	if err != nil {
		response := helper.ApiResponse("Failed to get LoveStorys", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Successfully Update LoveStory", http.StatusOK, "success", formatter.FormatLoveStory(updatedLoveStory))
	c.JSON(http.StatusOK, response)
}
func (h *lovestoryHandler) DeleteLoveStory(c *gin.Context) {
	param := c.Param("id")
	id, _ := strconv.Atoi(param)
	var inputID input.InputIDLoveStory
	inputID.ID = id
	_, err := h.service.LoveStoryServiceGetByID(inputID)
	if err != nil {
		response := helper.ApiResponse("Failed to get LoveStorys", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	_, err = h.service.LoveStoryServiceDeleteByID(inputID)
	if err != nil {
		response := helper.ApiResponse("Failed to get LoveStorys", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Successfully Delete LoveStory", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}
