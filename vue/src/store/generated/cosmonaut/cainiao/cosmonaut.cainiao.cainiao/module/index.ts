// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgAddOrder } from "./types/cainiao/tx";
import { MsgUpdateOrder } from "./types/cainiao/tx";
import { MsgUpdateOrderState } from "./types/cainiao/tx";
import { MsgReceiveOrder } from "./types/cainiao/tx";


const types = [
  ["/cosmonaut.cainiao.cainiao.MsgAddOrder", MsgAddOrder],
  ["/cosmonaut.cainiao.cainiao.MsgUpdateOrder", MsgUpdateOrder],
  ["/cosmonaut.cainiao.cainiao.MsgUpdateOrderState", MsgUpdateOrderState],
  ["/cosmonaut.cainiao.cainiao.MsgReceiveOrder", MsgReceiveOrder],
  
];
export const MissingWalletError = new Error("wallet is required");

export const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
}

interface SignAndBroadcastOptions {
  fee: StdFee,
  memo?: string
}

const txClient = async (wallet: OfflineSigner, { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }) => {
  if (!wallet) throw MissingWalletError;
  let client;
  if (addr) {
    client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  }else{
    client = await SigningStargateClient.offline( wallet, { registry });
  }
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions = {fee: defaultFee, memo: ""}) => client.signAndBroadcast(address, msgs, fee,memo),
    msgAddOrder: (data: MsgAddOrder): EncodeObject => ({ typeUrl: "/cosmonaut.cainiao.cainiao.MsgAddOrder", value: MsgAddOrder.fromPartial( data ) }),
    msgUpdateOrder: (data: MsgUpdateOrder): EncodeObject => ({ typeUrl: "/cosmonaut.cainiao.cainiao.MsgUpdateOrder", value: MsgUpdateOrder.fromPartial( data ) }),
    msgUpdateOrderState: (data: MsgUpdateOrderState): EncodeObject => ({ typeUrl: "/cosmonaut.cainiao.cainiao.MsgUpdateOrderState", value: MsgUpdateOrderState.fromPartial( data ) }),
    msgReceiveOrder: (data: MsgReceiveOrder): EncodeObject => ({ typeUrl: "/cosmonaut.cainiao.cainiao.MsgReceiveOrder", value: MsgReceiveOrder.fromPartial( data ) }),
    
  };
};

interface QueryClientOptions {
  addr: string
}

const queryClient = async ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

export {
  txClient,
  queryClient,
};
