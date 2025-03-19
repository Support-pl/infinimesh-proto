///
//  Generated code. Do not modify.
//  source: node/access/access.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

// ignore_for_file: UNDEFINED_SHOWN_NAME
import 'dart:core' as $core;
import 'package:protobuf/protobuf.dart' as $pb;

class Level extends $pb.ProtobufEnum {
  static const Level NONE = Level._(0, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'NONE');
  static const Level READ = Level._(1, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'READ');
  static const Level MGMT = Level._(2, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'MGMT');
  static const Level ADMIN = Level._(3, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'ADMIN');
  static const Level ROOT = Level._(4, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'ROOT');

  static const $core.List<Level> values = <Level> [
    NONE,
    READ,
    MGMT,
    ADMIN,
    ROOT,
  ];

  static final $core.Map<$core.int, Level> _byValue = $pb.ProtobufEnum.initByValue(values);
  static Level? valueOf($core.int value) => _byValue[value];

  const Level._($core.int v, $core.String n) : super(v, n);
}

class Role extends $pb.ProtobufEnum {
  static const Role UNSET = Role._(0, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'UNSET');
  static const Role OWNER = Role._(1, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'OWNER');

  static const $core.List<Role> values = <Role> [
    UNSET,
    OWNER,
  ];

  static final $core.Map<$core.int, Role> _byValue = $pb.ProtobufEnum.initByValue(values);
  static Role? valueOf($core.int value) => _byValue[value];

  const Role._($core.int v, $core.String n) : super(v, n);
}

