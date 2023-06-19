package middleware

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func MiddleDB(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("databaseConn", db)
		c.Next()
	}
}
