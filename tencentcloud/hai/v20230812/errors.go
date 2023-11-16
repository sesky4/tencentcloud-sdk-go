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

package v20230812

const (
	// 此产品的特有错误码

	// 欠费账户不能创建实例
	FAILEDOPERATION_ARREARSACCOUNTCANNOTRUNINSTANCES = "FailedOperation.ArrearsAccountCannotRunInstances"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 两个指定的入参每次只能使用其中一个
	INVALIDPARAMETER_ATMOSTONE = "InvalidParameter.AtMostOne"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 指定的应用不存在
	INVALIDPARAMETERVALUE_APPLICATIONIDNOTFOUND = "InvalidParameterValue.ApplicationIdNotFound"

	// 算力套餐类型不存在
	INVALIDPARAMETERVALUE_BUNDLETYPENOTFOUND = "InvalidParameterValue.BundleTypeNotFound"

	// 列表入参中存在重复值
	INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"

	// 指定的实例不存在
	INVALIDPARAMETERVALUE_INSTANCEIDNOTFOUND = "InvalidParameterValue.InstanceIdNotFound"

	// 实例名称过长
	INVALIDPARAMETERVALUE_INSTANCENAMETOOLONG = "InvalidParameterValue.InstanceNameTooLong"

	// 应用ID格式非法
	INVALIDPARAMETERVALUE_INVALIDAPPLICATIONIDMALFORMED = "InvalidParameterValue.InvalidApplicationIdMalformed"

	// 每次购买的实例数目不在合理范围内
	INVALIDPARAMETERVALUE_INVALIDINSTANCECOUNT = "InvalidParameterValue.InvalidInstanceCount"

	// 实例ID格式非法
	INVALIDPARAMETERVALUE_INVALIDINSTANCEIDMALFORMED = "InvalidParameterValue.InvalidInstanceIdMalformed"

	// 场景ID格式非法
	INVALIDPARAMETERVALUE_INVALIDSCENEIDMALFORMED = "InvalidParameterValue.InvalidSceneIdMalformed"

	// 指定实例有正在执行的操作，不能执行新的操作
	OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"

	// 指定的算力套餐库存不足
	RESOURCEINSUFFICIENT_BUNDLEINVENTORYSHORTAGE = "ResourceInsufficient.BundleInventoryShortage"

	// 用户账号的网络类型是传统型，不允许使用HAI
	UNAUTHORIZEDOPERATION_UNAUTHORIZEDNETWORKUSER = "UnauthorizedOperation.UnauthorizedNetworkUser"

	// 客户未授权使用本产品
	UNAUTHORIZEDOPERATION_UNAUTHORIZEDUSER = "UnauthorizedOperation.UnauthorizedUser"
)