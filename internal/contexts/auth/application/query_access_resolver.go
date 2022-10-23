package application

import (
	"context"
	"time"

	busDomain "github.com/hoffme/ddd-backend/internal/shared/bus/domain"

	"github.com/hoffme/ddd-backend/internal/contexts/auth/domain"
)

func (s *Service) CommandAccess(ctx context.Context, command busDomain.Command[domain.CommandAccessData]) error {
	refreshClaimsToken, err := s.ports.JWT.DecodeRefreshToken(ctx, command.Data.RefreshToken)
	if err != nil {
		return err
	}

	err = refreshClaimsToken.Verify()
	if err != nil {
		return err
	}

	expireAt := time.Now().Add(s.configs.AccessTokenExpiration).Unix()

	accessToken, err := s.ports.JWT.EncodeAccessToken(ctx, domain.JWTAccessTokenClaims{
		UserID:   refreshClaimsToken.UserID,
		ExpireAt: expireAt,
	})
	if err != nil {
		return err
	}

	event := domain.EventAccessDefinition.CreateEvent(domain.EventAccessData{
		CommandId:   command.ID,
		AccessToken: accessToken,
		ExpiredAt:   expireAt,
	})

	s.ports.EventBusEmitter.Emit(ctx, event)

	return nil
}
