package experimentService

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nedokyrill/ab-platform/internal/models"
	"net/http"
	"time"
)

// CreateExperiment создает новый эксперимент
func (s *ExperimentService) CreateExperiment(c *gin.Context) {
	var request struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description" binding:"required"`
		VariantA    string `json:"variant_a" binding:"required"`
		VariantB    string `json:"variant_b" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body: " + err.Error()})
		return
	}

	// Создаем новый эксперимент
	experiment := models.ExperimentModel{
		ExternalID:  uuid.New(),
		Name:        request.Name,
		Description: request.Description,
		VariantA:    request.VariantA,
		VariantB:    request.VariantB,
		CreatedAt:   time.Now(),
	}

	err := s.experimentRepo.CreateExperiment(experiment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create experiment"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":          experiment.ExternalID.String(),
		"name":        experiment.Name,
		"description": experiment.Description,
		"variant_a":   experiment.VariantA,
		"variant_b":   experiment.VariantB,
		"created_at":  experiment.CreatedAt,
	})
} 