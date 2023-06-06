# bosh-cpi-go: Library for writing BOSH CPIs in Go

See [docs/example.go](docs/example.go) for an example & [apiv1/interfaces.go](apiv1/interfaces.go) for interface details.

CPIs using this library:

- [Warden CPI](https://github.com/cppforlife/bosh-warden-cpi-release)
- [VirtualBox CPI](https://github.com/cppforlife/bosh-virtualbox-cpi-release)
- [Docker CPI](https://github.com/cppforlife/bosh-docker-cpi-release)
- [Kubernetes CPI](https://github.com/bosh-cpis/bosh-kubernetes-cpi-release)

## Test
run ginkgo v1
`go run github.com/onsi/ginkgo/ginkgo@v1.16.1 -r .`
## Todos

- rename apiv1 to api package
