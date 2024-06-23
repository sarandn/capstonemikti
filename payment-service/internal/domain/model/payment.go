package model

type Payment struct {
    ID        int     `json:"id"`
    Amount    float64 `json:"amount"`
    Status    string  `json:"status"`
    CreatedAt string  `json:"created_at"`
}
