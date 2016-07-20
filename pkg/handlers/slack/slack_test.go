/*
Copyright 2016 Skippbox, Ltd.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package slack

import (
	"fmt"
	"reflect"
	"testing"

	"k8s.io/kubernetes/pkg/watch"

	"github.com/skippbox/kubewatch/config"
)

func TestSlackInit(t *testing.T) {

	s := &Slack{}
	expectedError := fmt.Errorf(slackErrMsg, "Missing slack token or channel")

	var Tests = []struct {
		c   *config.Config
		err error
	}{
		{&config.Config{SlackToken: "foo", SlackChannel: "bar"}, nil},
		{&config.Config{SlackToken: "foo"}, expectedError},
		{&config.Config{SlackChannel: "foo"}, expectedError},
		{&config.Config{}, expectedError},
	}

	for _, tt := range Tests {
		if err := s.Init(tt.c); !reflect.DeepEqual(err, tt.err) {
			t.Fatalf("Init(): %v", err)
		}
	}
}

func TestSlackHandle(t *testing.T) {

	w := watch.Event{}
	s := &Slack{}

	err := s.Handle(w)

	if err == nil {
		t.Fatalf("Handle(): %v", err)
	}
}
