package dao

import (
	"configurator/internal/database"
	"configurator/internal/logger"
	"context"
)

func DropPollingProtocolDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.polling_protocol`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropAccessDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.access`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropVersionSnmpDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.version_snmp`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropAuthProtocolSnmpDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.auth_protocol_snmp`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropPrivacyProtocolSnmpDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.privacy_protocol_snmp`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropOidTypeDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.oid_type`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropLogicOperatorDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.logic_operator`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropAlarmLevelDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.alarm_level`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropVarTypeDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.var_type`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropPollingFrequencyDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.polling_frequency`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropOidAccessDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.oid_access`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropOidStatusDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.oid_status`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropAsn1TypeDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.asn1_type`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropVendorDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.vendor`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropComponentDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.component`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropParamDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.param`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropComponentParamDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.component_param`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropAgentCapabilitiesDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.agent_capabilities`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropAgentCapabilitiesModuleDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.agent_capabilities_module`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropAgentCapabilitiesModuleNotificationDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.agent_capabilities_module_notification`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropAgentCapabilitiesModuleObjectDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.agent_capabilities_module_object`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropChoiceDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.choice`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropExplicitDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.explicit`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropImplicitDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.implicit`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropImportDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.import`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropMibDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.mib`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropModuleComplianceDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.module_compliance`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropModuleComplianceModuleDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.module_compliance_module`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropModuleComplianceModuleGroupDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.module_compliance_module_group`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropModuleComplianceModuleObjectDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.module_compliance_module_object`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropModuleIdentityDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.module_identity`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropNotificationGroupDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.notification_group`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropNotificationTypeDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.notification_type`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropObjectGroupDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.object_group`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropObjectIdentifierDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.object_identifier`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropObjectIdentityDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.object_identity`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropObjectTypeDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.object_type`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropOidDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.oid`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropRevisionDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.revision`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropSequenceDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.sequence`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropTextualConventionDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.textual_convention`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropTrapTypeDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.trap_type`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropMibToAgentCapabilitiesDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.mib_to_agent_capabilities`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropMibToChoiceDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.mib_to_choice`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropMibToExplicitDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.mib_to_explicit`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropMibToImplicitDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.mib_to_implicit`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropMibToImportDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.mib_to_import`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropMibToModuleComplianceDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.mib_to_module_compliance`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropMibToModuleIdentityDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.mib_to_module_identity`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropMibToNotificationGroupDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.mib_to_notification_group`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropMibToNotificationTypeDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.mib_to_notification_type`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropMibToObjectGroupDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.mib_to_object_group`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropMibToObjectIdentifierDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.mib_to_object_identifier`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropMibToObjectIdentityDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.mib_to_object_identity`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropMibToObjectTypeDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.mib_to_object_type`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropMibToSequenceDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.mib_to_sequence`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropMibToTextualConventionDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.mib_to_textual_convention`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropMibToTrapTypeDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.mib_to_trap_type`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropDeviceIndicatorDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.device_indicator`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropParamIndicatorDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.param_indicator`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropMappingDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.mapping`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropDeviceComponentDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.device_component`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropDeviceComponentMappingDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.device_component_mapping`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropConfigurationDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.configuration`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropDefaultConfigurationDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.default_configuration`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropThresholdDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.threshold`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropDeviceSnmpDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.device_snmp`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropResultDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.result`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropAffectedThresholdDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.affected_threshold`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropAffectedParamDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.affected_param`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropConfigInProcessDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.config_in_process`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropChangeLogDao(ctx context.Context) error {
	var err error
	var query string
	query = `DROP TABLE IF EXISTS public.changelog`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropLogicOperatorDaoConstraints(ctx context.Context) error {
	var err error
	var query string
	query = `ALTER TABLE public.logic_operator
    DROP CONSTRAINT "logic_operator_value_key",
    DROP CONSTRAINT "logic_operator_type_check",
    DROP CONSTRAINT "logic_operator_arity_check";`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropOidAccessDaoConstraints(ctx context.Context) error {
	var err error
	var query string
	query = `ALTER TABLE public.oid_access
    DROP CONSTRAINT "oid_access_value_key";`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropOidStatusDaoConstraints(ctx context.Context) error {
	var err error
	var query string
	query = `ALTER TABLE public.oid_status
    DROP CONSTRAINT "oid_status_value_key";`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropAsn1TypeDaoConstraints(ctx context.Context) error {
	var err error
	var query string
	query = `ALTER TABLE public.asn1_type
    DROP CONSTRAINT "asn1_type_value_key";`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropVendorDaoConstraints(ctx context.Context) error {
	var err error
	var query string
	query = `ALTER TABLE public.vendor
    DROP CONSTRAINT "vendor_number_key";`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropMibDaoConstraints(ctx context.Context) error {
	var err error
	var query string
	query = `ALTER TABLE public.mib
    DROP CONSTRAINT "mib_path_key",
    DROP CONSTRAINT "mib_vendor_fkey";`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropComponentDaoConstraints(ctx context.Context) error {
	var err error
	var query string
	query = `ALTER TABLE public.component
    DROP CONSTRAINT "component_base_component_fkey",
    DROP CONSTRAINT "component_access_fkey";`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropParamDaoConstraints(ctx context.Context) error {
	var err error
	var query string
	query = `ALTER TABLE public.param
    DROP CONSTRAINT "param_type_fkey",
    DROP CONSTRAINT "param_access_fkey";`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropComponentParamDaoConstraints(ctx context.Context) error {
	var err error
	var query string
	query = `ALTER TABLE public.component_param
    DROP CONSTRAINT "component_param_component_id_fkey",
    DROP CONSTRAINT "component_param_param_id_fkey";`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropAgentCapabilitiesDaoConstraints(ctx context.Context) error {
	var err error
	var query string
	query = `ALTER TABLE public.agent_capabilities
    DROP CONSTRAINT "agent_capabilities_status_fkey";`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropAgentCapabilitiesModuleNotificationDaoConstraints(ctx context.Context) error {
	var err error
	var query string
	query = `ALTER TABLE public.agent_capabilities_module_notification
    DROP CONSTRAINT "agent_capabilities_module_notification_access_fkey";`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropAgentCapabilitiesModuleObjectDaoConstraints(ctx context.Context) error {
	var err error
	var query string
	query = `ALTER TABLE public.agent_capabilities_module_object
    DROP CONSTRAINT "agent_capabilities_module_object_access_fkey";`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropModuleComplianceDaoConstraints(ctx context.Context) error {
	var err error
	var query string
	query = `ALTER TABLE public.module_compliance
    DROP CONSTRAINT "module_compliance_status_fkey";`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropModuleComplianceModuleObjectDaoConstraints(ctx context.Context) error {
	var err error
	var query string
	query = `ALTER TABLE public.module_compliance_module_object
    DROP CONSTRAINT "module_compliance_module_object_access_fkey";`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropNotificationGroupDaoConstraints(ctx context.Context) error {
	var err error
	var query string
	query = `ALTER TABLE public.notification_group
    DROP CONSTRAINT "notification_group_status_fkey";`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropNotificationTypeDaoConstraints(ctx context.Context) error {
	var err error
	var query string
	query = `ALTER TABLE public.notification_type
    DROP CONSTRAINT "notification_type_status_fkey";`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropObjectGroupDaoConstraints(ctx context.Context) error {
	var err error
	var query string
	query = `ALTER TABLE public.object_group
    DROP CONSTRAINT "object_group_status_fkey";`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropObjectIdentifierDaoConstraints(ctx context.Context) error {
	var err error
	var query string
	query = `ALTER TABLE public.object_identifier
    DROP CONSTRAINT "object_identifier_type_fkey";`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropObjectIdentityDaoConstraints(ctx context.Context) error {
	var err error
	var query string
	query = `ALTER TABLE public.object_identity
    DROP CONSTRAINT "object_identity_status_fkey";`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropObjectTypeDaoConstraints(ctx context.Context) error {
	var err error
	var query string
	query = `ALTER TABLE public.object_type
    DROP CONSTRAINT "object_type_access_fkey",
    DROP CONSTRAINT "object_type_status_fkey";`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropOidDaoConstraints(ctx context.Context) error {
	var err error
	var query string
	query = `ALTER TABLE public.oid
    DROP CONSTRAINT "oid_mib_fkey",
    DROP CONSTRAINT "oid_type_fkey",
    DROP CONSTRAINT "oid_status_fkey",
    DROP CONSTRAINT "oid_access_fkey";`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropTextualConventionDaoConstraints(ctx context.Context) error {
	var err error
	var query string
	query = `ALTER TABLE public.textual_convention
    DROP CONSTRAINT "textual_convention_status_fkey";`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropMibToAgentCapabilitiesDaoConstraints(ctx context.Context) error {
	var err error
	var query string
	query = `ALTER TABLE public.mib_to_agent_capabilities
    DROP CONSTRAINT "mib_to_agent_capabilities_mib_id_fkey",
    DROP CONSTRAINT "mib_to_agent_capabilities_agent_capabilities_id_fkey",
    DROP CONSTRAINT "mib_to_agent_capabilities_agent_capabilities_module_id_fkey",
    DROP CONSTRAINT "mib_to_agent_capabilities_agent_capabilities_module_notification_id_fkey",
    DROP CONSTRAINT "mib_to_agent_capabilities_agent_capabilities_module_object_id_fkey",
    DROP CONSTRAINT "mib_ac";`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropMibToChoiceDaoConstraints(ctx context.Context) error {
	var err error
	var query string
	query = `ALTER TABLE public.mib_to_choice
    DROP CONSTRAINT "mib_to_choice_mib_id_fkey",
    DROP CONSTRAINT "mib_to_choice_choice_id_fkey",
    DROP CONSTRAINT "mib_c";`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropMibToExplicitDaoConstraints(ctx context.Context) error {
	var err error
	var query string
	query = `ALTER TABLE public.mib_to_explicit
    DROP CONSTRAINT "mib_to_explicit_mib_id_fkey",
    DROP CONSTRAINT "mib_to_explicit_explicit_id_fkey",
    DROP CONSTRAINT "mib_et";`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropMibToImplicitDaoConstraints(ctx context.Context) error {
	var err error
	var query string
	query = `ALTER TABLE public.mib_to_implicit
    DROP CONSTRAINT "mib_to_implicit_mib_id_fkey",
    DROP CONSTRAINT "mib_to_implicit_implicit_id_fkey",
    DROP CONSTRAINT "mib_it";`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropMibToImportDaoConstraints(ctx context.Context) error {
	var err error
	var query string
	query = `ALTER TABLE public.mib_to_import
    DROP CONSTRAINT "mib_to_import_mib_id_fkey",
    DROP CONSTRAINT "mib_to_import_import_id_fkey",
    DROP CONSTRAINT "mib_imp";`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropMibToModuleComplianceDaoConstraints(ctx context.Context) error {
	var err error
	var query string
	query = `ALTER TABLE public.mib_to_module_compliance
    DROP CONSTRAINT "mib_to_module_compliance_mib_id_fkey",
    DROP CONSTRAINT "mib_to_module_compliance_module_compliance_id_fkey",
    DROP CONSTRAINT "mib_to_module_compliance_module_compliance_module_id_fkey",
    DROP CONSTRAINT "mib_to_module_compliance_module_compliance_module_group_id_fkey",
    DROP CONSTRAINT "mib_to_module_compliance_module_compliance_module_object_id_fkey",
    DROP CONSTRAINT "mib_mc";`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropMibToModuleIdentityDaoConstraints(ctx context.Context) error {
	var err error
	var query string
	query = `ALTER TABLE public.mib_to_module_identity
    DROP CONSTRAINT "mib_to_module_identity_mib_id_fkey",
    DROP CONSTRAINT "mib_to_module_identity_module_identity_id_fkey",
    DROP CONSTRAINT "mib_to_module_identity_revision_id_fkey",
    DROP CONSTRAINT "mib_mi_rev";`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropMibToNotificationGroupDaoConstraints(ctx context.Context) error {
	var err error
	var query string
	query = `ALTER TABLE public.mib_to_notification_group
    DROP CONSTRAINT "mib_to_notification_group_mib_id_fkey",
    DROP CONSTRAINT "mib_to_notification_group_notification_group_id_fkey",
    DROP CONSTRAINT "mib_ng";`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropMibToNotificationTypeDaoConstraints(ctx context.Context) error {
	var err error
	var query string
	query = `ALTER TABLE public.mib_to_notification_type
    DROP CONSTRAINT "mib_to_notification_type_mib_id_fkey",
    DROP CONSTRAINT "mib_to_notification_type_notification_type_id_fkey",
    DROP CONSTRAINT "mib_nt";`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropMibToObjectGroupDaoConstraints(ctx context.Context) error {
	var err error
	var query string
	query = `ALTER TABLE public.mib_to_object_group
    DROP CONSTRAINT "mib_to_object_group_mib_id_fkey",
    DROP CONSTRAINT "mib_to_object_group_object_group_id_fkey",
    DROP CONSTRAINT "mib_og";`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropMibToObjectIdentifierDaoConstraints(ctx context.Context) error {
	var err error
	var query string
	query = `ALTER TABLE public.mib_to_object_identifier
    DROP CONSTRAINT "mib_to_object_identifier_mib_id_fkey",
    DROP CONSTRAINT "mib_to_object_identifier_object_identifier_id_fkey",
    DROP CONSTRAINT "mib_oif";`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropMibToObjectIdentityDaoConstraints(ctx context.Context) error {
	var err error
	var query string
	query = `ALTER TABLE public.mib_to_object_identity
    DROP CONSTRAINT "mib_to_object_identity_mib_id_fkey",
    DROP CONSTRAINT "mib_to_object_identity_object_identity_id_fkey",
    DROP CONSTRAINT "mib_oit";`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropMibToObjectTypeDaoConstraints(ctx context.Context) error {
	var err error
	var query string
	query = `ALTER TABLE public.mib_to_object_type
    DROP CONSTRAINT "mib_to_object_type_mib_id_fkey",
    DROP CONSTRAINT "mib_to_object_type_object_type_id_fkey",
    DROP CONSTRAINT "mib_ot";`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropMibToSequenceDaoConstraints(ctx context.Context) error {
	var err error
	var query string
	query = `ALTER TABLE public.mib_to_sequence
    DROP CONSTRAINT "mib_to_sequence_mib_id_fkey",
    DROP CONSTRAINT "mib_to_sequence_sequence_id_fkey",
    DROP CONSTRAINT "mib_sec";`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropMibToTextualConventionDaoConstraints(ctx context.Context) error {
	var err error
	var query string
	query = `ALTER TABLE public.mib_to_textual_convention
    DROP CONSTRAINT "mib_to_textual_convention_mib_id_fkey",
    DROP CONSTRAINT "mib_to_textual_convention_textual_convention_id_fkey",
    DROP CONSTRAINT "mib_tc";`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropMibToTrapTypeDaoConstraints(ctx context.Context) error {
	var err error
	var query string
	query = `ALTER TABLE public.mib_to_trap_type
    DROP CONSTRAINT "mib_to_trap_type_mib_id_fkey",
    DROP CONSTRAINT "mib_to_trap_type_trap_type_id_fkey",
    DROP CONSTRAINT "mib_tt";`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropDeviceIndicatorDaoConstraints(ctx context.Context) error {
	var err error
	var query string
	query = `ALTER TABLE public.device_indicator
    DROP CONSTRAINT "device_indicator_tt";`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropParamIndicatorDaoConstraints(ctx context.Context) error {
	var err error
	var query string
	query = `ALTER TABLE public.param_indicator
    DROP CONSTRAINT "param_indicator_oid_id_fkey";`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropMappingDaoConstraints(ctx context.Context) error {
	var err error
	var query string
	query = `ALTER TABLE public.mapping
    DROP CONSTRAINT "mapping_indicator_fkey",
    DROP CONSTRAINT "mapping_param_fkey",
    DROP CONSTRAINT "mapping_frequency_fkey",
    DROP CONSTRAINT "mapping_from_fkey",
    DROP CONSTRAINT "mapping_position_type_fkey";`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropDeviceComponentDaoConstraints(ctx context.Context) error {
	var err error
	var query string
	query = `ALTER TABLE public.device_component
    DROP CONSTRAINT "device_component_model_fkey",
    DROP CONSTRAINT "device_component_parent_fkey";`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropDeviceComponentMappingDaoConstraints(ctx context.Context) error {
	var err error
	var query string
	query = `ALTER TABLE public.device_component_mapping
    DROP CONSTRAINT "device_component_mapping_device_component_id_fkey",
    DROP CONSTRAINT "device_component_mapping_mapping_id_fkey";`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropConfigurationDaoConstraints(ctx context.Context) error {
	var err error
	var query string
	query = `ALTER TABLE public.configuration
    DROP CONSTRAINT "configuration_indicator_fkey",
    DROP CONSTRAINT "configuration_device_component_id_fkey";`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropDefaultConfigurationDaoConstraints(ctx context.Context) error {
	var err error
	var query string
	query = `ALTER TABLE public.default_configuration
    DROP CONSTRAINT "default_configuration_indicator_fkey",
    DROP CONSTRAINT "default_configuration_device_component_id_fkey";`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropThresholdDaoConstraints(ctx context.Context) error {
	var err error
	var query string
	query = `ALTER TABLE public.threshold
    DROP CONSTRAINT "threshold_query_key";`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropDeviceSnmpDaoConstraints(ctx context.Context) error {
	var err error
	var query string
	query = `ALTER TABLE public.device_snmp
    DROP CONSTRAINT "device_snmp_version_fkey",
    DROP CONSTRAINT "device_snmp_authentication_fkey",
    DROP CONSTRAINT "device_snmp_privacy_fkey",
    DROP CONSTRAINT "device_snmp_config_fkey";`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropResultDaoConstraints(ctx context.Context) error {
	var err error
	var query string
	query = `ALTER TABLE public.result
    DROP CONSTRAINT "result_component_fkey",
    DROP CONSTRAINT "result_param_fkey";`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropAffectedThresholdDaoConstraints(ctx context.Context) error {
	var err error
	var query string
	query = `ALTER TABLE public.affected_threshold
    DROP CONSTRAINT "affected_threshold_threshold_fkey",
    DROP CONSTRAINT "affected_threshold_host_port_fkey";`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropAffectedParamDaoConstraints(ctx context.Context) error {
	var err error
	var query string
	query = `ALTER TABLE public.affected_param
    DROP CONSTRAINT "affected_param_component_fkey",
    DROP CONSTRAINT "affected_param_param_fkey",
    DROP CONSTRAINT "affected_param_host_port_fkey";`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropConfigInProcessDaoConstraints(ctx context.Context) error {
	var err error
	var query string
	query = `ALTER TABLE public.config_in_process
    DROP CONSTRAINT "config_in_process_protocol_fkey";`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func DropPollingProtocolDaoCounter(ctx context.Context) error {
	var err error
	var query string
	query = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.polling_protocol', 'id'), COALESCE(MAX("id"), 1)) FROM public.polling_protocol;`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		logger.Errorf("Failed to reset public.polling_protocol counter: %v", err)
		return err
	}
	return nil
}

func DropAccessDaoCounter(ctx context.Context) error {
	var err error
	var query string
	query = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.access', 'id'), COALESCE(MAX("id"), 1)) FROM public.access;`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		logger.Errorf("Failed to reset public.access counter: %v", err)
		return err
	}
	return nil
}

func DropVersionSnmpDaoCounter(ctx context.Context) error {
	var err error
	var query string
	query = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.version_snmp', 'id'), COALESCE(MAX("id"), 1)) FROM public.version_snmp;`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		logger.Errorf("Failed to reset public.version_snmp counter: %v", err)
		return err
	}
	return nil
}

func DropAuthProtocolSnmpDaoCounter(ctx context.Context) error {
	var err error
	var query string
	query = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.auth_protocol_snmp', 'id'), COALESCE(MAX("id"), 1)) FROM public.auth_protocol_snmp;`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		logger.Errorf("Failed to reset public.auth_protocol_snmp counter: %v", err)
		return err
	}
	return nil
}

func DropPrivacyProtocolSnmpDaoCounter(ctx context.Context) error {
	var err error
	var query string
	query = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.privacy_protocol_snmp', 'id'), COALESCE(MAX("id"), 1)) FROM public.privacy_protocol_snmp;`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		logger.Errorf("Failed to reset public.privacy_protocol_snmp counter: %v", err)
		return err
	}
	return nil
}

func DropOidTypeDaoCounter(ctx context.Context) error {
	var err error
	var query string
	query = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.oid_type', 'id'), COALESCE(MAX("id"), 1)) FROM public.oid_type;`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		logger.Errorf("Failed to reset public.oid_type counter: %v", err)
		return err
	}
	return nil
}

func DropLogicOperatorDaoCounter(ctx context.Context) error {
	var err error
	var query string
	query = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.logic_operator', 'id'), COALESCE(MAX("id"), 1)) FROM public.logic_operator;`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		logger.Errorf("Failed to reset public.logic_operator counter: %v", err)
		return err
	}
	return nil
}

func DropAlarmLevelDaoCounter(ctx context.Context) error {
	var err error
	var query string
	query = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.alarm_level', 'id'), COALESCE(MAX("id"), 1)) FROM public.alarm_level;`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		logger.Errorf("Failed to reset public.alarm_level counter: %v", err)
		return err
	}
	return nil
}

func DropVarTypeDaoCounter(ctx context.Context) error {
	var err error
	var query string
	query = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.var_type', 'id'), COALESCE(MAX("id"), 1)) FROM public.var_type;`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		logger.Errorf("Failed to reset public.var_type counter: %v", err)
		return err
	}
	return nil
}

func DropPollingFrequencyDaoCounter(ctx context.Context) error {
	var err error
	var query string
	query = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.polling_frequency', 'id'), COALESCE(MAX("id"), 1)) FROM public.polling_frequency;`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		logger.Errorf("Failed to reset public.polling_frequency counter: %v", err)
		return err
	}
	return nil
}

func DropOidAccessDaoCounter(ctx context.Context) error {
	var err error
	var query string
	query = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.oid_access', 'id'), COALESCE(MAX("id"), 1)) FROM public.oid_access;`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		logger.Errorf("Failed to reset public.oid_access counter: %v", err)
		return err
	}
	return nil
}

func DropOidStatusDaoCounter(ctx context.Context) error {
	var err error
	var query string
	query = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.oid_status', 'id'), COALESCE(MAX("id"), 1)) FROM public.oid_status;`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		logger.Errorf("Failed to reset public.oid_status counter: %v", err)
		return err
	}
	return nil
}

func DropAsn1TypeDaoCounter(ctx context.Context) error {
	var err error
	var query string
	query = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.asn1_type', 'id'), COALESCE(MAX("id"), 1)) FROM public.asn1_type;`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		logger.Errorf("Failed to reset public.asn1_type counter: %v", err)
		return err
	}
	return nil
}

func DropVendorDaoCounter(ctx context.Context) error {
	var err error
	var query string
	query = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.vendor', 'id'), COALESCE(MAX("id"), 1)) FROM public.vendor;`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		logger.Errorf("Failed to reset public.vendor counter: %v", err)
		return err
	}
	return nil
}

func DropComponentDaoCounter(ctx context.Context) error {
	var err error
	var query string
	query = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.component', 'id'), COALESCE(MAX("id"), 1)) FROM public.component;`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		logger.Errorf("Failed to reset public.component counter: %v", err)
		return err
	}
	return nil
}

func DropParamDaoCounter(ctx context.Context) error {
	var err error
	var query string
	query = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.param', 'id'), COALESCE(MAX("id"), 1)) FROM public.param;`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		logger.Errorf("Failed to reset public.param counter: %v", err)
		return err
	}
	return nil
}

func DropAgentCapabilitiesDaoCounter(ctx context.Context) error {
	var err error
	var query string
	query = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.agent_capabilities', 'id'), COALESCE(MAX("id"), 1)) FROM public.agent_capabilities;`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		logger.Errorf("Failed to reset public.agent_capabilities counter: %v", err)
		return err
	}
	return nil
}

func DropAgentCapabilitiesModuleDaoCounter(ctx context.Context) error {
	var err error
	var query string
	query = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.agent_capabilities_module', 'id'), COALESCE(MAX("id"), 1)) FROM public.agent_capabilities_module;`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		logger.Errorf("Failed to reset public.agent_capabilities_module counter: %v", err)
		return err
	}
	return nil
}

func DropAgentCapabilitiesModuleNotificationDaoCounter(ctx context.Context) error {
	var err error
	var query string
	query = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.agent_capabilities_module_notification', 'id'), COALESCE(MAX("id"), 1)) FROM public.agent_capabilities_module_notification;`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		logger.Errorf("Failed to reset public.agent_capabilities_module_notification counter: %v", err)
		return err
	}
	return nil
}

func DropAgentCapabilitiesModuleObjectDaoCounter(ctx context.Context) error {
	var err error
	var query string
	query = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.agent_capabilities_module_object', 'id'), COALESCE(MAX("id"), 1)) FROM public.agent_capabilities_module_object;`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		logger.Errorf("Failed to reset public.agent_capabilities_module_object counter: %v", err)
		return err
	}
	return nil
}

func DropChoiceDaoCounter(ctx context.Context) error {
	var err error
	var query string
	query = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.choice', 'id'), COALESCE(MAX("id"), 1)) FROM public.choice;`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		logger.Errorf("Failed to reset public.choice counter: %v", err)
		return err
	}
	return nil
}

func DropExplicitDaoCounter(ctx context.Context) error {
	var err error
	var query string
	query = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.explicit', 'id'), COALESCE(MAX("id"), 1)) FROM public.explicit;`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		logger.Errorf("Failed to reset public.explicit counter: %v", err)
		return err
	}
	return nil
}

func DropImplicitDaoCounter(ctx context.Context) error {
	var err error
	var query string
	query = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.implicit', 'id'), COALESCE(MAX("id"), 1)) FROM public.implicit;`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		logger.Errorf("Failed to reset public.implicit counter: %v", err)
		return err
	}
	return nil
}

func DropImportDaoCounter(ctx context.Context) error {
	var err error
	var query string
	query = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.import', 'id'), COALESCE(MAX("id"), 1)) FROM public.import;`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		logger.Errorf("Failed to reset public.import counter: %v", err)
		return err
	}
	return nil
}

func DropMibDaoCounter(ctx context.Context) error {
	var err error
	var query string
	query = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.mib', 'id'), COALESCE(MAX("id"), 1)) FROM public.mib;`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		logger.Errorf("Failed to reset public.mib counter: %v", err)
		return err
	}
	return nil
}

func DropModuleComplianceDaoCounter(ctx context.Context) error {
	var err error
	var query string
	query = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.module_compliance', 'id'), COALESCE(MAX("id"), 1)) FROM public.module_compliance;`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		logger.Errorf("Failed to reset public.module_compliance counter: %v", err)
		return err
	}
	return nil
}

func DropModuleComplianceModuleDaoCounter(ctx context.Context) error {
	var err error
	var query string
	query = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.module_compliance_module', 'id'), COALESCE(MAX("id"), 1)) FROM public.module_compliance_module;`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		logger.Errorf("Failed to reset public.module_compliance_module counter: %v", err)
		return err
	}
	return nil
}

func DropModuleComplianceModuleGroupDaoCounter(ctx context.Context) error {
	var err error
	var query string
	query = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.module_compliance_module_group', 'id'), COALESCE(MAX("id"), 1)) FROM public.module_compliance_module_group;`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		logger.Errorf("Failed to reset public.module_compliance_module_group counter: %v", err)
		return err
	}
	return nil
}

func DropModuleComplianceModuleObjectDaoCounter(ctx context.Context) error {
	var err error
	var query string
	query = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.module_compliance_module_object', 'id'), COALESCE(MAX("id"), 1)) FROM public.module_compliance_module_object;`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		logger.Errorf("Failed to reset public.module_compliance_module_object counter: %v", err)
		return err
	}
	return nil
}

func DropModuleIdentityDaoCounter(ctx context.Context) error {
	var err error
	var query string
	query = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.module_identity', 'id'), COALESCE(MAX("id"), 1)) FROM public.module_identity;`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		logger.Errorf("Failed to reset public.module_identity counter: %v", err)
		return err
	}
	return nil
}

func DropNotificationGroupDaoCounter(ctx context.Context) error {
	var err error
	var query string
	query = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.notification_group', 'id'), COALESCE(MAX("id"), 1)) FROM public.notification_group;`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		logger.Errorf("Failed to reset public.notification_group counter: %v", err)
		return err
	}
	return nil
}

func DropNotificationTypeDaoCounter(ctx context.Context) error {
	var err error
	var query string
	query = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.notification_type', 'id'), COALESCE(MAX("id"), 1)) FROM public.notification_type;`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		logger.Errorf("Failed to reset public.notification_type counter: %v", err)
		return err
	}
	return nil
}

func DropObjectGroupDaoCounter(ctx context.Context) error {
	var err error
	var query string
	query = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.object_group', 'id'), COALESCE(MAX("id"), 1)) FROM public.object_group;`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		logger.Errorf("Failed to reset public.object_group counter: %v", err)
		return err
	}
	return nil
}

func DropObjectIdentifierDaoCounter(ctx context.Context) error {
	var err error
	var query string
	query = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.object_identifier', 'id'), COALESCE(MAX("id"), 1)) FROM public.object_identifier;`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		logger.Errorf("Failed to reset public.object_identifier counter: %v", err)
		return err
	}
	return nil
}

func DropObjectIdentityDaoCounter(ctx context.Context) error {
	var err error
	var query string
	query = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.object_identity', 'id'), COALESCE(MAX("id"), 1)) FROM public.object_identity;`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		logger.Errorf("Failed to reset public.object_identity counter: %v", err)
		return err
	}
	return nil
}

func DropObjectTypeDaoCounter(ctx context.Context) error {
	var err error
	var query string
	query = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.object_type', 'id'), COALESCE(MAX("id"), 1)) FROM public.object_type;`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		logger.Errorf("Failed to reset public.object_type counter: %v", err)
		return err
	}
	return nil
}

func DropRevisionDaoCounter(ctx context.Context) error {
	var err error
	var query string
	query = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.revision', 'id'), COALESCE(MAX("id"), 1)) FROM public.revision;`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		logger.Errorf("Failed to reset public.revision counter: %v", err)
		return err
	}
	return nil
}

func DropSequenceDaoCounter(ctx context.Context) error {
	var err error
	var query string
	query = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.sequence', 'id'), COALESCE(MAX("id"), 1)) FROM public.sequence;`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		logger.Errorf("Failed to reset public.sequence counter: %v", err)
		return err
	}
	return nil
}

func DropTextualConventionDaoCounter(ctx context.Context) error {
	var err error
	var query string
	query = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.textual_convention', 'id'), COALESCE(MAX("id"), 1)) FROM public.textual_convention;`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		logger.Errorf("Failed to reset public.textual_convention counter: %v", err)
		return err
	}
	return nil
}

func DropTrapTypeDaoCounter(ctx context.Context) error {
	var err error
	var query string
	query = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.trap_type', 'id'), COALESCE(MAX("id"), 1)) FROM public.trap_type;`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		logger.Errorf("Failed to reset public.trap_type counter: %v", err)
		return err
	}
	return nil
}

func DropMibToAgentCapabilitiesDaoCounter(ctx context.Context) error {
	var err error
	var query string
	query = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.mib_to_agent_capabilities', 'id'), COALESCE(MAX("id"), 1)) FROM public.mib_to_agent_capabilities;`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		logger.Errorf("Failed to reset public.mib_to_agent_capabilities counter: %v", err)
		return err
	}
	return nil
}

func DropMibToChoiceDaoCounter(ctx context.Context) error {
	var err error
	var query string
	query = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.mib_to_choice', 'id'), COALESCE(MAX("id"), 1)) FROM public.mib_to_choice;`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		logger.Errorf("Failed to reset public.mib_to_choice counter: %v", err)
		return err
	}
	return nil
}

func DropMibToExplicitDaoCounter(ctx context.Context) error {
	var err error
	var query string
	query = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.mib_to_explicit', 'id'), COALESCE(MAX("id"), 1)) FROM public.mib_to_explicit;`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		logger.Errorf("Failed to reset public.mib_to_explicit counter: %v", err)
		return err
	}
	return nil
}

func DropMibToImplicitDaoCounter(ctx context.Context) error {
	var err error
	var query string
	query = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.mib_to_implicit', 'id'), COALESCE(MAX("id"), 1)) FROM public.mib_to_implicit;`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		logger.Errorf("Failed to reset public.mib_to_implicit counter: %v", err)
		return err
	}
	return nil
}

func DropMibToImportDaoCounter(ctx context.Context) error {
	var err error
	var query string
	query = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.mib_to_import', 'id'), COALESCE(MAX("id"), 1)) FROM public.mib_to_import;`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		logger.Errorf("Failed to reset public.mib_to_import counter: %v", err)
		return err
	}
	return nil
}

func DropMibToModuleComplianceDaoCounter(ctx context.Context) error {
	var err error
	var query string
	query = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.mib_to_module_compliance', 'id'), COALESCE(MAX("id"), 1)) FROM public.mib_to_module_compliance;`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		logger.Errorf("Failed to reset public.mib_to_module_compliance counter: %v", err)
		return err
	}
	return nil
}

func DropMibToModuleIdentityDaoCounter(ctx context.Context) error {
	var err error
	var query string
	query = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.mib_to_module_identity', 'id'), COALESCE(MAX("id"), 1)) FROM public.mib_to_module_identity;`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		logger.Errorf("Failed to reset public.mib_to_module_identity counter: %v", err)
		return err
	}
	return nil
}

func DropMibToNotificationGroupDaoCounter(ctx context.Context) error {
	var err error
	var query string
	query = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.mib_to_notification_group', 'id'), COALESCE(MAX("id"), 1)) FROM public.mib_to_notification_group;`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		logger.Errorf("Failed to reset public.mib_to_notification_group counter: %v", err)
		return err
	}
	return nil
}

func DropMibToNotificationTypeDaoCounter(ctx context.Context) error {
	var err error
	var query string
	query = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.mib_to_notification_type', 'id'), COALESCE(MAX("id"), 1)) FROM public.mib_to_notification_type;`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		logger.Errorf("Failed to reset public.mib_to_notification_type counter: %v", err)
		return err
	}
	return nil
}

func DropMibToObjectGroupDaoCounter(ctx context.Context) error {
	var err error
	var query string
	query = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.mib_to_object_group', 'id'), COALESCE(MAX("id"), 1)) FROM public.mib_to_object_group;`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		logger.Errorf("Failed to reset public.mib_to_object_group counter: %v", err)
		return err
	}
	return nil
}

func DropMibToObjectIdentifierDaoCounter(ctx context.Context) error {
	var err error
	var query string
	query = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.mib_to_object_identifier', 'id'), COALESCE(MAX("id"), 1)) FROM public.mib_to_object_identifier;`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		logger.Errorf("Failed to reset public.mib_to_object_identifier counter: %v", err)
		return err
	}
	return nil
}

func DropMibToObjectIdentityDaoCounter(ctx context.Context) error {
	var err error
	var query string
	query = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.mib_to_object_identity', 'id'), COALESCE(MAX("id"), 1)) FROM public.mib_to_object_identity;`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		logger.Errorf("Failed to reset public.mib_to_object_identity counter: %v", err)
		return err
	}
	return nil
}

func DropMibObjectTypeDaoCounter(ctx context.Context) error {
	var err error
	var query string
	query = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.mib_to_object_type', 'id'), COALESCE(MAX("id"), 1)) FROM public.mib_to_object_type;`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		logger.Errorf("Failed to reset public.mib_to_object_type counter: %v", err)
		return err
	}
	return nil
}

func DropMibToSequenceDaoCounter(ctx context.Context) error {
	var err error
	var query string
	query = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.mib_to_sequence', 'id'), COALESCE(MAX("id"), 1)) FROM public.mib_to_sequence;`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		logger.Errorf("Failed to reset public.mib_to_sequence counter: %v", err)
		return err
	}
	return nil
}

func DropMibToTextualConventionDaoCounter(ctx context.Context) error {
	var err error
	var query string
	query = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.mib_to_textual_convention', 'id'), COALESCE(MAX("id"), 1)) FROM public.mib_to_textual_convention;`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		logger.Errorf("Failed to reset public.mib_to_textual_convention counter: %v", err)
		return err
	}
	return nil
}

func DropMibToTrapTypeDaoCounter(ctx context.Context) error {
	var err error
	var query string
	query = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.mib_to_trap_type', 'id'), COALESCE(MAX("id"), 1)) FROM public.mib_to_trap_type;`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		logger.Errorf("Failed to reset public.mib_to_trap_type counter: %v", err)
		return err
	}
	return nil
}

func DropDeviceIndicatorDaoCounter(ctx context.Context) error {
	var err error
	var query string
	query = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.device_indicator', 'id'), COALESCE(MAX("id"), 1)) FROM public.device_indicator;`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		logger.Errorf("Failed to reset public.device_indicator counter: %v", err)
		return err
	}
	return nil
}

func DropParamIndicatorDaoCounter(ctx context.Context) error {
	var err error
	var query string
	query = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.param_indicator', 'id'), COALESCE(MAX("id"), 1)) FROM public.param_indicator;`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		logger.Errorf("Failed to reset public.param_indicator counter: %v", err)
		return err
	}
	return nil
}

func DropMappingDaoCounter(ctx context.Context) error {
	var err error
	var query string
	query = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.mapping', 'id'), COALESCE(MAX("id"), 1)) FROM public.mapping;`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		logger.Errorf("Failed to reset public.mapping counter: %v", err)
		return err
	}
	return nil
}

func DropDeviceComponentDaoCounter(ctx context.Context) error {
	var err error
	var query string
	query = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.device_component', 'id'), COALESCE(MAX("id"), 1)) FROM public.device_component;`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		logger.Errorf("Failed to reset public.device_component counter: %v", err)
		return err
	}
	return nil
}

func DropConfigurationDaoCounter(ctx context.Context) error {
	var err error
	var query string
	query = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.configuration', 'id'), COALESCE(MAX("id"), 1)) FROM public.configuration;`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		logger.Errorf("Failed to reset public.configuration counter: %v", err)
		return err
	}
	return nil
}

func DropDefaultConfigurationDaoCounter(ctx context.Context) error {
	var err error
	var query string
	query = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.default_configuration', 'id'), COALESCE(MAX("id"), 1)) FROM public.default_configuration;`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		logger.Errorf("Failed to reset public.default_configuration counter: %v", err)
		return err
	}
	return nil
}

func DropThresholdDaoCounter(ctx context.Context) error {
	var err error
	var query string
	query = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.threshold', 'id'), COALESCE(MAX("id"), 1)) FROM public.threshold;`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		logger.Errorf("Failed to reset public.threshold counter: %v", err)
		return err
	}
	return nil
}

func DropAffectedThresholdDaoCounter(ctx context.Context) error {
	var err error
	var query string
	query = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.affected_threshold', 'id'), COALESCE(MAX("id"), 1)) FROM public.affected_threshold;`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		logger.Errorf("Failed to reset public.affected_threshold counter: %v", err)
		return err
	}
	return nil
}

func DropAffectedParamDaoCounter(ctx context.Context) error {
	var err error
	var query string
	query = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.affected_param', 'id'), COALESCE(MAX("id"), 1)) FROM public.affected_param;`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		logger.Errorf("Failed to reset public.affected_param counter: %v", err)
		return err
	}
	return nil
}

func DropChangeLogDaoCounter(ctx context.Context) error {
	var err error
	var query string
	query = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.changelog', 'id'), COALESCE(MAX("id"), 1)) FROM public.changelog;`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		logger.Errorf("Failed to reset public.changelog counter: %v", err)
		return err
	}
	return nil
}
