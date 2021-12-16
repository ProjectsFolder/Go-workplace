package services

import (
    "encoding/json"
    "io"
    "net/http"
    "net/url"
    "strconv"
    dto "workplace/internal/dto/billing"
)

type BillingProviderInterface interface {
    GetContract(id int) (*dto.Contract, error)
}

type BillingClient struct {
    client httpClientInterface
    billingUrl string
    billingUser string
    billingPassword string
}

type httpClientInterface interface {
    Do(req *http.Request) (*http.Response, error)
}

func NewBillingClient(
    httpClient httpClientInterface,
    billingUrl string,
    billingUser string,
    billingPassword string,
) *BillingClient {
    return &BillingClient{
        client: httpClient,
        billingUrl: billingUrl,
        billingUser: billingUser,
        billingPassword: billingPassword,
    }
}

func (billing *BillingClient) GetContract(id int) (*dto.Contract, error) {
    query := url.Values{}
    query.Add("module", "v2.contract")
    query.Add("action", "getDetailedById")
    query.Add("id", strconv.FormatInt(int64(id), 10))

    bytes, err := billing.sendRequest(query)
    if err != nil {
        return nil, err
    }
    var result dto.ContractResponse
    err = json.Unmarshal(bytes, &result)
    if err != nil {
        return nil, err
    }

    err = result.CheckResponse()
    if err != nil {
        return nil, err
    }

    var contract *dto.Contract
    if len(result.Data) > 0 {
        contract = &result.Data[0]
    }

    return contract, nil
}

func (billing *BillingClient) sendRequest(query url.Values) ([]byte, error) {
    request, err := http.NewRequest("GET", billing.billingUrl, nil)
    if err != nil {
        return nil, err
    }

    query.Add("user", billing.billingUser)
    query.Add("pswd", billing.billingPassword)
    query.Add("ct", "json")
    request.URL.RawQuery = query.Encode()

    response, err := billing.client.Do(request)
    if err != nil {
        return nil, err
    }
    
    bodyBytes, err := io.ReadAll(response.Body)
    if err != nil {
        return nil, err
    }

    return bodyBytes, nil
}
