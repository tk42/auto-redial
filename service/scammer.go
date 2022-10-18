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

	var tag_labels []string
	for _, tag := range req.Tags {
		_, err := s.db.CreateScammerTag(ctx, autoredial.CreateScammerTagParams{
			ScammerID: scammer.ID,
			Tag:       tag,
		})
		if err != nil {
			return nil, err
		}
		tag_labels = append(tag_labels, tag)
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

func (s ServiceServer) UpdateTag(ctx context.Context, req *bufbuild.UpdateTagRequest) (*bufbuild.UpdateTagResponse, error) {
	s.db.CreateScammerTag(ctx, autoredial.CreateScammerTagParams{
		ScammerID: req.GetId(),
		Tag:       req.GetTag(),
	})
	tags, err := s.db.ListScammerTag(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	var tag_labels []string
	for _, tag := range tags {
		tag_labels = append(tag_labels, tag.Tag)
	}
	return &bufbuild.UpdateTagResponse{
		Tags: tag_labels,
	}, nil
}

func (s ServiceServer) UpdateCall(ctx context.Context, req *bufbuild.UpdateCallRequest) (*bufbuild.UpdateCallResponse, error) {
	s.db.CreateScammerCall(ctx, autoredial.CreateScammerCallParams{
		ScammerID: req.GetId(),
		CallID:    req.GetCall(),
	})
	calls, err := s.db.ListScammerCallByScammerId(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	var call_labels []string
	for _, call := range calls {
		call_labels = append(call_labels, call.CallID)
	}
	return &bufbuild.UpdateCallResponse{
		Calls: call_labels,
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
