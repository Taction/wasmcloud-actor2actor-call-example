use wasmbus_rpc::actor::prelude::*;
use wasmcloud_interface_httpserver::{HttpRequest, HttpResponse, HttpServer, HttpServerReceiver};
use uppercase::*;

#[derive(Debug, Default, Actor, HealthResponder)]
#[services(Actor, Uppercase)]
struct ActorServerActor {}

/// Implementation of HttpServer trait methods
#[async_trait]
impl Uppercase for ActorServerActor {
    /// Returns a uppercase string in the response body.
    async fn upper(&self, ctx: &Context, arg: &UppercaseRequest) -> RpcResult<UppercaseResponse> {
        return RpcResult::Ok(UppercaseResponse{ data: Some(arg.data.to_uppercase()), success: true })

    }
}

