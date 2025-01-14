/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	common "github.com/upbound/provider-aws/config/common"
	client "sigs.k8s.io/controller-runtime/pkg/client"

	// ResolveReferences of this Policy.
	apisresolver "github.com/upbound/provider-aws/internal/apis"
)

func (mg *Policy) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("appautoscaling.aws.upbound.io", "v1beta1", "Target", "TargetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceID),
			Extract:      resource.ExtractParamPath("resource_id", false),
			Reference:    mg.Spec.ForProvider.ResourceIDRef,
			Selector:     mg.Spec.ForProvider.ResourceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceID")
	}
	mg.Spec.ForProvider.ResourceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("appautoscaling.aws.upbound.io", "v1beta1", "Target", "TargetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ScalableDimension),
			Extract:      resource.ExtractParamPath("scalable_dimension", false),
			Reference:    mg.Spec.ForProvider.ScalableDimensionRef,
			Selector:     mg.Spec.ForProvider.ScalableDimensionSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ScalableDimension")
	}
	mg.Spec.ForProvider.ScalableDimension = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ScalableDimensionRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("appautoscaling.aws.upbound.io", "v1beta1", "Target", "TargetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServiceNamespace),
			Extract:      resource.ExtractParamPath("service_namespace", false),
			Reference:    mg.Spec.ForProvider.ServiceNamespaceRef,
			Selector:     mg.Spec.ForProvider.ServiceNamespaceSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ServiceNamespace")
	}
	mg.Spec.ForProvider.ServiceNamespace = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServiceNamespaceRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ScheduledAction.
func (mg *ScheduledAction) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("appautoscaling.aws.upbound.io", "v1beta1", "Target", "TargetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceID),
			Extract:      resource.ExtractParamPath("resource_id", false),
			Reference:    mg.Spec.ForProvider.ResourceIDRef,
			Selector:     mg.Spec.ForProvider.ResourceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceID")
	}
	mg.Spec.ForProvider.ResourceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("appautoscaling.aws.upbound.io", "v1beta1", "Target", "TargetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ScalableDimension),
			Extract:      resource.ExtractParamPath("scalable_dimension", false),
			Reference:    mg.Spec.ForProvider.ScalableDimensionRef,
			Selector:     mg.Spec.ForProvider.ScalableDimensionSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ScalableDimension")
	}
	mg.Spec.ForProvider.ScalableDimension = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ScalableDimensionRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("appautoscaling.aws.upbound.io", "v1beta1", "Target", "TargetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServiceNamespace),
			Extract:      resource.ExtractParamPath("service_namespace", false),
			Reference:    mg.Spec.ForProvider.ServiceNamespaceRef,
			Selector:     mg.Spec.ForProvider.ServiceNamespaceSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ServiceNamespace")
	}
	mg.Spec.ForProvider.ServiceNamespace = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServiceNamespaceRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("appautoscaling.aws.upbound.io", "v1beta1", "Target", "TargetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ResourceID),
			Extract:      resource.ExtractParamPath("resource_id", false),
			Reference:    mg.Spec.InitProvider.ResourceIDRef,
			Selector:     mg.Spec.InitProvider.ResourceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ResourceID")
	}
	mg.Spec.InitProvider.ResourceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ResourceIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("appautoscaling.aws.upbound.io", "v1beta1", "Target", "TargetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ScalableDimension),
			Extract:      resource.ExtractParamPath("scalable_dimension", false),
			Reference:    mg.Spec.InitProvider.ScalableDimensionRef,
			Selector:     mg.Spec.InitProvider.ScalableDimensionSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ScalableDimension")
	}
	mg.Spec.InitProvider.ScalableDimension = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ScalableDimensionRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("appautoscaling.aws.upbound.io", "v1beta1", "Target", "TargetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ServiceNamespace),
			Extract:      resource.ExtractParamPath("service_namespace", false),
			Reference:    mg.Spec.InitProvider.ServiceNamespaceRef,
			Selector:     mg.Spec.InitProvider.ServiceNamespaceSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ServiceNamespace")
	}
	mg.Spec.InitProvider.ServiceNamespace = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ServiceNamespaceRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Target.
func (mg *Target) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RoleArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.ForProvider.RoleArnRef,
			Selector:     mg.Spec.ForProvider.RoleArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RoleArn")
	}
	mg.Spec.ForProvider.RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RoleArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RoleArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.InitProvider.RoleArnRef,
			Selector:     mg.Spec.InitProvider.RoleArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RoleArn")
	}
	mg.Spec.InitProvider.RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RoleArnRef = rsp.ResolvedReference

	return nil
}
