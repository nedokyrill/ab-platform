package experimentService

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

// GetExperimentByID получает эксперимент по ID
func (s *ExperimentService) GetExperimentByID(c *gin.Context) {
	experimentIDStr := c.Param("id")

	if experimentIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "experiment_id is required"})
		return
	}

	experimentID, err := uuid.Parse(experimentIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid experiment_id format"})
		return
	}

	// Получаем все эксперименты и ищем по ExternalID
	experiments, err := s.experimentRepo.GetAllExperiments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get experiments"})
		return
	}

	for _, experiment := range *experiments {
		if experiment.ExternalID == experimentID {
			c.JSON(http.StatusOK, gin.H{
				"id":          experiment.ExternalID.String(),
				"name":        experiment.Name,
				"description": experiment.Description,
				"variant_a":   experiment.VariantA,
				"variant_b":   experiment.VariantB,
				"created_at":  experiment.CreatedAt,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "experiment not found"})
} 