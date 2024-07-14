package saga

import "errors"

var (
	ErrNoKey = errors.New("no key")
)

type State map[string]any

func NewState() State {
	return State{}
}

func (s State) Get(key string) (any, error) {
	v, ok := s[key]
	if !ok {
		return nil, ErrNoKey
	}

	return v, nil
}

func (s State) Set(key string, value any) {
	s[key] = value
}
