package model

type Param struct {
	ID            int64   `json:"id" example:"1"`
	Title         string  `json:"title" example:"Макс. Напряжение"`
	NameEn        string  `json:"name_en" example:"max_voltage"`
	NameRu        string  `json:"name_ru" example:"макс_напряжение"`
	Type          VarType `json:"type" swaggertype:"string" example:"INT"`
	Value         string  `json:"value" example:"220"`
	DescriptionEn string  `json:"description_en,omitempty" example:"Maximum operational voltage"`
	DescriptionRu string  `json:"description_ru,omitempty" example:"Максимальное рабочее напряжение"`
	UnitsEn       string  `json:"units_en,omitempty" example:"V"`
	UnitsRu       string  `json:"units_ru,omitempty" example:"В"`
	Access        Access  `json:"access" swaggertype:"string" example:"ADMIN"`
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
	Access        Access  `json:"access" swaggertype:"string" example:"USER"`
	Params        []Param `json:"params,omitempty"`
}
