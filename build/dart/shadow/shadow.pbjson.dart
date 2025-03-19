///
//  Generated code. Do not modify.
//  source: shadow/shadow.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use stateKeyDescriptor instead')
const StateKey$json = const {
  '1': 'StateKey',
  '2': const [
    const {'1': 'REPORTED', '2': 0},
    const {'1': 'DESIRED', '2': 1},
  ],
};

/// Descriptor for `StateKey`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List stateKeyDescriptor = $convert.base64Decode('CghTdGF0ZUtleRIMCghSRVBPUlRFRBAAEgsKB0RFU0lSRUQQAQ==');
@$core.Deprecated('Use stateDescriptor instead')
const State$json = const {
  '1': 'State',
  '2': const [
    const {'1': 'timestamp', '3': 1, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'timestamp'},
    const {'1': 'data', '3': 2, '4': 1, '5': 11, '6': '.google.protobuf.Struct', '10': 'data'},
  ],
};

/// Descriptor for `State`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List stateDescriptor = $convert.base64Decode('CgVTdGF0ZRI4Cgl0aW1lc3RhbXAYASABKAsyGi5nb29nbGUucHJvdG9idWYuVGltZXN0YW1wUgl0aW1lc3RhbXASKwoEZGF0YRgCIAEoCzIXLmdvb2dsZS5wcm90b2J1Zi5TdHJ1Y3RSBGRhdGE=');
@$core.Deprecated('Use shadowDescriptor instead')
const Shadow$json = const {
  '1': 'Shadow',
  '2': const [
    const {'1': 'device', '3': 1, '4': 1, '5': 9, '10': 'device'},
    const {'1': 'reported', '3': 2, '4': 1, '5': 11, '6': '.infinimesh.shadow.State', '10': 'reported'},
    const {'1': 'desired', '3': 3, '4': 1, '5': 11, '6': '.infinimesh.shadow.State', '10': 'desired'},
  ],
};

/// Descriptor for `Shadow`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List shadowDescriptor = $convert.base64Decode('CgZTaGFkb3cSFgoGZGV2aWNlGAEgASgJUgZkZXZpY2USNAoIcmVwb3J0ZWQYAiABKAsyGC5pbmZpbmltZXNoLnNoYWRvdy5TdGF0ZVIIcmVwb3J0ZWQSMgoHZGVzaXJlZBgDIAEoCzIYLmluZmluaW1lc2guc2hhZG93LlN0YXRlUgdkZXNpcmVk');
@$core.Deprecated('Use getRequestDescriptor instead')
const GetRequest$json = const {
  '1': 'GetRequest',
  '2': const [
    const {'1': 'pool', '3': 1, '4': 3, '5': 9, '10': 'pool'},
  ],
};

/// Descriptor for `GetRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List getRequestDescriptor = $convert.base64Decode('CgpHZXRSZXF1ZXN0EhIKBHBvb2wYASADKAlSBHBvb2w=');
@$core.Deprecated('Use getResponseDescriptor instead')
const GetResponse$json = const {
  '1': 'GetResponse',
  '2': const [
    const {'1': 'shadows', '3': 1, '4': 3, '5': 11, '6': '.infinimesh.shadow.Shadow', '10': 'shadows'},
  ],
};

/// Descriptor for `GetResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List getResponseDescriptor = $convert.base64Decode('CgtHZXRSZXNwb25zZRIzCgdzaGFkb3dzGAEgAygLMhkuaW5maW5pbWVzaC5zaGFkb3cuU2hhZG93UgdzaGFkb3dz');
@$core.Deprecated('Use removeRequestDescriptor instead')
const RemoveRequest$json = const {
  '1': 'RemoveRequest',
  '2': const [
    const {'1': 'device', '3': 1, '4': 1, '5': 9, '10': 'device'},
    const {'1': 'state_key', '3': 2, '4': 1, '5': 14, '6': '.infinimesh.shadow.StateKey', '10': 'stateKey'},
    const {'1': 'key', '3': 3, '4': 1, '5': 9, '10': 'key'},
  ],
};

/// Descriptor for `RemoveRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List removeRequestDescriptor = $convert.base64Decode('Cg1SZW1vdmVSZXF1ZXN0EhYKBmRldmljZRgBIAEoCVIGZGV2aWNlEjgKCXN0YXRlX2tleRgCIAEoDjIbLmluZmluaW1lc2guc2hhZG93LlN0YXRlS2V5UghzdGF0ZUtleRIQCgNrZXkYAyABKAlSA2tleQ==');
@$core.Deprecated('Use streamShadowRequestDescriptor instead')
const StreamShadowRequest$json = const {
  '1': 'StreamShadowRequest',
  '2': const [
    const {'1': 'devices', '3': 1, '4': 3, '5': 9, '10': 'devices'},
    const {'1': 'only_delta', '3': 2, '4': 1, '5': 8, '10': 'onlyDelta'},
    const {'1': 'sync', '3': 3, '4': 1, '5': 8, '10': 'sync'},
  ],
};

/// Descriptor for `StreamShadowRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List streamShadowRequestDescriptor = $convert.base64Decode('ChNTdHJlYW1TaGFkb3dSZXF1ZXN0EhgKB2RldmljZXMYASADKAlSB2RldmljZXMSHQoKb25seV9kZWx0YRgCIAEoCFIJb25seURlbHRhEhIKBHN5bmMYAyABKAhSBHN5bmM=');
