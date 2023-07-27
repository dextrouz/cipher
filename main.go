package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ffiat/nostr"
)

func main() {

    // Keygen and decode flags

    keygen := flag.Bool("keygen", false, "display colorized output")
    decode := flag.String("decode", "", "decode nsec/npub")
    flag.Parse()

    if *keygen {
		fmt.Println("Generating key")
		sk := nostr.GeneratePrivateKey()
		pk, err := nostr.GetPublicKey(sk)
		if err != nil {
			log.Fatal("unable to generate public key")
		}
		ns, err := nostr.EncodePrivateKey(sk)
		if err != nil {
			log.Fatal("unable to generate public key")
		}
		np, err := nostr.EncodePublicKey(pk)
		if err != nil {
			log.Fatal("unable to generate public key")
		}
		fmt.Printf("nsec: %s\n", ns)
		fmt.Printf("npub: %s", np)
    }

    if *decode != "" {
		pubkey, err := nostr.DecodeBech32(*decode)
		if err != nil {
			log.Fatal("unable to generate public key")
		}
		fmt.Printf("%s", pubkey)
    }

    // Encoding flags

    var nsec, npub string
	fs := flag.NewFlagSet("encode", flag.ExitOnError)
	fs.StringVar(&nsec, "nsec", "", "subcommand 1 flag 'a'")
	fs.StringVar(&npub, "npub", "", "subcommand 1 flag 'b'")

	switch os.Args[1] {
	case "encode":
		err := fs.Parse(os.Args[2:])
        if err != nil {
            log.Fatalln("unsupported key type")
        }

        if nsec != "" {
            key, err := nostr.EncodePrivateKey(nsec)
            if err != nil {
                log.Fatal("unable to encode secret key")
            }
            fmt.Println(key)
        } else if npub != "" {
            key, err := nostr.EncodePublicKey(npub)
            if err != nil {
                log.Fatal("unable to encode public key")
            }
            fmt.Println(key)
        } else {
            log.Fatalln("unsupported key type")
        }
	}
}
