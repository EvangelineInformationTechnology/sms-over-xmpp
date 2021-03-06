package sms // import "github.com/mndrix/sms-over-xmpp"
import "net/http"

// SmsProvider describes a provider that is able to send and receive
// SMS messages.
type SmsProvider interface {
	// ReceiveSms retrieves the details of an incoming SMS based on an
	// HTTP request.
	ReceiveSms(r *http.Request) (*Sms, error)

	// SendSms sends the given SMS. It returns a unique identifier for
	// the outgoing message.  If possible, the identifier should
	// identify this message in the provider's logs.  If not, a random
	// identifier can be used.
	SendSms(*Sms) (string, error)
}

// CanSmsStatus describes a provider that can receive updates about
// the status of an SMS.
type CanSmsStatus interface {
	// SmsStatus returns true if the HTTP request contains an SMS status
	// update.  If so, it also returns the ID for that SMS and it's
	// status, respectively.  A status of "delivered" means that the
	// SMS has been delivered to its final destination.
	SmsStatus(*http.Request) (string, string, bool)
}
