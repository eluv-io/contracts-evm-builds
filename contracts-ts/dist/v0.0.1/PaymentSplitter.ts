import {
  ContractTransaction,
  ContractInterface,
  BytesLike as Arrayish,
  BigNumberish,
} from 'ethers';
import { BigNumber } from '@ethersproject/bignumber';
import { EthersContractContextV5 } from 'ethereum-abi-types-generator';

export type ContractContext = EthersContractContextV5<
  PaymentSplitter,
  PaymentSplitterMethodNames,
  PaymentSplitterEventsContext,
  PaymentSplitterEvents
>;

export declare type EventFilter = {
  address?: string;
  topics?: Array<string>;
  fromBlock?: string | number;
  toBlock?: string | number;
};

export interface ContractTransactionOverrides {
  /**
   * The maximum units of gas for the transaction to use
   */
  gasLimit?: number;
  /**
   * The price (in wei) per unit of gas
   */
  gasPrice?: BigNumber | string | number | Promise<any>;
  /**
   * The nonce to use in the transaction
   */
  nonce?: number;
  /**
   * The amount to send with the transaction (i.e. msg.value)
   */
  value?: BigNumber | string | number | Promise<any>;
  /**
   * The chain ID (or network ID) to use
   */
  chainId?: number;
}

export interface ContractCallOverrides {
  /**
   * The address to execute the call as
   */
  from?: string;
  /**
   * The maximum units of gas for the transaction to use
   */
  gasLimit?: number;
}
export type PaymentSplitterEvents =
  | 'ERC20PaymentReleased'
  | 'PayeeAdded'
  | 'PaymentReceived'
  | 'PaymentReleased';
export interface PaymentSplitterEventsContext {
  ERC20PaymentReleased(...parameters: any): EventFilter;
  PayeeAdded(...parameters: any): EventFilter;
  PaymentReceived(...parameters: any): EventFilter;
  PaymentReleased(...parameters: any): EventFilter;
}
export type PaymentSplitterMethodNames =
  | 'new'
  | 'payee'
  | 'releasable'
  | 'releasable'
  | 'release'
  | 'release'
  | 'released'
  | 'released'
  | 'shares'
  | 'totalReleased'
  | 'totalReleased'
  | 'totalShares';
export interface ERC20PaymentReleasedEventEmittedResponse {
  token: string;
  to: string;
  amount: BigNumberish;
}
export interface PayeeAddedEventEmittedResponse {
  account: string;
  shares: BigNumberish;
}
export interface PaymentReceivedEventEmittedResponse {
  from: string;
  amount: BigNumberish;
}
export interface PaymentReleasedEventEmittedResponse {
  to: string;
  amount: BigNumberish;
}
export interface PaymentSplitter {
  /**
   * Payable: true
   * Constant: false
   * StateMutability: payable
   * Type: constructor
   * @param payees Type: address[], Indexed: false
   * @param shares_ Type: uint256[], Indexed: false
   */
  'new'(
    payees: string[],
    shares_: BigNumberish[],
    overrides?: ContractTransactionOverrides
  ): Promise<ContractTransaction>;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   * @param index Type: uint256, Indexed: false
   */
  payee(
    index: BigNumberish,
    overrides?: ContractCallOverrides
  ): Promise<string>;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   * @param account Type: address, Indexed: false
   */
  releasable(
    account: string,
    overrides?: ContractCallOverrides
  ): Promise<BigNumber>;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   * @param token Type: address, Indexed: false
   * @param account Type: address, Indexed: false
   */
  releasable(
    token: string,
    account: string,
    overrides?: ContractCallOverrides
  ): Promise<BigNumber>;
  /**
   * Payable: false
   * Constant: false
   * StateMutability: nonpayable
   * Type: function
   * @param account Type: address, Indexed: false
   */
  release(
    account: string,
    overrides?: ContractTransactionOverrides
  ): Promise<ContractTransaction>;
  /**
   * Payable: false
   * Constant: false
   * StateMutability: nonpayable
   * Type: function
   * @param token Type: address, Indexed: false
   * @param account Type: address, Indexed: false
   */
  release(
    token: string,
    account: string,
    overrides?: ContractTransactionOverrides
  ): Promise<ContractTransaction>;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   * @param token Type: address, Indexed: false
   * @param account Type: address, Indexed: false
   */
  released(
    token: string,
    account: string,
    overrides?: ContractCallOverrides
  ): Promise<BigNumber>;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   * @param account Type: address, Indexed: false
   */
  released(
    account: string,
    overrides?: ContractCallOverrides
  ): Promise<BigNumber>;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   * @param account Type: address, Indexed: false
   */
  shares(
    account: string,
    overrides?: ContractCallOverrides
  ): Promise<BigNumber>;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   * @param token Type: address, Indexed: false
   */
  totalReleased(
    token: string,
    overrides?: ContractCallOverrides
  ): Promise<BigNumber>;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   */
  totalReleased(overrides?: ContractCallOverrides): Promise<BigNumber>;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   */
  totalShares(overrides?: ContractCallOverrides): Promise<BigNumber>;
}
