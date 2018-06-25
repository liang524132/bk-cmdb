/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.
 * Copyright (C) 2017-2018 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package api

import (
	"fmt"
	"net/http"

	"configcenter/src/common"
	"configcenter/src/common/blog"
	frtypes "configcenter/src/common/mapstr"
	"configcenter/src/common/metadata"
	"configcenter/src/scene_server/topo_server/core/types"
)

func init() {
	apiInst.initFuncs = append(apiInst.initFuncs, apiInst.initAssociation)
}

func (cli *topoAPI) initAssociation() {
	cli.actions = append(cli.actions, action{Method: http.MethodPost, Path: "/model/mainline", HandlerFunc: cli.CreateMainLineObject})
	cli.actions = append(cli.actions, action{Method: http.MethodDelete, Path: "/model/mainline/owners/{owner_id}/objectids/{obj_id}", HandlerFunc: cli.DeleteMainLineObject})
	cli.actions = append(cli.actions, action{Method: http.MethodGet, Path: "/model/{owner_id}", HandlerFunc: cli.SearchMainLineOBjectTopo})
	cli.actions = append(cli.actions, action{Method: http.MethodGet, Path: "/model/{owner_id}/{cls_id}/{obj_id}", HandlerFunc: cli.SearchObjectByClassificationID})
	cli.actions = append(cli.actions, action{Method: http.MethodGet, Path: "/inst/{owner_id}/{app_id}", HandlerFunc: cli.SearchBusinessTopo})
	cli.actions = append(cli.actions, action{Method: http.MethodGet, Path: "/inst/child/{owner_id}/{obj_id}/{app_id}/{inst_id}", HandlerFunc: cli.SearchMainLineChildInstTopo})
}

// CreateMainLineObject create a new object in the main line topo
func (cli *topoAPI) CreateMainLineObject(params types.LogicParams, pathParams, queryParams ParamsGetter, data frtypes.MapStr) (interface{}, error) {

	mainLineAssociation := &metadata.Association{}

	_, err := mainLineAssociation.Parse(data)
	if nil != err {
		blog.Errorf("[api-asst] failed to parse the data(%#v), error info is %s", data, err.Error())
	}

	return cli.core.AssociationOperation().CreateMainlineAssociation(params, mainLineAssociation)
}

// DeleteMainLineObject delete a object int the main line topo
func (cli *topoAPI) DeleteMainLineObject(params types.LogicParams, pathParams, queryParams ParamsGetter, data frtypes.MapStr) (interface{}, error) {

	objID := pathParams("obj_id")
	err := cli.core.AssociationOperation().DeleteMainlineAssociaton(params, objID)
	return nil, err
}

// SearchMainLineOBjectTopo search the main line topo
func (cli *topoAPI) SearchMainLineOBjectTopo(params types.LogicParams, pathParams, queryParams ParamsGetter, data frtypes.MapStr) (interface{}, error) {
	fmt.Println("search mainl line object topo")

	bizObj, err := cli.core.ObjectOperation().FindSingleObject(params, common.BKInnerObjIDApp)
	if nil != err {
		blog.Errorf("[api-asst] failed to find the biz object, error info is %s", err.Error())
		return nil, err
	}

	return cli.core.AssociationOperation().SearchMainlineAssociationTopo(params, bizObj)
}

// SearchObjectByClassificationID search the object by classification ID
func (cli *topoAPI) SearchObjectByClassificationID(params types.LogicParams, pathParams, queryParams ParamsGetter, data frtypes.MapStr) (interface{}, error) {
	fmt.Println("search object by classification id")
	return nil, nil
}

// SearchBusinessTopo search the business topo
func (cli *topoAPI) SearchBusinessTopo(params types.LogicParams, pathParams, queryParams ParamsGetter, data frtypes.MapStr) (interface{}, error) {

	paramPath := frtypes.MapStr{}
	paramPath.Set("id", pathParams("app_id"))
	id, err := paramPath.Int64("id")
	if nil != err {
		blog.Errorf("[api-asst] failed to parse the path params id(%s), error info is %s ", pathParams("app_id"), err.Error())
		return nil, err
	}

	return cli.core.AssociationOperation().SearchMainlineAssociationInstTopo(params, id)
}

// SearchMainLineChildInstTopo search the child inst topo by a inst
func (cli *topoAPI) SearchMainLineChildInstTopo(params types.LogicParams, pathParams, queryParams ParamsGetter, data frtypes.MapStr) (interface{}, error) {
	fmt.Println("search main line child inst topo")
	return nil, nil
}