package services

import (
	"context"
	"github.com/elsyarif/pms-api/internal/domain/entities"
	"github.com/elsyarif/pms-api/internal/domain/repository"
	"github.com/elsyarif/pms-api/pkg/uid"
)

type GroupService struct {
	groupRepo   repository.GroupRepository
	idGenerator uid.NanoGenerator
}

func NewGroupService(gr repository.GroupRepository, uid uid.NanoGenerator) GroupService {
	return GroupService{
		groupRepo:   gr,
		idGenerator: uid,
	}
}

func (s *GroupService) CreateGroup(ctx context.Context, group entities.Group) (*entities.Group, error) {
	group.Id = s.idGenerator.NanoId("group")
	err := s.groupRepo.AddGroup(ctx, group)
	if err != nil {
		return nil, err
	}

	return &group, nil
}

func (s *GroupService) GetAllGroup(ctx context.Context, name string) (*[]entities.Group, error) {
	return s.groupRepo.GetAllGroup(ctx, name)
}

func (s *GroupService) GetGroupById(ctx context.Context, groupId string) (*entities.Group, error) {
	return s.groupRepo.GetGroupById(ctx, groupId)
}

func (s *GroupService) EditGroup(ctx context.Context, groupId string, group entities.Group) error {
	_, err := s.groupRepo.GetGroupById(ctx, groupId)
	if err != nil {
		return err
	}

	return s.groupRepo.EditGroup(ctx, groupId, group)
}

func (s *GroupService) DeleteGroup(ctx context.Context, groupId string) error {
	_, err := s.groupRepo.GetGroupById(ctx, groupId)
	if err != nil {
		return err
	}

	return s.groupRepo.DeleteGroup(ctx, groupId)
}
