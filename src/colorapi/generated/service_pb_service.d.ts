// package: colorapi
// file: service.proto

import * as service_pb from "./service_pb";
import {grpc} from "@improbable-eng/grpc-web";

type ColorsPickAllColors = {
  readonly methodName: string;
  readonly service: typeof Colors;
  readonly requestStream: false;
  readonly responseStream: true;
  readonly requestType: typeof service_pb.StartIndex;
  readonly responseType: typeof service_pb.ColorResponse;
};

type ColorsAddColorFromHexOrRgb = {
  readonly methodName: string;
  readonly service: typeof Colors;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof service_pb.Color;
  readonly responseType: typeof service_pb.Reply;
};

export class Colors {
  static readonly serviceName: string;
  static readonly PickAllColors: ColorsPickAllColors;
  static readonly AddColorFromHexOrRgb: ColorsAddColorFromHexOrRgb;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }

interface UnaryResponse {
  cancel(): void;
}
interface ResponseStream<T> {
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): ResponseStream<T>;
  on(type: 'end', handler: (status?: Status) => void): ResponseStream<T>;
  on(type: 'status', handler: (status: Status) => void): ResponseStream<T>;
}
interface RequestStream<T> {
  write(message: T): RequestStream<T>;
  end(): void;
  cancel(): void;
  on(type: 'end', handler: (status?: Status) => void): RequestStream<T>;
  on(type: 'status', handler: (status: Status) => void): RequestStream<T>;
}
interface BidirectionalStream<ReqT, ResT> {
  write(message: ReqT): BidirectionalStream<ReqT, ResT>;
  end(): void;
  cancel(): void;
  on(type: 'data', handler: (message: ResT) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'end', handler: (status?: Status) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'status', handler: (status: Status) => void): BidirectionalStream<ReqT, ResT>;
}

export class ColorsClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  pickAllColors(requestMessage: service_pb.StartIndex, metadata?: grpc.Metadata): ResponseStream<service_pb.ColorResponse>;
  addColorFromHexOrRgb(
    requestMessage: service_pb.Color,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: service_pb.Reply|null) => void
  ): UnaryResponse;
  addColorFromHexOrRgb(
    requestMessage: service_pb.Color,
    callback: (error: ServiceError|null, responseMessage: service_pb.Reply|null) => void
  ): UnaryResponse;
}

