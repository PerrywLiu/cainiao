/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "cosmonaut.cainiao.cainiao";

export interface Orders {
  id: number;
  sendName: string;
  sendAddress: string;
  sendTel: string;
  targetName: string;
  targetAddress: string;
  targetTel: string;
  state: string;
  station: string;
  locationRouter: string[];
}

const baseOrders: object = {
  id: 0,
  sendName: "",
  sendAddress: "",
  sendTel: "",
  targetName: "",
  targetAddress: "",
  targetTel: "",
  state: "",
  station: "",
  locationRouter: "",
};

export const Orders = {
  encode(message: Orders, writer: Writer = Writer.create()): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.sendName !== "") {
      writer.uint32(18).string(message.sendName);
    }
    if (message.sendAddress !== "") {
      writer.uint32(26).string(message.sendAddress);
    }
    if (message.sendTel !== "") {
      writer.uint32(34).string(message.sendTel);
    }
    if (message.targetName !== "") {
      writer.uint32(42).string(message.targetName);
    }
    if (message.targetAddress !== "") {
      writer.uint32(50).string(message.targetAddress);
    }
    if (message.targetTel !== "") {
      writer.uint32(58).string(message.targetTel);
    }
    if (message.state !== "") {
      writer.uint32(66).string(message.state);
    }
    if (message.station !== "") {
      writer.uint32(74).string(message.station);
    }
    for (const v of message.locationRouter) {
      writer.uint32(82).string(v!);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Orders {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseOrders } as Orders;
    message.locationRouter = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.sendName = reader.string();
          break;
        case 3:
          message.sendAddress = reader.string();
          break;
        case 4:
          message.sendTel = reader.string();
          break;
        case 5:
          message.targetName = reader.string();
          break;
        case 6:
          message.targetAddress = reader.string();
          break;
        case 7:
          message.targetTel = reader.string();
          break;
        case 8:
          message.state = reader.string();
          break;
        case 9:
          message.station = reader.string();
          break;
        case 10:
          message.locationRouter.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Orders {
    const message = { ...baseOrders } as Orders;
    message.locationRouter = [];
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    if (object.sendName !== undefined && object.sendName !== null) {
      message.sendName = String(object.sendName);
    } else {
      message.sendName = "";
    }
    if (object.sendAddress !== undefined && object.sendAddress !== null) {
      message.sendAddress = String(object.sendAddress);
    } else {
      message.sendAddress = "";
    }
    if (object.sendTel !== undefined && object.sendTel !== null) {
      message.sendTel = String(object.sendTel);
    } else {
      message.sendTel = "";
    }
    if (object.targetName !== undefined && object.targetName !== null) {
      message.targetName = String(object.targetName);
    } else {
      message.targetName = "";
    }
    if (object.targetAddress !== undefined && object.targetAddress !== null) {
      message.targetAddress = String(object.targetAddress);
    } else {
      message.targetAddress = "";
    }
    if (object.targetTel !== undefined && object.targetTel !== null) {
      message.targetTel = String(object.targetTel);
    } else {
      message.targetTel = "";
    }
    if (object.state !== undefined && object.state !== null) {
      message.state = String(object.state);
    } else {
      message.state = "";
    }
    if (object.station !== undefined && object.station !== null) {
      message.station = String(object.station);
    } else {
      message.station = "";
    }
    if (object.locationRouter !== undefined && object.locationRouter !== null) {
      for (const e of object.locationRouter) {
        message.locationRouter.push(String(e));
      }
    }
    return message;
  },

  toJSON(message: Orders): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.sendName !== undefined && (obj.sendName = message.sendName);
    message.sendAddress !== undefined &&
      (obj.sendAddress = message.sendAddress);
    message.sendTel !== undefined && (obj.sendTel = message.sendTel);
    message.targetName !== undefined && (obj.targetName = message.targetName);
    message.targetAddress !== undefined &&
      (obj.targetAddress = message.targetAddress);
    message.targetTel !== undefined && (obj.targetTel = message.targetTel);
    message.state !== undefined && (obj.state = message.state);
    message.station !== undefined && (obj.station = message.station);
    if (message.locationRouter) {
      obj.locationRouter = message.locationRouter.map((e) => e);
    } else {
      obj.locationRouter = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<Orders>): Orders {
    const message = { ...baseOrders } as Orders;
    message.locationRouter = [];
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    if (object.sendName !== undefined && object.sendName !== null) {
      message.sendName = object.sendName;
    } else {
      message.sendName = "";
    }
    if (object.sendAddress !== undefined && object.sendAddress !== null) {
      message.sendAddress = object.sendAddress;
    } else {
      message.sendAddress = "";
    }
    if (object.sendTel !== undefined && object.sendTel !== null) {
      message.sendTel = object.sendTel;
    } else {
      message.sendTel = "";
    }
    if (object.targetName !== undefined && object.targetName !== null) {
      message.targetName = object.targetName;
    } else {
      message.targetName = "";
    }
    if (object.targetAddress !== undefined && object.targetAddress !== null) {
      message.targetAddress = object.targetAddress;
    } else {
      message.targetAddress = "";
    }
    if (object.targetTel !== undefined && object.targetTel !== null) {
      message.targetTel = object.targetTel;
    } else {
      message.targetTel = "";
    }
    if (object.state !== undefined && object.state !== null) {
      message.state = object.state;
    } else {
      message.state = "";
    }
    if (object.station !== undefined && object.station !== null) {
      message.station = object.station;
    } else {
      message.station = "";
    }
    if (object.locationRouter !== undefined && object.locationRouter !== null) {
      for (const e of object.locationRouter) {
        message.locationRouter.push(e);
      }
    }
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
