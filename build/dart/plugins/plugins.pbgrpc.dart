///
//  Generated code. Do not modify.
//  source: plugins/plugins.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'plugins.pb.dart' as $0;
export 'plugins.pb.dart';

class PluginsServiceClient extends $grpc.Client {
  static final _$get = $grpc.ClientMethod<$0.Plugin, $0.Plugin>(
      '/infinimesh.plugins.PluginsService/Get',
      ($0.Plugin value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Plugin.fromBuffer(value));
  static final _$list = $grpc.ClientMethod<$0.ListRequest, $0.Plugins>(
      '/infinimesh.plugins.PluginsService/List',
      ($0.ListRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Plugins.fromBuffer(value));
  static final _$create = $grpc.ClientMethod<$0.Plugin, $0.Plugin>(
      '/infinimesh.plugins.PluginsService/Create',
      ($0.Plugin value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Plugin.fromBuffer(value));
  static final _$update = $grpc.ClientMethod<$0.Plugin, $0.Plugin>(
      '/infinimesh.plugins.PluginsService/Update',
      ($0.Plugin value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Plugin.fromBuffer(value));
  static final _$delete = $grpc.ClientMethod<$0.Plugin, $0.Plugin>(
      '/infinimesh.plugins.PluginsService/Delete',
      ($0.Plugin value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Plugin.fromBuffer(value));

  PluginsServiceClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options, interceptors: interceptors);

  $grpc.ResponseFuture<$0.Plugin> get($0.Plugin request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$get, request, options: options);
  }

  $grpc.ResponseFuture<$0.Plugins> list($0.ListRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$list, request, options: options);
  }

  $grpc.ResponseFuture<$0.Plugin> create($0.Plugin request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$create, request, options: options);
  }

  $grpc.ResponseFuture<$0.Plugin> update($0.Plugin request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$update, request, options: options);
  }

  $grpc.ResponseFuture<$0.Plugin> delete($0.Plugin request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$delete, request, options: options);
  }
}

abstract class PluginsServiceBase extends $grpc.Service {
  $core.String get $name => 'infinimesh.plugins.PluginsService';

  PluginsServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.Plugin, $0.Plugin>(
        'Get',
        get_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.Plugin.fromBuffer(value),
        ($0.Plugin value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.ListRequest, $0.Plugins>(
        'List',
        list_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.ListRequest.fromBuffer(value),
        ($0.Plugins value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.Plugin, $0.Plugin>(
        'Create',
        create_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.Plugin.fromBuffer(value),
        ($0.Plugin value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.Plugin, $0.Plugin>(
        'Update',
        update_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.Plugin.fromBuffer(value),
        ($0.Plugin value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.Plugin, $0.Plugin>(
        'Delete',
        delete_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.Plugin.fromBuffer(value),
        ($0.Plugin value) => value.writeToBuffer()));
  }

  $async.Future<$0.Plugin> get_Pre(
      $grpc.ServiceCall call, $async.Future<$0.Plugin> request) async {
    return get(call, await request);
  }

  $async.Future<$0.Plugins> list_Pre(
      $grpc.ServiceCall call, $async.Future<$0.ListRequest> request) async {
    return list(call, await request);
  }

  $async.Future<$0.Plugin> create_Pre(
      $grpc.ServiceCall call, $async.Future<$0.Plugin> request) async {
    return create(call, await request);
  }

  $async.Future<$0.Plugin> update_Pre(
      $grpc.ServiceCall call, $async.Future<$0.Plugin> request) async {
    return update(call, await request);
  }

  $async.Future<$0.Plugin> delete_Pre(
      $grpc.ServiceCall call, $async.Future<$0.Plugin> request) async {
    return delete(call, await request);
  }

  $async.Future<$0.Plugin> get($grpc.ServiceCall call, $0.Plugin request);
  $async.Future<$0.Plugins> list(
      $grpc.ServiceCall call, $0.ListRequest request);
  $async.Future<$0.Plugin> create($grpc.ServiceCall call, $0.Plugin request);
  $async.Future<$0.Plugin> update($grpc.ServiceCall call, $0.Plugin request);
  $async.Future<$0.Plugin> delete($grpc.ServiceCall call, $0.Plugin request);
}
