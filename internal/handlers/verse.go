package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hendrasan/go-dhammapada-api/internal/models"
	"gorm.io/gorm"
)

type VerseHandler struct {
	DB *gorm.DB
}

func NewVerseHandler(db *gorm.DB) *VerseHandler {
	return &VerseHandler{
		DB: db,
	}
}

func (h *VerseHandler) GetVerses(c *gin.Context) {
	var verses []models.Verse

	// get pagination query parameters
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("page_size", "10")

	// convert query parameters to int
	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt < 1 {
		pageInt = 1
	}
	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil || pageSizeInt < 1 {
		pageSizeInt = 10
	}

	offset := (pageInt - 1) * pageSizeInt

	// fetch paginated results
	result := h.DB.Order("verse_number asc").Offset(offset).Limit(pageSizeInt).Find(&verses)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching verses"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": verses,
	})
}

func (h *VerseHandler) GetVerseByID(c *gin.Context) {
	var verse models.Verse

	id := c.Param("id")
	result := h.DB.First(&verse, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Verse not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching verse"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": verse,
	})
}
