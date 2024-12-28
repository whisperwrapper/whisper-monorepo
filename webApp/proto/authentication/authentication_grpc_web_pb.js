/**
 * @fileoverview gRPC-Web generated client stub for 
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.5.0
// 	protoc              v3.21.12
// source: authentication.proto


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js')

var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js')
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
 *   !proto.UserCredits,
 *   !proto.google.protobuf.Empty>}
 */
const methodDescriptor_ClientService_Register = new grpc.web.MethodDescriptor(
  '/ClientService/Register',
  grpc.web.MethodType.UNARY,
  proto.UserCredits,
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.UserCredits} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.UserCredits} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
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
 * @return {!Promise<!proto.google.protobuf.Empty>}
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


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.QueryParamethers,
 *   !proto.TranscriptionHistory>}
 */
const methodDescriptor_ClientService_GetTranscription = new grpc.web.MethodDescriptor(
  '/ClientService/GetTranscription',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.QueryParamethers,
  proto.TranscriptionHistory,
  /**
   * @param {!proto.QueryParamethers} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.TranscriptionHistory.deserializeBinary
);


/**
 * @param {!proto.QueryParamethers} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.TranscriptionHistory>}
 *     The XHR Node Readable Stream
 */
proto.ClientServiceClient.prototype.getTranscription =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/ClientService/GetTranscription',
      request,
      metadata || {},
      methodDescriptor_ClientService_GetTranscription);
};


/**
 * @param {!proto.QueryParamethers} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.TranscriptionHistory>}
 *     The XHR Node Readable Stream
 */
proto.ClientServicePromiseClient.prototype.getTranscription =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/ClientService/GetTranscription',
      request,
      metadata || {},
      methodDescriptor_ClientService_GetTranscription);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.NewTranscription,
 *   !proto.google.protobuf.Empty>}
 */
const methodDescriptor_ClientService_EditTranscription = new grpc.web.MethodDescriptor(
  '/ClientService/EditTranscription',
  grpc.web.MethodType.UNARY,
  proto.NewTranscription,
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.NewTranscription} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.NewTranscription} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.ClientServiceClient.prototype.editTranscription =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/ClientService/EditTranscription',
      request,
      metadata || {},
      methodDescriptor_ClientService_EditTranscription,
      callback);
};


/**
 * @param {!proto.NewTranscription} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     Promise that resolves to the response
 */
proto.ClientServicePromiseClient.prototype.editTranscription =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/ClientService/EditTranscription',
      request,
      metadata || {},
      methodDescriptor_ClientService_EditTranscription);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.Id,
 *   !proto.google.protobuf.Empty>}
 */
const methodDescriptor_ClientService_DeleteTranscription = new grpc.web.MethodDescriptor(
  '/ClientService/DeleteTranscription',
  grpc.web.MethodType.UNARY,
  proto.Id,
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.Id} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.Id} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.ClientServiceClient.prototype.deleteTranscription =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/ClientService/DeleteTranscription',
      request,
      metadata || {},
      methodDescriptor_ClientService_DeleteTranscription,
      callback);
};


/**
 * @param {!proto.Id} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     Promise that resolves to the response
 */
proto.ClientServicePromiseClient.prototype.deleteTranscription =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/ClientService/DeleteTranscription',
      request,
      metadata || {},
      methodDescriptor_ClientService_DeleteTranscription);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.QueryParamethers,
 *   !proto.TranslationHistory>}
 */
const methodDescriptor_ClientService_GetTranslation = new grpc.web.MethodDescriptor(
  '/ClientService/GetTranslation',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.QueryParamethers,
  proto.TranslationHistory,
  /**
   * @param {!proto.QueryParamethers} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.TranslationHistory.deserializeBinary
);


/**
 * @param {!proto.QueryParamethers} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.TranslationHistory>}
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
 * @param {!proto.QueryParamethers} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.TranslationHistory>}
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
 *   !proto.NewTranslation,
 *   !proto.google.protobuf.Empty>}
 */
const methodDescriptor_ClientService_EditTranslation = new grpc.web.MethodDescriptor(
  '/ClientService/EditTranslation',
  grpc.web.MethodType.UNARY,
  proto.NewTranslation,
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.NewTranslation} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.NewTranslation} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.ClientServiceClient.prototype.editTranslation =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/ClientService/EditTranslation',
      request,
      metadata || {},
      methodDescriptor_ClientService_EditTranslation,
      callback);
};


/**
 * @param {!proto.NewTranslation} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     Promise that resolves to the response
 */
proto.ClientServicePromiseClient.prototype.editTranslation =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/ClientService/EditTranslation',
      request,
      metadata || {},
      methodDescriptor_ClientService_EditTranslation);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.Id,
 *   !proto.google.protobuf.Empty>}
 */
const methodDescriptor_ClientService_DeleteTranslation = new grpc.web.MethodDescriptor(
  '/ClientService/DeleteTranslation',
  grpc.web.MethodType.UNARY,
  proto.Id,
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.Id} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.Id} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.ClientServiceClient.prototype.deleteTranslation =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/ClientService/DeleteTranslation',
      request,
      metadata || {},
      methodDescriptor_ClientService_DeleteTranslation,
      callback);
};


/**
 * @param {!proto.Id} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     Promise that resolves to the response
 */
proto.ClientServicePromiseClient.prototype.deleteTranslation =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/ClientService/DeleteTranslation',
      request,
      metadata || {},
      methodDescriptor_ClientService_DeleteTranslation);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.QueryParamethers,
 *   !proto.DiarizationHistory>}
 */
const methodDescriptor_ClientService_GetDiarization = new grpc.web.MethodDescriptor(
  '/ClientService/GetDiarization',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.QueryParamethers,
  proto.DiarizationHistory,
  /**
   * @param {!proto.QueryParamethers} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.DiarizationHistory.deserializeBinary
);


/**
 * @param {!proto.QueryParamethers} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.DiarizationHistory>}
 *     The XHR Node Readable Stream
 */
proto.ClientServiceClient.prototype.getDiarization =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/ClientService/GetDiarization',
      request,
      metadata || {},
      methodDescriptor_ClientService_GetDiarization);
};


/**
 * @param {!proto.QueryParamethers} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.DiarizationHistory>}
 *     The XHR Node Readable Stream
 */
proto.ClientServicePromiseClient.prototype.getDiarization =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/ClientService/GetDiarization',
      request,
      metadata || {},
      methodDescriptor_ClientService_GetDiarization);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.NewDiarization,
 *   !proto.google.protobuf.Empty>}
 */
const methodDescriptor_ClientService_EditDiarization = new grpc.web.MethodDescriptor(
  '/ClientService/EditDiarization',
  grpc.web.MethodType.UNARY,
  proto.NewDiarization,
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.NewDiarization} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.NewDiarization} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.ClientServiceClient.prototype.editDiarization =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/ClientService/EditDiarization',
      request,
      metadata || {},
      methodDescriptor_ClientService_EditDiarization,
      callback);
};


/**
 * @param {!proto.NewDiarization} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     Promise that resolves to the response
 */
proto.ClientServicePromiseClient.prototype.editDiarization =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/ClientService/EditDiarization',
      request,
      metadata || {},
      methodDescriptor_ClientService_EditDiarization);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.Id,
 *   !proto.google.protobuf.Empty>}
 */
const methodDescriptor_ClientService_DeleteDiarization = new grpc.web.MethodDescriptor(
  '/ClientService/DeleteDiarization',
  grpc.web.MethodType.UNARY,
  proto.Id,
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.Id} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.Id} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.ClientServiceClient.prototype.deleteDiarization =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/ClientService/DeleteDiarization',
      request,
      metadata || {},
      methodDescriptor_ClientService_DeleteDiarization,
      callback);
};


/**
 * @param {!proto.Id} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     Promise that resolves to the response
 */
proto.ClientServicePromiseClient.prototype.deleteDiarization =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/ClientService/DeleteDiarization',
      request,
      metadata || {},
      methodDescriptor_ClientService_DeleteDiarization);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.QueryParamethers,
 *   !proto.Combined>}
 */
const methodDescriptor_ClientService_GetTranscriptionAndDiarization = new grpc.web.MethodDescriptor(
  '/ClientService/GetTranscriptionAndDiarization',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.QueryParamethers,
  proto.Combined,
  /**
   * @param {!proto.QueryParamethers} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.Combined.deserializeBinary
);


/**
 * @param {!proto.QueryParamethers} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.Combined>}
 *     The XHR Node Readable Stream
 */
proto.ClientServiceClient.prototype.getTranscriptionAndDiarization =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/ClientService/GetTranscriptionAndDiarization',
      request,
      metadata || {},
      methodDescriptor_ClientService_GetTranscriptionAndDiarization);
};


/**
 * @param {!proto.QueryParamethers} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.Combined>}
 *     The XHR Node Readable Stream
 */
proto.ClientServicePromiseClient.prototype.getTranscriptionAndDiarization =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/ClientService/GetTranscriptionAndDiarization',
      request,
      metadata || {},
      methodDescriptor_ClientService_GetTranscriptionAndDiarization);
};


module.exports = proto;

