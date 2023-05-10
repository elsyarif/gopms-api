package usecases

import (
	"context"
	"github.com/elsyarif/pms-api/internal/domain/entities"
	"github.com/elsyarif/pms-api/internal/domain/services"
)

type GroupUseCase struct {
	groupService services.GroupService
}

func NewGroupUseCase(gu services.GroupService) GroupUseCase {
	return GroupUseCase{
		groupService: gu,
	}
}

func (u *GroupUseCase) AddGroup(ctx context.Context, group entities.Group) (*entities.Group, error) {
	return u.groupService.CreateGroup(ctx, group)
}

func (u *GroupUseCase) GetAllGroup(ctx context.Context, name string) (*[]entities.Group, error) {
	return u.groupService.GetAllGroup(ctx, name)
}

func (u *GroupUseCase) GetGroupById(ctx context.Context, groupId string) (*entities.Group, error) {
	return u.groupService.GetGroupById(ctx, groupId)
}

func (u *GroupUseCase) EditGroup(ctx context.Context, groupId string, group entities.Group) error {
	return u.groupService.EditGroup(ctx, groupId, group)
}

func (u *GroupUseCase) DeleteGroup(ctx context.Context, groupId string) error {
	return u.groupService.DeleteGroup(ctx, groupId)
}
