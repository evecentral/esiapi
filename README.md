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
