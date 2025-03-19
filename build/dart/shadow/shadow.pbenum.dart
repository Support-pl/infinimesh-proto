///
//  Generated code. Do not modify.
//  source: shadow/shadow.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

// ignore_for_file: UNDEFINED_SHOWN_NAME
import 'dart:core' as $core;
import 'package:protobuf/protobuf.dart' as $pb;

class StateKey extends $pb.ProtobufEnum {
  static const StateKey REPORTED = StateKey._(0, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'REPORTED');
  static const StateKey DESIRED = StateKey._(1, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'DESIRED');

  static const $core.List<StateKey> values = <StateKey> [
    REPORTED,
    DESIRED,
  ];

  static final $core.Map<$core.int, StateKey> _byValue = $pb.ProtobufEnum.initByValue(values);
  static StateKey? valueOf($core.int value) => _byValue[value];

  const StateKey._($core.int v, $core.String n) : super(v, n);
}

