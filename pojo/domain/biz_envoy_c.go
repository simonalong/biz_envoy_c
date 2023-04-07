package domain

type BizEnvoyC struct {
	Id          float64
	ServiceName string
	Times       string
}

func (BizEnvoyC) TableName() string {
	return "service_update"
}
