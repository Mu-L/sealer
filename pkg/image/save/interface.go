// Copyright © 2021 Alibaba Group Holding Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package save

import (
	"context"

	"github.com/docker/docker/pkg/progress"
	v1 "github.com/sealerio/sealer/types/api/v1"
)

// ImageSave can save a list of images of the specified platform
type ImageSave interface {
	// SaveImages is not concurrently safe
	SaveImages(images []string, dir string, platform v1.Platform) error
}

type Section struct {
	Registry string             `json:"registry,omitempty"`
	Username string             `json:"username,omitempty"`
	Password string             `json:"password,omitempty"`
	Images   map[string][]Named `json:"images,omitempty"`
}

type ImageListWithAuth []Section

type DefaultImageSaver struct {
	ctx            context.Context
	domainToImages map[string][]Named
	progressOut    progress.Output
}

func NewImageSaver(ctx context.Context) ImageSave {
	if ctx == nil {
		ctx = context.Background()
	}
	return &DefaultImageSaver{
		ctx:            ctx,
		domainToImages: make(map[string][]Named),
	}
}
