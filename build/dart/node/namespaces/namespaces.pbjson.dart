///
//  Generated code. Do not modify.
//  source: node/namespaces/namespaces.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use pluginDescriptor instead')
const Plugin$json = const {
  '1': 'Plugin',
  '2': const [
    const {'1': 'uuid', '3': 1, '4': 1, '5': 9, '10': 'uuid'},
    const {'1': 'vars', '3': 2, '4': 3, '5': 11, '6': '.infinimesh.node.namespaces.Plugin.VarsEntry', '10': 'vars'},
  ],
  '3': const [Plugin_VarsEntry$json],
};

@$core.Deprecated('Use pluginDescriptor instead')
const Plugin_VarsEntry$json = const {
  '1': 'VarsEntry',
  '2': const [
    const {'1': 'key', '3': 1, '4': 1, '5': 9, '10': 'key'},
    const {'1': 'value', '3': 2, '4': 1, '5': 9, '10': 'value'},
  ],
  '7': const {'7': true},
};

/// Descriptor for `Plugin`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List pluginDescriptor = $convert.base64Decode('CgZQbHVnaW4SEgoEdXVpZBgBIAEoCVIEdXVpZBJACgR2YXJzGAIgAygLMiwuaW5maW5pbWVzaC5ub2RlLm5hbWVzcGFjZXMuUGx1Z2luLlZhcnNFbnRyeVIEdmFycxo3CglWYXJzRW50cnkSEAoDa2V5GAEgASgJUgNrZXkSFAoFdmFsdWUYAiABKAlSBXZhbHVlOgI4AQ==');
@$core.Deprecated('Use namespaceDescriptor instead')
const Namespace$json = const {
  '1': 'Namespace',
  '2': const [
    const {'1': 'uuid', '3': 1, '4': 1, '5': 9, '10': 'uuid'},
    const {'1': 'title', '3': 2, '4': 1, '5': 9, '10': 'title'},
    const {'1': 'access', '3': 3, '4': 1, '5': 11, '6': '.infinimesh.node.access.Access', '9': 0, '10': 'access', '17': true},
    const {'1': 'plugin', '3': 4, '4': 1, '5': 11, '6': '.infinimesh.node.namespaces.Plugin', '9': 1, '10': 'plugin', '17': true},
  ],
  '8': const [
    const {'1': '_access'},
    const {'1': '_plugin'},
  ],
};

/// Descriptor for `Namespace`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List namespaceDescriptor = $convert.base64Decode('CglOYW1lc3BhY2USEgoEdXVpZBgBIAEoCVIEdXVpZBIUCgV0aXRsZRgCIAEoCVIFdGl0bGUSOwoGYWNjZXNzGAMgASgLMh4uaW5maW5pbWVzaC5ub2RlLmFjY2Vzcy5BY2Nlc3NIAFIGYWNjZXNziAEBEj8KBnBsdWdpbhgEIAEoCzIiLmluZmluaW1lc2gubm9kZS5uYW1lc3BhY2VzLlBsdWdpbkgBUgZwbHVnaW6IAQFCCQoHX2FjY2Vzc0IJCgdfcGx1Z2lu');
@$core.Deprecated('Use namespacesDescriptor instead')
const Namespaces$json = const {
  '1': 'Namespaces',
  '2': const [
    const {'1': 'namespaces', '3': 1, '4': 3, '5': 11, '6': '.infinimesh.node.namespaces.Namespace', '10': 'namespaces'},
  ],
};

/// Descriptor for `Namespaces`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List namespacesDescriptor = $convert.base64Decode('CgpOYW1lc3BhY2VzEkUKCm5hbWVzcGFjZXMYASADKAsyJS5pbmZpbmltZXNoLm5vZGUubmFtZXNwYWNlcy5OYW1lc3BhY2VSCm5hbWVzcGFjZXM=');
