### Requirements
You are supposed to know how to deploy actors to wasmcloud platform.

### What is this project for
In the [wasmcloud actor to actor call](https://wasmcloud.dev/app-dev/a2a/) documentation, there is a description of how to call actor via actor, but when looking at the [wasmcloud examples](https://github.com/wasmCloud/examples) and [interfaces](https://github.com/wasmCloud/interfaces) projects there is no such examples. So this project creates an example of an actor calling another actor from scratch.

This example involves an interface project that defines the call protocol between actor and actor, which contains only one function `Upper` to convert the input to uppercase; two actors, the actor client that send the request, and the actor server accepts the request and converts the request to uppercase.

This means that when testing this example, you can ignore the implementation language and deploy an actor server in any language as well as an actor client in any language.

The actor client uses wasmcloud's httpprovider to receive external requests and send the http body to the actor server.

There is [a document](https://taction.top/wasmcloud/a2acall/) about the details about this project but in Chinese.

### How to use this project
Compile wasm
```shell
cd actor-server && make
cd ../actor-client && make
```

Start two wasm, start httpprovier, link httpprovider to wasm client. And then call wasm client from httpserver.
```shell
curl your-httpprovier-server-and-port -d 'input'
```
And the `INPUT` is expected to show in your terminal.
