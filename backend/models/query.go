package models

import "time"

// QueryStatus represents the lifecycle state of a query
// Possible values: pending, checking, paused, fulfilled, notified, cancelled, error
// See consts below for canonical values

type QueryStatus string

type StatusChange struct {
	Status    QueryStatus `json:"status"`
	ChangedAt time.Time   `json:"changed_at"`
	Note      string      `json:"note,omitempty"`
}

type Query struct {
	ID          string         `json:"id"`
	UserID      string         `json:"user_id"`
	QueryText   string         `json:"query_text"`
	Status      QueryStatus    `json:"status"`
	Result      string         `json:"result,omitempty"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	NotifiedAt  *time.Time     `json:"notified_at,omitempty"`
	ErrorMsg    string         `json:"error_msg,omitempty"`
	History     []StatusChange `json:"history,omitempty"`
}

const (
	StatusPending   QueryStatus = "pending"
	StatusChecking  QueryStatus = "checking"
	StatusPaused    QueryStatus = "paused"
	StatusFulfilled QueryStatus = "fulfilled"
	StatusNotified  QueryStatus = "notified"
	StatusCancelled QueryStatus = "cancelled"
	StatusError     QueryStatus = "error"
)
