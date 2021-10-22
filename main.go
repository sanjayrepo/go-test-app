// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// // This will be the entry point for _GAE_, we'e using GO default http package to pass the request to a handler func
// func init() {
// 	http.HandleFunc("/", handler)
// }

// // func main() {
// // 	http.HandleFunc("/", handler)
// // }

// // This will handle requests on the default path and return Hello World on the response
// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Hello, world!!!")
// }

//////////////////////

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Title A", Artist: "Artist A", Price: 100.23},
	{ID: "2", Title: "Title B", Artist: "Artist B", Price: 64.69},
	{ID: "3", Title: "Title C", Artist: "Artist C", Price: 33.09},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}
func postAlbums(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.Run("localhost: 8080")
}
