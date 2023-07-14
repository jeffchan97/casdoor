// Copyright 2022 The Casdoor Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package controllers

import (
	"encoding/json"
	"github.com/casdoor/casdoor/object"
)

type UrlActionAuthzParams struct {
	object.PermissionRule
	Adapters []string `json:"adapters"`
}

func (c *ApiController) UrlActionAuthz() {
	//var permissionRule object.PermissionRule
	var urlActionAuthzParams UrlActionAuthzParams
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &urlActionAuthzParams)

	if err != nil {
		c.ResponseError(err.Error())
		return
	}

	c.Data["json"] = object.UrlActionAuthz(&urlActionAuthzParams.PermissionRule, urlActionAuthzParams.Adapters)
	c.ServeJSON()
}

func (c *ApiController) Enforce() {
	var permissionRule object.PermissionRule
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &permissionRule)
	if err != nil {
		c.ResponseError(err.Error())
		return
	}

	c.Data["json"] = object.Enforce(&permissionRule)
	c.ServeJSON()
}

func (c *ApiController) BatchEnforce() {
	var permissionRules []object.PermissionRule
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &permissionRules)
	if err != nil {
		c.ResponseError(err.Error())
		return
	}

	c.Data["json"] = object.BatchEnforce(permissionRules)
	c.ServeJSON()
}

func (c *ApiController) GetAllObjects() {
	userId := c.GetSessionUsername()
	if userId == "" {
		c.ResponseError(c.T("general:Please login first"))
		return
	}

	c.Data["json"] = object.GetAllObjects(userId)
	c.ServeJSON()
}

func (c *ApiController) GetAllActions() {
	userId := c.GetSessionUsername()
	if userId == "" {
		c.ResponseError(c.T("general:Please login first"))
		return
	}

	c.Data["json"] = object.GetAllActions(userId)
	c.ServeJSON()
}

func (c *ApiController) GetAllRoles() {
	userId := c.GetSessionUsername()
	if userId == "" {
		c.ResponseError(c.T("general:Please login first"))
		return
	}

	c.Data["json"] = object.GetAllRoles(userId)
	c.ServeJSON()
}
