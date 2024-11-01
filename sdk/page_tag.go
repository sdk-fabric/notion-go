
// PageTag automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app



import (
    "bytes"
    "encoding/json"
    "errors"
    "fmt"
    
    "io"
    "net/http"
    "net/url"
    
)

type PageTag struct {
    internal *sdkgen.TagAbstract
}



// Get Retrieves a Page object using the ID specified.
func (client *PageTag) Get(pageId string) (Page, error) {
    pathParams := make(map[string]interface{})
    pathParams["page_id"] = pageId

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/v1/pages/:page_id", pathParams))
    if err != nil {
        return Page{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return Page{}, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return Page{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return Page{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var data Page
        err := json.Unmarshal(respBody, &data)

        return data, err
    }

    var statusCode = resp.StatusCode
    return Page{}, errors.New(fmt.Sprint("The server returned an unknown status code: ", statusCode))
}

// Create Creates a new page that is a child of an existing page or database.
func (client *PageTag) Create(payload Page) (Page, error) {
    pathParams := make(map[string]interface{})

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/v1/pages", pathParams))
    if err != nil {
        return Page{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()

    raw, err := json.Marshal(payload)
    if err != nil {
        return Page{}, err
    }

    var reqBody = bytes.NewReader(raw)

    req, err := http.NewRequest("POST", u.String(), reqBody)
    if err != nil {
        return Page{}, err
    }

    req.Header.Set("Content-Type", "application/json")

    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return Page{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return Page{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var data Page
        err := json.Unmarshal(respBody, &data)

        return data, err
    }

    var statusCode = resp.StatusCode
    return Page{}, errors.New(fmt.Sprint("The server returned an unknown status code: ", statusCode))
}




func NewPageTag(httpClient *http.Client, parser *sdkgen.Parser) *PageTag {
	return &PageTag{
        internal: &sdkgen.TagAbstract{
            HttpClient: httpClient,
            Parser: parser,
        },
	}
}
