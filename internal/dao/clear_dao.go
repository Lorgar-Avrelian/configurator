package dao

import (
	"configurator/internal/database"
	"context"
)

func ClearPollingProtocolDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.polling_protocol`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearAccessDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.access`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearVersionSnmpDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.version_snmp`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearAuthProtocolSnmpDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.auth_protocol_snmp`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearPrivacyProtocolSnmpDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.privacy_protocol_snmp`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearOidTypeDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.oid_type`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearLogicOperatorDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.logic_operator`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearAlarmLevelDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.alarm_level`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearVarTypeDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.var_type`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearPollingFrequencyDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.polling_frequency`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearOidAccessDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.oid_access`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearOidStatusDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.oid_status`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearAsn1TypeDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.asn1_type`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearVendorDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.vendor`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearComponentDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.component`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearParamDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.param`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearComponentParamDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.component_param`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearAgentCapabilitiesDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.agent_capabilities`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearAgentCapabilitiesModuleDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.agent_capabilities_module`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearAgentCapabilitiesModuleNotificationDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.agent_capabilities_module_notification`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearAgentCapabilitiesModuleObjectDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.agent_capabilities_module_object`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearChoiceDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.choice`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearExplicitDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.explicit`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearImplicitDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.implicit`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearImportDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.import`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearMibDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.mib`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearModuleComplianceDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.module_compliance`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearModuleComplianceModuleDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.module_compliance_module`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearModuleComplianceModuleGroupDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.module_compliance_module_group`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearModuleComplianceModuleObjectDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.module_compliance_module_object`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearModuleIdentityDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.module_identity`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearNotificationGroupDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.notification_group`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearNotificationTypeDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.notification_type`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearObjectGroupDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.object_group`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearObjectIdentifierDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.object_identifier`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearObjectIdentityDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.object_identity`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearObjectTypeDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.object_type`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearOidDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.oid`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearRevisionDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.revision`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearSequenceDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.sequence`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearTextualConventionDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.textual_convention`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearTrapTypeDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.trap_type`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearMibToAgentCapabilitiesDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.mib_to_agent_capabilities`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearMibToChoiceDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.mib_to_choice`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearMibToExplicitDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.mib_to_explicit`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearMibToImplicitDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.mib_to_implicit`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearMibToImportDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.mib_to_import`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearMibToModuleComplianceDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.mib_to_module_compliance`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearMibToModuleIdentityDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.mib_to_module_identity`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearMibToNotificationGroupDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.mib_to_notification_group`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearMibToNotificationTypeDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.mib_to_notification_type`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearMibToObjectGroupDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.mib_to_object_group`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearMibToObjectIdentifierDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.mib_to_object_identifier`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearMibToObjectIdentityDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.mib_to_object_identity`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearMibToObjectTypeDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.mib_to_object_type`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearMibToSequenceDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.mib_to_sequence`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearMibToTextualConventionDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.mib_to_textual_convention`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearMibToTrapTypeDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.mib_to_trap_type`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearDeviceIndicatorDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.device_indicator`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearParamIndicatorDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.param_indicator`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearMappingDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.mapping`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearDeviceComponentDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.device_component`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearDeviceComponentMappingDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.device_component_mapping`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearConfigurationDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.configuration`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearDefaultConfigurationDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.default_configuration`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearThresholdDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.threshold`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearDeviceSnmpDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.device_snmp`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearResultDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.result`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearAffectedThresholdDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.affected_threshold`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearAffectedParamDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.affected_param`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func ClearConfigInProcessDao(ctx context.Context) error {
	var err error
	var query string
	query = `TRUNCATE TABLE public.config_in_process`
	_, err = database.Get().Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}
