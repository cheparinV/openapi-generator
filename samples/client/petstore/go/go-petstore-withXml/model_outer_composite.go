/*
 * OpenAPI Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package petstore

type OuterComposite struct {
	MyNumber  float32 `json:"my_number,omitempty" xml:"my_number"`
	MyString  string  `json:"my_string,omitempty" xml:"my_string"`
	MyBoolean bool    `json:"my_boolean,omitempty" xml:"my_boolean"`
}
