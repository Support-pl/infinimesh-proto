///
//  Generated code. Do not modify.
//  source: shadow/shadow.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'shadow.pb.dart' as $0;
export 'shadow.pb.dart';

class ShadowServiceClient extends $grpc.Client {
  static final _$get = $grpc.ClientMethod<$0.GetRequest, $0.GetResponse>(
      '/infinimesh.shadow.ShadowService/Get',
      ($0.GetRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.GetResponse.fromBuffer(value));
  static final _$patch = $grpc.ClientMethod<$0.Shadow, $0.Shadow>(
      '/infinimesh.shadow.ShadowService/Patch',
      ($0.Shadow value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Shadow.fromBuffer(value));
  static final _$remove = $grpc.ClientMethod<$0.RemoveRequest, $0.Shadow>(
      '/infinimesh.shadow.ShadowService/Remove',
      ($0.RemoveRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Shadow.fromBuffer(value));
  static final _$streamShadow =
      $grpc.ClientMethod<$0.StreamShadowRequest, $0.Shadow>(
          '/infinimesh.shadow.ShadowService/StreamShadow',
          ($0.StreamShadowRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $0.Shadow.fromBuffer(value));

  ShadowServiceClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options, interceptors: interceptors);

  $grpc.ResponseFuture<$0.GetResponse> get($0.GetRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$get, request, options: options);
  }

  $grpc.ResponseFuture<$0.Shadow> patch($0.Shadow request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$patch, request, options: options);
  }

  $grpc.ResponseFuture<$0.Shadow> remove($0.RemoveRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$remove, request, options: options);
  }

  $grpc.ResponseStream<$0.Shadow> streamShadow($0.StreamShadowRequest request,
      {$grpc.CallOptions? options}) {
    return $createStreamingCall(
        _$streamShadow, $async.Stream.fromIterable([request]),
        options: options);
  }
}

abstract class ShadowServiceBase extends $grpc.Service {
  $core.String get $name => 'infinimesh.shadow.ShadowService';

  ShadowServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.GetRequest, $0.GetResponse>(
        'Get',
        get_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.GetRequest.fromBuffer(value),
        ($0.GetResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.Shadow, $0.Shadow>(
        'Patch',
        patch_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.Shadow.fromBuffer(value),
        ($0.Shadow value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.RemoveRequest, $0.Shadow>(
        'Remove',
        remove_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.RemoveRequest.fromBuffer(value),
        ($0.Shadow value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.StreamShadowRequest, $0.Shadow>(
        'StreamShadow',
        streamShadow_Pre,
        false,
        true,
        ($core.List<$core.int> value) =>
            $0.StreamShadowRequest.fromBuffer(value),
        ($0.Shadow value) => value.writeToBuffer()));
  }

  $async.Future<$0.GetResponse> get_Pre(
      $grpc.ServiceCall call, $async.Future<$0.GetRequest> request) async {
    return get(call, await request);
  }

  $async.Future<$0.Shadow> patch_Pre(
      $grpc.ServiceCall call, $async.Future<$0.Shadow> request) async {
    return patch(call, await request);
  }

  $async.Future<$0.Shadow> remove_Pre(
      $grpc.ServiceCall call, $async.Future<$0.RemoveRequest> request) async {
    return remove(call, await request);
  }

  $async.Stream<$0.Shadow> streamShadow_Pre($grpc.ServiceCall call,
      $async.Future<$0.StreamShadowRequest> request) async* {
    yield* streamShadow(call, await request);
  }

  $async.Future<$0.GetResponse> get(
      $grpc.ServiceCall call, $0.GetRequest request);
  $async.Future<$0.Shadow> patch($grpc.ServiceCall call, $0.Shadow request);
  $async.Future<$0.Shadow> remove(
      $grpc.ServiceCall call, $0.RemoveRequest request);
  $async.Stream<$0.Shadow> streamShadow(
      $grpc.ServiceCall call, $0.StreamShadowRequest request);
}
