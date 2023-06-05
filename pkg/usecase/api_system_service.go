package usecase

import (
	"context"

	"github.com/traPtitech/neoshowcase/pkg/domain"
)

func (s *APIServerService) GetSystemPublicKey(_ context.Context) string {
	encoded := domain.Base64EncodedPublicKey(s.pubKey)
	return encoded + " NeoShowcase"
}

func (s *APIServerService) GetAvailableDomains(ctx context.Context) (domain.AvailableDomainSlice, error) {
	return s.controller.GetAvailableDomains(ctx)
}

func (s *APIServerService) GetAvailablePorts(ctx context.Context) (domain.AvailablePortSlice, error) {
	available, err := s.controller.GetAvailablePorts(ctx)
	if err != nil {
		return nil, err
	}
	return available, nil
}
