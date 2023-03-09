package model

import "time"

// ShortURL entity.
type ShortURL struct {
	ID        int       `json:"id,omitempty"`
	Code      string    `json:"code"`
	Original  string    `json:"original"`
	UserUUID  string    `json:"user_uuid,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	Deleted   bool      `json:"deleted,omitempty"`
}

// NewShortURL ShortURL entity constructor.
func NewShortURL(code, original string) *ShortURL {
	return &ShortURL{
		Code:     code,
		Original: original,
	}
}

// ShortURLBatchRequest entity.
type ShortURLBatchRequest struct {
	CorrelationID string `json:"correlation_id"`
	OriginalURL   string `json:"original_url"`
}

// ShortURLBatchResponse entity.
type ShortURLBatchResponse struct {
	CorrelationID string `json:"correlation_id"`
	ShortURL      string `json:"short_url"`
}
