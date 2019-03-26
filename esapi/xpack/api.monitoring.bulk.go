// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"io"
	"strings"
)

func newMonitoringBulkFunc(t Transport) MonitoringBulk {
	return func(body io.Reader, o ...func(*MonitoringBulkRequest)) (*Response, error) {
		var r = MonitoringBulkRequest{Body: body}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at http://www.elastic.co/guide/en/monitoring/current/appendix-api-bulk.html.
//
type MonitoringBulk func(body io.Reader, o ...func(*MonitoringBulkRequest)) (*Response, error)

// MonitoringBulkRequest configures the Monitoring Bulk API request.
//
type MonitoringBulkRequest struct {
	DocumentType string
	Body         io.Reader

	Interval         string
	SystemApiVersion string
	SystemID         string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r MonitoringBulkRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "POST"

	path.Grow(1 + len("_monitoring") + 1 + len(r.DocumentType) + 1 + len("bulk"))
	path.WriteString("/")
	path.WriteString("_monitoring")
	if r.DocumentType != "" {
		path.WriteString("/")
		path.WriteString(r.DocumentType)
	}
	path.WriteString("/")
	path.WriteString("bulk")

	params = make(map[string]string)

	if r.Interval != "" {
		params["interval"] = r.Interval
	}

	if r.SystemApiVersion != "" {
		params["system_api_version"] = r.SystemApiVersion
	}

	if r.SystemID != "" {
		params["system_id"] = r.SystemID
	}

	if r.Pretty {
		params["pretty"] = "true"
	}

	if r.Human {
		params["human"] = "true"
	}

	if r.ErrorTrace {
		params["error_trace"] = "true"
	}

	if len(r.FilterPath) > 0 {
		params["filter_path"] = strings.Join(r.FilterPath, ",")
	}

	req, _ := newRequest(method, path.String(), r.Body)

	if len(params) > 0 {
		q := req.URL.Query()
		for k, v := range params {
			q.Set(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	if r.Body != nil {
		req.Header[headerContentType] = headerContentTypeJSON
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	res, err := transport.Perform(req)
	if err != nil {
		return nil, err
	}

	response := Response{
		StatusCode: res.StatusCode,
		Body:       res.Body,
		Header:     res.Header,
	}

	return &response, nil
}

// WithContext sets the request context.
//
func (f MonitoringBulk) WithContext(v context.Context) func(*MonitoringBulkRequest) {
	return func(r *MonitoringBulkRequest) {
		r.ctx = v
	}
}

// WithDocumentType - default document type for items which don't provide one.
//
func (f MonitoringBulk) WithDocumentType(v string) func(*MonitoringBulkRequest) {
	return func(r *MonitoringBulkRequest) {
		r.DocumentType = v
	}
}

// WithInterval - collection interval (e.g., '10s' or '10000ms') of the payload.
//
func (f MonitoringBulk) WithInterval(v string) func(*MonitoringBulkRequest) {
	return func(r *MonitoringBulkRequest) {
		r.Interval = v
	}
}

// WithSystemApiVersion - api version of the monitored system.
//
func (f MonitoringBulk) WithSystemApiVersion(v string) func(*MonitoringBulkRequest) {
	return func(r *MonitoringBulkRequest) {
		r.SystemApiVersion = v
	}
}

// WithSystemID - identifier of the monitored system.
//
func (f MonitoringBulk) WithSystemID(v string) func(*MonitoringBulkRequest) {
	return func(r *MonitoringBulkRequest) {
		r.SystemID = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f MonitoringBulk) WithPretty() func(*MonitoringBulkRequest) {
	return func(r *MonitoringBulkRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f MonitoringBulk) WithHuman() func(*MonitoringBulkRequest) {
	return func(r *MonitoringBulkRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f MonitoringBulk) WithErrorTrace() func(*MonitoringBulkRequest) {
	return func(r *MonitoringBulkRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f MonitoringBulk) WithFilterPath(v ...string) func(*MonitoringBulkRequest) {
	return func(r *MonitoringBulkRequest) {
		r.FilterPath = v
	}
}