# go-audit

Audit API SDK.

Contains common audit configuration, structures and methods for calling audit API.

NB! go-audit is used as a dependency for both audit API and APIs that are making calls to the audit API.

## Built with Azugo Go Web Framework

This project is built using the [Azugo Go Web Framework](https://azugo.io), a powerful and flexible framework for building modern web applications in Go. Check out the [Azugo GitHub page](https://github.com/azugo) for more information and documentation.

## Usage

1. Add dependency

    ```sh
    go get -u github.com/nobid-lsp-latvia/go-audit
    ```

2. Set information about person data access in `AuditRequest` object.

3. Use `PersonRequest` method to make a call to audit API.
