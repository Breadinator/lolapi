package data

import "encoding/json"

func NewMatchDTO() *MatchDTO {
	match := new(MatchDTO)
	match.Info.Participants = make([]ParticipantDTO, 10)
	match.Info.Teams = make([]TeamDTO, 2)
	return match
}

func GetMatchDTO(data []byte) *MatchDTO {
	match := NewMatchDTO()
	json.Unmarshal(data, match)
	return match
}
