// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package history

import (
	h "github.com/uber/cadence/.gen/go/history"
	workflow "github.com/uber/cadence/.gen/go/shared"
	"github.com/uber/cadence/common/metrics"
	"github.com/uber/tchannel-go/thrift"
)

var _ Client = (*metricClient)(nil)

type metricClient struct {
	client        Client
	metricsClient metrics.Client
}

// NewMetricClient creates a new instance of Client that emits metrics
func NewMetricClient(client Client, metricsClient metrics.Client) Client {
	return &metricClient{
		client:        client,
		metricsClient: metricsClient,
	}
}

func (c *metricClient) StartWorkflowExecution(context thrift.Context,
	request *h.StartWorkflowExecutionRequest) (*workflow.StartWorkflowExecutionResponse, error) {
	c.metricsClient.IncCounter(metrics.HistoryClientStartWorkflowExecutionScope, metrics.CadenceRequests)

	sw := c.metricsClient.StartTimer(metrics.HistoryClientStartWorkflowExecutionScope, metrics.CadenceLatency)
	resp, err := c.client.StartWorkflowExecution(context, request)
	sw.Stop()

	if err != nil {
		c.metricsClient.IncCounter(metrics.HistoryClientStartWorkflowExecutionScope, metrics.CadenceFailures)
	}

	return resp, err
}

func (c *metricClient) GetWorkflowExecutionNextEventID(context thrift.Context,
	request *h.GetWorkflowExecutionNextEventIDRequest) (*h.GetWorkflowExecutionNextEventIDResponse, error) {
	c.metricsClient.IncCounter(metrics.HistoryClientGetWorkflowExecutionNextEventIDScope, metrics.CadenceRequests)

	sw := c.metricsClient.StartTimer(metrics.HistoryClientGetWorkflowExecutionNextEventIDScope, metrics.CadenceLatency)
	resp, err := c.client.GetWorkflowExecutionNextEventID(context, request)
	sw.Stop()

	if err != nil {
		c.metricsClient.IncCounter(metrics.HistoryClientGetWorkflowExecutionNextEventIDScope, metrics.CadenceFailures)
	}

	return resp, err
}

func (c *metricClient) RecordDecisionTaskStarted(context thrift.Context,
	request *h.RecordDecisionTaskStartedRequest) (*h.RecordDecisionTaskStartedResponse, error) {
	c.metricsClient.IncCounter(metrics.HistoryClientRecordDecisionTaskStartedScope, metrics.CadenceRequests)

	sw := c.metricsClient.StartTimer(metrics.HistoryClientRecordDecisionTaskStartedScope, metrics.CadenceLatency)
	resp, err := c.client.RecordDecisionTaskStarted(context, request)
	sw.Stop()

	if err != nil {
		c.metricsClient.IncCounter(metrics.HistoryClientRecordDecisionTaskStartedScope, metrics.CadenceFailures)
	}

	return resp, err
}

func (c *metricClient) RecordActivityTaskStarted(context thrift.Context,
	request *h.RecordActivityTaskStartedRequest) (*h.RecordActivityTaskStartedResponse, error) {
	c.metricsClient.IncCounter(metrics.HistoryClientRecordActivityTaskStartedScope, metrics.CadenceRequests)

	sw := c.metricsClient.StartTimer(metrics.HistoryClientRecordActivityTaskStartedScope, metrics.CadenceLatency)
	resp, err := c.client.RecordActivityTaskStarted(context, request)
	sw.Stop()

	if err != nil {
		c.metricsClient.IncCounter(metrics.HistoryClientRecordActivityTaskStartedScope, metrics.CadenceFailures)
	}

	return resp, err
}

func (c *metricClient) RespondDecisionTaskCompleted(context thrift.Context,
	request *h.RespondDecisionTaskCompletedRequest) error {
	c.metricsClient.IncCounter(metrics.HistoryClientRespondDecisionTaskCompletedScope, metrics.CadenceRequests)

	sw := c.metricsClient.StartTimer(metrics.HistoryClientRespondDecisionTaskCompletedScope, metrics.CadenceLatency)
	err := c.client.RespondDecisionTaskCompleted(context, request)
	sw.Stop()

	if err != nil {
		c.metricsClient.IncCounter(metrics.HistoryClientRespondDecisionTaskCompletedScope, metrics.CadenceFailures)
	}

	return err
}

func (c *metricClient) RespondActivityTaskCompleted(context thrift.Context,
	request *h.RespondActivityTaskCompletedRequest) error {
	c.metricsClient.IncCounter(metrics.HistoryClientRespondActivityTaskCompletedScope, metrics.CadenceRequests)

	sw := c.metricsClient.StartTimer(metrics.HistoryClientRespondActivityTaskCompletedScope, metrics.CadenceLatency)
	err := c.client.RespondActivityTaskCompleted(context, request)
	sw.Stop()

	if err != nil {
		c.metricsClient.IncCounter(metrics.HistoryClientRespondActivityTaskCompletedScope, metrics.CadenceFailures)
	}

	return err
}

func (c *metricClient) RespondActivityTaskFailed(context thrift.Context,
	request *h.RespondActivityTaskFailedRequest) error {
	c.metricsClient.IncCounter(metrics.HistoryClientRespondActivityTaskFailedScope, metrics.CadenceRequests)

	sw := c.metricsClient.StartTimer(metrics.HistoryClientRespondActivityTaskFailedScope, metrics.CadenceLatency)
	err := c.client.RespondActivityTaskFailed(context, request)
	sw.Stop()

	if err != nil {
		c.metricsClient.IncCounter(metrics.HistoryClientRespondActivityTaskFailedScope, metrics.CadenceFailures)
	}

	return err
}

func (c *metricClient) RespondActivityTaskCanceled(context thrift.Context,
	request *h.RespondActivityTaskCanceledRequest) error {
	c.metricsClient.IncCounter(metrics.HistoryClientRespondActivityTaskCanceledScope, metrics.CadenceRequests)

	sw := c.metricsClient.StartTimer(metrics.HistoryClientRespondActivityTaskCanceledScope, metrics.CadenceLatency)
	err := c.client.RespondActivityTaskCanceled(context, request)
	sw.Stop()

	if err != nil {
		c.metricsClient.IncCounter(metrics.HistoryClientRespondActivityTaskCanceledScope, metrics.CadenceFailures)
	}

	return err
}

func (c *metricClient) RecordActivityTaskHeartbeat(context thrift.Context,
	request *h.RecordActivityTaskHeartbeatRequest) (*workflow.RecordActivityTaskHeartbeatResponse, error) {
	c.metricsClient.IncCounter(metrics.HistoryClientRecordActivityTaskHeartbeatScope, metrics.CadenceRequests)

	sw := c.metricsClient.StartTimer(metrics.HistoryClientRecordActivityTaskHeartbeatScope, metrics.CadenceLatency)
	resp, err := c.client.RecordActivityTaskHeartbeat(context, request)
	sw.Stop()

	if err != nil {
		c.metricsClient.IncCounter(metrics.HistoryClientRecordActivityTaskHeartbeatScope, metrics.CadenceFailures)
	}

	return resp, err
}

func (c *metricClient) RequestCancelWorkflowExecution(context thrift.Context,
	request *h.RequestCancelWorkflowExecutionRequest) error {
	c.metricsClient.IncCounter(metrics.HistoryClientRequestCancelWorkflowExecutionScope, metrics.CadenceRequests)

	sw := c.metricsClient.StartTimer(metrics.HistoryClientRequestCancelWorkflowExecutionScope, metrics.CadenceLatency)
	err := c.client.RequestCancelWorkflowExecution(context, request)
	sw.Stop()

	if err != nil {
		c.metricsClient.IncCounter(metrics.HistoryClientRequestCancelWorkflowExecutionScope, metrics.CadenceFailures)
	}

	return err
}

func (c *metricClient) SignalWorkflowExecution(context thrift.Context,
	request *h.SignalWorkflowExecutionRequest) error {
	c.metricsClient.IncCounter(metrics.HistoryClientSignalWorkflowExecutionScope, metrics.CadenceRequests)

	sw := c.metricsClient.StartTimer(metrics.HistoryClientSignalWorkflowExecutionScope, metrics.CadenceLatency)
	err := c.client.SignalWorkflowExecution(context, request)
	sw.Stop()

	if err != nil {
		c.metricsClient.IncCounter(metrics.HistoryClientSignalWorkflowExecutionScope, metrics.CadenceFailures)
	}

	return err
}

func (c *metricClient) TerminateWorkflowExecution(context thrift.Context,
	request *h.TerminateWorkflowExecutionRequest) error {
	c.metricsClient.IncCounter(metrics.HistoryClientTerminateWorkflowExecutionScope, metrics.CadenceRequests)

	sw := c.metricsClient.StartTimer(metrics.HistoryClientTerminateWorkflowExecutionScope, metrics.CadenceLatency)
	err := c.client.TerminateWorkflowExecution(context, request)
	sw.Stop()

	if err != nil {
		c.metricsClient.IncCounter(metrics.HistoryClientTerminateWorkflowExecutionScope, metrics.CadenceFailures)
	}

	return err
}

func (c *metricClient) ScheduleDecisionTask(context thrift.Context,
	request *h.ScheduleDecisionTaskRequest) error {
	c.metricsClient.IncCounter(metrics.HistoryClientScheduleDecisionTaskScope, metrics.CadenceRequests)

	sw := c.metricsClient.StartTimer(metrics.HistoryClientScheduleDecisionTaskScope, metrics.CadenceLatency)
	err := c.client.ScheduleDecisionTask(context, request)
	sw.Stop()

	if err != nil {
		c.metricsClient.IncCounter(metrics.HistoryClientScheduleDecisionTaskScope, metrics.CadenceFailures)
	}

	return err
}

func (c *metricClient) RecordChildExecutionCompleted(context thrift.Context,
	request *h.RecordChildExecutionCompletedRequest) error {
	c.metricsClient.IncCounter(metrics.HistoryClientRecordChildExecutionCompletedScope, metrics.CadenceRequests)

	sw := c.metricsClient.StartTimer(metrics.HistoryClientRecordChildExecutionCompletedScope, metrics.CadenceLatency)
	err := c.client.RecordChildExecutionCompleted(context, request)
	sw.Stop()

	if err != nil {
		c.metricsClient.IncCounter(metrics.HistoryClientRecordChildExecutionCompletedScope, metrics.CadenceFailures)
	}

	return err
}
