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

type MibDto struct {
	ID     int64   `json:"id" example:"4912"`
	Path   string  `json:"path" example:"SNMPv2-MIB.mib"`
	Name   *string `json:"name" swaggertype:"string" extensions:"x-nullable" example:"SNMPv2-MIB"`
	Vendor *string `json:"vendor" swaggertype:"string" extensions:"x-nullable" example:"BASE"`
}

type OidDto struct {
	ID               string            `json:"id" example:"14e8713a-2f3c-3af1-8e6f-449d7a612227"`
	Mib              *MibDto           `json:"mib" extensions:"x-nullable"`
	Type             string            `json:"type" example:"OBJECT IDENTIFIER"`
	Name             string            `json:"name" example:"sysObjectID"`
	Number           *int32            `json:"number" swaggertype:"integer" extensions:"x-nullable" example:"2"`
	DotterNotation   string            `json:"dotter_notation" example:".1.3.6.1.2.1.1.2"`
	ObjectDescriptor string            `json:"object_descriptor" example:".iso.org.dod.internet.mgmt.mib-2.system.sysObjectID"`
	Syntax           *string           `json:"syntax,omitempty" swaggertype:"string" extensions:"x-nullable" example:"OBJECT IDENTIFIER"`
	Enum             map[string]string `json:"enum,omitempty" swaggertype:"object" extensions:"x-nullable"`
	Status           *string           `json:"status" extensions:"x-nullable" example:"current"`
	Access           *string           `json:"access" extensions:"x-nullable" example:"read-only"`
	Units            *string           `json:"units,omitempty" swaggertype:"string" extensions:"x-nullable"`
	Description      *string           `json:"description,omitempty" swaggertype:"string" extensions:"x-nullable" example:"The vendor's authoritative identification of the network management subsystem contained in the entity."`
	Category         *string           `json:"category,omitempty" swaggertype:"string" extensions:"x-nullable" example:"system"`
}

type DeviceIndicatorCreateDto struct {
	Description *string `json:"description" swaggertype:"string" extensions:"x-nullable" example:"Linux server-node-01 5.4.0-74-generic"`
	ObjectID    string  `json:"object_id" binding:"required" example:".1.3.6.1.4.1.8072.3.2.10"`
	Contact     *string `json:"contact" swaggertype:"string" extensions:"x-nullable" example:"sysadmin@company.com"`
	Name        *string `json:"name" swaggertype:"string" extensions:"x-nullable" example:"node-01.local"`
	Location    *string `json:"location" swaggertype:"string" extensions:"x-nullable" example:"Rack 04, Room 202"`
	Services    *int16  `json:"services" swaggertype:"integer" extensions:"x-nullable" example:"72"`
}

type DeviceIndicatorDto struct {
	ID          int64   `json:"id" example:"1"`
	Description *string `json:"description" swaggertype:"string" extensions:"x-nullable" example:"Linux server-node-01 5.4.0-74-generic"`
	ObjectID    string  `json:"object_id" example:".1.3.6.1.4.1.8072.3.2.10"`
	Contact     *string `json:"contact" swaggertype:"string" extensions:"x-nullable" example:"sysadmin@company.com"`
	Name        *string `json:"name" swaggertype:"string" extensions:"x-nullable" example:"node-01.local"`
	Location    *string `json:"location" swaggertype:"string" extensions:"x-nullable" example:"Rack 04, Room 202"`
	Services    *int16  `json:"services" swaggertype:"integer" extensions:"x-nullable" example:"72"`
}

type ParamIndicatorCreateDto struct {
	OidID          *string `json:"oid_id,omitempty" swaggertype:"string" extensions:"x-nullable" example:"14e8713a-2f3c-3af1-8e6f-449d7a612227"`
	DotterNotation *string `json:"dotter_notation,omitempty" swaggertype:"string" extensions:"x-nullable" example:".1.3.6.1.2.1.1.2"`
}

type ParamIndicatorDto struct {
	ID             int64   `json:"id" example:"1"`
	Oid            *OidDto `json:"oid" extensions:"x-nullable"`
	DotterNotation *string `json:"dotter_notation" swaggertype:"string" extensions:"x-nullable" example:".1.3.6.1.2.1.1.2"`
}

type MappingCreateDto struct {
	IndicatorID  int64             `json:"indicator" binding:"required" example:"1"`
	ParamID      int64             `json:"param" binding:"required" example:"2"`
	Frequency    string            `json:"frequency" binding:"required" example:"MEDIUM"`
	Value        *string           `json:"value,omitempty" swaggertype:"string" extensions:"x-nullable"`
	Coefficient  *float64          `json:"coefficient,omitempty" swaggertype:"number" extensions:"x-nullable"`
	Enum         map[string]string `json:"enum,omitempty" swaggertype:"object" extensions:"x-nullable"`
	Position     *int16            `json:"position,omitempty" swaggertype:"integer" extensions:"x-nullable" example:"1"`
	From         *int64            `json:"from,omitempty" swaggertype:"integer" extensions:"x-nullable"`
	PositionType *string           `json:"position_type,omitempty" swaggertype:"string" extensions:"x-nullable" example:"INTEGER"`
}

type MappingDto struct {
	ID           int64              `json:"id" example:"1"`
	Indicator    *ParamIndicatorDto `json:"indicator,omitempty" extensions:"x-nullable"`
	Param        *ParamDto          `json:"param,omitempty" extensions:"x-nullable"`
	Frequency    string             `json:"frequency" example:"MEDIUM"`
	Value        *string            `json:"value" swaggertype:"string" extensions:"x-nullable"`
	Coefficient  *float64           `json:"coefficient" swaggertype:"number" extensions:"x-nullable"`
	Enum         map[string]string  `json:"enum" swaggertype:"object" extensions:"x-nullable"`
	Position     *int16             `json:"position" swaggertype:"integer" extensions:"x-nullable" example:"1"`
	From         *int64             `json:"from" swaggertype:"integer" extensions:"x-nullable"`
	PositionType *string            `json:"position_type" swaggertype:"string" extensions:"x-nullable" example:"INTEGER"`
	Children     []MappingDto       `json:"children"`
}

type ComponentInfoDto struct {
	ID     int64  `json:"id" example:"1"`
	Title  string `json:"title" example:"Chassis"`
	NameEn string `json:"name_en" example:"chassis"`
	NameRu string `json:"name_ru" example:"Шасси"`
}

type DeviceComponentCreateDto struct {
	ModelID       int64  `json:"model" binding:"required" example:"1"`
	InternalOrder *int32 `json:"internal_order,omitempty" swaggertype:"integer" extensions:"x-nullable" example:"1"`
	Parent        *int64 `json:"parent,omitempty" swaggertype:"integer" extensions:"x-nullable"`
}

type DeviceComponentDto struct {
	ID            int64                `json:"id" example:"1"`
	Component     ComponentInfoDto     `json:"component"`
	InternalOrder *int32               `json:"internal_order" swaggertype:"integer" extensions:"x-nullable" example:"1"`
	Parent        *int64               `json:"parent" swaggertype:"integer" extensions:"x-nullable"`
	Mappings      []MappingDto         `json:"mappings"`
	Children      []DeviceComponentDto `json:"children"`
}

type DefaultConfigurationDto struct {
	ID              int64               `json:"id" example:"1"`
	Indicator       *DeviceIndicatorDto `json:"indicator,omitempty" extensions:"x-nullable"`
	DeviceComponent *DeviceComponentDto `json:"device_component,omitempty" extensions:"x-nullable"`
}

type ConfigurationDto struct {
	ID              int64               `json:"id" example:"1"`
	Indicator       *DeviceIndicatorDto `json:"indicator,omitempty" extensions:"x-nullable"`
	DeviceComponent *DeviceComponentDto `json:"device_component,omitempty" extensions:"x-nullable"`
}
