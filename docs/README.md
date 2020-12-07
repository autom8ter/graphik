# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [graphik.proto](#graphik.proto)
    - [AggFilter](#api.AggFilter)
    - [Authorizer](#api.Authorizer)
    - [Authorizers](#api.Authorizers)
    - [CFilter](#api.CFilter)
    - [ChanFilter](#api.ChanFilter)
    - [Change](#api.Change)
    - [Connection](#api.Connection)
    - [ConnectionConstructor](#api.ConnectionConstructor)
    - [ConnectionConstructors](#api.ConnectionConstructors)
    - [Connections](#api.Connections)
    - [Doc](#api.Doc)
    - [DocConstructor](#api.DocConstructor)
    - [DocConstructors](#api.DocConstructors)
    - [Docs](#api.Docs)
    - [EFilter](#api.EFilter)
    - [Edit](#api.Edit)
    - [ExprFilter](#api.ExprFilter)
    - [Filter](#api.Filter)
    - [Flags](#api.Flags)
    - [Graph](#api.Graph)
    - [Index](#api.Index)
    - [IndexConstructor](#api.IndexConstructor)
    - [Indexes](#api.Indexes)
    - [Message](#api.Message)
    - [Metadata](#api.Metadata)
    - [OutboundMessage](#api.OutboundMessage)
    - [Path](#api.Path)
    - [PathConstructor](#api.PathConstructor)
    - [Paths](#api.Paths)
    - [Pong](#api.Pong)
    - [Request](#api.Request)
    - [SConnectFilter](#api.SConnectFilter)
    - [Schema](#api.Schema)
    - [TFilter](#api.TFilter)
    - [Traversal](#api.Traversal)
    - [Traversals](#api.Traversals)
    - [TypeValidator](#api.TypeValidator)
    - [TypeValidators](#api.TypeValidators)
  
    - [Direction](#api.Direction)
  
    - [DatabaseService](#api.DatabaseService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="graphik.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## graphik.proto



<a name="api.AggFilter"></a>

### AggFilter



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| filter | [Filter](#api.Filter) |  |  |
| aggregate | [string](#string) |  |  |
| field | [string](#string) |  |  |






<a name="api.Authorizer"></a>

### Authorizer



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| expression | [string](#string) |  |  |






<a name="api.Authorizers"></a>

### Authorizers



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| authorizers | [Authorizer](#api.Authorizer) | repeated |  |






<a name="api.CFilter"></a>

### CFilter
CFilter is used to fetch connections related to a single noted


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| doc_path | [Path](#api.Path) |  | doc_path is the path to the target doc. (validator.field) = {msg_exists : true}] |
| gtype | [string](#string) |  | gtype is the type of connections to return. (validator.field) = {regex : &#34;^.{1,225}$&#34;} |
| expression | [string](#string) |  | expression is a CEL expression used to filter connections/modes |
| limit | [int32](#int32) |  | limit is the maximum number of items to return. (validator.field) = {int_gt : 0} |
| sort | [string](#string) |  | custom sorting of the results. (validator.field) = {regex : &#34;((^|, )(|path.gid|path.gtype|metadata.created_at|metadata.created_by|metadata.updated_at|metadata.updated_by|metadata.version|^attributes.(.*)))&#43;$&#34;} |
| seek | [string](#string) |  | seek to a specific key for pagination |
| reverse | [bool](#bool) |  | reverse the results |






<a name="api.ChanFilter"></a>

### ChanFilter
ChanFilter is used to filter messages in a pubsub channel


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| channel | [string](#string) |  | channel is the target channel to filter from |
| expression | [string](#string) |  | expression is CEL expression used to filter messages |






<a name="api.Change"></a>

### Change
Change represents a set of state changes in the graph


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| method | [string](#string) |  | method is the gRPC method invoked |
| identity | [Doc](#api.Doc) |  | identity is the identity invoking the change |
| timestamp | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | timestamp is when the change was made |
| paths_affected | [Paths](#api.Paths) |  | paths_affected are paths to docs/connections that have been affected by the change |






<a name="api.Connection"></a>

### Connection
Connection is a graph primitive that represents a relationship between two docs


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| path | [Path](#api.Path) |  | path is the path to the connection |
| attributes | [google.protobuf.Struct](#google.protobuf.Struct) |  | attributes are k/v pairs |
| directed | [bool](#bool) |  | directed is false if the connection is bi-directional |
| from | [Path](#api.Path) |  | from is the doc path that is the source of the connection |
| to | [Path](#api.Path) |  | to is the doc path that is the destination of the connection |
| metadata | [Metadata](#api.Metadata) |  | metadata is general metadata collected about the connection |






<a name="api.ConnectionConstructor"></a>

### ConnectionConstructor
ConnectionConstructor is used to create an Connection


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| path | [PathConstructor](#api.PathConstructor) |  | path is the path to the new Connection. If an id isn&#39;t present, one will be generated. |
| attributes | [google.protobuf.Struct](#google.protobuf.Struct) |  | attributes are k/v pairs |
| directed | [bool](#bool) |  | directed is false if the connection is bi-directional |
| from | [Path](#api.Path) |  | from is the doc path that is the root of the connection |
| to | [Path](#api.Path) |  | to is the doc path that is the destination of the connection |






<a name="api.ConnectionConstructors"></a>

### ConnectionConstructors
ConnectionConstructors is an array of ConnectionConstructor


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| connections | [ConnectionConstructor](#api.ConnectionConstructor) | repeated |  |






<a name="api.Connections"></a>

### Connections
Connections is an array of Connection


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| connections | [Connection](#api.Connection) | repeated |  |
| seek_next | [string](#string) |  |  |






<a name="api.Doc"></a>

### Doc
Doc is a Graph primitive representing a single entity/resource. It is connected to other docs via Connections


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| path | [Path](#api.Path) |  | path is the path to the doc |
| attributes | [google.protobuf.Struct](#google.protobuf.Struct) |  | k/v pairs |
| metadata | [Metadata](#api.Metadata) |  | metadata is general metadata collected about the doc |






<a name="api.DocConstructor"></a>

### DocConstructor
DocConstructor is used to create a doc


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| path | [PathConstructor](#api.PathConstructor) |  | path is the path to the new Doc. If an id isn&#39;t present, one will be generated. |
| attributes | [google.protobuf.Struct](#google.protobuf.Struct) |  | arbitrary k/v pairs |






<a name="api.DocConstructors"></a>

### DocConstructors
DocConstructor is used to create a batch of docs


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| docs | [DocConstructor](#api.DocConstructor) | repeated | docs is an array of doc constructors |






<a name="api.Docs"></a>

### Docs
Docs is an array of docs


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| docs | [Doc](#api.Doc) | repeated | docs is an array of docs |
| seek_next | [string](#string) |  |  |






<a name="api.EFilter"></a>

### EFilter
EFilter is used to patch/edit docs/connections


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| filter | [Filter](#api.Filter) |  | filter is used to filter docs/connections to patch |
| attributes | [google.protobuf.Struct](#google.protobuf.Struct) |  | attributes are k/v pairs used to overwrite k/v pairs on all docs/connections that pass the filter |






<a name="api.Edit"></a>

### Edit
Edit patches the attributes of a Doc or Connection


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| path | [Path](#api.Path) |  | path is the path to the target doc/connection to patch |
| attributes | [google.protobuf.Struct](#google.protobuf.Struct) |  | attributes are k/v pairs used to overwrite k/v pairs on a doc/connection |






<a name="api.ExprFilter"></a>

### ExprFilter



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| expression | [string](#string) |  | expression is a CEL expression used to filter connections/nodes |






<a name="api.Filter"></a>

### Filter
Filter is a generic filter using Common Expression Language


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| gtype | [string](#string) |  | gtype is the doc/connection type to be filtered |
| expression | [string](#string) |  | expression is a CEL expression used to filter connections/modes |
| limit | [int32](#int32) |  | limit is the maximum number of items to return |
| sort | [string](#string) |  | custom sorting of the results. |
| seek | [string](#string) |  | seek to a specific key for pagination |
| reverse | [bool](#bool) |  | reverse the results |
| index | [string](#string) |  | search in a specific index |






<a name="api.Flags"></a>

### Flags



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| open_id_discovery | [string](#string) |  | open id connect discovery uri ex: https://accounts.google.com/.well-known/openid-configuration (env: GRAPHIK_OPEN_ID) |
| storage_path | [string](#string) |  | persistant storage path (env: GRAPHIK_STORAGE_PATH) |
| metrics | [bool](#bool) |  | enable prometheus &amp; pprof metrics (emv: GRAPHIK_METRICS = true) |
| allow_headers | [string](#string) | repeated | cors allow headers (env: GRAPHIK_ALLOW_HEADERS) |
| allow_methods | [string](#string) | repeated | cors allow methods (env: GRAPHIK_ALLOW_METHODS) |
| allow_origins | [string](#string) | repeated | cors allow origins (env: GRAPHIK_ALLOW_ORIGINS) |
| root_users | [string](#string) | repeated | root user is a list of email addresses that bypass authorizers. (env: GRAPHIK_ROOT_USERS) |
| tls_cert | [string](#string) |  |  |
| tls_key | [string](#string) |  |  |
| playground_client_id | [string](#string) |  |  |
| playground_client_secret | [string](#string) |  |  |
| playground_redirect | [string](#string) |  |  |






<a name="api.Graph"></a>

### Graph
Graph is an array of docs and connections


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| docs | [Docs](#api.Docs) |  | docs are docs present in the graph |
| connections | [Connections](#api.Connections) |  | connections are connections present in the graph |






<a name="api.Index"></a>

### Index



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| gtype | [string](#string) |  | gtype is the doc/connection type to be filtered |
| expression | [string](#string) |  | expression is a CEL expression used to filter connections/modes |
| docs | [bool](#bool) |  | if docs is true, this index will be applied to documents. Either docs or connections may be true, but not both. |
| connections | [bool](#bool) |  | if docs is true, this index will be applied to connections. Either docs or connections may be true, but not both. |






<a name="api.IndexConstructor"></a>

### IndexConstructor



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| gtype | [string](#string) |  | gtype is the doc/connection type to be filtered |
| expression | [string](#string) |  | expression is a CEL expression used to filter connections/modes |
| docs | [bool](#bool) |  | if docs is true, this index will be applied to documents. Either docs or connections may be true, but not both. |
| connections | [bool](#bool) |  | if docs is true, this index will be applied to connections. Either docs or connections may be true, but not both. |






<a name="api.Indexes"></a>

### Indexes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| indexes | [Index](#api.Index) | repeated |  |






<a name="api.Message"></a>

### Message
Message is received on PubSub subscriptions


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| channel | [string](#string) |  | channel is the channel the message was sent to |
| data | [google.protobuf.Struct](#google.protobuf.Struct) |  | data is the data sent with the message |
| sender | [Path](#api.Path) |  | sender is the identity that sent the message |
| timestamp | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | timestamp is when the message was sent |






<a name="api.Metadata"></a>

### Metadata
Metadata is general metadata collected on docs/connections


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| created_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | created_at is the unix timestamp when the doc/connection was created |
| updated_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | updated_at is the unix timestamp when the doc/connection was last updated |
| version | [uint64](#uint64) |  | version iterates by 1 every time the doc/connection is modified |






<a name="api.OutboundMessage"></a>

### OutboundMessage
OutboundMessage is a message to be published to a pubsub channel


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| channel | [string](#string) |  | channel is the target channel to send the message to |
| data | [google.protobuf.Struct](#google.protobuf.Struct) |  | data is the data to send with the message |






<a name="api.Path"></a>

### Path
Path describes a doc/connection type &amp; id


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| gtype | [string](#string) |  | gtype is the type of the doc/connection ex: pet |
| gid | [string](#string) |  | gid is the unique id of the doc/connection within the context of it&#39;s type |






<a name="api.PathConstructor"></a>

### PathConstructor
PathConstructor creates a new Path


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| gtype | [string](#string) |  | gtype is the type of the doc/connection ex: pet |
| gid | [string](#string) |  | gid is the unique id of the doc/connection within the context of it&#39;s type |






<a name="api.Paths"></a>

### Paths
Paths is an array of paths


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| paths | [Path](#api.Path) | repeated |  |






<a name="api.Pong"></a>

### Pong
Pong returns PONG if the server is healthy


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| message | [string](#string) |  | message returns PONG if healthy |






<a name="api.Request"></a>

### Request



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| method | [string](#string) |  | method is the rpc method |
| identity | [Doc](#api.Doc) |  | identity is the identity making the request |
| timestamp | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | timestamp is when the intercept was received |
| request | [google.protobuf.Struct](#google.protobuf.Struct) |  | request is the intercepted request |






<a name="api.SConnectFilter"></a>

### SConnectFilter



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| filter | [Filter](#api.Filter) |  |  |
| gtype | [string](#string) |  |  |
| attributes | [google.protobuf.Struct](#google.protobuf.Struct) |  | attributes are k/v pairs |
| directed | [bool](#bool) |  | directed is false if the connection is bi-directional |
| from | [Path](#api.Path) |  | from is the doc path that is the root of the connection |






<a name="api.Schema"></a>

### Schema
Schema returns registered connection &amp; doc types


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| connection_types | [string](#string) | repeated | connection_types are the types of connections in the graph |
| doc_types | [string](#string) | repeated | doc_types are the types of docs in the graph |
| authorizers | [Authorizers](#api.Authorizers) |  |  |
| validators | [TypeValidators](#api.TypeValidators) |  |  |
| indexes | [Indexes](#api.Indexes) |  |  |






<a name="api.TFilter"></a>

### TFilter



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| root | [Path](#api.Path) |  |  |
| doc_expression | [string](#string) |  |  |
| connection_expression | [string](#string) |  |  |
| limit | [int32](#int32) |  |  |
| sort | [string](#string) |  | custom sorting of the results. |






<a name="api.Traversal"></a>

### Traversal



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| doc | [Doc](#api.Doc) |  |  |
| relative_path | [Paths](#api.Paths) |  |  |
| direction | [Direction](#api.Direction) |  |  |






<a name="api.Traversals"></a>

### Traversals



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| traversals | [Traversal](#api.Traversal) | repeated |  |






<a name="api.TypeValidator"></a>

### TypeValidator



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| gtype | [string](#string) |  |  |
| expression | [string](#string) |  |  |
| docs | [bool](#bool) |  | if docs is true, this validator will be applied to documents. Either docs or connections may be true, but not both. |
| connections | [bool](#bool) |  | if docs is true, this validator will be applied to connections. Either docs or connections may be true, but not both. |






<a name="api.TypeValidators"></a>

### TypeValidators



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| validators | [TypeValidator](#api.TypeValidator) | repeated |  |





 


<a name="api.Direction"></a>

### Direction


| Name | Number | Description |
| ---- | ------ | ----------- |
| None | 0 |  |
| From | 1 |  |
| To | 2 |  |


 

 


<a name="api.DatabaseService"></a>

### DatabaseService
DatabaseService is the primary database service

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Ping | [.google.protobuf.Empty](#google.protobuf.Empty) | [Pong](#api.Pong) | Ping returns PONG if the server is health |
| GetSchema | [.google.protobuf.Empty](#google.protobuf.Empty) | [Schema](#api.Schema) | GetSchema gets schema about the Graph doc &amp; connection types |
| SetAuthorizers | [Authorizers](#api.Authorizers) | [.google.protobuf.Empty](#google.protobuf.Empty) |  |
| SetIndexes | [Indexes](#api.Indexes) | [.google.protobuf.Empty](#google.protobuf.Empty) |  |
| SetTypeValidators | [TypeValidators](#api.TypeValidators) | [.google.protobuf.Empty](#google.protobuf.Empty) |  |
| Me | [.google.protobuf.Empty](#google.protobuf.Empty) | [Doc](#api.Doc) | Me returns a Doc of the currently logged in identity(the subject of the JWT) |
| CreateDoc | [DocConstructor](#api.DocConstructor) | [Doc](#api.Doc) | CreateDoc creates a doc in the graph |
| CreateDocs | [DocConstructors](#api.DocConstructors) | [Docs](#api.Docs) | CreateDocs creates a batch of docs in the graph |
| GetDoc | [Path](#api.Path) | [Doc](#api.Doc) | GetDoc gets a single doc in the graph |
| SearchDocs | [Filter](#api.Filter) | [Docs](#api.Docs) | SearchDocs searches the graph for docs |
| Traverse | [TFilter](#api.TFilter) | [Traversals](#api.Traversals) | Traverse executes a depth first search of the graph for docs |
| EditDoc | [Edit](#api.Edit) | [Doc](#api.Doc) | EditDoc patches a docs attributes |
| EditDocs | [EFilter](#api.EFilter) | [Docs](#api.Docs) | EditDocs patches a batch of docs attributes that pass the patch filter |
| DelDoc | [Path](#api.Path) | [.google.protobuf.Empty](#google.protobuf.Empty) | DelDoc deletes a doc &amp; all of it&#39;s connected connections |
| DelDocs | [Filter](#api.Filter) | [.google.protobuf.Empty](#google.protobuf.Empty) | DelDocs deletes a batch of docs that pass the filter |
| CreateConnection | [ConnectionConstructor](#api.ConnectionConstructor) | [Connection](#api.Connection) | CreateConnection creates an connection in the graph |
| CreateConnections | [ConnectionConstructors](#api.ConnectionConstructors) | [Connections](#api.Connections) | CreateConnections creates a batch of connections in the graph |
| SearchAndConnect | [SConnectFilter](#api.SConnectFilter) | [Connections](#api.Connections) |  |
| GetConnection | [Path](#api.Path) | [Connection](#api.Connection) | GetConnection gets a single connection in the graph |
| SearchConnections | [Filter](#api.Filter) | [Connections](#api.Connections) | SearchConnections searches the graph for connections |
| EditConnection | [Edit](#api.Edit) | [Connection](#api.Connection) | EditConnection patches an connections attributes |
| EditConnections | [EFilter](#api.EFilter) | [Connections](#api.Connections) | EditConnections patches a batch of connections attributes that pass the patch filter |
| DelConnection | [Path](#api.Path) | [.google.protobuf.Empty](#google.protobuf.Empty) | DelConnection deletes an connection from the graph |
| DelConnections | [Filter](#api.Filter) | [.google.protobuf.Empty](#google.protobuf.Empty) | DelConnections deletes a batch of connections that pass the filter |
| ConnectionsFrom | [CFilter](#api.CFilter) | [Connections](#api.Connections) | ConnectionsFrom returns connections that source from the given doc path that pass the filter |
| ConnectionsTo | [CFilter](#api.CFilter) | [Connections](#api.Connections) | ConnectionsTo returns connections that point to the given doc path that pass the filter |
| AggregateDocs | [AggFilter](#api.AggFilter) | [.google.protobuf.Value](#google.protobuf.Value) |  |
| AggregateConnections | [AggFilter](#api.AggFilter) | [.google.protobuf.Value](#google.protobuf.Value) |  |
| Publish | [OutboundMessage](#api.OutboundMessage) | [.google.protobuf.Empty](#google.protobuf.Empty) | Publish publishes a message to a pubsub channel |
| Subscribe | [ChanFilter](#api.ChanFilter) | [Message](#api.Message) stream | Subscribe subscribes to messages on a pubsub channel |
| SubscribeChanges | [ExprFilter](#api.ExprFilter) | [Change](#api.Change) stream |  |
| PushDocConstructors | [DocConstructor](#api.DocConstructor) stream | [Doc](#api.Doc) stream |  |
| PushConnectionConstructors | [ConnectionConstructor](#api.ConnectionConstructor) stream | [Connection](#api.Connection) stream |  |
| SeedDocs | [Doc](#api.Doc) stream | [.google.protobuf.Empty](#google.protobuf.Empty) |  |
| SeedConnections | [Connection](#api.Connection) stream | [.google.protobuf.Empty](#google.protobuf.Empty) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

