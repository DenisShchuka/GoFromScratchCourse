package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//API Methods///////////////////////////////
func createHotel(c *gin.Context) {
	//initialize DB
	conn := initializeDB()
	defer conn.Close(context.Background())

	var hotel Hotel

	if err := c.ShouldBindJSON(&hotel); err == nil {
		//&hotel
		row := conn.QueryRow(context.Background(), "INSERT INTO hotels (country, destination, hotel) VALUES ($1, $2, $3) RETURNING id", hotel.Country, hotel.Destination, hotel.Name) //тут ломаемся
		var id int
		err = row.Scan(&id)
		if err != nil {
			log.Printf("Error: %v\n", err)
		}
		//return OK
		c.JSON(http.StatusAccepted, gin.H{"status": "created"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"Unable to add new hotel due to ": err.Error()})
	}
}

func getAllHotels(c *gin.Context) {
	//initialize DB
	conn := initializeDB()
	defer conn.Close(context.Background())

	var results []Hotel
	//get all hotel objects from the database
	if hotels, err := conn.Query(context.Background(), "SELECT id, country, destination, hotel FROM hotels"); err != nil {
		fmt.Println("Unable to get list of hotels due to: ", err)
	} else {
		// deferring query closing
		defer hotels.Close()

		// Using tmp variable for reading
		var hotel Hotel

		// Next prepares the next row for reading.
		for hotels.Next() {
			// Scan reads the values from the current row into tmp
			hotels.Scan(&hotel.ID, &hotel.Country, &hotel.Destination, &hotel.Name)
			results = append(results, hotel)
		}

		if results != nil {
			c.JSON(http.StatusOK, gin.H{"result": results})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"result": hotel})
		}
		if hotels.Err() != nil {
			// if any error occurred while reading rows.
			fmt.Println("Error will reading hotels table: ", err)
		}
	}
}

func getHotelByID(c *gin.Context) {
	//initialize DB
	conn := initializeDB()
	defer conn.Close(context.Background())

	//get hotel id from URI path param
	hotelId := c.Param("id")

	//get hotel object from the database
	if hotels, err := conn.Query(context.Background(), "SELECT id, country, destination, hotel FROM hotels h WHERE h.ID = $1 ", hotelId); err != nil {
		fmt.Println("Unable to get the hotel by ID due to: ", err)
	} else {
		// deferring query closing
		defer hotels.Close()

		// Using tmp variable for reading
		var hotel Hotel

		// Next prepares the next row for reading.
		if hotels.Next() {
			// Scan reads the values from the current row into tmp
			hotels.Scan(&hotel.ID, &hotel.Country, &hotel.Destination, &hotel.Name)
			c.JSON(http.StatusOK, gin.H{"result": hotel})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"result": hotel})
		}
		if hotels.Err() != nil {
			// if any error occurred while reading rows.
			fmt.Println("Error will reading hotels table: ", err)
		}
	}
}

//END API methods//////////////////////////
