// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	"reflect"
	"github.com/petergtz/pegomock"
	http "net/http"
)

func AnyPtrToHttpClient() *http.Client {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(*http.Client))(nil)).Elem()))
	var nullValue *http.Client
	return nullValue
}

func EqPtrToHttpClient(value *http.Client) *http.Client {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue *http.Client
	return nullValue
}
