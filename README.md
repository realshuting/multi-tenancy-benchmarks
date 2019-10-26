# Multi-tenancy-benchmarks

Add description of this repo. 

[Proposal doc](https://docs.google.com/document/d/1O-G8jEpiJxOeYx9Pd2OuOSb8859dTRNmgBC5gJv0krE/edit?usp=sharing).

## Benchmarks
|    Level <img width=30/>    |       Type              |        Category     |             Test              |
|---------------------------------|--------------------------------|-------------------------|---------------------------|
|   Level 1     |     Behavioral    |  Control Plane Protection  |  [Ensure that Tenant A cannot list non namespaced resources](e2e/tests/tenantaccess)|
|   Level 1     |     Behavioral    |  Tenant Protection  |  [Ensure that Tenant A cannot list namespaced resources from Tenant B](e2e/tests/tenantprotection)|
|   Level 1     |     Configuration    |  Fairness  |  [Ensure that Tenant A cannot starve other tenants from cluster wide resources](e2e/tests/resourcequotas)|
|   Level 1     |     Behavioral    |  Tenant Isolation  |  [Ensure that users of Tenant A cannot modify Resource Quotas](e2e/tests/modify_resourcequotas)|
|   Level 1     |     Behavioral    |  Tenant Isolation  |  [Ensure that users of Tenant A cannot modify resources managed by the cluster administrator](e2e/tests/modify_admin_resource/README.md)|
|   Level 1     |     Behavioral    |  Network Protection & Isolation  |  [Ensure that users of Tenant A cannot connect to pods running in Tenant B](e2e/tests/network_isolation)|
|   Level 1     |     Behavioral    |  Host Protection  |  [Ensure that users of Tenant A cannot use hostpaths](e2e/tests/deny_hostpaths)|
|   Level 1     |     Behavioral    |  Host Protection  |  [Ensure that users of Tenant A cannot use NodePort](e2e/tests/deny_nodeports)|
|   Level 1     |     Behavioral    |  Host Protection  |  [Ensure that users of Tenant A cannot use HostPort](e2e/tests/deny_hostports/README.md)|


## Documentation
- [Multi-Tenancy Profile Definitions](documentation/definitions.md)

- [Benchmark Types](documentation/types.md)

- [Benchmark Categories](documentation/catagories.md)

- [Getting Started](documentation/run.md)

## License

## Roadmap

## Contributing