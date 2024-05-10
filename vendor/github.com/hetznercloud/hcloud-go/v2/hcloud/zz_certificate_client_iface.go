// Code generated by ifacemaker; DO NOT EDIT.

package hcloud

import (
	"context"
)

// ICertificateClient ...
type ICertificateClient interface {
	// GetByID retrieves a Certificate by its ID. If the Certificate does not exist, nil is returned.
	GetByID(ctx context.Context, id int64) (*Certificate, *Response, error)
	// GetByName retrieves a Certificate by its name. If the Certificate does not exist, nil is returned.
	GetByName(ctx context.Context, name string) (*Certificate, *Response, error)
	// Get retrieves a Certificate by its ID if the input can be parsed as an integer, otherwise it
	// retrieves a Certificate by its name. If the Certificate does not exist, nil is returned.
	Get(ctx context.Context, idOrName string) (*Certificate, *Response, error)
	// List returns a list of Certificates for a specific page.
	//
	// Please note that filters specified in opts are not taken into account
	// when their value corresponds to their zero value or when they are empty.
	List(ctx context.Context, opts CertificateListOpts) ([]*Certificate, *Response, error)
	// All returns all Certificates.
	All(ctx context.Context) ([]*Certificate, error)
	// AllWithOpts returns all Certificates for the given options.
	AllWithOpts(ctx context.Context, opts CertificateListOpts) ([]*Certificate, error)
	// Create creates a new uploaded certificate.
	//
	// Create returns an error for certificates of any other type. Use
	// CreateCertificate to create such certificates.
	Create(ctx context.Context, opts CertificateCreateOpts) (*Certificate, *Response, error)
	// CreateCertificate creates a new certificate of any type.
	CreateCertificate(ctx context.Context, opts CertificateCreateOpts) (CertificateCreateResult, *Response, error)
	// Update updates a Certificate.
	Update(ctx context.Context, certificate *Certificate, opts CertificateUpdateOpts) (*Certificate, *Response, error)
	// Delete deletes a certificate.
	Delete(ctx context.Context, certificate *Certificate) (*Response, error)
	// RetryIssuance retries the issuance of a failed managed certificate.
	RetryIssuance(ctx context.Context, certificate *Certificate) (*Action, *Response, error)
}
