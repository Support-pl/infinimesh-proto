///
//  Generated code. Do not modify.
//  source: handsfree/handsfree.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'handsfree.pb.dart' as $0;
export 'handsfree.pb.dart';

class HandsfreeServiceClient extends $grpc.Client {
  static final _$send = $grpc.ClientMethod<$0.ControlPacket, $0.ControlPacket>(
      '/infinimesh.handsfree.HandsfreeService/Send',
      ($0.ControlPacket value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.ControlPacket.fromBuffer(value));
  static final _$connect =
      $grpc.ClientMethod<$0.ConnectionRequest, $0.ControlPacket>(
          '/infinimesh.handsfree.HandsfreeService/Connect',
          ($0.ConnectionRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $0.ControlPacket.fromBuffer(value));

  HandsfreeServiceClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options, interceptors: interceptors);

  $grpc.ResponseFuture<$0.ControlPacket> send($0.ControlPacket request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$send, request, options: options);
  }

  $grpc.ResponseStream<$0.ControlPacket> connect($0.ConnectionRequest request,
      {$grpc.CallOptions? options}) {
    return $createStreamingCall(
        _$connect, $async.Stream.fromIterable([request]),
        options: options);
  }
}

abstract class HandsfreeServiceBase extends $grpc.Service {
  $core.String get $name => 'infinimesh.handsfree.HandsfreeService';

  HandsfreeServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.ControlPacket, $0.ControlPacket>(
        'Send',
        send_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.ControlPacket.fromBuffer(value),
        ($0.ControlPacket value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.ConnectionRequest, $0.ControlPacket>(
        'Connect',
        connect_Pre,
        false,
        true,
        ($core.List<$core.int> value) => $0.ConnectionRequest.fromBuffer(value),
        ($0.ControlPacket value) => value.writeToBuffer()));
  }

  $async.Future<$0.ControlPacket> send_Pre(
      $grpc.ServiceCall call, $async.Future<$0.ControlPacket> request) async {
    return send(call, await request);
  }

  $async.Stream<$0.ControlPacket> connect_Pre($grpc.ServiceCall call,
      $async.Future<$0.ConnectionRequest> request) async* {
    yield* connect(call, await request);
  }

  $async.Future<$0.ControlPacket> send(
      $grpc.ServiceCall call, $0.ControlPacket request);
  $async.Stream<$0.ControlPacket> connect(
      $grpc.ServiceCall call, $0.ConnectionRequest request);
}
