/*
Copyright 2020 The Kubermatic Kubernetes Platform contributors.

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

package common

import (
	"errors"

	kubermaticerrors "k8c.io/kubermatic/v2/pkg/util/errors"

	kerrors "k8s.io/apimachinery/pkg/api/errors"
)

// kubernetesErrorToHTTPError constructs HTTPError only if the given err is of type *StatusError.
// Otherwise unmodified err will be returned to the caller.
func KubernetesErrorToHTTPError(err error) error {
	var errStatus *kerrors.StatusError

	if errors.As(err, &errStatus) {
		httpCode := errStatus.Status().Code
		httpMessage := errStatus.Status().Message

		return kubermaticerrors.New(int(httpCode), httpMessage)
	}

	return err
}
