package field

import (
	"Be-Book-Padel/database"
	"Be-Book-Padel/models"
	"Be-Book-Padel/models/repository/field"
	fieldServices "Be-Book-Padel/models/service/field"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	FieldRepository = field.FieldRepository{}
	FieldService    = fieldServices.NewFieldService{
		FieldRepo: FieldRepository,
	}
)

type FieldHandler struct {
	DB *gorm.DB
}

func NewFieldHandler() *FieldHandler {
	return &FieldHandler{DB: database.DB}
}

func (h *FieldHandler) Create(c *gin.Context) {
	var req struct {
		Name          string `json:"name" binding:"required"`
		Length        int    `json:"length" binding:"required"`
		Width         int    `json:"width" binding:"required"`
		Description   string `json:"description" binding:"required"`
		IsActive      bool   `json:"is_active" binding:"required"`
		FieldPricings []struct {
			Price     float64        `json:"price" binding:"required"`
			StartTime string         `json:"start_time" binding:"required"`
			EndTime   string         `json:"end_time" binding:"required"`
			DayType   models.DayType `json:"day_type" binding:"required,oneof=weekday weekend"`
		} `json:"field_pricings" binding:"required,dive"`
	}

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error":   "invalid request",
			"details": err.Error(),
		})
		return
	}

	var fieldPricings []models.FieldPricing
	for _, fp := range req.FieldPricings {
		startTime, err := time.Parse("02-09-2006 15:04", fp.StartTime)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid start_time format, use dd-mm-yyy HH:MM"})
			return
		}
		endTime, err := time.Parse("02-09-2006 15:04", fp.EndTime)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid end_time format, use dd-mm-yyy HH:MM"})
			return
		}

		duration := endTime.Sub(startTime)
		minutes := int(duration.Minutes())

		fieldPricings = append(fieldPricings, models.FieldPricing{
			Price:           fp.Price,
			DurationMinutes: minutes,
			StartTime:       startTime,
			EndTime:         endTime,
			DayType:         fp.DayType,
		})
	}

	fieldModel := &models.Field{
		Name:          req.Name,
		Length:        req.Length,
		Width:         req.Width,
		Description:   req.Description,
		IsActive:      req.IsActive,
		FieldPricings: fieldPricings,
	}

	id, err := FieldService.Create(fieldModel)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{
		"id":      id,
		"message": "field created successfully",
	})
}
