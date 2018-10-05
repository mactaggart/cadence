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

namespace java com.uber.cadence.replicator

include "shared.thrift"
include "history.thrift"

enum ReplicationTaskType {
  Domain
  History
  SyncShardStatus
  SyncActivity
}

enum DomainOperation {
  Create
  Update
}

struct DomainTaskAttributes {
  05: optional DomainOperation domainOperation
  10: optional string id
  20: optional shared.DomainInfo info
  30: optional shared.DomainConfiguration config
  40: optional shared.DomainReplicationConfiguration replicationConfig
  50: optional i64 (js.type = "Long") configVersion
  60: optional i64 (js.type = "Long") failoverVersion
}

struct HistoryTaskAttributes {
  05: optional list<string> targetClusters
  10: optional string domainId
  20: optional string workflowId
  30: optional string runId
  40: optional i64 (js.type = "Long") firstEventId
  50: optional i64 (js.type = "Long") nextEventId
  60: optional i64 (js.type = "Long") version
  70: optional map<string, history.ReplicationInfo> replicationInfo
  80: optional shared.History history
  90: optional shared.History newRunHistory
}

struct SyncShardStatusTaskAttributes {
  10: optional string sourceCluster
  20: optional i64 (js.type = "Long") shardId
  30: optional i64 (js.type = "Long") timestamp
}

struct SyncActicvityTaskAttributes {
  10: optional string domainId
  20: optional string workflowId
  30: optional string runId
  40: optional i64 (js.type = "Long") version
  50: optional i64 (js.type = "Long") scheduledId
  60: optional i64 (js.type = "Long") scheduledTime
  70: optional i64 (js.type = "Long") startedId
  80: optional i64 (js.type = "Long") startedTime
  90: optional string requestId
  100: optional i64 (js.type = "Long") lastHeartbeatTime
  110: optional binary details
  120: optional i32 attempt
}

struct ReplicationTask {
  10: optional ReplicationTaskType taskType
  20: optional DomainTaskAttributes domainTaskAttributes
  30: optional HistoryTaskAttributes historyTaskAttributes
  40: optional SyncShardStatusTaskAttributes syncShardStatusTaskAttributes
  50: optional SyncActicvityTaskAttributes syncActicvityTaskAttributes
}

