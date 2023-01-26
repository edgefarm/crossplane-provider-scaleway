/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Backend.
func (mg *Backend) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LBID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.LBIDRef,
		Selector:     mg.Spec.ForProvider.LBIDSelector,
		To: reference.To{
			List:    &LBList{},
			Managed: &LB{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LBID")
	}
	mg.Spec.ForProvider.LBID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LBIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Certificate.
func (mg *Certificate) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LBID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.LBIDRef,
		Selector:     mg.Spec.ForProvider.LBIDSelector,
		To: reference.To{
			List:    &LBList{},
			Managed: &LB{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LBID")
	}
	mg.Spec.ForProvider.LBID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LBIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Frontend.
func (mg *Frontend) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.BackendID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.BackendIDRef,
		Selector:     mg.Spec.ForProvider.BackendIDSelector,
		To: reference.To{
			List:    &BackendList{},
			Managed: &Backend{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.BackendID")
	}
	mg.Spec.ForProvider.BackendID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BackendIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LBID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.LBIDRef,
		Selector:     mg.Spec.ForProvider.LBIDSelector,
		To: reference.To{
			List:    &LBList{},
			Managed: &LB{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LBID")
	}
	mg.Spec.ForProvider.LBID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LBIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this LB.
func (mg *LB) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.IPID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.IPIDRef,
		Selector:     mg.Spec.ForProvider.IPIDSelector,
		To: reference.To{
			List:    &IPList{},
			Managed: &IP{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.IPID")
	}
	mg.Spec.ForProvider.IPID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.IPIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Route.
func (mg *Route) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.BackendID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.BackendIDRef,
		Selector:     mg.Spec.ForProvider.BackendIDSelector,
		To: reference.To{
			List:    &BackendList{},
			Managed: &Backend{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.BackendID")
	}
	mg.Spec.ForProvider.BackendID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BackendIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FrontendID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.FrontendIDRef,
		Selector:     mg.Spec.ForProvider.FrontendIDSelector,
		To: reference.To{
			List:    &FrontendList{},
			Managed: &Frontend{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.FrontendID")
	}
	mg.Spec.ForProvider.FrontendID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FrontendIDRef = rsp.ResolvedReference

	return nil
}
