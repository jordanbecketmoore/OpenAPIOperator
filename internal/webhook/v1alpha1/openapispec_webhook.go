/*
Copyright 2025.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	"github.com/getkin/kin-openapi/openapi3"
	openapiv1alpha1 "github.com/jordanbecketmoore/OpenAPIOperator/api/v1alpha1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
)

// nolint:unused
// log is for logging in this package.
var openapispeclog = logf.Log.WithName("openapispec-resource")

// SetupOpenApiSpecWebhookWithManager registers the webhook for OpenApiSpec in the manager.
func SetupOpenApiSpecWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).For(&openapiv1alpha1.OpenApiSpec{}).
		WithValidator(&OpenApiSpecCustomValidator{}).
		WithDefaulter(&OpenApiSpecCustomDefaulter{}).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// +kubebuilder:webhook:path=/mutate-openapi-openapis-org-v1alpha1-openapispec,mutating=true,failurePolicy=fail,sideEffects=None,groups=openapi.openapis.org,resources=openapispecs,verbs=create;update,versions=v1alpha1,name=mopenapispec-v1alpha1.kb.io,admissionReviewVersions=v1

// OpenApiSpecCustomDefaulter struct is responsible for setting default values on the custom resource of the
// Kind OpenApiSpec when those are created or updated.
//
// NOTE: The +kubebuilder:object:generate=false marker prevents controller-gen from generating DeepCopy methods,
// as it is used only for temporary operations and does not need to be deeply copied.
type OpenApiSpecCustomDefaulter struct {
	// TODO(user): Add more fields as needed for defaulting
}

var _ webhook.CustomDefaulter = &OpenApiSpecCustomDefaulter{}

// Default implements webhook.CustomDefaulter so a webhook will be registered for the Kind OpenApiSpec.
func (d *OpenApiSpecCustomDefaulter) Default(ctx context.Context, obj runtime.Object) error {
	openapispec, ok := obj.(*openapiv1alpha1.OpenApiSpec)

	if !ok {
		return fmt.Errorf("expected an OpenApiSpec object but got %T", obj)
	}
	openapispeclog.Info("Defaulting for OpenApiSpec", "name", openapispec.GetName())

	// TODO(user): fill in your defaulting logic.

	return nil
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// NOTE: The 'path' attribute must follow a specific pattern and should not be modified directly here.
// Modifying the path for an invalid path can cause API server errors; failing to locate the webhook.
// +kubebuilder:webhook:path=/validate-openapi-openapis-org-v1alpha1-openapispec,mutating=false,failurePolicy=fail,sideEffects=None,groups=openapi.openapis.org,resources=openapispecs,verbs=create;update,versions=v1alpha1,name=vopenapispec-v1alpha1.kb.io,admissionReviewVersions=v1

// OpenApiSpecCustomValidator struct is responsible for validating the OpenApiSpec resource
// when it is created, updated, or deleted.
//
// NOTE: The +kubebuilder:object:generate=false marker prevents controller-gen from generating DeepCopy methods,
// as this struct is used only for temporary operations and does not need to be deeply copied.
type OpenApiSpecCustomValidator struct {
	// TODO(user): Add more fields as needed for validation
}

var _ webhook.CustomValidator = &OpenApiSpecCustomValidator{}

// ValidateCreate implements webhook.CustomValidator so a webhook will be registered for the type OpenApiSpec.
func (v *OpenApiSpecCustomValidator) ValidateCreate(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	openapispec, ok := obj.(*openapiv1alpha1.OpenApiSpec)
	if !ok {
		return nil, fmt.Errorf("expected a OpenApiSpec object but got %T", obj)
	}
	openapispeclog.Info("Validation for OpenApiSpec upon creation", "name", openapispec.GetName())

	return nil, validateOpenApiSpec(openapispec)
}

// ValidateUpdate implements webhook.CustomValidator so a webhook will be registered for the type OpenApiSpec.
func (v *OpenApiSpecCustomValidator) ValidateUpdate(ctx context.Context, oldObj, newObj runtime.Object) (admission.Warnings, error) {
	openapispec, ok := newObj.(*openapiv1alpha1.OpenApiSpec)
	if !ok {
		return nil, fmt.Errorf("expected a OpenApiSpec object for the newObj but got %T", newObj)
	}
	openapispeclog.Info("Validation for OpenApiSpec upon update", "name", openapispec.GetName())

	return nil, validateOpenApiSpec(openapispec)
}

// ValidateDelete implements webhook.CustomValidator so a webhook will be registered for the type OpenApiSpec.
func (v *OpenApiSpecCustomValidator) ValidateDelete(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	openapispec, ok := obj.(*openapiv1alpha1.OpenApiSpec)
	if !ok {
		return nil, fmt.Errorf("expected a OpenApiSpec object but got %T", obj)
	}
	openapispeclog.Info("Validation for OpenApiSpec upon deletion", "name", openapispec.GetName())

	// TODO(user): fill in your validation logic upon object deletion.

	return nil, nil
}

func validateOpenApiSpec(oapis *openapiv1alpha1.OpenApiSpec) error {
	var allErrs field.ErrorList
	if err := validateOpenApiSpecDocument(oapis); err != nil {
		allErrs = append(allErrs, err)
	}

	if len(allErrs) == 0 {
		return nil
	}

	return apierrors.NewInvalid(
		schema.GroupKind{
			Group: "openapi.openapis.org",
			Kind:  "OpenApiSpec",
		},
		oapis.Name,
		allErrs,
	)

}

func validateOpenApiSpecDocument(oapis *openapiv1alpha1.OpenApiSpec) *field.Error {
	documentString := oapis.Spec.Document
	// Create openapi3 loader
	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = false

	// Load document
	document, err := loader.LoadFromData([]byte(documentString))
	if err != nil {
		return field.Invalid(field.NewPath("spec").Child("document"), document, err.Error())
	}

	// Validate document
	if err := document.Validate(loader.Context); err != nil {
		return field.Invalid(field.NewPath("spec").Child("document"), document, err.Error())
	}
	return nil
}
