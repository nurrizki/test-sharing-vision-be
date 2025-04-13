package handler

import (
	"log"
	"net/http"
	"strconv"
	"test-sharing-vision/constant"
	"test-sharing-vision/service/controller"
	"test-sharing-vision/service/model"
	"test-sharing-vision/service/model/request"
	"test-sharing-vision/util"

	"github.com/gin-gonic/gin"
)

func InsertData(c *gin.Context) {

	var request request.InsertPosts
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Status:  constant.ResponseStatusFailed,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	err = util.ValidateRequest(request)
	if err != nil {
		log.Println("[InsertData] ValidateRequest - Error:", err.Error())
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Status:  constant.ResponseStatusFailed,
			Message: err.Error(),
		})
		return
	}

	// controller
	resp, err := controller.InsertData(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.GeneralResponse{
			Status:  constant.ResponseStatusFailed,
			Message: "Error when InsertData, Error: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}
func GetData(c *gin.Context) {

	var request request.GetPosts

	request.Title = c.Query("title")
	request.Content = c.Query("content")
	request.Category = c.Query("category")
	request.Status = c.Query("status")
	limitStr := c.Query("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10
	}
	request.Limit = limit

	offsetStr := c.Query("offset")
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		offset = 0
	}
	request.Offset = offset

	// set default limit
	if request.Limit < 1 {
		request.Limit = 10
	}

	// controller
	resp, err := controller.GetData(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.GeneralResponse{
			Status:  constant.ResponseStatusFailed,
			Message: "Error when Getdata, Error: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func GetDataById(c *gin.Context) {
	//get customerType from query param
	id := c.Param("id")

	// controller
	resp, err := controller.GetDataById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.GeneralResponse{
			Status:  constant.ResponseStatusFailed,
			Message: "Error when GetDataById, Error: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func DeleteDataById(c *gin.Context) {
	//get customerType from query param
	id := c.Param("id")

	// controller
	resp, err := controller.DeleteData(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.GeneralResponse{
			Status:  constant.ResponseStatusFailed,
			Message: "Error when DeleteDataById, Error: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func UpdateDataById(c *gin.Context) {
	//get customerType from query param
	id := c.Param("id")

	var request request.UpdatePosts
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Status:  constant.ResponseStatusFailed,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	request.Id = id

	err = util.ValidateRequest(request)
	if err != nil {
		log.Println("[UpdateDataById] ValidateRequest - Error:", err.Error())
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Status:  constant.ResponseStatusFailed,
			Message: err.Error(),
		})
		return
	}

	// controller
	resp, err := controller.UpdateData(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.GeneralResponse{
			Status:  constant.ResponseStatusFailed,
			Message: "Error when UpdateDataById, Error: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}
