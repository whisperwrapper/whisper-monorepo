/**
 * @fileoverview gRPC-Web generated client stub for 
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.5.0
// 	protoc              v5.26.0
// source: authentication.proto


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = require('./authentication_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.ClientServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname.replace(/\/+$/, '');

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.ClientServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname.replace(/\/+$/, '');

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.UserCredits,
 *   !proto.LoginResponse>}
 */
const methodDescriptor_ClientService_Login = new grpc.web.MethodDescriptor(
  '/ClientService/Login',
  grpc.web.MethodType.UNARY,
  proto.UserCredits,
  proto.LoginResponse,
  /**
   * @param {!proto.UserCredits} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.LoginResponse.deserializeBinary
);


/**
 * @param {!proto.UserCredits} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.LoginResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.LoginResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.ClientServiceClient.prototype.login =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/ClientService/Login',
      request,
      metadata || {},
      methodDescriptor_ClientService_Login,
      callback);
};


/**
 * @param {!proto.UserCredits} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.LoginResponse>}
 *     Promise that resolves to the response
 */
proto.ClientServicePromiseClient.prototype.login =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/ClientService/Login',
      request,
      metadata || {},
      methodDescriptor_ClientService_Login);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.Empty,
 *   !proto.TextHistory>}
 */
const methodDescriptor_ClientService_GetTranslation = new grpc.web.MethodDescriptor(
  '/ClientService/GetTranslation',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.Empty,
  proto.TextHistory,
  /**
   * @param {!proto.Empty} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.TextHistory.deserializeBinary
);


/**
 * @param {!proto.Empty} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.TextHistory>}
 *     The XHR Node Readable Stream
 */
proto.ClientServiceClient.prototype.getTranslation =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/ClientService/GetTranslation',
      request,
      metadata || {},
      methodDescriptor_ClientService_GetTranslation);
};


/**
 * @param {!proto.Empty} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.TextHistory>}
 *     The XHR Node Readable Stream
 */
proto.ClientServicePromiseClient.prototype.getTranslation =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/ClientService/GetTranslation',
      request,
      metadata || {},
      methodDescriptor_ClientService_GetTranslation);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.UserCredits,
 *   !proto.StatusResponse>}
 */
const methodDescriptor_ClientService_Register = new grpc.web.MethodDescriptor(
  '/ClientService/Register',
  grpc.web.MethodType.UNARY,
  proto.UserCredits,
  proto.StatusResponse,
  /**
   * @param {!proto.UserCredits} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.StatusResponse.deserializeBinary
);


/**
 * @param {!proto.UserCredits} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.StatusResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.StatusResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.ClientServiceClient.prototype.register =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/ClientService/Register',
      request,
      metadata || {},
      methodDescriptor_ClientService_Register,
      callback);
};


/**
 * @param {!proto.UserCredits} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.StatusResponse>}
 *     Promise that resolves to the response
 */
proto.ClientServicePromiseClient.prototype.register =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/ClientService/Register',
      request,
      metadata || {},
      methodDescriptor_ClientService_Register);
};


module.exports = proto;

