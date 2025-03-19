//
//Copyright © 2021-2023 Infinite Devices GmbH
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// @generated by protoc-gen-es v1.10.0 with parameter "target=js+dts"
// @generated from file node/devices/devices.proto (package infinimesh.node.devices, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage, Struct, Timestamp } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { Access } from "../access/access_pb.js";

/**
 * @generated from message infinimesh.node.devices.Device
 */
export declare class Device extends Message<Device> {
  /**
   * @generated from field: string uuid = 1;
   */
  uuid: string;

  /**
   * @generated from field: string title = 2;
   */
  title: string;

  /**
   * @generated from field: bool enabled = 3;
   */
  enabled: boolean;

  /**
   * @generated from field: infinimesh.node.devices.Certificate certificate = 4;
   */
  certificate?: Certificate;

  /**
   * @generated from field: repeated string tags = 5;
   */
  tags: string[];

  /**
   * @generated from field: bool basic_enabled = 6;
   */
  basicEnabled: boolean;

  /**
   * @generated from field: string token = 7;
   */
  token: string;

  /**
   * @generated from field: optional infinimesh.node.access.Access access = 8;
   */
  access?: Access;

  /**
   * @generated from field: google.protobuf.Struct config = 9;
   */
  config?: Struct;

  /**
   * @generated from field: google.protobuf.Timestamp last_updated = 10;
   */
  lastUpdated?: Timestamp;

  constructor(data?: PartialMessage<Device>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "infinimesh.node.devices.Device";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Device;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Device;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Device;

  static equals(a: Device | PlainMessage<Device> | undefined, b: Device | PlainMessage<Device> | undefined): boolean;
}

/**
 * @generated from message infinimesh.node.devices.Certificate
 */
export declare class Certificate extends Message<Certificate> {
  /**
   * @generated from field: string pem_data = 1;
   */
  pemData: string;

  /**
   * @generated from field: string algorithm = 2;
   */
  algorithm: string;

  /**
   * @generated from field: bytes fingerprint = 3;
   */
  fingerprint: Uint8Array;

  constructor(data?: PartialMessage<Certificate>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "infinimesh.node.devices.Certificate";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Certificate;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Certificate;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Certificate;

  static equals(a: Certificate | PlainMessage<Certificate> | undefined, b: Certificate | PlainMessage<Certificate> | undefined): boolean;
}

/**
 * @generated from message infinimesh.node.devices.Devices
 */
export declare class Devices extends Message<Devices> {
  /**
   * @generated from field: repeated infinimesh.node.devices.Device devices = 1;
   */
  devices: Device[];

  /**
   * @generated from field: int64 total = 2;
   */
  total: bigint;

  constructor(data?: PartialMessage<Devices>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "infinimesh.node.devices.Devices";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Devices;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Devices;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Devices;

  static equals(a: Devices | PlainMessage<Devices> | undefined, b: Devices | PlainMessage<Devices> | undefined): boolean;
}

/**
 * @generated from message infinimesh.node.devices.HandsfreeCreate
 */
export declare class HandsfreeCreate extends Message<HandsfreeCreate> {
  /**
   * @generated from field: string code = 1;
   */
  code: string;

  /**
   * @generated from field: repeated string payload = 2;
   */
  payload: string[];

  constructor(data?: PartialMessage<HandsfreeCreate>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "infinimesh.node.devices.HandsfreeCreate";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): HandsfreeCreate;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): HandsfreeCreate;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): HandsfreeCreate;

  static equals(a: HandsfreeCreate | PlainMessage<HandsfreeCreate> | undefined, b: HandsfreeCreate | PlainMessage<HandsfreeCreate> | undefined): boolean;
}

/**
 * @generated from message infinimesh.node.devices.CreateRequest
 */
export declare class CreateRequest extends Message<CreateRequest> {
  /**
   * @generated from field: infinimesh.node.devices.Device device = 1;
   */
  device?: Device;

  /**
   * Namespace to put device under
   *
   * @generated from field: string namespace = 2;
   */
  namespace: string;

  /**
   * if present, will attempt to obtain
   *
   * @generated from field: optional infinimesh.node.devices.HandsfreeCreate handsfree = 3;
   */
  handsfree?: HandsfreeCreate;

  constructor(data?: PartialMessage<CreateRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "infinimesh.node.devices.CreateRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateRequest;

  static equals(a: CreateRequest | PlainMessage<CreateRequest> | undefined, b: CreateRequest | PlainMessage<CreateRequest> | undefined): boolean;
}

/**
 * @generated from message infinimesh.node.devices.CreateResponse
 */
export declare class CreateResponse extends Message<CreateResponse> {
  /**
   * @generated from field: infinimesh.node.devices.Device device = 1;
   */
  device?: Device;

  constructor(data?: PartialMessage<CreateResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "infinimesh.node.devices.CreateResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateResponse;

  static equals(a: CreateResponse | PlainMessage<CreateResponse> | undefined, b: CreateResponse | PlainMessage<CreateResponse> | undefined): boolean;
}

/**
 * @generated from message infinimesh.node.devices.GetByFingerprintRequest
 */
export declare class GetByFingerprintRequest extends Message<GetByFingerprintRequest> {
  /**
   * @generated from field: bytes fingerprint = 1;
   */
  fingerprint: Uint8Array;

  constructor(data?: PartialMessage<GetByFingerprintRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "infinimesh.node.devices.GetByFingerprintRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetByFingerprintRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetByFingerprintRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetByFingerprintRequest;

  static equals(a: GetByFingerprintRequest | PlainMessage<GetByFingerprintRequest> | undefined, b: GetByFingerprintRequest | PlainMessage<GetByFingerprintRequest> | undefined): boolean;
}

