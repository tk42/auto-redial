package service

import (
	"context"
	"database/sql"

	bufbuild "github.com/tk42/auto-redial/gen/proto/golang/github.com/tk42/auto-redial"
	sqlc "github.com/tk42/auto-redial/gen/sqlc"
	"google.golang.org/protobuf/types/known/durationpb"
)

func NewMatchingServiceServer(client *sql.DB) bufbuild.MatchingStoreServiceServer {
	db := sqlc.New(client)
	return &ServiceServer{
		db: db,
	}
}

func (s ServiceServer) PutMatching(ctx context.Context, req *bufbuild.PutMatchingRequest) (*bufbuild.PutMatchingResponse, error) {
	matching, err := s.db.CreateMatching(ctx, sqlc.CreateMatchingParams{
		ID:           req.Id,
		CreatedAt:    ConvertDatetime2Time(req.CreatedAt),
		SerialNumber: req.SerialNumber,
		Matched:      req.Matched,
		Checked:      req.Checked,
		MatchingAt: sql.NullTime{
			Time:  ConvertDatetime2Time(req.MatchingAt),
			Valid: req.MatchingAt != nil,
		},
		TalkTime: sql.NullInt64{
			Int64: int64(req.TalkTime.AsDuration().Seconds()),
			Valid: req.TalkTime != nil,
		},
		Transcript: sql.NullString{
			String: *req.Transcript,
			Valid:  req.Transcript != nil,
		},
	})
	if err != nil {
		return nil, err
	}

	for _, id := range req.ScammerId {
		s.db.CreateScammerMatching(ctx, sqlc.CreateScammerMatchingParams{
			ScammerID:  id,
			MatchingID: matching.ID,
		})
	}

	return &bufbuild.PutMatchingResponse{
		Matching: &bufbuild.Matching{
			Id:           matching.ID,
			ScammerId:    req.ScammerId,
			CallId:       req.CallId,
			CreatedAt:    ConvertTime2Datetime(matching.CreatedAt),
			SerialNumber: matching.SerialNumber,
			Matched:      matching.Matched,
			Checked:      matching.Checked,
			MatchingAt:   ConvertTime2Datetime(matching.MatchingAt.Time),
			TalkTime: &durationpb.Duration{
				Seconds: matching.TalkTime.Int64,
			},
			Transcript: &matching.Transcript.String,
		},
	}, nil
}

func (s ServiceServer) GetMatching(ctx context.Context, req *bufbuild.GetMatchingRequest) (*bufbuild.GetMatchingResponse, error) {
	matching, err := s.db.GetMatching(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	scammermatchings, err := s.db.ListScammerMatchingByMatchingId(ctx, matching.ID)
	if err != nil {
		return nil, err
	}
	var scammer_ids, call_ids []string
	for _, scammermatching := range scammermatchings {
		scammer_ids = append(scammer_ids, scammermatching.ScammerID)
		calls, err := s.db.ListCallHistoryByScammerId(ctx, scammermatching.ScammerID)
		if err != nil {
			return nil, err
		}
		for _, call := range calls {
			call_ids = append(call_ids, call.ID)
		}
	}

	return &bufbuild.GetMatchingResponse{
		Matching: &bufbuild.Matching{
			Id:           matching.ID,
			ScammerId:    scammer_ids,
			CallId:       call_ids,
			CreatedAt:    ConvertTime2Datetime(matching.CreatedAt),
			SerialNumber: matching.SerialNumber,
			Matched:      matching.Matched,
			Checked:      matching.Checked,
			MatchingAt:   ConvertTime2Datetime(matching.MatchingAt.Time),
			TalkTime: &durationpb.Duration{
				Seconds: matching.TalkTime.Int64,
			},
			Transcript: &matching.Transcript.String,
		},
	}, nil
}

func (s ServiceServer) DeleteMatching(ctx context.Context, req *bufbuild.DeleteMatchingRequest) (*bufbuild.DeleteMatchingResponse, error) {
	err := s.db.DeleteMatching(ctx, req.MatchingId)
	if err != nil {
		return nil, err
	}

	return &bufbuild.DeleteMatchingResponse{}, nil
}
