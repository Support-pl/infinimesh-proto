///
//  Generated code. Do not modify.
//  source: plugins/plugins.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use pluginKindDescriptor instead')
const PluginKind$json = const {
  '1': 'PluginKind',
  '2': const [
    const {'1': 'UNKNOWN', '2': 0},
    const {'1': 'EMBEDDED', '2': 1},
  ],
};

/// Descriptor for `PluginKind`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List pluginKindDescriptor = $convert.base64Decode('CgpQbHVnaW5LaW5kEgsKB1VOS05PV04QABIMCghFTUJFRERFRBAB');
@$core.Deprecated('Use pluginDescriptor instead')
const Plugin$json = const {
  '1': 'Plugin',
  '2': const [
    const {'1': 'uuid', '3': 1, '4': 1, '5': 9, '10': 'uuid'},
    const {'1': 'logo', '3': 2, '4': 1, '5': 9, '10': 'logo'},
    const {'1': 'title', '3': 3, '4': 1, '5': 9, '10': 'title'},
    const {'1': 'description', '3': 4, '4': 1, '5': 9, '10': 'description'},
    const {'1': 'public', '3': 5, '4': 1, '5': 8, '10': 'public'},
    const {'1': 'namespace', '3': 6, '4': 1, '5': 9, '9': 0, '10': 'namespace', '17': true},
    const {'1': 'access', '3': 7, '4': 1, '5': 11, '6': '.infinimesh.node.access.Access', '9': 1, '10': 'access', '17': true},
    const {'1': 'kind', '3': 8, '4': 1, '5': 14, '6': '.infinimesh.plugins.PluginKind', '10': 'kind'},
    const {'1': 'embedded_conf', '3': 9, '4': 1, '5': 11, '6': '.infinimesh.plugins.EmbededPluginConf', '9': 2, '10': 'embeddedConf', '17': true},
    const {'1': 'vars', '3': 10, '4': 3, '5': 9, '10': 'vars'},
  ],
  '8': const [
    const {'1': '_namespace'},
    const {'1': '_access'},
    const {'1': '_embedded_conf'},
  ],
};

/// Descriptor for `Plugin`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List pluginDescriptor = $convert.base64Decode('CgZQbHVnaW4SEgoEdXVpZBgBIAEoCVIEdXVpZBISCgRsb2dvGAIgASgJUgRsb2dvEhQKBXRpdGxlGAMgASgJUgV0aXRsZRIgCgtkZXNjcmlwdGlvbhgEIAEoCVILZGVzY3JpcHRpb24SFgoGcHVibGljGAUgASgIUgZwdWJsaWMSIQoJbmFtZXNwYWNlGAYgASgJSABSCW5hbWVzcGFjZYgBARI7CgZhY2Nlc3MYByABKAsyHi5pbmZpbmltZXNoLm5vZGUuYWNjZXNzLkFjY2Vzc0gBUgZhY2Nlc3OIAQESMgoEa2luZBgIIAEoDjIeLmluZmluaW1lc2gucGx1Z2lucy5QbHVnaW5LaW5kUgRraW5kEk8KDWVtYmVkZGVkX2NvbmYYCSABKAsyJS5pbmZpbmltZXNoLnBsdWdpbnMuRW1iZWRlZFBsdWdpbkNvbmZIAlIMZW1iZWRkZWRDb25miAEBEhIKBHZhcnMYCiADKAlSBHZhcnNCDAoKX25hbWVzcGFjZUIJCgdfYWNjZXNzQhAKDl9lbWJlZGRlZF9jb25m');
@$core.Deprecated('Use pluginsDescriptor instead')
const Plugins$json = const {
  '1': 'Plugins',
  '2': const [
    const {'1': 'pool', '3': 1, '4': 3, '5': 11, '6': '.infinimesh.plugins.Plugin', '10': 'pool'},
  ],
};

/// Descriptor for `Plugins`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List pluginsDescriptor = $convert.base64Decode('CgdQbHVnaW5zEi4KBHBvb2wYASADKAsyGi5pbmZpbmltZXNoLnBsdWdpbnMuUGx1Z2luUgRwb29s');
@$core.Deprecated('Use embededPluginConfDescriptor instead')
const EmbededPluginConf$json = const {
  '1': 'EmbededPluginConf',
  '2': const [
    const {'1': 'frame_url', '3': 1, '4': 1, '5': 9, '10': 'frameUrl'},
  ],
};

/// Descriptor for `EmbededPluginConf`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List embededPluginConfDescriptor = $convert.base64Decode('ChFFbWJlZGVkUGx1Z2luQ29uZhIbCglmcmFtZV91cmwYASABKAlSCGZyYW1lVXJs');
@$core.Deprecated('Use listRequestDescriptor instead')
const ListRequest$json = const {
  '1': 'ListRequest',
  '2': const [
    const {'1': 'namespace', '3': 1, '4': 1, '5': 9, '9': 0, '10': 'namespace', '17': true},
  ],
  '8': const [
    const {'1': '_namespace'},
  ],
};

/// Descriptor for `ListRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List listRequestDescriptor = $convert.base64Decode('CgtMaXN0UmVxdWVzdBIhCgluYW1lc3BhY2UYASABKAlIAFIJbmFtZXNwYWNliAEBQgwKCl9uYW1lc3BhY2U=');
