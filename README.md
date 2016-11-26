# esiapi

`esiapi` is a Go based library and toolset for dealing with the
[EVE-Online ESI API](https://esi.tech.ccp.is/latest/).

A variety of tools are available which show the basics of the API in
addition to OAuth token management. Currently, backends for tokens
exist in Google Cloud Datastore and in a flat-file (JSON) on disk. By
using keys in Datastore, multiple worker machines can be used with a
shared per-user authentication system.

## cli/esi_token_dstool

Provision and refresh the tokens.
