# dns-notifier

A quick script that notify you each hour via pushover when your domain name is set on a DNS server.

## Usage with docker

Create `.env` file with the following content (replace with proper values):

```
PUSHOVER_APP_TOKEN=XXXXXX
PUSHOVER_USER_KEY=XXXXXX
DOMAIN=example.com

#Optional
DNS=8.8.8.8 # Default is Google DNS
```

Then it's just a matter of running `docker-compose up`.

## Standalone usage

### Install dependencies

```
go mod download && go mod verify
```

### Build

```
go build -o dns-notifier
```

### Run

```
PUSHOVER_APP_TOKEN=XXX PUSHOVER_USER_KEY=XXX DOMAIN=example.com ./dns-notifier
```
