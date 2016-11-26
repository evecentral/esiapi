package helper

import (
	"cloud.google.com/go/datastore"
	"context"
	"fmt"
	"github.com/theatrus/ooauth2"
)

type cloudDatastoreStore struct {
	client *datastore.Client
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
