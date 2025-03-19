///
//  Generated code. Do not modify.
//  source: node/devices/devices.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use deviceDescriptor instead')
const Device$json = const {
  '1': 'Device',
  '2': const [
    const {'1': 'uuid', '3': 1, '4': 1, '5': 9, '10': 'uuid'},
    const {'1': 'title', '3': 2, '4': 1, '5': 9, '10': 'title'},
    const {'1': 'enabled', '3': 3, '4': 1, '5': 8, '10': 'enabled'},
    const {'1': 'certificate', '3': 4, '4': 1, '5': 11, '6': '.infinimesh.node.devices.Certificate', '10': 'certificate'},
    const {'1': 'tags', '3': 5, '4': 3, '5': 9, '10': 'tags'},
    const {'1': 'basic_enabled', '3': 6, '4': 1, '5': 8, '10': 'basicEnabled'},
    const {'1': 'token', '3': 7, '4': 1, '5': 9, '10': 'token'},
    const {'1': 'access', '3': 8, '4': 1, '5': 11, '6': '.infinimesh.node.access.Access', '9': 0, '10': 'access', '17': true},
  ],
  '8': const [
    const {'1': '_access'},
  ],
};

/// Descriptor for `Device`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List deviceDescriptor = $convert.base64Decode('CgZEZXZpY2USEgoEdXVpZBgBIAEoCVIEdXVpZBIUCgV0aXRsZRgCIAEoCVIFdGl0bGUSGAoHZW5hYmxlZBgDIAEoCFIHZW5hYmxlZBJGCgtjZXJ0aWZpY2F0ZRgEIAEoCzIkLmluZmluaW1lc2gubm9kZS5kZXZpY2VzLkNlcnRpZmljYXRlUgtjZXJ0aWZpY2F0ZRISCgR0YWdzGAUgAygJUgR0YWdzEiMKDWJhc2ljX2VuYWJsZWQYBiABKAhSDGJhc2ljRW5hYmxlZBIUCgV0b2tlbhgHIAEoCVIFdG9rZW4SOwoGYWNjZXNzGAggASgLMh4uaW5maW5pbWVzaC5ub2RlLmFjY2Vzcy5BY2Nlc3NIAFIGYWNjZXNziAEBQgkKB19hY2Nlc3M=');
@$core.Deprecated('Use certificateDescriptor instead')
const Certificate$json = const {
  '1': 'Certificate',
  '2': const [
    const {'1': 'pem_data', '3': 1, '4': 1, '5': 9, '10': 'pemData'},
    const {'1': 'algorithm', '3': 2, '4': 1, '5': 9, '10': 'algorithm'},
    const {'1': 'fingerprint', '3': 3, '4': 1, '5': 12, '10': 'fingerprint'},
  ],
};

/// Descriptor for `Certificate`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List certificateDescriptor = $convert.base64Decode('CgtDZXJ0aWZpY2F0ZRIZCghwZW1fZGF0YRgBIAEoCVIHcGVtRGF0YRIcCglhbGdvcml0aG0YAiABKAlSCWFsZ29yaXRobRIgCgtmaW5nZXJwcmludBgDIAEoDFILZmluZ2VycHJpbnQ=');
@$core.Deprecated('Use devicesDescriptor instead')
const Devices$json = const {
  '1': 'Devices',
  '2': const [
    const {'1': 'devices', '3': 1, '4': 3, '5': 11, '6': '.infinimesh.node.devices.Device', '10': 'devices'},
  ],
};

/// Descriptor for `Devices`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List devicesDescriptor = $convert.base64Decode('CgdEZXZpY2VzEjkKB2RldmljZXMYASADKAsyHy5pbmZpbmltZXNoLm5vZGUuZGV2aWNlcy5EZXZpY2VSB2RldmljZXM=');
@$core.Deprecated('Use createRequestDescriptor instead')
const CreateRequest$json = const {
  '1': 'CreateRequest',
  '2': const [
    const {'1': 'device', '3': 1, '4': 1, '5': 11, '6': '.infinimesh.node.devices.Device', '10': 'device'},
    const {'1': 'namespace', '3': 2, '4': 1, '5': 9, '10': 'namespace'},
  ],
};

/// Descriptor for `CreateRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List createRequestDescriptor = $convert.base64Decode('Cg1DcmVhdGVSZXF1ZXN0EjcKBmRldmljZRgBIAEoCzIfLmluZmluaW1lc2gubm9kZS5kZXZpY2VzLkRldmljZVIGZGV2aWNlEhwKCW5hbWVzcGFjZRgCIAEoCVIJbmFtZXNwYWNl');
@$core.Deprecated('Use createResponseDescriptor instead')
const CreateResponse$json = const {
  '1': 'CreateResponse',
  '2': const [
    const {'1': 'device', '3': 1, '4': 1, '5': 11, '6': '.infinimesh.node.devices.Device', '10': 'device'},
  ],
};

/// Descriptor for `CreateResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List createResponseDescriptor = $convert.base64Decode('Cg5DcmVhdGVSZXNwb25zZRI3CgZkZXZpY2UYASABKAsyHy5pbmZpbmltZXNoLm5vZGUuZGV2aWNlcy5EZXZpY2VSBmRldmljZQ==');
@$core.Deprecated('Use getByFingerprintRequestDescriptor instead')
const GetByFingerprintRequest$json = const {
  '1': 'GetByFingerprintRequest',
  '2': const [
    const {'1': 'fingerprint', '3': 1, '4': 1, '5': 12, '10': 'fingerprint'},
  ],
};

/// Descriptor for `GetByFingerprintRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List getByFingerprintRequestDescriptor = $convert.base64Decode('ChdHZXRCeUZpbmdlcnByaW50UmVxdWVzdBIgCgtmaW5nZXJwcmludBgBIAEoDFILZmluZ2VycHJpbnQ=');
