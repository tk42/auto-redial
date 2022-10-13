package service

import (
	"context"
	"database/sql"

	bufbuild "github.com/tk42/auto-redial/gen/proto/golang/github.com/tk42/auto-redial"
	autoredial "github.com/tk42/auto-redial/gen/sqlc"
	"google.golang.org/protobuf/types/known/durationpb"
)

func NewCallHistoryServiceServer(client *sql.DB) bufbuild.CallHistoryStoreServiceServer {
	db := autoredial.New(client)
	return &ServiceServer{
		db: db,
	}
}

func (s ServiceServer) PutCallHistory(ctx context.Context, req *bufbuild.PutCallHistoryRequest) (*bufbuild.PutCallHistoryResponse, error) {
	params := autoredial.CreateCallHistoryParams{
		ID:        req.Id,
		ScammerID: req.ScammerId,
		CallAt:    ConvertDatetime2Time(req.CallAt),
		CallTime:  int64(req.CallTime.AsDuration().Seconds()),
		Result:    req.Result,
		TalkTime: sql.NullInt64{
			Int64: int64(req.TalkTime.AsDuration().Seconds()),
			Valid: req.TalkTime != nil,
		},
	}
	call, err := s.db.CreateCallHistory(ctx, params)
	if err != nil {
		return nil, err
	}

	matchings, err := s.db.ListMatchingTag(ctx, call.ScammerID)
	if err != nil {
		return nil, err
	}

	var tags []string
	for _, m := range matchings {
		tags = append(tags, m.Tag)
	}

	return &bufbuild.PutCallHistoryResponse{
		Call: &bufbuild.CallHistory{
			Id:        call.ID,
			ScammerId: call.ScammerID,
			CallAt:    ConvertTime2Datetime(call.CallAt),
			CallTime: &durationpb.Duration{
				Seconds: call.CallTime,
			},
			Result: call.Result,
			Tags:   tags,
			TalkTime: &durationpb.Duration{
				Seconds: params.TalkTime.Int64,
			},
		},
	}, nil
}

func (s ServiceServer) DeleteCallHistory(ctx context.Context, req *bufbuild.DeleteCallHistoryRequest) (*bufbuild.DeleteCallHistoryResponse, error) {
	for _, id := range req.GetId() {
		if err := s.db.DeleteCallHistory(ctx, id); err != nil {
			return nil, err
		}
	}
	return &bufbuild.DeleteCallHistoryResponse{}, nil
}

func (s ServiceServer) GetCallHistory(ctx context.Context, req *bufbuild.GetCallHistoryRequest) (*bufbuild.GetCallHistoryResponse, error) {
	var calls []*bufbuild.CallHistory
	for _, id := range req.GetId() {
		call, err := s.db.GetCallHistory(ctx, id)
		if err != nil {
			return nil, err
		}
		calls = append(calls, &bufbuild.CallHistory{
			Id:        call.ID,
			ScammerId: call.ScammerID,
			CallAt:    ConvertTime2Datetime(call.CallAt),
			CallTime: &durationpb.Duration{
				Seconds: call.CallTime,
			},
			Result: call.Result,
			TalkTime: &durationpb.Duration{
				Seconds: call.TalkTime.Int64,
			},
		})
	}
	return &bufbuild.GetCallHistoryResponse{
		Call: calls,
	}, nil
}
