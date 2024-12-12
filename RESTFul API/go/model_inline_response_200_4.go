/*
 * Production and Distribution Tracking API
 *
 * API for monitoring the production and distribution of documents or currency in real-time, ensuring transparency and security in the distribution chain.
 *
 * API version: 1.0.0
 * Contact: support@your-company.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type InlineResponse2004 struct {

	TotalDistributions int32 `json:"total_distributions,omitempty"`

	Delivered int32 `json:"delivered,omitempty"`
}