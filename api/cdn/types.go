package cdn

import "time"

type Domain struct {
	ID                string    `json:"id"`
	UserID            string    `json:"user_id"`
	Domain            string    `json:"domain"`
	Name              string    `json:"name"`
	PlanLevel         int       `json:"plan_level"`
	NSKeys            []string  `json:"ns_keys"`
	CurrentNS         []string  `json:"current_ns"`
	TargetCName       string    `json:"target_cname"`
	CustomCName       string    `json:"custom_cname"`
	Type              string    `json:"type"`
	Status            string    `json:"status"`
	DNSCloud          bool      `json:"dns_cloud"`
	Restriction       []string  `json:"restriction"`
	Transfer          Transfer  `json:"transfer"`
	FingerprintStatus bool      `json:"fingerprint_status"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type Transfer struct {
	Domain      string    `json:"domain"`
	AccountID   string    `json:"account_id"`
	AccountName string    `json:"account_name"`
	OwnerName   string    `json:"owner_name"`
	OwnerID     string    `json:"owner_id"`
	Time        time.Time `json:"time"`
	Incoming    bool      `json:"incoming"`
}
