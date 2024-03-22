package lib

import (
	"backend/src/helpers"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

var GoogleOAuthVerifier = oauth2.GenerateVerifier()

const CalendarName = "Jarvis[Beta]"

type googleCalendarState struct {
	UserSlug  string `json:"user"`
	Reference string `json:"reference"`
	// TODO: add more states or random numbers to validate if the backend create this state
}

func newCalendarState(userSlug string) *googleCalendarState {
	return &googleCalendarState{userSlug, "google-calendar"}
}

func (g *googleCalendarState) String() (string, error) {
	val, err := json.Marshal(g)
	if err != nil {
		return "", err
	}
	return string(val), nil
}

func ConvertCalendarState(code string) (*googleCalendarState, error) {
	var result googleCalendarState
	err := json.Unmarshal([]byte(code), &result)
	if err != nil {
		return nil, err
	}
	// TODO: Change reference and any other desired fields for match
	return &result, nil
}

type GoogleCalendarClient struct {
	Ctx      context.Context
	Calendar *calendar.Service
}

func getGoogleOAuthConfig() (*oauth2.Config, error) {
	// Reading tokens
	absPath, _ := filepath.Abs("config/google_oauth.json")
	byteValue, _ := os.ReadFile(absPath)

	// Get client
	return google.ConfigFromJSON(byteValue, calendar.CalendarScope) // TODO: Create a logic to switch scopes to client based on purpose
}

func GenerateGCalendarAuthorizationLink(userSlug string) (string, error) {
	config, _ := getGoogleOAuthConfig()
	stateToken, err := newCalendarState(userSlug).String()
	if err != nil {
		return "", errors.New("Trouble reading Google OAuth configuration")
	}
	return config.AuthCodeURL(stateToken, oauth2.AccessTypeOffline, oauth2.S256ChallengeOption(GoogleOAuthVerifier)), nil
}

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config, ctx context.Context, tok *oauth2.Token) *http.Client {
	// tok, _ := config.Exchange(context.TODO(), authCode)
	return config.Client(context.Background(), tok)
}

func NewCalendarClient(userSlug string) (*GoogleCalendarClient, error) {
	ctx := context.Background()
	config, err := getGoogleOAuthConfig()
	if err != nil {
		return nil, errors.New("trouble reading Google OAuth configuration")
	}

	authToken, _ := GetSavedCalendarToken(userSlug)

	client := getClient(config, ctx, authToken)
	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return nil, err
	}
	return &GoogleCalendarClient{ctx, srv}, nil
}

/*
Pseudo code:

for calendar in GoogleCalender:

	{id: string, name: string}
*/
func (client *GoogleCalendarClient) ListAllCalendarView() ([](map[string]string), error) {
	// List calendars
	calendarList, err := client.Calendar.CalendarList.List().Do()
	if err != nil {
		return make([]map[string]string, 0), err
	}

	var result []map[string]string
	// Loop through and print calendar details
	for _, calendar := range calendarList.Items {
		result = append(result, map[string]string{"id": calendar.Id, "name": calendar.Summary})
	}
	return result, nil
}

/*
Add event to primary calendar
*/
func (client *GoogleCalendarClient) AddEvent(calId string, eventData *calendar.Event) error {
	// List calendars
	_, err := client.Calendar.Events.Insert(calId, eventData).Do()
	return err
}

func (client *GoogleCalendarClient) FetchEventsWithFilters(startDate, endDate string) (GCalendarEventsGroup, error) {
	result := make(GCalendarEventsGroup, 0)
	events, err := client.Calendar.Events.List("primary").ShowDeleted(false).
		SingleEvents(true).TimeMin(startDate).TimeMax(endDate).OrderBy("startTime").Do()

	if err != nil {
		return result, err
	}

	if len(events.Items) != 0 {
		for _, item := range events.Items {
			startTime, _ := time.Parse(time.RFC3339, item.Start.DateTime)
			endTime, _ := time.Parse(time.RFC3339, item.End.DateTime)

			// Skipping all day events
			if endTime.Sub(startTime) >= 24*time.Hour {
				continue
			}

			result = append(result, NewGCalendarEvent(item.Summary, helpers.ConvertToTimezone(startTime, item.Start.TimeZone),
				helpers.ConvertToTimezone(endTime, item.End.TimeZone), false, helpers.CalendarIdToId(item.Id)))
		}
	}
	return result, nil
}

func (client *GoogleCalendarClient) FetchEvents(timezone string) (GCalendarEventsGroup, error) {
	result := make(GCalendarEventsGroup, 0)
	t := time.Now().Format(time.RFC3339)
	events, err := client.Calendar.Events.List("primary").ShowDeleted(false).
		SingleEvents(true).TimeMin(t).MaxResults(10).OrderBy("startTime").Do()

	if err != nil {
		return result, err
	}

	if len(events.Items) != 0 {
		for _, item := range events.Items {
			startTime, _ := time.Parse(time.RFC3339, item.Start.DateTime)
			endTime, _ := time.Parse(time.RFC3339, item.End.DateTime)
			result = append(result, NewGCalendarEvent(item.Summary, helpers.ConvertToTimezone(startTime, timezone), helpers.ConvertToTimezone(endTime, timezone), true, helpers.CalendarIdToId(item.Id)))
		}
	}
	return result, nil
}

func (client *GoogleCalendarClient) FindOrCreateJarvisCalendar(userSlug, timezone string) (string, error) {
	calendarNames, _ := client.ListAllCalendarView()
	for _, calendar := range calendarNames {
		if calendar["name"] == CalendarName {
			return calendar["id"], nil
		}
	}

	cal := &calendar.Calendar{
		Summary:  CalendarName,
		TimeZone: timezone,
	}

	createdCalendar, err := client.Calendar.Calendars.Insert(cal).Do()
	if err != nil {
		// log.Fatalf("Unable to create calendar: %v", err)
		return "", err
	}

	return createdCalendar.Id, nil
}

/*
date needs to be in the format "YYYY-MM-DD"
calendarId is the `id` for Jarvis in your account
*/
// TODO: Make sure it takes a function
func (client *GoogleCalendarClient) ClearEventsInCalendarByDate(calendarId, date string) error {
	// List events
	events, err := client.Calendar.Events.List(calendarId).TimeMin(date + "T00:00:00-00:00").TimeMax(date + "T23:59:59-00:00").Do()
	if err != nil {
		return err
	}

	// Delete each event
	for _, i := range events.Items {
		err := client.Calendar.Events.Delete(calendarId, i.Id).Do()
		if err != nil {
			return err
		}
	}

	return nil
}

/*
	Storing Calendar Token for User
*/

const CALENDAR_REDIS_PATTERN = "%s-calendar-token"

func SaveCalendarToken(user *googleCalendarState, token string) bool {
	config, _ := getGoogleOAuthConfig()
	tok, _ := config.Exchange(context.TODO(), token, oauth2.VerifierOption(GoogleOAuthVerifier))
	tokStr, _ := json.Marshal(tok)
	err := RedisClient.Store(fmt.Sprintf(CALENDAR_REDIS_PATTERN, user.UserSlug), string(tokStr))
	return err == nil
}

func GetSavedCalendarToken(userSlug string) (*oauth2.Token, error) {
	val, err := RedisClient.Get(fmt.Sprintf(CALENDAR_REDIS_PATTERN, userSlug))
	if err == nil {
		tok := &oauth2.Token{}
		json.Unmarshal([]byte(val), tok)
		return tok, nil
	}
	return nil, err
}
