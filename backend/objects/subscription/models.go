package subscription

type Subscription struct {
	Id      string
	Source  string
	Target  string
	Deleted bool
}

// TODO(noanguish) change it to source_id and target_id
type SourceTargetRequest struct {
	Source string `json:"source_id,omitempty"`
	Target string `json:"target_id,omitempty"`
}

type UserIdRequest struct {
	Id string `json:"id,omitempty"`
}

type UserIdWithSelfRequest struct {
	Id     string `json:"id,omitempty"`
	SelfId string `json:"self_id,omitempty"`
}

type SubscriptionIdResponse struct {
	Id string `json:"id"`
}

type UserIdsResponse struct {
	Ids []string `json:"ids"`
}

type SelfStatsResponse struct {
	SubscribersAmount   int32 `json:"subscribers_amount"`
	SubscriptionsAmount int32 `json:"subscriptions_amount"`
}

type GeneralStatsResponse struct {
	SubscribersAmount   int32 `json:"subscribers_amount"`
	SubscriptionsAmount int32 `json:"subscriptions_amount"`
	Subscribed          bool  `json:"is_subscribed"`
}
