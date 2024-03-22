package controllers

import (
	"backend/src/helpers"
	"backend/src/models"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CalendarShow(c *gin.Context) {
	calendarID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid skill ID"})
		return
	}
	user := CurrentUser(c)
	calendarEvents, err := models.CalendarShow(calendarID, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "events": calendarEvents})
}

func CalendarShowAll(c *gin.Context) {
	user := CurrentUser(c)

	calendars, err := models.CalendarShowAll(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	result := make([]models.CalendarJSON, 0)
	for _, calendar := range calendars {
		result = append(result, calendar.Marshal())
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "calendars": result})
}

func CalendarEvents(c *gin.Context) {
	userJWT := CurrentUser(c)

	user, err := models.UserFind(userJWT.UserId.String())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	timezone := user.TimeZone()
	userTimeZone, _ := time.LoadLocation(timezone)
	var startDate, endDate time.Time

	startDateParam := c.Query("startDate")
	if startDateParam != "" {
		startDate, _ = helpers.ConvertISOStringToTime(startDateParam, "UTC")
		startDate = helpers.ConvertToTimezone(startDate, timezone)
	} else {
		startDate = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, userTimeZone)
	}

	endDateParam := c.Query("endDate")
	if endDateParam != "" {
		endDate, _ = helpers.ConvertISOStringToTime(endDateParam, "UTC")
		endDate = helpers.ConvertToTimezone(endDate, timezone)
	} else {
		tomorrow := time.Now().AddDate(0, 0, 1)
		endDate = time.Date(tomorrow.Year(), tomorrow.Month(), tomorrow.Day(), 0, 0, 0, 0, userTimeZone)
	}

	log.Println("Start Date: ", startDate)
	log.Println("End Date: ", endDate)

	calendarEvents, err := user.CalendarEventsWithFilters(helpers.TimeFormatRFC3339(helpers.TimeToUTC(startDate)), helpers.TimeFormatRFC3339(helpers.TimeToUTC(endDate)))

	if err == nil {
		c.JSON(http.StatusOK, gin.H{"code": 200, "events": calendarEvents})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
