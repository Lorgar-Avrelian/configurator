package dto

import "filler/internal/model"

type ComponentCreate struct {
	Title         string `json:"title" binding:"required" example:"Модуль Питания"`
	NameEn        string `json:"name_en" binding:"required" example:"power_module"`
	NameRu        string `json:"name_ru" binding:"required" example:"модуль_питания"`
	BaseComponent *int64 `json:"base_component,omitempty" example:"1"`
	DescriptionEn string `json:"description_en,omitempty" example:"Main power controller unit"`
	DescriptionRu string `json:"description_ru,omitempty" example:"Главный контроллер питания"`
	Access        string `json:"access" binding:"required" example:"ADMIN"`
}

type ComponentUpdate struct {
	Title         string `json:"title" binding:"required" example:"Обновленный Модуль"`
	NameEn        string `json:"name_en" binding:"required" example:"power_module_v2"`
	NameRu        string `json:"name_ru" binding:"required" example:"модуль_питания_в2"`
	BaseComponent *int64 `json:"base_component,omitempty" example:"1"`
	DescriptionEn string `json:"description_en,omitempty" example:"Updated main power controller unit"`
	DescriptionRu string `json:"description_ru,omitempty" example:"Обновленный главный контроллер питания"`
	Access        string `json:"access" binding:"required" example:"OWNER"`
}

type ParamCreate struct {
	Title         string `json:"title" binding:"required" example:"Макс. Напряжение"`
	NameEn        string `json:"name_en" binding:"required" example:"max_voltage"`
	NameRu        string `json:"name_ru" binding:"required" example:"макс_напряжение"`
	Type          string `json:"type" binding:"required" example:"INT"`
	Value         string `json:"value,omitempty" example:"220"`
	DescriptionEn string `json:"description_en,omitempty" example:"Maximum operational voltage"`
	DescriptionRu string `json:"description_ru,omitempty" example:"Максимальное рабочее напряжение"`
	UnitsEn       string `json:"units_en,omitempty" example:"V"`
	UnitsRu       string `json:"units_ru,omitempty" example:"В"`
	Access        string `json:"access" binding:"required" example:"ADMIN"`
	Saved         bool   `json:"saved" example:"true"`
	Visible       bool   `json:"visible" example:"true"`
}

type ParamUpdate struct {
	Title         string `json:"title" binding:"required" example:"Макс. Напряжение"`
	NameEn        string `json:"name_en" binding:"required" example:"max_voltage"`
	NameRu        string `json:"name_ru" binding:"required" example:"макс_напряжение"`
	Type          string `json:"type" binding:"required" example:"INT"`
	Value         string `json:"value,omitempty" example:"240"`
	DescriptionEn string `json:"description_en,omitempty" example:"Maximum operational voltage"`
	DescriptionRu string `json:"description_ru,omitempty" example:"Максимальное рабочее напряжение"`
	UnitsEn       string `json:"units_en,omitempty" example:"V"`
	UnitsRu       string `json:"units_ru,omitempty" example:"В"`
	Access        string `json:"access" binding:"required" example:"ADMIN"`
	Saved         bool   `json:"saved" example:"true"`
	Visible       bool   `json:"visible" example:"true"`
}

type BindParamRequest struct {
	ComponentID int64 `json:"component_id" binding:"required" example:"10"`
	ParamID     int64 `json:"param_id" binding:"required" example:"5"`
}

type OidPageResponse struct {
	Page       int         `json:"page" example:"1"`
	PerPage    int         `json:"per_page" example:"100"`
	TotalItems int         `json:"total_items" example:"450"`
	Items      []model.Oid `json:"items"`
}
