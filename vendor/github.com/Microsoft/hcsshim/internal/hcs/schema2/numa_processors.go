// Autogenerated code; DO NOT EDIT.

/*
 * Schema Open API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package hcsschema

type NumaProcessors struct {
	CountPerNode  Range  `json:"count_per_node,omitempty"`
	NodePerSocket uint32 `json:"node_per_socket,omitempty"`
}

type Range struct {
	Max uint32 `json:"max,omitempty"`
}
