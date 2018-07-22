package wow

// Quest represents a World of Warcraft quest.
type Quest struct {
	ID                    int    `json:"id"`
	Title                 string `json:"title"`
	ReqLevel              int    `json:"reqLevel"`
	SuggestedPartyMembers int    `json:"suggestedPartyMembers"`
	Category              string `json:"category"`
	Level                 int    `json:"level"`
}
