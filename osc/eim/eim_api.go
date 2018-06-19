package eim

import (
	"context"
	"net/http"

	"github.com/terraform-providers/terraform-provider-outscale/osc"
)

//Operations defines all the operations needed for EIM
type Operations struct {
	client *osc.Client
}

//Service all the necessary actions for them EIM service
type Service interface {
	CreatePolicy(input *CreatePolicyInput) (*CreatePolicyOutput, error)
	GetPolicy(input *GetPolicyInput) (*GetPolicyOutput, error)
	GetPolicyVersion(input *GetPolicyVersionInput) (*GetPolicyVersionOutput, error)
	DeletePolicy(input *DeletePolicyInput) (*DeletePolicyOutput, error)
	DeletePolicyVersion(input *DeletePolicyVersionInput) (*DeletePolicyVersionOutput, error)
	ListPolicyVersions(input *ListPolicyVersionsInput) (*ListPolicyVersionsOutput, error)
	CreateGroup(input *CreateGroupInput) (*CreateGroupOutput, error)
	GetGroup(input *GetGroupInput) (*GetGroupOutput, error)
	UpdateGroup(input *UpdateGroupInput) (*UpdateGroupOutput, error)
	DeleteGroup(input *DeleteGroupInput) (*DeleteGroupOutput, error)
	AddUserToGroup(input *AddUserToGroupInput) (*AddUserToGroupOutput, error)
	RemoveUserFromGroup(input *RemoveUserFromGroupInput) (*RemoveUserFromGroupOutput, error)
	CreateUser(input *CreateUserInput) (*CreateUserOutput, error)
	GetUser(input *GetUserInput) (*GetUserOutput, error)
	UpdateUser(input *UpdateUserInput) (*UpdateUserOutput, error)
	ListGroupsForUserPages(input *ListGroupsForUserInput) (*ListGroupsForUserOutput, error)
	DeleteUser(input *DeleteUserInput) (*DeleteUserOutput, error)
	SetDefaultPolicyVersion(input *SetDefaultPolicyVersionInput) (*SetDefaultPolicyVersionOutput, error)
	AttachUserPolicy(input *AttachUserPolicyInput) (*AttachUserPolicyOutput, error)
	ListAttachedUserPolicies(input *ListAttachedUserPoliciesInput) (*ListAttachedUserPoliciesOutput, error)
	DetachUserPolicy(input *DetachUserPolicyInput) (*DetachUserPolicyOutput, error)
	GetUserPolicy(input *GetUserPolicyInput) (*GetUserPolicyOutput, error)
	ListUsers(input *ListUsersInput) (*ListUsersOutput, error)
	ListGroups(input *ListGroupsInput) (*ListGroupsOutput, error)
	ListGroupsForUser(input *ListGroupsForUserInput) (*ListGroupsForUserOutput, error)
	UploadServerCertificate(input *UploadServerCertificateInput) (*UploadServerCertificateOutput, error)
	GetServerCertificate(input *GetServerCertificateInput) (*GetServerCertificateOutput, error)
	DeleteServerCertificate(input *DeleteServerCertificateInput) (*DeleteServerCertificateOutput, error)
	ListServerCertificates(input *ListServerCertificatesInput) (*ListServerCertificatesOutput, error)
	UpdateServerCertificate(input *UpdateServerCertificateInput) (*UpdateServerCertificateOutput, error)
}

// CreatePolicy ...
func (v Operations) CreatePolicy(input *CreatePolicyInput) (*CreatePolicyOutput, error) {
	inURL := "/"
	endpoint := "CreatePolicy"
	output := &CreatePolicyOutput{}

	if input == nil {
		input = &CreatePolicyInput{}
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

// GetPolicy ...
func (v Operations) GetPolicy(input *GetPolicyInput) (*GetPolicyOutput, error) {
	inURL := "/"
	endpoint := "GetPolicy"
	output := &GetPolicyOutput{}

	if input == nil {
		input = &GetPolicyInput{}
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

// GetPolicyVersion ...
func (v Operations) GetPolicyVersion(input *GetPolicyVersionInput) (*GetPolicyVersionOutput, error) {
	inURL := "/"
	endpoint := "GetPolicyVersion"
	output := &GetPolicyVersionOutput{}

	if input == nil {
		input = &GetPolicyVersionInput{}
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

// DeletePolicy ...
func (v Operations) DeletePolicy(input *DeletePolicyInput) (*DeletePolicyOutput, error) {
	inURL := "/"
	endpoint := "DeletePolicy"
	output := &DeletePolicyOutput{}

	if input == nil {
		input = &DeletePolicyInput{}
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

// DeletePolicyVersion ...
func (v Operations) DeletePolicyVersion(input *DeletePolicyVersionInput) (*DeletePolicyVersionOutput, error) {
	inURL := "/"
	endpoint := "DeletePolicyVersion"
	output := &DeletePolicyVersionOutput{}

	if input == nil {
		input = &DeletePolicyVersionInput{}
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

// ListPolicyVersions ...
func (v Operations) ListPolicyVersions(input *ListPolicyVersionsInput) (*ListPolicyVersionsOutput, error) {
	inURL := "/"
	endpoint := "ListPolicyVersions"
	output := &ListPolicyVersionsOutput{}

	if input == nil {
		input = &ListPolicyVersionsInput{}
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

// CreateGroup ...
func (v Operations) CreateGroup(input *CreateGroupInput) (*CreateGroupOutput, error) {
	inURL := "/"
	endpoint := "CreateGroup"
	output := &CreateGroupOutput{}

	if input == nil {
		input = &CreateGroupInput{}
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

// UploadServerCertificate Uploads a server certificate and its matching private key.
func (v Operations) UploadServerCertificate(input *UploadServerCertificateInput) (*UploadServerCertificateOutput, error) {
	inURL := "/"
	endpoint := "UploadServerCertificate"
	output := &UploadServerCertificateOutput{}

	if input == nil {
		input = &UploadServerCertificateInput{}
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

// GetGroup ...
func (v Operations) GetGroup(input *GetGroupInput) (*GetGroupOutput, error) {
	inURL := "/"
	endpoint := "GetGroup"
	output := &GetGroupOutput{}

	if input == nil {
		input = &GetGroupInput{}
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

// GetServerCertificate Gets a server certificate and its matching private key.
func (v Operations) GetServerCertificate(input *GetServerCertificateInput) (*GetServerCertificateOutput, error) {
	inURL := "/"
	endpoint := "GetServerCertificate"
	output := &GetServerCertificateOutput{}

	if input == nil {
		input = &GetServerCertificateInput{}
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

// UpdateGroup ...
func (v Operations) UpdateGroup(input *UpdateGroupInput) (*UpdateGroupOutput, error) {
	inURL := "/"
	endpoint := "UpdateGroup"
	output := &UpdateGroupOutput{}

	if input == nil {
		input = &UpdateGroupInput{}
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

// DeleteServerCertificate Deletes a server certificate and its matching private key.
func (v Operations) DeleteServerCertificate(input *DeleteServerCertificateInput) (*DeleteServerCertificateOutput, error) {
	inURL := "/"
	endpoint := "DeleteServerCertificate"
	output := &DeleteServerCertificateOutput{}

	if input == nil {
		input = &DeleteServerCertificateInput{}
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

// DeleteGroup ...
func (v Operations) DeleteGroup(input *DeleteGroupInput) (*DeleteGroupOutput, error) {
	inURL := "/"
	endpoint := "DeleteGroup"
	output := &DeleteGroupOutput{}

	if input == nil {
		input = &DeleteGroupInput{}
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

// ListServerCertificates ...
func (v Operations) ListServerCertificates(input *ListServerCertificatesInput) (*ListServerCertificatesOutput, error) {
	inURL := "/"
	endpoint := "ListServerCertificates"
	output := &ListServerCertificatesOutput{}

	if input == nil {
		input = &ListServerCertificatesInput{}
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

// AddUserToGroup ...
func (v Operations) AddUserToGroup(input *AddUserToGroupInput) (*AddUserToGroupOutput, error) {
	inURL := "/"
	endpoint := "AddUserToGroup"
	output := &AddUserToGroupOutput{}

	if input == nil {
		input = &AddUserToGroupInput{}
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

// RemoveUserFromGroup ...
func (v Operations) RemoveUserFromGroup(input *RemoveUserFromGroupInput) (*RemoveUserFromGroupOutput, error) {
	inURL := "/"
	endpoint := "RemoveUserFromGroup"
	output := &RemoveUserFromGroupOutput{}

	if input == nil {
		input = &RemoveUserFromGroupInput{}
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

// CreateUser ...
func (v Operations) CreateUser(input *CreateUserInput) (*CreateUserOutput, error) {
	inURL := "/"
	endpoint := "CreateUser"
	output := &CreateUserOutput{}

	if input == nil {
		input = &CreateUserInput{}
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

// GetUser ...
func (v Operations) GetUser(input *GetUserInput) (*GetUserOutput, error) {
	inURL := "/"
	endpoint := "GetUser"
	output := &GetUserOutput{}

	if input == nil {
		input = &GetUserInput{}
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

// UpdateUser ...
func (v Operations) UpdateUser(input *UpdateUserInput) (*UpdateUserOutput, error) {
	inURL := "/"
	endpoint := "UpdateUser"
	output := &UpdateUserOutput{}

	if input == nil {
		input = &UpdateUserInput{}
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

// ListGroupsForUserPages ...
func (v Operations) ListGroupsForUserPages(input *ListGroupsForUserInput) (*ListGroupsForUserOutput, error) {
	inURL := "/"
	endpoint := "ListGroupsForUserPages"
	output := &ListGroupsForUserOutput{}

	if input == nil {
		input = &ListGroupsForUserInput{}
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

// DeleteUser ...
func (v Operations) DeleteUser(input *DeleteUserInput) (*DeleteUserOutput, error) {
	inURL := "/"
	endpoint := "DeleteUser"
	output := &DeleteUserOutput{}

	if input == nil {
		input = &DeleteUserInput{}
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

// SetDefaultPolicyVersion ...
func (v Operations) SetDefaultPolicyVersion(input *SetDefaultPolicyVersionInput) (*SetDefaultPolicyVersionOutput, error) {
	inURL := "/"
	endpoint := "SetDefaultPolicyVersion"
	output := &SetDefaultPolicyVersionOutput{}

	if input == nil {
		input = &SetDefaultPolicyVersionInput{}
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

// AttachUserPolicy ...
func (v Operations) AttachUserPolicy(input *AttachUserPolicyInput) (*AttachUserPolicyOutput, error) {
	inURL := "/"
	endpoint := "AttachUserPolicy"
	output := &AttachUserPolicyOutput{}

	if input == nil {
		input = &AttachUserPolicyInput{}
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

// ListAttachedUserPolicies ...
func (v Operations) ListAttachedUserPolicies(input *ListAttachedUserPoliciesInput) (*ListAttachedUserPoliciesOutput, error) {
	inURL := "/"
	endpoint := "ListAttachedUserPolicies"
	output := &ListAttachedUserPoliciesOutput{}

	if input == nil {
		input = &ListAttachedUserPoliciesInput{}
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

// DetachUserPolicy ...
func (v Operations) DetachUserPolicy(input *DetachUserPolicyInput) (*DetachUserPolicyOutput, error) {
	inURL := "/"
	endpoint := "DetachUserPolicy"
	output := &DetachUserPolicyOutput{}

	if input == nil {
		input = &DetachUserPolicyInput{}
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

// GetUserPolicy ...
func (v Operations) GetUserPolicy(input *GetUserPolicyInput) (*GetUserPolicyOutput, error) {
	inURL := "/"
	endpoint := "GetUserPolicy"
	output := &GetUserPolicyOutput{}

	if input == nil {
		input = &GetUserPolicyInput{}
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

// ListUsers ...
func (v Operations) ListUsers(input *ListUsersInput) (*ListUsersOutput, error) {
	inURL := "/"
	endpoint := "ListUsers"
	output := &ListUsersOutput{}

	if input == nil {
		input = &ListUsersInput{}
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

// ListGroups ...
func (v Operations) ListGroups(input *ListGroupsInput) (*ListGroupsOutput, error) {
	inURL := "/"
	endpoint := "ListGroups"
	output := &ListGroupsOutput{}

	if input == nil {
		input = &ListGroupsInput{}
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

// ListGroupsForUser ...
func (v Operations) ListGroupsForUser(input *ListGroupsForUserInput) (*ListGroupsForUserOutput, error) {
	inURL := "/"
	endpoint := "ListGroupsForUser"
	output := &ListGroupsForUserOutput{}

	if input == nil {
		input = &ListGroupsForUserInput{}
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

// UpdateServerCertificate ...
func (v Operations) UpdateServerCertificate(input *UpdateServerCertificateInput) (*UpdateServerCertificateOutput, error) {
	inURL := "/"
	endpoint := "UpdateServerCertificate"
	output := &UpdateServerCertificateOutput{}

	if input == nil {
		input = &UpdateServerCertificateInput{}
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
