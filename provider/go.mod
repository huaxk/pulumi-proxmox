module github.com/huaxk/pulumi-proxmox/provider

go 1.16

replace (
	github.com/hashicorp/go-getter v1.5.0 => github.com/hashicorp/go-getter v1.4.0
	github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20201219002103-53f910a1ea61
)

require (
	github.com/Telmate/terraform-provider-proxmox v0.0.0-20210518194742-0d6e7d75f408
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.2.1
	github.com/pulumi/pulumi/sdk/v3 v3.4.0
)
