package model

import (
	"database/sql"
	"fmt"
	"strings"
	"sync"
)

type Access int16
type AlarmLevel int16
type Asn1Type int16
type LogicOperator int16
type OidAccess int16
type OidStatus int16
type PollingFrequency int16
type VarType int16
type Vendor int64

type VendorData struct {
	ID        int64
	Name      string
	Number    int32
	Contact   sql.NullString
	Email     sql.NullString
	Directory sql.NullString
}

var (
	mu                      sync.RWMutex
	accessStrings           = make(map[Access]string)
	accessIds               = make(map[string]Access)
	alarmLevelStrings       = make(map[AlarmLevel]string)
	alarmLevelIds           = make(map[string]AlarmLevel)
	asn1TypeStrings         = make(map[Asn1Type]string)
	asn1TypeIds             = make(map[string]Asn1Type)
	logicOperatorStrings    = make(map[LogicOperator]string)
	logicOperatorIds        = make(map[string]LogicOperator)
	oidAccessStrings        = make(map[OidAccess]string)
	oidAccessIds            = make(map[string]OidAccess)
	oidStatusStrings        = make(map[OidStatus]string)
	oidStatusIds            = make(map[string]OidStatus)
	pollingFrequencyStrings = make(map[PollingFrequency]string)
	pollingFrequencyIds     = make(map[string]PollingFrequency)
	varTypeStrings          = make(map[VarType]string)
	varTypeIds              = make(map[string]VarType)
	vendorsMap              = make(map[Vendor]*VendorData)
	vendorNames             = make(map[string]Vendor)
	vendorDirectories       = make(map[string]Vendor)
)

func LoadRegistries(
	accessMap map[int16]string,
	varTypeMap map[int16]string,
	pollMap map[int16]string,
	asn1Map map[int16]string,
	statusMap map[int16]string,
	oidAccessMap map[int16]string,
	logicMap map[int16]string,
	alarmMap map[int16]string,
	vendors []map[string]interface{},
) {
	mu.Lock()
	defer mu.Unlock()
	for id, val := range accessMap {
		accessStrings[Access(id)] = val
		accessIds[strings.ToUpper(val)] = Access(id)
	}
	for id, val := range alarmMap {
		alarmLevelStrings[AlarmLevel(id)] = val
		alarmLevelIds[strings.ToUpper(val)] = AlarmLevel(id)
	}
	for id, val := range asn1Map {
		asn1TypeStrings[Asn1Type(id)] = val
		asn1TypeIds[strings.ToUpper(val)] = Asn1Type(id)
	}
	for id, val := range logicMap {
		logicOperatorStrings[LogicOperator(id)] = val
		logicOperatorIds[strings.ToUpper(val)] = LogicOperator(id)
	}
	for id, val := range oidAccessMap {
		oidAccessStrings[OidAccess(id)] = val
		oidAccessIds[strings.ToUpper(val)] = OidAccess(id)
	}
	for id, val := range statusMap {
		oidStatusStrings[OidStatus(id)] = val
		oidStatusIds[strings.ToUpper(val)] = OidStatus(id)
	}
	for id, val := range pollMap {
		pollingFrequencyStrings[PollingFrequency(id)] = val
		pollingFrequencyIds[strings.ToUpper(val)] = PollingFrequency(id)
	}
	for id, val := range varTypeMap {
		varTypeStrings[VarType(id)] = val
		varTypeIds[strings.ToUpper(val)] = VarType(id)
	}
	for _, v := range vendors {
		var vObj VendorData
		vObj.ID = v["id"].(int64)
		vObj.Name = v["name"].(string)
		vObj.Number = v["number"].(int32)
		vObj.Contact = v["contact"].(sql.NullString)
		vObj.Email = v["email"].(sql.NullString)
		vObj.Directory = v["directory"].(sql.NullString)
		var vKey Vendor
		vKey = Vendor(vObj.ID)
		vendorsMap[vKey] = &vObj
		var nameClean string
		nameClean = strings.ToUpper(strings.TrimSpace(vObj.Name))
		vendorNames[nameClean] = vKey
		if vObj.Directory.Valid && vObj.Directory.String != "" {
			var dirClean string
			dirClean = strings.ToUpper(strings.TrimSpace(vObj.Directory.String))
			vendorDirectories[dirClean] = vKey
		}
	}
}

func (a Access) String() string {
	mu.RLock()
	defer mu.RUnlock()
	var res string
	res = accessStrings[a]
	return res
}

func (a AlarmLevel) String() string {
	mu.RLock()
	defer mu.RUnlock()
	var res string
	res = alarmLevelStrings[a]
	return res
}

func (t Asn1Type) String() string {
	mu.RLock()
	defer mu.RUnlock()
	var res string
	res = asn1TypeStrings[t]
	return res
}

func (l LogicOperator) String() string {
	mu.RLock()
	defer mu.RUnlock()
	var res string
	res = logicOperatorStrings[l]
	return res
}

func (a OidAccess) String() string {
	mu.RLock()
	defer mu.RUnlock()
	var res string
	res = oidAccessStrings[a]
	return res
}

func (s OidStatus) String() string {
	mu.RLock()
	defer mu.RUnlock()
	var res string
	res = oidStatusStrings[s]
	return res
}

func (p PollingFrequency) String() string {
	mu.RLock()
	defer mu.RUnlock()
	var res string
	res = pollingFrequencyStrings[p]
	return res
}

func (v VarType) String() string {
	mu.RLock()
	defer mu.RUnlock()
	var res string
	res = varTypeStrings[v]
	return res
}

func (v Vendor) Data() *VendorData {
	mu.RLock()
	defer mu.RUnlock()
	var res *VendorData
	res = vendorsMap[v]
	return res
}

func (a Access) MarshalJSON() ([]byte, error) {
	var res []byte
	res = []byte(fmt.Sprintf("%q", a.String()))
	return res, nil
}

func (a AlarmLevel) MarshalJSON() ([]byte, error) {
	var res []byte
	res = []byte(fmt.Sprintf("%q", a.String()))
	return res, nil
}

func (t Asn1Type) MarshalJSON() ([]byte, error) {
	var res []byte
	res = []byte(fmt.Sprintf("%q", t.String()))
	return res, nil
}

func (l LogicOperator) MarshalJSON() ([]byte, error) {
	var res []byte
	res = []byte(fmt.Sprintf("%q", l.String()))
	return res, nil
}

func (a OidAccess) MarshalJSON() ([]byte, error) {
	var res []byte
	res = []byte(fmt.Sprintf("%q", a.String()))
	return res, nil
}

func (s OidStatus) MarshalJSON() ([]byte, error) {
	var res []byte
	res = []byte(fmt.Sprintf("%q", s.String()))
	return res, nil
}

func (p PollingFrequency) MarshalJSON() ([]byte, error) {
	var res []byte
	res = []byte(fmt.Sprintf("%q", p.String()))
	return res, nil
}

func (v VarType) MarshalJSON() ([]byte, error) {
	var res []byte
	res = []byte(fmt.Sprintf("%q", v.String()))
	return res, nil
}

func (a *Access) UnmarshalJSON(b []byte) error {
	var str string
	str = strings.Trim(string(b), `"`)
	*a = ParseAccess(str)
	return nil
}

func (a *AlarmLevel) UnmarshalJSON(b []byte) error {
	var str string
	str = strings.Trim(string(b), `"`)
	*a = ParseAlarmLevel(str)
	return nil
}

func (t *Asn1Type) UnmarshalJSON(b []byte) error {
	var str string
	str = strings.Trim(string(b), `"`)
	*t = ParseAsn1Type(str)
	return nil
}

func (l *LogicOperator) UnmarshalJSON(b []byte) error {
	var str string
	str = strings.Trim(string(b), `"`)
	*l = ParseLogicOperator(str)
	return nil
}

func (a *OidAccess) UnmarshalJSON(b []byte) error {
	var str string
	str = strings.Trim(string(b), `"`)
	*a = ParseOidAccess(str)
	return nil
}

func (s *OidStatus) UnmarshalJSON(b []byte) error {
	var str string
	str = strings.Trim(string(b), `"`)
	*s = ParseOidStatus(str)
	return nil
}

func (p *PollingFrequency) UnmarshalJSON(b []byte) error {
	var str string
	str = strings.Trim(string(b), `"`)
	*p = ParsePollingFrequency(str)
	return nil
}

func (v *VarType) UnmarshalJSON(b []byte) error {
	var str string
	str = strings.Trim(string(b), `"`)
	*v = ParseVarType(str)
	return nil
}

func ParseAccess(s string) Access {
	mu.RLock()
	defer mu.RUnlock()
	var clean string
	clean = strings.ToUpper(strings.TrimSpace(s))
	var id Access
	var ok bool
	id, ok = accessIds[clean]
	if ok {
		return id
	}
	return 1
}

func ParseAlarmLevel(s string) AlarmLevel {
	mu.RLock()
	defer mu.RUnlock()
	var clean string
	clean = strings.ToUpper(strings.TrimSpace(s))
	var id AlarmLevel
	var ok bool
	id, ok = alarmLevelIds[clean]
	if ok {
		return id
	}
	return 1
}

func ParseAsn1Type(s string) Asn1Type {
	mu.RLock()
	defer mu.RUnlock()
	var clean string
	clean = strings.ToUpper(strings.TrimSpace(s))
	var id Asn1Type
	var ok bool
	id, ok = asn1TypeIds[clean]
	if ok {
		return id
	}
	return 1
}

func ParseLogicOperator(s string) LogicOperator {
	mu.RLock()
	defer mu.RUnlock()
	var clean string
	clean = strings.ToUpper(strings.TrimSpace(s))
	var id LogicOperator
	var ok bool
	id, ok = logicOperatorIds[clean]
	if ok {
		return id
	}
	return 1
}

func ParseOidAccess(s string) OidAccess {
	mu.RLock()
	defer mu.RUnlock()
	var clean string
	clean = strings.ToUpper(strings.TrimSpace(s))
	var id OidAccess
	var ok bool
	id, ok = oidAccessIds[clean]
	if ok {
		return id
	}
	return 1
}

func ParseOidStatus(s string) OidStatus {
	mu.RLock()
	defer mu.RUnlock()
	var clean string
	clean = strings.ToUpper(strings.TrimSpace(s))
	var id OidStatus
	var ok bool
	id, ok = oidStatusIds[clean]
	if ok {
		return id
	}
	return 1
}

func ParsePollingFrequency(s string) PollingFrequency {
	mu.RLock()
	defer mu.RUnlock()
	var clean string
	clean = strings.ToUpper(strings.TrimSpace(s))
	var id PollingFrequency
	var ok bool
	id, ok = pollingFrequencyIds[clean]
	if ok {
		return id
	}
	return 1
}

func ParseVarType(s string) VarType {
	mu.RLock()
	defer mu.RUnlock()
	var clean string
	clean = strings.ToUpper(strings.TrimSpace(s))
	var id VarType
	var ok bool
	id, ok = varTypeIds[clean]
	if ok {
		return id
	}
	return 1
}

func ParseVendor(s string) Vendor {
	mu.RLock()
	defer mu.RUnlock()
	var clean string
	clean = strings.ToUpper(strings.TrimSpace(s))
	var id Vendor
	var ok bool
	id, ok = vendorNames[clean]
	if ok {
		return id
	}
	id, ok = vendorDirectories[clean]
	if ok {
		return id
	}
	return 0
}
