package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hendrasan/go-dhammapada-api/internal/models"
	"gorm.io/gorm"
)

type SearchHandler struct {
	DB *gorm.DB
}

func NewSearchHandler(db *gorm.DB) *SearchHandler {
	return &SearchHandler{
		DB: db,
	}
}

func (h *SearchHandler) Search(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Query is required"})
		return
	}

	result, err := h.searchAll(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error searching"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

func (h *SearchHandler) searchAll(query string) (*models.SearchResponse, error) {
	var chapters []models.Chapter
	var verses []models.Verse

	// Search in chapters
	result := h.DB.Where("title ILIKE ? OR english_title ILIKE ?", "%"+query+"%", "%"+query+"%").Find(&chapters)
	if result.Error != nil {
		return nil, result.Error
	}

	// Search in verses
	result = h.DB.Where("text ILIKE ? OR english_text ILIKE ? OR story_title ILIKE ? OR english_story_title ILIKE ? OR story ILIKE ? OR english_story ILIKE ?", "%"+query+"%", "%"+query+"%", "%"+query+"%", "%"+query+"%", "%"+query+"%", "%"+query+"%").Find(&verses)
	if result.Error != nil {
		return nil, result.Error
	}

	return &models.SearchResponse{
		Chapters: chapters,
		Verses:   verses,
	}, nil
}
