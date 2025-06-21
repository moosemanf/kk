package vault

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"os/user"
	"path/filepath"

	"filippo.io/age"
)

var (
	vaultPath     = filepath.Join(userHome(), ".kk.age")
	keyPath       = filepath.Join(userHome(), ".age", "key.txt")
	recipientPath = filepath.Join(userHome(), ".age", "recipient.txt")
)

func userHome() string {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}
	return u.HomeDir
}

func loadIdentity() (age.Identity, error) {
	f, err := os.Open(keyPath)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = errors.Join(err, f.Close())
	}()

	identities, err := age.ParseIdentities(f)
	if err != nil {
		return nil, err
	}
	if len(identities) == 0 {
		return nil, errors.New("no identities found in key file")
	}
	return identities[0], nil
}

func loadRecipient() (age.Recipient, error) {
	f, err := os.Open(recipientPath)
	if err != nil {
		fmt.Println("open recipientpath", err)
		return nil, err
	}

	defer func() {
		err = errors.Join(err, f.Close())
	}()

	recipients, err := age.ParseRecipients(f)
	if err != nil {
		fmt.Println("parse recipients", err)
		return nil, err
	}
	if len(recipients) == 0 {
		return nil, errors.New("no recipients found in key file")
	}
	return recipients[0], nil
}

func LoadVault() (map[string]string, error) {
	identity, err := loadIdentity()
	if err != nil {
		return nil, err
	}

	f, err := os.Open(vaultPath)
	if os.IsNotExist(err) {
		return make(map[string]string), nil
	}

	defer func() {
		err = errors.Join(err, f.Close())
	}()

	r, err := age.Decrypt(f, identity)
	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	var vault map[string]string
	if err := json.Unmarshal(data, &vault); err != nil {
		return nil, err
	}

	return vault, nil
}

func SaveVault(v map[string]string) error {
	recipient, err := loadRecipient()
	if err != nil {
		fmt.Println("loadRecipient")
		return err
	}

	var buf bytes.Buffer
	w, err := age.Encrypt(&buf, recipient)
	if err != nil {
		fmt.Println("encrypt")
		return err
	}
	jsonData, _ := json.MarshalIndent(v, "", "  ")
	if _, err := w.Write(jsonData); err != nil {
		fmt.Println("write json")
		return err
	}
	if err := w.Close(); err != nil {
		fmt.Println("close writer")
		return err
	}

	return os.WriteFile(vaultPath, buf.Bytes(), 0600)
}
