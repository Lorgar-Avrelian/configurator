package service

import (
	"configurator/internal/dao"
	"context"
)

func ChangeComponentData(ctx context.Context, prevId int64, newId int64) (bool, error) {
	var err error
	var ok bool
	var d1 []dao.PollingProtocolDao
	var d2 []dao.AccessDao
	var d3 []dao.VersionSnmpDao
	var d4 []dao.AuthProtocolSnmpDao
	var d5 []dao.PrivacyProtocolSnmpDao
	var d6 []dao.OidTypeDao
	var d7 []dao.LogicOperatorDao
	var d8 []dao.AlarmLevelDao
	var d9 []dao.VarTypeDao
	var d10 []dao.PollingFrequencyDao
	var d11 []dao.OidAccessDao
	var d12 []dao.OidStatusDao
	var d13 []dao.Asn1TypeDao
	var d14 []dao.VendorDao
	var d15 []dao.ComponentDao
	var d16 []dao.ParamDao
	var d17 []dao.ComponentParamDao
	var d18 []dao.AgentCapabilitiesDao
	var d19 []dao.AgentCapabilitiesModuleDao
	var d20 []dao.AgentCapabilitiesModuleNotificationDao
	var d21 []dao.AgentCapabilitiesModuleObjectDao
	var d22 []dao.ChoiceDao
	var d23 []dao.ExplicitDao
	var d24 []dao.ImplicitDao
	var d25 []dao.ImportDao
	var d26 []dao.MibDao
	var d27 []dao.ModuleComplianceDao
	var d28 []dao.ModuleComplianceModuleDao
	var d29 []dao.ModuleComplianceModuleGroupDao
	var d30 []dao.ModuleComplianceModuleObjectDao
	var d31 []dao.ModuleIdentityDao
	var d32 []dao.NotificationGroupDao
	var d33 []dao.NotificationTypeDao
	var d34 []dao.ObjectGroupDao
	var d35 []dao.ObjectIdentifierDao
	var d36 []dao.ObjectIdentityDao
	var d37 []dao.ObjectTypeDao
	var d38 []dao.OidDao
	var d39 []dao.RevisionDao
	var d40 []dao.SequenceDao
	var d41 []dao.TextualConventionDao
	var d42 []dao.TrapTypeDao
	var d43 []dao.MibToAgentCapabilitiesDao
	var d44 []dao.MibToChoiceDao
	var d45 []dao.MibToExplicitDao
	var d46 []dao.MibToImplicitDao
	var d47 []dao.MibToImportDao
	var d48 []dao.MibToModuleComplianceDao
	var d49 []dao.MibToModuleIdentityDao
	var d50 []dao.MibToNotificationGroupDao
	var d51 []dao.MibToNotificationTypeDao
	var d52 []dao.MibToObjectGroupDao
	var d53 []dao.MibToObjectIdentifierDao
	var d54 []dao.MibToObjectIdentityDao
	var d55 []dao.MibToObjectTypeDao
	var d56 []dao.MibToSequenceDao
	var d57 []dao.MibToTextualConventionDao
	var d58 []dao.MibToTrapTypeDao
	var d59 []dao.DeviceIndicatorDao
	var d60 []dao.ParamIndicatorDao
	var d61 []dao.MappingDao
	var d62 []dao.DeviceComponentDao
	var d63 []dao.DeviceComponentMappingDao
	var d64 []dao.ConfigurationDao
	var d65 []dao.DefaultConfigurationDao
	var d66 []dao.ThresholdDao
	var d67 []dao.DeviceSnmpDao
	var d68 []dao.ResultDao
	var d69 []dao.AffectedThresholdDao
	var d70 []dao.AffectedParamDao
	var d71 []dao.ConfigInProcessDao
	var d72 []dao.ChangeLogDao
	d1, d2, d3, d4, d5, d6, d7, d8, d9, d10, d11, d12, d13, d14, d15, d16, d17, d18, d19, d20, d21, d22, d23, d24, d25, d26, d27, d28, d29, d30, d31, d32, d33, d34, d35, d36, d37, d38, d39, d40, d41, d42, d43, d44, d45, d46, d47, d48, d49, d50, d51, d52, d53, d54, d55, d56, d57, d58, d59, d60, d61, d62, d63, d64, d65, d66, d67, d68, d69, d70, d71, d72, err = dao.GetAllData(ctx)
	if err != nil {
		ok = false
		return ok, err
	}
	dao.ClearAllData(ctx)
	err = dao.InsertAllData(ctx, d1, d2, d3, d4, d5, d6, d7, d8, d9, d10, d11, d12, d13, d14, d15, d16, d17, d18, d19, d20, d21, d22, d23, d24, d25, d26, d27, d28, d29, d30, d31, d32, d33, d34, d35, d36, d37, d38, d39, d40, d41, d42, d43, d44, d45, d46, d47, d48, d49, d50, d51, d52, d53, d54, d55, d56, d57, d58, d59, d60, d61, d62, d63, d64, d65, d66, d67, d68, d69, d70, d71, d72)
	if err != nil {
		ok = false
		return ok, err
	}
	ok = true
	return ok, nil
}
