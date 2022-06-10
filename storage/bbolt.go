package storage

import (
	"fmt"
	"github.com/nuts-foundation/go-stoabs"
	"github.com/nuts-foundation/go-stoabs/bbolt"
	"path"
)

// BBoltDatabaseType holds the type name of BBolt databases.
const BBoltDatabaseType DatabaseType = "bbolt"

type bboltDatabaseAdapter struct {
	datadir string
	config  DatabaseConfig
}

func (b bboltDatabaseAdapter) createStore(moduleName string, storeName string) (stoabs.KVStore, error) {
	var bboltConfig BBoltConfig
	err := b.config.UnmarshalConfig(&bboltConfig)
	if err != nil {
		return nil, fmt.Errorf("invalid BBolt database config: %w", err)
	}
	// TODO: Do something with the config (e.g. start backup background procedure)
	databasePath := path.Join(b.datadir, moduleName, storeName+".db")
	return bbolt.CreateBBoltStore(databasePath)
}

func (b bboltDatabaseAdapter) getClass() Class {
	return VolatileStorageClass
}

func (b bboltDatabaseAdapter) getType() DatabaseType {
	return BBoltDatabaseType
}
