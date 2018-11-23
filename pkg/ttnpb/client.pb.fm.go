// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

import (
	fmt "fmt"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	time "time"
)

var ClientFieldPathsNested = []string{
	"attributes",
	"contact_info",
	"created_at",
	"description",
	"endorsed",
	"grants",
	"ids",
	"ids.client_id",
	"name",
	"redirect_uris",
	"rights",
	"secret",
	"skip_authorization",
	"state",
	"updated_at",
}

var ClientFieldPathsTopLevel = []string{
	"attributes",
	"contact_info",
	"created_at",
	"description",
	"endorsed",
	"grants",
	"ids",
	"name",
	"redirect_uris",
	"rights",
	"secret",
	"skip_authorization",
	"state",
	"updated_at",
}

func (dst *Client) SetFields(src *Client, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "ids":
			if len(subs) > 0 {
				newDst := &dst.ClientIdentifiers
				var newSrc *ClientIdentifiers
				if src != nil {
					newSrc = &src.ClientIdentifiers
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.ClientIdentifiers = src.ClientIdentifiers
				} else {
					var zero ClientIdentifiers
					dst.ClientIdentifiers = zero
				}
			}
		case "created_at":
			if len(subs) > 0 {
				return fmt.Errorf("'created_at' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.CreatedAt = src.CreatedAt
			} else {
				var zero time.Time
				dst.CreatedAt = zero
			}
		case "updated_at":
			if len(subs) > 0 {
				return fmt.Errorf("'updated_at' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.UpdatedAt = src.UpdatedAt
			} else {
				var zero time.Time
				dst.UpdatedAt = zero
			}
		case "name":
			if len(subs) > 0 {
				return fmt.Errorf("'name' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Name = src.Name
			} else {
				var zero string
				dst.Name = zero
			}
		case "description":
			if len(subs) > 0 {
				return fmt.Errorf("'description' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Description = src.Description
			} else {
				var zero string
				dst.Description = zero
			}
		case "attributes":
			if len(subs) > 0 {
				return fmt.Errorf("'attributes' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Attributes = src.Attributes
			} else {
				dst.Attributes = nil
			}
		case "contact_info":
			if len(subs) > 0 {
				return fmt.Errorf("'contact_info' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ContactInfo = src.ContactInfo
			} else {
				dst.ContactInfo = nil
			}
		case "secret":
			if len(subs) > 0 {
				return fmt.Errorf("'secret' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Secret = src.Secret
			} else {
				var zero string
				dst.Secret = zero
			}
		case "redirect_uris":
			if len(subs) > 0 {
				return fmt.Errorf("'redirect_uris' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.RedirectURIs = src.RedirectURIs
			} else {
				dst.RedirectURIs = nil
			}
		case "state":
			if len(subs) > 0 {
				return fmt.Errorf("'state' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.State = src.State
			} else {
				var zero State
				dst.State = zero
			}
		case "skip_authorization":
			if len(subs) > 0 {
				return fmt.Errorf("'skip_authorization' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.SkipAuthorization = src.SkipAuthorization
			} else {
				var zero bool
				dst.SkipAuthorization = zero
			}
		case "endorsed":
			if len(subs) > 0 {
				return fmt.Errorf("'endorsed' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Endorsed = src.Endorsed
			} else {
				var zero bool
				dst.Endorsed = zero
			}
		case "grants":
			if len(subs) > 0 {
				return fmt.Errorf("'grants' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Grants = src.Grants
			} else {
				dst.Grants = nil
			}
		case "rights":
			if len(subs) > 0 {
				return fmt.Errorf("'rights' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Rights = src.Rights
			} else {
				dst.Rights = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

var ClientsFieldPathsNested = []string{
	"clients",
}

var ClientsFieldPathsTopLevel = []string{
	"clients",
}

func (dst *Clients) SetFields(src *Clients, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "clients":
			if len(subs) > 0 {
				return fmt.Errorf("'clients' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Clients = src.Clients
			} else {
				dst.Clients = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

var GetClientRequestFieldPathsNested = []string{
	"client_ids",
	"client_ids.client_id",
	"field_mask",
}

var GetClientRequestFieldPathsTopLevel = []string{
	"client_ids",
	"field_mask",
}

func (dst *GetClientRequest) SetFields(src *GetClientRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "client_ids":
			if len(subs) > 0 {
				newDst := &dst.ClientIdentifiers
				var newSrc *ClientIdentifiers
				if src != nil {
					newSrc = &src.ClientIdentifiers
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.ClientIdentifiers = src.ClientIdentifiers
				} else {
					var zero ClientIdentifiers
					dst.ClientIdentifiers = zero
				}
			}
		case "field_mask":
			if len(subs) > 0 {
				return fmt.Errorf("'field_mask' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.FieldMask = src.FieldMask
			} else {
				var zero github_com_gogo_protobuf_types.FieldMask
				dst.FieldMask = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

var ListClientsRequestFieldPathsNested = []string{
	"collaborator",
	"collaborator.ids.organization_ids",
	"collaborator.ids.organization_ids.organization_id",
	"collaborator.ids.user_ids",
	"collaborator.ids.user_ids.email",
	"collaborator.ids.user_ids.user_id",
	"field_mask",
	"ids",
	"limit",
	"order",
	"page",
}

var ListClientsRequestFieldPathsTopLevel = []string{
	"collaborator",
	"field_mask",
	"ids",
	"limit",
	"order",
	"page",
}

func (dst *ListClientsRequest) SetFields(src *ListClientsRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "collaborator":
			if len(subs) > 0 {
				newDst := dst.Collaborator
				if newDst == nil {
					newDst = &OrganizationOrUserIdentifiers{}
					dst.Collaborator = newDst
				}
				var newSrc *OrganizationOrUserIdentifiers
				if src != nil {
					newSrc = src.Collaborator
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Collaborator = src.Collaborator
				} else {
					dst.Collaborator = nil
				}
			}
		case "field_mask":
			if len(subs) > 0 {
				return fmt.Errorf("'field_mask' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.FieldMask = src.FieldMask
			} else {
				var zero github_com_gogo_protobuf_types.FieldMask
				dst.FieldMask = zero
			}
		case "order":
			if len(subs) > 0 {
				return fmt.Errorf("'order' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Order = src.Order
			} else {
				var zero string
				dst.Order = zero
			}
		case "limit":
			if len(subs) > 0 {
				return fmt.Errorf("'limit' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Limit = src.Limit
			} else {
				var zero uint32
				dst.Limit = zero
			}
		case "page":
			if len(subs) > 0 {
				return fmt.Errorf("'page' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Page = src.Page
			} else {
				var zero uint32
				dst.Page = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

var CreateClientRequestFieldPathsNested = []string{
	"client",
	"client.attributes",
	"client.contact_info",
	"client.created_at",
	"client.description",
	"client.endorsed",
	"client.grants",
	"client.ids",
	"client.ids.client_id",
	"client.name",
	"client.redirect_uris",
	"client.rights",
	"client.secret",
	"client.skip_authorization",
	"client.state",
	"client.updated_at",
	"collaborator",
	"collaborator.ids.organization_ids",
	"collaborator.ids.organization_ids.organization_id",
	"collaborator.ids.user_ids",
	"collaborator.ids.user_ids.email",
	"collaborator.ids.user_ids.user_id",
	"ids",
}

var CreateClientRequestFieldPathsTopLevel = []string{
	"client",
	"collaborator",
	"ids",
}

func (dst *CreateClientRequest) SetFields(src *CreateClientRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "client":
			if len(subs) > 0 {
				newDst := &dst.Client
				var newSrc *Client
				if src != nil {
					newSrc = &src.Client
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Client = src.Client
				} else {
					var zero Client
					dst.Client = zero
				}
			}
		case "collaborator":
			if len(subs) > 0 {
				newDst := &dst.Collaborator
				var newSrc *OrganizationOrUserIdentifiers
				if src != nil {
					newSrc = &src.Collaborator
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Collaborator = src.Collaborator
				} else {
					var zero OrganizationOrUserIdentifiers
					dst.Collaborator = zero
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

var UpdateClientRequestFieldPathsNested = []string{
	"client",
	"client.attributes",
	"client.contact_info",
	"client.created_at",
	"client.description",
	"client.endorsed",
	"client.grants",
	"client.ids",
	"client.ids.client_id",
	"client.name",
	"client.redirect_uris",
	"client.rights",
	"client.secret",
	"client.skip_authorization",
	"client.state",
	"client.updated_at",
	"field_mask",
}

var UpdateClientRequestFieldPathsTopLevel = []string{
	"client",
	"field_mask",
}

func (dst *UpdateClientRequest) SetFields(src *UpdateClientRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "client":
			if len(subs) > 0 {
				newDst := &dst.Client
				var newSrc *Client
				if src != nil {
					newSrc = &src.Client
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Client = src.Client
				} else {
					var zero Client
					dst.Client = zero
				}
			}
		case "field_mask":
			if len(subs) > 0 {
				return fmt.Errorf("'field_mask' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.FieldMask = src.FieldMask
			} else {
				var zero github_com_gogo_protobuf_types.FieldMask
				dst.FieldMask = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

var SetClientCollaboratorRequestFieldPathsNested = []string{
	"client_ids",
	"client_ids.client_id",
	"collaborator",
	"collaborator.ids",
	"collaborator.ids.ids.organization_ids",
	"collaborator.ids.ids.organization_ids.organization_id",
	"collaborator.ids.ids.user_ids",
	"collaborator.ids.ids.user_ids.email",
	"collaborator.ids.ids.user_ids.user_id",
	"collaborator.rights",
	"ids",
}

var SetClientCollaboratorRequestFieldPathsTopLevel = []string{
	"client_ids",
	"collaborator",
	"ids",
}

func (dst *SetClientCollaboratorRequest) SetFields(src *SetClientCollaboratorRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "client_ids":
			if len(subs) > 0 {
				newDst := &dst.ClientIdentifiers
				var newSrc *ClientIdentifiers
				if src != nil {
					newSrc = &src.ClientIdentifiers
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.ClientIdentifiers = src.ClientIdentifiers
				} else {
					var zero ClientIdentifiers
					dst.ClientIdentifiers = zero
				}
			}
		case "collaborator":
			if len(subs) > 0 {
				newDst := &dst.Collaborator
				var newSrc *Collaborator
				if src != nil {
					newSrc = &src.Collaborator
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Collaborator = src.Collaborator
				} else {
					var zero Collaborator
					dst.Collaborator = zero
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}
