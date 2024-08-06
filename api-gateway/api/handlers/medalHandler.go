package handlers

import (
	"fmt"

	pb "github.com/Bekzodbekk/protofiles/genproto/medals"

	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func MedalTypeFromString(s string) (pb.MedalType, error) {

	switch strings.ToUpper(s) {
	case "GOLD":
		return pb.MedalType_GOLD, nil
	case "SILVER":
		return pb.MedalType_SILVER, nil
	case "BRONZE":
		return pb.MedalType_BRONZE, nil
	default:
		return pb.MedalType(0), fmt.Errorf("invalid medal type: %s", s)
	}
}

func (h *Handlers) CreateMedal(ctx *gin.Context) {
	var req struct {
		CountryId string `json:"country_id"`
		Type      string `json:"type"`
		EventId   string `json:"event_id"`
		AthleteId string `json:"athlete_id"`
	}

	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	medalType, err := MedalTypeFromString(req.Type)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createReq := pb.CreateMedalRequest{
		CountryId: req.CountryId,
		Type:      medalType,
		EventId:   req.EventId,
		AthleteId: req.AthleteId,
	}

	resp, err := h.MedalService.CreateMedal(ctx, &createReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (h *Handlers) UpdateMedal(ctx *gin.Context) {
	var req struct {
		CountryId string `json:"country_id"`
		Type      string `json:"type"`
		EventId   string `json:"event_id"`
		AthleteId string `json:"athlete_id"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	medalType, err := MedalTypeFromString(req.Type)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateReq := pb.UpdateMedalRequest{
		Id:        ctx.Param("id"),
		CountryId: req.CountryId,
		Type:      medalType,
		EventId:   req.EventId,
		AthleteId: req.AthleteId,
	}

	resp, err := h.MedalService.UpdateMedal(ctx, &updateReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (h *Handlers) DeleteMedal(ctx *gin.Context) {
	id := ctx.Param("id")
	req := pb.DeleteMedalRequest{Id: id}
	resp, err := h.MedalService.DeleteMedal(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}
