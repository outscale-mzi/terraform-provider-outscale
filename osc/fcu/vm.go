package fcu

import (
	"context"
	"net/http"

	"github.com/terraform-providers/terraform-provider-outscale/osc"
)

//VMOperations defines all the operations needed for FCU VMs
type VMOperations struct {
	client *osc.Client
}

//VMService all the necessary actions for them VM service
type VMService interface {
	RunInstance(input *RunInstancesInput) (*Reservation, error)
	DescribeInstances(input *DescribeInstancesInput) (*DescribeInstancesOutput, error)
	GetPasswordData(input *GetPasswordDataInput) (*GetPasswordDataOutput, error)
	ModifyInstanceKeyPair(input *ModifyInstanceKeyPairInput) error
	ModifyInstanceAttribute(input *ModifyInstanceAttributeInput) (*ModifyInstanceAttributeOutput, error)
	TerminateInstances(input *TerminateInstancesInput) (*TerminateInstancesOutput, error)
	AllocateAddress(input *AllocateAddressInput) (*AllocateAddressOutput, error)
	DescribeAddressesRequest(input *DescribeAddressesInput) (*DescribeAddressesOutput, error)
	StopInstances(input *StopInstancesInput) (*StopInstancesOutput, error)
	StartInstances(input *StartInstancesInput) (*StartInstancesOutput, error)
	AssociateAddress(input *AssociateAddressInput) (*AssociateAddressOutput, error)
	DisassociateAddress(input *DisassociateAddressInput) (*DisassociateAddressOutput, error)
	ReleaseAddress(input *ReleaseAddressInput) (*ReleaseAddressOutput, error)
}

const opRunInstances = "RunInstances"

func (v VMOperations) RunInstance(input *RunInstancesInput) (*Reservation, error) {
	req, err := v.client.NewRequest(context.Background(), opRunInstances, http.MethodGet, "/", input)
	if err != nil {
		return nil, err
	}

	output := Reservation{}

	err = v.client.Do(context.Background(), req, &output)
	if err != nil {
		return nil, err
	}

	return &output, nil
}

const opDescribeInstances = "DescribeInstances"

// DescribeInstances method
func (v VMOperations) DescribeInstances(input *DescribeInstancesInput) (*DescribeInstancesOutput, error) {
	inURL := "/"
	endpoint := "DescribeInstances"
	output := &DescribeInstancesOutput{}

	if input == nil {
		input = &DescribeInstancesInput{}
	}

	req, err := v.client.NewRequest(context.TODO(), endpoint, http.MethodGet, inURL, input)

	if err != nil {
		return nil, err
	}

	err = v.client.Do(context.TODO(), req, output)
	if err != nil {
		return nil, err
	}

	return output, nil
}

// DescribeInstances method
func (v VMOperations) ModifyInstanceKeyPair(input *ModifyInstanceKeyPairInput) error {
	inURL := "/?Action=ModifyInstanceKeypair"
	endpoint := "ModifyInstanceKeypair"

	if input == nil {
		input = &ModifyInstanceKeyPairInput{}
	}

	req, err := v.client.NewRequest(context.TODO(), endpoint, http.MethodPost, inURL, input)

	if err != nil {
		return err
	}

	err = v.client.Do(context.TODO(), req, nil)
	if err != nil {
		return err
	}

	return nil
}

func (v VMOperations) ModifyInstanceAttribute(input *ModifyInstanceAttributeInput) (*ModifyInstanceAttributeOutput, error) {
	inURL := "/"
	endpoint := "ModifyInstanceAttribute"
	output := &ModifyInstanceAttributeOutput{}

	if input == nil {
		input = &ModifyInstanceAttributeInput{}
	}

	req, err := v.client.NewRequest(context.TODO(), endpoint, http.MethodGet, inURL, input)

	if err != nil {
		return nil, err
	}

	err = v.client.Do(context.TODO(), req, output)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (v VMOperations) GetPasswordData(input *GetPasswordDataInput) (*GetPasswordDataOutput, error) {
	inURL := "/"
	endpoint := "GetPasswordData"
	output := &GetPasswordDataOutput{}

	if input == nil {
		input = &GetPasswordDataInput{}
	}

	req, err := v.client.NewRequest(context.TODO(), endpoint, http.MethodGet, inURL, input)

	if err != nil {
		return nil, err
	}

	err = v.client.Do(context.TODO(), req, output)
	if err != nil {
		return nil, err
	}

	return output, nil
}

// DescribeInstances method
func (v VMOperations) TerminateInstances(input *TerminateInstancesInput) (*TerminateInstancesOutput, error) {
	inURL := "/"
	endpoint := "TerminateInstances"
	output := &TerminateInstancesOutput{}

	if input == nil {
		input = &TerminateInstancesInput{}
	}

	req, err := v.client.NewRequest(context.TODO(), endpoint, http.MethodGet, inURL, input)

	if err != nil {
		return nil, err
	}

	err = v.client.Do(context.TODO(), req, output)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (v VMOperations) AllocateAddress(input *AllocateAddressInput) (*AllocateAddressOutput, error) {
	inURL := "/"
	endpoint := "AllocateAddress"
	output := &AllocateAddressOutput{}

	if input == nil {
		input = &AllocateAddressInput{}
	}
	req, err := v.client.NewRequest(context.TODO(), endpoint, http.MethodGet, inURL, input)

	if err != nil {
		return nil, err
	}

	err = v.client.Do(context.TODO(), req, output)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (v VMOperations) StopInstances(input *StopInstancesInput) (*StopInstancesOutput, error) {
	inURL := "/"
	endpoint := "StopInstances"
	output := &StopInstancesOutput{}

	if input == nil {
		input = &StopInstancesInput{}
	}

	req, err := v.client.NewRequest(context.TODO(), endpoint, http.MethodGet, inURL, input)

	if err != nil {
		return nil, err
	}

	err = v.client.Do(context.TODO(), req, output)
	if err != nil {
		return nil, err
	}

	return output, nil
}

//DescribeAddresses
func (v VMOperations) DescribeAddressesRequest(input *DescribeAddressesInput) (*DescribeAddressesOutput, error) {
	inURL := "/"
	endpoint := "DescribeAddresses"
	output := &DescribeAddressesOutput{}

	if input == nil {
		input = &DescribeAddressesInput{}
	}

	req, err := v.client.NewRequest(context.TODO(), endpoint, http.MethodGet, inURL, input)

	if err != nil {
		return nil, err
	}

	err = v.client.Do(context.TODO(), req, output)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (v VMOperations) StartInstances(input *StartInstancesInput) (*StartInstancesOutput, error) {
	inURL := "/"
	endpoint := "StartInstances"
	output := &StartInstancesOutput{}

	if input == nil {
		input = &StartInstancesInput{}
	}

	req, err := v.client.NewRequest(context.TODO(), endpoint, http.MethodGet, inURL, input)

	if err != nil {
		return nil, err
	}

	err = v.client.Do(context.TODO(), req, output)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (v VMOperations) AssociateAddress(input *AssociateAddressInput) (*AssociateAddressOutput, error) {
	inURL := "/"
	endpoint := "AssociateAddress"
	output := &AssociateAddressOutput{}

	if input == nil {
		input = &AssociateAddressInput{}
	}
	req, err := v.client.NewRequest(context.TODO(), endpoint, http.MethodGet, inURL, input)

	if err != nil {
		return nil, err
	}

	err = v.client.Do(context.TODO(), req, output)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (v VMOperations) DisassociateAddress(input *DisassociateAddressInput) (*DisassociateAddressOutput, error) {
	inURL := "/"
	endpoint := "DisassociateAddress"
	output := &DisassociateAddressOutput{}

	if input == nil {
		input = &DisassociateAddressInput{}
	}
	req, err := v.client.NewRequest(context.TODO(), endpoint, http.MethodGet, inURL, input)

	if err != nil {
		return nil, err
	}

	err = v.client.Do(context.TODO(), req, output)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (v VMOperations) ReleaseAddress(input *ReleaseAddressInput) (*ReleaseAddressOutput, error) {
	inURL := "/"
	endpoint := "ReleaseAddress"
	output := &ReleaseAddressOutput{}

	if input == nil {
		input = &ReleaseAddressInput{}
	}
	req, err := v.client.NewRequest(context.TODO(), endpoint, http.MethodGet, inURL, input)

	if err != nil {
		return nil, err
	}

	err = v.client.Do(context.TODO(), req, output)
	if err != nil {
		return nil, err
	}

	return output, nil
}
