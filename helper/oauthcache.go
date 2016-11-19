package helper

import (
	"encoding/json"
	"github.com/theatrus/ooauth2"
	"io/ioutil"
	"log"
)

type FileTokenStore struct {
	Filename string
}

func FileToken(f *FileTokenStore) ooauth2.Option {
	return func(o *ooauth2.Options) error {
		o.TokenStore = f
		return nil
	}
}

func (o *FileTokenStore) ReadToken() (*ooauth2.Token, error) {
	fileContents, err := ioutil.ReadFile(o.Filename)
	if err != nil {
		return nil, err
	}
	var token ooauth2.Token
	err = json.Unmarshal(fileContents, &token)
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func (o *FileTokenStore) WriteToken(token *ooauth2.Token) {
	data, err := json.Marshal(token)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(o.Filename, data, 0600)
	if err != nil {
		log.Fatal(err)
	}

	return
}
