# Cipher

CLI to manage cryptography primitive for Nostr.

- Generate a new public/private key pair.

```shell
cipher -keygen
```

- Decode a secret/public key with NIP-19 to hex used in NIP-01.

```shell
cipher -decode nsec1uxv2kvx98rnanmsjr25pcw0clutd768s80r7yt4hp0a8s9whzqdqgxc2u7
```

- Encode a secret/public key from NIP-01 hex to NIP-19 bech32.

```shell
cipher encode -nsec e198ab30c538e7d9ee121aa81c39f8ff16df68f03bc7e22eb70bfa7815d7101a
```
