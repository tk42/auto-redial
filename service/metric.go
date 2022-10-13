package service

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	bufbuild "github.com/tk42/auto-redial/gen/proto/golang/github.com/tk42/auto-redial"
	autoredial "github.com/tk42/auto-redial/gen/sqlc"
	"google.golang.org/protobuf/types/known/durationpb"
)

func NewMetricServiceServer(client *sql.DB) bufbuild.MetricStoreServiceServer {
	db := autoredial.New(client)
	return &ServiceServer{
		db: db,
	}
}

func (s ServiceServer) PutMetric(ctx context.Context, req *bufbuild.PutMetricRequest) (*bufbuild.PutMetricResponse, error) {
	metric, err := s.db.CreateMetric(ctx, autoredial.CreateMetricParams{
		ID:        uuid.New().String(),
		CreatedAt: ConvertDatetime2Time(req.CreatedAt),
		Calls:     req.Calls,
		Scammers:  req.Scammers,
		Inactives: req.Inactives,
		CallTime:  int64(req.CallTime.AsDuration().Seconds()),
		TalkTime:  int64(req.TalkTime.AsDuration().Seconds()),
		Amount:    req.Amount,
	})
	if err != nil {
		return nil, err
	}
	return &bufbuild.PutMetricResponse{
		Metric: &bufbuild.Metric{
			CreatedAt: ConvertTime2Datetime(metric.CreatedAt),
			Calls:     metric.Calls,
			Scammers:  metric.Scammers,
			Inactives: metric.Inactives,
			CallTime: &durationpb.Duration{
				Seconds: metric.CallTime,
			},
			TalkTime: &durationpb.Duration{
				Seconds: metric.TalkTime,
			},
			Amount: metric.Amount,
		},
	}, nil
}

func (s ServiceServer) DeleteMetric(ctx context.Context, req *bufbuild.DeleteMetricRequest) (*bufbuild.DeleteMetricResponse, error) {
	if err := s.db.DeleteMetric(ctx, autoredial.DeleteMetricParams{
		CreatedAt:   ConvertDatetime2Time(req.GetFrom()),
		CreatedAt_2: ConvertDatetime2Time(req.GetTo()),
	}); err != nil {
		return nil, err
	}
	return &bufbuild.DeleteMetricResponse{}, nil
}

func (s ServiceServer) GetMetric(ctx context.Context, req *bufbuild.GetMetricRequest) (*bufbuild.GetMetricResponse, error) {
	metrics, err := s.db.ListMetric(ctx, autoredial.ListMetricParams{
		CreatedAt:   ConvertDatetime2Time(req.GetFrom()),
		CreatedAt_2: ConvertDatetime2Time(req.GetTo()),
	})
	if err != nil {
		return nil, err
	}
	res := make([]*bufbuild.Metric, len(metrics))
	for i, metric := range metrics {
		res[i] = &bufbuild.Metric{
			CreatedAt: ConvertTime2Datetime(metric.CreatedAt),
			Calls:     metric.Calls,
			Scammers:  metric.Scammers,
			Inactives: metric.Inactives,
			CallTime: &durationpb.Duration{
				Seconds: metric.CallTime,
			},
			TalkTime: &durationpb.Duration{
				Seconds: metric.TalkTime,
			},
			Amount: metric.Amount,
		}
	}
	return &bufbuild.GetMetricResponse{
		Metrics: res,
	}, nil
}
