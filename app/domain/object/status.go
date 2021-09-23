package object

type (
	StatusID    = int64
	AccountIDType   = int64
	// Status status
	Status struct {

		// The internal ID of the status
		ID StatusID `json:"id" db:"id"`

		// The account who creates status
		Account Account `json:"account,omitempty"`

		// The content of the status
		Content string `json:"content,omitempty" db:"content"`

		// The time the status was created
		CreateAt DateTime `json:"create_at,omitempty" db:"create_at"`

		// The account ID of status
		AccountID AccountIDType `json:"-" db:"account_id"`
	}
)
