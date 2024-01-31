package service

import (
	"context"

	v1 "real-kratosworld/api/kratosworld/v1"
	"real-kratosworld/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type KratosWorldService struct {
	v1.UnimplementedKratosWorldServer

	uc  *biz.GreeterUsecase
	log *log.Helper
}

func NewKratosWorldService(uc *biz.GreeterUsecase, logger log.Logger) *KratosWorldService {
	return &KratosWorldService{uc: uc, log: log.NewHelper(logger)}
}

func (s *KratosWorldService) Login(ctx context.Context, req *v1.LoginRequest) (*v1.UserReply, error) {
	return &v1.UserReply{
		User: &v1.UserReply_User{
			Username: "boom",
		},
	}, nil
}

func (s *KratosWorldService) Register(ctx context.Context, req *v1.RegisterRequest) (*v1.UserReply, error) {
	return &v1.UserReply{
		User: &v1.UserReply_User{
			Username: "boom",
		},
	}, nil
}

func (s *KratosWorldService) GetCurrentUser(ctx context.Context, req *v1.GetCurrentUserRequest) (*v1.UserReply, error) {
	return &v1.UserReply{
		User: &v1.UserReply_User{
			Username: "boom",
		},
	}, nil
}

func (s *KratosWorldService) UpdateUser(ctx context.Context, req *v1.UpdateUserRequest) (*v1.UserReply, error) {
	return &v1.UserReply{}, nil
}

func (s *KratosWorldService) GetProfile(ctx context.Context, req *v1.GetProfileRequest) (*v1.ProfileReply, error) {
	return &v1.ProfileReply{}, nil
}

func (s *KratosWorldService) FollowUser(ctx context.Context, req *v1.FollowUserRequest) (*v1.ProfileReply, error) {
	return &v1.ProfileReply{}, nil
}

func (s *KratosWorldService) UnfollowUser(ctx context.Context, req *v1.UnfollowUserRequest) (*v1.ProfileReply, error) {
	return &v1.ProfileReply{}, nil
}

func (s *KratosWorldService) ListArticles(ctx context.Context, req *v1.ListArticlesRequest) (*v1.MultipleAriclesReply, error) {
	return &v1.MultipleAriclesReply{}, nil
}

func (s *KratosWorldService) FeedArticles(ctx context.Context, req *v1.FeedArticlesRequest) (*v1.MultipleAriclesReply, error) {
	return &v1.MultipleAriclesReply{}, nil
}

func (s *KratosWorldService) GetArticles(ctx context.Context, req *v1.GetArticlesRequest) (*v1.SingleAricleReply, error) {
	return &v1.SingleAricleReply{}, nil
}

func (s *KratosWorldService) CreateArticle(ctx context.Context, req *v1.CreateArticleRequest) (*v1.SingleAricleReply, error) {
	return &v1.SingleAricleReply{}, nil
}

func (s *KratosWorldService) UpdateArticle(ctx context.Context, req *v1.UpdateArticleRequest) (*v1.SingleAricleReply, error) {
	return &v1.SingleAricleReply{}, nil
}

func (s *KratosWorldService) DeleteArticle(ctx context.Context, req *v1.DeleteArticleRequest) (*v1.SingleAricleReply, error) {
	return &v1.SingleAricleReply{}, nil
}

func (s *KratosWorldService) AddComment(ctx context.Context, req *v1.AddCommentRequest) (*v1.SingleCommentReply, error) {
	return &v1.SingleCommentReply{}, nil
}

func (s *KratosWorldService) GetComments(ctx context.Context, req *v1.AddCommentRequest) (*v1.MultipleCommentsReply, error) {
	return &v1.MultipleCommentsReply{}, nil
}

func (s *KratosWorldService) DeleteComment(ctx context.Context, req *v1.DeleteCommentRequest) (*v1.SingleCommentReply, error) {
	return &v1.SingleCommentReply{}, nil
}

func (s *KratosWorldService) FavoriteArticle(ctx context.Context, req *v1.FavoriteArticleRequest) (*v1.SingleAricleReply, error) {
	return &v1.SingleAricleReply{}, nil
}

func (s *KratosWorldService) UnfavoriteArticle(ctx context.Context, req *v1.UnfavoriteArticleRequest) (*v1.SingleAricleReply, error) {
	return &v1.SingleAricleReply{}, nil
}

func (s *KratosWorldService) GetTags(ctx context.Context, req *v1.GetTagsRequest) (*v1.TagListReply, error) {
	return &v1.TagListReply{}, nil
}
