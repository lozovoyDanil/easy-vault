package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"main.go/internal/model"
)

type unitsResponse struct {
	Data []model.StorageUnit `json:"data"`
}

func (h *Handler) groupUnits(ctx *gin.Context) {
	id, err := getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	groupId, err := strconv.Atoi(ctx.Param("group_id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid group id")
		return
	}

	units, err := h.services.GroupUnits(id, groupId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, unitsResponse{
		Data: units,
	})
}

func (h *Handler) unitById(ctx *gin.Context) {
	id, err := getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	unitId, err := strconv.Atoi(ctx.Param("unit_id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid unit id")
		return
	}

	unit, err := h.services.UnitById(id, unitId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, unitsResponse{
		Data: []model.StorageUnit{unit},
	})
}

func (h *Handler) createUnit(ctx *gin.Context) {
	id, err := getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	groupId, err := strconv.Atoi(ctx.Query("group_id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid group id")
		return
	}

	var input model.StorageUnit
	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	unitId, err := h.services.CreateUnit(id, groupId, input)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id": unitId,
	})

}

func (h *Handler) updateUnit(ctx *gin.Context) {
	id, err := getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	unitId, err := strconv.Atoi(ctx.Param("unit_id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid unit id")
		return
	}

	var input model.UpdateUnitInput
	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.UpdateUnit(id, unitId, input)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, statusResp{
		Status: "OK",
	})
}

func (h *Handler) deleteUnit(ctx *gin.Context) {
	id, err := getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	unitId, err := strconv.Atoi(ctx.Param("unit_id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid unit id")
		return
	}

	err = h.services.DeleteUnit(id, unitId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, statusResp{
		Status: "OK",
	})
}

func (h *Handler) reservedUnits(ctx *gin.Context) {
	id, err := getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	units, err := h.services.ReservedUnits(id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, unitsResponse{
		Data: units,
	})
}

func (h *Handler) reserveUnit(ctx *gin.Context) {
	fmt.Println("Получили запрос")
}

func (h *Handler) unitDetails(ctx *gin.Context) {
	fmt.Println("Получили запрос")
}
