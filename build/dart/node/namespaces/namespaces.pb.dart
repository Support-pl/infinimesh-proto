///
//  Generated code. Do not modify.
//  source: node/namespaces/namespaces.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import '../access/access.pb.dart' as $0;

class Plugin extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'Plugin', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'infinimesh.node.namespaces'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'uuid')
    ..m<$core.String, $core.String>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'vars', entryClassName: 'Plugin.VarsEntry', keyFieldType: $pb.PbFieldType.OS, valueFieldType: $pb.PbFieldType.OS, packageName: const $pb.PackageName('infinimesh.node.namespaces'))
    ..hasRequiredFields = false
  ;

  Plugin._() : super();
  factory Plugin({
    $core.String? uuid,
    $core.Map<$core.String, $core.String>? vars,
  }) {
    final _result = create();
    if (uuid != null) {
      _result.uuid = uuid;
    }
    if (vars != null) {
      _result.vars.addAll(vars);
    }
    return _result;
  }
  factory Plugin.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Plugin.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Plugin clone() => Plugin()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Plugin copyWith(void Function(Plugin) updates) => super.copyWith((message) => updates(message as Plugin)) as Plugin; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Plugin create() => Plugin._();
  Plugin createEmptyInstance() => create();
  static $pb.PbList<Plugin> createRepeated() => $pb.PbList<Plugin>();
  @$core.pragma('dart2js:noInline')
  static Plugin getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Plugin>(create);
  static Plugin? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get uuid => $_getSZ(0);
  @$pb.TagNumber(1)
  set uuid($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasUuid() => $_has(0);
  @$pb.TagNumber(1)
  void clearUuid() => clearField(1);

  @$pb.TagNumber(2)
  $core.Map<$core.String, $core.String> get vars => $_getMap(1);
}

class Namespace extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'Namespace', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'infinimesh.node.namespaces'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'uuid')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'title')
    ..aOM<$0.Access>(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'access', subBuilder: $0.Access.create)
    ..aOM<Plugin>(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'plugin', subBuilder: Plugin.create)
    ..hasRequiredFields = false
  ;

  Namespace._() : super();
  factory Namespace({
    $core.String? uuid,
    $core.String? title,
    $0.Access? access,
    Plugin? plugin,
  }) {
    final _result = create();
    if (uuid != null) {
      _result.uuid = uuid;
    }
    if (title != null) {
      _result.title = title;
    }
    if (access != null) {
      _result.access = access;
    }
    if (plugin != null) {
      _result.plugin = plugin;
    }
    return _result;
  }
  factory Namespace.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Namespace.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Namespace clone() => Namespace()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Namespace copyWith(void Function(Namespace) updates) => super.copyWith((message) => updates(message as Namespace)) as Namespace; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Namespace create() => Namespace._();
  Namespace createEmptyInstance() => create();
  static $pb.PbList<Namespace> createRepeated() => $pb.PbList<Namespace>();
  @$core.pragma('dart2js:noInline')
  static Namespace getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Namespace>(create);
  static Namespace? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get uuid => $_getSZ(0);
  @$pb.TagNumber(1)
  set uuid($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasUuid() => $_has(0);
  @$pb.TagNumber(1)
  void clearUuid() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get title => $_getSZ(1);
  @$pb.TagNumber(2)
  set title($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasTitle() => $_has(1);
  @$pb.TagNumber(2)
  void clearTitle() => clearField(2);

  @$pb.TagNumber(3)
  $0.Access get access => $_getN(2);
  @$pb.TagNumber(3)
  set access($0.Access v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasAccess() => $_has(2);
  @$pb.TagNumber(3)
  void clearAccess() => clearField(3);
  @$pb.TagNumber(3)
  $0.Access ensureAccess() => $_ensure(2);

  @$pb.TagNumber(4)
  Plugin get plugin => $_getN(3);
  @$pb.TagNumber(4)
  set plugin(Plugin v) { setField(4, v); }
  @$pb.TagNumber(4)
  $core.bool hasPlugin() => $_has(3);
  @$pb.TagNumber(4)
  void clearPlugin() => clearField(4);
  @$pb.TagNumber(4)
  Plugin ensurePlugin() => $_ensure(3);
}

class Namespaces extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'Namespaces', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'infinimesh.node.namespaces'), createEmptyInstance: create)
    ..pc<Namespace>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'namespaces', $pb.PbFieldType.PM, subBuilder: Namespace.create)
    ..hasRequiredFields = false
  ;

  Namespaces._() : super();
  factory Namespaces({
    $core.Iterable<Namespace>? namespaces,
  }) {
    final _result = create();
    if (namespaces != null) {
      _result.namespaces.addAll(namespaces);
    }
    return _result;
  }
  factory Namespaces.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Namespaces.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Namespaces clone() => Namespaces()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Namespaces copyWith(void Function(Namespaces) updates) => super.copyWith((message) => updates(message as Namespaces)) as Namespaces; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Namespaces create() => Namespaces._();
  Namespaces createEmptyInstance() => create();
  static $pb.PbList<Namespaces> createRepeated() => $pb.PbList<Namespaces>();
  @$core.pragma('dart2js:noInline')
  static Namespaces getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Namespaces>(create);
  static Namespaces? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<Namespace> get namespaces => $_getList(0);
}

