# cleanup-aws-access-keys
A cloud security tool to search and clean up unused AWS access keys, written in Go.

## Features:
* Find unused access keys (e.g: access keys unused for more than 90 days, access keys created both never used)
* Deactivate/activate access keys easily based on search criteria.
* Delete access keys based on search criteria.
* Auto-approve flag to run non-interactively (e.g: a cron job to deactivate access keys unused for more 90 days)

## What is an AWS access keys?
* Access keys are long-term credentials for an IAM user or the AWS account root user.
* You can use access keys to make programmatic calls to AWS via AWS CLI, AWS SDKs, or direct AWS API calls.
* An IAM user is only allowed to have maximum of two access keys (active or inactive) at a time.
* Access keys consist of two parts: an access key ID (e.g: `AKIAIOSFODNN7EXAMPLE`) and a secret access key (e.g: `wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY`). Like a user name and password, you must use both the access key ID and secret access key together to authenticate your requests. Manage your access keys as securely as you do your user name and password.
* If you lose or forget your secret key, you cannot retrieve it. Instead, create a new access key and make the old key inactive and delete it.


> **Warning:** Never post your secret access key on public platforms, such as GitHub. This can compromise your account security. As a best practice, it's recommended to rotate your keys frequently.

> __Best Practices:__ Use temporary security credentials (IAM roles) instead of access keys, and disable any AWS account root user access keys. For more information, see [Best Practices for Managing AWS Access Keys](https://docs.aws.amazon.com/general/latest/gr/aws-access-keys-best-practices.html) in the Amazon Web Services General Reference.

[Learn more about AWS access keys](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_access-keys.html?icmpid=docs_iam_console)

## Usage:
```
$ ./cleanup-aws-access-keys 
A cloud security tool to search and clean up unused AWS access keys (https://github.com/tuladhar/cleanup-aws-access-keys)

Usage:
  cleanup-aws-access-keys [command]

Available Commands:
  activate    Activate access key(s)
  completion  Generate the autocompletion script for the specified shell
  deactivate  Deactivate access key(s)
  delete      Delete access key(s)
  help        Help about any command
  search      Search for access key(s)

Flags:
  -h, --help      help for cleanup-aws-access-keys
  -v, --version   version for cleanup-aws-access-keys

Use "cleanup-aws-access-keys [command] --help" for more information about a command.
```

## Examples:

Search for active access keys unused for more than 90 days.
```
./cleanup-aws-access-keys search --last-used 90 --status active
```

Search for access keys created but never used.
```
./cleanup-aws-access-keys search --last-used -1
```

Search for inactive access keys.
```
./cleanup-aws-access-keys search --status inactive
```

Deactivate access keys unused for more than 90 days.
```
./cleanup-aws-access-keys deactivate --last-used 90
```
> Hint: Use `--auto-approve` flag to skip interactive prompt.

Deactivate access keys of specific username.
```
./cleanup-aws-access-keys deactivate --username jeff.bezos
```

Delete access keys unused for more than 180 days.
```
./cleanup-aws-access-keys delete --last-used 180
```

Delete inactive access keys of specific username.
```
./cleanup-aws-access-keys delete --status inactive --username jeff.bezos
```

## Installation
Binary is available for Linux, Windows and Mac OS (amd64 and arm64). Download the binary for your respective platform from the [releases page](https://github.com/tuladhar/cleanup-aws-access-keys/releases).

### Linux:
```
```

## Development
If you wish to contribute or compile from source code, you'll first need Go installed on your machine. Go version [Go v.1.9](https://go.dev/dl/)+ is required.

- Clone the repository
```
git clone https://github.com/tuladhar/cleanup-aws-access-keys
```
- Add missing modules
```
go mod tidy
```
- Modify the code, and build the binary or run directly
```
go run main.go
// or
go build
```

## Author
* Puru Tuladhar (https://github.com/tuladhar)