package dto

type ComponentCreateDto struct {
	Title         string `json:"title" binding:"required" example:"Component"`
	NameEn        string `json:"name_en" binding:"required" example:"component"`
	NameRu        string `json:"name_ru" binding:"required" example:"Компонент"`
	PluralNameEn  string `json:"plural_name_en" binding:"required" example:"Components"`
	PluralNameRu  string `json:"plural_name_ru" binding:"required" example:"Компоненты"`
	BaseComponent *int64 `json:"base_component,omitempty" swaggertype:"integer" extensions:"x-nullable"`
	DescriptionEn string `json:"description_en,omitempty" swaggertype:"string" extensions:"x-nullable"`
	DescriptionRu string `json:"description_ru,omitempty" swaggertype:"string" extensions:"x-nullable"`
	Access        string `json:"access" swaggertype:"string" example:"USER"`
}

type ComponentDto struct {
	ID            int64      `json:"id" example:"1"`
	Title         string     `json:"title" example:"Component"`
	NameEn        string     `json:"name_en" example:"component"`
	NameRu        string     `json:"name_ru" example:"Компонент"`
	PluralNameEn  string     `json:"plural_name_en" example:"Components"`
	PluralNameRu  string     `json:"plural_name_ru" example:"Компоненты"`
	BaseComponent *int64     `json:"base_component" swaggertype:"integer" extensions:"x-nullable"`
	DescriptionEn *string    `json:"description_en,omitempty" swaggertype:"string" extensions:"x-nullable"`
	DescriptionRu *string    `json:"description_ru,omitempty" swaggertype:"string" extensions:"x-nullable"`
	Access        string     `json:"access" swaggertype:"string" example:"USER"`
	Params        []ParamDto `json:"params"`
}

type ComponentUpdateDto struct {
	ID            int64      `json:"id" example:"1"`
	Title         string     `json:"title" example:"Component"`
	NameEn        string     `json:"name_en" example:"component"`
	NameRu        string     `json:"name_ru" example:"Компонент"`
	PluralNameEn  string     `json:"plural_name_en" example:"Components"`
	PluralNameRu  string     `json:"plural_name_ru" example:"Компоненты"`
	BaseComponent *int64     `json:"base_component" swaggertype:"integer" extensions:"x-nullable"`
	DescriptionEn *string    `json:"description_en,omitempty" swaggertype:"string" extensions:"x-nullable"`
	DescriptionRu *string    `json:"description_ru,omitempty" swaggertype:"string" extensions:"x-nullable"`
	Access        string     `json:"access" swaggertype:"string" example:"USER"`
	Params        []ParamDto `json:"params"`
}

type ParamCreateDto struct {
	Title         string `json:"title" binding:"required" example:"Component name"`
	NameEn        string `json:"name_en" binding:"required" example:"name"`
	NameRu        string `json:"name_ru" binding:"required" example:"Имя"`
	Type          string `json:"type" binding:"required" swaggertype:"string" example:"VARCHAR"`
	Value         string `json:"value,omitempty" swaggertype:"string" extensions:"x-nullable"`
	DescriptionEn string `json:"description_en,omitempty" swaggertype:"string" extensions:"x-nullable"`
	DescriptionRu string `json:"description_ru,omitempty" swaggertype:"string" extensions:"x-nullable"`
	UnitsEn       string `json:"units_en,omitempty" swaggertype:"string" extensions:"x-nullable"`
	UnitsRu       string `json:"units_ru,omitempty" swaggertype:"string" extensions:"x-nullable"`
	Access        string `json:"access" binding:"required" swaggertype:"string" example:"USER"`
	Saved         bool   `json:"saved" example:"true"`
	Visible       bool   `json:"visible" example:"true"`
}

type ParamDto struct {
	ID            int64   `json:"id" example:"2"`
	Title         string  `json:"title" example:"Component name"`
	NameEn        string  `json:"name_en" example:"name"`
	NameRu        string  `json:"name_ru" example:"Имя"`
	Type          string  `json:"type" swaggertype:"string" example:"VARCHAR"`
	Value         *string `json:"value" swaggertype:"string" extensions:"x-nullable"`
	DescriptionEn *string `json:"description_en,omitempty" swaggertype:"string" extensions:"x-nullable"`
	DescriptionRu *string `json:"description_ru,omitempty" swaggertype:"string" extensions:"x-nullable"`
	UnitsEn       *string `json:"units_en,omitempty" swaggertype:"string" extensions:"x-nullable"`
	UnitsRu       *string `json:"units_ru,omitempty" swaggertype:"string" extensions:"x-nullable"`
	Access        string  `json:"access" swaggertype:"string" example:"USER"`
	Saved         bool    `json:"saved" example:"true"`
	Visible       bool    `json:"visible" example:"true"`
}

type ParamUpdateDto struct {
	ID            int64   `json:"id" example:"2"`
	Title         string  `json:"title" example:"Component name"`
	NameEn        string  `json:"name_en" example:"name"`
	NameRu        string  `json:"name_ru" example:"Имя"`
	Type          string  `json:"type" swaggertype:"string" example:"VARCHAR"`
	Value         *string `json:"value" swaggertype:"string" extensions:"x-nullable"`
	DescriptionEn *string `json:"description_en,omitempty" swaggertype:"string" extensions:"x-nullable"`
	DescriptionRu *string `json:"description_ru,omitempty" swaggertype:"string" extensions:"x-nullable"`
	UnitsEn       *string `json:"units_en,omitempty" swaggertype:"string" extensions:"x-nullable"`
	UnitsRu       *string `json:"units_ru,omitempty" swaggertype:"string" extensions:"x-nullable"`
	Access        string  `json:"access" swaggertype:"string" example:"USER"`
	Saved         bool    `json:"saved" example:"true"`
	Visible       bool    `json:"visible" example:"true"`
}

type ComponentParamLinkDto struct {
	ComponentID int64 `json:"component_id" binding:"required" example:"1"`
	ParamID     int64 `json:"param_id" binding:"required" example:"2"`
}
