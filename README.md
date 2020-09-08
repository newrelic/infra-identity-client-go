[![Community Project header](https://github.com/newrelic/opensource-website/raw/master/src/images/categories/Community_Project.png)](https://opensource.newrelic.com/oss-category/#community-project)

# New Relic Infrastructure Identity - Go lang client [build badges go here when available]

Go auto-generated client for the OpenAPI Infrastructure Identity Client

## Installation

Install with go modules:

```
module your/module/path

go 1.13

require github.com/newrelic/infra-identity-client-go/identity v1.5.1
```

Install with go get:

```
go get github.com/newrelic/infra-identity-client-go/identity
```

**Please note**: the module path ends with `/identity`.

## Get Started

```go
package main

import (
    "context"
    "fmt"
    "github.com/newrelic/infra-identity-client-go/identity"
)

func main() {
    ac := identity.NewAPIClient(identity.NewConfiguration())
    connectResponse, httpResponse, err := ac.DefaultApi.ConnectPost(
        context.Background(), "user-agent", "license-key",
        identity.ConnectRequest{
            Fingerprint: identity.Fingerprint{
                FullHostname: "foo.example.org",
                Hostname: "foo",
            },
        }, nil)
    fmt.Printf("err = %v\n", err)
    fmt.Printf("connectResponse = %v\n", connectResponse)
    fmt.Printf("httpResponse = %v\n", httpResponse)
}
```

## Support

The New Relic Explorers Hub is an online forum where customers can interact with New Relic employees as well as other customers to get help and share best practices. Like all official New Relic open source projects, there's a related Community topic in the New Relic Explorers Hub. You can find this project's topic/threads here:

> https://discuss.newrelic.com/c/full-stack-observability/infrastructure

## Contributing
We encourage your contributions to improve New Relic Infrastructure Identity - Go lang client! Keep in mind when you submit your pull request, you'll need to sign the CLA via the click-through using CLA-Assistant. You only have to sign the CLA one time per project.
If you have any questions, or to execute our corporate CLA, required if your contribution is on behalf of a company,  please drop us an email at opensource@newrelic.com.

## License
New Relic Infrastructure Identity - Go lang client is licensed under the [Apache 2.0](http://apache.org/licenses/LICENSE-2.0.txt) License.
