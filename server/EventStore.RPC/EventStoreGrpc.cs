// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: event_store.proto
#region Designer generated code

using System;
using System.Threading;
using System.Threading.Tasks;
using grpc = global::Grpc.Core;

namespace EventStore.RPC {
  public static partial class EventStore
  {
    static readonly string __ServiceName = "eventstore.EventStore";

    static readonly grpc::Marshaller<global::EventStore.RPC.AppendToStreamRequest> __Marshaller_AppendToStreamRequest = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::EventStore.RPC.AppendToStreamRequest.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::EventStore.RPC.AppendToStreamResponse> __Marshaller_AppendToStreamResponse = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::EventStore.RPC.AppendToStreamResponse.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::EventStore.RPC.SubscribeToStreamFromRequest> __Marshaller_SubscribeToStreamFromRequest = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::EventStore.RPC.SubscribeToStreamFromRequest.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::EventStore.RPC.SubscribeToStreamFromResponse> __Marshaller_SubscribeToStreamFromResponse = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::EventStore.RPC.SubscribeToStreamFromResponse.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::EventStore.RPC.CreatePersistentSubscriptionRequest> __Marshaller_CreatePersistentSubscriptionRequest = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::EventStore.RPC.CreatePersistentSubscriptionRequest.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::EventStore.RPC.CreatePersistentSubscriptionResponse> __Marshaller_CreatePersistentSubscriptionResponse = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::EventStore.RPC.CreatePersistentSubscriptionResponse.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::EventStore.RPC.UpdatePersistentSubscriptionRequest> __Marshaller_UpdatePersistentSubscriptionRequest = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::EventStore.RPC.UpdatePersistentSubscriptionRequest.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::EventStore.RPC.UpdatePersistentSubscriptionResponse> __Marshaller_UpdatePersistentSubscriptionResponse = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::EventStore.RPC.UpdatePersistentSubscriptionResponse.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::EventStore.RPC.DeletePersistentSubscriptionRequest> __Marshaller_DeletePersistentSubscriptionRequest = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::EventStore.RPC.DeletePersistentSubscriptionRequest.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::EventStore.RPC.DeletePersistentSubscriptionResponse> __Marshaller_DeletePersistentSubscriptionResponse = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::EventStore.RPC.DeletePersistentSubscriptionResponse.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::EventStore.RPC.ConnectToPersistentSubscriptionRequest> __Marshaller_ConnectToPersistentSubscriptionRequest = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::EventStore.RPC.ConnectToPersistentSubscriptionRequest.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::EventStore.RPC.ConnectToPersistentSubscriptionResponse> __Marshaller_ConnectToPersistentSubscriptionResponse = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::EventStore.RPC.ConnectToPersistentSubscriptionResponse.Parser.ParseFrom);

    static readonly grpc::Method<global::EventStore.RPC.AppendToStreamRequest, global::EventStore.RPC.AppendToStreamResponse> __Method_AppendToStream = new grpc::Method<global::EventStore.RPC.AppendToStreamRequest, global::EventStore.RPC.AppendToStreamResponse>(
        grpc::MethodType.Unary,
        __ServiceName,
        "AppendToStream",
        __Marshaller_AppendToStreamRequest,
        __Marshaller_AppendToStreamResponse);

    static readonly grpc::Method<global::EventStore.RPC.SubscribeToStreamFromRequest, global::EventStore.RPC.SubscribeToStreamFromResponse> __Method_SubscribeToStreamFrom = new grpc::Method<global::EventStore.RPC.SubscribeToStreamFromRequest, global::EventStore.RPC.SubscribeToStreamFromResponse>(
        grpc::MethodType.DuplexStreaming,
        __ServiceName,
        "SubscribeToStreamFrom",
        __Marshaller_SubscribeToStreamFromRequest,
        __Marshaller_SubscribeToStreamFromResponse);

    static readonly grpc::Method<global::EventStore.RPC.CreatePersistentSubscriptionRequest, global::EventStore.RPC.CreatePersistentSubscriptionResponse> __Method_CreatePersistentSubscription = new grpc::Method<global::EventStore.RPC.CreatePersistentSubscriptionRequest, global::EventStore.RPC.CreatePersistentSubscriptionResponse>(
        grpc::MethodType.Unary,
        __ServiceName,
        "CreatePersistentSubscription",
        __Marshaller_CreatePersistentSubscriptionRequest,
        __Marshaller_CreatePersistentSubscriptionResponse);

    static readonly grpc::Method<global::EventStore.RPC.UpdatePersistentSubscriptionRequest, global::EventStore.RPC.UpdatePersistentSubscriptionResponse> __Method_UpdatePersistentSubscription = new grpc::Method<global::EventStore.RPC.UpdatePersistentSubscriptionRequest, global::EventStore.RPC.UpdatePersistentSubscriptionResponse>(
        grpc::MethodType.Unary,
        __ServiceName,
        "UpdatePersistentSubscription",
        __Marshaller_UpdatePersistentSubscriptionRequest,
        __Marshaller_UpdatePersistentSubscriptionResponse);

    static readonly grpc::Method<global::EventStore.RPC.DeletePersistentSubscriptionRequest, global::EventStore.RPC.DeletePersistentSubscriptionResponse> __Method_DeletePersistentSubscription = new grpc::Method<global::EventStore.RPC.DeletePersistentSubscriptionRequest, global::EventStore.RPC.DeletePersistentSubscriptionResponse>(
        grpc::MethodType.Unary,
        __ServiceName,
        "DeletePersistentSubscription",
        __Marshaller_DeletePersistentSubscriptionRequest,
        __Marshaller_DeletePersistentSubscriptionResponse);

    static readonly grpc::Method<global::EventStore.RPC.ConnectToPersistentSubscriptionRequest, global::EventStore.RPC.ConnectToPersistentSubscriptionResponse> __Method_ConnectToPersistentSubscription = new grpc::Method<global::EventStore.RPC.ConnectToPersistentSubscriptionRequest, global::EventStore.RPC.ConnectToPersistentSubscriptionResponse>(
        grpc::MethodType.DuplexStreaming,
        __ServiceName,
        "ConnectToPersistentSubscription",
        __Marshaller_ConnectToPersistentSubscriptionRequest,
        __Marshaller_ConnectToPersistentSubscriptionResponse);

    /// <summary>Service descriptor</summary>
    public static global::Google.Protobuf.Reflection.ServiceDescriptor Descriptor
    {
      get { return global::EventStore.RPC.EventStoreReflection.Descriptor.Services[0]; }
    }

    /// <summary>Base class for server-side implementations of EventStore</summary>
    public abstract partial class EventStoreBase
    {
      public virtual global::System.Threading.Tasks.Task<global::EventStore.RPC.AppendToStreamResponse> AppendToStream(global::EventStore.RPC.AppendToStreamRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      public virtual global::System.Threading.Tasks.Task SubscribeToStreamFrom(grpc::IAsyncStreamReader<global::EventStore.RPC.SubscribeToStreamFromRequest> requestStream, grpc::IServerStreamWriter<global::EventStore.RPC.SubscribeToStreamFromResponse> responseStream, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      public virtual global::System.Threading.Tasks.Task<global::EventStore.RPC.CreatePersistentSubscriptionResponse> CreatePersistentSubscription(global::EventStore.RPC.CreatePersistentSubscriptionRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      public virtual global::System.Threading.Tasks.Task<global::EventStore.RPC.UpdatePersistentSubscriptionResponse> UpdatePersistentSubscription(global::EventStore.RPC.UpdatePersistentSubscriptionRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      public virtual global::System.Threading.Tasks.Task<global::EventStore.RPC.DeletePersistentSubscriptionResponse> DeletePersistentSubscription(global::EventStore.RPC.DeletePersistentSubscriptionRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      public virtual global::System.Threading.Tasks.Task ConnectToPersistentSubscription(grpc::IAsyncStreamReader<global::EventStore.RPC.ConnectToPersistentSubscriptionRequest> requestStream, grpc::IServerStreamWriter<global::EventStore.RPC.ConnectToPersistentSubscriptionResponse> responseStream, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

    }

    /// <summary>Client for EventStore</summary>
    public partial class EventStoreClient : grpc::ClientBase<EventStoreClient>
    {
      /// <summary>Creates a new client for EventStore</summary>
      /// <param name="channel">The channel to use to make remote calls.</param>
      public EventStoreClient(grpc::Channel channel) : base(channel)
      {
      }
      /// <summary>Creates a new client for EventStore that uses a custom <c>CallInvoker</c>.</summary>
      /// <param name="callInvoker">The callInvoker to use to make remote calls.</param>
      public EventStoreClient(grpc::CallInvoker callInvoker) : base(callInvoker)
      {
      }
      /// <summary>Protected parameterless constructor to allow creation of test doubles.</summary>
      protected EventStoreClient() : base()
      {
      }
      /// <summary>Protected constructor to allow creation of configured clients.</summary>
      /// <param name="configuration">The client configuration.</param>
      protected EventStoreClient(ClientBaseConfiguration configuration) : base(configuration)
      {
      }

      public virtual global::EventStore.RPC.AppendToStreamResponse AppendToStream(global::EventStore.RPC.AppendToStreamRequest request, grpc::Metadata headers = null, DateTime? deadline = null, CancellationToken cancellationToken = default(CancellationToken))
      {
        return AppendToStream(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual global::EventStore.RPC.AppendToStreamResponse AppendToStream(global::EventStore.RPC.AppendToStreamRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_AppendToStream, null, options, request);
      }
      public virtual grpc::AsyncUnaryCall<global::EventStore.RPC.AppendToStreamResponse> AppendToStreamAsync(global::EventStore.RPC.AppendToStreamRequest request, grpc::Metadata headers = null, DateTime? deadline = null, CancellationToken cancellationToken = default(CancellationToken))
      {
        return AppendToStreamAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual grpc::AsyncUnaryCall<global::EventStore.RPC.AppendToStreamResponse> AppendToStreamAsync(global::EventStore.RPC.AppendToStreamRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_AppendToStream, null, options, request);
      }
      public virtual grpc::AsyncDuplexStreamingCall<global::EventStore.RPC.SubscribeToStreamFromRequest, global::EventStore.RPC.SubscribeToStreamFromResponse> SubscribeToStreamFrom(grpc::Metadata headers = null, DateTime? deadline = null, CancellationToken cancellationToken = default(CancellationToken))
      {
        return SubscribeToStreamFrom(new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual grpc::AsyncDuplexStreamingCall<global::EventStore.RPC.SubscribeToStreamFromRequest, global::EventStore.RPC.SubscribeToStreamFromResponse> SubscribeToStreamFrom(grpc::CallOptions options)
      {
        return CallInvoker.AsyncDuplexStreamingCall(__Method_SubscribeToStreamFrom, null, options);
      }
      public virtual global::EventStore.RPC.CreatePersistentSubscriptionResponse CreatePersistentSubscription(global::EventStore.RPC.CreatePersistentSubscriptionRequest request, grpc::Metadata headers = null, DateTime? deadline = null, CancellationToken cancellationToken = default(CancellationToken))
      {
        return CreatePersistentSubscription(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual global::EventStore.RPC.CreatePersistentSubscriptionResponse CreatePersistentSubscription(global::EventStore.RPC.CreatePersistentSubscriptionRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_CreatePersistentSubscription, null, options, request);
      }
      public virtual grpc::AsyncUnaryCall<global::EventStore.RPC.CreatePersistentSubscriptionResponse> CreatePersistentSubscriptionAsync(global::EventStore.RPC.CreatePersistentSubscriptionRequest request, grpc::Metadata headers = null, DateTime? deadline = null, CancellationToken cancellationToken = default(CancellationToken))
      {
        return CreatePersistentSubscriptionAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual grpc::AsyncUnaryCall<global::EventStore.RPC.CreatePersistentSubscriptionResponse> CreatePersistentSubscriptionAsync(global::EventStore.RPC.CreatePersistentSubscriptionRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_CreatePersistentSubscription, null, options, request);
      }
      public virtual global::EventStore.RPC.UpdatePersistentSubscriptionResponse UpdatePersistentSubscription(global::EventStore.RPC.UpdatePersistentSubscriptionRequest request, grpc::Metadata headers = null, DateTime? deadline = null, CancellationToken cancellationToken = default(CancellationToken))
      {
        return UpdatePersistentSubscription(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual global::EventStore.RPC.UpdatePersistentSubscriptionResponse UpdatePersistentSubscription(global::EventStore.RPC.UpdatePersistentSubscriptionRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_UpdatePersistentSubscription, null, options, request);
      }
      public virtual grpc::AsyncUnaryCall<global::EventStore.RPC.UpdatePersistentSubscriptionResponse> UpdatePersistentSubscriptionAsync(global::EventStore.RPC.UpdatePersistentSubscriptionRequest request, grpc::Metadata headers = null, DateTime? deadline = null, CancellationToken cancellationToken = default(CancellationToken))
      {
        return UpdatePersistentSubscriptionAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual grpc::AsyncUnaryCall<global::EventStore.RPC.UpdatePersistentSubscriptionResponse> UpdatePersistentSubscriptionAsync(global::EventStore.RPC.UpdatePersistentSubscriptionRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_UpdatePersistentSubscription, null, options, request);
      }
      public virtual global::EventStore.RPC.DeletePersistentSubscriptionResponse DeletePersistentSubscription(global::EventStore.RPC.DeletePersistentSubscriptionRequest request, grpc::Metadata headers = null, DateTime? deadline = null, CancellationToken cancellationToken = default(CancellationToken))
      {
        return DeletePersistentSubscription(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual global::EventStore.RPC.DeletePersistentSubscriptionResponse DeletePersistentSubscription(global::EventStore.RPC.DeletePersistentSubscriptionRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_DeletePersistentSubscription, null, options, request);
      }
      public virtual grpc::AsyncUnaryCall<global::EventStore.RPC.DeletePersistentSubscriptionResponse> DeletePersistentSubscriptionAsync(global::EventStore.RPC.DeletePersistentSubscriptionRequest request, grpc::Metadata headers = null, DateTime? deadline = null, CancellationToken cancellationToken = default(CancellationToken))
      {
        return DeletePersistentSubscriptionAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual grpc::AsyncUnaryCall<global::EventStore.RPC.DeletePersistentSubscriptionResponse> DeletePersistentSubscriptionAsync(global::EventStore.RPC.DeletePersistentSubscriptionRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_DeletePersistentSubscription, null, options, request);
      }
      public virtual grpc::AsyncDuplexStreamingCall<global::EventStore.RPC.ConnectToPersistentSubscriptionRequest, global::EventStore.RPC.ConnectToPersistentSubscriptionResponse> ConnectToPersistentSubscription(grpc::Metadata headers = null, DateTime? deadline = null, CancellationToken cancellationToken = default(CancellationToken))
      {
        return ConnectToPersistentSubscription(new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual grpc::AsyncDuplexStreamingCall<global::EventStore.RPC.ConnectToPersistentSubscriptionRequest, global::EventStore.RPC.ConnectToPersistentSubscriptionResponse> ConnectToPersistentSubscription(grpc::CallOptions options)
      {
        return CallInvoker.AsyncDuplexStreamingCall(__Method_ConnectToPersistentSubscription, null, options);
      }
      /// <summary>Creates a new instance of client from given <c>ClientBaseConfiguration</c>.</summary>
      protected override EventStoreClient NewInstance(ClientBaseConfiguration configuration)
      {
        return new EventStoreClient(configuration);
      }
    }

    /// <summary>Creates service definition that can be registered with a server</summary>
    /// <param name="serviceImpl">An object implementing the server-side handling logic.</param>
    public static grpc::ServerServiceDefinition BindService(EventStoreBase serviceImpl)
    {
      return grpc::ServerServiceDefinition.CreateBuilder()
          .AddMethod(__Method_AppendToStream, serviceImpl.AppendToStream)
          .AddMethod(__Method_SubscribeToStreamFrom, serviceImpl.SubscribeToStreamFrom)
          .AddMethod(__Method_CreatePersistentSubscription, serviceImpl.CreatePersistentSubscription)
          .AddMethod(__Method_UpdatePersistentSubscription, serviceImpl.UpdatePersistentSubscription)
          .AddMethod(__Method_DeletePersistentSubscription, serviceImpl.DeletePersistentSubscription)
          .AddMethod(__Method_ConnectToPersistentSubscription, serviceImpl.ConnectToPersistentSubscription).Build();
    }

  }
}
#endregion
