# Concourse Resource for Marketplace

Interact with the [VMware Marketplace](https://marketplace.cloud.vmware.com/) from concourse.

## Installing

The recommended method to use this resource is with
[resource_types](https://concourse-ci.org/resource-types.html) in the
pipeline config as follows:

```yaml
---
resource_types:
- name: marketplace
  type: registry-image
  source:
    repository: projects.registry.vmware.com/tanzu_isv_engineering/mkpcli_concourse_resource
```

## Source configuration

```yaml
resources:
- name: greenplum
  type: marketplace
  source:
    csp_api_token: {{api-token}}
    product_slug: vmware-tanzu-greenplum-r-11
```

* `csp_api_token`: *Required string.*

  API Token from your VMware Cloud Service Portal.

* `product_slug`: *Required string.*

  Slug of the product on the VMware Marketplace.

* `marketplace_env`: *Optional string.*

  Marketplace environment to use. Either `staging` or `production`

  Defaults to `production`.

## Behavior

### `check`: check for new product versions on the VMware Marketplace

Discovers all versions of the provided product.

### `in`: download a product asset from the VMware Marketplace

Downloads a product asset from the VMware Marketplace.

The details for the product is written to both `product.json` in the working directory (typically `/tmp/build/get`).
Use this to programmatically get information for the product.

A version file is written to `version`

#### Parameters

* `filename`: *Required string (unless `skip_download` is `true`).*

  The name of the file to use when saving the downloaded asset.

* `filter`: *Optional string.*

  A string to select a specific asset attached to a product.

* `accept_eula`: *Optional boolean.*

  Accepts the EULA for the product when downloading.

* `skip_download`: *Optional boolean.*

  If `true`, do not download an asset, but still get the product.json file.

```yaml
resource:
  - name: nginx
    type: marketplace
    source:
      csp_api_token: {{api-token}}
      product_slug: nginx

jobs:
- name: deploy-nginx-chart
  plan:
  - get: cluster
  - get: tasks
  - get: nginx
    params:
      filename: "nginx.tar"
      accept_eula: true
  - task: deploy-chart
    image: image
    file: tasks/deploy-chart.yml      
```

### `out`: upload and attach product assets to a product in the VMware Marketplace

Not yet implemented. It's on the roadmap!

## Developing

### Prerequisites

A valid installation of golang 1.18 is required.

### Dependencies

We use [go modules](https://github.com/golang/go/wiki/Modules) for dependencies, so you will have to make sure to turn them on with `GO111MODULE=on`.

### Running the tests

Run the tests with the make file:

```
make test
```

### Contributing

Please see our [Code of Conduct](CODE-OF-CONDUCT.md) and [Contributors guide](CONTRIBUTING.md).
