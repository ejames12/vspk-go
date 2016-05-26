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

// AutoDiscoverHypervisorFromDatacenterIdentity represents the Identity of the object
var AutoDiscoverHypervisorFromDatacenterIdentity = bambou.Identity{
	Name:     "autodiscoveredcomputeresource",
	Category: "autodiscoveredcomputeresources",
}

// AutoDiscoverHypervisorFromDatacentersList represents a list of AutoDiscoverHypervisorFromDatacenters
type AutoDiscoverHypervisorFromDatacentersList []*AutoDiscoverHypervisorFromDatacenter

// AutoDiscoverHypervisorFromDatacentersAncestor is the interface of an ancestor of a AutoDiscoverHypervisorFromDatacenter must implement.
type AutoDiscoverHypervisorFromDatacentersAncestor interface {
	AutoDiscoverHypervisorFromDatacenters(*bambou.FetchingInfo) (AutoDiscoverHypervisorFromDatacentersList, *bambou.Error)
	CreateAutoDiscoverHypervisorFromDatacenters(*AutoDiscoverHypervisorFromDatacenter) *bambou.Error
}

// AutoDiscoverHypervisorFromDatacenter represents the model of a autodiscoveredcomputeresource
type AutoDiscoverHypervisorFromDatacenter struct {
	ID                       string        `json:"ID,omitempty"`
	ParentID                 string        `json:"parentID,omitempty"`
	ParentType               string        `json:"parentType,omitempty"`
	Owner                    string        `json:"owner,omitempty"`
	NetworkList              []interface{} `json:"networkList,omitempty"`
	AssocVCenterDataCenterId string        `json:"assocVCenterDataCenterId,omitempty"`
	HypervisorIP             string        `json:"hypervisorIP,omitempty"`
}

// NewAutoDiscoverHypervisorFromDatacenter returns a new *AutoDiscoverHypervisorFromDatacenter
func NewAutoDiscoverHypervisorFromDatacenter() *AutoDiscoverHypervisorFromDatacenter {

	return &AutoDiscoverHypervisorFromDatacenter{}
}

// Identity returns the Identity of the object.
func (o *AutoDiscoverHypervisorFromDatacenter) Identity() bambou.Identity {

	return AutoDiscoverHypervisorFromDatacenterIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *AutoDiscoverHypervisorFromDatacenter) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *AutoDiscoverHypervisorFromDatacenter) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the AutoDiscoverHypervisorFromDatacenter from the server
func (o *AutoDiscoverHypervisorFromDatacenter) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the AutoDiscoverHypervisorFromDatacenter into the server
func (o *AutoDiscoverHypervisorFromDatacenter) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the AutoDiscoverHypervisorFromDatacenter from the server
func (o *AutoDiscoverHypervisorFromDatacenter) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}
