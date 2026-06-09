package v4

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/utils/ptr"
)

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

// TestGenericListEntities_CollectsItemsAddedWhilePaging ensures that entities
// added after the first page is fetched are still collected, by re-reading the
// total from each page instead of trusting the first page's (now stale) total.
func TestGenericListEntities_CollectsItemsAddedWhilePaging(t *testing.T) {
	// Page 0 reports a total of 4. While we are still paginating (before len has
	// caught up to that total) two items are added, and the total reported by
	// page 1 grows to 6, extending the walk to collect them.
	pages := []struct {
		data  []string
		total int
	}{
		{data: []string{"a", "b"}, total: 4}, // first page: total observed as 4
		{data: []string{"c", "d"}, total: 6}, // two items added; total now 6
		{data: []string{"e", "f"}, total: 6}, // the added items
	}

	callCount := 0
	apiCall := func(reqParams *V4ODataParams) (*fakeListResponse, error) {
		page := 0
		if reqParams.Page != nil {
			page = *reqParams.Page
		}
		callCount++
		require.Less(t, page, len(pages), "paged past the fixtures")
		return &fakeListResponse{
			Metadata: &fakeListMetadata{TotalAvailableResults: ptr.To(pages[page].total)},
			data:     pages[page].data,
		}, nil
	}

	result, err := GenericListEntities[*fakeListResponse, string](apiCall, nil, "fake")
	require.NoError(t, err)
	// Without refreshing totalCount the loop would have stopped after page 1
	// (len == 4 == stale total) and dropped "e" and "f".
	assert.Equal(t, []string{"a", "b", "c", "d", "e", "f"}, result)
}
