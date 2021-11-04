//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armalertsmanagement

const (
	module  = "armalertsmanagement"
	version = "v0.2.1"
)

// ActionRuleStatus - Indicates if the given action rule is enabled or disabled
type ActionRuleStatus string

const (
	ActionRuleStatusDisabled ActionRuleStatus = "Disabled"
	ActionRuleStatusEnabled  ActionRuleStatus = "Enabled"
)

// PossibleActionRuleStatusValues returns the possible values for the ActionRuleStatus const type.
func PossibleActionRuleStatusValues() []ActionRuleStatus {
	return []ActionRuleStatus{
		ActionRuleStatusDisabled,
		ActionRuleStatusEnabled,
	}
}

// ToPtr returns a *ActionRuleStatus pointing to the current value.
func (c ActionRuleStatus) ToPtr() *ActionRuleStatus {
	return &c
}

// ActionRuleType - Indicates type of action rule
type ActionRuleType string

const (
	ActionRuleTypeActionGroup ActionRuleType = "ActionGroup"
	ActionRuleTypeDiagnostics ActionRuleType = "Diagnostics"
	ActionRuleTypeSuppression ActionRuleType = "Suppression"
)

// PossibleActionRuleTypeValues returns the possible values for the ActionRuleType const type.
func PossibleActionRuleTypeValues() []ActionRuleType {
	return []ActionRuleType{
		ActionRuleTypeActionGroup,
		ActionRuleTypeDiagnostics,
		ActionRuleTypeSuppression,
	}
}

// ToPtr returns a *ActionRuleType pointing to the current value.
func (c ActionRuleType) ToPtr() *ActionRuleType {
	return &c
}

// AlertModificationEvent - Reason for the modification
type AlertModificationEvent string

const (
	AlertModificationEventAlertCreated           AlertModificationEvent = "AlertCreated"
	AlertModificationEventStateChange            AlertModificationEvent = "StateChange"
	AlertModificationEventMonitorConditionChange AlertModificationEvent = "MonitorConditionChange"
	AlertModificationEventSeverityChange         AlertModificationEvent = "SeverityChange"
	AlertModificationEventActionRuleTriggered    AlertModificationEvent = "ActionRuleTriggered"
	AlertModificationEventActionRuleSuppressed   AlertModificationEvent = "ActionRuleSuppressed"
	AlertModificationEventActionsTriggered       AlertModificationEvent = "ActionsTriggered"
	AlertModificationEventActionsSuppressed      AlertModificationEvent = "ActionsSuppressed"
	AlertModificationEventActionsFailed          AlertModificationEvent = "ActionsFailed"
)

// PossibleAlertModificationEventValues returns the possible values for the AlertModificationEvent const type.
func PossibleAlertModificationEventValues() []AlertModificationEvent {
	return []AlertModificationEvent{
		AlertModificationEventAlertCreated,
		AlertModificationEventStateChange,
		AlertModificationEventMonitorConditionChange,
		AlertModificationEventSeverityChange,
		AlertModificationEventActionRuleTriggered,
		AlertModificationEventActionRuleSuppressed,
		AlertModificationEventActionsTriggered,
		AlertModificationEventActionsSuppressed,
		AlertModificationEventActionsFailed,
	}
}

// ToPtr returns a *AlertModificationEvent pointing to the current value.
func (c AlertModificationEvent) ToPtr() *AlertModificationEvent {
	return &c
}

// AlertRuleState - The alert rule state.
type AlertRuleState string

const (
	AlertRuleStateDisabled AlertRuleState = "Disabled"
	AlertRuleStateEnabled  AlertRuleState = "Enabled"
)

// PossibleAlertRuleStateValues returns the possible values for the AlertRuleState const type.
func PossibleAlertRuleStateValues() []AlertRuleState {
	return []AlertRuleState{
		AlertRuleStateDisabled,
		AlertRuleStateEnabled,
	}
}

// ToPtr returns a *AlertRuleState pointing to the current value.
func (c AlertRuleState) ToPtr() *AlertRuleState {
	return &c
}

type AlertState string

const (
	AlertStateAcknowledged AlertState = "Acknowledged"
	AlertStateClosed       AlertState = "Closed"
	AlertStateNew          AlertState = "New"
)

// PossibleAlertStateValues returns the possible values for the AlertState const type.
func PossibleAlertStateValues() []AlertState {
	return []AlertState{
		AlertStateAcknowledged,
		AlertStateClosed,
		AlertStateNew,
	}
}

// ToPtr returns a *AlertState pointing to the current value.
func (c AlertState) ToPtr() *AlertState {
	return &c
}

type AlertsSortByFields string

const (
	AlertsSortByFieldsAlertState           AlertsSortByFields = "alertState"
	AlertsSortByFieldsLastModifiedDateTime AlertsSortByFields = "lastModifiedDateTime"
	AlertsSortByFieldsMonitorCondition     AlertsSortByFields = "monitorCondition"
	AlertsSortByFieldsName                 AlertsSortByFields = "name"
	AlertsSortByFieldsSeverity             AlertsSortByFields = "severity"
	AlertsSortByFieldsStartDateTime        AlertsSortByFields = "startDateTime"
	AlertsSortByFieldsTargetResource       AlertsSortByFields = "targetResource"
	AlertsSortByFieldsTargetResourceGroup  AlertsSortByFields = "targetResourceGroup"
	AlertsSortByFieldsTargetResourceName   AlertsSortByFields = "targetResourceName"
	AlertsSortByFieldsTargetResourceType   AlertsSortByFields = "targetResourceType"
)

// PossibleAlertsSortByFieldsValues returns the possible values for the AlertsSortByFields const type.
func PossibleAlertsSortByFieldsValues() []AlertsSortByFields {
	return []AlertsSortByFields{
		AlertsSortByFieldsAlertState,
		AlertsSortByFieldsLastModifiedDateTime,
		AlertsSortByFieldsMonitorCondition,
		AlertsSortByFieldsName,
		AlertsSortByFieldsSeverity,
		AlertsSortByFieldsStartDateTime,
		AlertsSortByFieldsTargetResource,
		AlertsSortByFieldsTargetResourceGroup,
		AlertsSortByFieldsTargetResourceName,
		AlertsSortByFieldsTargetResourceType,
	}
}

// ToPtr returns a *AlertsSortByFields pointing to the current value.
func (c AlertsSortByFields) ToPtr() *AlertsSortByFields {
	return &c
}

type AlertsSummaryGroupByFields string

const (
	AlertsSummaryGroupByFieldsAlertRule        AlertsSummaryGroupByFields = "alertRule"
	AlertsSummaryGroupByFieldsAlertState       AlertsSummaryGroupByFields = "alertState"
	AlertsSummaryGroupByFieldsMonitorCondition AlertsSummaryGroupByFields = "monitorCondition"
	AlertsSummaryGroupByFieldsMonitorService   AlertsSummaryGroupByFields = "monitorService"
	AlertsSummaryGroupByFieldsSeverity         AlertsSummaryGroupByFields = "severity"
	AlertsSummaryGroupByFieldsSignalType       AlertsSummaryGroupByFields = "signalType"
)

// PossibleAlertsSummaryGroupByFieldsValues returns the possible values for the AlertsSummaryGroupByFields const type.
func PossibleAlertsSummaryGroupByFieldsValues() []AlertsSummaryGroupByFields {
	return []AlertsSummaryGroupByFields{
		AlertsSummaryGroupByFieldsAlertRule,
		AlertsSummaryGroupByFieldsAlertState,
		AlertsSummaryGroupByFieldsMonitorCondition,
		AlertsSummaryGroupByFieldsMonitorService,
		AlertsSummaryGroupByFieldsSeverity,
		AlertsSummaryGroupByFieldsSignalType,
	}
}

// ToPtr returns a *AlertsSummaryGroupByFields pointing to the current value.
func (c AlertsSummaryGroupByFields) ToPtr() *AlertsSummaryGroupByFields {
	return &c
}

type Enum11 string

const (
	Enum11Asc  Enum11 = "asc"
	Enum11Desc Enum11 = "desc"
)

// PossibleEnum11Values returns the possible values for the Enum11 const type.
func PossibleEnum11Values() []Enum11 {
	return []Enum11{
		Enum11Asc,
		Enum11Desc,
	}
}

// ToPtr returns a *Enum11 pointing to the current value.
func (c Enum11) ToPtr() *Enum11 {
	return &c
}

type Identifier string

const (
	IdentifierMonitorServiceList Identifier = "MonitorServiceList"
)

// PossibleIdentifierValues returns the possible values for the Identifier const type.
func PossibleIdentifierValues() []Identifier {
	return []Identifier{
		IdentifierMonitorServiceList,
	}
}

// ToPtr returns a *Identifier pointing to the current value.
func (c Identifier) ToPtr() *Identifier {
	return &c
}

// MetadataIdentifier - Identification of the information to be retrieved by API call
type MetadataIdentifier string

const (
	MetadataIdentifierMonitorServiceList MetadataIdentifier = "MonitorServiceList"
)

// PossibleMetadataIdentifierValues returns the possible values for the MetadataIdentifier const type.
func PossibleMetadataIdentifierValues() []MetadataIdentifier {
	return []MetadataIdentifier{
		MetadataIdentifierMonitorServiceList,
	}
}

// ToPtr returns a *MetadataIdentifier pointing to the current value.
func (c MetadataIdentifier) ToPtr() *MetadataIdentifier {
	return &c
}

type MonitorCondition string

const (
	MonitorConditionFired    MonitorCondition = "Fired"
	MonitorConditionResolved MonitorCondition = "Resolved"
)

// PossibleMonitorConditionValues returns the possible values for the MonitorCondition const type.
func PossibleMonitorConditionValues() []MonitorCondition {
	return []MonitorCondition{
		MonitorConditionFired,
		MonitorConditionResolved,
	}
}

// ToPtr returns a *MonitorCondition pointing to the current value.
func (c MonitorCondition) ToPtr() *MonitorCondition {
	return &c
}

type MonitorService string

const (
	MonitorServiceActivityLogAdministrative MonitorService = "ActivityLog Administrative"
	MonitorServiceActivityLogAutoscale      MonitorService = "ActivityLog Autoscale"
	MonitorServiceActivityLogPolicy         MonitorService = "ActivityLog Policy"
	MonitorServiceActivityLogRecommendation MonitorService = "ActivityLog Recommendation"
	MonitorServiceActivityLogSecurity       MonitorService = "ActivityLog Security"
	MonitorServiceApplicationInsights       MonitorService = "Application Insights"
	MonitorServiceLogAnalytics              MonitorService = "Log Analytics"
	MonitorServiceNagios                    MonitorService = "Nagios"
	MonitorServicePlatform                  MonitorService = "Platform"
	MonitorServiceSCOM                      MonitorService = "SCOM"
	MonitorServiceServiceHealth             MonitorService = "ServiceHealth"
	MonitorServiceSmartDetector             MonitorService = "SmartDetector"
	MonitorServiceVMInsights                MonitorService = "VM Insights"
	MonitorServiceZabbix                    MonitorService = "Zabbix"
)

// PossibleMonitorServiceValues returns the possible values for the MonitorService const type.
func PossibleMonitorServiceValues() []MonitorService {
	return []MonitorService{
		MonitorServiceActivityLogAdministrative,
		MonitorServiceActivityLogAutoscale,
		MonitorServiceActivityLogPolicy,
		MonitorServiceActivityLogRecommendation,
		MonitorServiceActivityLogSecurity,
		MonitorServiceApplicationInsights,
		MonitorServiceLogAnalytics,
		MonitorServiceNagios,
		MonitorServicePlatform,
		MonitorServiceSCOM,
		MonitorServiceServiceHealth,
		MonitorServiceSmartDetector,
		MonitorServiceVMInsights,
		MonitorServiceZabbix,
	}
}

// ToPtr returns a *MonitorService pointing to the current value.
func (c MonitorService) ToPtr() *MonitorService {
	return &c
}

// Operator - operator for a given condition
type Operator string

const (
	OperatorContains       Operator = "Contains"
	OperatorDoesNotContain Operator = "DoesNotContain"
	OperatorEquals         Operator = "Equals"
	OperatorNotEquals      Operator = "NotEquals"
)

// PossibleOperatorValues returns the possible values for the Operator const type.
func PossibleOperatorValues() []Operator {
	return []Operator{
		OperatorContains,
		OperatorDoesNotContain,
		OperatorEquals,
		OperatorNotEquals,
	}
}

// ToPtr returns a *Operator pointing to the current value.
func (c Operator) ToPtr() *Operator {
	return &c
}

// ScopeType - type of target scope
type ScopeType string

const (
	ScopeTypeResource      ScopeType = "Resource"
	ScopeTypeResourceGroup ScopeType = "ResourceGroup"
	ScopeTypeSubscription  ScopeType = "Subscription"
)

// PossibleScopeTypeValues returns the possible values for the ScopeType const type.
func PossibleScopeTypeValues() []ScopeType {
	return []ScopeType{
		ScopeTypeResource,
		ScopeTypeResourceGroup,
		ScopeTypeSubscription,
	}
}

// ToPtr returns a *ScopeType pointing to the current value.
func (c ScopeType) ToPtr() *ScopeType {
	return &c
}

type Severity string

const (
	SeveritySev0 Severity = "Sev0"
	SeveritySev1 Severity = "Sev1"
	SeveritySev2 Severity = "Sev2"
	SeveritySev3 Severity = "Sev3"
	SeveritySev4 Severity = "Sev4"
)

// PossibleSeverityValues returns the possible values for the Severity const type.
func PossibleSeverityValues() []Severity {
	return []Severity{
		SeveritySev0,
		SeveritySev1,
		SeveritySev2,
		SeveritySev3,
		SeveritySev4,
	}
}

// ToPtr returns a *Severity pointing to the current value.
func (c Severity) ToPtr() *Severity {
	return &c
}

// SignalType - The type of signal the alert is based on, which could be metrics, logs or activity logs.
type SignalType string

const (
	SignalTypeLog     SignalType = "Log"
	SignalTypeMetric  SignalType = "Metric"
	SignalTypeUnknown SignalType = "Unknown"
)

// PossibleSignalTypeValues returns the possible values for the SignalType const type.
func PossibleSignalTypeValues() []SignalType {
	return []SignalType{
		SignalTypeLog,
		SignalTypeMetric,
		SignalTypeUnknown,
	}
}

// ToPtr returns a *SignalType pointing to the current value.
func (c SignalType) ToPtr() *SignalType {
	return &c
}

// SmartGroupModificationEvent - Reason for the modification
type SmartGroupModificationEvent string

const (
	SmartGroupModificationEventSmartGroupCreated SmartGroupModificationEvent = "SmartGroupCreated"
	SmartGroupModificationEventStateChange       SmartGroupModificationEvent = "StateChange"
	SmartGroupModificationEventAlertAdded        SmartGroupModificationEvent = "AlertAdded"
	SmartGroupModificationEventAlertRemoved      SmartGroupModificationEvent = "AlertRemoved"
)

// PossibleSmartGroupModificationEventValues returns the possible values for the SmartGroupModificationEvent const type.
func PossibleSmartGroupModificationEventValues() []SmartGroupModificationEvent {
	return []SmartGroupModificationEvent{
		SmartGroupModificationEventSmartGroupCreated,
		SmartGroupModificationEventStateChange,
		SmartGroupModificationEventAlertAdded,
		SmartGroupModificationEventAlertRemoved,
	}
}

// ToPtr returns a *SmartGroupModificationEvent pointing to the current value.
func (c SmartGroupModificationEvent) ToPtr() *SmartGroupModificationEvent {
	return &c
}

type SmartGroupsSortByFields string

const (
	SmartGroupsSortByFieldsAlertsCount          SmartGroupsSortByFields = "alertsCount"
	SmartGroupsSortByFieldsLastModifiedDateTime SmartGroupsSortByFields = "lastModifiedDateTime"
	SmartGroupsSortByFieldsSeverity             SmartGroupsSortByFields = "severity"
	SmartGroupsSortByFieldsStartDateTime        SmartGroupsSortByFields = "startDateTime"
	SmartGroupsSortByFieldsState                SmartGroupsSortByFields = "state"
)

// PossibleSmartGroupsSortByFieldsValues returns the possible values for the SmartGroupsSortByFields const type.
func PossibleSmartGroupsSortByFieldsValues() []SmartGroupsSortByFields {
	return []SmartGroupsSortByFields{
		SmartGroupsSortByFieldsAlertsCount,
		SmartGroupsSortByFieldsLastModifiedDateTime,
		SmartGroupsSortByFieldsSeverity,
		SmartGroupsSortByFieldsStartDateTime,
		SmartGroupsSortByFieldsState,
	}
}

// ToPtr returns a *SmartGroupsSortByFields pointing to the current value.
func (c SmartGroupsSortByFields) ToPtr() *SmartGroupsSortByFields {
	return &c
}

// State - Smart group state
type State string

const (
	StateAcknowledged State = "Acknowledged"
	StateClosed       State = "Closed"
	StateNew          State = "New"
)

// PossibleStateValues returns the possible values for the State const type.
func PossibleStateValues() []State {
	return []State{
		StateAcknowledged,
		StateClosed,
		StateNew,
	}
}

// ToPtr returns a *State pointing to the current value.
func (c State) ToPtr() *State {
	return &c
}

// SuppressionType - Specifies when the suppression should be applied
type SuppressionType string

const (
	SuppressionTypeAlways  SuppressionType = "Always"
	SuppressionTypeDaily   SuppressionType = "Daily"
	SuppressionTypeMonthly SuppressionType = "Monthly"
	SuppressionTypeOnce    SuppressionType = "Once"
	SuppressionTypeWeekly  SuppressionType = "Weekly"
)

// PossibleSuppressionTypeValues returns the possible values for the SuppressionType const type.
func PossibleSuppressionTypeValues() []SuppressionType {
	return []SuppressionType{
		SuppressionTypeAlways,
		SuppressionTypeDaily,
		SuppressionTypeMonthly,
		SuppressionTypeOnce,
		SuppressionTypeWeekly,
	}
}

// ToPtr returns a *SuppressionType pointing to the current value.
func (c SuppressionType) ToPtr() *SuppressionType {
	return &c
}

type TimeRange string

const (
	TimeRangeOneD    TimeRange = "1d"
	TimeRangeOneH    TimeRange = "1h"
	TimeRangeSevenD  TimeRange = "7d"
	TimeRangeThirtyD TimeRange = "30d"
)

// PossibleTimeRangeValues returns the possible values for the TimeRange const type.
func PossibleTimeRangeValues() []TimeRange {
	return []TimeRange{
		TimeRangeOneD,
		TimeRangeOneH,
		TimeRangeSevenD,
		TimeRangeThirtyD,
	}
}

// ToPtr returns a *TimeRange pointing to the current value.
func (c TimeRange) ToPtr() *TimeRange {
	return &c
}