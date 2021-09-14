package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func initializeDB() *pgx.Conn {
	connUrl := "postgres://postgres:postgre9090@localhost:5432/goworlddb"
	conn, err := pgx.Connect(context.Background(), connUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return conn
}

func runGinRouter() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("static/", "./static")

	usr := &User{FirstName: "Denis", LastName: "Shchuka", Email: "admin@goworld.com"}
	//web pages routes
	router.GET("/", func(c *gin.Context) {
		//render template
		c.HTML(
			http.StatusOK,
			"index.html",
			gin.H{"title": "GoWorld - Home", "payload": usr, "menuHome": true},
		)
	})
	router.GET("/index.html", func(c *gin.Context) {
		//render template
		c.HTML(
			http.StatusOK,
			"index.html",
			gin.H{"title": "GoWorld - Main page", "payload": usr, "menuHome": true},
		)
	})
	router.GET("/hotels", func(c *gin.Context) {
		conn := initializeDB()
		defer conn.Close(context.Background())

		var allHotels []Hotel
		if rows, err := conn.Query(context.Background(), "SELECT id, country, destination, hotel FROM hotels"); err != nil {
			fmt.Println("Unable to get list of hotels due to: ", err)
		} else {
			// deferring query closing
			defer rows.Close()

			// Using tmp variable for reading
			var hotel Hotel

			// Next prepares the next row for reading.
			for rows.Next() {
				// Scan reads the values from the current row into tmp
				rows.Scan(&hotel.ID, &hotel.Country, &hotel.Destination, &hotel.Name)
				allHotels = append(allHotels, hotel)
			}
		}

		//render template
		c.HTML(
			http.StatusOK,
			"hotels.html",
			gin.H{"title": "GoWorld - Hotels", "payload": usr, "hotels": allHotels, "menuHotels": true},
		)
	})
	router.GET("/contacts", func(c *gin.Context) {
		//render template
		c.HTML(
			http.StatusOK,
			"contacts.html",
			gin.H{"title": "GoWorld - Contacts", "payload": usr, "menuContacts": true},
		)
	})
	//API routes
	v1 := router.Group("/v1/hotels")
	{
		v1.POST("/", createHotel)
		v1.GET("/", getAllHotels)
		v1.GET("/:id", getHotelByID)
	}
	router.Run()
}

func main() {
	runGinRouter()
}
