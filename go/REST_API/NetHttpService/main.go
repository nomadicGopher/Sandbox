package main

import (
	_ "database/sql"
	"flag"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
	Rating uint8   `json:"rating:omitempty"`
}

func getAlbums(ginContext *gin.Context) {
	var albums []album

	// db, err := sql.Open("driver-name", "database=test.db")
	// if err != nil {
	// 	ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }
	// defer db.Close()

	// rows, err := db.Query("SELECT id, title, artist, price, rating FROM albums")
	// if err != nil {
	// 	ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }
	// defer rows.Close()

	// for rows.Next() {
	// 	var alb album
	// 	if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price, &alb.Rating); err != nil {
	// 		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 		return
	// 	}
	// 	albums = append(albums, alb)
	// }

	// if err := rows.Err(); err != nil {
	// 	ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	// For now, using hardcoded data for testing
	albums = []album{
		{ID: "1", Title: "KIDS", Artist: "Mac Miller", Price: 19.99, Rating: 10},
		{ID: "2", Title: "Might Delete Later", Artist: "J. Cole", Price: 14.79},
	}

	ginContext.IndentedJSON(http.StatusOK, albums)
}

func main() {
	port := flag.Int("port", 8080, "Port to run the server on")
	flag.Parse()

	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.Run(fmt.Sprintf("localhost:%d", *port))
}
