// package: colorapi
// file: service.proto

var service_pb = require("./service_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var Colors = (function () {
  function Colors() {}
  Colors.serviceName = "colorapi.Colors";
  return Colors;
}());

Colors.PickAllColors = {
  methodName: "PickAllColors",
  service: Colors,
  requestStream: false,
  responseStream: true,
  requestType: service_pb.StartIndex,
  responseType: service_pb.ColorResponse
};

Colors.AddColorFromHexOrRgb = {
  methodName: "AddColorFromHexOrRgb",
  service: Colors,
  requestStream: false,
  responseStream: false,
  requestType: service_pb.Color,
  responseType: service_pb.Reply
};

exports.Colors = Colors;

function ColorsClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

ColorsClient.prototype.pickAllColors = function pickAllColors(requestMessage, metadata) {
  var listeners = {
    data: [],
    end: [],
    status: []
  };
  var client = grpc.invoke(Colors.PickAllColors, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onMessage: function (responseMessage) {
      listeners.data.forEach(function (handler) {
        handler(responseMessage);
      });
    },
    onEnd: function (status, statusMessage, trailers) {
      listeners.status.forEach(function (handler) {
        handler({ code: status, details: statusMessage, metadata: trailers });
      });
      listeners.end.forEach(function (handler) {
        handler({ code: status, details: statusMessage, metadata: trailers });
      });
      listeners = null;
    }
  });
  return {
    on: function (type, handler) {
      listeners[type].push(handler);
      return this;
    },
    cancel: function () {
      listeners = null;
      client.close();
    }
  };
};

ColorsClient.prototype.addColorFromHexOrRgb = function addColorFromHexOrRgb(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Colors.AddColorFromHexOrRgb, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

exports.ColorsClient = ColorsClient;

