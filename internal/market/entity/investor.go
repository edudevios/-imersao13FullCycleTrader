package entity

type Investor struct {
	ID            string
	Name          string
	Assetposition []*InvestorAssetposition
}

func NewInvestor(id string) *Investor {
	return &Investor{
		ID:            id,
		Assetposition: []*InvestorAssetposition{},
	}
}

// Métodos em Go
func (m *Investor) AddAssetPosition(assetPosition *InvestorAssetposition) {
	m.Assetposition = append(m.Assetposition, assetPosition)
}

func (m *Investor) UpdateAssetPosition(assetID string, qtdShares int) {
	assetPosition := m.GetAssetPosition(assetID)
}

func (m *Investor) GetAssetPosition(assetID string) *InvestorAssetposition {

}

type InvestorAssetposition struct {
	AssetID string
	Shared  int
}
