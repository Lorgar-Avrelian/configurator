package model

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

type Param struct {
	ID            int64   `json:"id" example:"1"`
	Title         string  `json:"title" example:"Макс. Напряжение"`
	NameEn        string  `json:"name_en" example:"max_voltage"`
	NameRu        string  `json:"name_ru" example:"макс_напряжение"`
	Type          VarType `json:"type" swaggertype:"primitive,string" example:"INT"`
	Value         string  `json:"value" example:"220"`
	DescriptionEn string  `json:"description_en,omitempty" example:"Maximum operational voltage"`
	DescriptionRu string  `json:"description_ru,omitempty" example:"Максимальное рабочее напряжение"`
	UnitsEn       string  `json:"units_en,omitempty" example:"V"`
	UnitsRu       string  `json:"units_ru,omitempty" example:"В"`
	Access        Access  `json:"access" swaggertype:"primitive,string" example:"ADMIN"`
	Saved         bool    `json:"saved" example:"true"`
	Visible       bool    `json:"visible" example:"true"`
}

type Component struct {
	ID            int64   `json:"id" example:"10"`
	Title         string  `json:"title" example:"Модуль Питания"`
	NameEn        string  `json:"name_en" example:"power_module"`
	NameRu        string  `json:"name_ru" example:"модуль_питания"`
	BaseComponent *int64  `json:"base_component,omitempty" example:"1"`
	DescriptionEn string  `json:"description_en,omitempty" example:"Main power controller unit"`
	DescriptionRu string  `json:"description_ru,omitempty" example:"Главный контроллер питания"`
	Access        Access  `json:"access" swaggertype:"primitive,string" example:"USER"`
	Params        []Param `json:"params,omitempty"`
}

type Oid struct {
	ID               uuid.UUID       `json:"id" example:"00000000-0000-0000-0000-000000000000"`
	MibID            *int64          `json:"mib_id,omitempty" example:"1"`
	Type             Asn1Type        `json:"type" swaggertype:"primitive,string" example:"OBJECT-TYPE"`
	Name             string          `json:"name" example:"sysDescr"`
	Number           *int32          `json:"number,omitempty" example:"1"`
	DotterNotation   string          `json:"dotter_notation" example:"1.3.6.1.2.1.1.1"`
	ObjectDescriptor string          `json:"object_descriptor" example:"1.3.6.1.2.1.1.1"`
	Syntax           string          `json:"syntax,omitempty" example:"DisplayString"`
	Enum             json.RawMessage `json:"enum,omitempty" swaggertype:"primitive,object"`
	Status           *OidStatus      `json:"status,omitempty" swaggertype:"primitive,string" example:"current"`
	Access           *OidAccess      `json:"access,omitempty" swaggertype:"primitive,string" example:"read-only"`
	Units            string          `json:"units,omitempty" example:"seconds"`
	Description      string          `json:"description,omitempty" example:"A textual description of the entity."`
	Category         string          `json:"category,omitempty" example:"system"`
}

// CompareOids Сравнивает два OID по числовым сегментам dotter_notation.
// Возвращает true, если o1 должен идти раньше o2.
func CompareOids(o1, o2 string) bool {
	parts1 := strings.Split(o1, ".")
	parts2 := strings.Split(o2, ".")
	len1 := len(parts1)
	len2 := len(parts2)
	minLen := len1
	if len2 < minLen {
		minLen = len2
	}
	for i := 0; i < minLen; i++ {
		n1, err1 := strconv.Atoi(parts1[i])
		n2, err2 := strconv.Atoi(parts2[i])
		if err1 != nil || err2 != nil {
			if parts1[i] != parts2[i] {
				return parts1[i] < parts2[i]
			}
			continue
		}
		if n1 != n2 {
			return n1 < n2
		}
	}
	return len1 < len2
}
