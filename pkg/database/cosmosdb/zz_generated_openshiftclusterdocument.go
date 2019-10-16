// Code generated by github.com/jim-minter/go-cosmosdb, DO NOT EDIT.

package cosmosdb

import (
	"net/http"

	pkg "github.com/jim-minter/rp/pkg/api"
)

type openShiftClusterDocumentClient struct {
	*databaseClient
	path string
}

// OpenShiftClusterDocumentClient is a openShiftClusterDocument client
type OpenShiftClusterDocumentClient interface {
	Create(string, *pkg.OpenShiftClusterDocument) (*pkg.OpenShiftClusterDocument, error)
	List() OpenShiftClusterDocumentIterator
	ListAll() (*pkg.OpenShiftClusterDocuments, error)
	Get(string, string) (*pkg.OpenShiftClusterDocument, error)
	Replace(string, *pkg.OpenShiftClusterDocument) (*pkg.OpenShiftClusterDocument, error)
	Delete(string, *pkg.OpenShiftClusterDocument) error
	Query(string, *Query) OpenShiftClusterDocumentIterator
	QueryAll(string, *Query) (*pkg.OpenShiftClusterDocuments, error)
}

type openShiftClusterDocumentListIterator struct {
	*openShiftClusterDocumentClient
	continuation string
	done         bool
}

type openShiftClusterDocumentQueryIterator struct {
	*openShiftClusterDocumentClient
	partitionkey string
	query        *Query
	continuation string
	done         bool
}

// OpenShiftClusterDocumentIterator is a openShiftClusterDocument iterator
type OpenShiftClusterDocumentIterator interface {
	Next() (*pkg.OpenShiftClusterDocuments, error)
}

// NewOpenShiftClusterDocumentClient returns a new openShiftClusterDocument client
func NewOpenShiftClusterDocumentClient(collc CollectionClient, collid string) OpenShiftClusterDocumentClient {
	return &openShiftClusterDocumentClient{
		databaseClient: collc.(*collectionClient).databaseClient,
		path:           collc.(*collectionClient).path + "/colls/" + collid,
	}
}

func (c *openShiftClusterDocumentClient) all(i OpenShiftClusterDocumentIterator) (*pkg.OpenShiftClusterDocuments, error) {
	allopenShiftClusterDocuments := &pkg.OpenShiftClusterDocuments{}

	for {
		openShiftClusterDocuments, err := i.Next()
		if err != nil {
			return nil, err
		}
		if openShiftClusterDocuments == nil {
			break
		}

		allopenShiftClusterDocuments.Count += openShiftClusterDocuments.Count
		allopenShiftClusterDocuments.ResourceID = openShiftClusterDocuments.ResourceID
		allopenShiftClusterDocuments.OpenShiftClusterDocuments = append(allopenShiftClusterDocuments.OpenShiftClusterDocuments, openShiftClusterDocuments.OpenShiftClusterDocuments...)
	}

	return allopenShiftClusterDocuments, nil
}

func (c *openShiftClusterDocumentClient) Create(partitionkey string, newopenShiftClusterDocument *pkg.OpenShiftClusterDocument) (openShiftClusterDocument *pkg.OpenShiftClusterDocument, err error) {
	headers := http.Header{}
	if partitionkey != "" {
		headers.Set("X-Ms-Documentdb-Partitionkey", `["`+partitionkey+`"]`)
	}
	err = c.do(http.MethodPost, c.path+"/docs", "docs", c.path, http.StatusCreated, &newopenShiftClusterDocument, &openShiftClusterDocument, headers)
	return
}

func (c *openShiftClusterDocumentClient) List() OpenShiftClusterDocumentIterator {
	return &openShiftClusterDocumentListIterator{openShiftClusterDocumentClient: c}
}

func (c *openShiftClusterDocumentClient) ListAll() (*pkg.OpenShiftClusterDocuments, error) {
	return c.all(c.List())
}

func (c *openShiftClusterDocumentClient) Get(partitionkey, openShiftClusterDocumentid string) (openShiftClusterDocument *pkg.OpenShiftClusterDocument, err error) {
	headers := http.Header{}
	if partitionkey != "" {
		headers.Set("X-Ms-Documentdb-Partitionkey", `["`+partitionkey+`"]`)
	}
	err = c.do(http.MethodGet, c.path+"/docs/"+openShiftClusterDocumentid, "docs", c.path+"/docs/"+openShiftClusterDocumentid, http.StatusOK, nil, &openShiftClusterDocument, headers)
	return
}

func (c *openShiftClusterDocumentClient) Replace(partitionkey string, newopenShiftClusterDocument *pkg.OpenShiftClusterDocument) (openShiftClusterDocument *pkg.OpenShiftClusterDocument, err error) {
	if newopenShiftClusterDocument.ETag == "" {
		return nil, ErrETagRequired
	}
	headers := http.Header{}
	headers.Set("If-Match", newopenShiftClusterDocument.ETag)
	if partitionkey != "" {
		headers.Set("X-Ms-Documentdb-Partitionkey", `["`+partitionkey+`"]`)
	}
	err = c.do(http.MethodPut, c.path+"/docs/"+newopenShiftClusterDocument.ID, "docs", c.path+"/docs/"+newopenShiftClusterDocument.ID, http.StatusOK, &newopenShiftClusterDocument, &openShiftClusterDocument, headers)
	return
}

func (c *openShiftClusterDocumentClient) Delete(partitionkey string, openShiftClusterDocument *pkg.OpenShiftClusterDocument) error {
	if openShiftClusterDocument.ETag == "" {
		return ErrETagRequired
	}
	headers := http.Header{}
	headers.Set("If-Match", openShiftClusterDocument.ETag)
	if partitionkey != "" {
		headers.Set("X-Ms-Documentdb-Partitionkey", `["`+partitionkey+`"]`)
	}
	return c.do(http.MethodDelete, c.path+"/docs/"+openShiftClusterDocument.ID, "docs", c.path+"/docs/"+openShiftClusterDocument.ID, http.StatusNoContent, nil, nil, headers)
}

func (c *openShiftClusterDocumentClient) Query(partitionkey string, query *Query) OpenShiftClusterDocumentIterator {
	return &openShiftClusterDocumentQueryIterator{openShiftClusterDocumentClient: c, partitionkey: partitionkey, query: query}
}

func (c *openShiftClusterDocumentClient) QueryAll(partitionkey string, query *Query) (*pkg.OpenShiftClusterDocuments, error) {
	return c.all(c.Query(partitionkey, query))
}

func (i *openShiftClusterDocumentListIterator) Next() (openShiftClusterDocuments *pkg.OpenShiftClusterDocuments, err error) {
	if i.done {
		return
	}

	headers := http.Header{}
	if i.continuation != "" {
		headers.Set("X-Ms-Continuation", i.continuation)
	}

	err = i.do(http.MethodGet, i.path+"/docs", "docs", i.path, http.StatusOK, nil, &openShiftClusterDocuments, headers)
	if err != nil {
		return
	}

	i.continuation = headers.Get("X-Ms-Continuation")
	i.done = i.continuation == ""

	return
}

func (i *openShiftClusterDocumentQueryIterator) Next() (openShiftClusterDocuments *pkg.OpenShiftClusterDocuments, err error) {
	if i.done {
		return
	}

	headers := http.Header{}
	headers.Set("X-Ms-Documentdb-Isquery", "True")
	headers.Set("Content-Type", "application/query+json")
	if i.partitionkey != "" {
		headers.Set("X-Ms-Documentdb-Partitionkey", `["`+i.partitionkey+`"]`)
	} else {
		headers.Set("X-Ms-Documentdb-Query-Enablecrosspartition", "True")
	}
	if i.continuation != "" {
		headers.Set("X-Ms-Continuation", i.continuation)
	}

	err = i.do(http.MethodPost, i.path+"/docs", "docs", i.path, http.StatusOK, &i.query, &openShiftClusterDocuments, headers)
	if err != nil {
		return
	}

	i.continuation = headers.Get("X-Ms-Continuation")
	i.done = i.continuation == ""

	return
}