/*
 * HyperOne API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ProjectInvoices struct for ProjectInvoices
type ProjectInvoices struct {
	Id           string           `json:"id,omitempty"`
	InvoiceNo    string           `json:"invoiceNo,omitempty"`
	IssueDate    string           `json:"issueDate,omitempty"`
	Summary      string           `json:"summary,omitempty"`
	Project      string           `json:"project,omitempty"`
	Organisation string           `json:"organisation,omitempty"`
	Duplicate    ProjectDuplicate `json:"duplicate,omitempty"`
}
