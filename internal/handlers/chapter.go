package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hendrasan/go-dhammapada-api/internal/models"
	"gorm.io/gorm"
)

type ChapterHandler struct {
	DB *gorm.DB
}

func NewChapterHandler(db *gorm.DB) *ChapterHandler {
	return &ChapterHandler{
		DB: db,
	}
}

func (h *ChapterHandler) GetChapters(c *gin.Context) {
	var chapters []models.Chapter

	query := h.DB.Order("number asc")

	// Check if search query exists
	if searchQuery := c.Query("q"); searchQuery != "" {
		query = query.Where(
			"title ILIKE ? OR english_title ILIKE ?",
			"%"+searchQuery+"%",
			"%"+searchQuery+"%",
		)
	}

	result := query.Find(&chapters)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching chapters"})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusOK, gin.H{
			"data":    []models.Chapter{},
			"message": "No chapters found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": chapters,
	})
}

func (h *ChapterHandler) GetChapterByID(c *gin.Context) {
	var chapter models.Chapter
	var verses []models.Verse

	id := c.Param("id")

	// Eager load verses using Preload
	result := h.DB.Preload("Verses.Chapter").First(&chapter, id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Chapter not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching chapter"})
		return
	}

	verses = append(verses, chapter.Verses...)

	c.JSON(http.StatusOK, gin.H{
		"data":   chapter,
		"verses": verses,
	})
}
