// Copyright © 2022 Dell Inc. or its subsidiaries. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//      http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package drivers

import (
	"context"
	"fmt"
	"testing"

	csmv1 "github.com/dell/csm-operator/api/v1"
	"github.com/dell/csm-operator/tests/shared"
	"github.com/dell/csm-operator/tests/shared/crclient"
	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	powerScaleCSM            = csmForPowerScale()
	powerScaleCSMEmptyEnv    = csmForPowerScaleWithEmptyEnv()
	powerScaleCSMBadSkipCert = csmForPowerScaleBadSkipCert()
	powerScaleCSMBadCertCnt  = csmForPowerScaleBadCertCnt()
	powerScaleCSMBadVersion  = csmForPowerScaleBadVersion()
	objects                  = map[shared.StorageKey]runtime.Object{}
	powerScaleClient         = crclient.NewFakeClientNoInjector(objects)
	powerScaleSecret         = shared.MakeSecret("csm-creds", "driver-test", shared.ConfigVersion)

	powerScaleTests = []struct {
		// every single unit test name
		name string
		// csm object
		csm csmv1.ContainerStorageModule
		// client
		ct  client.Client
		sec *corev1.Secret
		// expected error
		expectedErr string
	}{
		{"happy path", powerScaleCSM, powerScaleClient, powerScaleSecret, ""},
		{"invalid value for skip cert validation", powerScaleCSMBadSkipCert, powerScaleClient, powerScaleSecret, "is an invalid value for X_CSI_ISI_SKIP_CERTIFICATE_VALIDATION"},
		{"invalid value for cert secret cnt", powerScaleCSMBadCertCnt, powerScaleClient, powerScaleSecret, "is an invalid value for CERT_SECRET_COUNT"},
	}

	preCheckPowerScaleTest = []struct {
		// every single unit test name
		name string
		// csm object
		csm csmv1.ContainerStorageModule
		// client
		ct client.Client
		// secret
		sec *corev1.Secret
		// expected error
		expectedErr string
	}{
		{"missing secret", powerScaleCSM, powerScaleClient, powerScaleSecret, "failed to find secret"},
		{"bad version", powerScaleCSMBadVersion, powerScaleClient, powerScaleSecret, "not supported"},
		{"missing envs", powerScaleCSMEmptyEnv, powerScaleClient, powerScaleSecret, "failed to find secret"},
	}

	powerScaleCommonEnvTest = []struct {
		name       string
		yamlString string
		csm        csmv1.ContainerStorageModule
		ct         client.Client
		sec        *corev1.Secret
		fileType   string
		expected   string
	}{
		{
			name:       "update GOISILON_DEBUG value for Controller",
			yamlString: "<GOISILON_DEBUG>",
			csm:        goisilonDebug("true"),
			ct:         powerScaleClient,
			sec:        powerScaleSecret,
			fileType:   "Controller",
			expected:   "true",
		},
		{
			name:       "update GOISILON_DEBUG value for Node",
			yamlString: "<GOISILON_DEBUG>",
			csm:        goisilonDebug("true"),
			ct:         powerScaleClient,
			sec:        powerScaleSecret,
			fileType:   "Node",
			expected:   "true",
		},
	}
)

func TestGetApplyCertVolume(t *testing.T) {
	for _, tt := range powerScaleTests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := getApplyCertVolume(tt.csm)
			if tt.expectedErr == "" {
				assert.Nil(t, err)
			} else {
				assert.Containsf(t, err.Error(), tt.expectedErr, "expected error containing %q, got %s", tt.expectedErr, err)
			}
		})
	}
}

func TestPrecheckPowerScale(t *testing.T) {
	ctx := context.Background()
	for _, tt := range preCheckPowerScaleTest {
		t.Run(tt.name, func(t *testing.T) { // #nosec G601 - Run waits for the call to complete.
			err := PrecheckPowerScale(ctx, &tt.csm, config, tt.ct)
			if tt.expectedErr == "" {
				assert.Nil(t, err)
			} else {
				assert.Containsf(t, err.Error(), tt.expectedErr, "expected error containing %q, got %s", tt.expectedErr, err)
			}
		})
	}

	// grab the first secret

	for _, tt := range powerScaleTests {
		// create secret for each run
		err := tt.ct.Create(ctx, tt.sec)
		if err != nil {
			assert.Nil(t, err)
		}

		t.Run(tt.name, func(t *testing.T) { // #nosec G601 - Run waits for the call to complete.
			err := PrecheckPowerScale(ctx, &tt.csm, config, tt.ct)
			if tt.expectedErr == "" {
				assert.Nil(t, err)
			} else {
				fmt.Printf("err: %+v\n", err)
				assert.Containsf(t, err.Error(), tt.expectedErr, "expected error containing %q, got %s", tt.expectedErr, err)
			}
		})

		// remove secret after each run
		err = tt.ct.Delete(ctx, tt.sec)
		if err != nil {
			assert.Nil(t, err)
		}
	}
}

func TestModifyPowerScaleCR(t *testing.T) {
	for _, tt := range powerScaleCommonEnvTest {
		t.Run(tt.name, func(t *testing.T) {
			result := ModifyPowerScaleCR(tt.yamlString, tt.csm, tt.fileType)
			if result != tt.expected {
				t.Errorf("expected %v, but got %v", tt.expected, result)
			}
		})
	}
}

// makes a csm object with tolerations
func csmForPowerScale() csmv1.ContainerStorageModule {
	res := shared.MakeCSM("csm", "driver-test", shared.ConfigVersion)

	// Add log level to cover some code in GetConfigMap
	envVarLogLevel1 := corev1.EnvVar{Name: "CERT_SECRET_COUNT", Value: "0"}
	envVarLogLevel2 := corev1.EnvVar{Name: "X_CSI_ISI_SKIP_CERTIFICATE_VALIDATION", Value: "false"}
	envVarLogLevel3 := corev1.EnvVar{Name: "GOISILON_DEBUG", Value: "false"}
	res.Spec.Driver.Common.Envs = []corev1.EnvVar{envVarLogLevel1, envVarLogLevel2, envVarLogLevel3}
	res.Spec.Driver.AuthSecret = "csm-creds"

	// Add pscale driver version
	res.Spec.Driver.ConfigVersion = shared.ConfigVersion
	res.Spec.Driver.CSIDriverType = csmv1.PowerScale

	return res
}

func csmForPowerScaleWithEmptyEnv() csmv1.ContainerStorageModule {
	res := shared.MakeCSM("csm", "driver-test", shared.ConfigVersion)

	res.Spec.Driver.Common.Envs = []corev1.EnvVar{}
	res.Spec.Driver.AuthSecret = "csm-creds"

	// Add pscale driver version
	res.Spec.Driver.ConfigVersion = shared.ConfigVersion
	res.Spec.Driver.CSIDriverType = csmv1.PowerScale

	return res
}

// makes a csm object with tolerations
func csmForPowerScaleBadSkipCert() csmv1.ContainerStorageModule {
	res := shared.MakeCSM("csm", "driver-test", shared.ConfigVersion)

	// Add log level to cover some code in GetConfigMap
	envVarLogLevel1 := corev1.EnvVar{Name: "CERT_SECRET_COUNT", Value: "2"}
	envVarLogLevel2 := corev1.EnvVar{Name: "X_CSI_ISI_SKIP_CERTIFICATE_VALIDATION", Value: "NotABool"}
	res.Spec.Driver.Common.Envs = []corev1.EnvVar{envVarLogLevel1, envVarLogLevel2}

	// Add pscale driver version
	res.Spec.Driver.ConfigVersion = shared.ConfigVersion
	res.Spec.Driver.CSIDriverType = csmv1.PowerScale

	return res
}

// makes a csm object with tolerations
func csmForPowerScaleBadCertCnt() csmv1.ContainerStorageModule {
	res := shared.MakeCSM("csm", "driver-test", shared.ConfigVersion)

	// Add log level to cover some code in GetConfigMap
	envVarLogLevel1 := corev1.EnvVar{Name: "CERT_SECRET_COUNT", Value: "thisIsNotANumber"}
	envVarLogLevel2 := corev1.EnvVar{Name: "X_CSI_ISI_SKIP_CERTIFICATE_VALIDATION", Value: "true"}
	res.Spec.Driver.Common.Envs = []corev1.EnvVar{envVarLogLevel1, envVarLogLevel2}

	// Add pscale driver version
	res.Spec.Driver.ConfigVersion = shared.ConfigVersion
	res.Spec.Driver.CSIDriverType = csmv1.PowerScale

	return res
}

// makes a csm object with tolerations
func csmForPowerScaleBadVersion() csmv1.ContainerStorageModule {
	res := shared.MakeCSM("csm", "driver-test", shared.ConfigVersion)

	// Add pscale driver version
	res.Spec.Driver.ConfigVersion = "v0"
	res.Spec.Driver.CSIDriverType = csmv1.PowerScale

	return res
}

func goisilonDebug(debug string) csmv1.ContainerStorageModule {
	cr := csmForPowerScale()
	cr.Spec.Driver.Common.Envs = []corev1.EnvVar{
		{Name: "GOISILON_DEBUG", Value: debug},
	}

	return cr
}
