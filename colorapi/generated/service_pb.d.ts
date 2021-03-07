// package: colorapi
// file: service.proto

import * as jspb from "google-protobuf";

export class ColorResponse extends jspb.Message {
  hasColorResp(): boolean;
  clearColorResp(): void;
  getColorResp(): Color | undefined;
  setColorResp(value?: Color): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ColorResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ColorResponse): ColorResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ColorResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ColorResponse;
  static deserializeBinaryFromReader(message: ColorResponse, reader: jspb.BinaryReader): ColorResponse;
}

export namespace ColorResponse {
  export type AsObject = {
    colorResp?: Color.AsObject,
  }
}

export class Reply extends jspb.Message {
  getMessage(): string;
  setMessage(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Reply.AsObject;
  static toObject(includeInstance: boolean, msg: Reply): Reply.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Reply, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Reply;
  static deserializeBinaryFromReader(message: Reply, reader: jspb.BinaryReader): Reply;
}

export namespace Reply {
  export type AsObject = {
    message: string,
  }
}

export class Color extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  getHex(): string;
  setHex(value: string): void;

  getRgb(): number;
  setRgb(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Color.AsObject;
  static toObject(includeInstance: boolean, msg: Color): Color.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Color, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Color;
  static deserializeBinaryFromReader(message: Color, reader: jspb.BinaryReader): Color;
}

export namespace Color {
  export type AsObject = {
    name: string,
    hex: string,
    rgb: number,
  }
}

export class StartIndex extends jspb.Message {
  getIndex(): number;
  setIndex(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StartIndex.AsObject;
  static toObject(includeInstance: boolean, msg: StartIndex): StartIndex.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: StartIndex, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StartIndex;
  static deserializeBinaryFromReader(message: StartIndex, reader: jspb.BinaryReader): StartIndex;
}

export namespace StartIndex {
  export type AsObject = {
    index: number,
  }
}

