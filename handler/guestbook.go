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

type guestbookHandler struct {
	service service.GuestBookService
}

func NewGuestBookHandler(service service.GuestBookService) *guestbookHandler {
	return &guestbookHandler{service}
}
func (h *guestbookHandler) GetGuestBook(c *gin.Context) {
	var input input.InputIDGuestBook
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.ApiResponse("Failed to get GuestBook", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	guestbookDetail, err := h.service.GuestBookServiceGetByID(input)
	if err != nil {
		response := helper.ApiResponse("Failed to get GuestBook", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Detail GuestBook", http.StatusOK, "success", formatter.FormatGuestBook(guestbookDetail))
	c.JSON(http.StatusOK, response)
}

func (h *guestbookHandler) GetGuestBooks(c *gin.Context) {
	guestbooks, err := h.service.GuestBookServiceGetAll()
	if err != nil {
		response := helper.ApiResponse("Failed to get GuestBooks", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("List of GuestBooks", http.StatusOK, "success", formatter.FormatGuestBooks(guestbooks))
	c.JSON(http.StatusOK, response)
}

func (h *guestbookHandler) CreateGuestBook(c *gin.Context) {
	var input input.GuestBookInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.ApiResponse("Create GuestBook failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	newGuestBook, err := h.service.GuestBookServiceCreate(input)
	if err != nil {
		response := helper.ApiResponse("Create GuestBook failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Successfully Create GuestBook", http.StatusOK, "success", formatter.FormatGuestBook(newGuestBook))
	c.JSON(http.StatusOK, response)
}
func (h *guestbookHandler) UpdateGuestBook(c *gin.Context) {
	var inputID input.InputIDGuestBook
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.ApiResponse("Failed to get GuestBooks", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	var inputData input.GuestBookInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.ApiResponse("Update GuestBook failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	updatedGuestBook, err := h.service.GuestBookServiceUpdate(inputID, inputData)
	if err != nil {
		response := helper.ApiResponse("Failed to get GuestBooks", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Successfully Update GuestBook", http.StatusOK, "success", formatter.FormatGuestBook(updatedGuestBook))
	c.JSON(http.StatusOK, response)
}
func (h *guestbookHandler) DeleteGuestBook(c *gin.Context) {
	param := c.Param("id")
	id, _ := strconv.Atoi(param)
	var inputID input.InputIDGuestBook
	inputID.ID = id
	_, err := h.service.GuestBookServiceGetByID(inputID)
	if err != nil {
		response := helper.ApiResponse("Failed to get GuestBooks", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	_, err = h.service.GuestBookServiceDeleteByID(inputID)
	if err != nil {
		response := helper.ApiResponse("Failed to get GuestBooks", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Successfully Delete GuestBook", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}
