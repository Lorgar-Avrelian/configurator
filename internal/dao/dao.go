package dao

import (
	"database/sql"
	"encoding/json"

	"github.com/google/uuid"
)

type PollingProtocolDao struct {
	ID    int16          `db:"id" json:"id"`
	Value sql.NullString `db:"value" json:"value"`
}

type AccessDao struct {
	ID    int16          `db:"id" json:"id"`
	Value sql.NullString `db:"value" json:"value"`
}

type VersionSnmpDao struct {
	ID    int16          `db:"id" json:"id"`
	Value sql.NullString `db:"value" json:"value"`
}

type AuthProtocolSnmpDao struct {
	ID    int16          `db:"id" json:"id"`
	Value sql.NullString `db:"value" json:"value"`
}

type PrivacyProtocolSnmpDao struct {
	ID    int16          `db:"id" json:"id"`
	Value sql.NullString `db:"value" json:"value"`
}

type OidTypeDao struct {
	ID    int16          `db:"id" json:"id"`
	Value sql.NullString `db:"value" json:"value"`
}

type LogicOperatorDao struct {
	ID         int16  `db:"id" json:"id"`
	Value      string `db:"value" json:"value"`
	Type       string `db:"type" json:"type"`
	Precedence int16  `db:"precedence" json:"precedence"`
	Arity      int16  `db:"arity" json:"arity"`
}

type AlarmLevelDao struct {
	ID    int16          `db:"id" json:"id"`
	Value sql.NullString `db:"value" json:"value"`
}

type VarTypeDao struct {
	ID    int16          `db:"id" json:"id"`
	Value sql.NullString `db:"value" json:"value"`
}

type PollingFrequencyDao struct {
	ID    int16          `db:"id" json:"id"`
	Value sql.NullString `db:"value" json:"value"`
}

type OidAccessDao struct {
	ID    int16  `db:"id" json:"id"`
	Value string `db:"value" json:"value"`
}

type OidStatusDao struct {
	ID    int16  `db:"id" json:"id"`
	Value string `db:"value" json:"value"`
}

type Asn1TypeDao struct {
	ID    int16  `db:"id" json:"id"`
	Value string `db:"value" json:"value"`
}

type VendorDao struct {
	ID        int64          `db:"id" json:"id"`
	Name      string         `db:"name" json:"name"`
	Number    int32          `db:"number" json:"number"`
	Contact   sql.NullString `db:"contact" json:"contact"`
	Email     sql.NullString `db:"email" json:"email"`
	Directory sql.NullString `db:"directory" json:"directory"`
}

type Param struct {
	ID            int64          `db:"id" json:"id"`
	Title         string         `db:"title" json:"title"`
	NameEn        string         `db:"name_en" json:"name_en"`
	NameRu        string         `db:"name_ru" json:"name_ru"`
	Type          int16          `db:"type" json:"type"`
	Value         sql.NullString `db:"value" json:"value"`
	DescriptionEn sql.NullString `db:"description_en" json:"description_en"`
	DescriptionRu sql.NullString `db:"description_ru" json:"description_ru"`
	UnitsEn       sql.NullString `db:"units_en" json:"units_en"`
	UnitsRu       sql.NullString `db:"units_ru" json:"units_ru"`
	Access        int16          `db:"access" json:"access"`
	Saved         bool           `db:"saved" json:"saved"`
	Visible       bool           `db:"visible" json:"visible"`
	Diagram       bool           `db:"diagram" json:"diagram"`
}

type Component struct {
	ID            int64          `db:"id" json:"id"`
	Title         string         `db:"title" json:"title"`
	NameEn        string         `db:"name_en" json:"name_en"`
	NameRu        string         `db:"name_ru" json:"name_ru"`
	PluralNameEn  string         `db:"plural_name_en" json:"plural_name_en"`
	PluralNameRu  string         `db:"plural_name_ru" json:"plural_name_ru"`
	BaseComponent sql.NullInt64  `db:"base_component" json:"base_component"`
	DescriptionEn sql.NullString `db:"description_en" json:"description_en"`
	DescriptionRu sql.NullString `db:"description_ru" json:"description_ru"`
	Access        int16          `db:"access" json:"access"`
	Params        []Param        `json:"params"`
}

type ComponentDao struct {
	ID            int64          `db:"id" json:"id"`
	Title         string         `db:"title" json:"title"`
	NameEn        string         `db:"name_en" json:"name_en"`
	NameRu        string         `db:"name_ru" json:"name_ru"`
	PluralNameEn  string         `db:"plural_name_en" json:"plural_name_en"`
	PluralNameRu  string         `db:"plural_name_ru" json:"plural_name_ru"`
	BaseComponent sql.NullInt64  `db:"base_component" json:"base_component"`
	DescriptionEn sql.NullString `db:"description_en" json:"description_en"`
	DescriptionRu sql.NullString `db:"description_ru" json:"description_ru"`
	Access        int16          `db:"access" json:"access"`
}

type ParamDao struct {
	ID            int64          `db:"id" json:"id"`
	Title         string         `db:"title" json:"title"`
	NameEn        string         `db:"name_en" json:"name_en"`
	NameRu        string         `db:"name_ru" json:"name_ru"`
	Type          int16          `db:"type" json:"type"`
	Value         sql.NullString `db:"value" json:"value"`
	DescriptionEn sql.NullString `db:"description_en" json:"description_en"`
	DescriptionRu sql.NullString `db:"description_ru" json:"description_ru"`
	UnitsEn       sql.NullString `db:"units_en" json:"units_en"`
	UnitsRu       sql.NullString `db:"units_ru" json:"units_ru"`
	Access        int16          `db:"access" json:"access"`
	Saved         bool           `db:"saved" json:"saved"`
	Visible       bool           `db:"visible" json:"visible"`
	Diagram       bool           `db:"diagram" json:"diagram"`
}

type ComponentParamDao struct {
	ComponentID int64 `db:"component_id" json:"component_id"`
	ParamID     int64 `db:"param_id" json:"param_id"`
}

type MibDao struct {
	ID     int64          `db:"id" json:"id"`
	Path   string         `db:"path" json:"path"`
	Name   sql.NullString `db:"name" json:"name"`
	Vendor sql.NullInt64  `db:"vendor" json:"vendor"`
}

type OidDao struct {
	ID               uuid.UUID       `db:"id" json:"id"`
	Mib              sql.NullInt64   `db:"mib" json:"mib"`
	Type             int16           `db:"type" json:"type"`
	Name             string          `db:"name" json:"name"`
	Number           sql.NullInt32   `db:"number" json:"number"`
	DotterNotation   string          `db:"dotter_notation" json:"dotter_notation"`
	ObjectDescriptor string          `db:"object_descriptor" json:"object_descriptor"`
	Syntax           sql.NullString  `db:"syntax" json:"syntax"`
	Enum             json.RawMessage `db:"enum" json:"enum"`
	Status           sql.NullInt16   `db:"status" json:"status"`
	Access           sql.NullInt16   `db:"access" json:"access"`
	Units            sql.NullString  `db:"units" json:"units"`
	Description      sql.NullString  `db:"description" json:"description"`
	Category         sql.NullString  `db:"category" json:"category"`
}

type Oid struct {
	ID               uuid.UUID       `db:"id"`
	MibID            sql.NullInt64   `db:"mib_id"`
	MibPath          sql.NullString  `db:"mib_path"`
	MibName          sql.NullString  `db:"mib_name"`
	MibVendor        sql.NullInt64   `db:"mib_vendor"`
	Type             int16           `db:"type"`
	Name             string          `db:"name"`
	Number           sql.NullInt32   `db:"number"`
	DotterNotation   string          `db:"dotter_notation"`
	ObjectDescriptor string          `db:"object_descriptor"`
	Syntax           sql.NullString  `db:"syntax"`
	Enum             json.RawMessage `db:"enum"`
	Status           sql.NullInt16   `db:"status"`
	Access           sql.NullInt16   `db:"access"`
	Units            sql.NullString  `db:"units"`
	Description      sql.NullString  `db:"description"`
	Category         sql.NullString  `db:"category"`
}

type DeviceIndicatorDao struct {
	ID          int64          `db:"id" json:"id"`
	Description sql.NullString `db:"description" json:"description"`
	ObjectID    string         `db:"object_id" json:"object_id"`
	Contact     sql.NullString `db:"contact" json:"contact"`
	Name        sql.NullString `db:"name" json:"name"`
	Location    sql.NullString `db:"location" json:"location"`
	Services    sql.NullInt16  `db:"services" json:"services"`
}

type ParamIndicatorDao struct {
	ID             int64          `db:"id" json:"id"`
	OidID          uuid.UUID      `db:"oid_id" json:"oid_id"`
	DotterNotation sql.NullString `db:"dotter_notation" json:"dotter_notation"`
}

type ParamIndicator struct {
	ID                  int64           `db:"id" json:"id"`
	OidID               *uuid.UUID      `db:"oid_id" json:"oid_id"`
	DotterNotation      sql.NullString  `db:"dotter_notation" json:"dotter_notation"`
	OidMibID            sql.NullInt64   `db:"oid_mib_id" json:"oid_mib_id"`
	OidMibPath          sql.NullString  `db:"oid_mib_path" json:"oid_mib_path"`
	OidMibName          sql.NullString  `db:"oid_mib_name" json:"oid_mib_name"`
	OidMibVendor        sql.NullInt64   `db:"oid_mib_vendor" json:"oid_mib_vendor"`
	OidType             sql.NullInt16   `db:"oid_type" json:"oid_type"`
	OidName             sql.NullString  `db:"oid_name" json:"oid_name"`
	OidNumber           sql.NullInt32   `db:"oid_number" json:"oid_number"`
	OidDotterNotation   sql.NullString  `db:"oid_dotter_notation" json:"oid_dotter_notation"`
	OidObjectDescriptor sql.NullString  `db:"oid_object_descriptor" json:"oid_object_descriptor"`
	OidSyntax           sql.NullString  `db:"oid_syntax" json:"oid_syntax"`
	OidEnum             json.RawMessage `db:"oid_enum" json:"oid_enum"`
	OidStatus           sql.NullInt16   `db:"oid_status" json:"oid_status"`
	OidAccess           sql.NullInt16   `db:"oid_access" json:"oid_access"`
	OidUnits            sql.NullString  `db:"oid_units" json:"oid_units"`
	OidDescription      sql.NullString  `db:"oid_description" json:"oid_description"`
	OidCategory         sql.NullString  `db:"oid_category" json:"oid_category"`
}

type MappingDao struct {
	ID           int64           `db:"id" json:"id"`
	Indicator    int64           `db:"indicator" json:"indicator"`
	Param        int64           `db:"param" json:"param"`
	Frequency    int16           `db:"frequency" json:"frequency"`
	Value        sql.NullString  `db:"value" json:"value"`
	Coefficient  sql.NullFloat64 `db:"coefficient" json:"coefficient"`
	Enum         json.RawMessage `db:"enum" json:"enum"`
	Position     sql.NullInt16   `db:"position" json:"position"`
	From         sql.NullInt64   `db:"from" json:"from"`
	PositionType sql.NullInt16   `db:"position_type" json:"position_type"`
}

type Mapping struct {
	ID                  int64           `db:"id" json:"id"`
	IndicatorID         int64           `db:"indicator" json:"indicator"`
	ParamID             int64           `db:"param" json:"param"`
	Frequency           int16           `db:"frequency" json:"frequency"`
	Value               sql.NullString  `db:"value" json:"value"`
	Coefficient         sql.NullFloat64 `db:"coefficient" json:"coefficient"`
	Enum                json.RawMessage `db:"enum" json:"enum"`
	Position            sql.NullInt16   `db:"position" json:"position"`
	From                sql.NullInt64   `db:"from" json:"from"`
	PositionType        sql.NullInt16   `db:"position_type" json:"position_type"`
	IndOidID            *uuid.UUID      `db:"ind_oid_id" json:"ind_oid_id"`
	IndDotterNotation   sql.NullString  `db:"ind_dotter_notation" json:"ind_dotter_notation"`
	OidMibID            sql.NullInt64   `db:"oid_mib_id" json:"oid_mib_id"`
	OidMibPath          sql.NullString  `db:"oid_mib_path" json:"oid_mib_path"`
	OidMibName          sql.NullString  `db:"oid_mib_name" json:"oid_mib_name"`
	OidMibVendor        sql.NullInt64   `db:"oid_mib_vendor" json:"oid_mib_vendor"`
	OidType             sql.NullInt16   `db:"oid_type" json:"oid_type"`
	OidName             sql.NullString  `db:"oid_name" json:"oid_name"`
	OidNumber           sql.NullInt32   `db:"oid_number" json:"oid_number"`
	OidDotterNotation   sql.NullString  `db:"oid_dotter_notation" json:"oid_dotter_notation"`
	OidObjectDescriptor sql.NullString  `db:"oid_object_descriptor" json:"oid_object_descriptor"`
	OidSyntax           sql.NullString  `db:"oid_syntax" json:"oid_syntax"`
	OidEnum             json.RawMessage `db:"oid_enum" json:"oid_enum"`
	OidStatus           sql.NullInt16   `db:"oid_status" json:"oid_status"`
	OidAccess           sql.NullInt16   `db:"oid_access" json:"oid_access"`
	OidUnits            sql.NullString  `db:"oid_units" json:"oid_units"`
	OidDescription      sql.NullString  `db:"oid_description" json:"oid_description"`
	OidCategory         sql.NullString  `db:"oid_category" json:"oid_category"`
	ParamTitle          string          `db:"param_title" json:"param_title"`
	ParamNameEn         string          `db:"param_name_en" json:"param_name_en"`
	ParamNameRu         string          `db:"param_name_ru" json:"param_name_ru"`
	ParamType           int16           `db:"param_type" json:"param_type"`
	ParamValue          sql.NullString  `db:"param_value" json:"param_value"`
	ParamDescriptionEn  sql.NullString  `db:"param_description_en" json:"param_description_en"`
	ParamDescriptionRu  sql.NullString  `db:"param_description_ru" json:"param_description_ru"`
	ParamUnitsEn        sql.NullString  `db:"param_units_en" json:"param_units_en"`
	ParamUnitsRu        sql.NullString  `db:"param_units_ru" json:"param_units_ru"`
	ParamAccess         int16           `db:"param_access" json:"param_access"`
	ParamSaved          bool            `db:"param_saved" json:"param_saved"`
	ParamVisible        bool            `db:"param_visible" json:"param_visible"`
	ParamDiagram        bool            `db:"param_diagram" json:"param_diagram"`
}

type DeviceComponentDao struct {
	ID            int64         `db:"id" json:"id"`
	Model         int64         `db:"model" json:"model"`
	InternalOrder sql.NullInt32 `db:"internal_order" json:"internal_order"`
	Parent        sql.NullInt64 `db:"parent" json:"parent"`
}

type DeviceComponent struct {
	ID            int64         `db:"id" json:"id"`
	ModelID       int64         `db:"model" json:"model"`
	InternalOrder sql.NullInt32 `db:"internal_order" json:"internal_order"`
	Parent        sql.NullInt64 `db:"parent" json:"parent"`
	CompTitle     string        `db:"comp_title" json:"comp_title"`
	CompNameEn    string        `db:"comp_name_en" json:"comp_name_en"`
	CompNameRu    string        `db:"comp_name_ru" json:"comp_name_ru"`
}

type DeviceComponentMappingDao struct {
	DeviceComponentID int64 `db:"device_component_id" json:"device_component_id"`
	MappingID         int64 `db:"mapping_id" json:"mapping_id"`
}

type ConfigurationDao struct {
	ID                int64         `db:"id" json:"id"`
	Indicator         int64         `db:"indicator" json:"indicator"`
	DeviceComponentID sql.NullInt64 `db:"device_component_id" json:"device_component_id"`
}

type Configuration struct {
	ID                int64           `db:"id"`
	IndicatorID       string          `db:"indicator_id"`
	IndDescription    sql.NullString  `db:"ind_description"`
	IndContact        sql.NullString  `db:"ind_contact"`
	IndName           sql.NullString  `db:"ind_name"`
	IndLocation       sql.NullString  `db:"ind_location"`
	IndServices       sql.NullInt16   `db:"ind_services"`
	DeviceComponentID sql.NullInt64   `db:"device_component_id"`
	CompModel         sql.NullInt64   `db:"comp_model"`
	CompInternalOrder sql.NullInt32   `db:"comp_internal_order"`
	CompParent        sql.NullInt64   `db:"comp_parent"`
	CompTitle         sql.NullString  `db:"comp_title"`
	CompNameEn        sql.NullString  `db:"comp_name_en"`
	CompNameRu        sql.NullString  `db:"comp_name_ru"`
	MapID             sql.NullInt64   `db:"map_id"`
	MapIndicatorID    sql.NullInt64   `db:"map_indicator_id"`
	MapParamID        sql.NullInt64   `db:"map_param_id"`
	MapFrequency      sql.NullInt16   `db:"map_frequency"`
	MapValue          sql.NullString  `db:"map_value"`
	MapCoefficient    sql.NullFloat64 `db:"map_coefficient"`
	MapEnum           []byte          `db:"map_enum"`
	MapPosition       sql.NullInt16   `db:"map_position"`
	MapFrom           sql.NullInt64   `db:"map_from"`
	MapPositionType   sql.NullInt16   `db:"map_position_type"`
	MapIndOidID       *uuid.UUID      `db:"map_ind_oid_id"`
	MapIndDotter      sql.NullString  `db:"map_ind_dotter"`
	MapOidMibID       sql.NullInt64   `db:"map_oid_mib_id"`
	MapOidMibPath     sql.NullString  `db:"map_oid_mib_path"`
	MapOidMibName     sql.NullString  `db:"map_oid_mib_name"`
	MapOidMibVendor   sql.NullInt64   `db:"map_oid_mib_vendor"`
	MapOidType        sql.NullInt16   `db:"map_oid_type"`
	MapOidName        sql.NullString  `db:"map_oid_name"`
	MapOidNumber      sql.NullInt32   `db:"map_oid_number"`
	MapOidDotter      sql.NullString  `db:"map_oid_dotter"`
	MapOidDescriptor  sql.NullString  `db:"map_oid_descriptor"`
	MapOidSyntax      sql.NullString  `db:"map_oid_syntax"`
	MapOidEnum        []byte          `db:"map_oid_enum"`
	MapOidStatus      sql.NullInt16   `db:"map_oid_status"`
	MapOidAccess      sql.NullInt16   `db:"map_oid_access"`
	MapOidUnits       sql.NullString  `db:"map_oid_units"`
	MapOidDescription sql.NullString  `db:"map_oid_description"`
	MapOidCategory    sql.NullString  `db:"map_oid_category"`
	MapParamTitle     sql.NullString  `db:"map_param_title"`
	MapParamNameEn    sql.NullString  `db:"map_param_name_en"`
	MapParamNameRu    sql.NullString  `db:"map_param_name_ru"`
	MapParamType      sql.NullInt16   `db:"map_param_type"`
	MapParamValue     sql.NullString  `db:"map_param_value"`
	MapParamDescEn    sql.NullString  `db:"map_param_desc_en"`
	MapParamDescRu    sql.NullString  `db:"map_param_desc_ru"`
	MapParamUnitsEn   sql.NullString  `db:"map_param_units_en"`
	MapParamUnitsRu   sql.NullString  `db:"map_param_units_ru"`
	MapParamAccess    sql.NullInt16   `db:"map_param_access"`
	MapParamSaved     sql.NullBool    `db:"map_param_saved"`
	MapParamVisible   sql.NullBool    `db:"map_param_visible"`
	MapParamDiagram   sql.NullBool    `db:"map_param_diagram"`
}

type DefaultConfigurationDao struct {
	ID                int64         `db:"id" json:"id"`
	Indicator         int64         `db:"indicator" json:"indicator"`
	DeviceComponentID sql.NullInt64 `db:"device_component_id" json:"device_component_id"`
}

type DefaultConfiguration struct {
	ID                int64           `db:"id"`
	IndicatorID       int64           `db:"indicator_id"`
	IndDescription    sql.NullString  `db:"ind_description"`
	IndObjectID       string          `db:"ind_object_id"`
	IndContact        sql.NullString  `db:"ind_contact"`
	IndName           sql.NullString  `db:"ind_name"`
	IndLocation       sql.NullString  `db:"ind_location"`
	IndServices       sql.NullInt16   `db:"ind_services"`
	DeviceComponentID sql.NullInt64   `db:"device_component_id"`
	CompModel         sql.NullInt64   `db:"comp_model"`
	CompInternalOrder sql.NullInt32   `db:"comp_internal_order"`
	CompParent        sql.NullInt64   `db:"comp_parent"`
	CompTitle         sql.NullString  `db:"comp_title"`
	CompNameEn        sql.NullString  `db:"comp_name_en"`
	CompNameRu        sql.NullString  `db:"comp_name_ru"`
	MapID             sql.NullInt64   `db:"map_id"`
	MapIndicatorID    sql.NullInt64   `db:"map_indicator_id"`
	MapParamID        sql.NullInt64   `db:"map_param_id"`
	MapFrequency      sql.NullInt16   `db:"map_frequency"`
	MapValue          sql.NullString  `db:"map_value"`
	MapCoefficient    sql.NullFloat64 `db:"map_coefficient"`
	MapEnum           []byte          `db:"map_enum"`
	MapPosition       sql.NullInt16   `db:"map_position"`
	MapFrom           sql.NullInt64   `db:"map_from"`
	MapPositionType   sql.NullInt16   `db:"map_position_type"`
	MapIndOidID       *uuid.UUID      `db:"map_ind_oid_id"`
	MapIndDotter      sql.NullString  `db:"map_ind_dotter"`
	MapOidMibID       sql.NullInt64   `db:"map_oid_mib_id"`
	MapOidMibPath     sql.NullString  `db:"map_oid_mib_path"`
	MapOidMibName     sql.NullString  `db:"map_oid_mib_name"`
	MapOidMibVendor   sql.NullInt64   `db:"map_oid_mib_vendor"`
	MapOidType        sql.NullInt16   `db:"map_oid_type"`
	MapOidName        sql.NullString  `db:"map_oid_name"`
	MapOidNumber      sql.NullInt32   `db:"map_oid_number"`
	MapOidDotter      sql.NullString  `db:"map_oid_dotter"`
	MapOidDescriptor  sql.NullString  `db:"map_oid_descriptor"`
	MapOidSyntax      sql.NullString  `db:"map_oid_syntax"`
	MapOidEnum        []byte          `db:"map_oid_enum"`
	MapOidStatus      sql.NullInt16   `db:"map_oid_status"`
	MapOidAccess      sql.NullInt16   `db:"map_oid_access"`
	MapOidUnits       sql.NullString  `db:"map_oid_units"`
	MapOidDescription sql.NullString  `db:"map_oid_description"`
	MapOidCategory    sql.NullString  `db:"map_oid_category"`
	MapParamTitle     sql.NullString  `db:"map_param_title"`
	MapParamNameEn    sql.NullString  `db:"map_param_name_en"`
	MapParamNameRu    sql.NullString  `db:"map_param_name_ru"`
	MapParamType      sql.NullInt16   `db:"map_param_type"`
	MapParamValue     sql.NullString  `db:"map_param_value"`
	MapParamDescEn    sql.NullString  `db:"map_param_desc_en"`
	MapParamDescRu    sql.NullString  `db:"map_param_desc_ru"`
	MapParamUnitsEn   sql.NullString  `db:"map_param_units_en"`
	MapParamUnitsRu   sql.NullString  `db:"map_param_units_ru"`
	MapParamAccess    sql.NullInt16   `db:"map_param_access"`
	MapParamSaved     sql.NullBool    `db:"map_param_saved"`
	MapParamVisible   sql.NullBool    `db:"map_param_visible"`
	MapParamDiagram   sql.NullBool    `db:"map_param_diagram"`
}

type ThresholdDao struct {
	ID          int64           `db:"id" json:"id"`
	Name        string          `db:"name" json:"name"`
	Description sql.NullString  `db:"description" json:"description"`
	Author      string          `db:"author" json:"author"`
	Created     sql.NullTime    `db:"created" json:"created"`
	Query       json.RawMessage `db:"query" json:"query"`
	Target      json.RawMessage `db:"target" json:"target"`
	Value       string          `db:"value" json:"value"`
}

type ResultDao struct {
	Host          string         `db:"host" json:"host"`
	Port          int32          `db:"port" json:"port"`
	Component     int64          `db:"component" json:"component"`
	InternalOrder sql.NullInt32  `db:"internal_order" json:"internal_order"`
	Param         int64          `db:"param" json:"param"`
	Value         sql.NullString `db:"value" json:"value"`
}

type Result struct {
	Host           string         `db:"host" json:"host"`
	Port           int32          `db:"port" json:"port"`
	ComponentTitle string         `db:"component_title" json:"component_title"`
	InternalOrder  sql.NullInt32  `db:"internal_order" json:"internal_order"`
	ParamTitle     string         `db:"param_title" json:"param_title"`
	Value          sql.NullString `db:"value" json:"value"`
}

type AffectedThresholdDao struct {
	ID        int64        `db:"id" json:"id"`
	Host      string       `db:"host" json:"host"`
	Port      int32        `db:"port" json:"port"`
	Threshold int64        `db:"threshold" json:"threshold"`
	Enabled   sql.NullBool `db:"enabled" json:"enabled"`
}

type AffectedParamDao struct {
	ID            int64         `db:"id" json:"id"`
	Host          string        `db:"host" json:"host"`
	Port          int32         `db:"port" json:"port"`
	Component     int64         `db:"component" json:"component"`
	InternalOrder sql.NullInt32 `db:"internal_order" json:"internal_order"`
	Param         int64         `db:"param" json:"param"`
}

type DeviceSnmpDao struct {
	Host           string         `db:"host" json:"host"`
	Port           int32          `db:"port" json:"port"`
	Community      string         `db:"community" json:"community"`
	Version        int16          `db:"version" json:"version"`
	Login          sql.NullString `db:"login" json:"login"`
	Password       sql.NullString `db:"password" json:"password"`
	Authentication int16          `db:"authentication" json:"authentication"`
	Privacy        int16          `db:"privacy" json:"privacy"`
	ID             uuid.UUID      `db:"id" json:"id"`
	Config         sql.NullInt64  `db:"config" json:"config"`
}

type ConfigInProcessDao struct {
	Host     string `db:"host" json:"host"`
	Port     int32  `db:"port" json:"port"`
	Protocol int16  `db:"protocol" json:"protocol"`
}

type AgentCapabilitiesDao struct {
	ID             int64          `db:"id" json:"id"`
	Name           string         `db:"name" json:"name"`
	ProductRelease sql.NullString `db:"product_release" json:"product_release"`
	Status         sql.NullInt16  `db:"status" json:"status"`
	Description    sql.NullString `db:"description" json:"description"`
	Reference      sql.NullString `db:"reference" json:"reference"`
}

type AgentCapabilitiesModuleDao struct {
	ID       int64          `db:"id" json:"id"`
	Supports sql.NullString `db:"supports" json:"supports"`
	Includes []string       `db:"includes" json:"includes"`
}

type AgentCapabilitiesModuleNotificationDao struct {
	ID          int64          `db:"id" json:"id"`
	Variation   sql.NullString `db:"variation" json:"variation"`
	Access      sql.NullInt16  `db:"access" json:"access"`
	Description sql.NullString `db:"description" json:"description"`
}

type AgentCapabilitiesModuleObjectDao struct {
	ID               int64          `db:"id" json:"id"`
	Variation        sql.NullString `db:"variation" json:"variation"`
	Syntax           sql.NullString `db:"syntax" json:"syntax"`
	WriteSyntax      sql.NullString `db:"write_syntax" json:"write_syntax"`
	Access           sql.NullInt16  `db:"access" json:"access"`
	CreationRequires []string       `db:"creation_requires" json:"creation_requires"`
	Defval           sql.NullString `db:"defval" json:"defval"`
	Description      sql.NullString `db:"description" json:"description"`
}

type ChoiceDao struct {
	ID      int64           `db:"id" json:"id"`
	Name    string          `db:"name" json:"name"`
	Choices json.RawMessage `db:"choices" json:"choices"`
}

type ExplicitDao struct {
	ID    int64  `db:"id" json:"id"`
	Name  string `db:"name" json:"name"`
	Value string `db:"value" json:"value"`
}

type ImplicitDao struct {
	ID          int64  `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	Application int16  `db:"application" json:"application"`
	Value       string `db:"value" json:"value"`
}

type ImportDao struct {
	ID    int64  `db:"id" json:"id"`
	Param string `db:"param" json:"param"`
	From  string `db:"from" json:"from"`
}

type ModuleComplianceDao struct {
	ID          int64          `db:"id" json:"id"`
	Name        string         `db:"name" json:"name"`
	Status      sql.NullInt16  `db:"status" json:"status"`
	Description sql.NullString `db:"description" json:"description"`
	Reference   sql.NullString `db:"reference" json:"reference"`
}

type ModuleComplianceModuleDao struct {
	ID              int64    `db:"id" json:"id"`
	Name            string   `db:"name" json:"name"`
	MandatoryGroups []string `db:"mandatory_groups" json:"mandatory_groups"`
}

type ModuleComplianceModuleGroupDao struct {
	ID          int64          `db:"id" json:"id"`
	Name        string         `db:"name" json:"name"`
	Description sql.NullString `db:"description" json:"description"`
}

type ModuleComplianceModuleObjectDao struct {
	ID          int64          `db:"id" json:"id"`
	Name        string         `db:"name" json:"name"`
	Syntax      sql.NullString `db:"syntax" json:"syntax"`
	WriteSyntax sql.NullString `db:"write_syntax" json:"write_syntax"`
	Access      sql.NullInt16  `db:"access" json:"access"`
	Description sql.NullString `db:"description" json:"description"`
}

type ModuleIdentityDao struct {
	ID           int64          `db:"id" json:"id"`
	Name         string         `db:"name" json:"name"`
	LastUpdated  sql.NullTime   `db:"last_updated" json:"last_updated"`
	Organization sql.NullString `db:"organization" json:"organization"`
	ContactInfo  sql.NullString `db:"contact_info" json:"contact_info"`
	Description  sql.NullString `db:"description" json:"description"`
}

type NotificationGroupDao struct {
	ID            int64          `db:"id" json:"id"`
	Name          string         `db:"name" json:"name"`
	Notifications []string       `db:"notifications" json:"notifications"`
	Status        sql.NullInt16  `db:"status" json:"status"`
	Description   sql.NullString `db:"description" json:"description"`
	Reference     sql.NullString `db:"reference" json:"reference"`
}

type NotificationTypeDao struct {
	ID          int64          `db:"id" json:"id"`
	Name        string         `db:"name" json:"name"`
	Objects     []string       `db:"objects" json:"objects"`
	Status      sql.NullInt16  `db:"status" json:"status"`
	Description sql.NullString `db:"description" json:"description"`
	Reference   sql.NullString `db:"reference" json:"reference"`
}

type ObjectGroupDao struct {
	ID          int64          `db:"id" json:"id"`
	Name        string         `db:"name" json:"name"`
	Objects     []string       `db:"objects" json:"objects"`
	Status      sql.NullInt16  `db:"status" json:"status"`
	Description sql.NullString `db:"description" json:"description"`
	Reference   sql.NullString `db:"reference" json:"reference"`
}

type ObjectIdentifierDao struct {
	ID     int64          `db:"id" json:"id"`
	Name   string         `db:"name" json:"name"`
	NumOid []*string      `db:"num_oid" json:"num_oid"`
	StrOid []*string      `db:"str_oid" json:"str_oid"`
	Parent sql.NullString `db:"parent" json:"parent"`
	Type   int16          `db:"type" json:"type"`
}

type ObjectIdentityDao struct {
	ID          int64          `db:"id" json:"id"`
	Name        string         `db:"name" json:"name"`
	Status      sql.NullInt16  `db:"status" json:"status"`
	Description sql.NullString `db:"description" json:"description"`
	Reference   sql.NullString `db:"reference" json:"reference"`
}

type ObjectTypeDao struct {
	ID           int64          `db:"id" json:"id"`
	Name         string         `db:"name" json:"name"`
	Syntax       sql.NullString `db:"syntax" json:"syntax"`
	Units        sql.NullString `db:"units" json:"units"`
	Access       sql.NullInt16  `db:"access" json:"access"`
	Status       sql.NullInt16  `db:"status" json:"status"`
	Description  sql.NullString `db:"description" json:"description"`
	Reference    sql.NullString `db:"reference" json:"reference"`
	Index        []string       `db:"index" json:"index"`
	Augments     []string       `db:"augments" json:"augments"`
	DefaultValue sql.NullString `db:"default_value" json:"default_value"`
}

type RevisionDao struct {
	ID          int64          `db:"id" json:"id"`
	Revision    sql.NullTime   `db:"revision" json:"revision"`
	Description sql.NullString `db:"description" json:"description"`
}

type SequenceDao struct {
	ID    int64           `db:"id" json:"id"`
	Name  string          `db:"name" json:"name"`
	Pairs json.RawMessage `db:"pairs" json:"pairs"`
}

type TextualConventionDao struct {
	ID          int64          `db:"id" json:"id"`
	Name        string         `db:"name" json:"name"`
	DisplayHint sql.NullString `db:"display_hint" json:"display_hint"`
	Status      sql.NullInt16  `db:"status" json:"status"`
	Description sql.NullString `db:"description" json:"description"`
	Reference   sql.NullString `db:"reference" json:"reference"`
	Syntax      sql.NullString `db:"syntax" json:"syntax"`
}

type TrapTypeDao struct {
	ID          int64          `db:"id" json:"id"`
	Name        string         `db:"name" json:"name"`
	Variables   []string       `db:"variables" json:"variables"`
	Description sql.NullString `db:"description" json:"description"`
	Reference   sql.NullString `db:"reference" json:"reference"`
}

type MibToAgentCapabilitiesDao struct {
	ID                                    int64         `db:"id" json:"id"`
	MibID                                 int64         `db:"mib_id" json:"mib_id"`
	AgentCapabilitiesID                   int64         `db:"agent_capabilities_id" json:"agent_capabilities_id"`
	AgentCapabilitiesModuleID             sql.NullInt64 `db:"agent_capabilities_module_id" json:"agent_capabilities_module_id"`
	AgentCapabilitiesModuleNotificationID sql.NullInt64 `db:"agent_capabilities_module_notification_id" json:"agent_capabilities_module_notification_id"`
	AgentCapabilitiesModuleObjectID       sql.NullInt64 `db:"agent_capabilities_module_object_id" json:"agent_capabilities_module_object_id"`
}

type MibToChoiceDao struct {
	ID       int64 `db:"id" json:"id"`
	MibID    int64 `db:"mib_id" json:"mib_id"`
	ChoiceID int64 `db:"choice_id" json:"choice_id"`
}

type MibToExplicitDao struct {
	ID         int64 `db:"id" json:"id"`
	MibID      int64 `db:"mib_id" json:"mib_id"`
	ExplicitID int64 `db:"explicit_id" json:"explicit_id"`
}

type MibToImplicitDao struct {
	ID         int64 `db:"id" json:"id"`
	MibID      int64 `db:"mib_id" json:"mib_id"`
	ImplicitID int64 `db:"implicit_id" json:"implicit_id"`
}

type MibToImportDao struct {
	ID       int64 `db:"id" json:"id"`
	MibID    int64 `db:"mib_id" json:"mib_id"`
	ImportID int64 `db:"import_id" json:"import_id"`
}

type MibToModuleComplianceDao struct {
	ID                             int64         `db:"id" json:"id"`
	MibID                          int64         `db:"mib_id" json:"mib_id"`
	ModuleComplianceID             int64         `db:"module_compliance_id" json:"module_compliance_id"`
	ModuleComplianceModuleID       sql.NullInt64 `db:"module_compliance_module_id" json:"module_compliance_module_id"`
	ModuleComplianceModuleGroupID  sql.NullInt64 `db:"module_compliance_module_group_id" json:"module_compliance_module_group_id"`
	ModuleComplianceModuleObjectID sql.NullInt64 `db:"module_compliance_module_object_id" json:"module_compliance_module_object_id"`
}

type MibToModuleIdentityDao struct {
	ID               int64         `db:"id" json:"id"`
	MibID            int64         `db:"mib_id" json:"mib_id"`
	ModuleIdentityID int64         `db:"module_identity_id" json:"module_identity_id"`
	RevisionID       sql.NullInt64 `db:"revision_id" json:"revision_id"`
}

type MibToNotificationGroupDao struct {
	ID                  int64 `db:"id" json:"id"`
	MibID               int64 `db:"mib_id" json:"mib_id"`
	NotificationGroupID int64 `db:"notification_group_id" json:"notification_group_id"`
}

type MibToNotificationTypeDao struct {
	ID                 int64 `db:"id" json:"id"`
	MibID              int64 `db:"mib_id" json:"mib_id"`
	NotificationTypeID int64 `db:"notification_type_id" json:"notification_type_id"`
}

type MibToObjectGroupDao struct {
	ID            int64 `db:"id" json:"id"`
	MibID         int64 `db:"mib_id" json:"mib_id"`
	ObjectGroupID int64 `db:"object_group_id" json:"object_group_id"`
}

type MibToObjectIdentifierDao struct {
	ID                 int64 `db:"id" json:"id"`
	MibID              int64 `db:"mib_id" json:"mib_id"`
	ObjectIdentifierID int64 `db:"object_identifier_id" json:"object_identifier_id"`
}

type MibToObjectIdentityDao struct {
	ID               int64 `db:"id" json:"id"`
	MibID            int64 `db:"mib_id" json:"mib_id"`
	ObjectIdentityID int64 `db:"object_identity_id" json:"object_identity_id"`
}

type MibToObjectTypeDao struct {
	ID           int64 `db:"id" json:"id"`
	MibID        int64 `db:"mib_id" json:"mib_id"`
	ObjectTypeID int64 `db:"object_type_id" json:"object_type_id"`
}

type MibToSequenceDao struct {
	ID         int64 `db:"id" json:"id"`
	MibID      int64 `db:"mib_id" json:"mib_id"`
	SequenceID int64 `db:"sequence_id" json:"sequence_id"`
}

type MibToTextualConventionDao struct {
	ID                  int64 `db:"id" json:"id"`
	MibID               int64 `db:"mib_id" json:"mib_id"`
	TextualConventionID int64 `db:"textual_convention_id" json:"textual_convention_id"`
}

type MibToTrapTypeDao struct {
	ID         int64 `db:"id" json:"id"`
	MibID      int64 `db:"mib_id" json:"mib_id"`
	TrapTypeID int64 `db:"trap_type_id" json:"trap_type_id"`
}

type DeviceSnmp struct {
	Device     DeviceSnmpDao
	Thresholds []ThresholdDao
	Components []Component
}

type ChangeLogDao struct {
	ID       int32  `db:"id" json:"id"`
	Script   string `db:"script" json:"script"`
	Executed bool   `db:"executed" json:"executed"`
}
