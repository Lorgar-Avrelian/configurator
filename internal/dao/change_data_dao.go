package dao

import (
	"configurator/internal/logger"
	"context"
	"fmt"
	"reflect"
	"sync"
)

var TableGroups = [][]string{
	{
		"public.polling_protocol",
		"public.access",
		"public.version_snmp",
		"public.auth_protocol_snmp",
		"public.privacy_protocol_snmp",
		"public.oid_type",
		"public.alarm_level",
		"public.var_type",
		"public.polling_frequency",
		"public.oid_access",
		"public.oid_status",
		"public.asn1_type",
		"public.logic_operator",
		"public.vendor",
	},
	{
		"public.mib",
		"public.component",
		"public.param",
		"public.threshold",
	},
	{
		"public.component_param",
		"public.agent_capabilities",
		"public.agent_capabilities_module",
		"public.agent_capabilities_module_notification",
		"public.agent_capabilities_module_object",
		"public.choice",
		"public.explicit",
		"public.implicit",
		"public.import",
		"public.module_compliance",
		"public.module_compliance_module",
		"public.module_compliance_module_group",
		"public.module_compliance_module_object",
		"public.module_identity",
		"public.notification_group",
		"public.notification_type",
		"public.object_group",
		"public.object_identifier",
		"public.object_identity",
		"public.object_type",
		"public.oid",
		"public.revision",
		"public.sequence",
		"public.textual_convention",
		"public.trap_type",
		"public.device_indicator",
	},
	{
		"public.mib_to_agent_capabilities",
		"public.mib_to_choice",
		"public.mib_to_explicit",
		"public.mib_to_implicit",
		"public.mib_to_import",
		"public.mib_to_module_compliance",
		"public.mib_to_module_identity",
		"public.mib_to_notification_group",
		"public.mib_to_notification_type",
		"public.mib_to_object_group",
		"public.mib_to_object_identifier",
		"public.mib_to_object_identity",
		"public.mib_to_object_type",
		"public.mib_to_sequence",
		"public.mib_to_textual_convention",
		"public.mib_to_trap_type",
		"public.param_indicator",
		"public.mapping",
		"public.device_component",
	},
	{
		"public.device_component_mapping",
		"public.configuration",
		"public.default_configuration",
	},
	{
		"public.device_snmp",
		"public.result",
		"public.config_in_process",
	},
	{
		"public.affected_threshold",
		"public.affected_param",
	},
}

func GetAllData(ctx context.Context) (
	[]PollingProtocolDao, []AccessDao, []VersionSnmpDao, []AuthProtocolSnmpDao, []PrivacyProtocolSnmpDao,
	[]OidTypeDao, []LogicOperatorDao, []AlarmLevelDao, []VarTypeDao, []PollingFrequencyDao,
	[]OidAccessDao, []OidStatusDao, []Asn1TypeDao, []VendorDao, []ComponentDao,
	[]ParamDao, []ComponentParamDao, []AgentCapabilitiesDao, []AgentCapabilitiesModuleDao, []AgentCapabilitiesModuleNotificationDao,
	[]AgentCapabilitiesModuleObjectDao, []ChoiceDao, []ExplicitDao, []ImplicitDao, []ImportDao,
	[]MibDao, []ModuleComplianceDao, []ModuleComplianceModuleDao, []ModuleComplianceModuleGroupDao, []ModuleComplianceModuleObjectDao,
	[]ModuleIdentityDao, []NotificationGroupDao, []NotificationTypeDao, []ObjectGroupDao, []ObjectIdentifierDao,
	[]ObjectIdentityDao, []ObjectTypeDao, []OidDao, []RevisionDao, []SequenceDao,
	[]TextualConventionDao, []TrapTypeDao, []MibToAgentCapabilitiesDao, []MibToChoiceDao, []MibToExplicitDao,
	[]MibToImplicitDao, []MibToImportDao, []MibToModuleComplianceDao, []MibToModuleIdentityDao, []MibToNotificationGroupDao,
	[]MibToNotificationTypeDao, []MibToObjectGroupDao, []MibToObjectIdentifierDao, []MibToObjectIdentityDao, []MibToObjectTypeDao,
	[]MibToSequenceDao, []MibToTextualConventionDao, []MibToTrapTypeDao, []DeviceIndicatorDao, []ParamIndicatorDao,
	[]MappingDao, []DeviceComponentDao, []DeviceComponentMappingDao, []ConfigurationDao, []DefaultConfigurationDao,
	[]ThresholdDao, []DeviceSnmpDao, []ResultDao, []AffectedThresholdDao, []AffectedParamDao,
	[]ConfigInProcessDao, error,
) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errResult error
	var callers map[string]any
	var results map[string]any
	var inArgs []reflect.Value
	var tName string
	var fn any
	var out1 []PollingProtocolDao
	var out2 []AccessDao
	var out3 []VersionSnmpDao
	var out4 []AuthProtocolSnmpDao
	var out5 []PrivacyProtocolSnmpDao
	var out6 []OidTypeDao
	var out7 []LogicOperatorDao
	var out8 []AlarmLevelDao
	var out9 []VarTypeDao
	var out10 []PollingFrequencyDao
	var out11 []OidAccessDao
	var out12 []OidStatusDao
	var out13 []Asn1TypeDao
	var out14 []VendorDao
	var out15 []ComponentDao
	var out16 []ParamDao
	var out17 []ComponentParamDao
	var out18 []AgentCapabilitiesDao
	var out19 []AgentCapabilitiesModuleDao
	var out20 []AgentCapabilitiesModuleNotificationDao
	var out21 []AgentCapabilitiesModuleObjectDao
	var out22 []ChoiceDao
	var out23 []ExplicitDao
	var out24 []ImplicitDao
	var out25 []ImportDao
	var out26 []MibDao
	var out27 []ModuleComplianceDao
	var out28 []ModuleComplianceModuleDao
	var out29 []ModuleComplianceModuleGroupDao
	var out30 []ModuleComplianceModuleObjectDao
	var out31 []ModuleIdentityDao
	var out32 []NotificationGroupDao
	var out33 []NotificationTypeDao
	var out34 []ObjectGroupDao
	var out35 []ObjectIdentifierDao
	var out36 []ObjectIdentityDao
	var out37 []ObjectTypeDao
	var out38 []OidDao
	var out39 []RevisionDao
	var out40 []SequenceDao
	var out41 []TextualConventionDao
	var out42 []TrapTypeDao
	var out43 []MibToAgentCapabilitiesDao
	var out44 []MibToChoiceDao
	var out45 []MibToExplicitDao
	var out46 []MibToImplicitDao
	var out47 []MibToImportDao
	var out48 []MibToModuleComplianceDao
	var out49 []MibToModuleIdentityDao
	var out50 []MibToNotificationGroupDao
	var out51 []MibToNotificationTypeDao
	var out52 []MibToObjectGroupDao
	var out53 []MibToObjectIdentifierDao
	var out54 []MibToObjectIdentityDao
	var out55 []MibToObjectTypeDao
	var out56 []MibToSequenceDao
	var out57 []MibToTextualConventionDao
	var out58 []MibToTrapTypeDao
	var out59 []DeviceIndicatorDao
	var out60 []ParamIndicatorDao
	var out61 []MappingDao
	var out62 []DeviceComponentDao
	var out63 []DeviceComponentMappingDao
	var out64 []ConfigurationDao
	var out65 []DefaultConfigurationDao
	var out66 []ThresholdDao
	var out67 []DeviceSnmpDao
	var out68 []ResultDao
	var out69 []AffectedThresholdDao
	var out70 []AffectedParamDao
	var out71 []ConfigInProcessDao
	inArgs = []reflect.Value{reflect.ValueOf(ctx)}
	results = make(map[string]any)
	callers = map[string]any{
		"public.polling_protocol":                       GetAllPollingProtocolDao,
		"public.access":                                 GetAllAccessDao,
		"public.version_snmp":                           GetAllVersionSnmpDao,
		"public.auth_protocol_snmp":                     GetAllAuthProtocolSnmpDao,
		"public.privacy_protocol_snmp":                  GetAllPrivacyProtocolSnmpDao,
		"public.oid_type":                               GetAllOidTypeDao,
		"public.logic_operator":                         GetAllLogicOperatorDao,
		"public.alarm_level":                            GetAllAlarmLevelDao,
		"public.var_type":                               GetAllVarTypeDao,
		"public.polling_frequency":                      GetAllPollingFrequencyDao,
		"public.oid_access":                             GetAllOidAccessDao,
		"public.oid_status":                             GetAllOidStatusDao,
		"public.asn1_type":                              GetAllAsn1TypeDao,
		"public.vendor":                                 GetAllVendorDao,
		"public.component":                              GetAllComponentDao,
		"public.param":                                  GetAllParamDao,
		"public.component_param":                        GetAllComponentParamDao,
		"public.agent_capabilities":                     GetAllAgentCapabilitiesDao,
		"public.agent_capabilities_module":              GetAllAgentCapabilitiesModuleDao,
		"public.agent_capabilities_module_notification": GetAllAgentCapabilitiesModuleNotificationDao,
		"public.agent_capabilities_module_object":       GetAllAgentCapabilitiesModuleObjectDao,
		"public.choice":                                 GetAllChoiceDao,
		"public.explicit":                               GetAllExplicitDao,
		"public.implicit":                               GetAllImplicitDao,
		"public.import":                                 GetAllImportDao,
		"public.mib":                                    GetAllMibDao,
		"public.module_compliance":                      GetAllModuleComplianceDao,
		"public.module_compliance_module":               GetAllModuleComplianceModuleDao,
		"public.module_compliance_module_group":         GetAllModuleComplianceModuleGroupDao,
		"public.module_compliance_module_object":        GetAllModuleComplianceModuleObjectDao,
		"public.module_identity":                        GetAllModuleIdentityDao,
		"public.notification_group":                     GetAllNotificationGroupDao,
		"public.notification_type":                      GetAllNotificationTypeDao,
		"public.object_group":                           GetAllObjectGroupDao,
		"public.object_identifier":                      GetAllObjectIdentifierDao,
		"public.object_identity":                        GetAllObjectIdentityDao,
		"public.object_type":                            GetAllObjectTypeDao,
		"public.oid":                                    GetAllOidDao,
		"public.revision":                               GetAllRevisionDao,
		"public.sequence":                               GetAllSequenceDao,
		"public.textual_convention":                     GetAllTextualConventionDao,
		"public.trap_type":                              GetAllTrapTypeDao,
		"public.mib_to_agent_capabilities":              GetAllMibToAgentCapabilitiesDao,
		"public.mib_to_choice":                          GetAllMibToChoiceDao,
		"public.mib_to_explicit":                        GetAllMibToExplicitDao,
		"public.mib_to_implicit":                        GetAllMibToImplicitDao,
		"public.mib_to_import":                          GetAllMibToImportDao,
		"public.mib_to_module_compliance":               GetAllMibToModuleComplianceDao,
		"public.mib_to_module_identity":                 GetAllMibToModuleIdentityDao,
		"public.mib_to_notification_group":              GetAllMibToNotificationGroupDao,
		"public.mib_to_notification_type":               GetAllMibToNotificationTypeDao,
		"public.mib_to_object_group":                    GetAllMibToObjectGroupDao,
		"public.mib_to_object_identifier":               GetAllMibToObjectIdentifierDao,
		"public.mib_to_object_identity":                 GetAllMibToObjectIdentityDao,
		"public.mib_to_object_type":                     GetAllMibToObjectTypeDao,
		"public.mib_to_sequence":                        GetAllMibToSequenceDao,
		"public.mib_to_textual_convention":              GetAllMibToTextualConventionDao,
		"public.mib_to_trap_type":                       GetAllMibToTrapTypeDao,
		"public.device_indicator":                       GetAllDeviceIndicatorDao,
		"public.param_indicator":                        GetAllParamIndicatorDao,
		"public.mapping":                                GetAllMappingDao,
		"public.device_component":                       GetAllDeviceComponentDao,
		"public.device_component_mapping":               GetAllDeviceComponentMappingDao,
		"public.configuration":                          GetAllConfigurationDao,
		"public.default_configuration":                  GetAllDefaultConfigurationDao,
		"public.threshold":                              GetAllThresholdDao,
		"public.device_snmp":                            GetAllDeviceSnmpDao,
		"public.result":                                 GetAllResultDao,
		"public.affected_threshold":                     GetAllAffectedThresholdDao,
		"public.affected_param":                         GetAllAffectedParamDao,
		"public.config_in_process":                      GetAllConfigInProcessDao,
	}
	for tName, fn = range callers {
		wg.Add(1)
		go func(tableName string, fetchFunc any) {
			var fVal reflect.Value
			var out []reflect.Value
			var dataSlice any
			var fetchErr any
			fVal = reflect.ValueOf(fetchFunc)
			out = fVal.Call(inArgs)
			dataSlice = out[0].Interface()
			fetchErr = out[1].Interface()
			wg.Done()
			if fetchErr != nil {
				mu.Lock()
				if errResult == nil {
					errResult = fmt.Errorf("error fetching %s: %v", tableName, fetchErr)
				}
				mu.Unlock()
				return
			}
			mu.Lock()
			results[tableName] = dataSlice
			mu.Unlock()
		}(tName, fn)
	}
	wg.Wait()
	if errResult != nil {
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, errResult
	}
	out1, _ = results["public.polling_protocol"].([]PollingProtocolDao)
	out2, _ = results["public.access"].([]AccessDao)
	out3, _ = results["public.version_snmp"].([]VersionSnmpDao)
	out4, _ = results["public.auth_protocol_snmp"].([]AuthProtocolSnmpDao)
	out5, _ = results["public.privacy_protocol_snmp"].([]PrivacyProtocolSnmpDao)
	out6, _ = results["public.oid_type"].([]OidTypeDao)
	out7, _ = results["public.logic_operator"].([]LogicOperatorDao)
	out8, _ = results["public.alarm_level"].([]AlarmLevelDao)
	out9, _ = results["public.var_type"].([]VarTypeDao)
	out10, _ = results["public.polling_frequency"].([]PollingFrequencyDao)
	out11, _ = results["public.oid_access"].([]OidAccessDao)
	out12, _ = results["public.oid_status"].([]OidStatusDao)
	out13, _ = results["public.asn1_type"].([]Asn1TypeDao)
	out14, _ = results["public.vendor"].([]VendorDao)
	out15, _ = results["public.component"].([]ComponentDao)
	out16, _ = results["public.param"].([]ParamDao)
	out17, _ = results["public.component_param"].([]ComponentParamDao)
	out18, _ = results["public.agent_capabilities"].([]AgentCapabilitiesDao)
	out19, _ = results["public.agent_capabilities_module"].([]AgentCapabilitiesModuleDao)
	out20, _ = results["public.agent_capabilities_module_notification"].([]AgentCapabilitiesModuleNotificationDao)
	out21, _ = results["public.agent_capabilities_module_object"].([]AgentCapabilitiesModuleObjectDao)
	out22, _ = results["public.choice"].([]ChoiceDao)
	out23, _ = results["public.explicit"].([]ExplicitDao)
	out24, _ = results["public.implicit"].([]ImplicitDao)
	out25, _ = results["public.import"].([]ImportDao)
	out26, _ = results["public.mib"].([]MibDao)
	out27, _ = results["public.module_compliance"].([]ModuleComplianceDao)
	out28, _ = results["public.module_compliance_module"].([]ModuleComplianceModuleDao)
	out29, _ = results["public.module_compliance_module_group"].([]ModuleComplianceModuleGroupDao)
	out30, _ = results["public.module_compliance_module_object"].([]ModuleComplianceModuleObjectDao)
	out31, _ = results["public.module_identity"].([]ModuleIdentityDao)
	out32, _ = results["public.notification_group"].([]NotificationGroupDao)
	out33, _ = results["public.notification_type"].([]NotificationTypeDao)
	out34, _ = results["public.object_group"].([]ObjectGroupDao)
	out35, _ = results["public.object_identifier"].([]ObjectIdentifierDao)
	out36, _ = results["public.object_identity"].([]ObjectIdentityDao)
	out37, _ = results["public.object_type"].([]ObjectTypeDao)
	out38, _ = results["public.oid"].([]OidDao)
	out39, _ = results["public.revision"].([]RevisionDao)
	out40, _ = results["public.sequence"].([]SequenceDao)
	out41, _ = results["public.textual_convention"].([]TextualConventionDao)
	out42, _ = results["public.trap_type"].([]TrapTypeDao)
	out43, _ = results["public.mib_to_agent_capabilities"].([]MibToAgentCapabilitiesDao)
	out44, _ = results["public.mib_to_choice"].([]MibToChoiceDao)
	out45, _ = results["public.mib_to_explicit"].([]MibToExplicitDao)
	out46, _ = results["public.mib_to_implicit"].([]MibToImplicitDao)
	out47, _ = results["public.mib_to_import"].([]MibToImportDao)
	out48, _ = results["public.mib_to_module_compliance"].([]MibToModuleComplianceDao)
	out49, _ = results["public.mib_to_module_identity"].([]MibToModuleIdentityDao)
	out50, _ = results["public.mib_to_notification_group"].([]MibToNotificationGroupDao)
	out51, _ = results["public.mib_to_notification_type"].([]MibToNotificationTypeDao)
	out52, _ = results["public.mib_to_object_group"].([]MibToObjectGroupDao)
	out53, _ = results["public.mib_to_object_identifier"].([]MibToObjectIdentifierDao)
	out54, _ = results["public.mib_to_object_identity"].([]MibToObjectIdentityDao)
	out55, _ = results["public.mib_to_object_type"].([]MibToObjectTypeDao)
	out56, _ = results["public.mib_to_sequence"].([]MibToSequenceDao)
	out57, _ = results["public.mib_to_textual_convention"].([]MibToTextualConventionDao)
	out58, _ = results["public.mib_to_trap_type"].([]MibToTrapTypeDao)
	out59, _ = results["public.device_indicator"].([]DeviceIndicatorDao)
	out60, _ = results["public.param_indicator"].([]ParamIndicatorDao)
	out61, _ = results["public.mapping"].([]MappingDao)
	out62, _ = results["public.device_component"].([]DeviceComponentDao)
	out63, _ = results["public.device_component_mapping"].([]DeviceComponentMappingDao)
	out64, _ = results["public.configuration"].([]ConfigurationDao)
	out65, _ = results["public.default_configuration"].([]DefaultConfigurationDao)
	out66, _ = results["public.threshold"].([]ThresholdDao)
	out67, _ = results["public.device_snmp"].([]DeviceSnmpDao)
	out68, _ = results["public.result"].([]ResultDao)
	out69, _ = results["public.affected_threshold"].([]AffectedThresholdDao)
	out70, _ = results["public.affected_param"].([]AffectedParamDao)
	out71, _ = results["public.config_in_process"].([]ConfigInProcessDao)
	return out1, out2, out3, out4, out5, out6, out7, out8, out9, out10, out11, out12, out13, out14, out15, out16, out17, out18, out19, out20, out21, out22, out23, out24, out25, out26, out27, out28, out29, out30, out31, out32, out33, out34, out35, out36, out37, out38, out39, out40, out41, out42, out43, out44, out45, out46, out47, out48, out49, out50, out51, out52, out53, out54, out55, out56, out57, out58, out59, out60, out61, out62, out63, out64, out65, out66, out67, out68, out69, out70, out71, nil
}

func ClearAllData(ctx context.Context) {
	var wg sync.WaitGroup
	var callers map[string]any
	var inArgs []reflect.Value
	var totalGroups int
	var groupIdx int
	var currentGroup []string
	var tName string
	var fn any
	inArgs = []reflect.Value{reflect.ValueOf(ctx)}
	callers = map[string]any{
		"public.polling_protocol":                       ClearPollingProtocolDao,
		"public.access":                                 ClearAccessDao,
		"public.version_snmp":                           ClearVersionSnmpDao,
		"public.auth_protocol_snmp":                     ClearAuthProtocolSnmpDao,
		"public.privacy_protocol_snmp":                  ClearPrivacyProtocolSnmpDao,
		"public.oid_type":                               ClearOidTypeDao,
		"public.logic_operator":                         ClearLogicOperatorDao,
		"public.alarm_level":                            ClearAlarmLevelDao,
		"public.var_type":                               ClearVarTypeDao,
		"public.polling_frequency":                      ClearPollingFrequencyDao,
		"public.oid_access":                             ClearOidAccessDao,
		"public.oid_status":                             ClearOidStatusDao,
		"public.asn1_type":                              ClearAsn1TypeDao,
		"public.vendor":                                 ClearVendorDao,
		"public.component":                              ClearComponentDao,
		"public.param":                                  ClearParamDao,
		"public.component_param":                        ClearComponentParamDao,
		"public.agent_capabilities":                     ClearAgentCapabilitiesDao,
		"public.agent_capabilities_module":              ClearAgentCapabilitiesModuleDao,
		"public.agent_capabilities_module_notification": ClearAgentCapabilitiesModuleNotificationDao,
		"public.agent_capabilities_module_object":       ClearAgentCapabilitiesModuleObjectDao,
		"public.choice":                                 ClearChoiceDao,
		"public.explicit":                               ClearExplicitDao,
		"public.implicit":                               ClearImplicitDao,
		"public.import":                                 ClearImportDao,
		"public.mib":                                    ClearMibDao,
		"public.module_compliance":                      ClearModuleComplianceDao,
		"public.module_compliance_module":               ClearModuleComplianceModuleDao,
		"public.module_compliance_module_group":         ClearModuleComplianceModuleGroupDao,
		"public.module_compliance_module_object":        ClearModuleComplianceModuleObjectDao,
		"public.module_identity":                        ClearModuleIdentityDao,
		"public.notification_group":                     ClearNotificationGroupDao,
		"public.notification_type":                      ClearNotificationTypeDao,
		"public.object_group":                           ClearObjectGroupDao,
		"public.object_identifier":                      ClearObjectIdentifierDao,
		"public.object_identity":                        ClearObjectIdentityDao,
		"public.object_type":                            ClearObjectTypeDao,
		"public.oid":                                    ClearOidDao,
		"public.revision":                               ClearRevisionDao,
		"public.sequence":                               ClearSequenceDao,
		"public.textual_convention":                     ClearTextualConventionDao,
		"public.trap_type":                              ClearTrapTypeDao,
		"public.mib_to_agent_capabilities":              ClearMibToAgentCapabilitiesDao,
		"public.mib_to_choice":                          ClearMibToChoiceDao,
		"public.mib_to_explicit":                        ClearMibToExplicitDao,
		"public.mib_to_implicit":                        ClearMibToImplicitDao,
		"public.mib_to_import":                          ClearMibToImportDao,
		"public.mib_to_module_compliance":               ClearMibToModuleComplianceDao,
		"public.mib_to_module_identity":                 ClearMibToModuleIdentityDao,
		"public.mib_to_notification_group":              ClearMibToNotificationGroupDao,
		"public.mib_to_notification_type":               ClearMibToNotificationTypeDao,
		"public.mib_to_object_group":                    ClearMibToObjectGroupDao,
		"public.mib_to_object_identifier":               ClearMibToObjectIdentifierDao,
		"public.mib_to_object_identity":                 ClearMibToObjectIdentityDao,
		"public.mib_to_object_type":                     ClearMibToObjectTypeDao,
		"public.mib_to_sequence":                        ClearMibToSequenceDao,
		"public.mib_to_textual_convention":              ClearMibToTextualConventionDao,
		"public.mib_to_trap_type":                       ClearMibToTrapTypeDao,
		"public.device_indicator":                       ClearDeviceIndicatorDao,
		"public.param_indicator":                        ClearParamIndicatorDao,
		"public.mapping":                                ClearMappingDao,
		"public.device_component":                       ClearDeviceComponentDao,
		"public.device_component_mapping":               ClearDeviceComponentMappingDao,
		"public.configuration":                          ClearConfigurationDao,
		"public.default_configuration":                  ClearDefaultConfigurationDao,
		"public.threshold":                              ClearThresholdDao,
		"public.device_snmp":                            ClearDeviceSnmpDao,
		"public.result":                                 ClearResultDao,
		"public.affected_threshold":                     ClearAffectedThresholdDao,
		"public.affected_param":                         ClearAffectedParamDao,
		"public.config_in_process":                      ClearConfigInProcessDao,
	}
	totalGroups = len(TableGroups)
	go func() {
		groupIdx = totalGroups
		for groupIdx > 0 {
			groupIdx--
			currentGroup = TableGroups[groupIdx]
			var tableIdx int
			var totalTables int
			totalTables = len(currentGroup)
			tableIdx = totalTables
			for tableIdx > 0 {
				tableIdx--
				tName = currentGroup[totalTables-1-tableIdx]
				fn = callers[tName]
				wg.Add(1)
				go func(tableName string, clearFunc any) {
					var fVal reflect.Value
					var out []reflect.Value
					var clearErr any
					fVal = reflect.ValueOf(clearFunc)
					out = fVal.Call(inArgs)
					if len(out) > 0 {
						clearErr = out[0].Interface()
					}
					wg.Done()
					if clearErr != nil {
						logger.Error("Error clearing %s: %v", tableName, clearErr)
					}
				}(tName, fn)
			}
			wg.Wait()
		}
	}()
}

func InsertAllData(
	ctx context.Context,
	in1 []PollingProtocolDao, in2 []AccessDao, in3 []VersionSnmpDao, in4 []AuthProtocolSnmpDao, in5 []PrivacyProtocolSnmpDao,
	in6 []OidTypeDao, in7 []LogicOperatorDao, in8 []AlarmLevelDao, in9 []VarTypeDao, in10 []PollingFrequencyDao,
	in11 []OidAccessDao, in12 []OidStatusDao, in13 []Asn1TypeDao, in14 []VendorDao, in15 []ComponentDao,
	in16 []ParamDao, in17 []ComponentParamDao, in18 []AgentCapabilitiesDao, in19 []AgentCapabilitiesModuleDao, in20 []AgentCapabilitiesModuleNotificationDao,
	in21 []AgentCapabilitiesModuleObjectDao, in22 []ChoiceDao, in23 []ExplicitDao, in24 []ImplicitDao, in25 []ImportDao,
	in26 []MibDao, in27 []ModuleComplianceDao, in28 []ModuleComplianceModuleDao, in29 []ModuleComplianceModuleGroupDao, in30 []ModuleComplianceModuleObjectDao,
	in31 []ModuleIdentityDao, in32 []NotificationGroupDao, in33 []NotificationTypeDao, in34 []ObjectGroupDao, in35 []ObjectIdentifierDao,
	in36 []ObjectIdentityDao, in37 []ObjectTypeDao, in38 []OidDao, in39 []RevisionDao, in40 []SequenceDao,
	in41 []TextualConventionDao, in42 []TrapTypeDao, in43 []MibToAgentCapabilitiesDao, in44 []MibToChoiceDao, in45 []MibToExplicitDao,
	in46 []MibToImplicitDao, in47 []MibToImportDao, in48 []MibToModuleComplianceDao, in49 []MibToModuleIdentityDao, in50 []MibToNotificationGroupDao,
	in51 []MibToNotificationTypeDao, in52 []MibToObjectGroupDao, in53 []MibToObjectIdentifierDao, in54 []MibToObjectIdentityDao, in55 []MibToObjectTypeDao,
	in56 []MibToSequenceDao, in57 []MibToTextualConventionDao, in58 []MibToTrapTypeDao, in59 []DeviceIndicatorDao, in60 []ParamIndicatorDao,
	in61 []MappingDao, in62 []DeviceComponentDao, in63 []DeviceComponentMappingDao, in64 []ConfigurationDao, in65 []DefaultConfigurationDao,
	in66 []ThresholdDao, in67 []DeviceSnmpDao, in68 []ResultDao, in69 []AffectedThresholdDao, in70 []AffectedParamDao,
	in71 []ConfigInProcessDao,
) error {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var callers map[string]any
	var args map[string][]reflect.Value
	var errResult error
	var totalGroups int
	var groupIdx int
	var currentGroup []string
	var tName string
	var fn any
	callers = map[string]any{
		"public.polling_protocol":                       ImportPollingProtocolDao,
		"public.access":                                 ImportAccessDao,
		"public.version_snmp":                           ImportVersionSnmpDao,
		"public.auth_protocol_snmp":                     ImportAuthProtocolSnmpDao,
		"public.privacy_protocol_snmp":                  ImportPrivacyProtocolSnmpDao,
		"public.oid_type":                               ImportOidTypeDao,
		"public.logic_operator":                         ImportLogicOperatorDao,
		"public.alarm_level":                            ImportAlarmLevelDao,
		"public.var_type":                               ImportVarTypeDao,
		"public.polling_frequency":                      ImportPollingFrequencyDao,
		"public.oid_access":                             ImportOidAccessDao,
		"public.oid_status":                             ImportOidStatusDao,
		"public.asn1_type":                              ImportAsn1TypeDao,
		"public.vendor":                                 ImportVendorDao,
		"public.component":                              ImportComponentDao,
		"public.param":                                  ImportParamDao,
		"public.component_param":                        ImportComponentParamDao,
		"public.agent_capabilities":                     ImportAgentCapabilitiesDao,
		"public.agent_capabilities_module":              ImportAgentCapabilitiesModuleDao,
		"public.agent_capabilities_module_notification": ImportAgentCapabilitiesModuleNotificationDao,
		"public.agent_capabilities_module_object":       ImportAgentCapabilitiesModuleObjectDao,
		"public.choice":                                 ImportChoiceDao,
		"public.explicit":                               ImportExplicitDao,
		"public.implicit":                               ImportImplicitDao,
		"public.import":                                 ImportImportDao,
		"public.mib":                                    ImportMibDao,
		"public.module_compliance":                      ImportModuleComplianceDao,
		"public.module_compliance_module":               ImportModuleComplianceModuleDao,
		"public.module_compliance_module_group":         ImportModuleComplianceModuleGroupDao,
		"public.module_compliance_module_object":        ImportModuleComplianceModuleObjectDao,
		"public.module_identity":                        ImportModuleIdentityDao,
		"public.notification_group":                     ImportNotificationGroupDao,
		"public.notification_type":                      ImportNotificationTypeDao,
		"public.object_group":                           ImportObjectGroupDao,
		"public.object_identifier":                      ImportObjectIdentifierDao,
		"public.object_identity":                        ImportObjectIdentityDao,
		"public.object_type":                            ImportObjectTypeDao,
		"public.oid":                                    ImportOidDao,
		"public.revision":                               ImportRevisionDao,
		"public.sequence":                               ImportSequenceDao,
		"public.textual_convention":                     ImportTextualConventionDao,
		"public.trap_type":                              ImportTrapTypeDao,
		"public.mib_to_agent_capabilities":              ImportMibToAgentCapabilitiesDao,
		"public.mib_to_choice":                          ImportMibToChoiceDao,
		"public.mib_to_explicit":                        ImportMibToExplicitDao,
		"public.mib_to_implicit":                        ImportMibToImplicitDao,
		"public.mib_to_import":                          ImportMibToImportDao,
		"public.mib_to_module_compliance":               ImportMibToModuleComplianceDao,
		"public.mib_to_module_identity":                 ImportMibToModuleIdentityDao,
		"public.mib_to_notification_group":              ImportMibToNotificationGroupDao,
		"public.mib_to_notification_type":               ImportMibToNotificationTypeDao,
		"public.mib_to_object_group":                    ImportMibToObjectGroupDao,
		"public.mib_to_object_identifier":               ImportMibToObjectIdentifierDao,
		"public.mib_to_object_identity":                 ImportMibToObjectIdentityDao,
		"public.mib_to_object_type":                     ImportMibToObjectTypeDao,
		"public.mib_to_sequence":                        ImportMibToSequenceDao,
		"public.mib_to_textual_convention":              ImportMibToTextualConventionDao,
		"public.mib_to_trap_type":                       ImportMibToTrapTypeDao,
		"public.device_indicator":                       ImportDeviceIndicatorDao,
		"public.param_indicator":                        ImportParamIndicatorDao,
		"public.mapping":                                ImportMappingDao,
		"public.device_component":                       ImportDeviceComponentDao,
		"public.device_component_mapping":               ImportDeviceComponentMappingDao,
		"public.configuration":                          ImportConfigurationDao,
		"public.default_configuration":                  ImportDefaultConfigurationDao,
		"public.threshold":                              ImportThresholdDao,
		"public.device_snmp":                            ImportDeviceSnmpDao,
		"public.result":                                 ImportResultDao,
		"public.affected_threshold":                     ImportAffectedThresholdDao,
		"public.affected_param":                         ImportAffectedParamDao,
		"public.config_in_process":                      ImportConfigInProcessDao,
	}
	args = map[string][]reflect.Value{
		"public.polling_protocol":                       {reflect.ValueOf(ctx), reflect.ValueOf(in1)},
		"public.access":                                 {reflect.ValueOf(ctx), reflect.ValueOf(in2)},
		"public.version_snmp":                           {reflect.ValueOf(ctx), reflect.ValueOf(in3)},
		"public.auth_protocol_snmp":                     {reflect.ValueOf(ctx), reflect.ValueOf(in4)},
		"public.privacy_protocol_snmp":                  {reflect.ValueOf(ctx), reflect.ValueOf(in5)},
		"public.oid_type":                               {reflect.ValueOf(ctx), reflect.ValueOf(in6)},
		"public.logic_operator":                         {reflect.ValueOf(ctx), reflect.ValueOf(in7)},
		"public.alarm_level":                            {reflect.ValueOf(ctx), reflect.ValueOf(in8)},
		"public.var_type":                               {reflect.ValueOf(ctx), reflect.ValueOf(in9)},
		"public.polling_frequency":                      {reflect.ValueOf(ctx), reflect.ValueOf(in10)},
		"public.oid_access":                             {reflect.ValueOf(ctx), reflect.ValueOf(in11)},
		"public.oid_status":                             {reflect.ValueOf(ctx), reflect.ValueOf(in12)},
		"public.asn1_type":                              {reflect.ValueOf(ctx), reflect.ValueOf(in13)},
		"public.vendor":                                 {reflect.ValueOf(ctx), reflect.ValueOf(in14)},
		"public.component":                              {reflect.ValueOf(ctx), reflect.ValueOf(in15)},
		"public.param":                                  {reflect.ValueOf(ctx), reflect.ValueOf(in16)},
		"public.component_param":                        {reflect.ValueOf(ctx), reflect.ValueOf(in17)},
		"public.agent_capabilities":                     {reflect.ValueOf(ctx), reflect.ValueOf(in18)},
		"public.agent_capabilities_module":              {reflect.ValueOf(ctx), reflect.ValueOf(in19)},
		"public.agent_capabilities_module_notification": {reflect.ValueOf(ctx), reflect.ValueOf(in20)},
		"public.agent_capabilities_module_object":       {reflect.ValueOf(ctx), reflect.ValueOf(in21)},
		"public.choice":                                 {reflect.ValueOf(ctx), reflect.ValueOf(in22)},
		"public.explicit":                               {reflect.ValueOf(ctx), reflect.ValueOf(in23)},
		"public.implicit":                               {reflect.ValueOf(ctx), reflect.ValueOf(in24)},
		"public.import":                                 {reflect.ValueOf(ctx), reflect.ValueOf(in25)},
		"public.mib":                                    {reflect.ValueOf(ctx), reflect.ValueOf(in26)},
		"public.module_compliance":                      {reflect.ValueOf(ctx), reflect.ValueOf(in27)},
		"public.module_compliance_module":               {reflect.ValueOf(ctx), reflect.ValueOf(in28)},
		"public.module_compliance_module_group":         {reflect.ValueOf(ctx), reflect.ValueOf(in29)},
		"public.module_compliance_module_object":        {reflect.ValueOf(ctx), reflect.ValueOf(in30)},
		"public.module_identity":                        {reflect.ValueOf(ctx), reflect.ValueOf(in31)},
		"public.notification_group":                     {reflect.ValueOf(ctx), reflect.ValueOf(in32)},
		"public.notification_type":                      {reflect.ValueOf(ctx), reflect.ValueOf(in33)},
		"public.object_group":                           {reflect.ValueOf(ctx), reflect.ValueOf(in34)},
		"public.object_identifier":                      {reflect.ValueOf(ctx), reflect.ValueOf(in35)},
		"public.object_identity":                        {reflect.ValueOf(ctx), reflect.ValueOf(in36)},
		"public.object_type":                            {reflect.ValueOf(ctx), reflect.ValueOf(in37)},
		"public.oid":                                    {reflect.ValueOf(ctx), reflect.ValueOf(in38)},
		"public.revision":                               {reflect.ValueOf(ctx), reflect.ValueOf(in39)},
		"public.sequence":                               {reflect.ValueOf(ctx), reflect.ValueOf(in40)},
		"public.textual_convention":                     {reflect.ValueOf(ctx), reflect.ValueOf(in41)},
		"public.trap_type":                              {reflect.ValueOf(ctx), reflect.ValueOf(in42)},
		"public.mib_to_agent_capabilities":              {reflect.ValueOf(ctx), reflect.ValueOf(in43)},
		"public.mib_to_choice":                          {reflect.ValueOf(ctx), reflect.ValueOf(in44)},
		"public.mib_to_explicit":                        {reflect.ValueOf(ctx), reflect.ValueOf(in45)},
		"public.mib_to_implicit":                        {reflect.ValueOf(ctx), reflect.ValueOf(in46)},
		"public.mib_to_import":                          {reflect.ValueOf(ctx), reflect.ValueOf(in47)},
		"public.mib_to_module_compliance":               {reflect.ValueOf(ctx), reflect.ValueOf(in48)},
		"public.mib_to_module_identity":                 {reflect.ValueOf(ctx), reflect.ValueOf(in49)},
		"public.mib_to_notification_group":              {reflect.ValueOf(ctx), reflect.ValueOf(in50)},
		"public.mib_to_notification_type":               {reflect.ValueOf(ctx), reflect.ValueOf(in51)},
		"public.mib_to_object_group":                    {reflect.ValueOf(ctx), reflect.ValueOf(in52)},
		"public.mib_to_object_identifier":               {reflect.ValueOf(ctx), reflect.ValueOf(in53)},
		"public.mib_to_object_identity":                 {reflect.ValueOf(ctx), reflect.ValueOf(in54)},
		"public.mib_to_object_type":                     {reflect.ValueOf(ctx), reflect.ValueOf(in55)},
		"public.mib_to_sequence":                        {reflect.ValueOf(ctx), reflect.ValueOf(in56)},
		"public.mib_to_textual_convention":              {reflect.ValueOf(ctx), reflect.ValueOf(in57)},
		"public.mib_to_trap_type":                       {reflect.ValueOf(ctx), reflect.ValueOf(in58)},
		"public.device_indicator":                       {reflect.ValueOf(ctx), reflect.ValueOf(in59)},
		"public.param_indicator":                        {reflect.ValueOf(ctx), reflect.ValueOf(in60)},
		"public.mapping":                                {reflect.ValueOf(ctx), reflect.ValueOf(in61)},
		"public.device_component":                       {reflect.ValueOf(ctx), reflect.ValueOf(in62)},
		"public.device_component_mapping":               {reflect.ValueOf(ctx), reflect.ValueOf(in63)},
		"public.configuration":                          {reflect.ValueOf(ctx), reflect.ValueOf(in64)},
		"public.default_configuration":                  {reflect.ValueOf(ctx), reflect.ValueOf(in65)},
		"public.threshold":                              {reflect.ValueOf(ctx), reflect.ValueOf(in66)},
		"public.device_snmp":                            {reflect.ValueOf(ctx), reflect.ValueOf(in67)},
		"public.result":                                 {reflect.ValueOf(ctx), reflect.ValueOf(in68)},
		"public.affected_threshold":                     {reflect.ValueOf(ctx), reflect.ValueOf(in69)},
		"public.affected_param":                         {reflect.ValueOf(ctx), reflect.ValueOf(in70)},
		"public.config_in_process":                      {reflect.ValueOf(ctx), reflect.ValueOf(in71)},
	}
	totalGroups = len(TableGroups)
	groupIdx = 0
	for totalGroups > groupIdx {
		currentGroup = TableGroups[groupIdx]
		var tableIdx int
		var totalTables int
		totalTables = len(currentGroup)
		tableIdx = totalTables
		for tableIdx > 0 {
			tableIdx--
			tName = currentGroup[totalTables-1-tableIdx]
			fn = callers[tName]
			wg.Add(1)
			go func(tableName string, insertFunc any) {
				var fVal reflect.Value
				var out []reflect.Value
				var insertErr any
				fVal = reflect.ValueOf(insertFunc)
				out = fVal.Call(args[tableName])
				if len(out) > 0 {
					insertErr = out[0].Interface()
				}
				wg.Done()
				if insertErr != nil {
					mu.Lock()
					if errResult == nil {
						errResult = fmt.Errorf("error inserting %s: %v", tableName, insertErr)
					}
					mu.Unlock()
				}
			}(tName, fn)
		}
		wg.Wait()
		if errResult != nil {
			return errResult
		}
		groupIdx++
	}
	return nil
}
