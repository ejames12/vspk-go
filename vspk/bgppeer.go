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

// BGPPeerIdentity represents the Identity of the object
var BGPPeerIdentity = bambou.Identity{
	Name:     "bgppeer",
	Category: "bgppeers",
}

// BGPPeersList represents a list of BGPPeers
type BGPPeersList []*BGPPeer

// BGPPeersAncestor is the interface of an ancestor of a BGPPeer must implement.
type BGPPeersAncestor interface {
	BGPPeers(*bambou.FetchingInfo) (BGPPeersList, *bambou.Error)
	CreateBGPPeers(*BGPPeer) *bambou.Error
}

// BGPPeer represents the model of a bgppeer
type BGPPeer struct {
	ID              string `json:"ID,omitempty"`
	ParentID        string `json:"parentID,omitempty"`
	ParentType      string `json:"parentType,omitempty"`
	Owner           string `json:"owner,omitempty"`
	LastStateChange int    `json:"lastStateChange,omitempty"`
	Address         string `json:"address,omitempty"`
	EntityScope     string `json:"entityScope,omitempty"`
	Status          string `json:"status,omitempty"`
	ExternalID      string `json:"externalID,omitempty"`
}

// NewBGPPeer returns a new *BGPPeer
func NewBGPPeer() *BGPPeer {

	return &BGPPeer{}
}

// Identity returns the Identity of the object.
func (o *BGPPeer) Identity() bambou.Identity {

	return BGPPeerIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *BGPPeer) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *BGPPeer) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the BGPPeer from the server
func (o *BGPPeer) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the BGPPeer into the server
func (o *BGPPeer) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the BGPPeer from the server
func (o *BGPPeer) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the BGPPeer
func (o *BGPPeer) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the BGPPeer
func (o *BGPPeer) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the BGPPeer
func (o *BGPPeer) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the BGPPeer
func (o *BGPPeer) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
