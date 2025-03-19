///
//  Generated code. Do not modify.
//  source: node/access/access.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use levelDescriptor instead')
const Level$json = const {
  '1': 'Level',
  '2': const [
    const {'1': 'NONE', '2': 0},
    const {'1': 'READ', '2': 1},
    const {'1': 'MGMT', '2': 2},
    const {'1': 'ADMIN', '2': 3},
    const {'1': 'ROOT', '2': 4},
  ],
};

/// Descriptor for `Level`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List levelDescriptor = $convert.base64Decode('CgVMZXZlbBIICgROT05FEAASCAoEUkVBRBABEggKBE1HTVQQAhIJCgVBRE1JThADEggKBFJPT1QQBA==');
@$core.Deprecated('Use roleDescriptor instead')
const Role$json = const {
  '1': 'Role',
  '2': const [
    const {'1': 'UNSET', '2': 0},
    const {'1': 'OWNER', '2': 1},
  ],
};

/// Descriptor for `Role`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List roleDescriptor = $convert.base64Decode('CgRSb2xlEgkKBVVOU0VUEAASCQoFT1dORVIQAQ==');
@$core.Deprecated('Use accessDescriptor instead')
const Access$json = const {
  '1': 'Access',
  '2': const [
    const {'1': 'level', '3': 1, '4': 1, '5': 14, '6': '.infinimesh.node.access.Level', '10': 'level'},
    const {'1': 'role', '3': 2, '4': 1, '5': 14, '6': '.infinimesh.node.access.Role', '10': 'role'},
    const {'1': 'namespace', '3': 3, '4': 1, '5': 9, '9': 0, '10': 'namespace', '17': true},
  ],
  '8': const [
    const {'1': '_namespace'},
  ],
};

/// Descriptor for `Access`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List accessDescriptor = $convert.base64Decode('CgZBY2Nlc3MSMwoFbGV2ZWwYASABKA4yHS5pbmZpbmltZXNoLm5vZGUuYWNjZXNzLkxldmVsUgVsZXZlbBIwCgRyb2xlGAIgASgOMhwuaW5maW5pbWVzaC5ub2RlLmFjY2Vzcy5Sb2xlUgRyb2xlEiEKCW5hbWVzcGFjZRgDIAEoCUgAUgluYW1lc3BhY2WIAQFCDAoKX25hbWVzcGFjZQ==');
@$core.Deprecated('Use nodeDescriptor instead')
const Node$json = const {
  '1': 'Node',
  '2': const [
    const {'1': 'node', '3': 1, '4': 1, '5': 9, '10': 'node'},
    const {'1': 'parent', '3': 2, '4': 1, '5': 9, '10': 'parent'},
    const {'1': 'edge', '3': 3, '4': 1, '5': 9, '10': 'edge'},
  ],
};

/// Descriptor for `Node`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List nodeDescriptor = $convert.base64Decode('CgROb2RlEhIKBG5vZGUYASABKAlSBG5vZGUSFgoGcGFyZW50GAIgASgJUgZwYXJlbnQSEgoEZWRnZRgDIAEoCVIEZWRnZQ==');
@$core.Deprecated('Use nodesDescriptor instead')
const Nodes$json = const {
  '1': 'Nodes',
  '2': const [
    const {'1': 'nodes', '3': 1, '4': 3, '5': 11, '6': '.infinimesh.node.access.Node', '10': 'nodes'},
  ],
};

/// Descriptor for `Nodes`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List nodesDescriptor = $convert.base64Decode('CgVOb2RlcxIyCgVub2RlcxgBIAMoCzIcLmluZmluaW1lc2gubm9kZS5hY2Nlc3MuTm9kZVIFbm9kZXM=');
