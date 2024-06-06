package model

type Ticket struct {
    ID     int    `json:"id"`
    Title  string `json:"title"`
    Status string `json:"status"`
}
