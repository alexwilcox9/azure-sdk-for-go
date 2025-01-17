//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armresourcemover

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type AffectedMoveResource.
func (a AffectedMoveResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", a.ID)
	populate(objectMap, "moveResources", a.MoveResources)
	populate(objectMap, "sourceId", a.SourceID)
	return json.Marshal(objectMap)
}

// GetResourceSettings implements the ResourceSettingsClassification interface for type AvailabilitySetResourceSettings.
func (a *AvailabilitySetResourceSettings) GetResourceSettings() *ResourceSettings {
	return &ResourceSettings{
		ResourceType:       a.ResourceType,
		TargetResourceName: a.TargetResourceName,
	}
}

// MarshalJSON implements the json.Marshaller interface for type AvailabilitySetResourceSettings.
func (a AvailabilitySetResourceSettings) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "faultDomain", a.FaultDomain)
	objectMap["resourceType"] = "Microsoft.Compute/availabilitySets"
	populate(objectMap, "tags", a.Tags)
	populate(objectMap, "targetResourceName", a.TargetResourceName)
	populate(objectMap, "updateDomain", a.UpdateDomain)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type AvailabilitySetResourceSettings.
func (a *AvailabilitySetResourceSettings) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "faultDomain":
			err = unpopulate(val, &a.FaultDomain)
			delete(rawMsg, key)
		case "resourceType":
			err = unpopulate(val, &a.ResourceType)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, &a.Tags)
			delete(rawMsg, key)
		case "targetResourceName":
			err = unpopulate(val, &a.TargetResourceName)
			delete(rawMsg, key)
		case "updateDomain":
			err = unpopulate(val, &a.UpdateDomain)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type BulkRemoveRequest.
func (b BulkRemoveRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "moveResourceInputType", b.MoveResourceInputType)
	populate(objectMap, "moveResources", b.MoveResources)
	populate(objectMap, "validateOnly", b.ValidateOnly)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type CloudErrorBody.
func (c CloudErrorBody) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "code", c.Code)
	populate(objectMap, "details", c.Details)
	populate(objectMap, "message", c.Message)
	populate(objectMap, "target", c.Target)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type CommitRequest.
func (c CommitRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "moveResourceInputType", c.MoveResourceInputType)
	populate(objectMap, "moveResources", c.MoveResources)
	populate(objectMap, "validateOnly", c.ValidateOnly)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type DiscardRequest.
func (d DiscardRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "moveResourceInputType", d.MoveResourceInputType)
	populate(objectMap, "moveResources", d.MoveResources)
	populate(objectMap, "validateOnly", d.ValidateOnly)
	return json.Marshal(objectMap)
}

// GetResourceSettings implements the ResourceSettingsClassification interface for type DiskEncryptionSetResourceSettings.
func (d *DiskEncryptionSetResourceSettings) GetResourceSettings() *ResourceSettings {
	return &ResourceSettings{
		ResourceType:       d.ResourceType,
		TargetResourceName: d.TargetResourceName,
	}
}

// MarshalJSON implements the json.Marshaller interface for type DiskEncryptionSetResourceSettings.
func (d DiskEncryptionSetResourceSettings) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	objectMap["resourceType"] = "Microsoft.Compute/diskEncryptionSets"
	populate(objectMap, "targetResourceName", d.TargetResourceName)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DiskEncryptionSetResourceSettings.
func (d *DiskEncryptionSetResourceSettings) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "resourceType":
			err = unpopulate(val, &d.ResourceType)
			delete(rawMsg, key)
		case "targetResourceName":
			err = unpopulate(val, &d.TargetResourceName)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// GetResourceSettings implements the ResourceSettingsClassification interface for type KeyVaultResourceSettings.
func (k *KeyVaultResourceSettings) GetResourceSettings() *ResourceSettings {
	return &ResourceSettings{
		ResourceType:       k.ResourceType,
		TargetResourceName: k.TargetResourceName,
	}
}

// MarshalJSON implements the json.Marshaller interface for type KeyVaultResourceSettings.
func (k KeyVaultResourceSettings) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	objectMap["resourceType"] = "Microsoft.KeyVault/vaults"
	populate(objectMap, "targetResourceName", k.TargetResourceName)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type KeyVaultResourceSettings.
func (k *KeyVaultResourceSettings) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "resourceType":
			err = unpopulate(val, &k.ResourceType)
			delete(rawMsg, key)
		case "targetResourceName":
			err = unpopulate(val, &k.TargetResourceName)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// GetResourceSettings implements the ResourceSettingsClassification interface for type LoadBalancerResourceSettings.
func (l *LoadBalancerResourceSettings) GetResourceSettings() *ResourceSettings {
	return &ResourceSettings{
		ResourceType:       l.ResourceType,
		TargetResourceName: l.TargetResourceName,
	}
}

// MarshalJSON implements the json.Marshaller interface for type LoadBalancerResourceSettings.
func (l LoadBalancerResourceSettings) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "backendAddressPools", l.BackendAddressPools)
	populate(objectMap, "frontendIPConfigurations", l.FrontendIPConfigurations)
	objectMap["resourceType"] = "Microsoft.Network/loadBalancers"
	populate(objectMap, "sku", l.SKU)
	populate(objectMap, "tags", l.Tags)
	populate(objectMap, "targetResourceName", l.TargetResourceName)
	populate(objectMap, "zones", l.Zones)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type LoadBalancerResourceSettings.
func (l *LoadBalancerResourceSettings) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "backendAddressPools":
			err = unpopulate(val, &l.BackendAddressPools)
			delete(rawMsg, key)
		case "frontendIPConfigurations":
			err = unpopulate(val, &l.FrontendIPConfigurations)
			delete(rawMsg, key)
		case "resourceType":
			err = unpopulate(val, &l.ResourceType)
			delete(rawMsg, key)
		case "sku":
			err = unpopulate(val, &l.SKU)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, &l.Tags)
			delete(rawMsg, key)
		case "targetResourceName":
			err = unpopulate(val, &l.TargetResourceName)
			delete(rawMsg, key)
		case "zones":
			err = unpopulate(val, &l.Zones)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type MoveCollection.
func (m MoveCollection) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "etag", m.Etag)
	populate(objectMap, "id", m.ID)
	populate(objectMap, "identity", m.Identity)
	populate(objectMap, "location", m.Location)
	populate(objectMap, "name", m.Name)
	populate(objectMap, "properties", m.Properties)
	populate(objectMap, "systemData", m.SystemData)
	populate(objectMap, "tags", m.Tags)
	populate(objectMap, "type", m.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type MoveCollectionResultList.
func (m MoveCollectionResultList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", m.NextLink)
	populate(objectMap, "value", m.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type MoveErrorInfo.
func (m MoveErrorInfo) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "moveResources", m.MoveResources)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type MoveResourceCollection.
func (m MoveResourceCollection) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", m.NextLink)
	populate(objectMap, "summaryCollection", m.SummaryCollection)
	populate(objectMap, "totalCount", m.TotalCount)
	populate(objectMap, "value", m.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type MoveResourceErrorBody.
func (m MoveResourceErrorBody) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "code", m.Code)
	populate(objectMap, "details", m.Details)
	populate(objectMap, "message", m.Message)
	populate(objectMap, "target", m.Target)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type MoveResourceProperties.
func (m MoveResourceProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "dependsOn", m.DependsOn)
	populate(objectMap, "dependsOnOverrides", m.DependsOnOverrides)
	populate(objectMap, "errors", m.Errors)
	populate(objectMap, "existingTargetId", m.ExistingTargetID)
	populate(objectMap, "isResolveRequired", m.IsResolveRequired)
	populate(objectMap, "moveStatus", m.MoveStatus)
	populate(objectMap, "provisioningState", m.ProvisioningState)
	populate(objectMap, "resourceSettings", m.ResourceSettings)
	populate(objectMap, "sourceId", m.SourceID)
	populate(objectMap, "sourceResourceSettings", m.SourceResourceSettings)
	populate(objectMap, "targetId", m.TargetID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MoveResourceProperties.
func (m *MoveResourceProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "dependsOn":
			err = unpopulate(val, &m.DependsOn)
			delete(rawMsg, key)
		case "dependsOnOverrides":
			err = unpopulate(val, &m.DependsOnOverrides)
			delete(rawMsg, key)
		case "errors":
			err = unpopulate(val, &m.Errors)
			delete(rawMsg, key)
		case "existingTargetId":
			err = unpopulate(val, &m.ExistingTargetID)
			delete(rawMsg, key)
		case "isResolveRequired":
			err = unpopulate(val, &m.IsResolveRequired)
			delete(rawMsg, key)
		case "moveStatus":
			err = unpopulate(val, &m.MoveStatus)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, &m.ProvisioningState)
			delete(rawMsg, key)
		case "resourceSettings":
			m.ResourceSettings, err = unmarshalResourceSettingsClassification(val)
			delete(rawMsg, key)
		case "sourceId":
			err = unpopulate(val, &m.SourceID)
			delete(rawMsg, key)
		case "sourceResourceSettings":
			m.SourceResourceSettings, err = unmarshalResourceSettingsClassification(val)
			delete(rawMsg, key)
		case "targetId":
			err = unpopulate(val, &m.TargetID)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// GetResourceSettings implements the ResourceSettingsClassification interface for type NetworkInterfaceResourceSettings.
func (n *NetworkInterfaceResourceSettings) GetResourceSettings() *ResourceSettings {
	return &ResourceSettings{
		ResourceType:       n.ResourceType,
		TargetResourceName: n.TargetResourceName,
	}
}

// MarshalJSON implements the json.Marshaller interface for type NetworkInterfaceResourceSettings.
func (n NetworkInterfaceResourceSettings) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "enableAcceleratedNetworking", n.EnableAcceleratedNetworking)
	populate(objectMap, "ipConfigurations", n.IPConfigurations)
	objectMap["resourceType"] = "Microsoft.Network/networkInterfaces"
	populate(objectMap, "tags", n.Tags)
	populate(objectMap, "targetResourceName", n.TargetResourceName)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type NetworkInterfaceResourceSettings.
func (n *NetworkInterfaceResourceSettings) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "enableAcceleratedNetworking":
			err = unpopulate(val, &n.EnableAcceleratedNetworking)
			delete(rawMsg, key)
		case "ipConfigurations":
			err = unpopulate(val, &n.IPConfigurations)
			delete(rawMsg, key)
		case "resourceType":
			err = unpopulate(val, &n.ResourceType)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, &n.Tags)
			delete(rawMsg, key)
		case "targetResourceName":
			err = unpopulate(val, &n.TargetResourceName)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// GetResourceSettings implements the ResourceSettingsClassification interface for type NetworkSecurityGroupResourceSettings.
func (n *NetworkSecurityGroupResourceSettings) GetResourceSettings() *ResourceSettings {
	return &ResourceSettings{
		ResourceType:       n.ResourceType,
		TargetResourceName: n.TargetResourceName,
	}
}

// MarshalJSON implements the json.Marshaller interface for type NetworkSecurityGroupResourceSettings.
func (n NetworkSecurityGroupResourceSettings) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	objectMap["resourceType"] = "Microsoft.Network/networkSecurityGroups"
	populate(objectMap, "securityRules", n.SecurityRules)
	populate(objectMap, "tags", n.Tags)
	populate(objectMap, "targetResourceName", n.TargetResourceName)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type NetworkSecurityGroupResourceSettings.
func (n *NetworkSecurityGroupResourceSettings) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "resourceType":
			err = unpopulate(val, &n.ResourceType)
			delete(rawMsg, key)
		case "securityRules":
			err = unpopulate(val, &n.SecurityRules)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, &n.Tags)
			delete(rawMsg, key)
		case "targetResourceName":
			err = unpopulate(val, &n.TargetResourceName)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type NicIPConfigurationResourceSettings.
func (n NicIPConfigurationResourceSettings) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "loadBalancerBackendAddressPools", n.LoadBalancerBackendAddressPools)
	populate(objectMap, "loadBalancerNatRules", n.LoadBalancerNatRules)
	populate(objectMap, "name", n.Name)
	populate(objectMap, "primary", n.Primary)
	populate(objectMap, "privateIpAddress", n.PrivateIPAddress)
	populate(objectMap, "privateIpAllocationMethod", n.PrivateIPAllocationMethod)
	populate(objectMap, "publicIp", n.PublicIP)
	populate(objectMap, "subnet", n.Subnet)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type OperationStatusError.
func (o OperationStatusError) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "additionalInfo", o.AdditionalInfo)
	populate(objectMap, "code", o.Code)
	populate(objectMap, "details", o.Details)
	populate(objectMap, "message", o.Message)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type OperationsDiscoveryCollection.
func (o OperationsDiscoveryCollection) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", o.NextLink)
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type PrepareRequest.
func (p PrepareRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "moveResourceInputType", p.MoveResourceInputType)
	populate(objectMap, "moveResources", p.MoveResources)
	populate(objectMap, "validateOnly", p.ValidateOnly)
	return json.Marshal(objectMap)
}

// GetResourceSettings implements the ResourceSettingsClassification interface for type PublicIPAddressResourceSettings.
func (p *PublicIPAddressResourceSettings) GetResourceSettings() *ResourceSettings {
	return &ResourceSettings{
		ResourceType:       p.ResourceType,
		TargetResourceName: p.TargetResourceName,
	}
}

// MarshalJSON implements the json.Marshaller interface for type PublicIPAddressResourceSettings.
func (p PublicIPAddressResourceSettings) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "domainNameLabel", p.DomainNameLabel)
	populate(objectMap, "fqdn", p.Fqdn)
	populate(objectMap, "publicIpAllocationMethod", p.PublicIPAllocationMethod)
	objectMap["resourceType"] = "Microsoft.Network/publicIPAddresses"
	populate(objectMap, "sku", p.SKU)
	populate(objectMap, "tags", p.Tags)
	populate(objectMap, "targetResourceName", p.TargetResourceName)
	populate(objectMap, "zones", p.Zones)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PublicIPAddressResourceSettings.
func (p *PublicIPAddressResourceSettings) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "domainNameLabel":
			err = unpopulate(val, &p.DomainNameLabel)
			delete(rawMsg, key)
		case "fqdn":
			err = unpopulate(val, &p.Fqdn)
			delete(rawMsg, key)
		case "publicIpAllocationMethod":
			err = unpopulate(val, &p.PublicIPAllocationMethod)
			delete(rawMsg, key)
		case "resourceType":
			err = unpopulate(val, &p.ResourceType)
			delete(rawMsg, key)
		case "sku":
			err = unpopulate(val, &p.SKU)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, &p.Tags)
			delete(rawMsg, key)
		case "targetResourceName":
			err = unpopulate(val, &p.TargetResourceName)
			delete(rawMsg, key)
		case "zones":
			err = unpopulate(val, &p.Zones)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RequiredForResourcesCollection.
func (r RequiredForResourcesCollection) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "sourceIds", r.SourceIDs)
	return json.Marshal(objectMap)
}

// GetResourceSettings implements the ResourceSettingsClassification interface for type ResourceGroupResourceSettings.
func (r *ResourceGroupResourceSettings) GetResourceSettings() *ResourceSettings {
	return &ResourceSettings{
		ResourceType:       r.ResourceType,
		TargetResourceName: r.TargetResourceName,
	}
}

// MarshalJSON implements the json.Marshaller interface for type ResourceGroupResourceSettings.
func (r ResourceGroupResourceSettings) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	objectMap["resourceType"] = "resourceGroups"
	populate(objectMap, "targetResourceName", r.TargetResourceName)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ResourceGroupResourceSettings.
func (r *ResourceGroupResourceSettings) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "resourceType":
			err = unpopulate(val, &r.ResourceType)
			delete(rawMsg, key)
		case "targetResourceName":
			err = unpopulate(val, &r.TargetResourceName)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ResourceMoveRequest.
func (r ResourceMoveRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "moveResourceInputType", r.MoveResourceInputType)
	populate(objectMap, "moveResources", r.MoveResources)
	populate(objectMap, "validateOnly", r.ValidateOnly)
	return json.Marshal(objectMap)
}

// GetResourceSettings implements the ResourceSettingsClassification interface for type ResourceSettings.
func (r *ResourceSettings) GetResourceSettings() *ResourceSettings { return r }

// GetResourceSettings implements the ResourceSettingsClassification interface for type SQLDatabaseResourceSettings.
func (s *SQLDatabaseResourceSettings) GetResourceSettings() *ResourceSettings {
	return &ResourceSettings{
		ResourceType:       s.ResourceType,
		TargetResourceName: s.TargetResourceName,
	}
}

// MarshalJSON implements the json.Marshaller interface for type SQLDatabaseResourceSettings.
func (s SQLDatabaseResourceSettings) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	objectMap["resourceType"] = "Microsoft.Sql/servers/databases"
	populate(objectMap, "tags", s.Tags)
	populate(objectMap, "targetResourceName", s.TargetResourceName)
	populate(objectMap, "zoneRedundant", s.ZoneRedundant)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SQLDatabaseResourceSettings.
func (s *SQLDatabaseResourceSettings) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "resourceType":
			err = unpopulate(val, &s.ResourceType)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, &s.Tags)
			delete(rawMsg, key)
		case "targetResourceName":
			err = unpopulate(val, &s.TargetResourceName)
			delete(rawMsg, key)
		case "zoneRedundant":
			err = unpopulate(val, &s.ZoneRedundant)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// GetResourceSettings implements the ResourceSettingsClassification interface for type SQLElasticPoolResourceSettings.
func (s *SQLElasticPoolResourceSettings) GetResourceSettings() *ResourceSettings {
	return &ResourceSettings{
		ResourceType:       s.ResourceType,
		TargetResourceName: s.TargetResourceName,
	}
}

// MarshalJSON implements the json.Marshaller interface for type SQLElasticPoolResourceSettings.
func (s SQLElasticPoolResourceSettings) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	objectMap["resourceType"] = "Microsoft.Sql/servers/elasticPools"
	populate(objectMap, "tags", s.Tags)
	populate(objectMap, "targetResourceName", s.TargetResourceName)
	populate(objectMap, "zoneRedundant", s.ZoneRedundant)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SQLElasticPoolResourceSettings.
func (s *SQLElasticPoolResourceSettings) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "resourceType":
			err = unpopulate(val, &s.ResourceType)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, &s.Tags)
			delete(rawMsg, key)
		case "targetResourceName":
			err = unpopulate(val, &s.TargetResourceName)
			delete(rawMsg, key)
		case "zoneRedundant":
			err = unpopulate(val, &s.ZoneRedundant)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// GetResourceSettings implements the ResourceSettingsClassification interface for type SQLServerResourceSettings.
func (s *SQLServerResourceSettings) GetResourceSettings() *ResourceSettings {
	return &ResourceSettings{
		ResourceType:       s.ResourceType,
		TargetResourceName: s.TargetResourceName,
	}
}

// MarshalJSON implements the json.Marshaller interface for type SQLServerResourceSettings.
func (s SQLServerResourceSettings) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	objectMap["resourceType"] = "Microsoft.Sql/servers"
	populate(objectMap, "targetResourceName", s.TargetResourceName)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SQLServerResourceSettings.
func (s *SQLServerResourceSettings) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "resourceType":
			err = unpopulate(val, &s.ResourceType)
			delete(rawMsg, key)
		case "targetResourceName":
			err = unpopulate(val, &s.TargetResourceName)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SummaryCollection.
func (s SummaryCollection) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "fieldName", s.FieldName)
	populate(objectMap, "summary", s.Summary)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SystemData.
func (s SystemData) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "createdAt", s.CreatedAt)
	populate(objectMap, "createdBy", s.CreatedBy)
	populate(objectMap, "createdByType", s.CreatedByType)
	populateTimeRFC3339(objectMap, "lastModifiedAt", s.LastModifiedAt)
	populate(objectMap, "lastModifiedBy", s.LastModifiedBy)
	populate(objectMap, "lastModifiedByType", s.LastModifiedByType)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SystemData.
func (s *SystemData) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "createdAt":
			err = unpopulateTimeRFC3339(val, &s.CreatedAt)
			delete(rawMsg, key)
		case "createdBy":
			err = unpopulate(val, &s.CreatedBy)
			delete(rawMsg, key)
		case "createdByType":
			err = unpopulate(val, &s.CreatedByType)
			delete(rawMsg, key)
		case "lastModifiedAt":
			err = unpopulateTimeRFC3339(val, &s.LastModifiedAt)
			delete(rawMsg, key)
		case "lastModifiedBy":
			err = unpopulate(val, &s.LastModifiedBy)
			delete(rawMsg, key)
		case "lastModifiedByType":
			err = unpopulate(val, &s.LastModifiedByType)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type UnresolvedDependencyCollection.
func (u UnresolvedDependencyCollection) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", u.NextLink)
	populate(objectMap, "summaryCollection", u.SummaryCollection)
	populate(objectMap, "totalCount", u.TotalCount)
	populate(objectMap, "value", u.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type UpdateMoveCollectionRequest.
func (u UpdateMoveCollectionRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "identity", u.Identity)
	populate(objectMap, "tags", u.Tags)
	return json.Marshal(objectMap)
}

// GetResourceSettings implements the ResourceSettingsClassification interface for type VirtualMachineResourceSettings.
func (v *VirtualMachineResourceSettings) GetResourceSettings() *ResourceSettings {
	return &ResourceSettings{
		ResourceType:       v.ResourceType,
		TargetResourceName: v.TargetResourceName,
	}
}

// MarshalJSON implements the json.Marshaller interface for type VirtualMachineResourceSettings.
func (v VirtualMachineResourceSettings) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	objectMap["resourceType"] = "Microsoft.Compute/virtualMachines"
	populate(objectMap, "tags", v.Tags)
	populate(objectMap, "targetAvailabilitySetId", v.TargetAvailabilitySetID)
	populate(objectMap, "targetAvailabilityZone", v.TargetAvailabilityZone)
	populate(objectMap, "targetResourceName", v.TargetResourceName)
	populate(objectMap, "targetVmSize", v.TargetVMSize)
	populate(objectMap, "userManagedIdentities", v.UserManagedIdentities)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type VirtualMachineResourceSettings.
func (v *VirtualMachineResourceSettings) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "resourceType":
			err = unpopulate(val, &v.ResourceType)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, &v.Tags)
			delete(rawMsg, key)
		case "targetAvailabilitySetId":
			err = unpopulate(val, &v.TargetAvailabilitySetID)
			delete(rawMsg, key)
		case "targetAvailabilityZone":
			err = unpopulate(val, &v.TargetAvailabilityZone)
			delete(rawMsg, key)
		case "targetResourceName":
			err = unpopulate(val, &v.TargetResourceName)
			delete(rawMsg, key)
		case "targetVmSize":
			err = unpopulate(val, &v.TargetVMSize)
			delete(rawMsg, key)
		case "userManagedIdentities":
			err = unpopulate(val, &v.UserManagedIdentities)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// GetResourceSettings implements the ResourceSettingsClassification interface for type VirtualNetworkResourceSettings.
func (v *VirtualNetworkResourceSettings) GetResourceSettings() *ResourceSettings {
	return &ResourceSettings{
		ResourceType:       v.ResourceType,
		TargetResourceName: v.TargetResourceName,
	}
}

// MarshalJSON implements the json.Marshaller interface for type VirtualNetworkResourceSettings.
func (v VirtualNetworkResourceSettings) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "addressSpace", v.AddressSpace)
	populate(objectMap, "dnsServers", v.DNSServers)
	populate(objectMap, "enableDdosProtection", v.EnableDdosProtection)
	objectMap["resourceType"] = "Microsoft.Network/virtualNetworks"
	populate(objectMap, "subnets", v.Subnets)
	populate(objectMap, "tags", v.Tags)
	populate(objectMap, "targetResourceName", v.TargetResourceName)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type VirtualNetworkResourceSettings.
func (v *VirtualNetworkResourceSettings) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "addressSpace":
			err = unpopulate(val, &v.AddressSpace)
			delete(rawMsg, key)
		case "dnsServers":
			err = unpopulate(val, &v.DNSServers)
			delete(rawMsg, key)
		case "enableDdosProtection":
			err = unpopulate(val, &v.EnableDdosProtection)
			delete(rawMsg, key)
		case "resourceType":
			err = unpopulate(val, &v.ResourceType)
			delete(rawMsg, key)
		case "subnets":
			err = unpopulate(val, &v.Subnets)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, &v.Tags)
			delete(rawMsg, key)
		case "targetResourceName":
			err = unpopulate(val, &v.TargetResourceName)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func populate(m map[string]interface{}, k string, v interface{}) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, v interface{}) error {
	if data == nil {
		return nil
	}
	return json.Unmarshal(data, v)
}
