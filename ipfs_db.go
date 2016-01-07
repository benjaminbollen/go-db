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
  // Note: this is not a correct dependency flow, this is a sandbox
  // and should be implemented in eris:db/tendermint
  // "github.com/eris-ltd/common/go/ipfs"

  // IPFS depends on jbenet multihash; include directly for validation
  // "github.com/jbenet/go-multihash"

  // . "github.com/tendermint/go-common"
)

type IpfsDB struct {
  // wrap around LevelDB and add functionality for backing up values to IPFS
  db_wrapper *LevelDB
}

func NewIpfsDB(name string) (*IpfsDB, error) {
  leveldb_database, err := NewLevelDB(name)
  if err != nil {
    return nil, err
  }
  database := &IpfsDB{db_wrapper: leveldb_database}
  return database, nil
}

func (db *IpfsDB) Get(key []byte) []byte {
  return db.db_wrapper.Get(key)
}

func (db *IpfsDB) Set(key []byte, value []byte) {
  db.db_wrapper.Set(key, value)
}

func (db *IpfsDB) SetSync(key []byte, value []byte) {
  db.db_wrapper.SetSync(key, value)
}

func (db *IpfsDB) Delete(key []byte) {
  db.db_wrapper.Delete(key)
}

func (db *IpfsDB) DeleteSync(key []byte) {
  db.db_wrapper.DeleteSync(key)
}

func (db *IpfsDB) Close() {
  db.db_wrapper.Close()
}

func (db *IpfsDB) Print() {
  db.db_wrapper.Print()
}
