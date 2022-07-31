package main

import (
  "encoding/base64"
  "fmt"
  "github.com/hashicorp/vault/shamir"
  "log"
  "os"
  "strings"
  "flag"
)

func handleError(err error) {
  if err != nil {
    log.Fatalln(err)
  }
}

var encoder *base64.Encoding = base64.URLEncoding
var secret, shardsCommaSep string
var k, t int

func init() {
  flag.StringVar(&secret, "split", "", "The secret to split into shards")
  flag.StringVar(&shardsCommaSep, "combine", "", "The shards to combine to get the secret separated by commas")
  flag.IntVar(&k, "k", 3, "The number of total key shards")
  flag.IntVar(&t, "t", 2, "The min number of shards needed to re-assemble the secret")
  flag.Parse()
}

func main() {
  if len(secret) > 0 && len(shardsCommaSep) > 0 {
    fmt.Println("ERROR: Cannot both split and combine in the same action")
    flag.Usage()
    os.Exit(1)
  } else if len(secret) == 0 && len(shardsCommaSep) == 0 {
    flag.Usage()
    os.Exit(1)
  }

  // Split Command selected
  if len(secret) > 0 {
    keys, err := shamir.Split([]byte(secret), k, t)
    if err != nil {
      log.Fatalln(err)
    }

    for _, shard := range keys {
      // RFC 4648 Encoding
      encoded := encoder.EncodeToString(shard)
      fmt.Println(encoded)
    }
  } else {
    shards := strings.Split(shardsCommaSep, ",")
    shardBytes := make([][]byte, len(shards))
    for i, shard := range shards {
      decodedShard, err := base64.URLEncoding.DecodeString(shard)
      handleError(err)
      shardBytes[i] = []byte(decodedShard)
    }

    decodedSecretBytes, err := shamir.Combine(shardBytes)
    handleError(err)
    decodedSecret := string(decodedSecretBytes)
    fmt.Println(decodedSecret)
  }
}
