///
//  Generated code. Do not modify.
//  source: plugins/plugins.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import '../node/access/access.pb.dart' as $1;

import 'plugins.pbenum.dart';

export 'plugins.pbenum.dart';

class Plugin extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'Plugin', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'infinimesh.plugins'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'uuid')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'logo')
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'title')
    ..aOS(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'description')
    ..aOB(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'public')
    ..aOS(6, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'namespace')
    ..aOM<$1.Access>(7, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'access', subBuilder: $1.Access.create)
    ..e<PluginKind>(8, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'kind', $pb.PbFieldType.OE, defaultOrMaker: PluginKind.UNKNOWN, valueOf: PluginKind.valueOf, enumValues: PluginKind.values)
    ..aOM<EmbededPluginConf>(9, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'embeddedConf', subBuilder: EmbededPluginConf.create)
    ..pPS(10, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'vars')
    ..hasRequiredFields = false
  ;

  Plugin._() : super();
  factory Plugin({
    $core.String? uuid,
    $core.String? logo,
    $core.String? title,
    $core.String? description,
    $core.bool? public,
    $core.String? namespace,
    $1.Access? access,
    PluginKind? kind,
    EmbededPluginConf? embeddedConf,
    $core.Iterable<$core.String>? vars,
  }) {
    final _result = create();
    if (uuid != null) {
      _result.uuid = uuid;
    }
    if (logo != null) {
      _result.logo = logo;
    }
    if (title != null) {
      _result.title = title;
    }
    if (description != null) {
      _result.description = description;
    }
    if (public != null) {
      _result.public = public;
    }
    if (namespace != null) {
      _result.namespace = namespace;
    }
    if (access != null) {
      _result.access = access;
    }
    if (kind != null) {
      _result.kind = kind;
    }
    if (embeddedConf != null) {
      _result.embeddedConf = embeddedConf;
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
  $core.String get logo => $_getSZ(1);
  @$pb.TagNumber(2)
  set logo($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasLogo() => $_has(1);
  @$pb.TagNumber(2)
  void clearLogo() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get title => $_getSZ(2);
  @$pb.TagNumber(3)
  set title($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasTitle() => $_has(2);
  @$pb.TagNumber(3)
  void clearTitle() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get description => $_getSZ(3);
  @$pb.TagNumber(4)
  set description($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasDescription() => $_has(3);
  @$pb.TagNumber(4)
  void clearDescription() => clearField(4);

  @$pb.TagNumber(5)
  $core.bool get public => $_getBF(4);
  @$pb.TagNumber(5)
  set public($core.bool v) { $_setBool(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasPublic() => $_has(4);
  @$pb.TagNumber(5)
  void clearPublic() => clearField(5);

  @$pb.TagNumber(6)
  $core.String get namespace => $_getSZ(5);
  @$pb.TagNumber(6)
  set namespace($core.String v) { $_setString(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasNamespace() => $_has(5);
  @$pb.TagNumber(6)
  void clearNamespace() => clearField(6);

  @$pb.TagNumber(7)
  $1.Access get access => $_getN(6);
  @$pb.TagNumber(7)
  set access($1.Access v) { setField(7, v); }
  @$pb.TagNumber(7)
  $core.bool hasAccess() => $_has(6);
  @$pb.TagNumber(7)
  void clearAccess() => clearField(7);
  @$pb.TagNumber(7)
  $1.Access ensureAccess() => $_ensure(6);

  @$pb.TagNumber(8)
  PluginKind get kind => $_getN(7);
  @$pb.TagNumber(8)
  set kind(PluginKind v) { setField(8, v); }
  @$pb.TagNumber(8)
  $core.bool hasKind() => $_has(7);
  @$pb.TagNumber(8)
  void clearKind() => clearField(8);

  @$pb.TagNumber(9)
  EmbededPluginConf get embeddedConf => $_getN(8);
  @$pb.TagNumber(9)
  set embeddedConf(EmbededPluginConf v) { setField(9, v); }
  @$pb.TagNumber(9)
  $core.bool hasEmbeddedConf() => $_has(8);
  @$pb.TagNumber(9)
  void clearEmbeddedConf() => clearField(9);
  @$pb.TagNumber(9)
  EmbededPluginConf ensureEmbeddedConf() => $_ensure(8);

  @$pb.TagNumber(10)
  $core.List<$core.String> get vars => $_getList(9);
}

class Plugins extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'Plugins', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'infinimesh.plugins'), createEmptyInstance: create)
    ..pc<Plugin>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'pool', $pb.PbFieldType.PM, subBuilder: Plugin.create)
    ..hasRequiredFields = false
  ;

  Plugins._() : super();
  factory Plugins({
    $core.Iterable<Plugin>? pool,
  }) {
    final _result = create();
    if (pool != null) {
      _result.pool.addAll(pool);
    }
    return _result;
  }
  factory Plugins.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Plugins.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Plugins clone() => Plugins()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Plugins copyWith(void Function(Plugins) updates) => super.copyWith((message) => updates(message as Plugins)) as Plugins; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Plugins create() => Plugins._();
  Plugins createEmptyInstance() => create();
  static $pb.PbList<Plugins> createRepeated() => $pb.PbList<Plugins>();
  @$core.pragma('dart2js:noInline')
  static Plugins getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Plugins>(create);
  static Plugins? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<Plugin> get pool => $_getList(0);
}

class EmbededPluginConf extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'EmbededPluginConf', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'infinimesh.plugins'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'frameUrl')
    ..hasRequiredFields = false
  ;

  EmbededPluginConf._() : super();
  factory EmbededPluginConf({
    $core.String? frameUrl,
  }) {
    final _result = create();
    if (frameUrl != null) {
      _result.frameUrl = frameUrl;
    }
    return _result;
  }
  factory EmbededPluginConf.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory EmbededPluginConf.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  EmbededPluginConf clone() => EmbededPluginConf()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  EmbededPluginConf copyWith(void Function(EmbededPluginConf) updates) => super.copyWith((message) => updates(message as EmbededPluginConf)) as EmbededPluginConf; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static EmbededPluginConf create() => EmbededPluginConf._();
  EmbededPluginConf createEmptyInstance() => create();
  static $pb.PbList<EmbededPluginConf> createRepeated() => $pb.PbList<EmbededPluginConf>();
  @$core.pragma('dart2js:noInline')
  static EmbededPluginConf getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<EmbededPluginConf>(create);
  static EmbededPluginConf? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get frameUrl => $_getSZ(0);
  @$pb.TagNumber(1)
  set frameUrl($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasFrameUrl() => $_has(0);
  @$pb.TagNumber(1)
  void clearFrameUrl() => clearField(1);
}

class ListRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'ListRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'infinimesh.plugins'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'namespace')
    ..hasRequiredFields = false
  ;

  ListRequest._() : super();
  factory ListRequest({
    $core.String? namespace,
  }) {
    final _result = create();
    if (namespace != null) {
      _result.namespace = namespace;
    }
    return _result;
  }
  factory ListRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ListRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ListRequest clone() => ListRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ListRequest copyWith(void Function(ListRequest) updates) => super.copyWith((message) => updates(message as ListRequest)) as ListRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ListRequest create() => ListRequest._();
  ListRequest createEmptyInstance() => create();
  static $pb.PbList<ListRequest> createRepeated() => $pb.PbList<ListRequest>();
  @$core.pragma('dart2js:noInline')
  static ListRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ListRequest>(create);
  static ListRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get namespace => $_getSZ(0);
  @$pb.TagNumber(1)
  set namespace($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasNamespace() => $_has(0);
  @$pb.TagNumber(1)
  void clearNamespace() => clearField(1);
}

