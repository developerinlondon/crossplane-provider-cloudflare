package providerconfig

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"
)

func Setup(mgr ctrl.Manager, o controller.Options) error {
	return nil
}

func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	return nil
}
