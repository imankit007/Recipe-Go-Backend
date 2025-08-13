package middleware

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"sample.com/crud-api/utils/logger"
)

const TxKey = "TX"

func TransactionMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		logger.Log.Print("Starting a new transaction")

		tx := db.Begin()
		if tx.Error != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": "Failed to start transaction"})
			return
		}

		// Store tx in context
		c.Set(TxKey, tx)

		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
				panic(r) // re-panic after rollback
			} else if len(c.Errors) > 0 {
				tx.Rollback()
			} else {
				tx.Commit()
			}
		}()

		c.Next()
	}
}
