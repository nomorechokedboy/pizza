// package: houseService
// file: src/house/house.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from 'google-protobuf'

export class House extends jspb.Message {
  getId(): number
  setId(value: number): House
  getStreet(): string
  setStreet(value: string): House
  getNumber(): string
  setNumber(value: string): House
  getSquare(): number
  setSquare(value: number): House

  serializeBinary(): Uint8Array
  toObject(includeInstance?: boolean): House.AsObject
  static toObject(includeInstance: boolean, msg: House): House.AsObject
  static extensions: { [key: number]: jspb.ExtensionFieldInfo<jspb.Message> }
  static extensionsBinary: {
    [key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>
  }
  static serializeBinaryToWriter(
    message: House,
    writer: jspb.BinaryWriter
  ): void
  static deserializeBinary(bytes: Uint8Array): House
  static deserializeBinaryFromReader(
    message: House,
    reader: jspb.BinaryReader
  ): House
}

export namespace House {
  export type AsObject = {
    id: number
    street: string
    number: string
    square: number
  }
}

export class HousesBySizeReq extends jspb.Message {
  getMinsquare(): number
  setMinsquare(value: number): HousesBySizeReq

  serializeBinary(): Uint8Array
  toObject(includeInstance?: boolean): HousesBySizeReq.AsObject
  static toObject(
    includeInstance: boolean,
    msg: HousesBySizeReq
  ): HousesBySizeReq.AsObject
  static extensions: { [key: number]: jspb.ExtensionFieldInfo<jspb.Message> }
  static extensionsBinary: {
    [key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>
  }
  static serializeBinaryToWriter(
    message: HousesBySizeReq,
    writer: jspb.BinaryWriter
  ): void
  static deserializeBinary(bytes: Uint8Array): HousesBySizeReq
  static deserializeBinaryFromReader(
    message: HousesBySizeReq,
    reader: jspb.BinaryReader
  ): HousesBySizeReq
}

export namespace HousesBySizeReq {
  export type AsObject = {
    minsquare: number
  }
}

export class HousesBySizeRes extends jspb.Message {
  clearIdsList(): void
  getIdsList(): Array<number>
  setIdsList(value: Array<number>): HousesBySizeRes
  addIds(value: number, index?: number): number

  serializeBinary(): Uint8Array
  toObject(includeInstance?: boolean): HousesBySizeRes.AsObject
  static toObject(
    includeInstance: boolean,
    msg: HousesBySizeRes
  ): HousesBySizeRes.AsObject
  static extensions: { [key: number]: jspb.ExtensionFieldInfo<jspb.Message> }
  static extensionsBinary: {
    [key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>
  }
  static serializeBinaryToWriter(
    message: HousesBySizeRes,
    writer: jspb.BinaryWriter
  ): void
  static deserializeBinary(bytes: Uint8Array): HousesBySizeRes
  static deserializeBinaryFromReader(
    message: HousesBySizeRes,
    reader: jspb.BinaryReader
  ): HousesBySizeRes
}

export namespace HousesBySizeRes {
  export type AsObject = {
    idsList: Array<number>
  }
}

export class HouseReq extends jspb.Message {
  getId(): number
  setId(value: number): HouseReq

  serializeBinary(): Uint8Array
  toObject(includeInstance?: boolean): HouseReq.AsObject
  static toObject(includeInstance: boolean, msg: HouseReq): HouseReq.AsObject
  static extensions: { [key: number]: jspb.ExtensionFieldInfo<jspb.Message> }
  static extensionsBinary: {
    [key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>
  }
  static serializeBinaryToWriter(
    message: HouseReq,
    writer: jspb.BinaryWriter
  ): void
  static deserializeBinary(bytes: Uint8Array): HouseReq
  static deserializeBinaryFromReader(
    message: HouseReq,
    reader: jspb.BinaryReader
  ): HouseReq
}

export namespace HouseReq {
  export type AsObject = {
    id: number
  }
}

export class HouseRes extends jspb.Message {
  hasHouse(): boolean
  clearHouse(): void
  getHouse(): House | undefined
  setHouse(value?: House): HouseRes

  serializeBinary(): Uint8Array
  toObject(includeInstance?: boolean): HouseRes.AsObject
  static toObject(includeInstance: boolean, msg: HouseRes): HouseRes.AsObject
  static extensions: { [key: number]: jspb.ExtensionFieldInfo<jspb.Message> }
  static extensionsBinary: {
    [key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>
  }
  static serializeBinaryToWriter(
    message: HouseRes,
    writer: jspb.BinaryWriter
  ): void
  static deserializeBinary(bytes: Uint8Array): HouseRes
  static deserializeBinaryFromReader(
    message: HouseRes,
    reader: jspb.BinaryReader
  ): HouseRes
}

export namespace HouseRes {
  export type AsObject = {
    house?: House.AsObject
  }
}

export class HousesReq extends jspb.Message {
  clearIdList(): void
  getIdList(): Array<number>
  setIdList(value: Array<number>): HousesReq
  addId(value: number, index?: number): number

  serializeBinary(): Uint8Array
  toObject(includeInstance?: boolean): HousesReq.AsObject
  static toObject(includeInstance: boolean, msg: HousesReq): HousesReq.AsObject
  static extensions: { [key: number]: jspb.ExtensionFieldInfo<jspb.Message> }
  static extensionsBinary: {
    [key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>
  }
  static serializeBinaryToWriter(
    message: HousesReq,
    writer: jspb.BinaryWriter
  ): void
  static deserializeBinary(bytes: Uint8Array): HousesReq
  static deserializeBinaryFromReader(
    message: HousesReq,
    reader: jspb.BinaryReader
  ): HousesReq
}

export namespace HousesReq {
  export type AsObject = {
    idList: Array<number>
  }
}

export class HousesRes extends jspb.Message {
  clearHousesList(): void
  getHousesList(): Array<House>
  setHousesList(value: Array<House>): HousesRes
  addHouses(value?: House, index?: number): House

  serializeBinary(): Uint8Array
  toObject(includeInstance?: boolean): HousesRes.AsObject
  static toObject(includeInstance: boolean, msg: HousesRes): HousesRes.AsObject
  static extensions: { [key: number]: jspb.ExtensionFieldInfo<jspb.Message> }
  static extensionsBinary: {
    [key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>
  }
  static serializeBinaryToWriter(
    message: HousesRes,
    writer: jspb.BinaryWriter
  ): void
  static deserializeBinary(bytes: Uint8Array): HousesRes
  static deserializeBinaryFromReader(
    message: HousesRes,
    reader: jspb.BinaryReader
  ): HousesRes
}

export namespace HousesRes {
  export type AsObject = {
    housesList: Array<House.AsObject>
  }
}
