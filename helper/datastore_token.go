package helper

import (
	"cloud.google.com/go/datastore"
	"context"
	"fmt"
	"github.com/theatrus/ooauth2"
)

type cloudDatastoreStore struct {
	client *datastore.Client
	lastToken *ooauth2.Token
	prefix string
}

func DatastoreToken(project, prefix string) ooauth2.Option {
	return func(o *ooauth2.Options) error {
		ctx := context.Background()

		client, err := datastore.NewClient(ctx, project)
		if err != nil {
			return err
		}

		t, err := client.NewTransaction(ctx)
		if err != nil {
			return fmt.Errorf("datastoredb: could not connect: %v", err)

		}

		if err := t.Rollback(); err != nil {
			return fmt.Errorf("datastoredb: could not connect: %v", err)

		}

		ts := &cloudDatastoreStore{
			prefix: prefix,
			client: client,
		}
		o.TokenStore = ts
		return nil
	}
}

func (o *cloudDatastoreStore) ReadToken() (*ooauth2.Token, error) {
	ctx := context.Background() // TODO context
	key := datastore.NameKey("oauth2_tokens", o.prefix, nil)
	// Load the token using the given name and project
	var t ooauth2.Token
	if err := o.client.Get(ctx, key, &t); err != nil {
		return nil, err
	}
	return &t, nil
}

func (o *cloudDatastoreStore) WriteToken(token *ooauth2.Token) {
	ctx := context.Background() // TODO context
	key := datastore.NameKey("oauth2_tokens", o.prefix, nil)
	// Load the token using the given name and project

	o.client.RunInTransaction(ctx, func (tx *datastore.Transaction) error {
		// TODO: Validate
		tx.Put(key, token)
		return nil
	})
	if err := o.client.Get(ctx, key, &t); err != nil {
		return nil, err
	}

}
