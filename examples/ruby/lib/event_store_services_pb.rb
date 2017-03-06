# Generated by the protocol buffer compiler.  DO NOT EDIT!
# Source: event_store.proto for package 'eventstore'

require 'grpc'
require 'event_store_pb'

module Eventstore
  module EventStore
    class Service

      include GRPC::GenericService

      self.marshal_class_method = :encode
      self.unmarshal_class_method = :decode
      self.service_name = 'eventstore.EventStore'

      rpc :AppendToStream, AppendToStreamRequest, AppendToStreamResponse
      rpc :SubscribeToStreamFrom, stream(SubscribeToStreamFromRequest), stream(SubscribeToStreamFromResponse)
    end

    Stub = Service.rpc_stub_class
  end
end
