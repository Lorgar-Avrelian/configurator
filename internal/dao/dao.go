package dao

import (
	"database/sql"
	"encoding/json"

	"github.com/google/uuid"
)

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

type DeviceComponentDao struct {
	ID            int64         `db:"id" json:"id"`
	Model         int64         `db:"model" json:"model"`
	InternalOrder sql.NullInt32 `db:"internal_order" json:"internal_order"`
	Parent        sql.NullInt64 `db:"parent" json:"parent"`
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

type DefaultConfigurationDao struct {
	ID                int64         `db:"id" json:"id"`
	Indicator         int64         `db:"indicator" json:"indicator"`
	DeviceComponentID sql.NullInt64 `db:"device_component_id" json:"device_component_id"`
}

type ThresholdDao struct {
	ID                  int64          `db:"id" json:"id"`
	SourceModel         int64          `db:"source_model" json:"source_model"`
	SourceInternalOrder int64          `db:"source_internal_order" json:"source_internal_order"`
	SourceParam         string         `db:"source_param" json:"source_param"`
	Value               string         `db:"value" json:"value"`
	Type                int16          `db:"type" json:"type"`
	Operator            int16          `db:"operator" json:"operator"`
	Enabled             sql.NullBool   `db:"enabled" json:"enabled"`
	TargetParam         sql.NullString `db:"target_param" json:"target_param"`
	TargetDevice        sql.NullInt64  `db:"target_device" json:"target_device"`
	Level               int16          `db:"level" json:"level"`
	PrevOperator        int16          `db:"prev_operator" json:"prev_operator"`
	Previous            sql.NullInt64  `db:"previous" json:"previous"`
}

type ConfigurationThresholdDao struct {
	ConfigurationID int64 `db:"configuration_id" json:"configuration_id"`
	ThresholdID     int64 `db:"threshold_id" json:"threshold_id"`
}

type DefaultConfigurationThresholdDao struct {
	DefaultConfigurationID int64 `db:"default_configuration_id" json:"default_configuration_id"`
	ThresholdID            int64 `db:"threshold_id" json:"threshold_id"`
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
	NumOid []string       `db:"num_oid" json:"num_oid"`
	StrOid []string       `db:"str_oid" json:"str_oid"`
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
