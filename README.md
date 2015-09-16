This is our custom plugin for the null provider / null_resource that we will need to use until https://github.com/hashicorp/terraform/pull/3244 is merged.

Just run `go install` in the root directory and it will overwrite the default null_resource definition on your GOPATH.
