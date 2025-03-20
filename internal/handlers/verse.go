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
	var total int64

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

	// create base query
	baseQuery := h.DB.Model(&models.Verse{})

	// apply search filter to base query if search parameter exists
	if searchQuery := c.Query("q"); searchQuery != "" {
		searchPattern := "%" + searchQuery + "%"

		baseQuery = baseQuery.Where(
			"text ILIKE ? OR english_text ILIKE ? OR story_title ILIKE ? OR english_story_title ILIKE ? OR story ILIKE ? OR english_story ILIKE ?",
			searchPattern, searchPattern, searchPattern, searchPattern, searchPattern, searchPattern,
		)
	}

	// clone base query for pagination and add pagination-specific operations
	query := baseQuery.Session(&gorm.Session{}).Preload("Chapter").Order("verse_number asc").Offset(offset).Limit(pageSizeInt)

	// fetch paginated results
	result := query.Find(&verses)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching verses"})
		return
	}

	// clone base query for count
	countQuery := baseQuery.Session(&gorm.Session{})

	// get total count of verses
	countQuery.Count(&total)

	// calculate total pages
	totalPages := total / int64(pageSizeInt)
	if int(total)%pageSizeInt != 0 {
		totalPages++
	}

	c.JSON(http.StatusOK, gin.H{
		"data": verses,
		"meta": gin.H{
			"current_page":  pageInt,
			"per_page":      pageSizeInt,
			"total_records": total,
			"total_pages":   totalPages,
			"has_next":      pageInt < int(totalPages),
			"has_prev":      pageInt > 1,
		},
	})
}

func (h *VerseHandler) GetVerseByID(c *gin.Context) {
	var verse models.Verse

	id := c.Param("id")
	result := h.DB.Preload("Chapter").First(&verse, id)
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

func (h *VerseHandler) GetRandomVerse(c *gin.Context) {
	var verse models.Verse

	result := h.DB.Preload("Chapter").Order("RANDOM()").First(&verse)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching random verse"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": verse,
	})
}
