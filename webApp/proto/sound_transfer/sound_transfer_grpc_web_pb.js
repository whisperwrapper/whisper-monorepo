/**
 * @fileoverview gRPC-Web generated client stub for 
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.5.0
// 	protoc              v5.26.0
// source: sound_transfer.proto


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = require('./sound_transfer_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.SoundServiceClient =
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
proto.SoundServicePromiseClient =
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
 *   !proto.TextMessage,
 *   !proto.TextMessage>}
 */
const methodDescriptor_SoundService_TestConnection = new grpc.web.MethodDescriptor(
  '/SoundService/TestConnection',
  grpc.web.MethodType.UNARY,
  proto.TextMessage,
  proto.TextMessage,
  /**
   * @param {!proto.TextMessage} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.TextMessage.deserializeBinary
);


/**
 * @param {!proto.TextMessage} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.TextMessage)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.TextMessage>|undefined}
 *     The XHR Node Readable Stream
 */
proto.SoundServiceClient.prototype.testConnection =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/SoundService/TestConnection',
      request,
      metadata || {},
      methodDescriptor_SoundService_TestConnection,
      callback);
};


/**
 * @param {!proto.TextMessage} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.TextMessage>}
 *     Promise that resolves to the response
 */
proto.SoundServicePromiseClient.prototype.testConnection =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/SoundService/TestConnection',
      request,
      metadata || {},
      methodDescriptor_SoundService_TestConnection);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.TranscriptionRequest,
 *   !proto.SoundResponse>}
 */
const methodDescriptor_SoundService_TranscribeFile = new grpc.web.MethodDescriptor(
  '/SoundService/TranscribeFile',
  grpc.web.MethodType.UNARY,
  proto.TranscriptionRequest,
  proto.SoundResponse,
  /**
   * @param {!proto.TranscriptionRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.SoundResponse.deserializeBinary
);


/**
 * @param {!proto.TranscriptionRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.SoundResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.SoundResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.SoundServiceClient.prototype.transcribeFile =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/SoundService/TranscribeFile',
      request,
      metadata || {},
      methodDescriptor_SoundService_TranscribeFile,
      callback);
};


/**
 * @param {!proto.TranscriptionRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.SoundResponse>}
 *     Promise that resolves to the response
 */
proto.SoundServicePromiseClient.prototype.transcribeFile =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/SoundService/TranscribeFile',
      request,
      metadata || {},
      methodDescriptor_SoundService_TranscribeFile);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.TranscriptionRequest,
 *   !proto.SoundStreamResponse>}
 */
const methodDescriptor_SoundService_TranscribeLiveWeb = new grpc.web.MethodDescriptor(
  '/SoundService/TranscribeLiveWeb',
  grpc.web.MethodType.UNARY,
  proto.TranscriptionRequest,
  proto.SoundStreamResponse,
  /**
   * @param {!proto.TranscriptionRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.SoundStreamResponse.deserializeBinary
);


/**
 * @param {!proto.TranscriptionRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.SoundStreamResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.SoundStreamResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.SoundServiceClient.prototype.transcribeLiveWeb =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/SoundService/TranscribeLiveWeb',
      request,
      metadata || {},
      methodDescriptor_SoundService_TranscribeLiveWeb,
      callback);
};


/**
 * @param {!proto.TranscriptionRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.SoundStreamResponse>}
 *     Promise that resolves to the response
 */
proto.SoundServicePromiseClient.prototype.transcribeLiveWeb =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/SoundService/TranscribeLiveWeb',
      request,
      metadata || {},
      methodDescriptor_SoundService_TranscribeLiveWeb);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.TranslationRequest,
 *   !proto.SoundResponse>}
 */
const methodDescriptor_SoundService_TranslateFile = new grpc.web.MethodDescriptor(
  '/SoundService/TranslateFile',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.TranslationRequest,
  proto.SoundResponse,
  /**
   * @param {!proto.TranslationRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.SoundResponse.deserializeBinary
);


/**
 * @param {!proto.TranslationRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.SoundResponse>}
 *     The XHR Node Readable Stream
 */
proto.SoundServiceClient.prototype.translateFile =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/SoundService/TranslateFile',
      request,
      metadata || {},
      methodDescriptor_SoundService_TranslateFile);
};


/**
 * @param {!proto.TranslationRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.SoundResponse>}
 *     The XHR Node Readable Stream
 */
proto.SoundServicePromiseClient.prototype.translateFile =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/SoundService/TranslateFile',
      request,
      metadata || {},
      methodDescriptor_SoundService_TranslateFile);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.TranscriptionRequest,
 *   !proto.SpeakerAndLineResponse>}
 */
const methodDescriptor_SoundService_DiarizateFile = new grpc.web.MethodDescriptor(
  '/SoundService/DiarizateFile',
  grpc.web.MethodType.UNARY,
  proto.TranscriptionRequest,
  proto.SpeakerAndLineResponse,
  /**
   * @param {!proto.TranscriptionRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.SpeakerAndLineResponse.deserializeBinary
);


/**
 * @param {!proto.TranscriptionRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.SpeakerAndLineResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.SpeakerAndLineResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.SoundServiceClient.prototype.diarizateFile =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/SoundService/DiarizateFile',
      request,
      metadata || {},
      methodDescriptor_SoundService_DiarizateFile,
      callback);
};


/**
 * @param {!proto.TranscriptionRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.SpeakerAndLineResponse>}
 *     Promise that resolves to the response
 */
proto.SoundServicePromiseClient.prototype.diarizateFile =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/SoundService/DiarizateFile',
      request,
      metadata || {},
      methodDescriptor_SoundService_DiarizateFile);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.TextAndId,
 *   !proto.TextMessage>}
 */
const methodDescriptor_SoundService_TranslateText = new grpc.web.MethodDescriptor(
  '/SoundService/TranslateText',
  grpc.web.MethodType.UNARY,
  proto.TextAndId,
  proto.TextMessage,
  /**
   * @param {!proto.TextAndId} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.TextMessage.deserializeBinary
);


/**
 * @param {!proto.TextAndId} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.TextMessage)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.TextMessage>|undefined}
 *     The XHR Node Readable Stream
 */
proto.SoundServiceClient.prototype.translateText =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/SoundService/TranslateText',
      request,
      metadata || {},
      methodDescriptor_SoundService_TranslateText,
      callback);
};


/**
 * @param {!proto.TextAndId} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.TextMessage>}
 *     Promise that resolves to the response
 */
proto.SoundServicePromiseClient.prototype.translateText =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/SoundService/TranslateText',
      request,
      metadata || {},
      methodDescriptor_SoundService_TranslateText);
};


module.exports = proto;

