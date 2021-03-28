# 注册服务

**服务的可插拔**需要一个服务注册功能， 以service_name、service_info为key value。

gRpc维护注册的服务信息。当收到客户端发来的携带service_name、service_info信息的消息时，会调**Protocol Buffers**提供的反序列方法，然后找到对应的service进行相应method的调用。最后对消息进行回复，该回复应该有**rpc_id**进行标记。

# 客户端存根

客户端调用远程服务时，应该由rpc_id维护一个rpc调用的信息。

调用时传的信息应该携带service_name、service_method等信息，参数由**Protocol Buffers**进行序列化

