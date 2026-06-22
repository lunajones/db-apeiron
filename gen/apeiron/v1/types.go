package apeironv1

// Compatibility helpers for historical server code.
// Canonical DB/server API fields are generated from proto/apeiron/v1/*.proto.

func (s *Skill) GetComboIndex() int64 {
	if s == nil {
		return 0
	}
	return int64(s.GetComboStep())
}
