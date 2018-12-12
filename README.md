# `sham` - Shamir secret sharing for everyone

Brings the Shamir secret sharing algorithm from Hashicorp Vault to a CLI utility for usage outside Vault. Shamir breaks a secret key into shards such that a certain threshold is required to re-assemble the keys. This allows splitting a key into cryptographically isolated keys where no one key gives information about the final secret.

## Installation

```
go get github.com/onetwopunch/sham
```

## Usage

```
$ sham --help
Usage of sham:
  -combine string
    	The shards to combine to get the secret separated by commas
  -k int
    	The number of total key shards (default 3)
  -split string
    	The secret to split into shards
  -t int
    	The min number of shards needed to re-assemble the secret (default 2)
```
## Examples

### Splitting a secret key
```
$ sham --split 'supersecret' -k 4 -t 3
hvXh_5lEQ7-4f79f
Ee0n0a6bGBu7UesJ
N4WJV4j11Ys9WuVE
SGPS5nQMmNeg-Nnb
```

### Combining a secret key
```
$ sham --combine hvXh_5lEQ7-4f79f,Ee0n0a6bGBu7UesJ,N4WJV4j11Ys9WuVE
supersecret
```
