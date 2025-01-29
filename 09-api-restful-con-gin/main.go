package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Artista string `json:"artista"`
	Price   string `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Dark Side of the Moon", Artista: "Pink Floyd", Price: "29.99"},
	{ID: "2", Title: "Thriller", Artista: "Michael Jackson", Price: "24.99"},
	{ID: "3", Title: "Back in Black", Artista: "AC/DC", Price: "19.99"},
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

func getAlbumsById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message:": "album no encontrado"})
}

func main() {

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "Hola señor, esto es Vinyl App",
		})
	})

	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumsById)

	router.Run(":8080")
}
