package controller

import (
	"context"
	"net/url"

	infrav1 "github.com/ionos-cloud/cluster-api-provider-proxmox/api/v1alpha1"
	"github.com/ionos-cloud/cluster-api-provider-proxmox/pkg/proxmox"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// ProxmoxImportReconciler reconciles a ProxmoxImport object
type ProxmoxImportReconciler struct {
	client.Client
	Scheme        *runtime.Scheme
	Recorder      record.EventRecorder
	ProxmoxClient proxmox.Client
}

// SetupWithManager sets up the controller with the Manager.
func (r *ProxmoxImportReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&infrav1.ProxmoxImport{}).
		Complete(r)
}

//+kubebuilder:rbac:groups=infrastructure.cluster.x-k8s.io,resources=proxmoximports,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=infrastructure.cluster.x-k8s.io,resources=proxmoximports/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=infrastructure.cluster.x-k8s.io,resources=proxmoximports/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.16.0/pkg/reconcile
func (r *ProxmoxImportReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	proxmoxImport := &infrav1.ProxmoxImport{}
	err := r.Get(ctx, req.NamespacedName, proxmoxImport)
	if err != nil {
		if apierrors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}

	importUrl, err := url.Parse(proxmoxImport.Spec.Url)
	if err != nil {
		return ctrl.Result{}, err
	}

	task, err := r.ProxmoxClient.Download(ctx, "pve01", importUrl)
	if err != nil {
		return ctrl.Result{}, err
	}

	task, err = r.ProxmoxClient.GetTask(ctx, string(task.UPID))
	if err != nil {
		return ctrl.Result{}, err
	}

	logger.Info("task status", task)

	return ctrl.Result{}, nil
}
