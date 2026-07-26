package v4

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/utils/ptr"
)

// TestHeaderArgs_NoHeaders ensures a context without any headers yields nil args,
// preserving the pre-existing behaviour.
func TestHeaderArgs_NoHeaders(t *testing.T) {
	assert.Nil(t, headerArgs(context.Background()))
}

// TestWithRequestID_HeaderArgs ensures WithRequestID puts the request id on the
// context.
func TestWithRequestID_HeaderArgs(t *testing.T) {
	const id = "idempotency-key-123"
	ctx := WithRequestID(context.Background(), id)

	args := headerArgs(ctx)
	require.Len(t, args, 1)

	val, ok := args[requestIDHeader]
	require.True(t, ok, "expected header %q to be present", requestIDHeader)

	// The SDK only honours values that assert to *string.
	strPtr, ok := val.(*string)
	require.True(t, ok, "header value must be a *string, got %T", val)
	assert.Equal(t, id, *strPtr)
}

// TestWithHeader_AccumulateAndOverride ensures multiple headers accumulate and a
// later value overrides an earlier one for the same key.
func TestWithHeader_AccumulateAndOverride(t *testing.T) {
	ctx := withHeader(context.Background(), "A", "1")
	ctx = withHeader(ctx, "B", "2")
	ctx = withHeader(ctx, "A", "3") // override A

	args := headerArgs(ctx)
	require.Len(t, args, 2)
	assert.Equal(t, "3", *(args["A"].(*string)))
	assert.Equal(t, "2", *(args["B"].(*string)))
}

// TestWithHeader_DoesNotMutateParent ensures deriving a child context does not
// mutate the header map carried by the parent (copy-on-write).
func TestWithHeader_DoesNotMutateParent(t *testing.T) {
	parent := withHeader(context.Background(), "A", "1")
	_ = withHeader(parent, "B", "2") // child adds B

	parentArgs := headerArgs(parent)
	assert.Len(t, parentArgs, 1, "parent context must not see the child's header")
	_, hasB := parentArgs["B"]
	assert.False(t, hasB)
}

// TestHeaderArgs_DistinctPointers ensures each header maps to its own *string and
// values are not aliased across keys.
func TestHeaderArgs_DistinctPointers(t *testing.T) {
	ctx := withHeader(context.Background(), "A", "1")
	ctx = withHeader(ctx, "B", "2")

	args := headerArgs(ctx)
	assert.NotSame(t, args["A"].(*string), args["B"].(*string))
	assert.Equal(t, "1", *(args["A"].(*string)))
	assert.Equal(t, "2", *(args["B"].(*string)))
}

func TestRequestIDHeader_UsesSDKCasing(t *testing.T) {
	assert.Equal(t, "NTNX-Request-Id", requestIDHeader)

	ctx := WithRequestID(context.Background(), "id")
	_, ok := headerArgs(ctx)["NTNX-Request-Id"]
	assert.True(t, ok, "WithRequestID must set the header under the exact SDK key")
}

// TestWithRequestID_LastWins ensures re-stamping a context with a new request id
// overrides the earlier one, so retries carry a single, deterministic key.
func TestWithRequestID_LastWins(t *testing.T) {
	ctx := WithRequestID(context.Background(), "first")
	ctx = WithRequestID(ctx, "second")

	args := headerArgs(ctx)
	require.Len(t, args, 1)
	assert.Equal(t, "second", *(args[requestIDHeader].(*string)))
}

// TestHeaderArgs_ReturnedMapIsIsolated ensures mutating the map returned by
// headerArgs does not affect what a later call derives from the same context.
func TestHeaderArgs_ReturnedMapIsIsolated(t *testing.T) {
	ctx := WithRequestID(context.Background(), "id")

	first := headerArgs(ctx)
	delete(first, requestIDHeader)

	second := headerArgs(ctx)
	_, ok := second[requestIDHeader]
	assert.True(t, ok, "a second headerArgs call must not be affected by mutating the first result")
}

// fakeListMetadata mimics the generated *ApiResponseMetadata shape that
// GetMetadataTotalResults reflects over (it only needs a TotalAvailableResults
// *int field).
type fakeListMetadata struct {
	TotalAvailableResults *int
}

// fakeListResponse mimics a generated *List<Entity>ApiResponse: it exposes a
// Metadata field for reflection and a GetData accessor for the page payload.
type fakeListResponse struct {
	Metadata *fakeListMetadata
	data     []string
}

func (r *fakeListResponse) GetData() any {
	if r.data == nil {
		return nil
	}
	return r.data
}

// TestGenericListEntities_StopsOnEmptyPage ensures the list-all pagination loop
// terminates when a page returns no items, even though the cumulative number of
// collected items never reaches the totalCount snapshot taken from the first
// page. This reproduces the runaway pagination observed against a PC whose
// image set shrank (or whose TotalAvailableResults was inconsistent with the
// page data) during a list, which previously caused the loop to request pages
// far beyond the last page of real data.
func TestGenericListEntities_StopsOnEmptyPage(t *testing.T) {
	const total = 10 // reported by the first page's metadata
	// Only pages 0 and 1 carry data (6 items total); every later page is empty,
	// so len(result) never reaches total. Without the empty-page guard the loop
	// would page forever.
	pageData := map[int][]string{
		0: {"a", "b", "c"},
		1: {"d", "e", "f"},
	}

	const maxPagesBeforeFailing = 100
	callCount := 0

	apiCall := func(reqParams *V4ODataParams) (*fakeListResponse, error) {
		callCount++
		require.LessOrEqual(t, callCount, maxPagesBeforeFailing,
			"GenericListEntities paged past the available data; the empty-page guard is not working")

		page := 0
		if reqParams.Page != nil {
			page = *reqParams.Page
		}
		return &fakeListResponse{
			Metadata: &fakeListMetadata{TotalAvailableResults: ptr.To(total)},
			data:     pageData[page],
		}, nil
	}

	result, err := GenericListEntities[*fakeListResponse, string](apiCall, nil, "fake")
	require.NoError(t, err)
	assert.Equal(t, []string{"a", "b", "c", "d", "e", "f"}, result)
}

// TestGenericListEntities_PagesUntilTotalCount ensures the loop still walks
// every page until totalCount is reached when the data set is consistent.
func TestGenericListEntities_PagesUntilTotalCount(t *testing.T) {
	const total = 5
	pageData := map[int][]string{
		0: {"a", "b"},
		1: {"c", "d"},
		2: {"e"},
	}

	callCount := 0
	apiCall := func(reqParams *V4ODataParams) (*fakeListResponse, error) {
		callCount++
		page := 0
		if reqParams.Page != nil {
			page = *reqParams.Page
		}
		return &fakeListResponse{
			Metadata: &fakeListMetadata{TotalAvailableResults: ptr.To(total)},
			data:     pageData[page],
		}, nil
	}

	result, err := GenericListEntities[*fakeListResponse, string](apiCall, nil, "fake")
	require.NoError(t, err)
	assert.Equal(t, []string{"a", "b", "c", "d", "e"}, result)
	// Pages 0, 1 and 2 are fetched; once len(result) == total the loop stops.
	assert.Equal(t, 3, callCount)
}
