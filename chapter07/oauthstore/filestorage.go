package oauthstore

import(
	"errors"
	"os"
	"sync"
	"encoding/json"
	"golang.org/x/oauth2"
)

//FileStorage satisfaz nossa interface de armazenamento
type FileStorage struct{
	Path string
	mu sync.RWMutex
}

//GetToken recupera um token de um arquivo
func (f *FileStorage) GetToken() (*oauth2.Token, error)  {
	
	f.mu.RLock()

	defer f.mu.RUnlock()

	in, err := os.Open(f.Path)
	if err != nil {
		return nil, err
	}

	defer in.Close()

	var t *oauth2.Token
	data := json.NewDecoder(in)

	return t, data.Decode(&t)
}

//SetToken cria, trunca e armazena um token
// em um arquivo
func (f *FileStorage) SetToken(t *oauth2.Token) error  {
	
	if t == nil || !t.Valid() {
		return errors.New("bad token")
	}

	f.mu.Lock()

	defer f.mu.Unlock()

	out, err := os.OpenFile(f.Path, os.O_RDWR | os.O_CREATE | os.O_TRUNC, 0755)
	if err != nil {
		return err
	}

	defer out.Close()

	data, err := json.Marshal(&t)
	if err != nil {
		return  err
	}

	_, err = out.Write(data)

	return err
}