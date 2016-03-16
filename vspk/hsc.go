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

// HSCIdentity represents the Identity of the object
var HSCIdentity = bambou.Identity{
	Name:     "hsc",
	Category: "hscs",
}

// HSCsList represents a list of HSCs
type HSCsList []*HSC

// HSCsAncestor is the interface of an ancestor of a HSC must implement.
type HSCsAncestor interface {
	HSCs(*bambou.FetchingInfo) (HSCsList, *bambou.Error)
	CreateHSCs(*HSC) *bambou.Error
}

// HSC represents the model of a hsc
type HSC struct {
	ID                          string        `json:"ID,omitempty"`
	ParentID                    string        `json:"parentID,omitempty"`
	ParentType                  string        `json:"parentType,omitempty"`
	Owner                       string        `json:"owner,omitempty"`
	Address                     string        `json:"address,omitempty"`
	AlreadyMarkedForUnavailable bool          `json:"alreadyMarkedForUnavailable"`
	AverageCPUUsage             float64       `json:"averageCPUUsage,omitempty"`
	AverageMemoryUsage          float64       `json:"averageMemoryUsage,omitempty"`
	CurrentCPUUsage             float64       `json:"currentCPUUsage,omitempty"`
	CurrentMemoryUsage          float64       `json:"currentMemoryUsage,omitempty"`
	Description                 string        `json:"description,omitempty"`
	Disks                       []interface{} `json:"disks,omitempty"`
	EntityScope                 string        `json:"entityScope,omitempty"`
	ExternalID                  string        `json:"externalID,omitempty"`
	LastStateChange             int           `json:"lastStateChange,omitempty"`
	LastUpdatedBy               string        `json:"lastUpdatedBy,omitempty"`
	Location                    string        `json:"location,omitempty"`
	ManagementIP                string        `json:"managementIP,omitempty"`
	Messages                    []interface{} `json:"messages,omitempty"`
	Model                       string        `json:"model,omitempty"`
	Name                        string        `json:"name,omitempty"`
	PeakCPUUsage                float64       `json:"peakCPUUsage,omitempty"`
	PeakMemoryUsage             float64       `json:"peakMemoryUsage,omitempty"`
	ProductVersion              string        `json:"productVersion,omitempty"`
	Status                      string        `json:"status,omitempty"`
	Type                        string        `json:"type,omitempty"`
	UnavailableTimestamp        int           `json:"unavailableTimestamp,omitempty"`
	Vsds                        []interface{} `json:"vsds,omitempty"`
}

// NewHSC returns a new *HSC
func NewHSC() *HSC {

	return &HSC{}
}

// Identity returns the Identity of the object.
func (o *HSC) Identity() bambou.Identity {

	return HSCIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *HSC) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *HSC) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the HSC from the server
func (o *HSC) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the HSC into the server
func (o *HSC) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the HSC from the server
func (o *HSC) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Alarms retrieves the list of child Alarms of the HSC
func (o *HSC) Alarms(info *bambou.FetchingInfo) (AlarmsList, *bambou.Error) {

	var list AlarmsList
	err := bambou.CurrentSession().FetchChildren(o, AlarmIdentity, &list, info)
	return list, err
}

// CreateAlarm creates a new child Alarm under the HSC
func (o *HSC) CreateAlarm(child *Alarm) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// BGPPeers retrieves the list of child BGPPeers of the HSC
func (o *HSC) BGPPeers(info *bambou.FetchingInfo) (BGPPeersList, *bambou.Error) {

	var list BGPPeersList
	err := bambou.CurrentSession().FetchChildren(o, BGPPeerIdentity, &list, info)
	return list, err
}

// CreateBGPPeer creates a new child BGPPeer under the HSC
func (o *HSC) CreateBGPPeer(child *BGPPeer) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// EventLogs retrieves the list of child EventLogs of the HSC
func (o *HSC) EventLogs(info *bambou.FetchingInfo) (EventLogsList, *bambou.Error) {

	var list EventLogsList
	err := bambou.CurrentSession().FetchChildren(o, EventLogIdentity, &list, info)
	return list, err
}

// CreateEventLog creates a new child EventLog under the HSC
func (o *HSC) CreateEventLog(child *EventLog) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the HSC
func (o *HSC) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the HSC
func (o *HSC) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Jobs retrieves the list of child Jobs of the HSC
func (o *HSC) Jobs(info *bambou.FetchingInfo) (JobsList, *bambou.Error) {

	var list JobsList
	err := bambou.CurrentSession().FetchChildren(o, JobIdentity, &list, info)
	return list, err
}

// CreateJob creates a new child Job under the HSC
func (o *HSC) CreateJob(child *Job) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Metadatas retrieves the list of child Metadatas of the HSC
func (o *HSC) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the HSC
func (o *HSC) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// MonitoringPorts retrieves the list of child MonitoringPorts of the HSC
func (o *HSC) MonitoringPorts(info *bambou.FetchingInfo) (MonitoringPortsList, *bambou.Error) {

	var list MonitoringPortsList
	err := bambou.CurrentSession().FetchChildren(o, MonitoringPortIdentity, &list, info)
	return list, err
}

// CreateMonitoringPort creates a new child MonitoringPort under the HSC
func (o *HSC) CreateMonitoringPort(child *MonitoringPort) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// VRSs retrieves the list of child VRSs of the HSC
func (o *HSC) VRSs(info *bambou.FetchingInfo) (VRSsList, *bambou.Error) {

	var list VRSsList
	err := bambou.CurrentSession().FetchChildren(o, VRSIdentity, &list, info)
	return list, err
}

// CreateVRS creates a new child VRS under the HSC
func (o *HSC) CreateVRS(child *VRS) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}