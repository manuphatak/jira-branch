# jira-branch

## Install

**Brew**

```sh
brew install manuphatak/tap/jira-branch
```

**Go**

```
go get github.com/manuphatak/jira-branch
```

**Manual**

Checkout https://github.com/manuphatak/jira-branch/releases

## Setup

Set the following environment variables:

- `JIRA_USERNAME` This is the username/email you use to sign in.
- `JIRA_API_KEY` Can be created by navigating to https://id.atlassian.com/manage-profile/security/api-tokens.

## Usage

```sh
jira-branch <ticketUrl>
```

**Example**

```sh
$ jira-branch https://basisfin.atlassian.net/browse/BASIS-2185
Switched to branch 'BASIS-2185/payroll_migration_drop_unique_payrollconnectionid'
```
