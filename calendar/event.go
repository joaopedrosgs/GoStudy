package calendar

type Event struct {
	Summary     string
	Description string
	Date        string
	IsProva     bool
	IsTrabalho  bool
}

//Events contains the event list
var eventList []Event

func GetEventList() []Event {
	if eventList == nil {
		Pool()
	}
	return eventList
}
