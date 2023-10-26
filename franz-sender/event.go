package main

type RawEvent struct {
        Source string           `json:"source"`
        EventName string        `json:"event_name"`
        EventStatus string      `json:"event_status"`
        Created string          `json:"created"`
        Payout string           `json:"payout"`
}
