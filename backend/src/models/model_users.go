package models

 
// User of the system.
type Query struct {
    Id          int      `json:"id"`
    Words    string    `json:"words,omitempty"`

}