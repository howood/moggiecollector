package requestid

import (
	"net/http"

	"github.com/howood/moggiecollector/infrastructure/uuid"
)

type RequestContextKey string

// KeyRequestID is XRequestId key
const KeyRequestID = "X-Request-ID"

func generateRequestID() string {
	return uuid.GetUUID(uuid.SatoriUUID)
}

// GetRequestID returns XRequestId
func GetRequestID(r *http.Request) string {
	if r.Header.Get(KeyRequestID) != "" {
		return r.Header.Get(KeyRequestID)
	}
	return generateRequestID()
}

func GetRequestIDKey() RequestContextKey {
	return RequestContextKey(KeyRequestID)
}
