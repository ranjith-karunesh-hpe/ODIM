//(C) Copyright [2020] Hewlett Packard Enterprise Development LP
//
//Licensed under the Apache License, Version 2.0 (the "License"); you may
//not use this file except in compliance with the License. You may obtain
//a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
//WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
//License for the specific language governing permissions and limitations
// under the License.

//Package mgrresponse ...
package mgrresponse

import (
	dmtf "github.com/ODIM-Project/ODIM/lib-dmtf/model"
)

//ManagersCollection for odimra
type ManagersCollection struct {
	OdataContext    string      `json:"@odata.context"`
	Etag            string      `json:"@odata.etag,omitempty"`
	OdataID         string      `json:"@odata.id"`
	OdataType       string      `json:"@odata.type"`
	Description     string      `json:"Description,omitempty"`
	Capabilities    string      `json:"Capabilities,omitempty"`
	Name            string      `json:"Name"`
	Members         []dmtf.Link `json:"Members"`
	MembersCount    int         `json:"Members@odata.count"`
	MembersNextLink string      `json:"Members@odata.nextLink,omitempty"`
	Oem             *dmtf.Oem   `json:"Oem,omitempty"`
}
