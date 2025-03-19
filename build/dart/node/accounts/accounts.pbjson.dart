///
//  Generated code. Do not modify.
//  source: node/accounts/accounts.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use credentialsDescriptor instead')
const Credentials$json = const {
  '1': 'Credentials',
  '2': const [
    const {'1': 'type', '3': 1, '4': 1, '5': 9, '10': 'type'},
    const {'1': 'data', '3': 2, '4': 3, '5': 9, '10': 'data'},
  ],
};

/// Descriptor for `Credentials`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List credentialsDescriptor = $convert.base64Decode('CgtDcmVkZW50aWFscxISCgR0eXBlGAEgASgJUgR0eXBlEhIKBGRhdGEYAiADKAlSBGRhdGE=');
@$core.Deprecated('Use accountDescriptor instead')
const Account$json = const {
  '1': 'Account',
  '2': const [
    const {'1': 'uuid', '3': 1, '4': 1, '5': 9, '10': 'uuid'},
    const {'1': 'title', '3': 2, '4': 1, '5': 9, '10': 'title'},
    const {'1': 'enabled', '3': 3, '4': 1, '5': 8, '10': 'enabled'},
    const {'1': 'default_namespace', '3': 4, '4': 1, '5': 9, '10': 'defaultNamespace'},
    const {'1': 'access', '3': 5, '4': 1, '5': 11, '6': '.infinimesh.node.access.Access', '9': 0, '10': 'access', '17': true},
  ],
  '8': const [
    const {'1': '_access'},
  ],
};

/// Descriptor for `Account`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List accountDescriptor = $convert.base64Decode('CgdBY2NvdW50EhIKBHV1aWQYASABKAlSBHV1aWQSFAoFdGl0bGUYAiABKAlSBXRpdGxlEhgKB2VuYWJsZWQYAyABKAhSB2VuYWJsZWQSKwoRZGVmYXVsdF9uYW1lc3BhY2UYBCABKAlSEGRlZmF1bHROYW1lc3BhY2USOwoGYWNjZXNzGAUgASgLMh4uaW5maW5pbWVzaC5ub2RlLmFjY2Vzcy5BY2Nlc3NIAFIGYWNjZXNziAEBQgkKB19hY2Nlc3M=');
@$core.Deprecated('Use createRequestDescriptor instead')
const CreateRequest$json = const {
  '1': 'CreateRequest',
  '2': const [
    const {'1': 'account', '3': 1, '4': 1, '5': 11, '6': '.infinimesh.node.accounts.Account', '10': 'account'},
    const {'1': 'credentials', '3': 2, '4': 1, '5': 11, '6': '.infinimesh.node.accounts.Credentials', '10': 'credentials'},
    const {'1': 'namespace', '3': 3, '4': 1, '5': 9, '10': 'namespace'},
    const {'1': 'access', '3': 4, '4': 1, '5': 5, '9': 0, '10': 'access', '17': true},
  ],
  '8': const [
    const {'1': '_access'},
  ],
};

/// Descriptor for `CreateRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List createRequestDescriptor = $convert.base64Decode('Cg1DcmVhdGVSZXF1ZXN0EjsKB2FjY291bnQYASABKAsyIS5pbmZpbmltZXNoLm5vZGUuYWNjb3VudHMuQWNjb3VudFIHYWNjb3VudBJHCgtjcmVkZW50aWFscxgCIAEoCzIlLmluZmluaW1lc2gubm9kZS5hY2NvdW50cy5DcmVkZW50aWFsc1ILY3JlZGVudGlhbHMSHAoJbmFtZXNwYWNlGAMgASgJUgluYW1lc3BhY2USGwoGYWNjZXNzGAQgASgFSABSBmFjY2Vzc4gBAUIJCgdfYWNjZXNz');
@$core.Deprecated('Use createResponseDescriptor instead')
const CreateResponse$json = const {
  '1': 'CreateResponse',
  '2': const [
    const {'1': 'account', '3': 1, '4': 1, '5': 11, '6': '.infinimesh.node.accounts.Account', '10': 'account'},
  ],
};

/// Descriptor for `CreateResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List createResponseDescriptor = $convert.base64Decode('Cg5DcmVhdGVSZXNwb25zZRI7CgdhY2NvdW50GAEgASgLMiEuaW5maW5pbWVzaC5ub2RlLmFjY291bnRzLkFjY291bnRSB2FjY291bnQ=');
@$core.Deprecated('Use accountsDescriptor instead')
const Accounts$json = const {
  '1': 'Accounts',
  '2': const [
    const {'1': 'accounts', '3': 1, '4': 3, '5': 11, '6': '.infinimesh.node.accounts.Account', '10': 'accounts'},
  ],
};

/// Descriptor for `Accounts`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List accountsDescriptor = $convert.base64Decode('CghBY2NvdW50cxI9CghhY2NvdW50cxgBIAMoCzIhLmluZmluaW1lc2gubm9kZS5hY2NvdW50cy5BY2NvdW50UghhY2NvdW50cw==');
