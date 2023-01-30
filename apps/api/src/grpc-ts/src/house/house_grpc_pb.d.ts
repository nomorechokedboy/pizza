// package: houseService
// file: src/house/house.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from 'grpc'
import * as src_house_house_pb from '../../src/house/house_pb'

interface IHouseServiceService
  extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
  getBySize: IHouseServiceService_IGetBySize
  getHouse: IHouseServiceService_IGetHouse
  getHouses: IHouseServiceService_IGetHouses
}

interface IHouseServiceService_IGetBySize
  extends grpc.MethodDefinition<
    src_house_house_pb.HousesBySizeReq,
    src_house_house_pb.HousesBySizeRes
  > {
  path: '/houseService.HouseService/GetBySize'
  requestStream: false
  responseStream: false
  requestSerialize: grpc.serialize<src_house_house_pb.HousesBySizeReq>
  requestDeserialize: grpc.deserialize<src_house_house_pb.HousesBySizeReq>
  responseSerialize: grpc.serialize<src_house_house_pb.HousesBySizeRes>
  responseDeserialize: grpc.deserialize<src_house_house_pb.HousesBySizeRes>
}
interface IHouseServiceService_IGetHouse
  extends grpc.MethodDefinition<
    src_house_house_pb.HouseReq,
    src_house_house_pb.HouseRes
  > {
  path: '/houseService.HouseService/GetHouse'
  requestStream: false
  responseStream: false
  requestSerialize: grpc.serialize<src_house_house_pb.HouseReq>
  requestDeserialize: grpc.deserialize<src_house_house_pb.HouseReq>
  responseSerialize: grpc.serialize<src_house_house_pb.HouseRes>
  responseDeserialize: grpc.deserialize<src_house_house_pb.HouseRes>
}
interface IHouseServiceService_IGetHouses
  extends grpc.MethodDefinition<
    src_house_house_pb.HousesReq,
    src_house_house_pb.HousesRes
  > {
  path: '/houseService.HouseService/GetHouses'
  requestStream: false
  responseStream: false
  requestSerialize: grpc.serialize<src_house_house_pb.HousesReq>
  requestDeserialize: grpc.deserialize<src_house_house_pb.HousesReq>
  responseSerialize: grpc.serialize<src_house_house_pb.HousesRes>
  responseDeserialize: grpc.deserialize<src_house_house_pb.HousesRes>
}

export const HouseServiceService: IHouseServiceService

export interface IHouseServiceServer {
  getBySize: grpc.handleUnaryCall<
    src_house_house_pb.HousesBySizeReq,
    src_house_house_pb.HousesBySizeRes
  >
  getHouse: grpc.handleUnaryCall<
    src_house_house_pb.HouseReq,
    src_house_house_pb.HouseRes
  >
  getHouses: grpc.handleUnaryCall<
    src_house_house_pb.HousesReq,
    src_house_house_pb.HousesRes
  >
}

export interface IHouseServiceClient {
  getBySize(
    request: src_house_house_pb.HousesBySizeReq,
    callback: (
      error: grpc.ServiceError | null,
      response: src_house_house_pb.HousesBySizeRes
    ) => void
  ): grpc.ClientUnaryCall
  getBySize(
    request: src_house_house_pb.HousesBySizeReq,
    metadata: grpc.Metadata,
    callback: (
      error: grpc.ServiceError | null,
      response: src_house_house_pb.HousesBySizeRes
    ) => void
  ): grpc.ClientUnaryCall
  getBySize(
    request: src_house_house_pb.HousesBySizeReq,
    metadata: grpc.Metadata,
    options: Partial<grpc.CallOptions>,
    callback: (
      error: grpc.ServiceError | null,
      response: src_house_house_pb.HousesBySizeRes
    ) => void
  ): grpc.ClientUnaryCall
  getHouse(
    request: src_house_house_pb.HouseReq,
    callback: (
      error: grpc.ServiceError | null,
      response: src_house_house_pb.HouseRes
    ) => void
  ): grpc.ClientUnaryCall
  getHouse(
    request: src_house_house_pb.HouseReq,
    metadata: grpc.Metadata,
    callback: (
      error: grpc.ServiceError | null,
      response: src_house_house_pb.HouseRes
    ) => void
  ): grpc.ClientUnaryCall
  getHouse(
    request: src_house_house_pb.HouseReq,
    metadata: grpc.Metadata,
    options: Partial<grpc.CallOptions>,
    callback: (
      error: grpc.ServiceError | null,
      response: src_house_house_pb.HouseRes
    ) => void
  ): grpc.ClientUnaryCall
  getHouses(
    request: src_house_house_pb.HousesReq,
    callback: (
      error: grpc.ServiceError | null,
      response: src_house_house_pb.HousesRes
    ) => void
  ): grpc.ClientUnaryCall
  getHouses(
    request: src_house_house_pb.HousesReq,
    metadata: grpc.Metadata,
    callback: (
      error: grpc.ServiceError | null,
      response: src_house_house_pb.HousesRes
    ) => void
  ): grpc.ClientUnaryCall
  getHouses(
    request: src_house_house_pb.HousesReq,
    metadata: grpc.Metadata,
    options: Partial<grpc.CallOptions>,
    callback: (
      error: grpc.ServiceError | null,
      response: src_house_house_pb.HousesRes
    ) => void
  ): grpc.ClientUnaryCall
}

export class HouseServiceClient
  extends grpc.Client
  implements IHouseServiceClient
{
  constructor(
    address: string,
    credentials: grpc.ChannelCredentials,
    options?: object
  )
  public getBySize(
    request: src_house_house_pb.HousesBySizeReq,
    callback: (
      error: grpc.ServiceError | null,
      response: src_house_house_pb.HousesBySizeRes
    ) => void
  ): grpc.ClientUnaryCall
  public getBySize(
    request: src_house_house_pb.HousesBySizeReq,
    metadata: grpc.Metadata,
    callback: (
      error: grpc.ServiceError | null,
      response: src_house_house_pb.HousesBySizeRes
    ) => void
  ): grpc.ClientUnaryCall
  public getBySize(
    request: src_house_house_pb.HousesBySizeReq,
    metadata: grpc.Metadata,
    options: Partial<grpc.CallOptions>,
    callback: (
      error: grpc.ServiceError | null,
      response: src_house_house_pb.HousesBySizeRes
    ) => void
  ): grpc.ClientUnaryCall
  public getHouse(
    request: src_house_house_pb.HouseReq,
    callback: (
      error: grpc.ServiceError | null,
      response: src_house_house_pb.HouseRes
    ) => void
  ): grpc.ClientUnaryCall
  public getHouse(
    request: src_house_house_pb.HouseReq,
    metadata: grpc.Metadata,
    callback: (
      error: grpc.ServiceError | null,
      response: src_house_house_pb.HouseRes
    ) => void
  ): grpc.ClientUnaryCall
  public getHouse(
    request: src_house_house_pb.HouseReq,
    metadata: grpc.Metadata,
    options: Partial<grpc.CallOptions>,
    callback: (
      error: grpc.ServiceError | null,
      response: src_house_house_pb.HouseRes
    ) => void
  ): grpc.ClientUnaryCall
  public getHouses(
    request: src_house_house_pb.HousesReq,
    callback: (
      error: grpc.ServiceError | null,
      response: src_house_house_pb.HousesRes
    ) => void
  ): grpc.ClientUnaryCall
  public getHouses(
    request: src_house_house_pb.HousesReq,
    metadata: grpc.Metadata,
    callback: (
      error: grpc.ServiceError | null,
      response: src_house_house_pb.HousesRes
    ) => void
  ): grpc.ClientUnaryCall
  public getHouses(
    request: src_house_house_pb.HousesReq,
    metadata: grpc.Metadata,
    options: Partial<grpc.CallOptions>,
    callback: (
      error: grpc.ServiceError | null,
      response: src_house_house_pb.HousesRes
    ) => void
  ): grpc.ClientUnaryCall
}
