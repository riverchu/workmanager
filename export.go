package workmanager

import (
	"context"

	"golang.org/x/time/rate"
)

type (
	// Work ...
	Work func(target WorkTarget, configs map[WorkerName]WorkerConfig) (results []WorkTarget, err error)

	// WorkerBuilder ...
	WorkerBuilder func(ctx context.Context, args map[string]interface{}) Worker

	// StepRunner ...
	StepRunner func(work Work, workTarget WorkTarget, nexts ...func(WorkTarget))

	// StepCallback ...
	StepCallback func(...WorkTarget) []WorkTarget
)

// ================================================
// ================= Register API =================
// ================================================

// Register register worker and step runner/processor
func Register(
	from WorkStep,
	runner StepRunner,
	workers map[WorkerName]WorkerBuilder,
	to ...WorkStep,
) {
	defaultWorkerMgr.Register(from, runner, workers, to...)
}

// RegisterWorker register worker
func RegisterWorker(name WorkerName, builder WorkerBuilder) {
	defaultWorkerMgr.RegisterWorker(name, builder)
}

// RegisterStep register step runner and processor
func RegisterStep(from WorkStep, runner StepRunner, to ...WorkStep) {
	defaultWorkerMgr.RegisterStep(from, runner, to...)
}

// RegisterBeforeCallbacks ...
func RegisterBeforeCallbacks(step WorkStep, callbacks ...StepCallback) {
	defaultWorkerMgr.RegisterBeforeCallbacks(step, callbacks...)
}

// RegisterAfterCallbacks ...
func RegisterAfterCallbacks(step WorkStep, callbacks ...StepCallback) {
	defaultWorkerMgr.RegisterAfterCallbacks(step, callbacks...)
}

// ================================================
// ================== Server API ==================
// ================================================

// Serve daemon serve goroutine
func Serve(steps ...WorkStep) { defaultWorkerMgr.Serve(steps...) }

// Recv ...
func Recv(step WorkStep, target WorkTarget) error { return defaultWorkerMgr.Recv(step, target) }

// RecvFrom recv from chan
func RecvFrom(step WorkStep, recv <-chan WorkTarget) error {
	return defaultWorkerMgr.RecvFrom(step, recv)
}

// SetCacher set default work manager cacher
func SetCacher(c Cacher) { defaultWorkerMgr.SetCacher(c) }

// ================================================
// ================ Step Operation ================
// ================================================

// ListSteps list all steps
func ListSteps() []WorkStep { return defaultWorkerMgr.ListSteps() }

// PoolStastus return pool status
func PoolStatus(step WorkStep) (num, size int) { return defaultWorkerMgr.PoolStatus(step) }

// SetPool set pool size
func SetPool(size int, steps ...WorkStep) { defaultWorkerMgr.SetPool(size, steps...) }

// SetDefaultPool set default pool
func SetDefaultPool(size int) { defaultWorkerMgr.SetDefaultPool(size) }

// SetLimiter set limiter
func SetLimiter(rate rate.Limit, burst int, steps ...WorkStep) {
	defaultWorkerMgr.SetLimiter(rate, burst, steps...)
}

// SetDefaultLimiter set default limiter
func SetDefaultLimiter(rate rate.Limit, burst int) { defaultWorkerMgr.SetDefaultLimiter(rate, burst) }

// ================================================
// ================ Task Operation ================
// ================================================

// AddTask ...
func AddTask(task WorkTask) { defaultWorkerMgr.AddTask(task) }

// GetTask ...
func GetTask(token string) WorkTask { return defaultWorkerMgr.GetTask(token) }

// CancelTask ...
func CancelTask(token string) error { return defaultWorkerMgr.CancelTask(token) }

// ================================================
// ================ Pipe Operation ================
// ================================================

// GetSendChan ...
func GetSendChan(step WorkStep) chan<- WorkTarget {
	return defaultWorkerMgr.GetSendChan(step)
}

// SetSendChan ...
func SetSendChan(step WorkStep, ch chan<- WorkTarget) { defaultWorkerMgr.SetSendChan(step, ch) }

// MITMSendChan ...
func MITMSendChan(step WorkStep, newSendCh chan<- WorkTarget) (oldSendCh chan<- WorkTarget) {
	return defaultWorkerMgr.MITMSendChan(step, newSendCh)
}

// SetPipe ...
func SetPipe(step WorkStep, opts ...PipeOption) { defaultWorkerMgr.SetPipe(step, opts...) }
