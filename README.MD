## Info

Worklog jira time in CLI.

## Install

### requirements:

`go(lang) version higher than 1.16`

### steps:

* `make install` compile and cp binaries to `/usr/local/bin/`
* `jlog install`
    * enter domain (domain.atlassian.net)
    * enter jira username/email (myjiraname@organisation.com)
    * enter jira api token (generated token https://id.atlassian.com/manage-profile/security/api-tokens)

Everything will be saved in the configuration file: `~/.jlog/config/config.yaml`

### usage:

* `jlog <taskNr> <time>` for example: `jlog DDD-2323 15m`