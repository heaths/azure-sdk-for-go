//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package internal

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
	"strings"
	"time"
)

const (
	rfc1123JSON = `"` + time.RFC1123 + `"`
)

type timeRFC1123 time.Time

func (t timeRFC1123) MarshalJSON() ([]byte, error) {
	b := []byte(time.Time(t).Format(rfc1123JSON))
	return b, nil
}

func (t timeRFC1123) MarshalText() ([]byte, error) {
	b := []byte(time.Time(t).Format(time.RFC1123))
	return b, nil
}

func (t *timeRFC1123) UnmarshalJSON(data []byte) error {
	p, err := time.Parse(rfc1123JSON, strings.ToUpper(string(data)))
	*t = timeRFC1123(p)
	return err
}

func (t *timeRFC1123) UnmarshalText(data []byte) error {
	p, err := time.Parse(time.RFC1123, string(data))
	*t = timeRFC1123(p)
	return err
}

func populateTimeRFC1123(m map[string]interface{}, k string, t *time.Time) {
	if t == nil {
		return
	} else if azcore.IsNullValue(t) {
		m[k] = nil
		return
	} else if reflect.ValueOf(t).IsNil() {
		return
	}
	m[k] = (*timeRFC1123)(t)
}

func unpopulateTimeRFC1123(data json.RawMessage, t **time.Time) error {
	if data == nil || strings.EqualFold(string(data), "null") {
		return nil
	}
	var aux timeRFC1123
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	*t = (*time.Time)(&aux)
	return nil
}
