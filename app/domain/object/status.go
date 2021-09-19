package object

type (
	StatusID    = int64
	AccountIDType   = int64
	// Status status
	Status struct {

		// The internal ID of the status
		ID StatusID `json:"-"`

		// The content of the status
		Content string `json:"content,omitempty"`

		// The time the status was created
		CreateAt DateTime `json:"create_at,omitempty" db:"create_at"`

		// The account ID of status
		AccountID AccountIDType `json:"account_id,omitempty" db:"account_id"`
	}
)
