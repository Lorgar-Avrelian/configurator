package dao

import (
	"configurator/internal/database"
	"configurator/internal/logger"
	"configurator/internal/model"
	"context"
	"database/sql"
	"encoding/json"
)

func LoadEnumsFromDB(ctx context.Context) error {
	conn := database.Get()
	accessMap := make(map[int16]string)
	aRows, err := conn.Query(ctx, `SELECT id, value FROM public.access`)
	if err == nil {
		defer aRows.Close()
		for aRows.Next() {
			var id int16
			var val string
			_ = aRows.Scan(&id, &val)
			accessMap[id] = val
		}
	}
	varTypeMap := make(map[int16]string)
	vRows, err := conn.Query(ctx, `SELECT id, value FROM public.var_type`)
	if err == nil {
		defer vRows.Close()
		for vRows.Next() {
			var id int16
			var val string
			_ = vRows.Scan(&id, &val)
			varTypeMap[id] = val
		}
	}
	pollMap := make(map[int16]string)
	pRows, err := conn.Query(ctx, `SELECT id, value FROM public.polling_frequency`)
	if err == nil {
		defer pRows.Close()
		for pRows.Next() {
			var id int16
			var val string
			_ = pRows.Scan(&id, &val)
			pollMap[id] = val
		}
	}
	asn1Map := make(map[int16]string)
	asRows, err := conn.Query(ctx, `SELECT id, value FROM public.asn1_type`)
	if err == nil {
		defer asRows.Close()
		for asRows.Next() {
			var id int16
			var val string
			_ = asRows.Scan(&id, &val)
			asn1Map[id] = val
		}
	}
	statusMap := make(map[int16]string)
	stRows, err := conn.Query(ctx, `SELECT id, value FROM public.oid_status`)
	if err == nil {
		defer stRows.Close()
		for stRows.Next() {
			var id int16
			var val string
			_ = stRows.Scan(&id, &val)
			statusMap[id] = val
		}
	}
	oidAccessMap := make(map[int16]string)
	oaRows, err := conn.Query(ctx, `SELECT id, value FROM public.oid_access`)
	if err == nil {
		defer oaRows.Close()
		for oaRows.Next() {
			var id int16
			var val string
			_ = oaRows.Scan(&id, &val)
			oidAccessMap[id] = val
		}
	}
	logicMap := make(map[int16]string)
	loRows, err := conn.Query(ctx, `SELECT id, value FROM public.logic_operator`)
	if err == nil {
		defer loRows.Close()
		for loRows.Next() {
			var id int16
			var val string
			_ = loRows.Scan(&id, &val)
			logicMap[id] = val
		}
	}
	alarmMap := make(map[int16]string)
	alRows, err := conn.Query(ctx, `SELECT id, value FROM public.alarm_level`)
	if err == nil {
		defer alRows.Close()
		for alRows.Next() {
			var id int16
			var val string
			_ = alRows.Scan(&id, &val)
			alarmMap[id] = val
		}
	}
	var vendors []map[string]interface{}
	vdRows, err := conn.Query(ctx, `SELECT id, name, directory FROM public.vendor`)
	if err == nil {
		defer vdRows.Close()
		for vdRows.Next() {
			var id int64
			var name string
			var dir sql.NullString
			if err := vdRows.Scan(&id, &name, &dir); err == nil {
				vendors = append(vendors, map[string]interface{}{
					"id":        id,
					"name":      name,
					"directory": dir.String,
				})
			}
		}
	}
	model.LoadRegistries(
		accessMap,
		varTypeMap,
		pollMap,
		asn1Map,
		statusMap,
		oidAccessMap,
		logicMap,
		alarmMap,
		vendors,
	)
	logger.Info("Все системные справочники успешно загружены из БД в память приложения (%d access, %d types, %d frequencies, %d asn1, %d statuses, %d oid_access, %d logic_ops, %d alarms, %d vendors)",
		len(accessMap), len(varTypeMap), len(pollMap), len(asn1Map), len(statusMap), len(oidAccessMap), len(logicMap), len(alarmMap), len(vendors))
	return nil
}

func stringToNull(s string) sql.NullString {
	return sql.NullString{
		String: s,
		Valid:  s != "",
	}
}

func mapToJSONB(m *map[string]string) ([]byte, error) {
	if m == nil || len(*m) == 0 {
		return nil, nil
	}
	bytes, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

type PollingProtocolData struct {
	ID    int16  `db:"id" json:"id"`
	Value string `db:"value" json:"value"`
}

type AccessData struct {
	ID    int16  `db:"id" json:"id"`
	Value string `db:"value" json:"value"`
}

type VersionSnmpData struct {
	ID    int16  `db:"id" json:"id"`
	Value string `db:"value" json:"value"`
}

type AuthProtocolSnmpData struct {
	ID    int16  `db:"id" json:"id"`
	Value string `db:"value" json:"value"`
}

type PrivacyProtocolSnmpData struct {
	ID    int16  `db:"id" json:"id"`
	Value string `db:"value" json:"value"`
}

type OidTypeData struct {
	ID    int16  `db:"id" json:"id"`
	Value string `db:"value" json:"value"`
}

type LogicOperatorData struct {
	ID    int16  `db:"id" json:"id"`
	Value string `db:"value" json:"value"`
}

type AlarmLevelData struct {
	ID    int16  `db:"id" json:"id"`
	Value string `db:"value" json:"value"`
}

type VarTypeData struct {
	ID    int16  `db:"id" json:"id"`
	Value string `db:"value" json:"value"`
}

type PollingFrequencyData struct {
	ID    int16  `db:"id" json:"id"`
	Value string `db:"value" json:"value"`
}

type OidAccessData struct {
	ID    int16  `db:"id" json:"id"`
	Value string `db:"value" json:"value"`
}

type OidStatusData struct {
	ID    int16  `db:"id" json:"id"`
	Value string `db:"value" json:"value"`
}

type Asn1TypeData struct {
	ID    int16  `db:"id" json:"id"`
	Value string `db:"value" json:"value"`
}

type VendorData struct {
	ID        int64          `db:"id" json:"id"`
	Name      string         `db:"name" json:"name"`
	Number    int32          `db:"number" json:"number"`
	Contact   sql.NullString `db:"contact" json:"contact"`
	Email     sql.NullString `db:"email" json:"email"`
	Directory sql.NullString `db:"directory" json:"directory"`
}

type ComponentData struct {
	ID            int64          `db:"id" json:"id"`
	Title         string         `db:"title" json:"title"`
	NameEn        string         `db:"name_en" json:"name_en"`
	NameRu        string         `db:"name_ru" json:"name_ru"`
	BaseComponent sql.NullInt64  `db:"base_component" json:"base_component"`
	DescriptionEn sql.NullString `db:"description_en" json:"description_en"`
	DescriptionRu sql.NullString `db:"description_ru" json:"description_ru"`
	Access        int16          `db:"access" json:"access"`
}

type ParamData struct {
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

type ComponentParamData struct {
	ComponentID int64 `db:"component_id" json:"component_id"`
	ParamID     int64 `db:"param_id" json:"param_id"`
}

type AgentCapabilitiesData struct {
	ID             int64          `db:"id" json:"id"`
	Name           string         `db:"name" json:"name"`
	ProductRelease sql.NullString `db:"product_release" json:"product_release"`
	Status         sql.NullInt16  `db:"status" json:"status"`
	Description    sql.NullString `db:"description" json:"description"`
	Reference      sql.NullString `db:"reference" json:"reference"`
}

type AgentCapabilitiesModuleData struct {
	ID       int64            `db:"id" json:"id"`
	Supports sql.NullString   `db:"supports" json:"supports"`
	Includes []sql.NullString `db:"includes" json:"includes"`
}

type AgentCapabilitiesModuleNotificationData struct {
	ID          int64          `db:"id" json:"id"`
	Variation   sql.NullString `db:"variation" json:"variation"`
	Access      sql.NullInt16  `db:"access" json:"access"`
	Description sql.NullString `db:"description" json:"description"`
}

type AgentCapabilitiesModuleObjectData struct {
	ID               int64            `db:"id" json:"id"`
	Variation        sql.NullString   `db:"variation" json:"variation"`
	Syntax           sql.NullString   `db:"syntax" json:"syntax"`
	WriteSyntax      sql.NullString   `db:"write_syntax" json:"write_syntax"`
	Access           sql.NullInt16    `db:"access" json:"access"`
	CreationRequires []sql.NullString `db:"creation_requires" json:"creation_requires"`
	Defval           sql.NullString   `db:"defval" json:"defval"`
	Description      sql.NullString   `db:"description" json:"description"`
}

type ChoiceData struct {
	ID      int64  `db:"id" json:"id"`
	Name    string `db:"name" json:"name"`
	Choices string `db:"choices" json:"choices"`
}

type ExplicitData struct {
	ID    int64  `db:"id" json:"id"`
	Name  string `db:"name" json:"name"`
	Value string `db:"value" json:"value"`
}

type ImplicitData struct {
	ID          int64  `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	Application int16  `db:"application" json:"application"`
	Value       string `db:"value" json:"value"`
}

type ImportData struct {
	ID    int64  `db:"id" json:"id"`
	Param string `db:"param" json:"param"`
	From  string `db:"from" json:"from"`
}

type MibData struct {
	ID     int64          `db:"id" json:"id"`
	Path   string         `db:"path" json:"path"`
	Name   sql.NullString `db:"name" json:"name"`
	Vendor sql.NullInt64  `db:"vendor" json:"vendor"`
}

type ModuleComplianceData struct {
	ID          int64          `db:"id" json:"id"`
	Name        string         `db:"name" json:"name"`
	Status      sql.NullInt16  `db:"status" json:"status"`
	Description sql.NullString `db:"description" json:"description"`
	Reference   sql.NullString `db:"reference" json:"reference"`
}

type ModuleComplianceModuleData struct {
	ID              int64            `db:"id" json:"id"`
	Name            string           `db:"name" json:"name"`
	MandatoryGroups []sql.NullString `db:"mandatory_groups" json:"mandatory_groups"`
}

type ModuleComplianceModuleGroupData struct {
	ID          int64          `db:"id" json:"id"`
	Name        string         `db:"name" json:"name"`
	Description sql.NullString `db:"description" json:"description"`
}

type ModuleComplianceModuleObjectData struct {
	ID          int64          `db:"id" json:"id"`
	Name        string         `db:"name" json:"name"`
	Syntax      sql.NullString `db:"syntax" json:"syntax"`
	WriteSyntax sql.NullString `db:"write_syntax" json:"write_syntax"`
	Access      sql.NullInt16  `db:"access" json:"access"`
	Description sql.NullString `db:"description" json:"description"`
}

type ModuleIdentityData struct {
	ID           int64          `db:"id" json:"id"`
	Name         string         `db:"name" json:"name"`
	LastUpdated  sql.NullTime   `db:"last_updated" json:"last_updated"`
	Organization sql.NullString `db:"organization" json:"organization"`
	ContactInfo  sql.NullString `db:"contact_info" json:"contact_info"`
	Description  sql.NullString `db:"description" json:"description"`
}

type NotificationGroupData struct {
	ID            int64            `db:"id" json:"id"`
	Name          string           `db:"name" json:"name"`
	Notifications []sql.NullString `db:"notifications" json:"notifications"`
	Status        sql.NullInt16    `db:"status" json:"status"`
	Description   sql.NullString   `db:"description" json:"description"`
	Reference     sql.NullString   `db:"reference" json:"reference"`
}

type NotificationTypeData struct {
	ID          int64            `db:"id" json:"id"`
	Name        string           `db:"name" json:"name"`
	Objects     []sql.NullString `db:"objects" json:"objects"`
	Status      sql.NullInt16    `db:"status" json:"status"`
	Description sql.NullString   `db:"description" json:"description"`
	Reference   sql.NullString   `db:"reference" json:"reference"`
}

type ObjectGroupData struct {
	ID          int64            `db:"id" json:"id"`
	Name        string           `db:"name" json:"name"`
	Objects     []sql.NullString `db:"objects" json:"objects"`
	Status      sql.NullInt16    `db:"status" json:"status"`
	Description sql.NullString   `db:"description" json:"description"`
	Reference   sql.NullString   `db:"reference" json:"reference"`
}

type ObjectIdentifierData struct {
	ID     int64            `db:"id" json:"id"`
	Name   string           `db:"name" json:"name"`
	NumOid []sql.NullString `db:"num_oid" json:"num_oid"`
	StrOid []sql.NullString `db:"str_oid" json:"str_oid"`
	Parent sql.NullString   `db:"parent" json:"parent"`
	Type   int16            `db:"type" json:"type"`
}

type ObjectIdentityData struct {
	ID          int64          `db:"id" json:"id"`
	Name        string         `db:"name" json:"name"`
	Status      sql.NullInt16  `db:"status" json:"status"`
	Description sql.NullString `db:"description" json:"description"`
	Reference   sql.NullString `db:"reference" json:"reference"`
}

type ObjectTypeData struct {
	ID           int64            `db:"id" json:"id"`
	Name         string           `db:"name" json:"name"`
	Syntax       sql.NullString   `db:"syntax" json:"syntax"`
	Units        sql.NullString   `db:"units" json:"units"`
	Access       sql.NullInt16    `db:"access" json:"access"`
	Status       sql.NullInt16    `db:"status" json:"status"`
	Description  sql.NullString   `db:"description" json:"description"`
	Reference    sql.NullString   `db:"reference" json:"reference"`
	Index        []sql.NullString `db:"index" json:"index"`
	Augments     []sql.NullString `db:"augments" json:"augments"`
	DefaultValue sql.NullString   `db:"default_value" json:"default_value"`
}

type OidData struct {
	ID               string         `db:"id" json:"id"`
	Mib              sql.NullInt64  `db:"mib" json:"mib"`
	Type             int16          `db:"type" json:"type"`
	Name             string         `db:"name" json:"name"`
	Number           sql.NullInt32  `db:"number" json:"number"`
	DotterNotation   string         `db:"dotter_notation" json:"dotter_notation"`
	ObjectDescriptor string         `db:"object_descriptor" json:"object_descriptor"`
	Syntax           sql.NullString `db:"syntax" json:"syntax"`
	Enum             sql.NullString `db:"enum" json:"enum"`
	Status           sql.NullInt16  `db:"status" json:"status"`
	Access           sql.NullInt16  `db:"access" json:"access"`
	Units            sql.NullString `db:"units" json:"units"`
	Description      sql.NullString `db:"description" json:"description"`
	Category         sql.NullString `db:"category" json:"category"`
}

type RevisionData struct {
	ID          int64          `db:"id" json:"id"`
	Revision    sql.NullTime   `db:"revision" json:"revision"`
	Description sql.NullString `db:"description" json:"description"`
}

type SequenceData struct {
	ID    int64  `db:"id" json:"id"`
	Name  string `db:"name" json:"name"`
	Pairs string `db:"pairs" json:"pairs"`
}

type TextualConventionData struct {
	ID          int64          `db:"id" json:"id"`
	Name        string         `db:"name" json:"name"`
	DisplayHint sql.NullString `db:"display_hint" json:"display_hint"`
	Status      sql.NullInt16  `db:"status" json:"status"`
	Description sql.NullString `db:"description" json:"description"`
	Reference   sql.NullString `db:"reference" json:"reference"`
	Syntax      sql.NullString `db:"syntax" json:"syntax"`
}

type TrapTypeData struct {
	ID          int64            `db:"id" json:"id"`
	Name        string           `db:"name" json:"name"`
	Variables   []sql.NullString `db:"variables" json:"variables"`
	Description sql.NullString   `db:"description" json:"description"`
	Reference   sql.NullString   `db:"reference" json:"reference"`
}

type MibToAgentCapabilitiesData struct {
	ID                                    int64         `db:"id" json:"id"`
	MibID                                 int64         `db:"mib_id" json:"mib_id"`
	AgentCapabilitiesID                   int64         `db:"agent_capabilities_id" json:"agent_capabilities_id"`
	AgentCapabilitiesModuleID             sql.NullInt64 `db:"agent_capabilities_module_id" json:"agent_capabilities_module_id"`
	AgentCapabilitiesModuleNotificationID sql.NullInt64 `db:"agent_capabilities_module_notification_id" json:"agent_capabilities_module_notification_id"`
	AgentCapabilitiesModuleObjectID       sql.NullInt64 `db:"agent_capabilities_module_object_id" json:"agent_capabilities_module_object_id"`
}

type MibToChoiceData struct {
	ID       int64 `db:"id" json:"id"`
	MibID    int64 `db:"mib_id" json:"mib_id"`
	ChoiceID int64 `db:"choice_id" json:"choice_id"`
}

type MibToExplicitData struct {
	ID         int64 `db:"id" json:"id"`
	MibID      int64 `db:"mib_id" json:"mib_id"`
	ExplicitID int64 `db:"explicit_id" json:"explicit_id"`
}

type MibToImplicitData struct {
	ID         int64 `db:"id" json:"id"`
	MibID      int64 `db:"mib_id" json:"mib_id"`
	ImplicitID int64 `db:"implicit_id" json:"implicit_id"`
}

type MibToImportData struct {
	ID       int64 `db:"id" json:"id"`
	MibID    int64 `db:"mib_id" json:"mib_id"`
	ImportID int64 `db:"import_id" json:"import_id"`
}

type MibToModuleComplianceData struct {
	ID                             int64         `db:"id" json:"id"`
	MibID                          int64         `db:"mib_id" json:"mib_id"`
	ModuleComplianceID             int64         `db:"module_compliance_id" json:"module_compliance_id"`
	ModuleComplianceModuleID       sql.NullInt64 `db:"module_compliance_module_id" json:"module_compliance_module_id"`
	ModuleComplianceModuleGroupID  sql.NullInt64 `db:"module_compliance_module_group_id" json:"module_compliance_module_group_id"`
	ModuleComplianceModuleObjectID sql.NullInt64 `db:"module_compliance_module_object_id" json:"module_compliance_module_object_id"`
}

type MibToModuleIdentityData struct {
	ID               int64         `db:"id" json:"id"`
	MibID            int64         `db:"mib_id" json:"mib_id"`
	ModuleIdentityID int64         `db:"module_identity_id" json:"module_identity_id"`
	RevisionID       sql.NullInt64 `db:"revision_id" json:"revision_id"`
}

type MibToNotificationGroupData struct {
	ID                  int64 `db:"id" json:"id"`
	MibID               int64 `db:"mib_id" json:"mib_id"`
	NotificationGroupID int64 `db:"notification_group_id" json:"notification_group_id"`
}

type MibToNotificationTypeData struct {
	ID                 int64 `db:"id" json:"id"`
	MibID              int64 `db:"mib_id" json:"mib_id"`
	NotificationTypeID int64 `db:"notification_type_id" json:"notification_type_id"`
}

type MibToObjectGroupData struct {
	ID            int64 `db:"id" json:"id"`
	MibID         int64 `db:"mib_id" json:"mib_id"`
	ObjectGroupID int64 `db:"object_group_id" json:"object_group_id"`
}

type MibToObjectIdentifierData struct {
	ID                 int64 `db:"id" json:"id"`
	MibID              int64 `db:"mib_id" json:"mib_id"`
	ObjectIdentifierID int64 `db:"object_identifier_id" json:"object_identifier_id"`
}

type MibToObjectIdentityData struct {
	ID               int64 `db:"id" json:"id"`
	MibID            int64 `db:"mib_id" json:"mib_id"`
	ObjectIdentityID int64 `db:"object_identity_id" json:"object_identity_id"`
}

type MibToObjectTypeData struct {
	ID           int64 `db:"id" json:"id"`
	MibID        int64 `db:"mib_id" json:"mib_id"`
	ObjectTypeID int64 `db:"object_type_id" json:"object_type_id"`
}

type MibToSequenceData struct {
	ID         int64 `db:"id" json:"id"`
	MibID      int64 `db:"mib_id" json:"mib_id"`
	SequenceID int64 `db:"sequence_id" json:"sequence_id"`
}

type MibToTextualConventionData struct {
	ID                  int64 `db:"id" json:"id"`
	MibID               int64 `db:"mib_id" json:"mib_id"`
	TextualConventionID int64 `db:"textual_convention_id" json:"textual_convention_id"`
}

type MibToTrapTypeData struct {
	ID         int64 `db:"id" json:"id"`
	MibID      int64 `db:"mib_id" json:"mib_id"`
	TrapTypeID int64 `db:"trap_type_id" json:"trap_type_id"`
}

type DeviceIndicatorData struct {
	ID          int64          `db:"id" json:"id"`
	Description sql.NullString `db:"description" json:"description"`
	ObjectID    string         `db:"object_id" json:"object_id"`
	Contact     sql.NullString `db:"contact" json:"contact"`
	Name        sql.NullString `db:"name" json:"name"`
	Location    sql.NullString `db:"location" json:"location"`
	Services    sql.NullInt16  `db:"services" json:"services"`
}

type ParamIndicatorData struct {
	ID             int64          `db:"id" json:"id"`
	OidID          string         `db:"oid_id" json:"oid_id"`
	DotterNotation sql.NullString `db:"dotter_notation" json:"dotter_notation"`
}

type MappingData struct {
	ID          int64           `db:"id" json:"id"`
	Indicator   int64           `db:"indicator" json:"indicator"`
	Param       int64           `db:"param" json:"param"`
	Frequency   int16           `db:"frequency" json:"frequency"`
	Coefficient sql.NullFloat64 `db:"coefficient" json:"coefficient"`
	Enum        sql.NullString  `db:"enum" json:"enum"`
}

type DeviceComponentData struct {
	ID            int64         `db:"id" json:"id"`
	Model         int64         `db:"model" json:"model"`
	InternalOrder sql.NullInt32 `db:"internal_order" json:"internal_order"`
	Parent        sql.NullInt64 `db:"parent" json:"parent"`
}

type DeviceComponentMappingData struct {
	DeviceComponentID int64 `db:"device_component_id" json:"device_component_id"`
	MappingID         int64 `db:"mapping_id" json:"mapping_id"`
}

type ConfigurationData struct {
	ID                int64         `db:"id" json:"id"`
	Indicator         int64         `db:"indicator" json:"indicator"`
	DeviceComponentID sql.NullInt64 `db:"device_component_id" json:"device_component_id"`
}

type DefaultConfigurationData struct {
	ID                int64         `db:"id" json:"id"`
	Indicator         int64         `db:"indicator" json:"indicator"`
	DeviceComponentID sql.NullInt64 `db:"device_component_id" json:"device_component_id"`
}

type ThresholdData struct {
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

type ConfigurationThresholdData struct {
	ConfigurationID int64 `db:"configuration_id" json:"configuration_id"`
	ThresholdID     int64 `db:"threshold_id" json:"threshold_id"`
}

type DefaultConfigurationThresholdData struct {
	DefaultConfigurationID int64 `db:"default_configuration_id" json:"default_configuration_id"`
	ThresholdID            int64 `db:"threshold_id" json:"threshold_id"`
}

type DeviceSnmpData struct {
	Host           string         `db:"host" json:"host"`
	Port           int32          `db:"port" json:"port"`
	Community      string         `db:"community" json:"community"`
	Version        int16          `db:"version" json:"version"`
	Login          sql.NullString `db:"login" json:"login"`
	Password       sql.NullString `db:"password" json:"password"`
	Authentication int16          `db:"authentication" json:"authentication"`
	Privacy        int16          `db:"privacy" json:"privacy"`
	ID             string         `db:"id" json:"id"`
	Config         sql.NullInt64  `db:"config" json:"config"`
}

type ConfigInProcessData struct {
	Host     string `db:"host" json:"host"`
	Port     int32  `db:"port" json:"port"`
	Protocol int16  `db:"protocol" json:"protocol"`
}
