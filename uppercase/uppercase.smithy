// uppercase.smithy
//

// Tell the code generator how to reference symbols defined in this namespace
metadata package = [ { namespace: "com.oneitfarm.wasm.uppercase", crate: "uppercase" } ]

namespace com.oneitfarm.wasm.uppercase

use org.wasmcloud.model#wasmbus

/// Description of Uppercase service
@wasmbus( actorReceive: true )

service Uppercase {
  version: "0.1",
  operations: [ Upper ]
}

/// Uppercase - Execute transaction
operation Upper {
    input: UppercaseRequest,
    output: UppercaseResponse,
}

/// Parameters sent for Uppercase
structure UppercaseRequest {
    /// method to be Uppercased
    @required
    data: String,
}

/// Response to Uppercase
structure UppercaseResponse {
    /// Indicates a successful authorization
    @required
    success: Boolean,

    /// response data
    data: String,
}
