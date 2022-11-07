use wasmbus_rpc::actor::prelude::*;
use wasmcloud_interface_httpserver::{HttpRequest, HttpResponse, HttpServer, HttpServerReceiver};
use uppercase::*;
use std::str;

#[derive(Debug, Default, Actor, HealthResponder)]
#[services(Actor, HttpServer)]
struct ActorClientActor {}

/// Implementation of HttpServer trait methods
#[async_trait]
impl HttpServer for ActorClientActor {
    /// Returns a greeting, "Hello World", in the response body.
    /// If the request contains a query parameter 'name=NAME', the
    /// response is changed to "Hello NAME"
    async fn handle_request(&self, _ctx: &Context, req: &HttpRequest) -> RpcResult<HttpResponse> {
        // let text = form_urlencoded::parse(req.query_string.as_bytes())
        //     .find(|(n, _)| n == "name")
        //     .map(|(_, v)| v.to_string())
        //     .unwrap_or_else(|| "World".to_string());
        let text: String = str::from_utf8(&req.body).unwrap_or(&"").to_string();

        // let provider = UppercaseSender::to_actor("MBVITL2QJQ42KDTVVQJ2WZM6K3BMZ2CVJXLFYYQKYCLEPS5R5WOUKXIU");
        let provider = UppercaseSender::to_actor("oneitfarm/actor_server");
        let res = provider.upper(_ctx, &UppercaseRequest{ data: text }).await?;

        Ok(HttpResponse {
            body: format!("{}", res.data.map_or("no name content".to_string() ,|v| v)).as_bytes().to_vec(),
            ..Default::default()
        })
    }
}

