# Azure Data Box Module for Go

[![PkgGoDev](https://pkg.go.dev/badge/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox)](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox)

The `armdatabox` module provides operations for working with Azure Data Box.

[Source code](https://github.com/Azure/azure-sdk-for-go/tree/main/sdk/resourcemanager/databox/armdatabox)

# Getting started

## Prerequisites

- an [Azure subscription](https://azure.microsoft.com/free/)
- Go 1.16 or above

## Install the package

This project uses [Go modules](https://github.com/golang/go/wiki/Modules) for versioning and dependency management.

Install the Azure Data Box module:

```sh
go get github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox
```

## Authorization

When creating a client, you will need to provide a credential for authenticating with Azure Data Box.  The `azidentity` module provides facilities for various ways of authenticating with Azure including client/secret, certificate, managed identity, and more.

```go
cred, err := azidentity.NewDefaultAzureCredential(nil)
```

For more information on authentication, please see the documentation for `azidentity` at [pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/azidentity](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/azidentity).

## Clients

Azure Data Box modules consist of one or more clients.  A client groups a set of related APIs, providing access to its functionality within the specified subscription.  Create one or more clients to access the APIs you require using your credential.

```go
client := armdatabox.NewJobsClient(<subscription ID>, cred, nil)
```

You can use `ClientOptions` in package `github.com/Azure/azure-sdk-for-go/sdk/azcore/arm` to set endpoint to connect with public and sovereign clouds as well as Azure Stack. For more information, please see the documentation for `azcore` at [pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/azcore](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/azcore).

```go
options = arm.ClientOptions{
    Host: arm.AzureChina,
}
client := armdatabox.NewJobsClient(<subscription ID>, cred, &options)
```

## Provide Feedback

If you encounter bugs or have suggestions, please
[open an issue](https://github.com/Azure/azure-sdk-for-go/issues) and assign the `Data Box` label.

# Contributing

This project welcomes contributions and suggestions. Most contributions require
you to agree to a Contributor License Agreement (CLA) declaring that you have
the right to, and actually do, grant us the rights to use your contribution.
For details, visit [https://cla.microsoft.com](https://cla.microsoft.com).

When you submit a pull request, a CLA-bot will automatically determine whether
you need to provide a CLA and decorate the PR appropriately (e.g., label,
comment). Simply follow the instructions provided by the bot. You will only
need to do this once across all repos using our CLA.

This project has adopted the
[Microsoft Open Source Code of Conduct](https://opensource.microsoft.com/codeofconduct/).
For more information, see the
[Code of Conduct FAQ](https://opensource.microsoft.com/codeofconduct/faq/)
or contact [opencode@microsoft.com](mailto:opencode@microsoft.com) with any
additional questions or comments.