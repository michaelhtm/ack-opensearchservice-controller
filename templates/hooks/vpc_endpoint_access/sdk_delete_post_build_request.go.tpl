	input.Account, input.Service = vpcEndpointAccessPrincipal(aws.ToString(r.ko.Spec.Principal))
	if r.ko.Spec.ServiceOptions != nil && len(r.ko.Spec.ServiceOptions.SupportedRegions) > 0 {
		input.ServiceOptions = &svcsdktypes.ServiceOptions{
			SupportedRegions: aws.ToStringSlice(r.ko.Spec.ServiceOptions.SupportedRegions),
		}
	}
