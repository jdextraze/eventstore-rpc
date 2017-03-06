# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: event_store.proto

require 'google/protobuf'

Google::Protobuf::DescriptorPool.generated_pool.build do
  add_message "eventstore.AppendToStreamRequest" do
    optional :stream_id, :string, 1
    optional :expected_version, :int32, 2
    repeated :events, :message, 3, "eventstore.EventData"
    optional :user_credentials, :message, 4, "eventstore.UserCredentials"
  end
  add_message "eventstore.AppendToStreamResponse" do
    optional :next_expected_version, :int32, 1
    optional :position, :message, 2, "eventstore.Position"
    optional :error, :message, 3, "eventstore.Error"
  end
  add_message "eventstore.SubscribeToStreamFromRequest" do
    optional :stream_id, :string, 1
    optional :last_checkpoint, :int32, 2
    optional :user_credentials, :message, 3, "eventstore.UserCredentials"
  end
  add_message "eventstore.SubscribeToStreamFromResponse" do
    optional :event, :message, 1, "eventstore.ResolvedEvent"
    optional :drop_reason, :enum, 2, "eventstore.SubscribeToStreamFromResponse.DropReason"
    optional :error, :message, 3, "eventstore.Error"
  end
  add_enum "eventstore.SubscribeToStreamFromResponse.DropReason" do
    value :UserInitiated, 0
    value :NotAuthenticated, 1
    value :AccessDenied, 2
    value :SubscribingError, 3
    value :ServerError, 4
    value :ConnectionClosed, 5
    value :CatchUpError, 6
    value :ProcessingQueueOverflow, 7
    value :EventHandlerException, 8
    value :MaxSubscribersReached, 9
    value :PersistentSubscriptionDeleted, 10
    value :Unknown, 100
    value :NotFound, 101
  end
  add_message "eventstore.EventData" do
    optional :event_id, :bytes, 1
    optional :event_type, :string, 2
    optional :is_json, :bool, 3
    optional :data, :bytes, 4
    optional :metadata, :bytes, 5
  end
  add_message "eventstore.UserCredentials" do
    optional :username, :string, 1
    optional :password, :string, 2
  end
  add_message "eventstore.Position" do
    optional :commit_position, :int64, 1
    optional :prepare_position, :int64, 2
  end
  add_message "eventstore.Error" do
    optional :type, :string, 1
    optional :text, :string, 2
  end
  add_message "eventstore.RecordedEvent" do
    optional :event_stream_id, :string, 1
    optional :event_id, :bytes, 2
    optional :event_number, :int32, 3
    optional :event_type, :string, 4
    optional :data, :bytes, 5
    optional :metadata, :bytes, 6
    optional :is_json, :bool, 7
    optional :created, :int64, 8
    optional :created_epoch, :int64, 9
  end
  add_message "eventstore.ResolvedEvent" do
    optional :event, :message, 1, "eventstore.RecordedEvent"
    optional :position, :message, 2, "eventstore.Position"
  end
end

module Eventstore
  AppendToStreamRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("eventstore.AppendToStreamRequest").msgclass
  AppendToStreamResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("eventstore.AppendToStreamResponse").msgclass
  SubscribeToStreamFromRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("eventstore.SubscribeToStreamFromRequest").msgclass
  SubscribeToStreamFromResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("eventstore.SubscribeToStreamFromResponse").msgclass
  SubscribeToStreamFromResponse::DropReason = Google::Protobuf::DescriptorPool.generated_pool.lookup("eventstore.SubscribeToStreamFromResponse.DropReason").enummodule
  EventData = Google::Protobuf::DescriptorPool.generated_pool.lookup("eventstore.EventData").msgclass
  UserCredentials = Google::Protobuf::DescriptorPool.generated_pool.lookup("eventstore.UserCredentials").msgclass
  Position = Google::Protobuf::DescriptorPool.generated_pool.lookup("eventstore.Position").msgclass
  Error = Google::Protobuf::DescriptorPool.generated_pool.lookup("eventstore.Error").msgclass
  RecordedEvent = Google::Protobuf::DescriptorPool.generated_pool.lookup("eventstore.RecordedEvent").msgclass
  ResolvedEvent = Google::Protobuf::DescriptorPool.generated_pool.lookup("eventstore.ResolvedEvent").msgclass
end