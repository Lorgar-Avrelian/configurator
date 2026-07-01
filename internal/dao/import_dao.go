package dao

import (
	"configurator/internal/database"
	"context"

	"github.com/jackc/pgx/v5"
)

const batchSize = 5000

func dropArray[T any](list []T) [][]T {
	var batches [][]T
	var start int
	var end int
	if batchSize > len(list) {
		batches = [][]T{list}
		return batches
	}
	batches = make([][]T, 0, (len(list)+batchSize-1)/batchSize)
	start = 0
	for len(list) > start {
		end = start + batchSize
		if end > len(list) {
			end = len(list)
		}
		batches = append(batches, list[start:end])
		start = start + batchSize
	}
	return batches
}

func executeBatchInsert[T any](ctx context.Context, batches [][]T, insertFunction func(pgx.Batch, T) pgx.Batch) error {
	var br pgx.BatchResults
	var err error
	var batch pgx.Batch
	var chunk []T
	var item T
	var i int
	var j int
	var k int
	i = 0
	for len(batches) > i {
		chunk = batches[i]
		batch = pgx.Batch{}
		j = 0
		for len(chunk) > j {
			item = chunk[j]
			batch = insertFunction(batch, item)
			j = j + 1
		}
		br = database.Get().SendBatch(ctx, &batch)
		k = 0
		for batch.Len() > k {
			_, err = br.Exec()
			if err != nil {
				br.Close()
				return err
			}
			k = k + 1
		}
		err = br.Close()
		if err != nil {
			return err
		}
		i = i + 1
	}
	return nil
}

func ImportPollingProtocolDao(ctx context.Context, list []PollingProtocolDao) error {
	var batches [][]PollingProtocolDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.polling_protocol ("id", "value") 
			 VALUES ($1, $2)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item PollingProtocolDao) pgx.Batch {
		b.Queue(query, item.ID, item.Value)
		return b
	})
}

func ImportAccessDao(ctx context.Context, list []AccessDao) error {
	var batches [][]AccessDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.access ("id", "value") 
			 VALUES ($1, $2)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item AccessDao) pgx.Batch {
		b.Queue(query, item.ID, item.Value)
		return b
	})
}

func ImportVersionSnmpDao(ctx context.Context, list []VersionSnmpDao) error {
	var batches [][]VersionSnmpDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.version_snmp ("id", "value") 
			 VALUES ($1, $2)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item VersionSnmpDao) pgx.Batch {
		b.Queue(query, item.ID, item.Value)
		return b
	})
}

func ImportAuthProtocolSnmpDao(ctx context.Context, list []AuthProtocolSnmpDao) error {
	var batches [][]AuthProtocolSnmpDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.auth_protocol_snmp ("id", "value") 
			 VALUES ($1, $2)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item AuthProtocolSnmpDao) pgx.Batch {
		b.Queue(query, item.ID, item.Value)
		return b
	})
}

func ImportPrivacyProtocolSnmpDao(ctx context.Context, list []PrivacyProtocolSnmpDao) error {
	var batches [][]PrivacyProtocolSnmpDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.privacy_protocol_snmp ("id", "value") 
			 VALUES ($1, $2)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item PrivacyProtocolSnmpDao) pgx.Batch {
		b.Queue(query, item.ID, item.Value)
		return b
	})
}

func ImportOidTypeDao(ctx context.Context, list []OidTypeDao) error {
	var batches [][]OidTypeDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.oid_type ("id", "value") 
			 VALUES ($1, $2)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item OidTypeDao) pgx.Batch {
		b.Queue(query, item.ID, item.Value)
		return b
	})
}

func ImportLogicOperatorDao(ctx context.Context, list []LogicOperatorDao) error {
	var batches [][]LogicOperatorDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.logic_operator ("id", "value", "type", "precedence", "arity") 
			 VALUES ($1, $2, $3, $4, $5)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item LogicOperatorDao) pgx.Batch {
		b.Queue(query, item.ID, item.Value, item.Type, item.Precedence, item.Arity)
		return b
	})
}

func ImportAlarmLevelDao(ctx context.Context, list []AlarmLevelDao) error {
	var batches [][]AlarmLevelDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.alarm_level ("id", "value") 
			 VALUES ($1, $2)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item AlarmLevelDao) pgx.Batch {
		b.Queue(query, item.ID, item.Value)
		return b
	})
}

func ImportVarTypeDao(ctx context.Context, list []VarTypeDao) error {
	var batches [][]VarTypeDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.var_type ("id", "value") 
			 VALUES ($1, $2)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item VarTypeDao) pgx.Batch {
		b.Queue(query, item.ID, item.Value)
		return b
	})
}

func ImportPollingFrequencyDao(ctx context.Context, list []PollingFrequencyDao) error {
	var batches [][]PollingFrequencyDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.polling_frequency ("id", "value") 
			 VALUES ($1, $2)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item PollingFrequencyDao) pgx.Batch {
		b.Queue(query, item.ID, item.Value)
		return b
	})
}

func ImportOidAccessDao(ctx context.Context, list []OidAccessDao) error {
	var batches [][]OidAccessDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.oid_access ("id", "value") 
			 VALUES ($1, $2)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item OidAccessDao) pgx.Batch {
		b.Queue(query, item.ID, item.Value)
		return b
	})
}

func ImportOidStatusDao(ctx context.Context, list []OidStatusDao) error {
	var batches [][]OidStatusDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.oid_status ("id", "value") 
			 VALUES ($1, $2)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item OidStatusDao) pgx.Batch {
		b.Queue(query, item.ID, item.Value)
		return b
	})
}

func ImportAsn1TypeDao(ctx context.Context, list []Asn1TypeDao) error {
	var batches [][]Asn1TypeDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.asn1_type ("id", "value") 
			 VALUES ($1, $2)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item Asn1TypeDao) pgx.Batch {
		b.Queue(query, item.ID, item.Value)
		return b
	})
}

func ImportVendorDao(ctx context.Context, list []VendorDao) error {
	var batches [][]VendorDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.vendor ("id", "name", "number", "contact", "email", "directory") 
			 VALUES ($1, $2, $3, $4, $5, $6)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item VendorDao) pgx.Batch {
		b.Queue(query, item.ID, item.Name, item.Number, item.Contact, item.Email, item.Directory)
		return b
	})
}

func ImportComponentDao(ctx context.Context, list []ComponentDao) error {
	var batches [][]ComponentDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.component ("id", "title", "name_en", "name_ru", "plural_name_en", "plural_name_ru", "base_component", "description_en", "description_ru", "access") 
			 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item ComponentDao) pgx.Batch {
		b.Queue(query, item.ID, item.Title, item.NameEn, item.NameRu, item.PluralNameEn, item.PluralNameRu, item.BaseComponent, item.DescriptionEn, item.DescriptionRu, item.Access)
		return b
	})
}

func ImportParamDao(ctx context.Context, list []ParamDao) error {
	var batches [][]ParamDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.param ("id", "title", "name_en", "name_ru", "type", "value", "description_en", "description_ru", "units_en", "units_ru", "access", "saved", "visible") 
			 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item ParamDao) pgx.Batch {
		b.Queue(query, item.ID, item.Title, item.NameEn, item.NameRu, item.Type, item.Value, item.DescriptionEn, item.DescriptionRu, item.UnitsEn, item.UnitsRu, item.Access, item.Saved, item.Visible)
		return b
	})
}

func ImportComponentParamDao(ctx context.Context, list []ComponentParamDao) error {
	var batches [][]ComponentParamDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.component_param ("component_id", "param_id") 
			 VALUES ($1, $2)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item ComponentParamDao) pgx.Batch {
		b.Queue(query, item.ComponentID, item.ParamID)
		return b
	})
}

func ImportAgentCapabilitiesDao(ctx context.Context, list []AgentCapabilitiesDao) error {
	var batches [][]AgentCapabilitiesDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.agent_capabilities ("id", "name", "product_release", "status", "description", "reference") 
			 VALUES ($1, $2, $3, $4, $5, $6)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item AgentCapabilitiesDao) pgx.Batch {
		b.Queue(query, item.ID, item.Name, item.ProductRelease, item.Status, item.Description, item.Reference)
		return b
	})
}

func ImportAgentCapabilitiesModuleDao(ctx context.Context, list []AgentCapabilitiesModuleDao) error {
	var batches [][]AgentCapabilitiesModuleDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.agent_capabilities_module ("id", "supports", "includes") 
			 VALUES ($1, $2, $3)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item AgentCapabilitiesModuleDao) pgx.Batch {
		b.Queue(query, item.ID, item.Supports, item.Includes)
		return b
	})
}

func ImportAgentCapabilitiesModuleNotificationDao(ctx context.Context, list []AgentCapabilitiesModuleNotificationDao) error {
	var batches [][]AgentCapabilitiesModuleNotificationDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.agent_capabilities_module_notification ("id", "variation", "access", "description") 
			 VALUES ($1, $2, $3, $4)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item AgentCapabilitiesModuleNotificationDao) pgx.Batch {
		b.Queue(query, item.ID, item.Variation, item.Access, item.Description)
		return b
	})
}

func ImportAgentCapabilitiesModuleObjectDao(ctx context.Context, list []AgentCapabilitiesModuleObjectDao) error {
	var batches [][]AgentCapabilitiesModuleObjectDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.agent_capabilities_module_object ("id", "variation", "syntax", "write_syntax", "access", "creation_requires", "defval", "description") 
			 VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item AgentCapabilitiesModuleObjectDao) pgx.Batch {
		b.Queue(query, item.ID, item.Variation, item.Syntax, item.WriteSyntax, item.Access, item.CreationRequires, item.Defval, item.Description)
		return b
	})
}

func ImportChoiceDao(ctx context.Context, list []ChoiceDao) error {
	var batches [][]ChoiceDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.choice ("id", "name", "choices") 
			 VALUES ($1, $2, $3)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item ChoiceDao) pgx.Batch {
		var val interface{}
		val = nil
		if len(item.Choices) > 0 {
			val = string(item.Choices)
		}
		b.Queue(query, item.ID, item.Name, val)
		return b
	})
}

func ImportExplicitDao(ctx context.Context, list []ExplicitDao) error {
	var batches [][]ExplicitDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.explicit ("id", "name", "value") 
			 VALUES ($1, $2, $3)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item ExplicitDao) pgx.Batch {
		b.Queue(query, item.ID, item.Name, item.Value)
		return b
	})
}

func ImportImplicitDao(ctx context.Context, list []ImplicitDao) error {
	var batches [][]ImplicitDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.implicit ("id", "name", "application", "value") 
			 VALUES ($1, $2, $3, $4)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item ImplicitDao) pgx.Batch {
		b.Queue(query, item.ID, item.Name, item.Application, item.Value)
		return b
	})
}

func ImportImportDao(ctx context.Context, list []ImportDao) error {
	var batches [][]ImportDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.import ("id", "param", "from") 
			 VALUES ($1, $2, $3)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item ImportDao) pgx.Batch {
		b.Queue(query, item.ID, item.Param, item.From)
		return b
	})
}

func ImportMibDao(ctx context.Context, list []MibDao) error {
	var batches [][]MibDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.mib ("id", "path", "name", "vendor") 
			 VALUES ($1, $2, $3, $4)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item MibDao) pgx.Batch {
		b.Queue(query, item.ID, item.Path, item.Name, item.Vendor)
		return b
	})
}

func ImportModuleComplianceDao(ctx context.Context, list []ModuleComplianceDao) error {
	var batches [][]ModuleComplianceDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.module_compliance ("id", "name", "status", "description", "reference") 
			 VALUES ($1, $2, $3, $4, $5)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item ModuleComplianceDao) pgx.Batch {
		b.Queue(query, item.ID, item.Name, item.Status, item.Description, item.Reference)
		return b
	})
}

func ImportModuleComplianceModuleDao(ctx context.Context, list []ModuleComplianceModuleDao) error {
	var batches [][]ModuleComplianceModuleDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.module_compliance_module ("id", "name", "mandatory_groups") 
			 VALUES ($1, $2, $3)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item ModuleComplianceModuleDao) pgx.Batch {
		b.Queue(query, item.ID, item.Name, item.MandatoryGroups)
		return b
	})
}

func ImportModuleComplianceModuleGroupDao(ctx context.Context, list []ModuleComplianceModuleGroupDao) error {
	var batches [][]ModuleComplianceModuleGroupDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.module_compliance_module_group ("id", "name", "description") 
			 VALUES ($1, $2, $3)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item ModuleComplianceModuleGroupDao) pgx.Batch {
		b.Queue(query, item.ID, item.Name, item.Description)
		return b
	})
}

func ImportModuleComplianceModuleObjectDao(ctx context.Context, list []ModuleComplianceModuleObjectDao) error {
	var batches [][]ModuleComplianceModuleObjectDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.module_compliance_module_object ("id", "name", "syntax", "write_syntax", "access", "description") 
			 VALUES ($1, $2, $3, $4, $5, $6)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item ModuleComplianceModuleObjectDao) pgx.Batch {
		b.Queue(query, item.ID, item.Name, item.Syntax, item.WriteSyntax, item.Access, item.Description)
		return b
	})
}

func ImportModuleIdentityDao(ctx context.Context, list []ModuleIdentityDao) error {
	var batches [][]ModuleIdentityDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.module_identity ("id", "name", "last_updated", "organization", "contact_info", "description") 
			 VALUES ($1, $2, $3, $4, $5, $6)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item ModuleIdentityDao) pgx.Batch {
		b.Queue(query, item.ID, item.Name, item.LastUpdated, item.Organization, item.ContactInfo, item.Description)
		return b
	})
}

func ImportNotificationGroupDao(ctx context.Context, list []NotificationGroupDao) error {
	var batches [][]NotificationGroupDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.notification_group ("id", "name", "notifications", "status", "description", "reference") 
			 VALUES ($1, $2, $3, $4, $5, $6)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item NotificationGroupDao) pgx.Batch {
		b.Queue(query, item.ID, item.Name, item.Notifications, item.Status, item.Description, item.Reference)
		return b
	})
}

func ImportNotificationTypeDao(ctx context.Context, list []NotificationTypeDao) error {
	var batches [][]NotificationTypeDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.notification_type ("id", "name", "objects", "status", "description", "reference") 
			 VALUES ($1, $2, $3, $4, $5, $6)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item NotificationTypeDao) pgx.Batch {
		b.Queue(query, item.ID, item.Name, item.Objects, item.Status, item.Description, item.Reference)
		return b
	})
}

func ImportObjectGroupDao(ctx context.Context, list []ObjectGroupDao) error {
	var batches [][]ObjectGroupDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.object_group ("id", "name", "objects", "status", "description", "reference") 
			 VALUES ($1, $2, $3, $4, $5, $6)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item ObjectGroupDao) pgx.Batch {
		b.Queue(query, item.ID, item.Name, item.Objects, item.Status, item.Description, item.Reference)
		return b
	})
}

func ImportObjectIdentifierDao(ctx context.Context, list []ObjectIdentifierDao) error {
	var batches [][]ObjectIdentifierDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.object_identifier ("id", "name", "num_oid", "str_oid", "parent", "type") 
			 VALUES ($1, $2, $3, $4, $5, $6)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item ObjectIdentifierDao) pgx.Batch {
		var flatNum []interface{}
		var flatStr []interface{}
		var m int
		flatNum = make([]interface{}, len(item.NumOid))
		m = 0
		for len(item.NumOid) > m {
			flatNum[m] = nil
			if item.NumOid[m] != nil {
				flatNum[m] = *item.NumOid[m]
			}
			m = m + 1
		}
		flatStr = make([]interface{}, len(item.StrOid))
		m = 0
		for len(item.StrOid) > m {
			flatStr[m] = nil
			if item.StrOid[m] != nil {
				flatStr[m] = *item.StrOid[m]
			}
			m = m + 1
		}
		b.Queue(query, item.ID, item.Name, flatNum, flatStr, item.Parent, item.Type)
		return b
	})
}

func ImportObjectIdentityDao(ctx context.Context, list []ObjectIdentityDao) error {
	var batches [][]ObjectIdentityDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.object_identity ("id", "name", "status", "description", "reference") 
			 VALUES ($1, $2, $3, $4, $5)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item ObjectIdentityDao) pgx.Batch {
		b.Queue(query, item.ID, item.Name, item.Status, item.Description, item.Reference)
		return b
	})
}

func ImportObjectTypeDao(ctx context.Context, list []ObjectTypeDao) error {
	var batches [][]ObjectTypeDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.object_type ("id", "name", "syntax", "units", "access", "status", "description", "reference", "index", "augments", "default_value") 
			 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item ObjectTypeDao) pgx.Batch {
		b.Queue(query, item.ID, item.Name, item.Syntax, item.Units, item.Access, item.Status, item.Description, item.Reference, item.Index, item.Augments, item.DefaultValue)
		return b
	})
}

func ImportOidDao(ctx context.Context, list []OidDao) error {
	var batches [][]OidDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.oid ("id", "mib", "type", "name", "number", "dotter_notation", "object_descriptor", "syntax", "enum", "status", "access", "units", "description", "category") 
			 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item OidDao) pgx.Batch {
		var val interface{}
		val = nil
		if len(item.Enum) > 0 {
			val = string(item.Enum)
		}
		b.Queue(query, item.ID, item.Mib, item.Type, item.Name, item.Number, item.DotterNotation, item.ObjectDescriptor, item.Syntax, val, item.Status, item.Access, item.Units, item.Description, item.Category)
		return b
	})
}

func ImportRevisionDao(ctx context.Context, list []RevisionDao) error {
	var batches [][]RevisionDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.revision ("id", "revision", "description") 
			 VALUES ($1, $2, $3)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item RevisionDao) pgx.Batch {
		b.Queue(query, item.ID, item.Revision, item.Description)
		return b
	})
}

func ImportSequenceDao(ctx context.Context, list []SequenceDao) error {
	var batches [][]SequenceDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.sequence ("id", "name", "pairs") 
			 VALUES ($1, $2, $3)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item SequenceDao) pgx.Batch {
		var val interface{}
		val = nil
		if len(item.Pairs) > 0 {
			val = string(item.Pairs)
		}
		b.Queue(query, item.ID, item.Name, val)
		return b
	})
}

func ImportTextualConventionDao(ctx context.Context, list []TextualConventionDao) error {
	var batches [][]TextualConventionDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.textual_convention ("id", "name", "display_hint", "status", "description", "reference", "syntax") 
			 VALUES ($1, $2, $3, $4, $5, $6, $7)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item TextualConventionDao) pgx.Batch {
		b.Queue(query, item.ID, item.Name, item.DisplayHint, item.Status, item.Description, item.Reference, item.Syntax)
		return b
	})
}

func ImportTrapTypeDao(ctx context.Context, list []TrapTypeDao) error {
	var batches [][]TrapTypeDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.trap_type ("id", "name", "variables", "description", "reference") 
			 VALUES ($1, $2, $3, $4, $5)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item TrapTypeDao) pgx.Batch {
		b.Queue(query, item.ID, item.Name, item.Variables, item.Description, item.Reference)
		return b
	})
}

func ImportMibToAgentCapabilitiesDao(ctx context.Context, list []MibToAgentCapabilitiesDao) error {
	var batches [][]MibToAgentCapabilitiesDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.mib_to_agent_capabilities ("id", "mib_id", "agent_capabilities_id", "agent_capabilities_module_id", "agent_capabilities_module_notification_id", "agent_capabilities_module_object_id") 
			 VALUES ($1, $2, $3, $4, $5, $6)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item MibToAgentCapabilitiesDao) pgx.Batch {
		b.Queue(query, item.ID, item.MibID, item.AgentCapabilitiesID, item.AgentCapabilitiesModuleID, item.AgentCapabilitiesModuleNotificationID, item.AgentCapabilitiesModuleObjectID)
		return b
	})
}

func ImportMibToChoiceDao(ctx context.Context, list []MibToChoiceDao) error {
	var batches [][]MibToChoiceDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.mib_to_choice ("id", "mib_id", "choice_id") 
			 VALUES ($1, $2, $3)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item MibToChoiceDao) pgx.Batch {
		b.Queue(query, item.ID, item.MibID, item.ChoiceID)
		return b
	})
}

func ImportMibToExplicitDao(ctx context.Context, list []MibToExplicitDao) error {
	var batches [][]MibToExplicitDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.mib_to_explicit ("id", "mib_id", "explicit_id") 
			 VALUES ($1, $2, $3)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item MibToExplicitDao) pgx.Batch {
		b.Queue(query, item.ID, item.MibID, item.ExplicitID)
		return b
	})
}

func ImportMibToImplicitDao(ctx context.Context, list []MibToImplicitDao) error {
	var batches [][]MibToImplicitDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.mib_to_implicit ("id", "mib_id", "implicit_id") 
			 VALUES ($1, $2, $3)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item MibToImplicitDao) pgx.Batch {
		b.Queue(query, item.ID, item.MibID, item.ImplicitID)
		return b
	})
}

func ImportMibToImportDao(ctx context.Context, list []MibToImportDao) error {
	var batches [][]MibToImportDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.mib_to_import ("id", "mib_id", "import_id") 
			 VALUES ($1, $2, $3)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item MibToImportDao) pgx.Batch {
		b.Queue(query, item.ID, item.MibID, item.ImportID)
		return b
	})
}

func ImportMibToModuleComplianceDao(ctx context.Context, list []MibToModuleComplianceDao) error {
	var batches [][]MibToModuleComplianceDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.mib_to_module_compliance ("id", "mib_id", "module_compliance_id", "module_compliance_module_id", "module_compliance_module_group_id", "module_compliance_module_object_id") 
			 VALUES ($1, $2, $3, $4, $5, $6)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item MibToModuleComplianceDao) pgx.Batch {
		b.Queue(query, item.ID, item.MibID, item.ModuleComplianceID, item.ModuleComplianceModuleID, item.ModuleComplianceModuleGroupID, item.ModuleComplianceModuleObjectID)
		return b
	})
}

func ImportMibToModuleIdentityDao(ctx context.Context, list []MibToModuleIdentityDao) error {
	var batches [][]MibToModuleIdentityDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.mib_to_module_identity ("id", "mib_id", "module_identity_id", "revision_id") 
			 VALUES ($1, $2, $3, $4)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item MibToModuleIdentityDao) pgx.Batch {
		b.Queue(query, item.ID, item.MibID, item.ModuleIdentityID, item.RevisionID)
		return b
	})
}

func ImportMibToNotificationGroupDao(ctx context.Context, list []MibToNotificationGroupDao) error {
	var batches [][]MibToNotificationGroupDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.mib_to_notification_group ("id", "mib_id", "notification_group_id") 
			 VALUES ($1, $2, $3)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item MibToNotificationGroupDao) pgx.Batch {
		b.Queue(query, item.ID, item.MibID, item.NotificationGroupID)
		return b
	})
}

func ImportMibToNotificationTypeDao(ctx context.Context, list []MibToNotificationTypeDao) error {
	var batches [][]MibToNotificationTypeDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.mib_to_notification_type ("id", "mib_id", "notification_type_id") 
			 VALUES ($1, $2, $3)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item MibToNotificationTypeDao) pgx.Batch {
		b.Queue(query, item.ID, item.MibID, item.NotificationTypeID)
		return b
	})
}

func ImportMibToObjectGroupDao(ctx context.Context, list []MibToObjectGroupDao) error {
	var batches [][]MibToObjectGroupDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.mib_to_object_group ("id", "mib_id", "object_group_id") 
			 VALUES ($1, $2, $3)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item MibToObjectGroupDao) pgx.Batch {
		b.Queue(query, item.ID, item.MibID, item.ObjectGroupID)
		return b
	})
}

func ImportMibToObjectIdentifierDao(ctx context.Context, list []MibToObjectIdentifierDao) error {
	var batches [][]MibToObjectIdentifierDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.mib_to_object_identifier ("id", "mib_id", "object_identifier_id") 
			 VALUES ($1, $2, $3)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item MibToObjectIdentifierDao) pgx.Batch {
		b.Queue(query, item.ID, item.MibID, item.ObjectIdentifierID)
		return b
	})
}

func ImportMibToObjectIdentityDao(ctx context.Context, list []MibToObjectIdentityDao) error {
	var batches [][]MibToObjectIdentityDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.mib_to_object_identity ("id", "mib_id", "object_identity_id") 
			 VALUES ($1, $2, $3)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item MibToObjectIdentityDao) pgx.Batch {
		b.Queue(query, item.ID, item.MibID, item.ObjectIdentityID)
		return b
	})
}

func ImportMibToObjectTypeDao(ctx context.Context, list []MibToObjectTypeDao) error {
	var batches [][]MibToObjectTypeDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.mib_to_object_type ("id", "mib_id", "object_type_id") 
			 VALUES ($1, $2, $3)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item MibToObjectTypeDao) pgx.Batch {
		b.Queue(query, item.ID, item.MibID, item.ObjectTypeID)
		return b
	})
}

func ImportMibToSequenceDao(ctx context.Context, list []MibToSequenceDao) error {
	var batches [][]MibToSequenceDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.mib_to_sequence ("id", "mib_id", "sequence_id") 
			 VALUES ($1, $2, $3)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item MibToSequenceDao) pgx.Batch {
		b.Queue(query, item.ID, item.MibID, item.SequenceID)
		return b
	})
}

func ImportMibToTextualConventionDao(ctx context.Context, list []MibToTextualConventionDao) error {
	var batches [][]MibToTextualConventionDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.mib_to_textual_convention ("id", "mib_id", "textual_convention_id") 
			 VALUES ($1, $2, $3)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item MibToTextualConventionDao) pgx.Batch {
		b.Queue(query, item.ID, item.MibID, item.TextualConventionID)
		return b
	})
}

func ImportMibToTrapTypeDao(ctx context.Context, list []MibToTrapTypeDao) error {
	var batches [][]MibToTrapTypeDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.mib_to_trap_type ("id", "mib_id", "trap_type_id") 
			 VALUES ($1, $2, $3)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item MibToTrapTypeDao) pgx.Batch {
		b.Queue(query, item.ID, item.MibID, item.TrapTypeID)
		return b
	})
}

func ImportDeviceIndicatorDao(ctx context.Context, list []DeviceIndicatorDao) error {
	var batches [][]DeviceIndicatorDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.device_indicator ("id", "description", "object_id", "contact", "name", "location", "services") 
			 VALUES ($1, $2, $3, $4, $5, $6, $7)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item DeviceIndicatorDao) pgx.Batch {
		b.Queue(query, item.ID, item.Description, item.ObjectID, item.Contact, item.Name, item.Location, item.Services)
		return b
	})
}

func ImportParamIndicatorDao(ctx context.Context, list []ParamIndicatorDao) error {
	var batches [][]ParamIndicatorDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.param_indicator ("id", "oid_id", "dotter_notation") 
			 VALUES ($1, $2, $3)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item ParamIndicatorDao) pgx.Batch {
		b.Queue(query, item.ID, item.OidID, item.DotterNotation)
		return b
	})
}

func ImportDeviceComponentDao(ctx context.Context, list []DeviceComponentDao) error {
	var batches [][]DeviceComponentDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.device_component ("id", "model", "internal_order", "parent") 
			 VALUES ($1, $2, $3, $4)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item DeviceComponentDao) pgx.Batch {
		b.Queue(query, item.ID, item.Model, item.InternalOrder, item.Parent)
		return b
	})
}

func ImportDeviceComponentMappingDao(ctx context.Context, list []DeviceComponentMappingDao) error {
	var batches [][]DeviceComponentMappingDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.device_component_mapping ("device_component_id", "mapping_id") 
			 VALUES ($1, $2)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item DeviceComponentMappingDao) pgx.Batch {
		b.Queue(query, item.DeviceComponentID, item.MappingID)
		return b
	})
}

func ImportConfigurationDao(ctx context.Context, list []ConfigurationDao) error {
	var batches [][]ConfigurationDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.configuration ("id", "indicator", "device_component_id") 
			 VALUES ($1, $2, $3)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item ConfigurationDao) pgx.Batch {
		b.Queue(query, item.ID, item.Indicator, item.DeviceComponentID)
		return b
	})
}

func ImportDefaultConfigurationDao(ctx context.Context, list []DefaultConfigurationDao) error {
	var batches [][]DefaultConfigurationDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.default_configuration ("id", "indicator", "device_component_id") 
			 VALUES ($1, $2, $3)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item DefaultConfigurationDao) pgx.Batch {
		b.Queue(query, item.ID, item.Indicator, item.DeviceComponentID)
		return b
	})
}

func ImportThresholdDao(ctx context.Context, list []ThresholdDao) error {
	var batches [][]ThresholdDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.threshold ("id", "name", "description", "author", "created", "query") 
			 VALUES ($1, $2, $3, $4, $5, $6)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item ThresholdDao) pgx.Batch {
		b.Queue(query, item.ID, item.Name, item.Description, item.Author, item.Created, item.Query)
		return b
	})
}

func ImportDeviceSnmpDao(ctx context.Context, list []DeviceSnmpDao) error {
	var batches [][]DeviceSnmpDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.device_snmp ("host", "port", "community", "version", "login", "password", "authentication", "privacy", "id", "config") 
			 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item DeviceSnmpDao) pgx.Batch {
		b.Queue(query, item.Host, item.Port, item.Community, item.Version, item.Login, item.Password, item.Authentication, item.Privacy, item.ID, item.Config)
		return b
	})
}

func ImportResultDao(ctx context.Context, list []ResultDao) error {
	var batches [][]ResultDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.result ("host", "port", "component", "internal_order", "param", "value") 
			 VALUES ($1, $2, $3, $4, $5, $6)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item ResultDao) pgx.Batch {
		b.Queue(query, item.Host, item.Port, item.Component, item.InternalOrder, item.Param, item.Value)
		return b
	})
}

func ImportAffectedThresholdDao(ctx context.Context, list []AffectedThresholdDao) error {
	var batches [][]AffectedThresholdDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.affected_threshold ("id", "host", "port", "threshold", "enabled") 
			 VALUES ($1, $2, $3, $4, $5)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item AffectedThresholdDao) pgx.Batch {
		b.Queue(query, item.ID, item.Host, item.Port, item.Threshold, item.Enabled)
		return b
	})
}

func ImportAffectedParamDao(ctx context.Context, list []AffectedParamDao) error {
	var batches [][]AffectedParamDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.affected_param ("id", "host", "port", "component", "internal_order", "param") 
			 VALUES ($1, $2, $3, $4, $5, $6)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item AffectedParamDao) pgx.Batch {
		b.Queue(query, item.ID, item.Host, item.Port, item.Component, item.InternalOrder, item.Param)
		return b
	})
}

func ImportConfigInProcessDao(ctx context.Context, list []ConfigInProcessDao) error {
	var batches [][]ConfigInProcessDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.config_in_process ("host", "port", "protocol") 
			 VALUES ($1, $2, $3)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item ConfigInProcessDao) pgx.Batch {
		b.Queue(query, item.Host, item.Port, item.Protocol)
		return b
	})
}

func ImportMappingDao(ctx context.Context, list []MappingDao) error {
	var batches [][]MappingDao
	var query string
	batches = dropArray(list)
	query = `INSERT INTO public.mapping ("id", "indicator", "param", "frequency", "value", "coefficient", "enum", "position", "from", "position_type") 
			 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	return executeBatchInsert(ctx, batches, func(b pgx.Batch, item MappingDao) pgx.Batch {
		var val interface{}
		val = nil
		if len(item.Enum) > 0 {
			val = string(item.Enum)
		}
		b.Queue(query, item.ID, item.Indicator, item.Param, item.Frequency, item.Value, item.Coefficient, val, item.Position, item.From, item.PositionType)
		return b
	})
}
