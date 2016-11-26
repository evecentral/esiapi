# esiapi

`esiapi` is a Go based library and toolset for dealing with the
[EVE-Online ESI API](https://esi.tech.ccp.is/latest/).

A variety of tools are available which show the basics of the API in
addition to OAuth token management. Currently, backends for tokens
exist in Google Cloud Datastore and in a flat-file (JSON) on disk. By
using keys in Datastore, multiple worker machines can be used with a
shared per-user authentication system.

Dependencies in this package can be managed by
using [Glide](http://glide.sh), a Go package manager. Running `glide
up && glide install` in this directory will vendor all dependencies
from upstreams.

# Clients and OAuth2

The clients in this package are built
with [Go-Swagger](https://github.com/go-swagger/go-swagger) using a
reduced version of the Swagger client definition from CCP. There is a
bug in Go-Swagger which refuses to process the current CCP swagger
definition. See this issue for details
https://github.com/go-swagger/go-swagger/issues/733

Helpers are provided to navigate the OAuth2 flows, and several CLI
tools exist to navigate the credentials.

Most of the tools are built with the assumption that you will be using
Google Cloud Datastore as a backing store, but there is no limitation
in the API for how credentials and tokens are stored.

## cli/esi_token_dstool

Provision and refresh the tokens in a Google Cloud Datastore entity. 

Parameters:

 * `-esi.settings=path/to/file` - A path to a JSON file containing the credentials
   to use to bootstrap the OAuth2 flow. See the
   `settings.json.example` file in this repository for reference.
   
 * `-esi.ds.project=project-name` - The project name to use for the
   Cloud Datastore table. Authentication is handled out of band - use
   the `gcloud auth` workflow to authenticate on a client machine, or
   use the service account already provisioned. 
   
 * `-esi.token.name=token-name` - What to name this token. Allows
   using multiple per-user tokens.
