# Search and clean up unused AWS access keys
A cloud security tool to search and clean up unused AWS access keys, written in Go.

<image src="https://user-images.githubusercontent.com/5674762/188233291-1723a0b8-b00c-4ea8-8cc2-672934ea09e4.png" height=420 width=500 />

## Features
* Find unused access keys (e.g: access keys unused for more than 90 days, access keys created both never used)
* Deactivate/activate access keys easily based on search criteria.
* Delete access keys based on search criteria.
* Auto-approve flag to run non-interactively (e.g: integrate as cron job or Lambda to deactivate access keys unused for more 90 days)

## What is an AWS access keys?
* Access keys are long-term credentials for an IAM user or the AWS account root user.
* You can use access keys to make programmatic calls to AWS via AWS CLI, AWS SDKs, or direct AWS API calls.
* An IAM user is only allowed to have maximum of two access keys (active or inactive) at a time.
* Access keys consist of two parts: an access key ID (e.g: `AKIAIOSFODNN7EXAMPLE`) and a secret access key (e.g: `wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY`). Like a user name and password, you must use both the access key ID and secret access key together to authenticate your requests. Manage your access keys as securely as you do your user name and password.
* If you lose or forget your secret key, you cannot retrieve it. Instead, create a new access key and make the old key inactive and delete it.


> **Warning:** Never post your secret access key on public platforms, such as GitHub. This can compromise your account security. As a best practice, it's recommended to rotate your keys frequently.

> __Best Practices:__ Use temporary security credentials (IAM roles) instead of access keys, and disable any AWS account root user access keys. For more information, see [Best Practices for Managing AWS Access Keys](https://docs.aws.amazon.com/general/latest/gr/aws-access-keys-best-practices.html) in the Amazon Web Services General Reference.

[Learn more about AWS access keys](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_access-keys.html?icmpid=docs_iam_console)

## Usage
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

## Examples

### Search for active access keys unused for more than 90 days.
```
./cleanup-aws-access-keys search --last-used 90 --status active
```
![2022-09-03_00-34](https://user-images.githubusercontent.com/5674762/188224200-272d5b1c-c5bc-44ce-821f-1d63d473d05d.png)

### Search for access keys created but never used.
```
./cleanup-aws-access-keys search --last-used -1
```
![2022-09-03_00-37](https://user-images.githubusercontent.com/5674762/188224291-ad0f7132-e4bf-41e4-9dd0-b5f71d3a849c.png)

### Search for inactive access keys.
```
./cleanup-aws-access-keys search --status inactive
```
![2022-09-03_00-39](https://user-images.githubusercontent.com/5674762/188224305-a8b8bf4e-e24d-4e59-9528-2e49fe8a395c.png)

### Deactivate access keys unused for more than 90 days.
```
./cleanup-aws-access-keys deactivate --last-used 90
```
> Hint: Use `--auto-approve` flag to skip interactive prompt.
![2022-09-03_01-19](https://user-images.githubusercontent.com/5674762/188224695-6cbf8564-993f-474a-8596-b24dae41c10d.png)

### Deactivate access keys of specific username.
```
./cleanup-aws-access-keys deactivate --username jeff.bezos
```

### Delete access keys unused for more than 180 days.
```
./cleanup-aws-access-keys delete --last-used 180
```
![2022-09-03_01-21](https://user-images.githubusercontent.com/5674762/188224980-280fe611-0f70-48c4-acac-c4fed98b0756.png)

### Delete inactive access keys of specific username.
```
./cleanup-aws-access-keys delete --status inactive --username jeff.bezos
```

## Installation
Binary is available for Linux, Windows and Mac OS (amd64 and arm64). Download the binary for your respective platform from the [releases page](https://github.com/tuladhar/cleanup-aws-access-keys/releases).

### Linux:
```
curl -sSLO https://github.com/tuladhar/cleanup-aws-access-keys/releases/download/v1.0/cleanup-aws-access-keys-v1.0-linux-amd64.tar.gz
# or
wget https://github.com/tuladhar/cleanup-aws-access-keys/releases/download/v1.0/cleanup-aws-access-keys-v1.0-linux-amd64.tar.gz

tar zxf cleanup-aws-access-keys-v1.0-linux-amd64.tar.gz
chmod +x ./cleanup-aws-access-keys
./cleanup-aws-access-keys --version
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
