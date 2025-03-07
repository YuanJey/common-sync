package third_data

import "common-sync/pkg/http_client"

type ThirdData interface {
	GetThirdDeptList() ([]ThirdDept, error)
	GetThirdUserList() ([]ThirdUser, error)
}
type ThirdApiData struct {
	httpClient *http_client.HttpClient
}

func (t *ThirdApiData) GetThirdDeptList() ([]ThirdDept, error) {
	//TODO implement me
	panic("implement me")
}

func (t *ThirdApiData) GetThirdUserList() ([]ThirdUser, error) {
	//TODO implement me
	panic("implement me")
}
