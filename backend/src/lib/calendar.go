package lib

import (
        "os"
		// "log"
		"fmt"
        "time"
		"errors"
		"context"
		"net/http"
		"encoding/json"
		"path/filepath"
        "golang.org/x/oauth2"
        "golang.org/x/oauth2/google"
		"google.golang.org/api/option"
        "google.golang.org/api/calendar/v3"
)

type googleCalendarState struct {
	UserSlug  string `json:"user"`
	reference string `json:"reference"`
	// TODO: add more states or random numbers to validate if the backend create this state
}

func newCalendarState(userSlug string) *googleCalendarState{
	return &googleCalendarState{userSlug, "google-calendar"}
}

func (g *googleCalendarState) String() (string, error){
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
	Ctx 	  context.Context
	Calendar *calendar.Service
}

func getGoogleOAuthConfig() (*oauth2.Config, error) {
	// Reading tokens
	absPath, _ := filepath.Abs("config/google_oauth.json")
	byteValue, _ := os.ReadFile(absPath)

	// Get client
	return google.ConfigFromJSON(byteValue, calendar.CalendarReadonlyScope)
}

func GenerateGCalendarAuthorizationLink(userSlug string) (string, error) {
	config, err := getGoogleOAuthConfig()
	stateToken, err := newCalendarState(userSlug).String()
	if err != nil {
		return "", errors.New("Trouble reading Google OAuth configuration")
	}
	return config.AuthCodeURL(stateToken, oauth2.AccessTypeOffline), nil
}

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config, ctx context.Context, authCode string) *http.Client {
	tok, _ := config.Exchange(context.TODO(), authCode)
	return config.Client(context.Background(), tok)
}

func NewCalendarClient(authCode string) (*GoogleCalendarClient, error) {
	ctx := context.Background()
	config, err := getGoogleOAuthConfig()
	if err != nil {
		return nil, errors.New("Trouble reading Google OAuth configuration")
	}

	client := getClient(config, ctx, authCode)
	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return nil, err
	}
	return &GoogleCalendarClient{ctx, srv}, nil
}

func (client *GoogleCalendarClient) FetchEvents() ([]string, error) {
	result := make([]string, 0)
	t := time.Now().Format(time.RFC3339)
	events, err := client.Calendar.Events.List("primary").ShowDeleted(false).
			SingleEvents(true).TimeMin(t).MaxResults(10).OrderBy("startTime").Do()

	if err != nil {
		// log.Fatalf("Unable to retrieve next ten of the user's events: %v", err)
		return result, err
	}

	if len(events.Items) != 0 {
		for _, item := range events.Items {
				date := item.Start.DateTime
				if date == "" {
						date = item.Start.Date
				}
				result = append(result, (fmt.Sprintf("%v (%v)\n", item.Summary, date)))
		}
	}
	return result, nil
}

/*
	Storing Calendar Token for User
*/

const CALENDAR_REDIS_PATTERN = "%s-calendar-token"


func SaveCalendarToken(user *googleCalendarState, token string) bool {
	err := RedisClient.Store(fmt.Sprintf(CALENDAR_REDIS_PATTERN, user.UserSlug), token)
	return err == nil
}

func GetSavedCalendar(userSlug string) (string, error) {
	val, err := RedisClient.Get(fmt.Sprintf(CALENDAR_REDIS_PATTERN, userSlug))
	if err == nil {
		return val, nil
	}
	return "", err
}