package subscription

import (
	"github.com/NoAnguish/PearlerBackend/backend/utils/api_errors"
	"github.com/NoAnguish/PearlerBackend/backend/utils/database"
	"github.com/NoAnguish/PearlerBackend/backend/utils/formatters"
)

func CreateSubscriptionHandler(request SourceTargetRequest) (*SubscriptionIdResponse, *api_errors.Error) {
	if request.Source == request.Target {
		return nil, api_errors.NewBadRequestError("source and target are the same")
	}

	s, err_ := database.PrepareDefaultWriteSession()
	if err_ != nil {
		return nil, api_errors.NewDatabaseConnectionError(err_)
	}
	defer s.Close()

	subscription := Subscription{
		Id:      formatters.GenerateId(),
		Source:  request.Source,
		Target:  request.Target,
		Deleted: false,
	}

	flag, err := Exists(s, request.Source, request.Target)
	if err != nil {
		return nil, err
	}

	if flag {
		existedSub, err := GetBySourceAndTarget(s, request.Source, request.Target)
		if err != nil {
			return nil, err
		}

		if !existedSub.Deleted {
			return nil, api_errors.NewBadRequestError("subscription already exist")
		}

		subscription.Id = existedSub.Id
		err = Update(s, subscription)
		if err != nil {
			return nil, err
		}
	} else {
		err = Insert(s, subscription)
		if err != nil {
			return nil, err
		}
	}

	return &SubscriptionIdResponse{Id: subscription.Id}, nil
}

func DeleteSubscriptionHandler(request SourceTargetRequest) (*SubscriptionIdResponse, *api_errors.Error) {
	if request.Source == request.Target {
		return nil, api_errors.NewBadRequestError("source and target are the same")
	}

	s, err_ := database.PrepareDefaultWriteSession()
	if err_ != nil {
		return nil, api_errors.NewDatabaseConnectionError(err_)
	}
	defer s.Close()

	flag, err := Exists(s, request.Source, request.Target)
	if err != nil {
		return nil, err
	}

	if !flag {
		return nil, api_errors.NewNotFoundError("subscription does not exist")
	}

	subscription, err := GetBySourceAndTarget(s, request.Source, request.Target)
	if err != nil {
		return nil, err
	}
	if subscription.Deleted {
		return nil, api_errors.NewBadRequestError("subscription has already deleted")
	}

	subscription.Deleted = true
	err = Update(s, *subscription)
	if err != nil {
		return nil, err
	}

	return &SubscriptionIdResponse{Id: subscription.Id}, nil
}

func GetSelfStatsByUserIdHandler(request UserIdRequest) (*SelfStatsResponse, *api_errors.Error) {
	s, err_ := database.PrepareDefaultScanSession()
	if err_ != nil {
		return nil, api_errors.NewDatabaseConnectionError(err_)
	}
	defer s.Close()

	subscribersAmount, err := GetSubsAmountByTargetId(s, request.Id)
	if err != nil {
		return nil, err
	}

	subscriptionsAmount, err := GetSubsAmountBySourceId(s, request.Id)
	if err != nil {
		return nil, err
	}

	response := SelfStatsResponse{
		SubscribersAmount:   *subscribersAmount,
		SubscriptionsAmount: *subscriptionsAmount,
	}

	return &response, nil
}

func GetGeneralStatsByUserIdHandler(request UserIdWithSelfRequest) (*GeneralStatsResponse, *api_errors.Error) {
	s, err_ := database.PrepareDefaultScanSession()
	if err_ != nil {
		return nil, api_errors.NewDatabaseConnectionError(err_)
	}
	defer s.Close()

	subscribersAmount, err := GetSubsAmountByTargetId(s, request.Id)
	if err != nil {
		return nil, err
	}

	subscriptionsAmount, err := GetSubsAmountBySourceId(s, request.Id)
	if err != nil {
		return nil, err
	}

	response := GeneralStatsResponse{
		SubscribersAmount:   *subscribersAmount,
		SubscriptionsAmount: *subscriptionsAmount,
	}

	flag, err := Exists(s, request.SelfId, request.Id)
	if err != nil {
		return nil, err
	}

	if !flag {
		response.Subscribed = false
	} else {
		sub, err := GetBySourceAndTarget(s, request.SelfId, request.Id)
		if err != nil {
			return nil, err
		}

		response.Subscribed = !sub.Deleted
	}

	return &response, nil
}
