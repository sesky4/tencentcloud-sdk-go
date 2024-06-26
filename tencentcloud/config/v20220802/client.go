// Copyright (c) 2017-2018 THL A29 Limited, a Tencent company. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v20220802

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2022-08-02"

type Client struct {
    common.Client
}

// Deprecated
func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
    cpf := profile.NewClientProfile()
    client = &Client{}
    client.Init(region).WithSecretId(secretId, secretKey).WithProfile(cpf)
    return
}

func NewClient(credential common.CredentialIface, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewListAggregateConfigRulesRequest() (request *ListAggregateConfigRulesRequest) {
    request = &ListAggregateConfigRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "ListAggregateConfigRules")
    
    
    return
}

func NewListAggregateConfigRulesResponse() (response *ListAggregateConfigRulesResponse) {
    response = &ListAggregateConfigRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListAggregateConfigRules
// 账号组获取规则列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ACCOUNTGROUPISNOTEXIST = "ResourceNotFound.AccountGroupIsNotExist"
func (c *Client) ListAggregateConfigRules(request *ListAggregateConfigRulesRequest) (response *ListAggregateConfigRulesResponse, err error) {
    return c.ListAggregateConfigRulesWithContext(context.Background(), request)
}

// ListAggregateConfigRules
// 账号组获取规则列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ACCOUNTGROUPISNOTEXIST = "ResourceNotFound.AccountGroupIsNotExist"
func (c *Client) ListAggregateConfigRulesWithContext(ctx context.Context, request *ListAggregateConfigRulesRequest) (response *ListAggregateConfigRulesResponse, err error) {
    if request == nil {
        request = NewListAggregateConfigRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListAggregateConfigRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewListAggregateConfigRulesResponse()
    err = c.Send(request, response)
    return
}

func NewListConfigRulesRequest() (request *ListConfigRulesRequest) {
    request = &ListConfigRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "ListConfigRules")
    
    
    return
}

func NewListConfigRulesResponse() (response *ListConfigRulesResponse) {
    response = &ListConfigRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListConfigRules
// 获取规则列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ListConfigRules(request *ListConfigRulesRequest) (response *ListConfigRulesResponse, err error) {
    return c.ListConfigRulesWithContext(context.Background(), request)
}

// ListConfigRules
// 获取规则列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ListConfigRulesWithContext(ctx context.Context, request *ListConfigRulesRequest) (response *ListConfigRulesResponse, err error) {
    if request == nil {
        request = NewListConfigRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListConfigRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewListConfigRulesResponse()
    err = c.Send(request, response)
    return
}

func NewPutEvaluationsRequest() (request *PutEvaluationsRequest) {
    request = &PutEvaluationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "PutEvaluations")
    
    
    return
}

func NewPutEvaluationsResponse() (response *PutEvaluationsResponse) {
    response = &PutEvaluationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// PutEvaluations
// 上报自定义规则评估结果
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ACCOUNTGROUPISNOTEXIST = "ResourceNotFound.AccountGroupIsNotExist"
//  RESOURCENOTFOUND_RESOURCENOTEXIST = "ResourceNotFound.ResourceNotExist"
//  RESOURCENOTFOUND_RULEISNOTEXIST = "ResourceNotFound.RuleIsNotExist"
func (c *Client) PutEvaluations(request *PutEvaluationsRequest) (response *PutEvaluationsResponse, err error) {
    return c.PutEvaluationsWithContext(context.Background(), request)
}

// PutEvaluations
// 上报自定义规则评估结果
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ACCOUNTGROUPISNOTEXIST = "ResourceNotFound.AccountGroupIsNotExist"
//  RESOURCENOTFOUND_RESOURCENOTEXIST = "ResourceNotFound.ResourceNotExist"
//  RESOURCENOTFOUND_RULEISNOTEXIST = "ResourceNotFound.RuleIsNotExist"
func (c *Client) PutEvaluationsWithContext(ctx context.Context, request *PutEvaluationsRequest) (response *PutEvaluationsResponse, err error) {
    if request == nil {
        request = NewPutEvaluationsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("PutEvaluations require credential")
    }

    request.SetContext(ctx)
    
    response = NewPutEvaluationsResponse()
    err = c.Send(request, response)
    return
}
