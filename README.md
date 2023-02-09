# azure-keyvault-env

Azure Keyvault Env provides a way to source variables from Azure KeyVault and pass these to other processes within the same shell.

## Usage

```bash
Usage:
  kvenv [flags]

Flags:
  -h, --help             help for kvenv
  -s, --secret strings   name of secret
  -k, --keyvault string  name of keyvault
```

```bash
kvenv -s test -u vault "/home/user/start.sh"
```

## Credits

The work is based on [Charlie Getzen's](https://github.com/cgetzen/secretsmanagerenv) work on AWS Env.
