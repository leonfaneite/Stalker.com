package models

 
// User of the system.
type Query struct {
    Id          int64      `json:"id,string,omitempty"`
    Words    string    `json:"words,omitempty"`

}