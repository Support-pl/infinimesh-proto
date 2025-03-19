///
//  Generated code. Do not modify.
//  source: handsfree/handsfree.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use codeDescriptor instead')
const Code$json = const {
  '1': 'Code',
  '2': const [
    const {'1': 'UNKNOWN', '2': 0},
    const {'1': 'AUTH', '2': 1},
    const {'1': 'DATA', '2': 2},
    const {'1': 'SUCCESS', '2': 3},
    const {'1': 'ERROR', '2': 4},
  ],
};

/// Descriptor for `Code`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List codeDescriptor = $convert.base64Decode('CgRDb2RlEgsKB1VOS05PV04QABIICgRBVVRIEAESCAoEREFUQRACEgsKB1NVQ0NFU1MQAxIJCgVFUlJPUhAE');
@$core.Deprecated('Use connectionRequestDescriptor instead')
const ConnectionRequest$json = const {
  '1': 'ConnectionRequest',
  '2': const [
    const {'1': 'app_id', '3': 1, '4': 1, '5': 9, '10': 'appId'},
  ],
};

/// Descriptor for `ConnectionRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List connectionRequestDescriptor = $convert.base64Decode('ChFDb25uZWN0aW9uUmVxdWVzdBIVCgZhcHBfaWQYASABKAlSBWFwcElk');
@$core.Deprecated('Use controlPacketDescriptor instead')
const ControlPacket$json = const {
  '1': 'ControlPacket',
  '2': const [
    const {'1': 'code', '3': 1, '4': 1, '5': 14, '6': '.infinimesh.handsfree.Code', '10': 'code'},
    const {'1': 'payload', '3': 2, '4': 3, '5': 9, '10': 'payload'},
  ],
};

/// Descriptor for `ControlPacket`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List controlPacketDescriptor = $convert.base64Decode('Cg1Db250cm9sUGFja2V0Ei4KBGNvZGUYASABKA4yGi5pbmZpbmltZXNoLmhhbmRzZnJlZS5Db2RlUgRjb2RlEhgKB3BheWxvYWQYAiADKAlSB3BheWxvYWQ=');
