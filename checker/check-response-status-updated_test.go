package checker_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tufin/oasdiff/checker"
	"github.com/tufin/oasdiff/diff"
)

// CL: Adding a success response status
func TestResponseSuccessStatusAdded(t *testing.T) {
	s1, _ := open("../data/checker/response_status_base.yaml")
	s2, err := open("../data/checker/response_status_base.yaml")
	require.NoError(t, err)

	// Add new success response
	s2.Spec.Paths["/api/v1.0/groups"].Post.Responses["201"] = s2.Spec.Paths["/api/v1.0/groups"].Post.Responses["200"]

	d, osm, err := diff.GetWithOperationsSourcesMap(getConfig(), s1, s2)
	require.NoError(t, err)
	errs := checker.CheckBackwardCompatibilityUntilLevel(singleCheckConfig(checker.ResponseSuccessStatusUpdated), d, osm, checker.INFO)
	require.NotEmpty(t, errs)
	require.Equal(t, checker.BackwardCompatibilityErrors{
		{
			Id:          "response-success-status-added",
			Text:        "added the success response with the status '201'",
			Comment:     "",
			Level:       checker.INFO,
			Operation:   "POST",
			Path:        "/api/v1.0/groups",
			Source:      "../data/checker/response_status_base.yaml",
			OperationId: "createOneGroup",
		},
	}, errs)
}

// CL: Adding a non-success response status
func TestResponseNonSuccessStatusAdded(t *testing.T) {
	s1, _ := open("../data/checker/response_status_base.yaml")
	s2, err := open("../data/checker/response_status_base.yaml")
	require.NoError(t, err)

	// Add new non-success response
	s2.Spec.Paths["/api/v1.0/groups"].Post.Responses["400"] = s2.Spec.Paths["/api/v1.0/groups"].Post.Responses["409"]

	d, osm, err := diff.GetWithOperationsSourcesMap(getConfig(), s1, s2)
	require.NoError(t, err)
	errs := checker.CheckBackwardCompatibilityUntilLevel(singleCheckConfig(checker.ResponseNonSuccessStatusUpdated), d, osm, checker.INFO)
	require.NotEmpty(t, errs)
	require.Equal(t, checker.BackwardCompatibilityErrors{
		{
			Id:          "response-non-success-status-added",
			Text:        "added the non-success response with the status '400'",
			Comment:     "",
			Level:       checker.INFO,
			Operation:   "POST",
			Path:        "/api/v1.0/groups",
			Source:      "../data/checker/response_status_base.yaml",
			OperationId: "createOneGroup",
		},
	}, errs)
}

// CL: Removing a non-success response status
func TestResponseNonSuccessStatusRemoved(t *testing.T) {
	s1, _ := open("../data/checker/response_status_base.yaml")
	s2, err := open("../data/checker/response_status_base.yaml")
	require.NoError(t, err)

	delete(s2.Spec.Paths["/api/v1.0/groups"].Post.Responses, "409")

	d, osm, err := diff.GetWithOperationsSourcesMap(getConfig(), s1, s2)
	require.NoError(t, err)
	errs := checker.CheckBackwardCompatibilityUntilLevel(singleCheckConfig(checker.ResponseNonSuccessStatusUpdated), d, osm, checker.INFO)
	require.NotEmpty(t, errs)
	require.Len(t, errs, 1)
	require.Equal(t, checker.BackwardCompatibilityErrors{
		{
			Id:          "response-non-success-status-removed",
			Text:        "removed the non-success response with the status '409'",
			Comment:     "",
			Level:       checker.INFO,
			Operation:   "POST",
			Path:        "/api/v1.0/groups",
			Source:      "../data/checker/response_status_base.yaml",
			OperationId: "createOneGroup",
		},
	}, errs)
}

// BC: Removing a success status is breaking
func TestResponseSuccessStatusRemoved(t *testing.T) {
	s1, _ := open("../data/checker/response_status_base.yaml")
	s2, err := open("../data/checker/response_status_base.yaml")
	require.NoError(t, err)

	delete(s2.Spec.Paths["/api/v1.0/groups"].Post.Responses, "200")

	d, osm, err := diff.GetWithOperationsSourcesMap(getConfig(), s1, s2)
	require.NoError(t, err)
	errs := checker.CheckBackwardCompatibilityUntilLevel(singleCheckConfig(checker.ResponseSuccessStatusUpdated), d, osm, checker.INFO)
	require.NotEmpty(t, errs)
	require.Len(t, errs, 1)
	require.Equal(t, checker.BackwardCompatibilityErrors{
		{
			Id:          "response-success-status-removed",
			Text:        "removed the success response with the status '200'",
			Comment:     "",
			Level:       checker.ERR,
			Operation:   "POST",
			Path:        "/api/v1.0/groups",
			Source:      "../data/checker/response_status_base.yaml",
			OperationId: "createOneGroup",
		},
	}, errs)
}