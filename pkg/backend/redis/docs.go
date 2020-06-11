// Licensed to Preferred Networks, Inc. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Preferred Networks, Inc. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package redis

// Redis Task QueueBackend:
//
// Keys (prefix=_pftaskqueue:{keyprefix})
// ========================================================================
// {prefix}:queues                               (HASH) queueName -> queueUID
// {prefix}:queue:{uid}                          (String) TaskQueue
// {prefix}:queue:{uid}:tasks                    (Set) All taskUIDs
// {prefix}:queue:{uid}:task:pending             (List) Pending TaskUIDs
// {prefix}:queue:{uid}:task:deadletter          (List) TaskToDeadLetterErrors
// {prefix}:queue:{uid}:task:{uid}               (String) Task
// {prefix}:queue:{uid}:workers                  (Set) WorkerUIDs
// {prefix}:queue:{uid}:worker:{uid}             (String) Worker
// {prefix}:queue:{uid}:worker:{uid}:pending     (List) TaskUIDs (worker level pending queue)
// {prefix}:queue:{uid}:worker:{uid}:tasks       (Set) TaskUIDs
//
