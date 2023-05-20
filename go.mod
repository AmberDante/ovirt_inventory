module ovirt_inventory

go 1.19

require (
	github.com/ovirt/go-ovirt-client-log/v3 v3.0.0
	github.com/ovirt/go-ovirt-client/v3 v3.0.0-alpha1
	golang.org/x/oauth2 v0.5.0
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/ovirt/go-ovirt v0.0.0-20220427092237-114c47f2835c // indirect
	golang.org/x/net v0.7.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)

replace github.com/ovirt/go-ovirt-client/v3 v3.0.0-alpha1 => github.com/AmberDante/go-ovirt-client/v3 v3.0.1-alpha1
