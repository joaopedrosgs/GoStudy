package calendar

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/user"
	"strings"
	"time"

	"github.com/dustin/go-humanize"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	user, err := user.Current()
	if err != nil {
		log.Fatal("Failed to get current user: ", err.Error())
	}
	tokFile := user.HomeDir + "/.config/go-study/token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(oauth2.NoContext, authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	defer f.Close()
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	json.NewEncoder(f).Encode(token)
}

// Pool update the event list from google calendar
func Pool() {
	if eventList == nil {
		eventList = make([]Event, 0, 10)
	}
	user, err := user.Current()
	if err != nil {
		log.Fatal("Failed to get current user: ", err.Error())
	}
	b, err := ioutil.ReadFile(user.HomeDir + "/.config/go-study/credentials.json")
	if err != nil {
		log.Fatal("Unable to read client secret file: ", err)
		return
	}

	config, err := google.ConfigFromJSON(b, calendar.CalendarReadonlyScope)
	if err != nil {
		log.Fatal("Unable to parse client secret file to config: ", err)
		return
	}
	client := getClient(config)

	srv, err := calendar.New(client)
	if err != nil {
		log.Fatal("Unable to retrieve Calendar client: ", err)
		return
	}

	t := time.Now().Format(time.RFC3339)
	events, err := srv.Events.List("primary").ShowDeleted(false).
		SingleEvents(true).TimeMin(t).MaxResults(200).OrderBy("startTime").Do()
	if err != nil {
		log.Fatal("Unable to retrieve next ten of the user's events: ", err)
		return
	}
	for _, item := range events.Items {
		event := Event{Summary: strings.Title(strings.ToLower(item.Summary)), Description: item.Description}
		date := item.Start.DateTime
		if date == "" {
			date = item.Start.Date
		}
		timeDate, _ := time.Parse(time.RFC3339, date)
		event.Date = humanize.Time(timeDate)
		if strings.Contains(event.Summary, "Prova") {
			event.IsProva = true
		}
		if strings.Contains(event.Summary, "Trabalho") {
			event.IsTrabalho = true
		}

		eventList = append(eventList, event)
	}

}
