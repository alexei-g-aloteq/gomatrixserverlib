module github.com/alexei-g-aloteq/gomatrixserverlib

// To update Go proxy cache:
// GOPROXY=proxy.golang.org go list -m github.com/alexei-g-aloteq/gomatrixserverlib@vX.Y.Z
// BEFORE first usage

require (
	github.com/frankban/quicktest v1.7.2 // indirect
	github.com/google/go-cmp v0.4.0
	github.com/konsorten/go-windows-terminal-sequences v1.0.2 // indirect
	github.com/kr/pretty v0.2.0 // indirect
	github.com/matrix-org/gomatrix v0.0.0-20190528120928-7df988a63f26
	github.com/matrix-org/util v0.0.0-20190711121626-527ce5ddefc7
	github.com/miekg/dns v1.1.25
	github.com/sirupsen/logrus v1.4.2
	github.com/stretchr/testify v1.4.0 // indirect
	github.com/tidwall/gjson v1.12.1
	github.com/tidwall/sjson v1.0.3
	golang.org/x/crypto v0.0.0-20200221231518-2aa609cf4a9d
	golang.org/x/net v0.0.0-20200226121028-0de0cce0169b // indirect
	golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e // indirect
	golang.org/x/sys v0.0.0-20220405052023-b1e9470b6e64 // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/h2non/gock.v1 v1.0.14
	gopkg.in/macaroon.v2 v2.1.0
	gopkg.in/yaml.v2 v2.2.8
)

go 1.13
