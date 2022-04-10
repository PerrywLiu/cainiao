/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

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

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  AddOrder(request: MsgAddOrder): Promise<MsgAddOrderResponse>;
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
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

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
