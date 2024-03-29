package statistics

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterRoutes(router gin.IRouter, db *mongo.Database, excludeAddrs []string) {
	router.GET("/statistics", HandlerGetStatistics(db, excludeAddrs))
}
