package routes

import (
	"backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse events Id"})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Assuming `UserID` should come from the request or authentication context
	// event.UserID = someLogicToGetUserID() // Example placeholder

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event, please try again"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event Created!", "event": event})
}

// LogError logs the error with detailed information including file and line number

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {

		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse events Ids"})
		return
	}

	_, err = models.GetEventById(eventId)

	if err != nil {

		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fatch the events Id"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {

		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {

		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not Update request events Id"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event Updated successfully!"})

}
