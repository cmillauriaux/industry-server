package model

type Offer struct {
	OfferID    string
	ClientID   string
	ProviderID string
	ProductID  string
	Quantity   int
	Price      int
	Quality    int
}
