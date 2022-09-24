// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var src_house_house_pb = require('../../src/house/house_pb.js');

function serialize_houseService_HouseReq(arg) {
  if (!(arg instanceof src_house_house_pb.HouseReq)) {
    throw new Error('Expected argument of type houseService.HouseReq');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_houseService_HouseReq(buffer_arg) {
  return src_house_house_pb.HouseReq.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_houseService_HouseRes(arg) {
  if (!(arg instanceof src_house_house_pb.HouseRes)) {
    throw new Error('Expected argument of type houseService.HouseRes');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_houseService_HouseRes(buffer_arg) {
  return src_house_house_pb.HouseRes.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_houseService_HousesBySizeReq(arg) {
  if (!(arg instanceof src_house_house_pb.HousesBySizeReq)) {
    throw new Error('Expected argument of type houseService.HousesBySizeReq');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_houseService_HousesBySizeReq(buffer_arg) {
  return src_house_house_pb.HousesBySizeReq.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_houseService_HousesBySizeRes(arg) {
  if (!(arg instanceof src_house_house_pb.HousesBySizeRes)) {
    throw new Error('Expected argument of type houseService.HousesBySizeRes');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_houseService_HousesBySizeRes(buffer_arg) {
  return src_house_house_pb.HousesBySizeRes.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_houseService_HousesReq(arg) {
  if (!(arg instanceof src_house_house_pb.HousesReq)) {
    throw new Error('Expected argument of type houseService.HousesReq');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_houseService_HousesReq(buffer_arg) {
  return src_house_house_pb.HousesReq.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_houseService_HousesRes(arg) {
  if (!(arg instanceof src_house_house_pb.HousesRes)) {
    throw new Error('Expected argument of type houseService.HousesRes');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_houseService_HousesRes(buffer_arg) {
  return src_house_house_pb.HousesRes.deserializeBinary(new Uint8Array(buffer_arg));
}


var HouseServiceService = exports.HouseServiceService = {
  getBySize: {
    path: '/houseService.HouseService/GetBySize',
    requestStream: false,
    responseStream: false,
    requestType: src_house_house_pb.HousesBySizeReq,
    responseType: src_house_house_pb.HousesBySizeRes,
    requestSerialize: serialize_houseService_HousesBySizeReq,
    requestDeserialize: deserialize_houseService_HousesBySizeReq,
    responseSerialize: serialize_houseService_HousesBySizeRes,
    responseDeserialize: deserialize_houseService_HousesBySizeRes,
  },
  getHouse: {
    path: '/houseService.HouseService/GetHouse',
    requestStream: false,
    responseStream: false,
    requestType: src_house_house_pb.HouseReq,
    responseType: src_house_house_pb.HouseRes,
    requestSerialize: serialize_houseService_HouseReq,
    requestDeserialize: deserialize_houseService_HouseReq,
    responseSerialize: serialize_houseService_HouseRes,
    responseDeserialize: deserialize_houseService_HouseRes,
  },
  getHouses: {
    path: '/houseService.HouseService/GetHouses',
    requestStream: false,
    responseStream: false,
    requestType: src_house_house_pb.HousesReq,
    responseType: src_house_house_pb.HousesRes,
    requestSerialize: serialize_houseService_HousesReq,
    requestDeserialize: deserialize_houseService_HousesReq,
    responseSerialize: serialize_houseService_HousesRes,
    responseDeserialize: deserialize_houseService_HousesRes,
  },
};

exports.HouseServiceClient = grpc.makeGenericClientConstructor(HouseServiceService);
