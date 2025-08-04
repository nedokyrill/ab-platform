package experimentService

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetExperiments получает список всех экспериментов
func (s *ExperimentService) GetExperiments(c *gin.Context) {
	experiments, err := s.experimentRepo.GetAllExperiments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get experiments"})
		return
	}

	var response []gin.H
	for _, experiment := range *experiments {
		response = append(response, gin.H{
			"id":          experiment.ExternalID.String(),
			"name":        experiment.Name,
			"description": experiment.Description,
			"variant_a":   experiment.VariantA,
			"variant_b":   experiment.VariantB,
			"created_at":  experiment.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"experiments": response,
		"count":       len(response),
	})
} 