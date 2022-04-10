/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";

export const protobufPackage = "cosmonaut.cainiao.cainiao";

export interface MsgAddOrder {
  creator: string;
  sendName: string;
  sendAddress: string;
  sendTel: string;
  receiveName: string;
  receiveAddress: string;
  receiveTel: string;
}

export interface MsgAddOrderResponse {}

export interface MsgUpdateOrder {
  creator: string;
  station: string;
  id: number;
}

export interface MsgUpdateOrderResponse {}

export interface MsgUpdateOrderState {
  creator: string;
  id: string;
  station: string;
}

export interface MsgUpdateOrderStateResponse {}

const baseMsgAddOrder: object = {
  creator: "",
  sendName: "",
  sendAddress: "",
  sendTel: "",
  receiveName: "",
  receiveAddress: "",
  receiveTel: "",
};

export const MsgAddOrder = {
  encode(message: MsgAddOrder, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
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
    if (message.receiveName !== "") {
      writer.uint32(42).string(message.receiveName);
    }
    if (message.receiveAddress !== "") {
      writer.uint32(50).string(message.receiveAddress);
    }
    if (message.receiveTel !== "") {
      writer.uint32(58).string(message.receiveTel);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAddOrder {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgAddOrder } as MsgAddOrder;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
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
          message.receiveName = reader.string();
          break;
        case 6:
          message.receiveAddress = reader.string();
          break;
        case 7:
          message.receiveTel = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAddOrder {
    const message = { ...baseMsgAddOrder } as MsgAddOrder;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
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
    if (object.receiveName !== undefined && object.receiveName !== null) {
      message.receiveName = String(object.receiveName);
    } else {
      message.receiveName = "";
    }
    if (object.receiveAddress !== undefined && object.receiveAddress !== null) {
      message.receiveAddress = String(object.receiveAddress);
    } else {
      message.receiveAddress = "";
    }
    if (object.receiveTel !== undefined && object.receiveTel !== null) {
      message.receiveTel = String(object.receiveTel);
    } else {
      message.receiveTel = "";
    }
    return message;
  },

  toJSON(message: MsgAddOrder): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.sendName !== undefined && (obj.sendName = message.sendName);
    message.sendAddress !== undefined &&
      (obj.sendAddress = message.sendAddress);
    message.sendTel !== undefined && (obj.sendTel = message.sendTel);
    message.receiveName !== undefined &&
      (obj.receiveName = message.receiveName);
    message.receiveAddress !== undefined &&
      (obj.receiveAddress = message.receiveAddress);
    message.receiveTel !== undefined && (obj.receiveTel = message.receiveTel);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgAddOrder>): MsgAddOrder {
    const message = { ...baseMsgAddOrder } as MsgAddOrder;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
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
    if (object.receiveName !== undefined && object.receiveName !== null) {
      message.receiveName = object.receiveName;
    } else {
      message.receiveName = "";
    }
    if (object.receiveAddress !== undefined && object.receiveAddress !== null) {
      message.receiveAddress = object.receiveAddress;
    } else {
      message.receiveAddress = "";
    }
    if (object.receiveTel !== undefined && object.receiveTel !== null) {
      message.receiveTel = object.receiveTel;
    } else {
      message.receiveTel = "";
    }
    return message;
  },
};

const baseMsgAddOrderResponse: object = {};

export const MsgAddOrderResponse = {
  encode(_: MsgAddOrderResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAddOrderResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgAddOrderResponse } as MsgAddOrderResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgAddOrderResponse {
    const message = { ...baseMsgAddOrderResponse } as MsgAddOrderResponse;
    return message;
  },

  toJSON(_: MsgAddOrderResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgAddOrderResponse>): MsgAddOrderResponse {
    const message = { ...baseMsgAddOrderResponse } as MsgAddOrderResponse;
    return message;
  },
};

const baseMsgUpdateOrder: object = { creator: "", station: "", id: 0 };

export const MsgUpdateOrder = {
  encode(message: MsgUpdateOrder, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.station !== "") {
      writer.uint32(18).string(message.station);
    }
    if (message.id !== 0) {
      writer.uint32(24).uint64(message.id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateOrder {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUpdateOrder } as MsgUpdateOrder;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.station = reader.string();
          break;
        case 3:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateOrder {
    const message = { ...baseMsgUpdateOrder } as MsgUpdateOrder;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.station !== undefined && object.station !== null) {
      message.station = String(object.station);
    } else {
      message.station = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: MsgUpdateOrder): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.station !== undefined && (obj.station = message.station);
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgUpdateOrder>): MsgUpdateOrder {
    const message = { ...baseMsgUpdateOrder } as MsgUpdateOrder;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.station !== undefined && object.station !== null) {
      message.station = object.station;
    } else {
      message.station = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseMsgUpdateOrderResponse: object = {};

export const MsgUpdateOrderResponse = {
  encode(_: MsgUpdateOrderResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateOrderResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUpdateOrderResponse } as MsgUpdateOrderResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgUpdateOrderResponse {
    const message = { ...baseMsgUpdateOrderResponse } as MsgUpdateOrderResponse;
    return message;
  },

  toJSON(_: MsgUpdateOrderResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgUpdateOrderResponse>): MsgUpdateOrderResponse {
    const message = { ...baseMsgUpdateOrderResponse } as MsgUpdateOrderResponse;
    return message;
  },
};

const baseMsgUpdateOrderState: object = { creator: "", id: "", station: "" };

export const MsgUpdateOrderState = {
  encode(
    message: MsgUpdateOrderState,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== "") {
      writer.uint32(18).string(message.id);
    }
    if (message.station !== "") {
      writer.uint32(26).string(message.station);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateOrderState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUpdateOrderState } as MsgUpdateOrderState;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.id = reader.string();
          break;
        case 3:
          message.station = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateOrderState {
    const message = { ...baseMsgUpdateOrderState } as MsgUpdateOrderState;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = String(object.id);
    } else {
      message.id = "";
    }
    if (object.station !== undefined && object.station !== null) {
      message.station = String(object.station);
    } else {
      message.station = "";
    }
    return message;
  },

  toJSON(message: MsgUpdateOrderState): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = message.id);
    message.station !== undefined && (obj.station = message.station);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgUpdateOrderState>): MsgUpdateOrderState {
    const message = { ...baseMsgUpdateOrderState } as MsgUpdateOrderState;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = "";
    }
    if (object.station !== undefined && object.station !== null) {
      message.station = object.station;
    } else {
      message.station = "";
    }
    return message;
  },
};

const baseMsgUpdateOrderStateResponse: object = {};

export const MsgUpdateOrderStateResponse = {
  encode(
    _: MsgUpdateOrderStateResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgUpdateOrderStateResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgUpdateOrderStateResponse,
    } as MsgUpdateOrderStateResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgUpdateOrderStateResponse {
    const message = {
      ...baseMsgUpdateOrderStateResponse,
    } as MsgUpdateOrderStateResponse;
    return message;
  },

  toJSON(_: MsgUpdateOrderStateResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgUpdateOrderStateResponse>
  ): MsgUpdateOrderStateResponse {
    const message = {
      ...baseMsgUpdateOrderStateResponse,
    } as MsgUpdateOrderStateResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  AddOrder(request: MsgAddOrder): Promise<MsgAddOrderResponse>;
  UpdateOrder(request: MsgUpdateOrder): Promise<MsgUpdateOrderResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  UpdateOrderState(
    request: MsgUpdateOrderState
  ): Promise<MsgUpdateOrderStateResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  AddOrder(request: MsgAddOrder): Promise<MsgAddOrderResponse> {
    const data = MsgAddOrder.encode(request).finish();
    const promise = this.rpc.request(
      "cosmonaut.cainiao.cainiao.Msg",
      "AddOrder",
      data
    );
    return promise.then((data) => MsgAddOrderResponse.decode(new Reader(data)));
  }

  UpdateOrder(request: MsgUpdateOrder): Promise<MsgUpdateOrderResponse> {
    const data = MsgUpdateOrder.encode(request).finish();
    const promise = this.rpc.request(
      "cosmonaut.cainiao.cainiao.Msg",
      "UpdateOrder",
      data
    );
    return promise.then((data) =>
      MsgUpdateOrderResponse.decode(new Reader(data))
    );
  }

  UpdateOrderState(
    request: MsgUpdateOrderState
  ): Promise<MsgUpdateOrderStateResponse> {
    const data = MsgUpdateOrderState.encode(request).finish();
    const promise = this.rpc.request(
      "cosmonaut.cainiao.cainiao.Msg",
      "UpdateOrderState",
      data
    );
    return promise.then((data) =>
      MsgUpdateOrderStateResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

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
