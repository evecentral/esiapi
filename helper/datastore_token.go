package helper

import (
	"cloud.google.com/go/datastore"
	"context"
	"fmt"
	"github.com/theatrus/ooauth2"
	"time"
)

type cloudDatastoreStore struct {
	client       *datastore.Client
	lastToken    *ooauth2.Token
	expireWindow time.Duration
	prefix       string
}

func NewDatastoreTokenStore(project, prefix string, expireWindow time.Duration) (ooauth2.TokenStore, error) {
	ctx := context.Background()

	client, err := datastore.NewClient(ctx, project)
	if err != nil {
		return nil, err
	}

	t, err := client.NewTransaction(ctx)
	if err != nil {
		return nil, fmt.Errorf("datastoredb: could not connect: %v", err)

	}

	if err := t.Rollback(); err != nil {
		return nil, fmt.Errorf("datastoredb: could not connect: %v", err)

	}

	ts := &cloudDatastoreStore{
		prefix:       prefix,
		client:       client,
		expireWindow: expireWindow,
	}
	return ts, nil
}

func (o *cloudDatastoreStore) ReadToken() (*ooauth2.Token, error) {
	if o.lastToken != nil && !o.lastToken.ExpiringWithin(o.expireWindow) {
		return o.lastToken, nil
	}

	ctx := context.Background() // TODO context
	key := datastore.NameKey("oauth2_tokens", o.prefix, nil)
	// Load the token using the given name and project
	var t ooauth2.Token

	if err := o.client.Get(ctx, key, &t); err != nil {
		return nil, err
	}
	o.lastToken = &t

	return &t, nil
}

func (o *cloudDatastoreStore) WriteToken(token *ooauth2.Token) {
	ctx := context.Background() // TODO context
	key := datastore.NameKey("oauth2_tokens", o.prefix, nil)
	// Load the token using the given name and project

	o.client.RunInTransaction(ctx, func(tx *datastore.Transaction) error {
		// TODO: Validate
		tx.Put(key, token)
		return nil
	})

	o.lastToken = token
	return
}
