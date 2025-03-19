///
//  Generated code. Do not modify.
//  source: node/node.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'node.pb.dart' as $1;
import 'accounts/accounts.pb.dart' as $2;
import 'access/access.pb.dart' as $3;
import 'namespaces/namespaces.pb.dart' as $4;
import 'devices/devices.pb.dart' as $5;
import '../shadow/shadow.pb.dart' as $0;
export 'node.pb.dart';

class AccountsServiceClient extends $grpc.Client {
  static final _$token = $grpc.ClientMethod<$1.TokenRequest, $1.TokenResponse>(
      '/infinimesh.node.AccountsService/Token',
      ($1.TokenRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $1.TokenResponse.fromBuffer(value));
  static final _$get = $grpc.ClientMethod<$2.Account, $2.Account>(
      '/infinimesh.node.AccountsService/Get',
      ($2.Account value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $2.Account.fromBuffer(value));
  static final _$list = $grpc.ClientMethod<$1.EmptyMessage, $2.Accounts>(
      '/infinimesh.node.AccountsService/List',
      ($1.EmptyMessage value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $2.Accounts.fromBuffer(value));
  static final _$create =
      $grpc.ClientMethod<$2.CreateRequest, $2.CreateResponse>(
          '/infinimesh.node.AccountsService/Create',
          ($2.CreateRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $2.CreateResponse.fromBuffer(value));
  static final _$update = $grpc.ClientMethod<$2.Account, $2.Account>(
      '/infinimesh.node.AccountsService/Update',
      ($2.Account value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $2.Account.fromBuffer(value));
  static final _$toggle = $grpc.ClientMethod<$2.Account, $2.Account>(
      '/infinimesh.node.AccountsService/Toggle',
      ($2.Account value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $2.Account.fromBuffer(value));
  static final _$delete = $grpc.ClientMethod<$2.Account, $1.DeleteResponse>(
      '/infinimesh.node.AccountsService/Delete',
      ($2.Account value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $1.DeleteResponse.fromBuffer(value));
  static final _$deletables = $grpc.ClientMethod<$2.Account, $3.Nodes>(
      '/infinimesh.node.AccountsService/Deletables',
      ($2.Account value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $3.Nodes.fromBuffer(value));
  static final _$setCredentials =
      $grpc.ClientMethod<$1.SetCredentialsRequest, $1.SetCredentialsResponse>(
          '/infinimesh.node.AccountsService/SetCredentials',
          ($1.SetCredentialsRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $1.SetCredentialsResponse.fromBuffer(value));

  AccountsServiceClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options, interceptors: interceptors);

  $grpc.ResponseFuture<$1.TokenResponse> token($1.TokenRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$token, request, options: options);
  }

  $grpc.ResponseFuture<$2.Account> get($2.Account request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$get, request, options: options);
  }

  $grpc.ResponseFuture<$2.Accounts> list($1.EmptyMessage request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$list, request, options: options);
  }

  $grpc.ResponseFuture<$2.CreateResponse> create($2.CreateRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$create, request, options: options);
  }

  $grpc.ResponseFuture<$2.Account> update($2.Account request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$update, request, options: options);
  }

  $grpc.ResponseFuture<$2.Account> toggle($2.Account request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$toggle, request, options: options);
  }

  $grpc.ResponseFuture<$1.DeleteResponse> delete($2.Account request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$delete, request, options: options);
  }

  $grpc.ResponseFuture<$3.Nodes> deletables($2.Account request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$deletables, request, options: options);
  }

  $grpc.ResponseFuture<$1.SetCredentialsResponse> setCredentials(
      $1.SetCredentialsRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$setCredentials, request, options: options);
  }
}

abstract class AccountsServiceBase extends $grpc.Service {
  $core.String get $name => 'infinimesh.node.AccountsService';

  AccountsServiceBase() {
    $addMethod($grpc.ServiceMethod<$1.TokenRequest, $1.TokenResponse>(
        'Token',
        token_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $1.TokenRequest.fromBuffer(value),
        ($1.TokenResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$2.Account, $2.Account>(
        'Get',
        get_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $2.Account.fromBuffer(value),
        ($2.Account value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.EmptyMessage, $2.Accounts>(
        'List',
        list_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $1.EmptyMessage.fromBuffer(value),
        ($2.Accounts value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$2.CreateRequest, $2.CreateResponse>(
        'Create',
        create_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $2.CreateRequest.fromBuffer(value),
        ($2.CreateResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$2.Account, $2.Account>(
        'Update',
        update_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $2.Account.fromBuffer(value),
        ($2.Account value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$2.Account, $2.Account>(
        'Toggle',
        toggle_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $2.Account.fromBuffer(value),
        ($2.Account value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$2.Account, $1.DeleteResponse>(
        'Delete',
        delete_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $2.Account.fromBuffer(value),
        ($1.DeleteResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$2.Account, $3.Nodes>(
        'Deletables',
        deletables_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $2.Account.fromBuffer(value),
        ($3.Nodes value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.SetCredentialsRequest,
            $1.SetCredentialsResponse>(
        'SetCredentials',
        setCredentials_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $1.SetCredentialsRequest.fromBuffer(value),
        ($1.SetCredentialsResponse value) => value.writeToBuffer()));
  }

  $async.Future<$1.TokenResponse> token_Pre(
      $grpc.ServiceCall call, $async.Future<$1.TokenRequest> request) async {
    return token(call, await request);
  }

  $async.Future<$2.Account> get_Pre(
      $grpc.ServiceCall call, $async.Future<$2.Account> request) async {
    return get(call, await request);
  }

  $async.Future<$2.Accounts> list_Pre(
      $grpc.ServiceCall call, $async.Future<$1.EmptyMessage> request) async {
    return list(call, await request);
  }

  $async.Future<$2.CreateResponse> create_Pre(
      $grpc.ServiceCall call, $async.Future<$2.CreateRequest> request) async {
    return create(call, await request);
  }

  $async.Future<$2.Account> update_Pre(
      $grpc.ServiceCall call, $async.Future<$2.Account> request) async {
    return update(call, await request);
  }

  $async.Future<$2.Account> toggle_Pre(
      $grpc.ServiceCall call, $async.Future<$2.Account> request) async {
    return toggle(call, await request);
  }

  $async.Future<$1.DeleteResponse> delete_Pre(
      $grpc.ServiceCall call, $async.Future<$2.Account> request) async {
    return delete(call, await request);
  }

  $async.Future<$3.Nodes> deletables_Pre(
      $grpc.ServiceCall call, $async.Future<$2.Account> request) async {
    return deletables(call, await request);
  }

  $async.Future<$1.SetCredentialsResponse> setCredentials_Pre(
      $grpc.ServiceCall call,
      $async.Future<$1.SetCredentialsRequest> request) async {
    return setCredentials(call, await request);
  }

  $async.Future<$1.TokenResponse> token(
      $grpc.ServiceCall call, $1.TokenRequest request);
  $async.Future<$2.Account> get($grpc.ServiceCall call, $2.Account request);
  $async.Future<$2.Accounts> list(
      $grpc.ServiceCall call, $1.EmptyMessage request);
  $async.Future<$2.CreateResponse> create(
      $grpc.ServiceCall call, $2.CreateRequest request);
  $async.Future<$2.Account> update($grpc.ServiceCall call, $2.Account request);
  $async.Future<$2.Account> toggle($grpc.ServiceCall call, $2.Account request);
  $async.Future<$1.DeleteResponse> delete(
      $grpc.ServiceCall call, $2.Account request);
  $async.Future<$3.Nodes> deletables(
      $grpc.ServiceCall call, $2.Account request);
  $async.Future<$1.SetCredentialsResponse> setCredentials(
      $grpc.ServiceCall call, $1.SetCredentialsRequest request);
}

class NamespacesServiceClient extends $grpc.Client {
  static final _$get = $grpc.ClientMethod<$4.Namespace, $4.Namespace>(
      '/infinimesh.node.NamespacesService/Get',
      ($4.Namespace value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $4.Namespace.fromBuffer(value));
  static final _$list = $grpc.ClientMethod<$1.EmptyMessage, $4.Namespaces>(
      '/infinimesh.node.NamespacesService/List',
      ($1.EmptyMessage value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $4.Namespaces.fromBuffer(value));
  static final _$create = $grpc.ClientMethod<$4.Namespace, $4.Namespace>(
      '/infinimesh.node.NamespacesService/Create',
      ($4.Namespace value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $4.Namespace.fromBuffer(value));
  static final _$update = $grpc.ClientMethod<$4.Namespace, $4.Namespace>(
      '/infinimesh.node.NamespacesService/Update',
      ($4.Namespace value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $4.Namespace.fromBuffer(value));
  static final _$delete = $grpc.ClientMethod<$4.Namespace, $1.DeleteResponse>(
      '/infinimesh.node.NamespacesService/Delete',
      ($4.Namespace value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $1.DeleteResponse.fromBuffer(value));
  static final _$deletables = $grpc.ClientMethod<$4.Namespace, $3.Nodes>(
      '/infinimesh.node.NamespacesService/Deletables',
      ($4.Namespace value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $3.Nodes.fromBuffer(value));
  static final _$joins = $grpc.ClientMethod<$4.Namespace, $2.Accounts>(
      '/infinimesh.node.NamespacesService/Joins',
      ($4.Namespace value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $2.Accounts.fromBuffer(value));
  static final _$join = $grpc.ClientMethod<$1.JoinRequest, $2.Accounts>(
      '/infinimesh.node.NamespacesService/Join',
      ($1.JoinRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $2.Accounts.fromBuffer(value));

  NamespacesServiceClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options, interceptors: interceptors);

  $grpc.ResponseFuture<$4.Namespace> get($4.Namespace request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$get, request, options: options);
  }

  $grpc.ResponseFuture<$4.Namespaces> list($1.EmptyMessage request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$list, request, options: options);
  }

  $grpc.ResponseFuture<$4.Namespace> create($4.Namespace request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$create, request, options: options);
  }

  $grpc.ResponseFuture<$4.Namespace> update($4.Namespace request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$update, request, options: options);
  }

  $grpc.ResponseFuture<$1.DeleteResponse> delete($4.Namespace request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$delete, request, options: options);
  }

  $grpc.ResponseFuture<$3.Nodes> deletables($4.Namespace request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$deletables, request, options: options);
  }

  $grpc.ResponseFuture<$2.Accounts> joins($4.Namespace request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$joins, request, options: options);
  }

  $grpc.ResponseFuture<$2.Accounts> join($1.JoinRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$join, request, options: options);
  }
}

abstract class NamespacesServiceBase extends $grpc.Service {
  $core.String get $name => 'infinimesh.node.NamespacesService';

  NamespacesServiceBase() {
    $addMethod($grpc.ServiceMethod<$4.Namespace, $4.Namespace>(
        'Get',
        get_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $4.Namespace.fromBuffer(value),
        ($4.Namespace value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.EmptyMessage, $4.Namespaces>(
        'List',
        list_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $1.EmptyMessage.fromBuffer(value),
        ($4.Namespaces value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$4.Namespace, $4.Namespace>(
        'Create',
        create_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $4.Namespace.fromBuffer(value),
        ($4.Namespace value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$4.Namespace, $4.Namespace>(
        'Update',
        update_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $4.Namespace.fromBuffer(value),
        ($4.Namespace value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$4.Namespace, $1.DeleteResponse>(
        'Delete',
        delete_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $4.Namespace.fromBuffer(value),
        ($1.DeleteResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$4.Namespace, $3.Nodes>(
        'Deletables',
        deletables_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $4.Namespace.fromBuffer(value),
        ($3.Nodes value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$4.Namespace, $2.Accounts>(
        'Joins',
        joins_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $4.Namespace.fromBuffer(value),
        ($2.Accounts value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.JoinRequest, $2.Accounts>(
        'Join',
        join_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $1.JoinRequest.fromBuffer(value),
        ($2.Accounts value) => value.writeToBuffer()));
  }

  $async.Future<$4.Namespace> get_Pre(
      $grpc.ServiceCall call, $async.Future<$4.Namespace> request) async {
    return get(call, await request);
  }

  $async.Future<$4.Namespaces> list_Pre(
      $grpc.ServiceCall call, $async.Future<$1.EmptyMessage> request) async {
    return list(call, await request);
  }

  $async.Future<$4.Namespace> create_Pre(
      $grpc.ServiceCall call, $async.Future<$4.Namespace> request) async {
    return create(call, await request);
  }

  $async.Future<$4.Namespace> update_Pre(
      $grpc.ServiceCall call, $async.Future<$4.Namespace> request) async {
    return update(call, await request);
  }

  $async.Future<$1.DeleteResponse> delete_Pre(
      $grpc.ServiceCall call, $async.Future<$4.Namespace> request) async {
    return delete(call, await request);
  }

  $async.Future<$3.Nodes> deletables_Pre(
      $grpc.ServiceCall call, $async.Future<$4.Namespace> request) async {
    return deletables(call, await request);
  }

  $async.Future<$2.Accounts> joins_Pre(
      $grpc.ServiceCall call, $async.Future<$4.Namespace> request) async {
    return joins(call, await request);
  }

  $async.Future<$2.Accounts> join_Pre(
      $grpc.ServiceCall call, $async.Future<$1.JoinRequest> request) async {
    return join(call, await request);
  }

  $async.Future<$4.Namespace> get($grpc.ServiceCall call, $4.Namespace request);
  $async.Future<$4.Namespaces> list(
      $grpc.ServiceCall call, $1.EmptyMessage request);
  $async.Future<$4.Namespace> create(
      $grpc.ServiceCall call, $4.Namespace request);
  $async.Future<$4.Namespace> update(
      $grpc.ServiceCall call, $4.Namespace request);
  $async.Future<$1.DeleteResponse> delete(
      $grpc.ServiceCall call, $4.Namespace request);
  $async.Future<$3.Nodes> deletables(
      $grpc.ServiceCall call, $4.Namespace request);
  $async.Future<$2.Accounts> joins(
      $grpc.ServiceCall call, $4.Namespace request);
  $async.Future<$2.Accounts> join(
      $grpc.ServiceCall call, $1.JoinRequest request);
}

class DevicesServiceClient extends $grpc.Client {
  static final _$get = $grpc.ClientMethod<$5.Device, $5.Device>(
      '/infinimesh.node.DevicesService/Get',
      ($5.Device value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $5.Device.fromBuffer(value));
  static final _$list = $grpc.ClientMethod<$1.EmptyMessage, $5.Devices>(
      '/infinimesh.node.DevicesService/List',
      ($1.EmptyMessage value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $5.Devices.fromBuffer(value));
  static final _$create =
      $grpc.ClientMethod<$5.CreateRequest, $5.CreateResponse>(
          '/infinimesh.node.DevicesService/Create',
          ($5.CreateRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $5.CreateResponse.fromBuffer(value));
  static final _$update = $grpc.ClientMethod<$5.Device, $5.Device>(
      '/infinimesh.node.DevicesService/Update',
      ($5.Device value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $5.Device.fromBuffer(value));
  static final _$delete = $grpc.ClientMethod<$5.Device, $1.DeleteResponse>(
      '/infinimesh.node.DevicesService/Delete',
      ($5.Device value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $1.DeleteResponse.fromBuffer(value));
  static final _$toggle = $grpc.ClientMethod<$5.Device, $5.Device>(
      '/infinimesh.node.DevicesService/Toggle',
      ($5.Device value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $5.Device.fromBuffer(value));
  static final _$toggleBasic = $grpc.ClientMethod<$5.Device, $5.Device>(
      '/infinimesh.node.DevicesService/ToggleBasic',
      ($5.Device value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $5.Device.fromBuffer(value));
  static final _$makeDevicesToken =
      $grpc.ClientMethod<$1.DevicesTokenRequest, $1.TokenResponse>(
          '/infinimesh.node.DevicesService/MakeDevicesToken',
          ($1.DevicesTokenRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $1.TokenResponse.fromBuffer(value));
  static final _$getByToken = $grpc.ClientMethod<$5.Device, $5.Device>(
      '/infinimesh.node.DevicesService/GetByToken',
      ($5.Device value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $5.Device.fromBuffer(value));
  static final _$getByFingerprint =
      $grpc.ClientMethod<$5.GetByFingerprintRequest, $5.Device>(
          '/infinimesh.node.DevicesService/GetByFingerprint',
          ($5.GetByFingerprintRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $5.Device.fromBuffer(value));

  DevicesServiceClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options, interceptors: interceptors);

  $grpc.ResponseFuture<$5.Device> get($5.Device request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$get, request, options: options);
  }

  $grpc.ResponseFuture<$5.Devices> list($1.EmptyMessage request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$list, request, options: options);
  }

  $grpc.ResponseFuture<$5.CreateResponse> create($5.CreateRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$create, request, options: options);
  }

  $grpc.ResponseFuture<$5.Device> update($5.Device request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$update, request, options: options);
  }

  $grpc.ResponseFuture<$1.DeleteResponse> delete($5.Device request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$delete, request, options: options);
  }

  $grpc.ResponseFuture<$5.Device> toggle($5.Device request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$toggle, request, options: options);
  }

  $grpc.ResponseFuture<$5.Device> toggleBasic($5.Device request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$toggleBasic, request, options: options);
  }

  $grpc.ResponseFuture<$1.TokenResponse> makeDevicesToken(
      $1.DevicesTokenRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$makeDevicesToken, request, options: options);
  }

  $grpc.ResponseFuture<$5.Device> getByToken($5.Device request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$getByToken, request, options: options);
  }

  $grpc.ResponseFuture<$5.Device> getByFingerprint(
      $5.GetByFingerprintRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$getByFingerprint, request, options: options);
  }
}

abstract class DevicesServiceBase extends $grpc.Service {
  $core.String get $name => 'infinimesh.node.DevicesService';

  DevicesServiceBase() {
    $addMethod($grpc.ServiceMethod<$5.Device, $5.Device>(
        'Get',
        get_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $5.Device.fromBuffer(value),
        ($5.Device value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.EmptyMessage, $5.Devices>(
        'List',
        list_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $1.EmptyMessage.fromBuffer(value),
        ($5.Devices value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$5.CreateRequest, $5.CreateResponse>(
        'Create',
        create_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $5.CreateRequest.fromBuffer(value),
        ($5.CreateResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$5.Device, $5.Device>(
        'Update',
        update_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $5.Device.fromBuffer(value),
        ($5.Device value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$5.Device, $1.DeleteResponse>(
        'Delete',
        delete_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $5.Device.fromBuffer(value),
        ($1.DeleteResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$5.Device, $5.Device>(
        'Toggle',
        toggle_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $5.Device.fromBuffer(value),
        ($5.Device value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$5.Device, $5.Device>(
        'ToggleBasic',
        toggleBasic_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $5.Device.fromBuffer(value),
        ($5.Device value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.DevicesTokenRequest, $1.TokenResponse>(
        'MakeDevicesToken',
        makeDevicesToken_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $1.DevicesTokenRequest.fromBuffer(value),
        ($1.TokenResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$5.Device, $5.Device>(
        'GetByToken',
        getByToken_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $5.Device.fromBuffer(value),
        ($5.Device value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$5.GetByFingerprintRequest, $5.Device>(
        'GetByFingerprint',
        getByFingerprint_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $5.GetByFingerprintRequest.fromBuffer(value),
        ($5.Device value) => value.writeToBuffer()));
  }

  $async.Future<$5.Device> get_Pre(
      $grpc.ServiceCall call, $async.Future<$5.Device> request) async {
    return get(call, await request);
  }

  $async.Future<$5.Devices> list_Pre(
      $grpc.ServiceCall call, $async.Future<$1.EmptyMessage> request) async {
    return list(call, await request);
  }

  $async.Future<$5.CreateResponse> create_Pre(
      $grpc.ServiceCall call, $async.Future<$5.CreateRequest> request) async {
    return create(call, await request);
  }

  $async.Future<$5.Device> update_Pre(
      $grpc.ServiceCall call, $async.Future<$5.Device> request) async {
    return update(call, await request);
  }

  $async.Future<$1.DeleteResponse> delete_Pre(
      $grpc.ServiceCall call, $async.Future<$5.Device> request) async {
    return delete(call, await request);
  }

  $async.Future<$5.Device> toggle_Pre(
      $grpc.ServiceCall call, $async.Future<$5.Device> request) async {
    return toggle(call, await request);
  }

  $async.Future<$5.Device> toggleBasic_Pre(
      $grpc.ServiceCall call, $async.Future<$5.Device> request) async {
    return toggleBasic(call, await request);
  }

  $async.Future<$1.TokenResponse> makeDevicesToken_Pre($grpc.ServiceCall call,
      $async.Future<$1.DevicesTokenRequest> request) async {
    return makeDevicesToken(call, await request);
  }

  $async.Future<$5.Device> getByToken_Pre(
      $grpc.ServiceCall call, $async.Future<$5.Device> request) async {
    return getByToken(call, await request);
  }

  $async.Future<$5.Device> getByFingerprint_Pre($grpc.ServiceCall call,
      $async.Future<$5.GetByFingerprintRequest> request) async {
    return getByFingerprint(call, await request);
  }

  $async.Future<$5.Device> get($grpc.ServiceCall call, $5.Device request);
  $async.Future<$5.Devices> list(
      $grpc.ServiceCall call, $1.EmptyMessage request);
  $async.Future<$5.CreateResponse> create(
      $grpc.ServiceCall call, $5.CreateRequest request);
  $async.Future<$5.Device> update($grpc.ServiceCall call, $5.Device request);
  $async.Future<$1.DeleteResponse> delete(
      $grpc.ServiceCall call, $5.Device request);
  $async.Future<$5.Device> toggle($grpc.ServiceCall call, $5.Device request);
  $async.Future<$5.Device> toggleBasic(
      $grpc.ServiceCall call, $5.Device request);
  $async.Future<$1.TokenResponse> makeDevicesToken(
      $grpc.ServiceCall call, $1.DevicesTokenRequest request);
  $async.Future<$5.Device> getByToken(
      $grpc.ServiceCall call, $5.Device request);
  $async.Future<$5.Device> getByFingerprint(
      $grpc.ServiceCall call, $5.GetByFingerprintRequest request);
}

class ShadowServiceClient extends $grpc.Client {
  static final _$get = $grpc.ClientMethod<$0.GetRequest, $0.GetResponse>(
      '/infinimesh.node.ShadowService/Get',
      ($0.GetRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.GetResponse.fromBuffer(value));
  static final _$patch = $grpc.ClientMethod<$0.Shadow, $0.Shadow>(
      '/infinimesh.node.ShadowService/Patch',
      ($0.Shadow value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Shadow.fromBuffer(value));
  static final _$remove = $grpc.ClientMethod<$0.RemoveRequest, $0.Shadow>(
      '/infinimesh.node.ShadowService/Remove',
      ($0.RemoveRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Shadow.fromBuffer(value));
  static final _$streamShadow =
      $grpc.ClientMethod<$0.StreamShadowRequest, $0.Shadow>(
          '/infinimesh.node.ShadowService/StreamShadow',
          ($0.StreamShadowRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $0.Shadow.fromBuffer(value));
  static final _$streamShadowSync =
      $grpc.ClientMethod<$0.StreamShadowRequest, $0.Shadow>(
          '/infinimesh.node.ShadowService/StreamShadowSync',
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

  $grpc.ResponseStream<$0.Shadow> streamShadowSync(
      $0.StreamShadowRequest request,
      {$grpc.CallOptions? options}) {
    return $createStreamingCall(
        _$streamShadowSync, $async.Stream.fromIterable([request]),
        options: options);
  }
}

abstract class ShadowServiceBase extends $grpc.Service {
  $core.String get $name => 'infinimesh.node.ShadowService';

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
    $addMethod($grpc.ServiceMethod<$0.StreamShadowRequest, $0.Shadow>(
        'StreamShadowSync',
        streamShadowSync_Pre,
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

  $async.Stream<$0.Shadow> streamShadowSync_Pre($grpc.ServiceCall call,
      $async.Future<$0.StreamShadowRequest> request) async* {
    yield* streamShadowSync(call, await request);
  }

  $async.Future<$0.GetResponse> get(
      $grpc.ServiceCall call, $0.GetRequest request);
  $async.Future<$0.Shadow> patch($grpc.ServiceCall call, $0.Shadow request);
  $async.Future<$0.Shadow> remove(
      $grpc.ServiceCall call, $0.RemoveRequest request);
  $async.Stream<$0.Shadow> streamShadow(
      $grpc.ServiceCall call, $0.StreamShadowRequest request);
  $async.Stream<$0.Shadow> streamShadowSync(
      $grpc.ServiceCall call, $0.StreamShadowRequest request);
}

class InternalServiceClient extends $grpc.Client {
  static final _$getLDAPProviders =
      $grpc.ClientMethod<$1.EmptyMessage, $1.LDAPProviders>(
          '/infinimesh.node.InternalService/GetLDAPProviders',
          ($1.EmptyMessage value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $1.LDAPProviders.fromBuffer(value));

  InternalServiceClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options, interceptors: interceptors);

  $grpc.ResponseFuture<$1.LDAPProviders> getLDAPProviders(
      $1.EmptyMessage request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$getLDAPProviders, request, options: options);
  }
}

abstract class InternalServiceBase extends $grpc.Service {
  $core.String get $name => 'infinimesh.node.InternalService';

  InternalServiceBase() {
    $addMethod($grpc.ServiceMethod<$1.EmptyMessage, $1.LDAPProviders>(
        'GetLDAPProviders',
        getLDAPProviders_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $1.EmptyMessage.fromBuffer(value),
        ($1.LDAPProviders value) => value.writeToBuffer()));
  }

  $async.Future<$1.LDAPProviders> getLDAPProviders_Pre(
      $grpc.ServiceCall call, $async.Future<$1.EmptyMessage> request) async {
    return getLDAPProviders(call, await request);
  }

  $async.Future<$1.LDAPProviders> getLDAPProviders(
      $grpc.ServiceCall call, $1.EmptyMessage request);
}
