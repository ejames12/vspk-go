/*
  Copyright (c) 2015, Alcatel-Lucent Inc
  All rights reserved.

  Redistribution and use in source and binary forms, with or without
  modification, are permitted provided that the following conditions are met:
      * Redistributions of source code must retain the above copyright
        notice, this list of conditions and the following disclaimer.
      * Redistributions in binary form must reproduce the above copyright
        notice, this list of conditions and the following disclaimer in the
        documentation and/or other materials provided with the distribution.
      * Neither the name of the copyright holder nor the names of its contributors
        may be used to endorse or promote products derived from this software without
        specific prior written permission.

  THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
  ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
  WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
  DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY
  DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
  (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
  LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
  ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
  (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
  SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// NSGatewayIdentity represents the Identity of the object
var NSGatewayIdentity = bambou.Identity{
	Name:     "nsgateway",
	Category: "nsgateways",
}

// NSGatewaysList represents a list of NSGateways
type NSGatewaysList []*NSGateway

// NSGatewaysAncestor is the interface of an ancestor of a NSGateway must implement.
type NSGatewaysAncestor interface {
	NSGateways(*bambou.FetchingInfo) (NSGatewaysList, *bambou.Error)
	CreateNSGateways(*NSGateway) *bambou.Error
}

// NSGateway represents the model of a nsgateway
type NSGateway struct {
	ID                                 string `json:"ID,omitempty"`
	ParentID                           string `json:"parentID,omitempty"`
	ParentType                         string `json:"parentType,omitempty"`
	Owner                              string `json:"owner,omitempty"`
	NATTraversalEnabled                bool   `json:"NATTraversalEnabled"`
	TPMStatus                          string `json:"TPMStatus,omitempty"`
	Name                               string `json:"name,omitempty"`
	LastConfigurationReloadTimestamp   int    `json:"lastConfigurationReloadTimestamp,omitempty"`
	LastUpdatedBy                      string `json:"lastUpdatedBy,omitempty"`
	DatapathID                         string `json:"datapathID,omitempty"`
	RedundancyGroupID                  string `json:"redundancyGroupID,omitempty"`
	TemplateID                         string `json:"templateID,omitempty"`
	Pending                            bool   `json:"pending"`
	PermittedAction                    string `json:"permittedAction,omitempty"`
	Personality                        string `json:"personality,omitempty"`
	Description                        string `json:"description,omitempty"`
	EnterpriseID                       string `json:"enterpriseID,omitempty"`
	EntityScope                        string `json:"entityScope,omitempty"`
	LocationID                         string `json:"locationID,omitempty"`
	ConfigurationReloadState           string `json:"configurationReloadState,omitempty"`
	ConfigurationStatus                string `json:"configurationStatus,omitempty"`
	BootstrapID                        string `json:"bootstrapID,omitempty"`
	BootstrapStatus                    string `json:"bootstrapStatus,omitempty"`
	AssociatedGatewaySecurityID        string `json:"associatedGatewaySecurityID,omitempty"`
	AssociatedGatewaySecurityProfileID string `json:"associatedGatewaySecurityProfileID,omitempty"`
	AutoDiscGatewayID                  string `json:"autoDiscGatewayID,omitempty"`
	ExternalID                         string `json:"externalID,omitempty"`
	SystemID                           string `json:"systemID,omitempty"`
}

// NewNSGateway returns a new *NSGateway
func NewNSGateway() *NSGateway {

	return &NSGateway{}
}

// Identity returns the Identity of the object.
func (o *NSGateway) Identity() bambou.Identity {

	return NSGatewayIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *NSGateway) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *NSGateway) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the NSGateway from the server
func (o *NSGateway) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the NSGateway into the server
func (o *NSGateway) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the NSGateway from the server
func (o *NSGateway) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// GatewaySecurities retrieves the list of child GatewaySecurities of the NSGateway
func (o *NSGateway) GatewaySecurities(info *bambou.FetchingInfo) (GatewaySecuritiesList, *bambou.Error) {

	var list GatewaySecuritiesList
	err := bambou.CurrentSession().FetchChildren(o, GatewaySecurityIdentity, &list, info)
	return list, err
}

// CreateGatewaySecurity creates a new child GatewaySecurity under the NSGateway
func (o *NSGateway) CreateGatewaySecurity(child *GatewaySecurity) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// PATNATPools retrieves the list of child PATNATPools of the NSGateway
func (o *NSGateway) PATNATPools(info *bambou.FetchingInfo) (PATNATPoolsList, *bambou.Error) {

	var list PATNATPoolsList
	err := bambou.CurrentSession().FetchChildren(o, PATNATPoolIdentity, &list, info)
	return list, err
}

// AssignPATNATPools assigns the list of PATNATPools to the NSGateway
func (o *NSGateway) AssignPATNATPools(children PATNATPoolsList) *bambou.Error {

	list := []bambou.Identifiable{}
	for _, c := range children {
		list = append(list, c)
	}

	return bambou.CurrentSession().AssignChildren(o, list, PATNATPoolIdentity)
}

// Permissions retrieves the list of child Permissions of the NSGateway
func (o *NSGateway) Permissions(info *bambou.FetchingInfo) (PermissionsList, *bambou.Error) {

	var list PermissionsList
	err := bambou.CurrentSession().FetchChildren(o, PermissionIdentity, &list, info)
	return list, err
}

// CreatePermission creates a new child Permission under the NSGateway
func (o *NSGateway) CreatePermission(child *Permission) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Metadatas retrieves the list of child Metadatas of the NSGateway
func (o *NSGateway) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the NSGateway
func (o *NSGateway) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Alarms retrieves the list of child Alarms of the NSGateway
func (o *NSGateway) Alarms(info *bambou.FetchingInfo) (AlarmsList, *bambou.Error) {

	var list AlarmsList
	err := bambou.CurrentSession().FetchChildren(o, AlarmIdentity, &list, info)
	return list, err
}

// CreateAlarm creates a new child Alarm under the NSGateway
func (o *NSGateway) CreateAlarm(child *Alarm) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the NSGateway
func (o *NSGateway) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the NSGateway
func (o *NSGateway) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// InfrastructureConfigs retrieves the list of child InfrastructureConfigs of the NSGateway
func (o *NSGateway) InfrastructureConfigs(info *bambou.FetchingInfo) (InfrastructureConfigsList, *bambou.Error) {

	var list InfrastructureConfigsList
	err := bambou.CurrentSession().FetchChildren(o, InfrastructureConfigIdentity, &list, info)
	return list, err
}

// CreateInfrastructureConfig creates a new child InfrastructureConfig under the NSGateway
func (o *NSGateway) CreateInfrastructureConfig(child *InfrastructureConfig) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// EnterprisePermissions retrieves the list of child EnterprisePermissions of the NSGateway
func (o *NSGateway) EnterprisePermissions(info *bambou.FetchingInfo) (EnterprisePermissionsList, *bambou.Error) {

	var list EnterprisePermissionsList
	err := bambou.CurrentSession().FetchChildren(o, EnterprisePermissionIdentity, &list, info)
	return list, err
}

// CreateEnterprisePermission creates a new child EnterprisePermission under the NSGateway
func (o *NSGateway) CreateEnterprisePermission(child *EnterprisePermission) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Jobs retrieves the list of child Jobs of the NSGateway
func (o *NSGateway) Jobs(info *bambou.FetchingInfo) (JobsList, *bambou.Error) {

	var list JobsList
	err := bambou.CurrentSession().FetchChildren(o, JobIdentity, &list, info)
	return list, err
}

// CreateJob creates a new child Job under the NSGateway
func (o *NSGateway) CreateJob(child *Job) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Locations retrieves the list of child Locations of the NSGateway
func (o *NSGateway) Locations(info *bambou.FetchingInfo) (LocationsList, *bambou.Error) {

	var list LocationsList
	err := bambou.CurrentSession().FetchChildren(o, LocationIdentity, &list, info)
	return list, err
}

// CreateLocation creates a new child Location under the NSGateway
func (o *NSGateway) CreateLocation(child *Location) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Bootstraps retrieves the list of child Bootstraps of the NSGateway
func (o *NSGateway) Bootstraps(info *bambou.FetchingInfo) (BootstrapsList, *bambou.Error) {

	var list BootstrapsList
	err := bambou.CurrentSession().FetchChildren(o, BootstrapIdentity, &list, info)
	return list, err
}

// CreateBootstrap creates a new child Bootstrap under the NSGateway
func (o *NSGateway) CreateBootstrap(child *Bootstrap) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// BootstrapActivations retrieves the list of child BootstrapActivations of the NSGateway
func (o *NSGateway) BootstrapActivations(info *bambou.FetchingInfo) (BootstrapActivationsList, *bambou.Error) {

	var list BootstrapActivationsList
	err := bambou.CurrentSession().FetchChildren(o, BootstrapActivationIdentity, &list, info)
	return list, err
}

// CreateBootstrapActivation creates a new child BootstrapActivation under the NSGateway
func (o *NSGateway) CreateBootstrapActivation(child *BootstrapActivation) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// NSPorts retrieves the list of child NSPorts of the NSGateway
func (o *NSGateway) NSPorts(info *bambou.FetchingInfo) (NSPortsList, *bambou.Error) {

	var list NSPortsList
	err := bambou.CurrentSession().FetchChildren(o, NSPortIdentity, &list, info)
	return list, err
}

// CreateNSPort creates a new child NSPort under the NSGateway
func (o *NSGateway) CreateNSPort(child *NSPort) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Subnets retrieves the list of child Subnets of the NSGateway
func (o *NSGateway) Subnets(info *bambou.FetchingInfo) (SubnetsList, *bambou.Error) {

	var list SubnetsList
	err := bambou.CurrentSession().FetchChildren(o, SubnetIdentity, &list, info)
	return list, err
}

// CreateSubnet creates a new child Subnet under the NSGateway
func (o *NSGateway) CreateSubnet(child *Subnet) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// EventLogs retrieves the list of child EventLogs of the NSGateway
func (o *NSGateway) EventLogs(info *bambou.FetchingInfo) (EventLogsList, *bambou.Error) {

	var list EventLogsList
	err := bambou.CurrentSession().FetchChildren(o, EventLogIdentity, &list, info)
	return list, err
}

// CreateEventLog creates a new child EventLog under the NSGateway
func (o *NSGateway) CreateEventLog(child *EventLog) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
