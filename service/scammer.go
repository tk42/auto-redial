package service

import (
	"context"
	"database/sql"

	bufbuild "github.com/tk42/auto-redial/gen/proto/golang/github.com/tk42/auto-redial"
	autoredial "github.com/tk42/auto-redial/gen/sqlc"
)

func NewScammerServiceServer(client *sql.DB) bufbuild.ScammerStoreServiceServer {
	db := autoredial.New(client)
	return &ServiceServer{
		db: db,
	}
}

func (s ServiceServer) PutScammer(ctx context.Context, req *bufbuild.PutScammerRequest) (*bufbuild.PutScammerResponse, error) {
	scammer, err := s.db.CreateScammer(ctx, autoredial.CreateScammerParams{
		ID:       req.Id,
		Name:     req.Name,
		Tel:      req.Tel,
		IsActive: req.IsActive,
	})
	if err != nil {
		return nil, err
	}

	calls, err := s.db.ListScammerCallByScammerId(ctx, scammer.ID)
	if err != nil {
		return nil, err
	}
	var call_ids []string
	for _, c := range calls {
		call_ids = append(call_ids, c.CallID)
	}

	tags, err := s.db.ListScammerTag(ctx, scammer.ID)
	if err != nil {
		return nil, err
	}
	var tag_labels []string
	for _, t := range tags {
		tag_labels = append(tag_labels, t.Tag)
	}

	return &bufbuild.PutScammerResponse{
		Scammer: &bufbuild.Scammer{
			Id:       scammer.ID,
			Name:     scammer.Name,
			Tel:      scammer.Tel,
			IsActive: scammer.IsActive,
			Tags:     tag_labels,
			Calls:    call_ids,
		},
	}, nil
}

func (s ServiceServer) DeleteScammer(ctx context.Context, req *bufbuild.DeleteScammerRequest) (*bufbuild.DeleteScammerResponse, error) {
	for _, id := range req.GetId() {
		if err := s.db.DeleteScammer(ctx, id); err != nil {
			return nil, err
		}
	}
	return &bufbuild.DeleteScammerResponse{}, nil
}

func (s ServiceServer) GetScammer(ctx context.Context, req *bufbuild.GetScammerRequest) (*bufbuild.GetScammerResponse, error) {
	var scammers []*bufbuild.Scammer
	for _, id := range req.GetId() {
		scammer, err := s.db.GetScammer(ctx, id)
		if err != nil {
			return nil, err
		}

		calls, err := s.db.ListScammerCallByScammerId(ctx, scammer.ID)
		if err != nil {
			return nil, err
		}
		var call_ids []string
		for _, c := range calls {
			call_ids = append(call_ids, c.CallID)
		}

		tags, err := s.db.ListScammerTag(ctx, scammer.ID)
		if err != nil {
			return nil, err
		}
		var tag_labels []string
		for _, t := range tags {
			tag_labels = append(tag_labels, t.Tag)
		}

		scammers = append(scammers, &bufbuild.Scammer{
			Id:       scammer.ID,
			Name:     scammer.Name,
			Tel:      scammer.Tel,
			IsActive: scammer.IsActive,
			Tags:     tag_labels,
			Calls:    call_ids,
		})
	}

	return &bufbuild.GetScammerResponse{
		Scammer: scammers,
	}, nil
}
