// Notes: this is a draft and should be moved to eris-ltd/tendermint
// 1. problem with IAVL+ tree rotation and go-db structure, stale objects?
// 2. efficiency of accessing ipfs over http PostAPICall?
// 3. if ipfs is accessed over http call, this process should keep a copy?
// 4. implemented at db level, there is no knowledge of a graph structure
//    arguably the merkle structure can be reflected in the IPFS DAG

// proposal:
// ipfs_db acts as a persistent (levelDB) db that additionally stores
// the values to IPFS, and stores the retrieved IPFS hash for the stored value,
// under the key.
// It is proposed that only the value []byte are stored to IPFS,
// and not the key []byte + value []byte.  The main argument is that the
// key []byte + value IPFS hash, can be separately stored to IPFS; however,
// at the user level the structure of the data might desire a more optimal
// graph structure where the key-value is absorbed into a named link of a
// higher object (commit or transaction block).

package db

import (
  "fmt"

  "github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/errors"
	"github.com/syndtr/goleveldb/leveldb/opt"

  // Note: this is not a correct dependency flow, this is a sandbox
  // and should be implemented in eris:db/tendermint
  "github.com/eris-ltd/common/go/ipfs"

  // IPFS depends on jbenet multihash; include directly for validation
  "github.com/jbenet/go-multihash"

  . "github.com/tendermint/go-common"
)

type IpfsDB struct {
  // first copy the mem-db structure, probably redundant by passing
  // IPFS calls into goroutines
  db map[string][]byte
}

func NewIpfsDB(name string) (*IpfsDB, error) {

  database := &IpfsDB{db: make(map[string][]byte)}
  return database
}

func (db *IpfsDB) Get(key []byte) []byte {
  return db.db[string(key)]
}

func (db *IpfsDB) Set(key []byte, value []byte) {
  db.db[string(key)] = value
}

func (db *IpfsDB) SetSync(key []byte, value []byte) {
  db.db[string(key)] = value
}

func (db *IpfsDB) Delete(key []byte) {
  delete(db.db, string(key))
}

func (db *IpfsDB) DeleteSync(key []byte) {
  delete(db.db, string(key))
}

func (db *IpfsDB) Close() {
  db = nil
}

func (db *IpfsDB) Print() {
  for key, value := range db.db {
    fmt.Printf("[%X]:\t[%X]\n", []byte(key), value)
  }
}
