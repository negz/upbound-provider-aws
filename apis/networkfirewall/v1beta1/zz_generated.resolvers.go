/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta1 "github.com/upbound/provider-aws/apis/ec2/v1beta1"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this FirewallPolicy.
func (mg *FirewallPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.FirewallPolicy); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.FirewallPolicy[i3].StatelessRuleGroupReference); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FirewallPolicy[i3].StatelessRuleGroupReference[i4].ResourceArn),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.ForProvider.FirewallPolicy[i3].StatelessRuleGroupReference[i4].ResourceArnRef,
				Selector:     mg.Spec.ForProvider.FirewallPolicy[i3].StatelessRuleGroupReference[i4].ResourceArnSelector,
				To: reference.To{
					List:    &RuleGroupList{},
					Managed: &RuleGroup{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.FirewallPolicy[i3].StatelessRuleGroupReference[i4].ResourceArn")
			}
			mg.Spec.ForProvider.FirewallPolicy[i3].StatelessRuleGroupReference[i4].ResourceArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.FirewallPolicy[i3].StatelessRuleGroupReference[i4].ResourceArnRef = rsp.ResolvedReference

		}
	}

	return nil
}

// ResolveReferences of this RuleGroup.
func (mg *RuleGroup) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.RuleGroup); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.RuleGroup[i3].ReferenceSets); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.RuleGroup[i3].ReferenceSets[i4].IPSetReferences); i5++ {
				for i6 := 0; i6 < len(mg.Spec.ForProvider.RuleGroup[i3].ReferenceSets[i4].IPSetReferences[i5].IPSetReference); i6++ {
					rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
						CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RuleGroup[i3].ReferenceSets[i4].IPSetReferences[i5].IPSetReference[i6].ReferenceArn),
						Extract:      resource.ExtractParamPath("arn", true),
						Reference:    mg.Spec.ForProvider.RuleGroup[i3].ReferenceSets[i4].IPSetReferences[i5].IPSetReference[i6].ReferenceArnRef,
						Selector:     mg.Spec.ForProvider.RuleGroup[i3].ReferenceSets[i4].IPSetReferences[i5].IPSetReference[i6].ReferenceArnSelector,
						To: reference.To{
							List:    &v1beta1.ManagedPrefixListList{},
							Managed: &v1beta1.ManagedPrefixList{},
						},
					})
					if err != nil {
						return errors.Wrap(err, "mg.Spec.ForProvider.RuleGroup[i3].ReferenceSets[i4].IPSetReferences[i5].IPSetReference[i6].ReferenceArn")
					}
					mg.Spec.ForProvider.RuleGroup[i3].ReferenceSets[i4].IPSetReferences[i5].IPSetReference[i6].ReferenceArn = reference.ToPtrValue(rsp.ResolvedValue)
					mg.Spec.ForProvider.RuleGroup[i3].ReferenceSets[i4].IPSetReferences[i5].IPSetReference[i6].ReferenceArnRef = rsp.ResolvedReference

				}
			}
		}
	}

	return nil
}
