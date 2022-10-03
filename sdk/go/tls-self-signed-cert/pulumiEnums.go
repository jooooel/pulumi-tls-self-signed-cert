// Code generated by Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tlsselfsignedcert

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Algorithm string

const (
	AlgorithmRSA     = Algorithm("RSA")
	AlgorithmECDSA   = Algorithm("ECDSA")
	AlgorithmED25519 = Algorithm("ED25519")
)

func (Algorithm) ElementType() reflect.Type {
	return reflect.TypeOf((*Algorithm)(nil)).Elem()
}

func (e Algorithm) ToAlgorithmOutput() AlgorithmOutput {
	return pulumi.ToOutput(e).(AlgorithmOutput)
}

func (e Algorithm) ToAlgorithmOutputWithContext(ctx context.Context) AlgorithmOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AlgorithmOutput)
}

func (e Algorithm) ToAlgorithmPtrOutput() AlgorithmPtrOutput {
	return e.ToAlgorithmPtrOutputWithContext(context.Background())
}

func (e Algorithm) ToAlgorithmPtrOutputWithContext(ctx context.Context) AlgorithmPtrOutput {
	return Algorithm(e).ToAlgorithmOutputWithContext(ctx).ToAlgorithmPtrOutputWithContext(ctx)
}

func (e Algorithm) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Algorithm) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Algorithm) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Algorithm) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AlgorithmOutput struct{ *pulumi.OutputState }

func (AlgorithmOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Algorithm)(nil)).Elem()
}

func (o AlgorithmOutput) ToAlgorithmOutput() AlgorithmOutput {
	return o
}

func (o AlgorithmOutput) ToAlgorithmOutputWithContext(ctx context.Context) AlgorithmOutput {
	return o
}

func (o AlgorithmOutput) ToAlgorithmPtrOutput() AlgorithmPtrOutput {
	return o.ToAlgorithmPtrOutputWithContext(context.Background())
}

func (o AlgorithmOutput) ToAlgorithmPtrOutputWithContext(ctx context.Context) AlgorithmPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Algorithm) *Algorithm {
		return &v
	}).(AlgorithmPtrOutput)
}

func (o AlgorithmOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AlgorithmOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Algorithm) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AlgorithmOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AlgorithmOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Algorithm) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AlgorithmPtrOutput struct{ *pulumi.OutputState }

func (AlgorithmPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Algorithm)(nil)).Elem()
}

func (o AlgorithmPtrOutput) ToAlgorithmPtrOutput() AlgorithmPtrOutput {
	return o
}

func (o AlgorithmPtrOutput) ToAlgorithmPtrOutputWithContext(ctx context.Context) AlgorithmPtrOutput {
	return o
}

func (o AlgorithmPtrOutput) Elem() AlgorithmOutput {
	return o.ApplyT(func(v *Algorithm) Algorithm {
		if v != nil {
			return *v
		}
		var ret Algorithm
		return ret
	}).(AlgorithmOutput)
}

func (o AlgorithmPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AlgorithmPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Algorithm) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// AlgorithmInput is an input type that accepts AlgorithmArgs and AlgorithmOutput values.
// You can construct a concrete instance of `AlgorithmInput` via:
//
//          AlgorithmArgs{...}
type AlgorithmInput interface {
	pulumi.Input

	ToAlgorithmOutput() AlgorithmOutput
	ToAlgorithmOutputWithContext(context.Context) AlgorithmOutput
}

var algorithmPtrType = reflect.TypeOf((**Algorithm)(nil)).Elem()

type AlgorithmPtrInput interface {
	pulumi.Input

	ToAlgorithmPtrOutput() AlgorithmPtrOutput
	ToAlgorithmPtrOutputWithContext(context.Context) AlgorithmPtrOutput
}

type algorithmPtr string

func AlgorithmPtr(v string) AlgorithmPtrInput {
	return (*algorithmPtr)(&v)
}

func (*algorithmPtr) ElementType() reflect.Type {
	return algorithmPtrType
}

func (in *algorithmPtr) ToAlgorithmPtrOutput() AlgorithmPtrOutput {
	return pulumi.ToOutput(in).(AlgorithmPtrOutput)
}

func (in *algorithmPtr) ToAlgorithmPtrOutputWithContext(ctx context.Context) AlgorithmPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AlgorithmPtrOutput)
}

type AllowedUses string

const (
	AllowedUses_Any_extended                      = AllowedUses("any_extended")
	AllowedUses_Cert_signing                      = AllowedUses("cert_signing")
	AllowedUses_Client_auth                       = AllowedUses("client_auth")
	AllowedUses_Code_signing                      = AllowedUses("code_signing")
	AllowedUses_Content_commitment                = AllowedUses("content_commitment")
	AllowedUses_Crl_signing                       = AllowedUses("crl_signing")
	AllowedUses_Data_encipherment                 = AllowedUses("data_encipherment")
	AllowedUses_Decipher_only                     = AllowedUses("decipher_only")
	AllowedUses_Digital_signature                 = AllowedUses("digital_signature")
	AllowedUses_Email_protection                  = AllowedUses("email_protection")
	AllowedUses_Encipher_only                     = AllowedUses("encipher_only")
	AllowedUses_Ipsec_end_system                  = AllowedUses("ipsec_end_system")
	AllowedUses_Ipsec_tunnel                      = AllowedUses("ipsec_tunnel")
	AllowedUses_Ipsec_user                        = AllowedUses("ipsec_user")
	AllowedUses_Key_agreement                     = AllowedUses("key_agreement")
	AllowedUses_Key_encipherment                  = AllowedUses("key_encipherment")
	AllowedUses_Microsoft_commercial_code_signing = AllowedUses("microsoft_commercial_code_signing")
	AllowedUses_Microsoft_kernel_code_signing     = AllowedUses("microsoft_kernel_code_signing")
	AllowedUses_Microsoft_server_gated_crypto     = AllowedUses("microsoft_server_gated_crypto")
	AllowedUses_Netscape_server_gated_crypto      = AllowedUses("netscape_server_gated_crypto")
	AllowedUses_Ocsp_signing                      = AllowedUses("ocsp_signing")
	AllowedUses_Server_auth                       = AllowedUses("server_auth")
	AllowedUsesTimestamping                       = AllowedUses("timestamping")
)

func (AllowedUses) ElementType() reflect.Type {
	return reflect.TypeOf((*AllowedUses)(nil)).Elem()
}

func (e AllowedUses) ToAllowedUsesOutput() AllowedUsesOutput {
	return pulumi.ToOutput(e).(AllowedUsesOutput)
}

func (e AllowedUses) ToAllowedUsesOutputWithContext(ctx context.Context) AllowedUsesOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AllowedUsesOutput)
}

func (e AllowedUses) ToAllowedUsesPtrOutput() AllowedUsesPtrOutput {
	return e.ToAllowedUsesPtrOutputWithContext(context.Background())
}

func (e AllowedUses) ToAllowedUsesPtrOutputWithContext(ctx context.Context) AllowedUsesPtrOutput {
	return AllowedUses(e).ToAllowedUsesOutputWithContext(ctx).ToAllowedUsesPtrOutputWithContext(ctx)
}

func (e AllowedUses) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AllowedUses) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AllowedUses) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AllowedUses) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AllowedUsesOutput struct{ *pulumi.OutputState }

func (AllowedUsesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AllowedUses)(nil)).Elem()
}

func (o AllowedUsesOutput) ToAllowedUsesOutput() AllowedUsesOutput {
	return o
}

func (o AllowedUsesOutput) ToAllowedUsesOutputWithContext(ctx context.Context) AllowedUsesOutput {
	return o
}

func (o AllowedUsesOutput) ToAllowedUsesPtrOutput() AllowedUsesPtrOutput {
	return o.ToAllowedUsesPtrOutputWithContext(context.Background())
}

func (o AllowedUsesOutput) ToAllowedUsesPtrOutputWithContext(ctx context.Context) AllowedUsesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AllowedUses) *AllowedUses {
		return &v
	}).(AllowedUsesPtrOutput)
}

func (o AllowedUsesOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AllowedUsesOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AllowedUses) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AllowedUsesOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AllowedUsesOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AllowedUses) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AllowedUsesPtrOutput struct{ *pulumi.OutputState }

func (AllowedUsesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AllowedUses)(nil)).Elem()
}

func (o AllowedUsesPtrOutput) ToAllowedUsesPtrOutput() AllowedUsesPtrOutput {
	return o
}

func (o AllowedUsesPtrOutput) ToAllowedUsesPtrOutputWithContext(ctx context.Context) AllowedUsesPtrOutput {
	return o
}

func (o AllowedUsesPtrOutput) Elem() AllowedUsesOutput {
	return o.ApplyT(func(v *AllowedUses) AllowedUses {
		if v != nil {
			return *v
		}
		var ret AllowedUses
		return ret
	}).(AllowedUsesOutput)
}

func (o AllowedUsesPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AllowedUsesPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AllowedUses) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// AllowedUsesInput is an input type that accepts AllowedUsesArgs and AllowedUsesOutput values.
// You can construct a concrete instance of `AllowedUsesInput` via:
//
//          AllowedUsesArgs{...}
type AllowedUsesInput interface {
	pulumi.Input

	ToAllowedUsesOutput() AllowedUsesOutput
	ToAllowedUsesOutputWithContext(context.Context) AllowedUsesOutput
}

var allowedUsesPtrType = reflect.TypeOf((**AllowedUses)(nil)).Elem()

type AllowedUsesPtrInput interface {
	pulumi.Input

	ToAllowedUsesPtrOutput() AllowedUsesPtrOutput
	ToAllowedUsesPtrOutputWithContext(context.Context) AllowedUsesPtrOutput
}

type allowedUsesPtr string

func AllowedUsesPtr(v string) AllowedUsesPtrInput {
	return (*allowedUsesPtr)(&v)
}

func (*allowedUsesPtr) ElementType() reflect.Type {
	return allowedUsesPtrType
}

func (in *allowedUsesPtr) ToAllowedUsesPtrOutput() AllowedUsesPtrOutput {
	return pulumi.ToOutput(in).(AllowedUsesPtrOutput)
}

func (in *allowedUsesPtr) ToAllowedUsesPtrOutputWithContext(ctx context.Context) AllowedUsesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AllowedUsesPtrOutput)
}

// AllowedUsesArrayInput is an input type that accepts AllowedUsesArray and AllowedUsesArrayOutput values.
// You can construct a concrete instance of `AllowedUsesArrayInput` via:
//
//          AllowedUsesArray{ AllowedUsesArgs{...} }
type AllowedUsesArrayInput interface {
	pulumi.Input

	ToAllowedUsesArrayOutput() AllowedUsesArrayOutput
	ToAllowedUsesArrayOutputWithContext(context.Context) AllowedUsesArrayOutput
}

type AllowedUsesArray []AllowedUses

func (AllowedUsesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AllowedUses)(nil)).Elem()
}

func (i AllowedUsesArray) ToAllowedUsesArrayOutput() AllowedUsesArrayOutput {
	return i.ToAllowedUsesArrayOutputWithContext(context.Background())
}

func (i AllowedUsesArray) ToAllowedUsesArrayOutputWithContext(ctx context.Context) AllowedUsesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AllowedUsesArrayOutput)
}

type AllowedUsesArrayOutput struct{ *pulumi.OutputState }

func (AllowedUsesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AllowedUses)(nil)).Elem()
}

func (o AllowedUsesArrayOutput) ToAllowedUsesArrayOutput() AllowedUsesArrayOutput {
	return o
}

func (o AllowedUsesArrayOutput) ToAllowedUsesArrayOutputWithContext(ctx context.Context) AllowedUsesArrayOutput {
	return o
}

func (o AllowedUsesArrayOutput) Index(i pulumi.IntInput) AllowedUsesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AllowedUses {
		return vs[0].([]AllowedUses)[vs[1].(int)]
	}).(AllowedUsesOutput)
}

type EcdsaCurve string

const (
	EcdsaCurveP224 = EcdsaCurve("P224")
	EcdsaCurveP256 = EcdsaCurve("P256")
	EcdsaCurveP384 = EcdsaCurve("P384")
	EcdsaCurveP521 = EcdsaCurve("P521")
)

func (EcdsaCurve) ElementType() reflect.Type {
	return reflect.TypeOf((*EcdsaCurve)(nil)).Elem()
}

func (e EcdsaCurve) ToEcdsaCurveOutput() EcdsaCurveOutput {
	return pulumi.ToOutput(e).(EcdsaCurveOutput)
}

func (e EcdsaCurve) ToEcdsaCurveOutputWithContext(ctx context.Context) EcdsaCurveOutput {
	return pulumi.ToOutputWithContext(ctx, e).(EcdsaCurveOutput)
}

func (e EcdsaCurve) ToEcdsaCurvePtrOutput() EcdsaCurvePtrOutput {
	return e.ToEcdsaCurvePtrOutputWithContext(context.Background())
}

func (e EcdsaCurve) ToEcdsaCurvePtrOutputWithContext(ctx context.Context) EcdsaCurvePtrOutput {
	return EcdsaCurve(e).ToEcdsaCurveOutputWithContext(ctx).ToEcdsaCurvePtrOutputWithContext(ctx)
}

func (e EcdsaCurve) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EcdsaCurve) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EcdsaCurve) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EcdsaCurve) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type EcdsaCurveOutput struct{ *pulumi.OutputState }

func (EcdsaCurveOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EcdsaCurve)(nil)).Elem()
}

func (o EcdsaCurveOutput) ToEcdsaCurveOutput() EcdsaCurveOutput {
	return o
}

func (o EcdsaCurveOutput) ToEcdsaCurveOutputWithContext(ctx context.Context) EcdsaCurveOutput {
	return o
}

func (o EcdsaCurveOutput) ToEcdsaCurvePtrOutput() EcdsaCurvePtrOutput {
	return o.ToEcdsaCurvePtrOutputWithContext(context.Background())
}

func (o EcdsaCurveOutput) ToEcdsaCurvePtrOutputWithContext(ctx context.Context) EcdsaCurvePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EcdsaCurve) *EcdsaCurve {
		return &v
	}).(EcdsaCurvePtrOutput)
}

func (o EcdsaCurveOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EcdsaCurveOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EcdsaCurve) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EcdsaCurveOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EcdsaCurveOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EcdsaCurve) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EcdsaCurvePtrOutput struct{ *pulumi.OutputState }

func (EcdsaCurvePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EcdsaCurve)(nil)).Elem()
}

func (o EcdsaCurvePtrOutput) ToEcdsaCurvePtrOutput() EcdsaCurvePtrOutput {
	return o
}

func (o EcdsaCurvePtrOutput) ToEcdsaCurvePtrOutputWithContext(ctx context.Context) EcdsaCurvePtrOutput {
	return o
}

func (o EcdsaCurvePtrOutput) Elem() EcdsaCurveOutput {
	return o.ApplyT(func(v *EcdsaCurve) EcdsaCurve {
		if v != nil {
			return *v
		}
		var ret EcdsaCurve
		return ret
	}).(EcdsaCurveOutput)
}

func (o EcdsaCurvePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EcdsaCurvePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *EcdsaCurve) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// EcdsaCurveInput is an input type that accepts EcdsaCurveArgs and EcdsaCurveOutput values.
// You can construct a concrete instance of `EcdsaCurveInput` via:
//
//          EcdsaCurveArgs{...}
type EcdsaCurveInput interface {
	pulumi.Input

	ToEcdsaCurveOutput() EcdsaCurveOutput
	ToEcdsaCurveOutputWithContext(context.Context) EcdsaCurveOutput
}

var ecdsaCurvePtrType = reflect.TypeOf((**EcdsaCurve)(nil)).Elem()

type EcdsaCurvePtrInput interface {
	pulumi.Input

	ToEcdsaCurvePtrOutput() EcdsaCurvePtrOutput
	ToEcdsaCurvePtrOutputWithContext(context.Context) EcdsaCurvePtrOutput
}

type ecdsaCurvePtr string

func EcdsaCurvePtr(v string) EcdsaCurvePtrInput {
	return (*ecdsaCurvePtr)(&v)
}

func (*ecdsaCurvePtr) ElementType() reflect.Type {
	return ecdsaCurvePtrType
}

func (in *ecdsaCurvePtr) ToEcdsaCurvePtrOutput() EcdsaCurvePtrOutput {
	return pulumi.ToOutput(in).(EcdsaCurvePtrOutput)
}

func (in *ecdsaCurvePtr) ToEcdsaCurvePtrOutputWithContext(ctx context.Context) EcdsaCurvePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(EcdsaCurvePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AlgorithmInput)(nil)).Elem(), Algorithm("RSA"))
	pulumi.RegisterInputType(reflect.TypeOf((*AlgorithmPtrInput)(nil)).Elem(), Algorithm("RSA"))
	pulumi.RegisterInputType(reflect.TypeOf((*AllowedUsesInput)(nil)).Elem(), AllowedUses("any_extended"))
	pulumi.RegisterInputType(reflect.TypeOf((*AllowedUsesPtrInput)(nil)).Elem(), AllowedUses("any_extended"))
	pulumi.RegisterInputType(reflect.TypeOf((*AllowedUsesArrayInput)(nil)).Elem(), AllowedUsesArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EcdsaCurveInput)(nil)).Elem(), EcdsaCurve("P224"))
	pulumi.RegisterInputType(reflect.TypeOf((*EcdsaCurvePtrInput)(nil)).Elem(), EcdsaCurve("P224"))
	pulumi.RegisterOutputType(AlgorithmOutput{})
	pulumi.RegisterOutputType(AlgorithmPtrOutput{})
	pulumi.RegisterOutputType(AllowedUsesOutput{})
	pulumi.RegisterOutputType(AllowedUsesPtrOutput{})
	pulumi.RegisterOutputType(AllowedUsesArrayOutput{})
	pulumi.RegisterOutputType(EcdsaCurveOutput{})
	pulumi.RegisterOutputType(EcdsaCurvePtrOutput{})
}