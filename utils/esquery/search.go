package esquery

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/elastic/go-elasticsearch/v8/esapi"
)

// SearchRequest represents a request to ElasticSearch's Search API, described
// in https://www.elastic.co/guide/en/elasticsearch/reference/current/search.html.
// Not all features of the search API are currently supported, but a request can
// currently include a query, aggregations, and more.
type SearchRequest struct {
	client *elasticsearch.Client
	index  []string

	aggs           []Aggregation
	explain        *bool
	from           *uint64
	highlight      Mappable
	searchAfter    []interface{}
	postFilter     Mappable
	query          Mappable
	size           *uint64
	sort           Sort
	source         Source
	trackTotalHits *bool
	timeout        *time.Duration
}

// SearchResult is the result of a search in Elasticsearch.
type SearchResult struct {
	Header          http.Header                `json:"-"`
	TookInMillis    int64                      `json:"took,omitempty"`             // search time in milliseconds
	TerminatedEarly bool                       `json:"terminated_early,omitempty"` // request terminated early
	NumReducePhases int                        `json:"num_reduce_phases,omitempty"`
	Clusters        *SearchResultCluster       `json:"_clusters,omitempty"`    // 6.1.0+
	ScrollId        string                     `json:"_scroll_id,omitempty"`   // only used with Scroll and Scan operations
	Hits            *SearchHits                `json:"hits,omitempty"`         // the actual search hits
	Suggest         SearchSuggest              `json:"suggest,omitempty"`      // results from suggesters
	Aggregations    map[string]json.RawMessage `json:"aggregations,omitempty"` // results from aggregations
	TimedOut        bool                       `json:"timed_out,omitempty"`    // true if the search timed out
	Error           *ErrorDetails              `json:"error,omitempty"`        // only used in MultiGet
	Status          int                        `json:"status,omitempty"`       // used in MultiSearch
	PitId           string                     `json:"pit_id,omitempty"`       // Point In Time ID
}

// SearchHits specifies the list of search hits.
type SearchHits struct {
	TotalHits *TotalHits   `json:"total,omitempty"`     // total number of hits found
	MaxScore  *float64     `json:"max_score,omitempty"` // maximum score of all hits
	Hits      []*SearchHit `json:"hits,omitempty"`      // the actual hits returned
}

// TotalHits specifies total number of hits and its relation
type TotalHits struct {
	Value    int64  `json:"value"`    // value of the total hit count
	Relation string `json:"relation"` // how the value should be interpreted: accurate ("eq") or a lower bound ("gte")
}

// SearchHit is a single hit.
type SearchHit struct {
	Score          *float64        `json:"_score,omitempty"`   // computed score
	Index          string          `json:"_index,omitempty"`   // index name
	Type           string          `json:"_type,omitempty"`    // type meta field
	Id             string          `json:"_id,omitempty"`      // external or internal
	Uid            string          `json:"_uid,omitempty"`     // uid meta field (see MapperService.java for all meta fields)
	Routing        string          `json:"_routing,omitempty"` // routing meta field
	Parent         string          `json:"_parent,omitempty"`  // parent meta field
	Version        *int64          `json:"_version,omitempty"` // version number, when Version is set to true in SearchService
	SeqNo          *int64          `json:"_seq_no"`
	PrimaryTerm    *int64          `json:"_primary_term"`
	Sort           []interface{}   `json:"sort,omitempty"`            // sort information
	Source         json.RawMessage `json:"_source,omitempty"`         // stored document source
	MatchedQueries []string        `json:"matched_queries,omitempty"` // matched queries
	Shard          string          `json:"_shard,omitempty"`          // used e.g. in Search Explain
	Node           string          `json:"_node,omitempty"`           // used e.g. in Search Explain

	// HighlightFields
	// SortValues
	// MatchedFilters
}

// SearchResultCluster holds information about a search response
// from a cluster.
type SearchResultCluster struct {
	Successful int `json:"successful,omitempty"`
	Total      int `json:"total,omitempty"`
	Skipped    int `json:"skipped,omitempty"`
}

// Suggest

// SearchSuggest is a map of suggestions.
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.0/search-suggesters.html.
type SearchSuggest map[string][]SearchSuggestion

// SearchSuggestion is a single search suggestion.
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.0/search-suggesters.html.
type SearchSuggestion struct {
	Text    string                   `json:"text"`
	Offset  int                      `json:"offset"`
	Length  int                      `json:"length"`
	Options []SearchSuggestionOption `json:"options"`
}

// SearchSuggestionOption is an option of a SearchSuggestion.
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.0/search-suggesters.html.
type SearchSuggestionOption struct {
	Text            string              `json:"text"`
	Index           string              `json:"_index"`
	Type            string              `json:"_type"`
	Id              string              `json:"_id"`
	Score           float64             `json:"score"`  // term and phrase suggesters uses "score" as of 6.2.4
	ScoreUnderscore float64             `json:"_score"` // completion and context suggesters uses "_score" as of 6.2.4
	Highlighted     string              `json:"highlighted"`
	CollateMatch    bool                `json:"collate_match"`
	Freq            int                 `json:"freq"` // from TermSuggestion.Option in Java API
	Source          json.RawMessage     `json:"_source"`
	Contexts        map[string][]string `json:"contexts,omitempty"`
}

// ErrorDetails encapsulate error details from Elasticsearch.
// It is used in e.g. elastic.Error and elastic.BulkResponseItem.
type ErrorDetails struct {
	Type         string                   `json:"type"`
	Reason       string                   `json:"reason"`
	ResourceType string                   `json:"resource.type,omitempty"`
	ResourceId   string                   `json:"resource.id,omitempty"`
	Index        string                   `json:"index,omitempty"`
	Phase        string                   `json:"phase,omitempty"`
	Grouped      bool                     `json:"grouped,omitempty"`
	CausedBy     map[string]interface{}   `json:"caused_by,omitempty"`
	RootCause    []*ErrorDetails          `json:"root_cause,omitempty"`
	Suppressed   []*ErrorDetails          `json:"suppressed,omitempty"`
	FailedShards []map[string]interface{} `json:"failed_shards,omitempty"`
	Header       map[string]interface{}   `json:"header,omitempty"`

	// ScriptException adds the information in the following block.

	ScriptStack []string             `json:"script_stack,omitempty"` // from ScriptException
	Script      string               `json:"script,omitempty"`       // from ScriptException
	Lang        string               `json:"lang,omitempty"`         // from ScriptException
	Position    *ScriptErrorPosition `json:"position,omitempty"`     // from ScriptException (7.7+)
}

// ScriptErrorPosition specifies the position of the error
// in a script. It is used in ErrorDetails for scripting errors.
type ScriptErrorPosition struct {
	Offset int `json:"offset"`
	Start  int `json:"start"`
	End    int `json:"end"`
}

// Search creates a new SearchRequest object, to be filled via method chaining.
func Search() *SearchRequest {
	return &SearchRequest{}
}

func (req *SearchRequest) SetClient(es *elasticsearch.Client) {
	req.client = es
}

func (req *SearchRequest) Index(v ...string) *SearchRequest {
	req.index = v
	return req
}

// Query sets a query for the request.
func (req *SearchRequest) Query(q Mappable) *SearchRequest {
	req.query = q
	return req
}

// Aggs sets one or more aggregations for the request.
func (req *SearchRequest) Aggs(aggs ...Aggregation) *SearchRequest {
	req.aggs = append(req.aggs, aggs...)
	return req
}

// PostFilter sets a post_filter for the request.
func (req *SearchRequest) PostFilter(filter Mappable) *SearchRequest {
	req.postFilter = filter
	return req
}

// From sets a document offset to start from.
func (req *SearchRequest) From(offset uint64) *SearchRequest {
	req.from = &offset
	return req
}

// Size sets the number of hits to return. The default - according to the ES
// documentation - is 10.
func (req *SearchRequest) Size(size uint64) *SearchRequest {
	req.size = &size
	return req
}

// Sort sets how the results should be sorted.
func (req *SearchRequest) Sort(sorter Sort) *SearchRequest {
	req.sort = sorter

	return req
}

// SortByName sets how the results should be sorted.
func (req *SearchRequest) SortByName(name string, order Order) *SearchRequest {
	req.sort = append(req.sort, map[string]interface{}{
		name: map[string]interface{}{
			"order": order,
		},
	})

	return req
}

// SearchAfter retrieve the sorted result
func (req *SearchRequest) SearchAfter(s ...interface{}) *SearchRequest {
	req.searchAfter = append(req.searchAfter, s...)
	return req
}

// Explain sets whether the ElasticSearch API should return an explanation for
// how each hit's score was calculated.
func (req *SearchRequest) Explain(b bool) *SearchRequest {
	req.explain = &b
	return req
}

// Timeout sets a timeout for the request.
func (req *SearchRequest) Timeout(dur time.Duration) *SearchRequest {
	req.timeout = &dur
	return req
}

// SourceIncludes sets the keys to return from the matching documents.
func (req *SearchRequest) SourceIncludes(keys ...string) *SearchRequest {
	req.source.includes = keys
	return req
}

// SourceExcludes sets the keys to not return from the matching documents.
func (req *SearchRequest) SourceExcludes(keys ...string) *SearchRequest {
	req.source.excludes = keys
	return req
}

// WithoutSource exclude all sources
func (req *SearchRequest) WithoutSource() *SearchRequest {
	req.source.without = true
	return req
}

// WithTrackTotalHits return total count > 10000
func (req *SearchRequest) WithTrackTotalHits() *SearchRequest {
	trackTotalHits := true
	req.trackTotalHits = &trackTotalHits
	return req
}

// Highlight sets a highlight for the request.
func (req *SearchRequest) Highlight(highlight Mappable) *SearchRequest {
	req.highlight = highlight
	return req
}

// Map implements the Mappable interface. It converts the request to into a
// nested map[string]interface{}, as expected by the go-elasticsearch library.
func (req *SearchRequest) Map() map[string]interface{} {
	m := make(map[string]interface{})

	//by default, track_total_hits is true
	//https://www.elastic.co/guide/en/elasticsearch/reference/master/search-your-data.html#track-total-hits
	if req.trackTotalHits != nil {
		m["track_total_hits"] = *req.trackTotalHits
	}

	if req.query != nil {
		m["query"] = req.query.Map()
	}
	if len(req.aggs) > 0 {
		aggs := make(map[string]interface{})
		for _, agg := range req.aggs {
			aggs[agg.Name()] = agg.Map()
		}

		m["aggs"] = aggs
	}
	if req.postFilter != nil {
		m["post_filter"] = req.postFilter.Map()
	}
	if req.size != nil {
		m["size"] = *req.size
	}
	if len(req.sort) > 0 {
		m["sort"] = req.sort
	}
	if req.from != nil {
		m["from"] = *req.from
	}
	if req.explain != nil {
		m["explain"] = *req.explain
	}
	if req.timeout != nil {
		m["timeout"] = fmt.Sprintf("%.0fs", req.timeout.Seconds())
	}
	if req.highlight != nil {
		m["highlight"] = req.highlight.Map()
	}
	if req.searchAfter != nil {
		m["search_after"] = req.searchAfter
	}

	switch req.source.Map().(type) {
	case map[string]interface{}:
		source := req.source.Map().(map[string]interface{})
		if len(source) > 0 {
			m["_source"] = source
		}
	case bool:
		if req.source.Map().(bool) {
			m["_source"] = false
		}
	}

	return m
}

// MarshalJSON implements the json.Marshaler interface. It returns a JSON
// representation of the map generated by the SearchRequest's Map method.
func (req *SearchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(req.Map())
}

// Run executes the request using the provided ElasticSearch client. Zero or
// more search options can be provided as well. It returns the standard Response
// type of the official Go client.
func (req *SearchRequest) Run(
	api *elasticsearch.Client,
	o ...func(*esapi.SearchRequest),
) (res *esapi.Response, err error) {
	return req.RunSearch(api.Search, o...)
}

func (req *SearchRequest) Do(ctx context.Context) (res *SearchResult, err error) {
	var option []func(*esapi.SearchRequest)

	option = append(option, req.client.Search.WithContext(ctx))
	option = append(option, req.client.Search.WithIndex(req.index...))

	result, err := req.RunSearch(req.client.Search, option...)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(body), &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

//func (req *SearchRequest) Do(
//	api *elasticsearch.Client,
//	o ...func(*esapi.SearchRequest),
//) (res *SearchResult, err error) {
//	result, err := req.RunSearch(api.Search, o...)
//	if err != nil {
//		return nil, err
//	}
//
//	body, err := ioutil.ReadAll(result.Body)
//	if err != nil {
//		return nil, err
//	}
//
//	err = json.Unmarshal([]byte(body), &res)
//	if err != nil {
//		return nil, err
//	}
//
//	return res, nil
//}

func (r *SearchResult) TotalHits() int64 {
	if r != nil && r.Hits != nil && r.Hits.TotalHits != nil {
		return r.Hits.TotalHits.Value
	}
	return 0
}

// RunSearch is the same as the Run method, except that it accepts a value of
// type esapi.Search (usually this is the Search field of an elasticsearch.Client
// object). Since the ElasticSearch client does not provide an interface type
// for its API (which would allow implementation of mock clients), this provides
// a workaround. The Search function in the ES client is actually a field of a
// function type.
func (req *SearchRequest) RunSearch(
	search esapi.Search,
	o ...func(*esapi.SearchRequest),
) (res *esapi.Response, err error) {
	var b bytes.Buffer
	err = json.NewEncoder(&b).Encode(req.Map())
	if err != nil {
		return nil, err
	}

	opts := append([]func(*esapi.SearchRequest){search.WithBody(&b)}, o...)

	return search(opts...)
}

// Query is a shortcut for creating a SearchRequest with only a query. It is
// mostly included to maintain the API provided by esquery in early releases.
func Query(q Mappable) *SearchRequest {
	return Search().Query(q)
}

// Aggregate is a shortcut for creating a SearchRequest with aggregations. It is
// mostly included to maintain the API provided by esquery in early releases.
func Aggregate(aggs ...Aggregation) *SearchRequest {
	return Search().Aggs(aggs...)
}
