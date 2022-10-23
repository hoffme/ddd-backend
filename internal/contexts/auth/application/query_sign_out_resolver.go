package application

import (
	"context"

	busDomain "github.com/hoffme/ddd-backend/internal/shared/bus/domain"

	"github.com/hoffme/ddd-backend/internal/contexts/auth/domain"
)

func (s *Service) CommandSignOut(ctx context.Context, command busDomain.Command[domain.CommandSignOutData]) error {
	refreshClaimsToken, err := s.ports.JWT.DecodeRefreshToken(ctx, command.Data.RefreshToken)
	if err != nil {
		return err
	}

	err = refreshClaimsToken.Verify()
	if err != nil {
		return err
	}

	event := domain.EventSignOutDefinition.CreateEvent(domain.EventSignOutData{
		CommandId:    command.ID,
		RefreshToken: command.Data.RefreshToken,
	})

	s.ports.EventBusEmitter.Emit(ctx, event)

	return nil
}
