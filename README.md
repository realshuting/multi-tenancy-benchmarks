# Multi-tenancy-benchmarks

Add description of this repo. 

[Proposal doc](https://docs.google.com/document/d/1O-G8jEpiJxOeYx9Pd2OuOSb8859dTRNmgBC5gJv0krE/edit?usp=sharing).

## Benchmarks
|    Level <img width=30/>    |       Type              |        Category     |             Test              |
|---------------------------------|--------------------------------|-------------------------|---------------------------|
|   Level 1     |     Behavioral    |  Control Plane Protection  |  [Ensure that Tenant A cannot list non namespaced resources](e2e/tests/tenantaccess/README.md)|
|   Level 1     |     Behavioral    |  Tenant Protection  |  Ensure that Tenant A cannot list namespaced resources from Tenant B|
|   Level 1     |     Configuration    |  Fairness  |  Ensure that Tenant A cannot starve other tenants from cluster wide resources|
|   Level 1     |     Behavioral    |  Tenant Isolation  |  [Ensure that users of Tenant A cannot modify Resource Quotas](e2e/tests/resourcequotas)|
|   Level 1     |     Behavioral    |  Tenant Isolation  |  Ensure that users of Tenant A cannot modify resources managed by the cluster administrator|
|   Level 1     |     Behavioral    |  Network Protection & Isolation  |  Ensure that users of Tenant A cannot connect to pods running in Tenant B|
|   Level 1     |     Behavioral    |  Host Protection  |  Ensure that users of Tenant A cannot use hostpaths|
|   Level 1     |     Behavioral    |  Host Protection  |  Ensure that users of Tenant A cannot use HostPort|


## Documentation
- [Multi-Tenancy Profile Definitions](documentation/definitions.md)

- [Benchmark Types](documentation/types.md)

- [Benchmark Categories](documentation/catagories.md)

- [Getting Started](documentation/run.md)

## License

## Roadmap

## Contributing