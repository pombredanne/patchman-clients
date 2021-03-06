# Go API client for rbac

The API for Role Based Access Control.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
```

Put the package under your project folder and add the following in import:

```golang
import "./rbac"
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost/api/rbac/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccessApi* | [**GetPrincipalAccess**](docs/AccessApi.md#getprincipalaccess) | **Get** /access/ | Get the permitted access for a principal in the tenant (defaults to principal from the identity header)
*GroupApi* | [**AddPrincipalToGroup**](docs/GroupApi.md#addprincipaltogroup) | **Post** /groups/{uuid}/principals/ | Add a principal to a group in the tenant
*GroupApi* | [**AddRoleToGroup**](docs/GroupApi.md#addroletogroup) | **Post** /groups/{uuid}/roles/ | Add a role to a group in the tenant
*GroupApi* | [**CreateGroup**](docs/GroupApi.md#creategroup) | **Post** /groups/ | Create a group in a tenant
*GroupApi* | [**DeleteGroup**](docs/GroupApi.md#deletegroup) | **Delete** /groups/{uuid}/ | Delete a group in the tenant
*GroupApi* | [**DeletePrincipalFromGroup**](docs/GroupApi.md#deleteprincipalfromgroup) | **Delete** /groups/{uuid}/principals/ | Remove a principal from a group in the tenant
*GroupApi* | [**DeleteRoleFromGroup**](docs/GroupApi.md#deleterolefromgroup) | **Delete** /groups/{uuid}/roles/ | Remove a role from a group in the tenant
*GroupApi* | [**GetGroup**](docs/GroupApi.md#getgroup) | **Get** /groups/{uuid}/ | Get a group in the tenant
*GroupApi* | [**GetPrincipalsFromGroup**](docs/GroupApi.md#getprincipalsfromgroup) | **Get** /groups/{uuid}/principals/ | Get a list of principals from a group in the tenant
*GroupApi* | [**ListGroups**](docs/GroupApi.md#listgroups) | **Get** /groups/ | List the groups for a tenant
*GroupApi* | [**ListRolesForGroup**](docs/GroupApi.md#listrolesforgroup) | **Get** /groups/{uuid}/roles/ | List the roles for a group in the tenant
*GroupApi* | [**UpdateGroup**](docs/GroupApi.md#updategroup) | **Put** /groups/{uuid}/ | Udate a group in the tenant
*PermissionApi* | [**ListPermissionOptions**](docs/PermissionApi.md#listpermissionoptions) | **Get** /permissions/options/ | List the available options for fields of permissions for a tenant
*PermissionApi* | [**ListPermissions**](docs/PermissionApi.md#listpermissions) | **Get** /permissions/ | List the permissions for a tenant
*PolicyApi* | [**CreatePolicies**](docs/PolicyApi.md#createpolicies) | **Post** /policies/ | Create a policy in a tenant
*PolicyApi* | [**DeletePolicy**](docs/PolicyApi.md#deletepolicy) | **Delete** /policies/{uuid}/ | Delete a policy in the tenant
*PolicyApi* | [**GetPolicy**](docs/PolicyApi.md#getpolicy) | **Get** /policies/{uuid}/ | Get a policy in the tenant
*PolicyApi* | [**ListPolicies**](docs/PolicyApi.md#listpolicies) | **Get** /policies/ | List the policies in the tenant
*PolicyApi* | [**UpdatePolicy**](docs/PolicyApi.md#updatepolicy) | **Put** /policies/{uuid}/ | Update a policy in the tenant
*PrincipalApi* | [**ListPrincipals**](docs/PrincipalApi.md#listprincipals) | **Get** /principals/ | List the principals for a tenant
*RoleApi* | [**CreateRoles**](docs/RoleApi.md#createroles) | **Post** /roles/ | Create a roles for a tenant
*RoleApi* | [**DeleteRole**](docs/RoleApi.md#deleterole) | **Delete** /roles/{uuid}/ | Delete a role in the tenant
*RoleApi* | [**GetRole**](docs/RoleApi.md#getrole) | **Get** /roles/{uuid}/ | Get a role in the tenant
*RoleApi* | [**GetRoleAccess**](docs/RoleApi.md#getroleaccess) | **Get** /roles/{uuid}/access/ | Get access for a role in the tenant
*RoleApi* | [**ListRoles**](docs/RoleApi.md#listroles) | **Get** /roles/ | List the roles for a tenant
*RoleApi* | [**UpdateRole**](docs/RoleApi.md#updaterole) | **Put** /roles/{uuid}/ | Update a Role in the tenant
*StatusApi* | [**GetStatus**](docs/StatusApi.md#getstatus) | **Get** /status/ | Obtain server status


## Documentation For Models

 - [Access](docs/Access.md)
 - [AccessPagination](docs/AccessPagination.md)
 - [AccessPaginationAllOf](docs/AccessPaginationAllOf.md)
 - [AdditionalGroup](docs/AdditionalGroup.md)
 - [Error](docs/Error.md)
 - [Error403](docs/Error403.md)
 - [Error403Errors](docs/Error403Errors.md)
 - [ErrorErrors](docs/ErrorErrors.md)
 - [Group](docs/Group.md)
 - [GroupOut](docs/GroupOut.md)
 - [GroupOutAllOf](docs/GroupOutAllOf.md)
 - [GroupPagination](docs/GroupPagination.md)
 - [GroupPaginationAllOf](docs/GroupPaginationAllOf.md)
 - [GroupPrincipalIn](docs/GroupPrincipalIn.md)
 - [GroupRoleIn](docs/GroupRoleIn.md)
 - [GroupRolesPagination](docs/GroupRolesPagination.md)
 - [GroupWithPrincipals](docs/GroupWithPrincipals.md)
 - [GroupWithPrincipalsAllOf](docs/GroupWithPrincipalsAllOf.md)
 - [GroupWithPrincipalsAndRoles](docs/GroupWithPrincipalsAndRoles.md)
 - [GroupWithPrincipalsAndRolesAllOf](docs/GroupWithPrincipalsAndRolesAllOf.md)
 - [InlineResponse200](docs/InlineResponse200.md)
 - [ListPagination](docs/ListPagination.md)
 - [PaginationLinks](docs/PaginationLinks.md)
 - [PaginationMeta](docs/PaginationMeta.md)
 - [Permission](docs/Permission.md)
 - [PermissionOptionsPagination](docs/PermissionOptionsPagination.md)
 - [PermissionOptionsPaginationAllOf](docs/PermissionOptionsPaginationAllOf.md)
 - [PermissionPagination](docs/PermissionPagination.md)
 - [PermissionPaginationAllOf](docs/PermissionPaginationAllOf.md)
 - [Policy](docs/Policy.md)
 - [PolicyExtended](docs/PolicyExtended.md)
 - [PolicyExtendedAllOf](docs/PolicyExtendedAllOf.md)
 - [PolicyIn](docs/PolicyIn.md)
 - [PolicyInAllOf](docs/PolicyInAllOf.md)
 - [PolicyPagination](docs/PolicyPagination.md)
 - [PolicyPaginationAllOf](docs/PolicyPaginationAllOf.md)
 - [Principal](docs/Principal.md)
 - [PrincipalIn](docs/PrincipalIn.md)
 - [PrincipalOut](docs/PrincipalOut.md)
 - [PrincipalPagination](docs/PrincipalPagination.md)
 - [PrincipalPaginationAllOf](docs/PrincipalPaginationAllOf.md)
 - [ResourceDefinition](docs/ResourceDefinition.md)
 - [ResourceDefinitionFilter](docs/ResourceDefinitionFilter.md)
 - [Role](docs/Role.md)
 - [RoleIn](docs/RoleIn.md)
 - [RoleInAllOf](docs/RoleInAllOf.md)
 - [RoleOut](docs/RoleOut.md)
 - [RoleOutAllOf](docs/RoleOutAllOf.md)
 - [RoleOutDynamic](docs/RoleOutDynamic.md)
 - [RoleOutDynamicAllOf](docs/RoleOutDynamicAllOf.md)
 - [RolePagination](docs/RolePagination.md)
 - [RolePaginationDynamic](docs/RolePaginationDynamic.md)
 - [RolePaginationDynamicAllOf](docs/RolePaginationDynamicAllOf.md)
 - [RoleWithAccess](docs/RoleWithAccess.md)
 - [Status](docs/Status.md)
 - [Timestamped](docs/Timestamped.md)
 - [Uuid](docs/Uuid.md)


## Documentation For Authorization



## basic_auth

- **Type**: HTTP basic authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
    UserName: "username",
    Password: "password",
})
r, err := client.Service.Operation(auth, args)
```



## Author



