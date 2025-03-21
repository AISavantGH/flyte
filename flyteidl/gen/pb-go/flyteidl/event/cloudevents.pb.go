// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: flyteidl/event/cloudevents.proto

package event

import (
	core "github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/core"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is the cloud event parallel to the raw WorkflowExecutionEvent message. It's filled in with additional
// information that downstream consumers may find useful.
type CloudEventWorkflowExecution struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RawEvent        *WorkflowExecutionEvent `protobuf:"bytes,1,opt,name=raw_event,json=rawEvent,proto3" json:"raw_event,omitempty"`
	OutputInterface *core.TypedInterface    `protobuf:"bytes,2,opt,name=output_interface,json=outputInterface,proto3" json:"output_interface,omitempty"`
	// The following are ExecutionMetadata fields
	// We can't have the ExecutionMetadata object directly because of import cycle
	ArtifactIds        []*core.ArtifactID                `protobuf:"bytes,3,rep,name=artifact_ids,json=artifactIds,proto3" json:"artifact_ids,omitempty"`
	ReferenceExecution *core.WorkflowExecutionIdentifier `protobuf:"bytes,4,opt,name=reference_execution,json=referenceExecution,proto3" json:"reference_execution,omitempty"`
	Principal          string                            `protobuf:"bytes,5,opt,name=principal,proto3" json:"principal,omitempty"`
	// The ID of the LP that generated the execution that generated the Artifact.
	// Here for provenance information.
	// Launch plan IDs are easier to get than workflow IDs so we'll use these for now.
	LaunchPlanId *core.Identifier `protobuf:"bytes,6,opt,name=launch_plan_id,json=launchPlanId,proto3" json:"launch_plan_id,omitempty"`
	// We can't have the ExecutionMetadata object directly because of import cycle
	Labels map[string]string `protobuf:"bytes,7,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *CloudEventWorkflowExecution) Reset() {
	*x = CloudEventWorkflowExecution{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_event_cloudevents_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloudEventWorkflowExecution) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudEventWorkflowExecution) ProtoMessage() {}

func (x *CloudEventWorkflowExecution) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_event_cloudevents_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudEventWorkflowExecution.ProtoReflect.Descriptor instead.
func (*CloudEventWorkflowExecution) Descriptor() ([]byte, []int) {
	return file_flyteidl_event_cloudevents_proto_rawDescGZIP(), []int{0}
}

func (x *CloudEventWorkflowExecution) GetRawEvent() *WorkflowExecutionEvent {
	if x != nil {
		return x.RawEvent
	}
	return nil
}

func (x *CloudEventWorkflowExecution) GetOutputInterface() *core.TypedInterface {
	if x != nil {
		return x.OutputInterface
	}
	return nil
}

func (x *CloudEventWorkflowExecution) GetArtifactIds() []*core.ArtifactID {
	if x != nil {
		return x.ArtifactIds
	}
	return nil
}

func (x *CloudEventWorkflowExecution) GetReferenceExecution() *core.WorkflowExecutionIdentifier {
	if x != nil {
		return x.ReferenceExecution
	}
	return nil
}

func (x *CloudEventWorkflowExecution) GetPrincipal() string {
	if x != nil {
		return x.Principal
	}
	return ""
}

func (x *CloudEventWorkflowExecution) GetLaunchPlanId() *core.Identifier {
	if x != nil {
		return x.LaunchPlanId
	}
	return nil
}

func (x *CloudEventWorkflowExecution) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

type CloudEventNodeExecution struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RawEvent *NodeExecutionEvent `protobuf:"bytes,1,opt,name=raw_event,json=rawEvent,proto3" json:"raw_event,omitempty"`
	// The relevant task execution if applicable
	TaskExecId *core.TaskExecutionIdentifier `protobuf:"bytes,2,opt,name=task_exec_id,json=taskExecId,proto3" json:"task_exec_id,omitempty"`
	// The typed interface for the task that produced the event.
	OutputInterface *core.TypedInterface `protobuf:"bytes,3,opt,name=output_interface,json=outputInterface,proto3" json:"output_interface,omitempty"`
	// The following are ExecutionMetadata fields
	// We can't have the ExecutionMetadata object directly because of import cycle
	ArtifactIds []*core.ArtifactID `protobuf:"bytes,4,rep,name=artifact_ids,json=artifactIds,proto3" json:"artifact_ids,omitempty"`
	Principal   string             `protobuf:"bytes,5,opt,name=principal,proto3" json:"principal,omitempty"`
	// The ID of the LP that generated the execution that generated the Artifact.
	// Here for provenance information.
	// Launch plan IDs are easier to get than workflow IDs so we'll use these for now.
	LaunchPlanId *core.Identifier `protobuf:"bytes,6,opt,name=launch_plan_id,json=launchPlanId,proto3" json:"launch_plan_id,omitempty"`
	// We can't have the ExecutionMetadata object directly because of import cycle
	Labels map[string]string `protobuf:"bytes,7,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *CloudEventNodeExecution) Reset() {
	*x = CloudEventNodeExecution{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_event_cloudevents_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloudEventNodeExecution) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudEventNodeExecution) ProtoMessage() {}

func (x *CloudEventNodeExecution) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_event_cloudevents_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudEventNodeExecution.ProtoReflect.Descriptor instead.
func (*CloudEventNodeExecution) Descriptor() ([]byte, []int) {
	return file_flyteidl_event_cloudevents_proto_rawDescGZIP(), []int{1}
}

func (x *CloudEventNodeExecution) GetRawEvent() *NodeExecutionEvent {
	if x != nil {
		return x.RawEvent
	}
	return nil
}

func (x *CloudEventNodeExecution) GetTaskExecId() *core.TaskExecutionIdentifier {
	if x != nil {
		return x.TaskExecId
	}
	return nil
}

func (x *CloudEventNodeExecution) GetOutputInterface() *core.TypedInterface {
	if x != nil {
		return x.OutputInterface
	}
	return nil
}

func (x *CloudEventNodeExecution) GetArtifactIds() []*core.ArtifactID {
	if x != nil {
		return x.ArtifactIds
	}
	return nil
}

func (x *CloudEventNodeExecution) GetPrincipal() string {
	if x != nil {
		return x.Principal
	}
	return ""
}

func (x *CloudEventNodeExecution) GetLaunchPlanId() *core.Identifier {
	if x != nil {
		return x.LaunchPlanId
	}
	return nil
}

func (x *CloudEventNodeExecution) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

type CloudEventTaskExecution struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RawEvent *TaskExecutionEvent `protobuf:"bytes,1,opt,name=raw_event,json=rawEvent,proto3" json:"raw_event,omitempty"`
	// We can't have the ExecutionMetadata object directly because of import cycle
	Labels map[string]string `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *CloudEventTaskExecution) Reset() {
	*x = CloudEventTaskExecution{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_event_cloudevents_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloudEventTaskExecution) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudEventTaskExecution) ProtoMessage() {}

func (x *CloudEventTaskExecution) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_event_cloudevents_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudEventTaskExecution.ProtoReflect.Descriptor instead.
func (*CloudEventTaskExecution) Descriptor() ([]byte, []int) {
	return file_flyteidl_event_cloudevents_proto_rawDescGZIP(), []int{2}
}

func (x *CloudEventTaskExecution) GetRawEvent() *TaskExecutionEvent {
	if x != nil {
		return x.RawEvent
	}
	return nil
}

func (x *CloudEventTaskExecution) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

// This event is to be sent by Admin after it creates an execution.
type CloudEventExecutionStart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The execution created.
	ExecutionId *core.WorkflowExecutionIdentifier `protobuf:"bytes,1,opt,name=execution_id,json=executionId,proto3" json:"execution_id,omitempty"`
	// The launch plan used.
	LaunchPlanId *core.Identifier `protobuf:"bytes,2,opt,name=launch_plan_id,json=launchPlanId,proto3" json:"launch_plan_id,omitempty"`
	WorkflowId   *core.Identifier `protobuf:"bytes,3,opt,name=workflow_id,json=workflowId,proto3" json:"workflow_id,omitempty"`
	// Artifact inputs to the workflow execution for which we have the full Artifact ID. These are likely the result of artifact queries that are run.
	ArtifactIds []*core.ArtifactID `protobuf:"bytes,4,rep,name=artifact_ids,json=artifactIds,proto3" json:"artifact_ids,omitempty"`
	// Artifact inputs to the workflow execution for which we only have the tracking bit that's installed into the Literal's metadata by the Artifact service.
	ArtifactTrackers []string `protobuf:"bytes,5,rep,name=artifact_trackers,json=artifactTrackers,proto3" json:"artifact_trackers,omitempty"`
	Principal        string   `protobuf:"bytes,6,opt,name=principal,proto3" json:"principal,omitempty"`
}

func (x *CloudEventExecutionStart) Reset() {
	*x = CloudEventExecutionStart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_event_cloudevents_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloudEventExecutionStart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudEventExecutionStart) ProtoMessage() {}

func (x *CloudEventExecutionStart) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_event_cloudevents_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudEventExecutionStart.ProtoReflect.Descriptor instead.
func (*CloudEventExecutionStart) Descriptor() ([]byte, []int) {
	return file_flyteidl_event_cloudevents_proto_rawDescGZIP(), []int{3}
}

func (x *CloudEventExecutionStart) GetExecutionId() *core.WorkflowExecutionIdentifier {
	if x != nil {
		return x.ExecutionId
	}
	return nil
}

func (x *CloudEventExecutionStart) GetLaunchPlanId() *core.Identifier {
	if x != nil {
		return x.LaunchPlanId
	}
	return nil
}

func (x *CloudEventExecutionStart) GetWorkflowId() *core.Identifier {
	if x != nil {
		return x.WorkflowId
	}
	return nil
}

func (x *CloudEventExecutionStart) GetArtifactIds() []*core.ArtifactID {
	if x != nil {
		return x.ArtifactIds
	}
	return nil
}

func (x *CloudEventExecutionStart) GetArtifactTrackers() []string {
	if x != nil {
		return x.ArtifactTrackers
	}
	return nil
}

func (x *CloudEventExecutionStart) GetPrincipal() string {
	if x != nil {
		return x.Principal
	}
	return ""
}

var File_flyteidl_event_cloudevents_proto protoreflect.FileDescriptor

var file_flyteidl_event_cloudevents_proto_rawDesc = []byte{
	0x0a, 0x20, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x1a, 0x1a, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6c, 0x69,
	0x74, 0x65, 0x72, 0x61, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x66, 0x6c,
	0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x66, 0x6c, 0x79,
	0x74, 0x65, 0x69, 0x64, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x66,
	0x61, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x66, 0x6c,
	0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb2, 0x04,
	0x0a, 0x1b, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x57, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x43, 0x0a,
	0x09, 0x72, 0x61, 0x77, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x26, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x72, 0x61, 0x77, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x48, 0x0a, 0x10, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x66,
	0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x54, 0x79, 0x70,
	0x65, 0x64, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x52, 0x0f, 0x6f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x0c,
	0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x49, 0x44, 0x52, 0x0b, 0x61,
	0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x49, 0x64, 0x73, 0x12, 0x5b, 0x0a, 0x13, 0x72, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69,
	0x64, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x52, 0x12, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x45, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x69, 0x6e, 0x63,
	0x69, 0x70, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x69, 0x6e,
	0x63, 0x69, 0x70, 0x61, 0x6c, 0x12, 0x3f, 0x0a, 0x0e, 0x6c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x5f,
	0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x0c, 0x6c, 0x61, 0x75, 0x6e, 0x63, 0x68,
	0x50, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x4f, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64,
	0x6c, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0x93, 0x04, 0x0a, 0x17, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x4e, 0x6f, 0x64, 0x65, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3f,
	0x0a, 0x09, 0x72, 0x61, 0x77, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x72, 0x61, 0x77, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12,
	0x48, 0x0a, 0x0c, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x0a, 0x74,
	0x61, 0x73, 0x6b, 0x45, 0x78, 0x65, 0x63, 0x49, 0x64, 0x12, 0x48, 0x0a, 0x10, 0x6f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x64, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x52, 0x0f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x0c, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x5f,
	0x69, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x66, 0x6c, 0x79, 0x74,
	0x65, 0x69, 0x64, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x49, 0x44, 0x52, 0x0b, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x49, 0x64,
	0x73, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x12,
	0x3f, 0x0a, 0x0e, 0x6c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69,
	0x64, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x52, 0x0c, 0x6c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x50, 0x6c, 0x61, 0x6e, 0x49, 0x64,
	0x12, 0x4b, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x33, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x6f, 0x64, 0x65,
	0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x1a, 0x39, 0x0a,
	0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xe2, 0x01, 0x0a, 0x17, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3f, 0x0a, 0x09, 0x72, 0x61, 0x77, 0x5f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69,
	0x64, 0x6c, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x72, 0x61, 0x77,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x4b, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c,
	0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x54, 0x61, 0x73, 0x6b, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xef, 0x02,
	0x0a, 0x18, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x45, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x4d, 0x0a, 0x0c, 0x65, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2a, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x0b, 0x65, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x3f, 0x0a, 0x0e, 0x6c, 0x61, 0x75,
	0x6e, 0x63, 0x68, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x0c, 0x6c, 0x61,
	0x75, 0x6e, 0x63, 0x68, 0x50, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x3a, 0x0a, 0x0b, 0x77, 0x6f,
	0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x49, 0x64, 0x12, 0x3c, 0x0a, 0x0c, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x66,
	0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x41, 0x72, 0x74,
	0x69, 0x66, 0x61, 0x63, 0x74, 0x49, 0x44, 0x52, 0x0b, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63,
	0x74, 0x49, 0x64, 0x73, 0x12, 0x2b, 0x0a, 0x11, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74,
	0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x10, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72,
	0x73, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x42,
	0xbc, 0x01, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c,
	0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x10, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x6f, 0x72, 0x67, 0x2f,
	0x66, 0x6c, 0x79, 0x74, 0x65, 0x2f, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x70, 0x62, 0x2d, 0x67, 0x6f, 0x2f, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64,
	0x6c, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0xa2, 0x02, 0x03, 0x46, 0x45, 0x58, 0xaa, 0x02, 0x0e,
	0x46, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0xca, 0x02,
	0x0e, 0x46, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x5c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0xe2,
	0x02, 0x1a, 0x46, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x5c, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0f, 0x46,
	0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x3a, 0x3a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_flyteidl_event_cloudevents_proto_rawDescOnce sync.Once
	file_flyteidl_event_cloudevents_proto_rawDescData = file_flyteidl_event_cloudevents_proto_rawDesc
)

func file_flyteidl_event_cloudevents_proto_rawDescGZIP() []byte {
	file_flyteidl_event_cloudevents_proto_rawDescOnce.Do(func() {
		file_flyteidl_event_cloudevents_proto_rawDescData = protoimpl.X.CompressGZIP(file_flyteidl_event_cloudevents_proto_rawDescData)
	})
	return file_flyteidl_event_cloudevents_proto_rawDescData
}

var file_flyteidl_event_cloudevents_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_flyteidl_event_cloudevents_proto_goTypes = []interface{}{
	(*CloudEventWorkflowExecution)(nil),      // 0: flyteidl.event.CloudEventWorkflowExecution
	(*CloudEventNodeExecution)(nil),          // 1: flyteidl.event.CloudEventNodeExecution
	(*CloudEventTaskExecution)(nil),          // 2: flyteidl.event.CloudEventTaskExecution
	(*CloudEventExecutionStart)(nil),         // 3: flyteidl.event.CloudEventExecutionStart
	nil,                                      // 4: flyteidl.event.CloudEventWorkflowExecution.LabelsEntry
	nil,                                      // 5: flyteidl.event.CloudEventNodeExecution.LabelsEntry
	nil,                                      // 6: flyteidl.event.CloudEventTaskExecution.LabelsEntry
	(*WorkflowExecutionEvent)(nil),           // 7: flyteidl.event.WorkflowExecutionEvent
	(*core.TypedInterface)(nil),              // 8: flyteidl.core.TypedInterface
	(*core.ArtifactID)(nil),                  // 9: flyteidl.core.ArtifactID
	(*core.WorkflowExecutionIdentifier)(nil), // 10: flyteidl.core.WorkflowExecutionIdentifier
	(*core.Identifier)(nil),                  // 11: flyteidl.core.Identifier
	(*NodeExecutionEvent)(nil),               // 12: flyteidl.event.NodeExecutionEvent
	(*core.TaskExecutionIdentifier)(nil),     // 13: flyteidl.core.TaskExecutionIdentifier
	(*TaskExecutionEvent)(nil),               // 14: flyteidl.event.TaskExecutionEvent
}
var file_flyteidl_event_cloudevents_proto_depIdxs = []int32{
	7,  // 0: flyteidl.event.CloudEventWorkflowExecution.raw_event:type_name -> flyteidl.event.WorkflowExecutionEvent
	8,  // 1: flyteidl.event.CloudEventWorkflowExecution.output_interface:type_name -> flyteidl.core.TypedInterface
	9,  // 2: flyteidl.event.CloudEventWorkflowExecution.artifact_ids:type_name -> flyteidl.core.ArtifactID
	10, // 3: flyteidl.event.CloudEventWorkflowExecution.reference_execution:type_name -> flyteidl.core.WorkflowExecutionIdentifier
	11, // 4: flyteidl.event.CloudEventWorkflowExecution.launch_plan_id:type_name -> flyteidl.core.Identifier
	4,  // 5: flyteidl.event.CloudEventWorkflowExecution.labels:type_name -> flyteidl.event.CloudEventWorkflowExecution.LabelsEntry
	12, // 6: flyteidl.event.CloudEventNodeExecution.raw_event:type_name -> flyteidl.event.NodeExecutionEvent
	13, // 7: flyteidl.event.CloudEventNodeExecution.task_exec_id:type_name -> flyteidl.core.TaskExecutionIdentifier
	8,  // 8: flyteidl.event.CloudEventNodeExecution.output_interface:type_name -> flyteidl.core.TypedInterface
	9,  // 9: flyteidl.event.CloudEventNodeExecution.artifact_ids:type_name -> flyteidl.core.ArtifactID
	11, // 10: flyteidl.event.CloudEventNodeExecution.launch_plan_id:type_name -> flyteidl.core.Identifier
	5,  // 11: flyteidl.event.CloudEventNodeExecution.labels:type_name -> flyteidl.event.CloudEventNodeExecution.LabelsEntry
	14, // 12: flyteidl.event.CloudEventTaskExecution.raw_event:type_name -> flyteidl.event.TaskExecutionEvent
	6,  // 13: flyteidl.event.CloudEventTaskExecution.labels:type_name -> flyteidl.event.CloudEventTaskExecution.LabelsEntry
	10, // 14: flyteidl.event.CloudEventExecutionStart.execution_id:type_name -> flyteidl.core.WorkflowExecutionIdentifier
	11, // 15: flyteidl.event.CloudEventExecutionStart.launch_plan_id:type_name -> flyteidl.core.Identifier
	11, // 16: flyteidl.event.CloudEventExecutionStart.workflow_id:type_name -> flyteidl.core.Identifier
	9,  // 17: flyteidl.event.CloudEventExecutionStart.artifact_ids:type_name -> flyteidl.core.ArtifactID
	18, // [18:18] is the sub-list for method output_type
	18, // [18:18] is the sub-list for method input_type
	18, // [18:18] is the sub-list for extension type_name
	18, // [18:18] is the sub-list for extension extendee
	0,  // [0:18] is the sub-list for field type_name
}

func init() { file_flyteidl_event_cloudevents_proto_init() }
func file_flyteidl_event_cloudevents_proto_init() {
	if File_flyteidl_event_cloudevents_proto != nil {
		return
	}
	file_flyteidl_event_event_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_flyteidl_event_cloudevents_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloudEventWorkflowExecution); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_flyteidl_event_cloudevents_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloudEventNodeExecution); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_flyteidl_event_cloudevents_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloudEventTaskExecution); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_flyteidl_event_cloudevents_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloudEventExecutionStart); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_flyteidl_event_cloudevents_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_flyteidl_event_cloudevents_proto_goTypes,
		DependencyIndexes: file_flyteidl_event_cloudevents_proto_depIdxs,
		MessageInfos:      file_flyteidl_event_cloudevents_proto_msgTypes,
	}.Build()
	File_flyteidl_event_cloudevents_proto = out.File
	file_flyteidl_event_cloudevents_proto_rawDesc = nil
	file_flyteidl_event_cloudevents_proto_goTypes = nil
	file_flyteidl_event_cloudevents_proto_depIdxs = nil
}
