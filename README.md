harmony-go-sdk-sample
---

This is a small sample of an application that uses [harmony-one/go-sdk](https://github.com/harmony-one/go-sdk) with Go Contract Bindings.

Here, we investigated the cause of duplicate symbols error based on cgo dependency that occurred during naive development.

The branch where the error occurs:
https://github.com/datachainlab/harmony-go-sdk-sample/tree/duplicate-symbol-error

Sample fix to avoid duplicate symbols:
https://github.com/datachainlab/harmony-go-sdk-sample/pull/2

You can check the build status in both cases from Github actions.
