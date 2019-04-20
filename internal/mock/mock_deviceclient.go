// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2018 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package mock

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/edgexfoundry/go-mod-core-contracts/models"
)

const (
	InvalidDeviceId                = "1ef435eb-5060-49b0-8d55-8d4e43239800"
	RandomIntegerGeneratorDeviceId = "79ef304e-3b1e-4765-b509-798109c2467d"
	RandomFloatGeneratorDeviceId   = "9a2b3ef2-6aed-41db-b868-79196c2a8b0a"
)

var (
	ValidDeviceRandomBoolGenerator            = models.Device{}
	ValidDeviceRandomIntegerGenerator         = models.Device{}
	ValidDeviceRandomUnsignedIntegerGenerator = models.Device{}
	ValidDeviceRandomFloatGenerator           = models.Device{}
	DuplicateDeviceRandomFloatGenerator       = models.Device{}
	NewValidDevice                            = models.Device{}
)

type DeviceClientMock struct{}

func (dc *DeviceClientMock) Add(dev *models.Device, ctx context.Context) (string, error) {
	panic("implement me")
}

func (dc *DeviceClientMock) Delete(id string, ctx context.Context) error {
	panic("implement me")
}

func (dc *DeviceClientMock) DeleteByName(name string, ctx context.Context) error {
	panic("implement me")
}

func (dc *DeviceClientMock) CheckForDevice(token string, ctx context.Context) (models.Device, error) {
	panic("implement me")
}

func (dc *DeviceClientMock) Device(id string, ctx context.Context) (models.Device, error) {
	if id == InvalidDeviceId {
		return models.Device{}, fmt.Errorf("invalid id")
	}
	return models.Device{}, nil
}

func (dc *DeviceClientMock) DeviceForName(name string, ctx context.Context) (models.Device, error) {
	var device = models.Device{Id: "5b977c62f37ba10e36673802", Name: name}
	var err error = nil
	if name == "" {
		err = errors.New("Item not found")
	}

	return device, err
}

func (dc *DeviceClientMock) Devices(ctx context.Context) ([]models.Device, error) {
	panic("implement me")
}

func (dc *DeviceClientMock) DevicesByLabel(label string, ctx context.Context) ([]models.Device, error) {
	panic("implement me")
}

func (dc *DeviceClientMock) DevicesForProfile(profileid string, ctx context.Context) ([]models.Device, error) {
	panic("implement me")
}

func (dc *DeviceClientMock) DevicesForProfileByName(profileName string, ctx context.Context) ([]models.Device, error) {
	panic("implement me")
}

func (dc *DeviceClientMock) DevicesForService(serviceid string, ctx context.Context) ([]models.Device, error) {
	panic("implement me")
}

func (dc *DeviceClientMock) DevicesForServiceByName(serviceName string, ctx context.Context) ([]models.Device, error) {
	populateDeviceMock()
	return []models.Device{
		ValidDeviceRandomBoolGenerator,
		ValidDeviceRandomIntegerGenerator,
		ValidDeviceRandomUnsignedIntegerGenerator,
		ValidDeviceRandomFloatGenerator,
	}, nil
}

func (dc *DeviceClientMock) Update(dev models.Device, ctx context.Context) error {
	return nil
}

func (dc *DeviceClientMock) UpdateAdminState(id string, adminState string, ctx context.Context) error {
	return nil
}

func (dc *DeviceClientMock) UpdateAdminStateByName(name string, adminState string, ctx context.Context) error {
	return nil
}

func (dc *DeviceClientMock) UpdateLastConnected(id string, time int64, ctx context.Context) error {
	return nil
}

func (dc *DeviceClientMock) UpdateLastConnectedByName(name string, time int64, ctx context.Context) error {
	return nil
}

func (dc *DeviceClientMock) UpdateLastReported(id string, time int64, ctx context.Context) error {
	return nil
}

func (dc *DeviceClientMock) UpdateLastReportedByName(name string, time int64, ctx context.Context) error {
	return nil
}

func (dc *DeviceClientMock) UpdateOpState(id string, opState string, ctx context.Context) error {
	return nil
}

func (dc *DeviceClientMock) UpdateOpStateByName(name string, opState string, ctx context.Context) error {
	return nil
}

func populateDeviceMock() {
	deviceDataRandomBoolGenerator := `{"description":"","id":"4a37f97f-4fdb-4082-b8e1-872fa1715a78","name":"Random-Boolean-Generator01","adminState":"UNLOCKED","operatingState":"ENABLED","protocols":{"other":{"Address":"device-virtual-bool-01","Port":"300"}},"labels":["device-virtual-example"],"service":{"created":1553712641011,"modified":1553712641011,"origin":1553712641003,"description":"","id":"c58ed126-c713-4cf8-9750-b38c1fd9b43e","name":"device-virtual","lastConnected":0,"lastReported":0,"operatingState":"ENABLED","labels":[],"addressable":{"created":1553712640993,"modified":1553712640993,"origin":1553712640987,"id":"6ec2a3bb-d5c6-49a8-92a5-e41d3fe6ceb7","name":"device-virtual","protocol":"HTTP","method":"POST","address":"192.168.56.101","port":49988,"path":"/api/v1/callback","baseURL":"HTTP://192.168.56.101:49988","url":"HTTP://192.168.56.101:49988/api/v1/callback"},"adminState":"UNLOCKED"},"profile":{"created":1553712641093,"modified":1553712641093,"description":"Example of Device-Virtual","id":"79e5c532-15d2-4267-a7b4-574490eae8a6","name":"Random-Boolean-Generator","manufacturer":"IOTech","model":"Device-Virtual-01","labels":["device-virtual-example"],"deviceResources":[{"description":"used to decide whether to re-generate a random value","name":"Enable_Randomization","properties":{"value":{"type":"Bool","readWrite":"W","defaultValue":"true"},"units":{"type":"String","readWrite":"R","defaultValue":"Random"}}},{"description":"Generate random boolean value","name":"RandomValue_Bool","properties":{"value":{"type":"Bool","readWrite":"R","defaultValue":"true"},"units":{"type":"String","readWrite":"R","defaultValue":"random bool value"}}}],"deviceCommands":[{"name":"RandomValue_Bool","get":[{"operation":"get","object":"RandomValue_Bool","parameter":"RandomValue_Bool"}],"set":[{"operation":"set","object":"Enable_Randomization","parameter":"Enable_Randomization","resource":"RandomValue_Bool"},{"operation":"set","object":"RandomValue_Bool","parameter":"RandomValue_Bool","resource":"RandomValue_Bool"}]}],"coreCommands":[{"created":1553712641071,"modified":1553712641071,"id":"c6d15b2a-2f7f-43a8-af52-0e896c25b662","name":"RandomValue_Bool","get":{"path":"/api/v1/device/{deviceId}/RandomValue_Bool","responses":[{"code":"503","description":"service unavailable"}]},"put":{"path":"/api/v1/device/{deviceId}/RandomValue_Bool","parameterNames":["RandomValue_Bool","Enable_Randomization"]}}]}}`
	deviceDataRandomIntegerGenerator := `{"description":"","id":"` + RandomIntegerGeneratorDeviceId + `","name":"Random-Integer-Generator01","adminState":"UNLOCKED","operatingState":"ENABLED","protocols":{"other":{"Address":"device-virtual-int-01","Protocol":"300"}},"labels":["device-virtual-example"],"service":{"created":1553721583438,"modified":1553721583438,"origin":1553721583423,"description":"","id":"88e4ec9a-5f7a-425d-876a-a246e44a14d6","name":"device-virtual","lastConnected":0,"lastReported":0,"operatingState":"ENABLED","labels":[],"addressable":{"created":1553721583413,"modified":1553721583413,"origin":1553721583413,"id":"77735758-379d-481d-9887-f07001c05d65","name":"device-virtual","protocol":"HTTP","method":"POST","address":"192.168.56.101","port":49988,"path":"/api/v1/callback","baseURL":"HTTP://192.168.56.101:49988","url":"HTTP://192.168.56.101:49988/api/v1/callback"},"adminState":"UNLOCKED"},"profile":{"created":1553721583525,"modified":1553721583525,"description":"Example of Device-Virtual","id":"1da96459-4fef-4860-9980-a653c28c5ff6","name":"Random-Integer-Generator","manufacturer":"IOTech","model":"Device-Virtual-01","labels":["device-virtual-example"],"deviceResources":[{"description":"used to decide whether to re-generate a random value","name":"Enable_Randomization","properties":{"value":{"type":"Bool","readWrite":"W","defaultValue":"true"},"units":{"type":"String","readWrite":"R","defaultValue":"Random"}}},{"description":"Generate random int8 value","name":"RandomValue_Int8","properties":{"value":{"type":"Int8","readWrite":"R","defaultValue":"0"},"units":{"type":"String","readWrite":"R","defaultValue":"random int8 value"}}},{"description":"Generate random int16 value","name":"RandomValue_Int16","properties":{"value":{"type":"Int16","readWrite":"R","defaultValue":"0"},"units":{"type":"String","readWrite":"R","defaultValue":"random int16 value"}}},{"description":"Generate random int32 value","name":"RandomValue_Int32","properties":{"value":{"type":"Int32","readWrite":"R","defaultValue":"0"},"units":{"type":"String","readWrite":"R","defaultValue":"random int32 value"}}},{"description":"Generate random int64 value","name":"RandomValue_Int64","properties":{"value":{"type":"Int64","readWrite":"R","defaultValue":"0"},"units":{"type":"String","readWrite":"R","defaultValue":"random int64 value"}}},{"description":"ResourceTestTransform_Pass","name":"ResourceTestTransform_Pass","properties":{"value":{"type":"Int8","readWrite":"RW","defaultValue":"0","offset":"1.0"},"units":{"type":"String","readWrite":"R"}}},{"description":"ResourceTestTransform_Fail","name":"ResourceTestTransform_Fail","properties":{"value":{"type":"Int8","readWrite":"RW","defaultValue":"0","offset":"error"},"units":{"type":"String","readWrite":"R"}}},{"description":"ResourceTestAssertion_Pass","name":"ResourceTestAssertion_Pass","properties":{"value":{"type":"Int8","readWrite":"RW","defaultValue":"0","assertion":"123"},"units":{"type":"String","readWrite":"R"}}},{"description":"ResourceTestAssertion_Fail","name":"ResourceTestAssertion_Fail","properties":{"value":{"type":"Int8","readWrite":"RW","defaultValue":"0","assertion":"12"},"units":{"type":"String","readWrite":"R"}}},{"description":"NoDeviceResourceForResult","name":"NoDeviceResourceForResult","properties":{"value":{"type":"String","readWrite":"RW"},"units":{"type":"String","readWrite":"R"}}},{"description":"Error","name":"Error","properties":{"value":{"type":"String","readWrite":"RW"},"units":{"type":"String","readWrite":"R"}}}],"deviceCommands":[{"name":"RandomValue_Int8","get":[{"operation":"get","object":"RandomValue_Int8","parameter":"RandomValue_Int8"}],"set":[{"operation":"set","object":"RandomValue_Int8","parameter":"RandomValue_Int8","resource":"RandomValue_Int8"},{"operation":"set","object":"Enable_Randomization","parameter":"Enable_Randomization","resource":"RandomValue_Int8"}]},{"name":"RandomValue_Int16","get":[{"operation":"get","object":"RandomValue_Int16","parameter":"RandomValue_Int16"}],"set":[{"operation":"set","object":"RandomValue_Int16","parameter":"RandomValue_Int16","resource":"RandomValue_Int16"},{"operation":"set","object":"Enable_Randomization","parameter":"Enable_Randomization","resource":"RandomValue_Int16"}]},{"name":"RandomValue_Int32","get":[{"operation":"get","object":"RandomValue_Int32","parameter":"RandomValue_Int32"}],"set":[{"operation":"set","object":"RandomValue_Int32","parameter":"RandomValue_Int32","resource":"RandomValue_Int32"},{"operation":"set","object":"Enable_Randomization","parameter":"Enable_Randomization","resource":"RandomValue_Int32"}]},{"name":"RandomValue_Int64","get":[{"operation":"get","object":"RandomValue_Int64","parameter":"RandomValue_Int64"}],"set":[{"operation":"set","object":"RandomValue_Int64","parameter":"RandomValue_Int64","resource":"RandomValue_Int64"},{"operation":"set","object":"Enable_Randomization","parameter":"Enable_Randomization","resource":"RandomValue_Int64"}]},{"name":"ResourceTestTransform_Fail","get":[{"operation":"get","object":"ResourceTestTransform_Fail","parameter":"ResourceTestTransform_Fail"}],"set":[{"operation":"set","object":"ResourceTestTransform_Fail","parameter":"ResourceTestTransform_Fail"}]},{"name":"ResourceTestAssertion_Pass","get":[{"operation":"get","object":"ResourceTestAssertion_Pass","parameter":"ResourceTestAssertion_Pass"}],"set":[{"operation":"set","object":"ResourceTestAssertion_Pass","parameter":"ResourceTestAssertion_Pass"}]},{"name":"ResourceTestAssertion_Fail","get":[{"operation":"get","object":"ResourceTestAssertion_Fail","parameter":"ResourceTestAssertion_Fail"}],"set":[{"operation":"set","object":"ResourceTestAssertion_Fail","parameter":"ResourceTestAssertion_Fail"}]},{"name":"ResourceTestMapping_Pass","get":[{"operation":"get","object":"RandomValue_Int8","parameter":"RandomValue_Int8","mappings":{"123":"Pass"}}],"set":[{"operation":"get","object":"RandomValue_Int8","parameter":"RandomValue_Int8","mappings":{"Pass":"123"}}]},{"name":"ResourceTestMapping_Fail","get":[{"operation":"get","object":"RandomValue_Int8","parameter":"RandomValue_Int8","mappings":{"12":"Pass"}}],"set":[{"operation":"get","object":"RandomValue_Int8","parameter":"RandomValue_Int8","mappings":{"Pass":"12"}}]},{"name":"NoDeviceResourceForOperation","get":[{"operation":"get","object":"ResourceNotFound","parameter":"Error"}],"set":[{"operation":"set","object":"ResourceNotFound","parameter":"Error"}]},{"name":"NoDeviceResourceForResult","get":[{"operation":"get","object":"NoDeviceResourceForResult","parameter":"NoDeviceResourceForResult"}]},{"name":"Error","get":[{"index":"1","operation":"get","object":"Error","parameter":"Error"},{"index":"2","operation":"get","object":"Error","parameter":"Error"}],"set":[{"index":"1","operation":"set","object":"Error","parameter":"Error"},{"index":"2","operation":"set","object":"Error","parameter":"Error"}]}],"coreCommands":[{"created":1553721583522,"modified":1553721583522,"id":"6b2c4f1d-9537-4cf6-a25b-9f0bf7761be6","name":"RandomValue_Int8","get":{"path":"/api/v1/device/{deviceId}/RandomValue_Int8","responses":[{"code":"503","description":"service unavailable"}]},"put":{"path":"/api/v1/device/{deviceId}/RandomValue_Int8","parameterNames":["RandomValue_Int8","Enable_Randomization"]}},{"created":1553721583523,"modified":1553721583523,"id":"b5b3d990-7e86-4222-a92a-cacde97a23a2","name":"RandomValue_Int16","get":{"path":"/api/v1/device/{deviceId}/RandomValue_Int16","responses":[{"code":"503","description":"service unavailable"}]},"put":{"path":"/api/v1/device/{deviceId}/RandomValue_Int16","parameterNames":["RandomValue_Int16","Enable_Randomization"]}},{"created":1553721583523,"modified":1553721583523,"id":"d0ce104e-0582-4651-9c8a-86e51c4b1180","name":"RandomValue_Int32","get":{"path":"/api/v1/device/{deviceId}/RandomValue_Int32","responses":[{"code":"503","description":"service unavailable"}]},"put":{"path":"/api/v1/device/{deviceId}/RandomValue_Int32","parameterNames":["RandomValue_Int32","Enable_Randomization"]}},{"created":1553721583523,"modified":1553721583523,"id":"8105e902-194a-49b1-a71f-296a0057dfff","name":"RandomValue_Int64","get":{"path":"/api/v1/device/{deviceId}/RandomValue_Int64","responses":[{"code":"503","description":"service unavailable"}]},"put":{"path":"/api/v1/device/{deviceId}/RandomValue_Int64","parameterNames":["RandomValue_Int64","Enable_Randomization"]}},{"created":1553721583523,"modified":1553721583523,"id":"86f7074f-4350-4b2c-913b-69288abeec79","name":"ResourceTestTransform_Fail","get":{"path":"/api/v1/device/{deviceId}/ResourceTestTransform_Fail","responses":[{"code":"503","description":"service unavailable"}]},"put":{"path":"/api/v1/device/{deviceId}/ResourceTestTransform_Fail","parameterNames":["ResourceTestTransform_Fail"]}},{"created":1553721583523,"modified":1553721583523,"id":"02819de2-9011-4991-9ce6-f4f7f1990a79","name":"ResourceTestAssertion_Fail","get":{"path":"/api/v1/device/{deviceId}/ResourceTestAssertion_Fail","responses":[{"code":"503","description":"service unavailable"}]},"put":{"path":"/api/v1/device/{deviceId}/ResourceTestAssertion_Fail","parameterNames":["ResourceTestAssertion_Fail"]}},{"created":1553721583523,"modified":1553721583523,"id":"cc625b93-7655-4e5b-920c-77387a4ec3a2","name":"ResourceTestMapping_Pass","get":{"path":"/api/v1/device/{deviceId}/ResourceTestMapping_Pass","responses":[{"code":"503","description":"service unavailable"}]},"put":{"path":"/api/v1/device/{deviceId}/ResourceTestMapping_Pass","parameterNames":["RandomValue_Int8"]}},{"created":1553721583523,"modified":1553721583523,"id":"b0b4d373-430d-46f7-88cc-cd392aaa5c90","name":"ResourceTestMapping_Fail","get":{"path":"/api/v1/device/{deviceId}/ResourceTestMapping_Fail","responses":[{"code":"503","description":"service unavailable"}]},"put":{"path":"/api/v1/device/{deviceId}/ResourceTestMapping_Fail","parameterNames":["RandomValue_Int8"]}},{"created":1553721583523,"modified":1553721583523,"id":"ffa2c711-49ea-4e8a-84ad-0e00ef69de3f","name":"NoDeviceResourceForOperation","get":{"path":"/api/v1/device/{deviceId}/NoDeviceResourceForOperation","responses":[{"code":"503","description":"service unavailable"}]},"put":{"path":"/api/v1/device/{deviceId}/NoDeviceResourceForOperation","parameterNames":["Error"]}},{"created":1553721583524,"modified":1553721583524,"id":"32841000-5fa1-4a27-975f-18456f564149","name":"NoDeviceResourceForResult","get":{"path":"/api/v1/device/{deviceId}/NoDeviceResourceForResult","responses":[{"code":"503","description":"service unavailable"}]},"put":{"path":"/api/v1/device/{deviceId}/NoDeviceResourceForResult","parameterNames":["NoDeviceResourceForResult"]}},{"created":1553721583524,"modified":1553721583524,"id":"2f2cc0cc-7cd1-4a18-9202-df0d049b94de","name":"Error","get":{"path":"/api/v1/device/{deviceId}/Error","responses":[{"code":"503","description":"service unavailable"}]},"put":{"path":"/api/v1/device/{deviceId}/Error","parameterNames":["Error","Error"]}}]}}`
	deviceDataRandomUnsignedIntegerGenerator := `{"description":"","id":"8243da6a-d955-41ab-a612-8200b7ef0c26","name":"Random-UnsignedInteger-Generator01","adminState":"UNLOCKED","operatingState":"ENABLED","protocols":{"other":{"Address":"device-virtual-uint-01","Protocol":"300"}},"labels":["device-virtual-example"],"service":{"created":1553712641011,"modified":1553712641011,"origin":1553712641003,"description":"","id":"c58ed126-c713-4cf8-9750-b38c1fd9b43e","name":"device-virtual","lastConnected":0,"lastReported":0,"operatingState":"ENABLED","labels":[],"addressable":{"created":1553712640993,"modified":1553712640993,"origin":1553712640987,"id":"6ec2a3bb-d5c6-49a8-92a5-e41d3fe6ceb7","name":"device-virtual","protocol":"HTTP","method":"POST","address":"192.168.56.101","port":49988,"path":"/api/v1/callback","baseURL":"HTTP://192.168.56.101:49988","url":"HTTP://192.168.56.101:49988/api/v1/callback"},"adminState":"UNLOCKED"},"profile":{"created":1553712641135,"modified":1553712641135,"description":"Example of Device-Virtual","id":"761a27de-1cdd-4f3a-a4df-55da32239827","name":"Random-UnsignedInteger-Generator","manufacturer":"IOTech","model":"Device-Virtual-01","labels":["device-virtual-example"],"deviceResources":[{"description":"used to decide whether to re-generate a random value","name":"Enable_Randomization","properties":{"value":{"type":"Bool","readWrite":"W","defaultValue":"true"},"units":{"type":"String","readWrite":"R","defaultValue":"Random"}}},{"description":"Generate random uint8 value","name":"RandomValue_Uint8","properties":{"value":{"type":"Uint8","readWrite":"R","defaultValue":"0"},"units":{"type":"String","readWrite":"R","defaultValue":"random uint8 value"}}},{"description":"Generate random uint16 value","name":"RandomValue_Uint16","properties":{"value":{"type":"Uint16","readWrite":"R","defaultValue":"0"},"units":{"type":"String","readWrite":"R","defaultValue":"random uint16 value"}}},{"description":"Generate random uint32 value","name":"RandomValue_Uint32","properties":{"value":{"type":"Uint32","readWrite":"R","defaultValue":"0"},"units":{"type":"String","readWrite":"R","defaultValue":"random uint32 value"}}},{"description":"Generate random uint64 value","name":"RandomValue_Uint64","properties":{"value":{"type":"Uint64","readWrite":"R","defaultValue":"0"},"units":{"type":"String","readWrite":"R","defaultValue":"random uint64 value"}}}],"deviceCommands":[{"name":"RandomValue_Uint8","get":[{"operation":"get","object":"RandomValue_Uint8","parameter":"RandomValue_Uint8"}],"set":[{"operation":"set","object":"RandomValue_Uint8","parameter":"RandomValue_Uint8","resource":"RandomValue_Uint8"},{"operation":"set","object":"Enable_Randomization","parameter":"Enable_Randomization","resource":"RandomValue_Uint8"}]},{"name":"RandomValue_Uint16","get":[{"operation":"get","object":"RandomValue_Uint16","parameter":"RandomValue_Uint16"}],"set":[{"operation":"set","object":"RandomValue_Uint16","parameter":"RandomValue_Uint16","resource":"RandomValue_Uint16"},{"operation":"set","object":"Enable_Randomization","parameter":"Enable_Randomization","resource":"RandomValue_Uint16"}]},{"name":"RandomValue_Uint32","get":[{"operation":"get","object":"RandomValue_Uint32","parameter":"RandomValue_Uint32"}],"set":[{"operation":"set","object":"RandomValue_Uint32","parameter":"RandomValue_Uint32","resource":"RandomValue_Uint32"},{"operation":"set","object":"Enable_Randomization","parameter":"Enable_Randomization","resource":"RandomValue_Uint32"}]},{"name":"RandomValue_Uint64","get":[{"operation":"get","object":"RandomValue_Uint64","parameter":"RandomValue_Uint64"}],"set":[{"operation":"set","object":"RandomValue_Uint64","parameter":"RandomValue_Uint64","resource":"RandomValue_Uint64"},{"operation":"set","object":"Enable_Randomization","parameter":"Enable_Randomization","resource":"RandomValue_Uint64"}]}],"coreCommands":[{"created":1553712641133,"modified":1553712641133,"id":"f388b262-9514-446c-81a4-11951862f9e4","name":"RandomValue_Uint8","get":{"path":"/api/v1/device/{deviceId}/RandomValue_Uint8","responses":[{"code":"503","description":"service unavailable"}]},"put":{"path":"/api/v1/device/{deviceId}/RandomValue_Uint8","parameterNames":["RandomValue_Uint8","Enable_Randomization"]}},{"created":1553712641134,"modified":1553712641134,"id":"4e438f91-fc47-45bf-aefd-4c6444fd60aa","name":"RandomValue_Uint16","get":{"path":"/api/v1/device/{deviceId}/RandomValue_Uint16","responses":[{"code":"503","description":"service unavailable"}]},"put":{"path":"/api/v1/device/{deviceId}/RandomValue_Uint16","parameterNames":["RandomValue_Uint16","Enable_Randomization"]}},{"created":1553712641134,"modified":1553712641134,"id":"93976c0d-5f38-4690-88f7-343de7cacb89","name":"RandomValue_Uint32","get":{"path":"/api/v1/device/{deviceId}/RandomValue_Uint32","responses":[{"code":"503","description":"service unavailable"}]},"put":{"path":"/api/v1/device/{deviceId}/RandomValue_Uint32","parameterNames":["RandomValue_Uint32","Enable_Randomization"]}},{"created":1553712641134,"modified":1553712641134,"id":"4bc4bb97-0690-44b6-9bff-ecac78cd69a2","name":"RandomValue_Uint64","get":{"path":"/api/v1/device/{deviceId}/RandomValue_Uint64","responses":[{"code":"503","description":"service unavailable"}]},"put":{"path":"/api/v1/device/{deviceId}/RandomValue_Uint64","parameterNames":["RandomValue_Uint64","Enable_Randomization"]}}]}}`
	deviceDataRandomFloatGenerator := `{"description":"","id":"` + RandomFloatGeneratorDeviceId + `","name":"Random-Float-Generator01","adminState":"UNLOCKED","operatingState":"ENABLED","protocols":{"other":{"Address":"device-virtual-float-01","Protocol":"300"}},"labels":["device-virtual-example"],"service":{"created":1553712641011,"modified":1553712641011,"origin":1553712641003,"description":"","id":"c58ed126-c713-4cf8-9750-b38c1fd9b43e","name":"device-virtual","lastConnected":0,"lastReported":0,"operatingState":"ENABLED","labels":[],"addressable":{"created":1553712640993,"modified":1553712640993,"origin":1553712640987,"id":"6ec2a3bb-d5c6-49a8-92a5-e41d3fe6ceb7","name":"device-virtual","protocol":"HTTP","method":"POST","address":"192.168.56.101","port":49988,"path":"/api/v1/callback","baseURL":"HTTP://192.168.56.101:49988","url":"HTTP://192.168.56.101:49988/api/v1/callback"},"adminState":"UNLOCKED"},"profile":{"created":1553712641105,"modified":1553712641105,"description":"Example of Device-Virtual","id":"bf0c812e-b028-4501-9263-972217c4a230","name":"Random-Float-Generator","manufacturer":"IOTech","model":"Device-Virtual-01","labels":["device-virtual-example"],"deviceResources":[{"description":"used to decide whether to re-generate a random value","name":"Enable_Randomization","properties":{"value":{"type":"Bool","readWrite":"W","defaultValue":"true"},"units":{"type":"String","readWrite":"R","defaultValue":"Random"}}},{"description":"Generate random float32 value","name":"RandomValue_Float32","properties":{"value":{"type":"Float32","readWrite":"R","defaultValue":"0"},"units":{"type":"String","readWrite":"R","defaultValue":"random float32 value"}}},{"description":"Generate random float64 value","name":"RandomValue_Float64","properties":{"value":{"type":"Float64","readWrite":"R","defaultValue":"0"},"units":{"type":"String","readWrite":"R","defaultValue":"random float64 value"}}}],"deviceCommands":[{"name":"RandomValue_Float32","get":[{"operation":"get","object":"RandomValue_Float32","parameter":"RandomValue_Float32"}],"set":[{"operation":"set","object":"RandomValue_Float32","parameter":"RandomValue_Float32","resource":"RandomValue_Float32"},{"operation":"set","object":"Enable_Randomization","parameter":"Enable_Randomization","resource":"RandomValue_Float32"}]},{"name":"RandomValue_Float64","get":[{"operation":"get","object":"RandomValue_Float64","parameter":"RandomValue_Float64"}],"set":[{"operation":"set","object":"RandomValue_Float64","parameter":"RandomValue_Float64","resource":"RandomValue_Float64"},{"operation":"set","object":"Enable_Randomization","parameter":"Enable_Randomization","resource":"RandomValue_Float64"}]}],"coreCommands":[{"created":1553712641104,"modified":1553712641104,"id":"75c2b3b3-b2f8-431b-9e81-9e2dcb7b1ff6","name":"RandomValue_Float32","get":{"path":"/api/v1/device/{deviceId}/RandomValue_Float32","responses":[{"code":"503","description":"service unavailable"}]},"put":{"path":"/api/v1/device/{deviceId}/RandomValue_Float32","parameterNames":["RandomValue_Float32","Enable_Randomization"]}},{"created":1553712641105,"modified":1553712641105,"id":"5a0a9d40-a34f-4950-871c-c6402520dd28","name":"RandomValue_Float64","get":{"path":"/api/v1/device/{deviceId}/RandomValue_Float64","responses":[{"code":"503","description":"service unavailable"}]},"put":{"path":"/api/v1/device/{deviceId}/RandomValue_Float64","parameterNames":["RandomValue_Float64","Enable_Randomization"]}}]}}`
	newValidDeviceData := `{"created":1552129341782,"modified":1552129341782,"origin":1552129341676,"description":"Auto-generate this virtual device. GS1 AC Drive","id":"37ee8e20-d914-4588-95e6-0088fceb4a99","name":"GS1-AC-Drive01","adminState":"UNLOCKED","operatingState":"ENABLED","addressable":{"created":1552129341674,"modified":0,"origin":1552129341642,"id":"a09cf6c1-427c-4517-bc28-b4cb37df97e3","name":"GS1-AC-Drive01-virtual-addressable","protocol":"OTHER","method":"POST","address":"edgex-device-virtual","port":49990,"path":null,"publisher":null,"user":null,"password":null,"topic":null,"baseURL":"OTHER://edgex-device-virtual:49990","url":"OTHER://edgex-device-virtual:49990"},"lastConnected":0,"lastReported":0,"labels":["modbus","industrial"],"location":null,"service":{"created":1552129332730,"modified":1552129332730,"origin":1552129332645,"description":"","id":"2ce0393b-9814-4bd1-9ea6-3e95140c4b72","name":"edgex-device-virtual","lastConnected":0,"lastReported":0,"operatingState":"ENABLED","labels":["virtual"],"addressable":{"created":1552129332610,"modified":0,"origin":1552129332114,"id":"35fa45f1-f531-4789-b5f6-ec640dc1670c","name":"edgex-device-virtual","protocol":"HTTP","method":"POST","address":"edgex-device-virtual","port":49990,"path":"/api/v1/callback","publisher":null,"user":null,"password":null,"topic":null,"baseURL":"HTTP://edgex-device-virtual:49990","url":"HTTP://edgex-device-virtual:49990/api/v1/callback"},"adminState":"UNLOCKED"},"profile":{"created":1552129341568,"modified":1552129341568,"origin":1552129341469,"description":"GS1 AC Drive","id":"cce09f07-ce9a-40c8-aff4-9b023f09adae","name":"GS1-AC-Drive","manufacturer":"Automation Direct","model":"GS1-10P5","labels":["modbus","industrial"],"objects":null,"deviceResources":[{"description":"Get the OutputVoltage.","name":"HoldingRegister_8455","tag":null,"properties":{"value":{"type":"Float","readWrite":"R","minimum":"0.00","maximum":"300","defaultValue":"0.00","size":"4","word":"2","lsb":null,"mask":"0x00","shift":"0","scale":"1.0","offset":"0.0","base":"0","assertion":null,"signed":true,"precision":"3.2"},"units":{"type":"String","readWrite":"R","defaultValue":"rpm"}},"attributes":{"instance":"8455","type":"doublebyteInt"}},{"description":"Get the RPM.","name":"HoldingRegister_8454","tag":null,"properties":{"value":{"type":"Float","readWrite":"R","minimum":"0.00","maximum":"300","defaultValue":"0.00","size":"4","word":"2","lsb":null,"mask":"0x00","shift":"0","scale":"1.0","offset":"0.0","base":"0","assertion":null,"signed":true,"precision":"3.2"},"units":{"type":"String","readWrite":"R","defaultValue":"Volts"}},"attributes":{"instance":"8454","type":"doublebyteInt"}},{"description":"The status of the device.","name":"HoldingRegister_2331","tag":null,"properties":{"value":{"type":"Boolean","readWrite":"R","minimum":null,"maximum":null,"defaultValue":"false","size":"1","word":"2","lsb":null,"mask":"0x00","shift":"0","scale":"1.0","offset":"0.0","base":"0","assertion":null,"signed":true,"precision":null},"units":{"type":"String","readWrite":"R","defaultValue":"OFF"}},"attributes":{"instance":"8454","type":"doublebyteInt"}},{"description":"whether generating random value in each collection cycle","name":"enableRandomization","tag":null,"properties":{"value":{"type":"Boolean","readWrite":"W","minimum":null,"maximum":null,"defaultValue":"true","size":"1","word":"2","lsb":null,"mask":"0x00","shift":"0","scale":"1.0","offset":"0.0","base":"0","assertion":null,"signed":true,"precision":null},"units":{"type":"String","readWrite":"R","defaultValue":"Random"}},"attributes":null},{"description":"the frequency of collection","name":"collectionFrequency","tag":null,"properties":{"value":{"type":"Integer","readWrite":"W","minimum":"0","maximum":"600","defaultValue":"15","size":"4","word":"2","lsb":null,"mask":"0x00","shift":"0","scale":"1.0","offset":"0.0","base":"0","assertion":null,"signed":true,"precision":"3"},"units":{"type":"String","readWrite":"R","defaultValue":"Seconds"}},"attributes":null}],"deviceCommands":[{"name":"OutputVoltage","get":[{"index":null,"operation":"get","object":"HoldingRegister_8455","property":"value","parameter":"OutputVoltage","resource":null,"secondary":[],"mappings":{}}],"set":[{"index":null,"operation":"set","object":"enableRandomization","property":"value","parameter":"enableRandomization","resource":null,"secondary":[],"mappings":{}},{"index":null,"operation":"set","object":"collectionFrequency","property":"value","parameter":"collectionFrequency","resource":null,"secondary":[],"mappings":{}}]},{"name":"RPM","get":[{"index":null,"operation":"get","object":"HoldingRegister_8454","property":"value","parameter":"RPM","resource":null,"secondary":[],"mappings":{}}],"set":[{"index":null,"operation":"set","object":"enableRandomization","property":"value","parameter":"enableRandomization","resource":null,"secondary":[],"mappings":{}},{"index":null,"operation":"set","object":"collectionFrequency","property":"value","parameter":"collectionFrequency","resource":null,"secondary":[],"mappings":{}}]},{"name":"Status","get":[{"index":null,"operation":"get","object":"HoldingRegister_2331","property":"value","parameter":"Status","resource":null,"secondary":[],"mappings":{}}],"set":[{"index":null,"operation":"set","object":"enableRandomization","property":"value","parameter":"enableRandomization","resource":null,"secondary":[],"mappings":{}},{"index":null,"operation":"set","object":"collectionFrequency","property":"value","parameter":"collectionFrequency","resource":null,"secondary":[],"mappings":{}}]}],"coreCommands":[{"created":1552129341567,"modified":0,"origin":0,"id":"21548fad-69ff-4671-8f4b-3a7d5ab18dec","name":"OutputVoltage","get":{"path":"/api/v1/device/{deviceId}/OutputVoltage","responses":[{"code":"200","description":null,"expectedValues":["HoldingRegister_8455"]},{"code":"503","description":"service unavailable","expectedValues":[]}]},"put":{"path":"/api/v1/device/{deviceId}/OutputVoltage","responses":[{"code":"204","description":"valid and accepted","expectedValues":[]},{"code":"400","description":"bad request","expectedValues":[]},{"code":"503","description":"service unavailable","expectedValues":[]}],"parameterNames":["enableRandomization","collectionFrequency"]}},{"created":1552129341568,"modified":0,"origin":0,"id":"67360b6e-b426-443d-b53c-6490237456e4","name":"Status","get":{"path":"/api/v1/device/{deviceId}/Status","responses":[{"code":"200","description":null,"expectedValues":["HoldingRegister_2331"]},{"code":"503","description":"service unavailable","expectedValues":[]}]},"put":{"path":"/api/v1/device/{deviceId}/Status","responses":[{"code":"204","description":"valid and accepted","expectedValues":[]},{"code":"400","description":"bad request","expectedValues":[]},{"code":"503","description":"service unavailable","expectedValues":[]}],"parameterNames":["enableRandomization","collectionFrequency"]}},{"created":1552129341568,"modified":0,"origin":0,"id":"3d81d381-5055-4928-a8bf-d120d646a2c2","name":"RPM","get":{"path":"/api/v1/device/{deviceId}/RPM","responses":[{"code":"200","description":null,"expectedValues":["HoldingRegister_8454"]},{"code":"503","description":"service unavailable","expectedValues":[]}]},"put":{"path":"/api/v1/device/{deviceId}/RPM","responses":[{"code":"204","description":"valid and accepted","expectedValues":[]},{"code":"400","description":"bad request","expectedValues":[]},{"code":"503","description":"service unavailable","expectedValues":[]}],"parameterNames":["enableRandomization","collectionFrequency"]}}]}}`
	json.Unmarshal([]byte(deviceDataRandomBoolGenerator), &ValidDeviceRandomBoolGenerator)
	json.Unmarshal([]byte(deviceDataRandomIntegerGenerator), &ValidDeviceRandomIntegerGenerator)
	json.Unmarshal([]byte(deviceDataRandomUnsignedIntegerGenerator), &ValidDeviceRandomUnsignedIntegerGenerator)
	json.Unmarshal([]byte(deviceDataRandomFloatGenerator), &ValidDeviceRandomFloatGenerator)
	json.Unmarshal([]byte(deviceDataRandomFloatGenerator), &DuplicateDeviceRandomFloatGenerator)
	json.Unmarshal([]byte(newValidDeviceData), &NewValidDevice)
}
