package handlers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	ch := NewChapterHandler(db)
	vh := NewVerseHandler(db)
	sh := NewSearchHandler(db)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	v1 := r.Group("/api/v1")
	{
		v1.GET("/chapters", ch.GetChapters)
		v1.GET("/chapters/:id", ch.GetChapterByID)

		v1.GET("/verses", vh.GetVerses)
		v1.GET("/verses/:id", vh.GetVerseByID)

		v1.GET("/search", sh.Search)
	}
}
