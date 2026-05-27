package model

import (
	"fmt"
	"strings"
	"sync"
)

type Access int16
type VarType int16
type PollingFrequency int16

var (
	mu                      sync.RWMutex
	accessStrings           = make(map[Access]string)
	accessIds               = make(map[string]Access)
	varTypeStrings          = make(map[VarType]string)
	varTypeIds              = make(map[string]VarType)
	pollingFrequencyStrings = make(map[PollingFrequency]string)
	pollingFrequencyIds     = make(map[string]PollingFrequency)
)

// LoadRegistries заполняет карты данными, полученными из слоя БД
func LoadRegistries(accessMap map[int16]string, varTypeMap map[int16]string, pollMap map[int16]string) {
	mu.Lock()
	defer mu.Unlock()
	for id, val := range accessMap {
		accessStrings[Access(id)] = val
		accessIds[strings.ToUpper(val)] = Access(id)
	}
	for id, val := range varTypeMap {
		varTypeStrings[VarType(id)] = val
		varTypeIds[strings.ToUpper(val)] = VarType(id)
	}
	for id, val := range pollMap {
		pollingFrequencyStrings[PollingFrequency(id)] = val
		pollingFrequencyIds[strings.ToUpper(val)] = PollingFrequency(id)
	}
}

func ParseAccess(s string) Access {
	mu.RLock()
	defer mu.RUnlock()
	s = strings.ToUpper(strings.TrimSpace(s))
	if id, ok := accessIds[s]; ok {
		return id
	}
	return 1
}

func ParseVarType(s string) VarType {
	mu.RLock()
	defer mu.RUnlock()
	s = strings.ToUpper(strings.TrimSpace(s))
	if id, ok := varTypeIds[s]; ok {
		return id
	}
	return 1
}

func ParsePollingFrequency(s string) PollingFrequency {
	mu.RLock()
	defer mu.RUnlock()
	s = strings.ToUpper(strings.TrimSpace(s))
	if id, ok := pollingFrequencyIds[s]; ok {
		return id
	}
	return 1 // Дефолтный ID (например, LOW)
}

func (a Access) String() string {
	mu.RLock()
	defer mu.RUnlock()
	if val, ok := accessStrings[a]; ok {
		return val
	}
	return fmt.Sprintf("UNKNOWN_ACCESS_%d", a)
}

func (v VarType) String() string {
	mu.RLock()
	defer mu.RUnlock()
	if val, ok := varTypeStrings[v]; ok {
		return val
	}
	return fmt.Sprintf("UNKNOWN_VARTYPE_%d", v)
}

func (p PollingFrequency) String() string {
	mu.RLock()
	defer mu.RUnlock()
	if val, ok := pollingFrequencyStrings[p]; ok {
		return val
	}
	return fmt.Sprintf("UNKNOWN_POLLING_%d", p)
}

func (a Access) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", a.String())), nil
}

func (a *Access) UnmarshalJSON(b []byte) error {
	str := strings.Trim(string(b), `"`)
	*a = ParseAccess(str)
	return nil
}

func (v VarType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", v.String())), nil
}

func (v *VarType) UnmarshalJSON(b []byte) error {
	str := strings.Trim(string(b), `"`)
	*v = ParseVarType(str)
	return nil
}
