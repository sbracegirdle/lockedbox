# Lockedbox

Lockedbox is a small tool to make it very difficult to access a secret (e.g. Steam or Netflix password), except at a given time of day or week.

This can be useful if you need a little help in being disciplined around usage of certain services (e.g. gaming, entertainment, etc).


## Installation and pre-requisites

If you wish to build lockedbox for yourself, you need golang; https://golang.org/doc/install.

## How to build

**NOTE: The program has the blocked/unblocked times hard-coded into it. You must change these in the `lockedbox.go` file first before building.**

**NOTE: The key used to encrypt/decrypt your secrets is generated mathematically by a formula on line 66. Tweak these values slightly after building to ensure you cannot modify the program after encryption to easily decrypt your secrets.**

1. Open lockedbox.go
2. Go to line 46 ("Decryption only allowed at certain times")
3. Adjust the logic to suit your needs (e.g. I want it allowed every day after 8PM)
4. Run build in terminal:

```s
go build
```

5. Adjust values on line 66 slightly, save file

This should create an executable suitable for usage on your system.

## How to use

```s
lockedbox [encrypt|decrypt] [secret|cypher]
```

### Encrypting

```s
lockedbox encrypt MySteamPassword
```

This will encrypt your secret.

**WARNING: Please do not lose your encrypted secret.**

**WARNING: Please test decryption and give a backup to a family member or friend before deleting your unecrypted copy of the secret**

### Decrypting

On the configured days and hours, you will be able to decrypt your secret:

```s
lockedbox decrypt 3423-fsdf3-sdfaerASDGjhasgdhjsdagsda
```


## License

MIT. See [LICENSE](LICENSE) for details.
