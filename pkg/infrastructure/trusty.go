// Copyright 2019 Red Hat, Inc. and/or its affiliates
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package infrastructure

import (
	"github.com/kiegroup/kogito-cloud-operator/pkg/apis/app/v1alpha1"
	"github.com/kiegroup/kogito-cloud-operator/pkg/client"
)

const (
	// DefaultTrustyImageName is just the image name for the Trusty Service
	DefaultTrustyImageName = "kogito-trusty"
	// DefaultTrustyName is the default name for the Trusty instance service
	DefaultTrustyName = "trusty"

	trustyHTTPRouteEnv = "KOGITO_TRUSTY_HTTP_URL"
	trustyWSRouteEnv   = "KOGITO_TRUSTY_WS_URL"
)

// InjectTrustyURLIntoKogitoApps will query for every KogitoRuntime in the given namespace to inject the Trusty route to each one
// Won't trigger an update if the KogitoRuntime already has the route set to avoid unnecessary reconciliation triggers
func InjectTrustyURLIntoKogitoApps(client *client.Client, namespace string) error {
	log.Debugf("Injecting Trusty Route in kogito apps")
	return injectURLIntoKogitoApps(client, namespace, trustyHTTPRouteEnv, trustyWSRouteEnv, &v1alpha1.KogitoTrustyList{})
}
