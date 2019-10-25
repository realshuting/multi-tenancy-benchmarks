# Multi-tenancy-benchmarks

Proposal doc: https://docs.google.com/document/d/1O-G8jEpiJxOeYx9Pd2OuOSb8859dTRNmgBC5gJv0krE/edit?usp=sharing

## Multi-Tenancy Profile Definitions

**Level 1**
- isolate and protect the kubernetes control plane from tenants
- use standard Kubernetes resources
- may inhibit select Kubernetes features. For example, a tenant may not be allowed to install a CRD

**Level 2**
- may require multi-tenancy related CRDs or other Kubernetes extensions
- provides self-service creation of tenant namespaces
- provides self-service management of other namespace resources like network policies, roles, and role bindings

**Level 3**
- are intended for environments or use cases where a higher-level of multi-tenancy is paramount
- allows of all Kubernetes features. For example, a tenant can install their own CRD and different tenants may have different versions


## Benchmarks
|              Level              |              Type              |        Category     |             Test              |
|---------------------------------|--------------------------------|-------------------------|---------------------------|
|   Level 1     |     Behavioral    |  Control Plane Protection  |    [Ensure that Tenant A cannot list non namespaced resources](test/e2e/tenantaccess/README.md)|



## Write configurations

To set up the test configuration, edit this [config](test/e2e/manifest/config.yaml) file. 

For all supported configurations, refer to this [config](https://github.com/realshuting/multi-tenancy-benchmarks/blob/a4f0e1a601928c12470c6b53802d50a4f4ca6b44/test/e2e/config.go#L12) file.


### 