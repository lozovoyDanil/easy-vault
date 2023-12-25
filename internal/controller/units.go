package controller

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"main.go/internal/model"
	"main.go/internal/service"
)

type unitsResponse struct {
	Data []model.StorageUnit `json:"data"`
}

func (h *Handler) groupUnits(ctx *gin.Context) {
	user, err := getUserIdentity(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	groupId, err := strconv.Atoi(ctx.Query("group_id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid group id")
		return
	}

	units, err := h.services.GroupUnits(*user, groupId)
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
	user, err := getUserIdentity(ctx)
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

	unitId, err := h.services.CreateUnit(*user, groupId, input)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id": unitId,
	})

}

func (h *Handler) updateUnit(ctx *gin.Context) {
	user, err := getUserIdentity(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	unitId, err := strconv.Atoi(ctx.Param("unit_id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid unit id")
		return
	}

	var input model.UnitInput
	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.UpdateUnit(*user, unitId, input)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, statusResp{
		Status: "OK",
	})
}

func (h *Handler) deleteUnit(ctx *gin.Context) {
	user, err := getUserIdentity(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	unitId, err := strconv.Atoi(ctx.Param("unit_id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid unit id")
		return
	}

	err = h.services.DeleteUnit(*user, unitId)
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

func (h *Handler) unitDetails(ctx *gin.Context) {
	user, err := getUserIdentity(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	unitId, err := strconv.Atoi(ctx.Param("unit_id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid unit id")
		return
	}

	unit, err := h.services.UnitDetails(*user, unitId)
	if errors.Is(err, service.ErrOwnershipViolation) {
		newErrorResponse(ctx, http.StatusForbidden, err.Error())
		return
	}
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, unit)
}

func (h *Handler) reserveUnit(ctx *gin.Context) {
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

	var reservInfo model.UnitInput
	if err := ctx.BindJSON(&reservInfo); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.ReserveUnit(id, unitId, reservInfo)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, statusResp{
		Status: "OK",
	})
}

func (h *Handler) cancelReserv(ctx *gin.Context) {

}

func (h *Handler) lockUnit(ctx *gin.Context) {
	fmt.Println("Получили запрос")
}
func (h *Handler) unlockUnit(ctx *gin.Context) {
	fmt.Println("Получили запрос")
}
