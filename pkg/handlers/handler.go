/*
 * Copyright (c) 2024. Devtron Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package handlers

import (
	"github.com/devtron-labs/kubewatch/config"
	"github.com/devtron-labs/kubewatch/pkg/handlers/flock"
	"github.com/devtron-labs/kubewatch/pkg/handlers/hipchat"
	"github.com/devtron-labs/kubewatch/pkg/handlers/mattermost"
	"github.com/devtron-labs/kubewatch/pkg/handlers/slack"
	"github.com/devtron-labs/kubewatch/pkg/handlers/webhook"
)

// Handler is implemented by any handler.
// The Handle method is used to process event
type Handler interface {
	Init(c *config.Config) error
	ObjectCreated(obj interface{})
	ObjectDeleted(obj interface{})
	ObjectUpdated(oldObj, newObj interface{})
}

// Map maps each event handler function to a name for easily lookup
var Map = map[string]interface{}{
	"default":    &Default{},
	"slack":      &slack.Slack{},
	"hipchat":    &hipchat.Hipchat{},
	"mattermost": &mattermost.Mattermost{},
	"flock":      &flock.Flock{},
	"webhook":    &webhook.Webhook{},
}

// Default handler implements Handler interface,
// print each event with JSON format
type Default struct {
}

// Init initializes handler configuration
// Do nothing for default handler
func (d *Default) Init(c *config.Config) error {
	return nil
}

func (d *Default) ObjectCreated(obj interface{}) {
	//e := event.New(obj, "created")
	//fmt.Printf("Created: %+v\n", e)
}

func (d *Default) ObjectDeleted(obj interface{}) {

}

func (d *Default) ObjectUpdated(oldObj, newObj interface{}) {
	//e := event.New(oldObj, "updated")
	//fmt.Printf("Updated: %+v\n", e)
}
