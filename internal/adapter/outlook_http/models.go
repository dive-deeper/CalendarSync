package outlook_http

import (
	"github.com/inovex/CalendarSync/internal/models"
)

// https://learn.microsoft.com/en-us/graph/api/resources/event?view=graph-rest-1.0

type EventList struct {
	NextLink string  `json:"@odata.nextLink"`
	Events   []Event `json:"value"`
}

type Event struct {
	ID                         string       `json:"id"`
	UID                        string       `json:"iCalUId"`
	ChangeKey                  string       `json:"changeKey"`
	HtmlLink                   string       `json:"webLink"`
	Subject                    string       `json:"subject"`
	Start                      Time         `json:"start"`
	End                        Time         `json:"end"`
	Body                       Body         `json:"body,omitempty"`
	Attendees                  []Attendee   `json:"attendees,omitempty"`
	Location                   Location     `json:"location"`
	IsReminderOn               bool         `json:"isReminderOn"`
	ReminderMinutesBeforeStart int          `json:"reminderMinutesBeforeStart"`
	Extensions                 []Extensions `json:"extensions"`
	IsAllDay                   bool         `json:"isAllDay"`
	OnlineMeetingUrl           string       `json:"onlineMeetingUrl"`
}

type Extensions struct {
	OdataType     string `json:"@odata.type"`
	ExtensionName string `json:"extensionName"`
	// needs to be embedded, Microsoft returns a 500 on an non-embedded object
	models.Metadata
}

type Body struct {
	ContentType string `json:"contentType,omitempty"`
	Content     string `json:"content,omitempty"`
}

type Time struct {
	DateTime string `json:"dateTime"`
	TimeZone string `json:"timeZone"`
}

type Attendee struct {
	EmailAdress EmailAddress `json:"emailAddress,omitempty"`
}

type EmailAddress struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type Location struct {
	Name string `json:"displayName"`
}
