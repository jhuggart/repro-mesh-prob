## Reproduce a problem encountered while trying to use relfection to serve graphql endpoints with graphql-mesh

1. Install Go
1. Install npm
1. Install yarn
1. Run `yarn` to install dependencies
1. Run `go run main.go` to start up gRPC server
1. Run `yarn graphql-mesh serve` to attempt to start graphql server

After this, you should see `Unable to start GraphQL Mesh: Channel credentials must be a ChannelCredentials object`