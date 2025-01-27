package main

import( "github.com/gin-gonic/gin"
)

type event struct {
	ID          string
	Title       string
	Description string
}

type allEvents []event

var events = allEvents{
	{
		ID:          "1",
		Title:       "Introduction to Golang",
		Description: "Come join us for a chance to learn how golang works and get to eventually try it out",
	},
	{
		ID:          "2",
		Title:       "java",
		Description: "Come join us for a chance to learn how golang works and get to eventually try it out",
	},
}

func getAllEvents(c *gin.Context) {
	// Json (http.statusok, gin.h {"status": "login successful"})
	// Json (http.statusok, map [string] interface {} {"status": "login successful"})

	c.JSON(200, gin.H{
		"data": events,
	})
}

func getOneEvent(c *gin.Context) {
	eventID := c.Param("id")
	for _, singleEvent := range events {
		if singleEvent.ID == eventID {
			c.JSON(200, gin.H{
				"data": singleEvent,
			})
		}
	}
}

func createEvent(c *gin.Context) {
	id := c.PostForm("ID")
	title := c.PostForm("Title")
	description := c.PostForm("Description")

	newEvent := event{id, title, description}

	events = append(events, newEvent)

	c.JSON(201, gin.H{
		"status":  201,
		"message": "new Event was created",
		"event":   newEvent,
	})
}

func updateEvent(c *gin.Context) {
	eventID := c.Param("id")

	id := c.PostForm("ID")
	title := c.PostForm("Title")
	description := c.PostForm("Description")

	for i, singleEvent := range events {
		if singleEvent.ID == eventID {
			singleEvent.ID = id
			singleEvent.Title = title
			singleEvent.Description = description

			events = append(events[:i], singleEvent)

			c.JSON(200, gin.H{
				"status":  200,
				"message": "new Event was updated",
				"event":   singleEvent,
			})
		}
	}
}

func deleteEvent(c *gin.Context) {

	eventID := c.Param("id")

	for i, singleEvent := range events {
		if singleEvent.ID == eventID {

			events = append(events[:i], events[i+1:]...)

			c.JSON(200, gin.H{
				"status":  200,
				"message": "Event has been deleted successfully",
				"event":   singleEvent,
			})
		}
	}
}

func main() {
	r := gin.Default()

	{
		// /api/v1/
		r.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hello world",
			})
		})
		r.GET("/events", getAllEvents)
		r.GET("/events/:id", getOneEvent)
		r.POST("/event", createEvent)
		r.PATCH("/events/:id", updateEvent)
		r.DELETE("/events/:id", deleteEvent)
	}

	r.Run()

}