// Code generated by "enumer -type=State"; DO NOT EDIT.

package fsrep

import (
	"fmt"
)

const _StateName = "ReadyCompleted"

var _StateIndex = [...]uint8{0, 5, 14}

func (i State) String() string {
	i -= 1
	if i >= State(len(_StateIndex)-1) {
		return fmt.Sprintf("State(%d)", i+1)
	}
	return _StateName[_StateIndex[i]:_StateIndex[i+1]]
}

var _StateValues = []State{1, 2}

var _StateNameToValueMap = map[string]State{
	_StateName[0:5]:  1,
	_StateName[5:14]: 2,
}

// StateString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func StateString(s string) (State, error) {
	if val, ok := _StateNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to State values", s)
}

// StateValues returns all values of the enum
func StateValues() []State {
	return _StateValues
}

// IsAState returns "true" if the value is listed in the enum definition. "false" otherwise
func (i State) IsAState() bool {
	for _, v := range _StateValues {
		if i == v {
			return true
		}
	}
	return false
}