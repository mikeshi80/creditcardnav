package datastruct

import (
	"time"
)

type Bank struct {
	Name string // the name of the bank
	Desc string // the description of the bank
}

type Card struct {
	Bank *Bank
	Name string // the name of the card
	Desc string // the description of the card
}

type Position struct {
	Latitude  float32
	Longitude float32
}

type Business struct {
	Name     string   // It is the name get from dianping.com
	Aliases  []string // a business has a lot of names in credit card websites.
	Addr     string   // address
	Phones   []string
	Position *Position // the GPS position which can help mobil apps work better
}

type Event struct {
	Bank   *Bank      // the event bank
	Cards  []*Card    // the available cards, if empty, all cards available
	Name   string     // the event name, maybe empty
	Start  *time.Time // event start time, may be nil if not exact start day
	End    *time.Time // event end time, may be nil if not exact end day
	Cond   string     // the condition that you can apply this event
	Create *time.Time // The first creation time when Clawer gets it
	Update *time.Time // The latest update time
}
