///
//  Generated code. Do not modify.
//  source: node/node.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use tokenRequestDescriptor instead')
const TokenRequest$json = const {
  '1': 'TokenRequest',
  '2': const [
    const {'1': 'auth', '3': 1, '4': 1, '5': 11, '6': '.infinimesh.node.accounts.Credentials', '10': 'auth'},
    const {'1': 'exp', '3': 2, '4': 1, '5': 3, '10': 'exp'},
    const {'1': 'uuid', '3': 3, '4': 1, '5': 9, '9': 0, '10': 'uuid', '17': true},
    const {'1': 'inf', '3': 4, '4': 1, '5': 8, '9': 1, '10': 'inf', '17': true},
  ],
  '8': const [
    const {'1': '_uuid'},
    const {'1': '_inf'},
  ],
};

/// Descriptor for `TokenRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List tokenRequestDescriptor = $convert.base64Decode('CgxUb2tlblJlcXVlc3QSOQoEYXV0aBgBIAEoCzIlLmluZmluaW1lc2gubm9kZS5hY2NvdW50cy5DcmVkZW50aWFsc1IEYXV0aBIQCgNleHAYAiABKANSA2V4cBIXCgR1dWlkGAMgASgJSABSBHV1aWSIAQESFQoDaW5mGAQgASgISAFSA2luZogBAUIHCgVfdXVpZEIGCgRfaW5m');
@$core.Deprecated('Use tokenResponseDescriptor instead')
const TokenResponse$json = const {
  '1': 'TokenResponse',
  '2': const [
    const {'1': 'token', '3': 1, '4': 1, '5': 9, '10': 'token'},
  ],
};

/// Descriptor for `TokenResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List tokenResponseDescriptor = $convert.base64Decode('Cg1Ub2tlblJlc3BvbnNlEhQKBXRva2VuGAEgASgJUgV0b2tlbg==');
@$core.Deprecated('Use setCredentialsRequestDescriptor instead')
const SetCredentialsRequest$json = const {
  '1': 'SetCredentialsRequest',
  '2': const [
    const {'1': 'uuid', '3': 1, '4': 1, '5': 9, '10': 'uuid'},
    const {'1': 'credentials', '3': 2, '4': 1, '5': 11, '6': '.infinimesh.node.accounts.Credentials', '10': 'credentials'},
  ],
};

/// Descriptor for `SetCredentialsRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List setCredentialsRequestDescriptor = $convert.base64Decode('ChVTZXRDcmVkZW50aWFsc1JlcXVlc3QSEgoEdXVpZBgBIAEoCVIEdXVpZBJHCgtjcmVkZW50aWFscxgCIAEoCzIlLmluZmluaW1lc2gubm9kZS5hY2NvdW50cy5DcmVkZW50aWFsc1ILY3JlZGVudGlhbHM=');
@$core.Deprecated('Use setCredentialsResponseDescriptor instead')
const SetCredentialsResponse$json = const {
  '1': 'SetCredentialsResponse',
};

/// Descriptor for `SetCredentialsResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List setCredentialsResponseDescriptor = $convert.base64Decode('ChZTZXRDcmVkZW50aWFsc1Jlc3BvbnNl');
@$core.Deprecated('Use deleteResponseDescriptor instead')
const DeleteResponse$json = const {
  '1': 'DeleteResponse',
};

/// Descriptor for `DeleteResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List deleteResponseDescriptor = $convert.base64Decode('Cg5EZWxldGVSZXNwb25zZQ==');
@$core.Deprecated('Use emptyMessageDescriptor instead')
const EmptyMessage$json = const {
  '1': 'EmptyMessage',
};

/// Descriptor for `EmptyMessage`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List emptyMessageDescriptor = $convert.base64Decode('CgxFbXB0eU1lc3NhZ2U=');
@$core.Deprecated('Use joinRequestDescriptor instead')
const JoinRequest$json = const {
  '1': 'JoinRequest',
  '2': const [
    const {'1': 'namespace', '3': 1, '4': 1, '5': 9, '10': 'namespace'},
    const {'1': 'account', '3': 2, '4': 1, '5': 9, '10': 'account'},
    const {'1': 'access', '3': 3, '4': 1, '5': 14, '6': '.infinimesh.node.access.Level', '10': 'access'},
  ],
};

/// Descriptor for `JoinRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List joinRequestDescriptor = $convert.base64Decode('CgtKb2luUmVxdWVzdBIcCgluYW1lc3BhY2UYASABKAlSCW5hbWVzcGFjZRIYCgdhY2NvdW50GAIgASgJUgdhY2NvdW50EjUKBmFjY2VzcxgDIAEoDjIdLmluZmluaW1lc2gubm9kZS5hY2Nlc3MuTGV2ZWxSBmFjY2Vzcw==');
@$core.Deprecated('Use devicesTokenRequestDescriptor instead')
const DevicesTokenRequest$json = const {
  '1': 'DevicesTokenRequest',
  '2': const [
    const {'1': 'devices', '3': 1, '4': 3, '5': 9, '10': 'devices'},
    const {'1': 'exp', '3': 2, '4': 1, '5': 3, '10': 'exp'},
    const {'1': 'post', '3': 3, '4': 1, '5': 8, '10': 'post'},
  ],
};

/// Descriptor for `DevicesTokenRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List devicesTokenRequestDescriptor = $convert.base64Decode('ChNEZXZpY2VzVG9rZW5SZXF1ZXN0EhgKB2RldmljZXMYASADKAlSB2RldmljZXMSEAoDZXhwGAIgASgDUgNleHASEgoEcG9zdBgDIAEoCFIEcG9zdA==');
@$core.Deprecated('Use lDAPProvidersDescriptor instead')
const LDAPProviders$json = const {
  '1': 'LDAPProviders',
  '2': const [
    const {'1': 'providers', '3': 1, '4': 3, '5': 11, '6': '.infinimesh.node.LDAPProviders.ProvidersEntry', '10': 'providers'},
  ],
  '3': const [LDAPProviders_ProvidersEntry$json],
};

@$core.Deprecated('Use lDAPProvidersDescriptor instead')
const LDAPProviders_ProvidersEntry$json = const {
  '1': 'ProvidersEntry',
  '2': const [
    const {'1': 'key', '3': 1, '4': 1, '5': 9, '10': 'key'},
    const {'1': 'value', '3': 2, '4': 1, '5': 9, '10': 'value'},
  ],
  '7': const {'7': true},
};

/// Descriptor for `LDAPProviders`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List lDAPProvidersDescriptor = $convert.base64Decode('Cg1MREFQUHJvdmlkZXJzEksKCXByb3ZpZGVycxgBIAMoCzItLmluZmluaW1lc2gubm9kZS5MREFQUHJvdmlkZXJzLlByb3ZpZGVyc0VudHJ5Uglwcm92aWRlcnMaPAoOUHJvdmlkZXJzRW50cnkSEAoDa2V5GAEgASgJUgNrZXkSFAoFdmFsdWUYAiABKAlSBXZhbHVlOgI4AQ==');
