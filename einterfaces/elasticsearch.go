// Copyright (c) 2017-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package einterfaces

import (
	"time"

	"github.com/mattermost/mattermost-server/model"
)

type ElasticsearchInterface interface {
	Start() *model.AppError
	IndexPost(post *model.Post, teamId string) *model.AppError
	SearchPosts(channels *model.ChannelList, searchParams []*model.SearchParams) ([]string, *model.AppError)
	DeletePost(post *model.Post) *model.AppError
	TestConfig(cfg *model.Config) *model.AppError
	PurgeIndexes() *model.AppError
	DataRetentionDeleteIndexes(cutoff time.Time) *model.AppError
}

var theElasticsearchInterface ElasticsearchInterface

func RegisterElasticsearchInterface(newInterface ElasticsearchInterface) {
	theElasticsearchInterface = newInterface
}

func GetElasticsearchInterface() ElasticsearchInterface {
	return theElasticsearchInterface
}
