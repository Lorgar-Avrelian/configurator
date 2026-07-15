package dao

import (
	"configurator/internal/database"
	"configurator/internal/logger"
	"context"

	"github.com/jackc/pgx/v5"
)

func GetAllPollingProtocolDao(ctx context.Context) ([]PollingProtocolDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "value"
			 FROM public.polling_protocol
			 ORDER BY "id" ASC`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []PollingProtocolDao
	list = make([]PollingProtocolDao, 0)
	for rows.Next() {
		var d PollingProtocolDao
		err = rows.Scan(
			&d.ID,
			&d.Value,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllAccessDao(ctx context.Context) ([]AccessDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "value"
			 FROM public.access
			 ORDER BY "id" ASC`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []AccessDao
	list = make([]AccessDao, 0)
	for rows.Next() {
		var d AccessDao
		err = rows.Scan(
			&d.ID,
			&d.Value,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllVersionSnmpDao(ctx context.Context) ([]VersionSnmpDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "value"
			 FROM public.version_snmp
			 ORDER BY "id" ASC`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []VersionSnmpDao
	list = make([]VersionSnmpDao, 0)
	for rows.Next() {
		var d VersionSnmpDao
		err = rows.Scan(
			&d.ID,
			&d.Value,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllAuthProtocolSnmpDao(ctx context.Context) ([]AuthProtocolSnmpDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "value"
			 FROM public.auth_protocol_snmp
			 ORDER BY "id" ASC`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []AuthProtocolSnmpDao
	list = make([]AuthProtocolSnmpDao, 0)
	for rows.Next() {
		var d AuthProtocolSnmpDao
		err = rows.Scan(
			&d.ID,
			&d.Value,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllPrivacyProtocolSnmpDao(ctx context.Context) ([]PrivacyProtocolSnmpDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "value"
			 FROM public.privacy_protocol_snmp
			 ORDER BY "id" ASC`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []PrivacyProtocolSnmpDao
	list = make([]PrivacyProtocolSnmpDao, 0)
	for rows.Next() {
		var d PrivacyProtocolSnmpDao
		err = rows.Scan(
			&d.ID,
			&d.Value,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllOidTypeDao(ctx context.Context) ([]OidTypeDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "value"
			 FROM public.oid_type
			 ORDER BY "id" ASC`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []OidTypeDao
	list = make([]OidTypeDao, 0)
	for rows.Next() {
		var d OidTypeDao
		err = rows.Scan(
			&d.ID,
			&d.Value,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllLogicOperatorDao(ctx context.Context) ([]LogicOperatorDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "value", "type", "precedence", "arity"
			 FROM public.logic_operator
			 ORDER BY "id" ASC`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []LogicOperatorDao
	list = make([]LogicOperatorDao, 0)
	for rows.Next() {
		var d LogicOperatorDao
		err = rows.Scan(
			&d.ID,
			&d.Value,
			&d.Type,
			&d.Precedence,
			&d.Arity,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllAlarmLevelDao(ctx context.Context) ([]AlarmLevelDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "value"
			 FROM public.alarm_level
			 ORDER BY "id" ASC`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []AlarmLevelDao
	list = make([]AlarmLevelDao, 0)
	for rows.Next() {
		var d AlarmLevelDao
		err = rows.Scan(
			&d.ID,
			&d.Value,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllVarTypeDao(ctx context.Context) ([]VarTypeDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "value"
			 FROM public.var_type
			 ORDER BY "id" ASC`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []VarTypeDao
	list = make([]VarTypeDao, 0)
	for rows.Next() {
		var d VarTypeDao
		err = rows.Scan(
			&d.ID,
			&d.Value,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllPollingFrequencyDao(ctx context.Context) ([]PollingFrequencyDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "value"
			 FROM public.polling_frequency
			 ORDER BY "id" ASC`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []PollingFrequencyDao
	list = make([]PollingFrequencyDao, 0)
	for rows.Next() {
		var d PollingFrequencyDao
		err = rows.Scan(
			&d.ID,
			&d.Value,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllOidAccessDao(ctx context.Context) ([]OidAccessDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "value"
			 FROM public.oid_access
			 ORDER BY "id" ASC`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []OidAccessDao
	list = make([]OidAccessDao, 0)
	for rows.Next() {
		var d OidAccessDao
		err = rows.Scan(
			&d.ID,
			&d.Value,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllOidStatusDao(ctx context.Context) ([]OidStatusDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "value"
			 FROM public.oid_status
			 ORDER BY "id" ASC`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []OidStatusDao
	list = make([]OidStatusDao, 0)
	for rows.Next() {
		var d OidStatusDao
		err = rows.Scan(
			&d.ID,
			&d.Value,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllAsn1TypeDao(ctx context.Context) ([]Asn1TypeDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "value"
			 FROM public.asn1_type
			 ORDER BY "id" ASC`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []Asn1TypeDao
	list = make([]Asn1TypeDao, 0)
	for rows.Next() {
		var d Asn1TypeDao
		err = rows.Scan(
			&d.ID,
			&d.Value,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllVendorDao(ctx context.Context) ([]VendorDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "name", "number", "contact", "email", "directory"
			 FROM public.vendor
			 ORDER BY "id" ASC`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []VendorDao
	list = make([]VendorDao, 0)
	for rows.Next() {
		var d VendorDao
		err = rows.Scan(
			&d.ID,
			&d.Name,
			&d.Number,
			&d.Contact,
			&d.Email,
			&d.Directory,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllComponentDao(ctx context.Context) ([]ComponentDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "title", "name_en", "name_ru", "plural_name_en", "plural_name_ru", "base_component", "description_en", "description_ru", "access"
			 FROM public.component
			 ORDER BY "id" ASC`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []ComponentDao
	list = make([]ComponentDao, 0)
	for rows.Next() {
		var d ComponentDao
		err = rows.Scan(
			&d.ID,
			&d.Title,
			&d.NameEn,
			&d.NameRu,
			&d.PluralNameEn,
			&d.PluralNameRu,
			&d.BaseComponent,
			&d.DescriptionEn,
			&d.DescriptionRu,
			&d.Access,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllParamDao(ctx context.Context) ([]ParamDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "title", "name_en", "name_ru", "type", "value", "description_en", "description_ru", "units_en", "units_ru", "access", "saved", "visible", "diagram"
			 FROM public.param
			 ORDER BY "id" ASC`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []ParamDao
	list = make([]ParamDao, 0)
	for rows.Next() {
		var d ParamDao
		err = rows.Scan(
			&d.ID,
			&d.Title,
			&d.NameEn,
			&d.NameRu,
			&d.Type,
			&d.Value,
			&d.DescriptionEn,
			&d.DescriptionRu,
			&d.UnitsEn,
			&d.UnitsRu,
			&d.Access,
			&d.Saved,
			&d.Visible,
			&d.Diagram,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllComponentParamDao(ctx context.Context) ([]ComponentParamDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "component_id", "param_id"
			 FROM public.component_param`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []ComponentParamDao
	list = make([]ComponentParamDao, 0)
	for rows.Next() {
		var d ComponentParamDao
		err = rows.Scan(
			&d.ComponentID,
			&d.ParamID,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllAgentCapabilitiesDao(ctx context.Context) ([]AgentCapabilitiesDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "name", "product_release", "status", "description", "reference"
			 FROM public.agent_capabilities
			 ORDER BY "id" ASC`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []AgentCapabilitiesDao
	list = make([]AgentCapabilitiesDao, 0)
	for rows.Next() {
		var d AgentCapabilitiesDao
		err = rows.Scan(
			&d.ID,
			&d.Name,
			&d.ProductRelease,
			&d.Status,
			&d.Description,
			&d.Reference,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllAgentCapabilitiesModuleDao(ctx context.Context) ([]AgentCapabilitiesModuleDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "supports", "includes"
			 FROM public.agent_capabilities_module
			 ORDER BY "id" ASC`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []AgentCapabilitiesModuleDao
	list = make([]AgentCapabilitiesModuleDao, 0)
	for rows.Next() {
		var d AgentCapabilitiesModuleDao
		err = rows.Scan(
			&d.ID,
			&d.Supports,
			&d.Includes,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllAgentCapabilitiesModuleNotificationDao(ctx context.Context) ([]AgentCapabilitiesModuleNotificationDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "variation", "access", "description"
			 FROM public.agent_capabilities_module_notification
			 ORDER BY "id" ASC`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []AgentCapabilitiesModuleNotificationDao
	list = make([]AgentCapabilitiesModuleNotificationDao, 0)
	for rows.Next() {
		var d AgentCapabilitiesModuleNotificationDao
		err = rows.Scan(
			&d.ID,
			&d.Variation,
			&d.Access,
			&d.Description,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllAgentCapabilitiesModuleObjectDao(ctx context.Context) ([]AgentCapabilitiesModuleObjectDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "variation", "syntax", "write_syntax", "access", "creation_requires", "defval", "description"
			 FROM public.agent_capabilities_module_object
			 ORDER BY "id" ASC`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []AgentCapabilitiesModuleObjectDao
	list = make([]AgentCapabilitiesModuleObjectDao, 0)
	for rows.Next() {
		var d AgentCapabilitiesModuleObjectDao
		err = rows.Scan(
			&d.ID,
			&d.Variation,
			&d.Syntax,
			&d.WriteSyntax,
			&d.Access,
			&d.CreationRequires,
			&d.Defval,
			&d.Description,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllChoiceDao(ctx context.Context) ([]ChoiceDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "name", "choices"
			 FROM public.choice
			 ORDER BY "id" ASC`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []ChoiceDao
	list = make([]ChoiceDao, 0)
	for rows.Next() {
		var d ChoiceDao
		err = rows.Scan(
			&d.ID,
			&d.Name,
			&d.Choices,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllExplicitDao(ctx context.Context) ([]ExplicitDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "name", "value"
			 FROM public.explicit
			 ORDER BY "id" ASC`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []ExplicitDao
	list = make([]ExplicitDao, 0)
	for rows.Next() {
		var d ExplicitDao
		err = rows.Scan(
			&d.ID,
			&d.Name,
			&d.Value,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllImplicitDao(ctx context.Context) ([]ImplicitDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "name", "application", "value"
			 FROM public.implicit
			 ORDER BY "id" ASC`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []ImplicitDao
	list = make([]ImplicitDao, 0)
	for rows.Next() {
		var d ImplicitDao
		err = rows.Scan(
			&d.ID,
			&d.Name,
			&d.Application,
			&d.Value,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllImportDao(ctx context.Context) ([]ImportDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "param", "from"
			 FROM public.import
			 ORDER BY "id" ASC`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []ImportDao
	list = make([]ImportDao, 0)
	for rows.Next() {
		var d ImportDao
		err = rows.Scan(
			&d.ID,
			&d.Param,
			&d.From,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllMibDao(ctx context.Context) ([]MibDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "path", "name", "vendor"
			 FROM public.mib
			 ORDER BY "id" ASC`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []MibDao
	list = make([]MibDao, 0)
	for rows.Next() {
		var d MibDao
		err = rows.Scan(
			&d.ID,
			&d.Path,
			&d.Name,
			&d.Vendor,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllModuleComplianceDao(ctx context.Context) ([]ModuleComplianceDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "name", "status", "description", "reference"
			 FROM public.module_compliance
			 ORDER BY "id" ASC`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []ModuleComplianceDao
	list = make([]ModuleComplianceDao, 0)
	for rows.Next() {
		var d ModuleComplianceDao
		err = rows.Scan(
			&d.ID,
			&d.Name,
			&d.Status,
			&d.Description,
			&d.Reference,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllModuleComplianceModuleDao(ctx context.Context) ([]ModuleComplianceModuleDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "name", "mandatory_groups"
			 FROM public.module_compliance_module
			 ORDER BY "id" ASC`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []ModuleComplianceModuleDao
	list = make([]ModuleComplianceModuleDao, 0)
	for rows.Next() {
		var d ModuleComplianceModuleDao
		err = rows.Scan(
			&d.ID,
			&d.Name,
			&d.MandatoryGroups,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllModuleComplianceModuleGroupDao(ctx context.Context) ([]ModuleComplianceModuleGroupDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "name", "description"
			 FROM public.module_compliance_module_group
			 ORDER BY "id" ASC`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []ModuleComplianceModuleGroupDao
	list = make([]ModuleComplianceModuleGroupDao, 0)
	for rows.Next() {
		var d ModuleComplianceModuleGroupDao
		err = rows.Scan(
			&d.ID,
			&d.Name,
			&d.Description,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllModuleComplianceModuleObjectDao(ctx context.Context) ([]ModuleComplianceModuleObjectDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "name", "syntax", "write_syntax", "access", "description"
			 FROM public.module_compliance_module_object
			 ORDER BY "id" ASC`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []ModuleComplianceModuleObjectDao
	list = make([]ModuleComplianceModuleObjectDao, 0)
	for rows.Next() {
		var d ModuleComplianceModuleObjectDao
		err = rows.Scan(
			&d.ID,
			&d.Name,
			&d.Syntax,
			&d.WriteSyntax,
			&d.Access,
			&d.Description,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllModuleIdentityDao(ctx context.Context) ([]ModuleIdentityDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "name", "last_updated", "organization", "contact_info", "description"
			 FROM public.module_identity
			 ORDER BY "id" ASC`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []ModuleIdentityDao
	list = make([]ModuleIdentityDao, 0)
	for rows.Next() {
		var d ModuleIdentityDao
		err = rows.Scan(
			&d.ID,
			&d.Name,
			&d.LastUpdated,
			&d.Organization,
			&d.ContactInfo,
			&d.Description,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllNotificationGroupDao(ctx context.Context) ([]NotificationGroupDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "name", "notifications", "status", "description", "reference"
			 FROM public.notification_group
			 ORDER BY "id" ASC`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []NotificationGroupDao
	list = make([]NotificationGroupDao, 0)
	for rows.Next() {
		var d NotificationGroupDao
		err = rows.Scan(
			&d.ID,
			&d.Name,
			&d.Notifications,
			&d.Status,
			&d.Description,
			&d.Reference,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllNotificationTypeDao(ctx context.Context) ([]NotificationTypeDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "name", "objects", "status", "description", "reference"
			 FROM public.notification_type
			 ORDER BY "id" ASC`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []NotificationTypeDao
	list = make([]NotificationTypeDao, 0)
	for rows.Next() {
		var d NotificationTypeDao
		err = rows.Scan(
			&d.ID,
			&d.Name,
			&d.Objects,
			&d.Status,
			&d.Description,
			&d.Reference,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllObjectGroupDao(ctx context.Context) ([]ObjectGroupDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "name", "objects", "status", "description", "reference"
			 FROM public.object_group
			 ORDER BY "id" ASC`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []ObjectGroupDao
	list = make([]ObjectGroupDao, 0)
	for rows.Next() {
		var d ObjectGroupDao
		err = rows.Scan(
			&d.ID,
			&d.Name,
			&d.Objects,
			&d.Status,
			&d.Description,
			&d.Reference,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllObjectIdentifierDao(ctx context.Context) ([]ObjectIdentifierDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "name", "num_oid", "str_oid", "parent", "type"
			 FROM public.object_identifier
			 ORDER BY "id" ASC`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []ObjectIdentifierDao
	list = make([]ObjectIdentifierDao, 0)
	for rows.Next() {
		var d ObjectIdentifierDao
		err = rows.Scan(
			&d.ID,
			&d.Name,
			&d.NumOid,
			&d.StrOid,
			&d.Parent,
			&d.Type,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllObjectIdentityDao(ctx context.Context) ([]ObjectIdentityDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "name", "status", "description", "reference"
			 FROM public.object_identity
			 ORDER BY "id" ASC`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []ObjectIdentityDao
	list = make([]ObjectIdentityDao, 0)
	for rows.Next() {
		var d ObjectIdentityDao
		err = rows.Scan(
			&d.ID,
			&d.Name,
			&d.Status,
			&d.Description,
			&d.Reference,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllObjectTypeDao(ctx context.Context) ([]ObjectTypeDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "name", "syntax", "units", "access", "status", "description", "reference", "index", "augments", "default_value"
			 FROM public.object_type
			 ORDER BY "id" ASC`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []ObjectTypeDao
	list = make([]ObjectTypeDao, 0)
	for rows.Next() {
		var d ObjectTypeDao
		err = rows.Scan(
			&d.ID,
			&d.Name,
			&d.Syntax,
			&d.Units,
			&d.Access,
			&d.Status,
			&d.Description,
			&d.Reference,
			&d.Index,
			&d.Augments,
			&d.DefaultValue,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllOidDao(ctx context.Context) ([]OidDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "mib", "type", "name", "number", "dotter_notation", "object_descriptor", "syntax", "enum", "status", "access", "units", "description", "category"
			 FROM public.oid
			 ORDER BY "id" ASC`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []OidDao
	list = make([]OidDao, 0)
	for rows.Next() {
		var d OidDao
		err = rows.Scan(
			&d.ID,
			&d.Mib,
			&d.Type,
			&d.Name,
			&d.Number,
			&d.DotterNotation,
			&d.ObjectDescriptor,
			&d.Syntax,
			&d.Enum,
			&d.Status,
			&d.Access,
			&d.Units,
			&d.Description,
			&d.Category,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllRevisionDao(ctx context.Context) ([]RevisionDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "revision", "description"
			 FROM public.revision
			 ORDER BY "id" ASC`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []RevisionDao
	list = make([]RevisionDao, 0)
	for rows.Next() {
		var d RevisionDao
		err = rows.Scan(
			&d.ID,
			&d.Revision,
			&d.Description,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllSequenceDao(ctx context.Context) ([]SequenceDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "name", "pairs"
			 FROM public.sequence
			 ORDER BY "id" ASC`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []SequenceDao
	list = make([]SequenceDao, 0)
	for rows.Next() {
		var d SequenceDao
		err = rows.Scan(
			&d.ID,
			&d.Name,
			&d.Pairs,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllTextualConventionDao(ctx context.Context) ([]TextualConventionDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "name", "display_hint", "status", "description", "reference", "syntax"
			 FROM public.textual_convention
			 ORDER BY "id" ASC`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []TextualConventionDao
	list = make([]TextualConventionDao, 0)
	for rows.Next() {
		var d TextualConventionDao
		err = rows.Scan(
			&d.ID,
			&d.Name,
			&d.DisplayHint,
			&d.Status,
			&d.Description,
			&d.Reference,
			&d.Syntax,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllTrapTypeDao(ctx context.Context) ([]TrapTypeDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "name", "variables", "description", "reference"
			 FROM public.trap_type
			 ORDER BY "id" ASC`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []TrapTypeDao
	list = make([]TrapTypeDao, 0)
	for rows.Next() {
		var d TrapTypeDao
		err = rows.Scan(
			&d.ID,
			&d.Name,
			&d.Variables,
			&d.Description,
			&d.Reference,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllMibToAgentCapabilitiesDao(ctx context.Context) ([]MibToAgentCapabilitiesDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "mib_id", "agent_capabilities_id", "agent_capabilities_module_id", "agent_capabilities_module_notification_id", "agent_capabilities_module_object_id"
			 FROM public.mib_to_agent_capabilities
			 ORDER BY "id" ASC`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []MibToAgentCapabilitiesDao
	list = make([]MibToAgentCapabilitiesDao, 0)
	for rows.Next() {
		var d MibToAgentCapabilitiesDao
		err = rows.Scan(
			&d.ID,
			&d.MibID,
			&d.AgentCapabilitiesID,
			&d.AgentCapabilitiesModuleID,
			&d.AgentCapabilitiesModuleNotificationID,
			&d.AgentCapabilitiesModuleObjectID,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllMibToChoiceDao(ctx context.Context) ([]MibToChoiceDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "mib_id", "choice_id"
			 FROM public.mib_to_choice
			 ORDER BY "id" ASC`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []MibToChoiceDao
	list = make([]MibToChoiceDao, 0)
	for rows.Next() {
		var d MibToChoiceDao
		err = rows.Scan(
			&d.ID,
			&d.MibID,
			&d.ChoiceID,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllMibToExplicitDao(ctx context.Context) ([]MibToExplicitDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "mib_id", "explicit_id"
			 FROM public.mib_to_explicit
			 ORDER BY "id" ASC`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []MibToExplicitDao
	list = make([]MibToExplicitDao, 0)
	for rows.Next() {
		var d MibToExplicitDao
		err = rows.Scan(
			&d.ID,
			&d.MibID,
			&d.ExplicitID,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllMibToImplicitDao(ctx context.Context) ([]MibToImplicitDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "mib_id", "implicit_id"
			 FROM public.mib_to_implicit
			 ORDER BY "id" ASC`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []MibToImplicitDao
	list = make([]MibToImplicitDao, 0)
	for rows.Next() {
		var d MibToImplicitDao
		err = rows.Scan(
			&d.ID,
			&d.MibID,
			&d.ImplicitID,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllMibToImportDao(ctx context.Context) ([]MibToImportDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "mib_id", "import_id"
			 FROM public.mib_to_import
			 ORDER BY "id" ASC`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []MibToImportDao
	list = make([]MibToImportDao, 0)
	for rows.Next() {
		var d MibToImportDao
		err = rows.Scan(
			&d.ID,
			&d.MibID,
			&d.ImportID,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllMibToModuleComplianceDao(ctx context.Context) ([]MibToModuleComplianceDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "mib_id", "module_compliance_id", "module_compliance_module_id", "module_compliance_module_group_id", "module_compliance_module_object_id"
			 FROM public.mib_to_module_compliance
			 ORDER BY "id" ASC`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []MibToModuleComplianceDao
	list = make([]MibToModuleComplianceDao, 0)
	for rows.Next() {
		var d MibToModuleComplianceDao
		err = rows.Scan(
			&d.ID,
			&d.MibID,
			&d.ModuleComplianceID,
			&d.ModuleComplianceModuleID,
			&d.ModuleComplianceModuleGroupID,
			&d.ModuleComplianceModuleObjectID,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllMibToModuleIdentityDao(ctx context.Context) ([]MibToModuleIdentityDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "mib_id", "module_identity_id", "revision_id"
			 FROM public.mib_to_module_identity
			 ORDER BY "id" ASC`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []MibToModuleIdentityDao
	list = make([]MibToModuleIdentityDao, 0)
	for rows.Next() {
		var d MibToModuleIdentityDao
		err = rows.Scan(
			&d.ID,
			&d.MibID,
			&d.ModuleIdentityID,
			&d.RevisionID,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllMibToNotificationGroupDao(ctx context.Context) ([]MibToNotificationGroupDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "mib_id", "notification_group_id"
			 FROM public.mib_to_notification_group
			 ORDER BY "id" ASC`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []MibToNotificationGroupDao
	list = make([]MibToNotificationGroupDao, 0)
	for rows.Next() {
		var d MibToNotificationGroupDao
		err = rows.Scan(
			&d.ID,
			&d.MibID,
			&d.NotificationGroupID,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllMibToNotificationTypeDao(ctx context.Context) ([]MibToNotificationTypeDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "mib_id", "notification_type_id"
			 FROM public.mib_to_notification_type
			 ORDER BY "id" ASC`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []MibToNotificationTypeDao
	list = make([]MibToNotificationTypeDao, 0)
	for rows.Next() {
		var d MibToNotificationTypeDao
		err = rows.Scan(
			&d.ID,
			&d.MibID,
			&d.NotificationTypeID,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllMibToObjectGroupDao(ctx context.Context) ([]MibToObjectGroupDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "mib_id", "object_group_id"
			 FROM public.mib_to_object_group
			 ORDER BY "id" ASC`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []MibToObjectGroupDao
	list = make([]MibToObjectGroupDao, 0)
	for rows.Next() {
		var d MibToObjectGroupDao
		err = rows.Scan(
			&d.ID,
			&d.MibID,
			&d.ObjectGroupID,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllMibToObjectIdentifierDao(ctx context.Context) ([]MibToObjectIdentifierDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "mib_id", "object_identifier_id"
			 FROM public.mib_to_object_identifier
			 ORDER BY "id" ASC`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []MibToObjectIdentifierDao
	list = make([]MibToObjectIdentifierDao, 0)
	for rows.Next() {
		var d MibToObjectIdentifierDao
		err = rows.Scan(
			&d.ID,
			&d.MibID,
			&d.ObjectIdentifierID,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllMibToObjectIdentityDao(ctx context.Context) ([]MibToObjectIdentityDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "mib_id", "object_identity_id"
			 FROM public.mib_to_object_identity
			 ORDER BY "id" ASC`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []MibToObjectIdentityDao
	list = make([]MibToObjectIdentityDao, 0)
	for rows.Next() {
		var d MibToObjectIdentityDao
		err = rows.Scan(
			&d.ID,
			&d.MibID,
			&d.ObjectIdentityID,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllMibToObjectTypeDao(ctx context.Context) ([]MibToObjectTypeDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "mib_id", "object_type_id"
			 FROM public.mib_to_object_type
			 ORDER BY "id" ASC`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []MibToObjectTypeDao
	list = make([]MibToObjectTypeDao, 0)
	for rows.Next() {
		var d MibToObjectTypeDao
		err = rows.Scan(
			&d.ID,
			&d.MibID,
			&d.ObjectTypeID,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllMibToSequenceDao(ctx context.Context) ([]MibToSequenceDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "mib_id", "sequence_id"
			 FROM public.mib_to_sequence
			 ORDER BY "id" ASC`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []MibToSequenceDao
	list = make([]MibToSequenceDao, 0)
	for rows.Next() {
		var d MibToSequenceDao
		err = rows.Scan(
			&d.ID,
			&d.MibID,
			&d.SequenceID,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllMibToTextualConventionDao(ctx context.Context) ([]MibToTextualConventionDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "mib_id", "textual_convention_id"
			 FROM public.mib_to_textual_convention
			 ORDER BY "id" ASC`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []MibToTextualConventionDao
	list = make([]MibToTextualConventionDao, 0)
	for rows.Next() {
		var d MibToTextualConventionDao
		err = rows.Scan(
			&d.ID,
			&d.MibID,
			&d.TextualConventionID,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllMibToTrapTypeDao(ctx context.Context) ([]MibToTrapTypeDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "mib_id", "trap_type_id"
			 FROM public.mib_to_trap_type
			 ORDER BY "id" ASC`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []MibToTrapTypeDao
	list = make([]MibToTrapTypeDao, 0)
	for rows.Next() {
		var d MibToTrapTypeDao
		err = rows.Scan(
			&d.ID,
			&d.MibID,
			&d.TrapTypeID,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllDeviceIndicatorDao(ctx context.Context) ([]DeviceIndicatorDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "description", "object_id", "contact", "name", "location", "services"
			 FROM public.device_indicator
			 ORDER BY "id" ASC`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []DeviceIndicatorDao
	list = make([]DeviceIndicatorDao, 0)
	for rows.Next() {
		var d DeviceIndicatorDao
		err = rows.Scan(
			&d.ID,
			&d.Description,
			&d.ObjectID,
			&d.Contact,
			&d.Name,
			&d.Location,
			&d.Services,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllParamIndicatorDao(ctx context.Context) ([]ParamIndicatorDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "oid_id", "dotter_notation"
			 FROM public.param_indicator
			 ORDER BY "id" ASC`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []ParamIndicatorDao
	list = make([]ParamIndicatorDao, 0)
	for rows.Next() {
		var d ParamIndicatorDao
		err = rows.Scan(
			&d.ID,
			&d.OidID,
			&d.DotterNotation,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllMappingDao(ctx context.Context) ([]MappingDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "indicator", "param", "frequency", "value", "coefficient", "enum", "position", "from", "position_type"
			 FROM public.mapping
			 ORDER BY "id" ASC`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []MappingDao
	list = make([]MappingDao, 0)
	for rows.Next() {
		var d MappingDao
		err = rows.Scan(
			&d.ID,
			&d.Indicator,
			&d.Param,
			&d.Frequency,
			&d.Value,
			&d.Coefficient,
			&d.Enum,
			&d.Position,
			&d.From,
			&d.PositionType,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllDeviceComponentDao(ctx context.Context) ([]DeviceComponentDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "model", "internal_order", "parent"
			 FROM public.device_component
			 ORDER BY "id" ASC`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []DeviceComponentDao
	list = make([]DeviceComponentDao, 0)
	for rows.Next() {
		var d DeviceComponentDao
		err = rows.Scan(
			&d.ID,
			&d.Model,
			&d.InternalOrder,
			&d.Parent,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllDeviceComponentMappingDao(ctx context.Context) ([]DeviceComponentMappingDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "device_component_id", "mapping_id"
			 FROM public.device_component_mapping`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []DeviceComponentMappingDao
	list = make([]DeviceComponentMappingDao, 0)
	for rows.Next() {
		var d DeviceComponentMappingDao
		err = rows.Scan(
			&d.DeviceComponentID,
			&d.MappingID,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllConfigurationDao(ctx context.Context) ([]ConfigurationDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "indicator", "device_component_id"
			 FROM public.configuration
			 ORDER BY "id" ASC`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []ConfigurationDao
	list = make([]ConfigurationDao, 0)
	for rows.Next() {
		var d ConfigurationDao
		err = rows.Scan(
			&d.ID,
			&d.Indicator,
			&d.DeviceComponentID,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllDefaultConfigurationDao(ctx context.Context) ([]DefaultConfigurationDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "indicator", "device_component_id"
			 FROM public.default_configuration
			 ORDER BY "id" ASC`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []DefaultConfigurationDao
	list = make([]DefaultConfigurationDao, 0)
	for rows.Next() {
		var d DefaultConfigurationDao
		err = rows.Scan(
			&d.ID,
			&d.Indicator,
			&d.DeviceComponentID,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllThresholdDao(ctx context.Context) ([]ThresholdDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "name", "description", "author", "created", "query"::text, "target"::text, "value"
	FROM public.threshold
	ORDER BY "id" ASC`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		logger.Errorf("Failed to execute select query for all thresholds: %v", err)
		return nil, err
	}
	defer rows.Close()
	var list []ThresholdDao
	list = []ThresholdDao{}
	for rows.Next() {
		var d ThresholdDao
		err = rows.Scan(
			&d.ID,
			&d.Name,
			&d.Description,
			&d.Author,
			&d.Created,
			&d.Query,
			&d.Target,
			&d.Value,
		)
		if err != nil {
			logger.Errorf("Failed to scan row into threshold dao structure: %v", err)
			return nil, err
		}
		list = append(list, d)
	}
	err = rows.Err()
	if err != nil {
		logger.Errorf("Threshold rows iterator processing error encountered: %v", err)
		return nil, err
	}
	return list, nil
}

func GetAllDeviceSnmpDao(ctx context.Context) ([]DeviceSnmpDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "host", "port", "community", "version", "login", "password", "authentication", "privacy", "id", "config"
			 FROM public.device_snmp`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []DeviceSnmpDao
	list = make([]DeviceSnmpDao, 0)
	for rows.Next() {
		var d DeviceSnmpDao
		err = rows.Scan(
			&d.Host,
			&d.Port,
			&d.Community,
			&d.Version,
			&d.Login,
			&d.Password,
			&d.Authentication,
			&d.Privacy,
			&d.ID,
			&d.Config,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllResultDao(ctx context.Context) ([]ResultDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "host", "port", "component", "internal_order", "param", "value"
			 FROM public.result`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []ResultDao
	list = make([]ResultDao, 0)
	for rows.Next() {
		var d ResultDao
		err = rows.Scan(
			&d.Host,
			&d.Port,
			&d.Component,
			&d.InternalOrder,
			&d.Param,
			&d.Value,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllAffectedThresholdDao(ctx context.Context) ([]AffectedThresholdDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "host", "port", "threshold", "enabled"
			 FROM public.affected_threshold
			 ORDER BY "id" ASC`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []AffectedThresholdDao
	list = make([]AffectedThresholdDao, 0)
	for rows.Next() {
		var d AffectedThresholdDao
		err = rows.Scan(
			&d.ID,
			&d.Host,
			&d.Port,
			&d.Threshold,
			&d.Enabled,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllAffectedParamDao(ctx context.Context) ([]AffectedParamDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "host", "port", "component", "internal_order", "param"
			 FROM public.affected_param
			 ORDER BY "id" ASC`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []AffectedParamDao
	list = make([]AffectedParamDao, 0)
	for rows.Next() {
		var d AffectedParamDao
		err = rows.Scan(
			&d.ID,
			&d.Host,
			&d.Port,
			&d.Component,
			&d.InternalOrder,
			&d.Param,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllConfigInProcessDao(ctx context.Context) ([]ConfigInProcessDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "host", "port", "protocol"
			 FROM public.config_in_process`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []ConfigInProcessDao
	list = make([]ConfigInProcessDao, 0)
	for rows.Next() {
		var d ConfigInProcessDao
		err = rows.Scan(
			&d.Host,
			&d.Port,
			&d.Protocol,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func GetAllCChangeLogDao(ctx context.Context) ([]ChangeLogDao, error) {
	var rows pgx.Rows
	var err error
	var query string
	query = `SELECT "id", "script", "executed"
			 FROM public.changelog`
	rows, err = database.Get().Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []ChangeLogDao
	list = make([]ChangeLogDao, 0)
	for rows.Next() {
		var d ChangeLogDao
		err = rows.Scan(
			&d.ID,
			&d.Script,
			&d.Executed,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}
