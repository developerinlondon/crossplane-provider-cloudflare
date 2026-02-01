// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	providerconfig "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/providerconfig"
	bucket "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/r2/bucket"
	bucketcors "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/r2/bucketcors"
	bucketeventnotification "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/r2/bucketeventnotification"
	bucketlifecycle "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/r2/bucketlifecycle"
	bucketlock "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/r2/bucketlock"
	bucketsippy "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/r2/bucketsippy"
	customdomain "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/r2/customdomain"
	manageddomain "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/r2/manageddomain"
	dnssec "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/zone/dnssec"
	setting "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/zone/setting"
	zone "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/zone/zone"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		providerconfig.Setup,
		bucket.Setup,
		bucketcors.Setup,
		bucketeventnotification.Setup,
		bucketlifecycle.Setup,
		bucketlock.Setup,
		bucketsippy.Setup,
		customdomain.Setup,
		manageddomain.Setup,
		dnssec.Setup,
		setting.Setup,
		zone.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		providerconfig.SetupGated,
		bucket.SetupGated,
		bucketcors.SetupGated,
		bucketeventnotification.SetupGated,
		bucketlifecycle.SetupGated,
		bucketlock.SetupGated,
		bucketsippy.SetupGated,
		customdomain.SetupGated,
		manageddomain.SetupGated,
		dnssec.SetupGated,
		setting.SetupGated,
		zone.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
