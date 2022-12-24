package domain

type Deal struct {
	Id                int     `json:"id"`
	NeedId            int     `json:"needId"`
	OfferId           int     `json:"offerId"`
	CompanyCommission float32 `json:"companyCommission"`
	RealtorCommission float32 `json:"realtorCommission"`
	ClientCommission  float32 `json:"clientCommission"`
}
