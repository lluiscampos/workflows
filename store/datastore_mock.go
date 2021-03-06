// Copyright 2019 Northern.tech AS
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package store

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/mendersoftware/workflows/model"
)

// DataStoreMock is a mocked data storage service
type DataStoreMock struct {
	// Jobs contains the list of queued jobs
	Jobs    []model.Job
	channel chan *model.Job
}

// NewDataStoreMock initializes a DataStore mock object
func NewDataStoreMock() *DataStoreMock {

	return &DataStoreMock{
		channel: make(chan *model.Job),
	}
}

// InsertJob inserts the job in the queue
func (db *DataStoreMock) InsertJob(ctx context.Context, job *model.Job) (*model.Job, error) {
	job.ID = primitive.NewObjectID().Hex()
	db.Jobs = append(db.Jobs, *job)

	return job, nil
}

// GetJobs returns a channel of Jobs
func (db *DataStoreMock) GetJobs(ctx context.Context) <-chan *model.Job {
	return db.channel
}

// GetJobStatus returns the status of a Job
func (db *DataStoreMock) GetJobStatus(ctx context.Context, job *model.Job, fromStatus string, toStatus string) (*model.JobStatus, error) {
	return nil, nil
}

// UpdateJobAddResult add a task execution result to a job status
func (db *DataStoreMock) UpdateJobAddResult(ctx context.Context, jobStatus *model.JobStatus, data bson.M) error {
	return nil
}

// UpdateJobStatus set the task execution status for a job status
func (db *DataStoreMock) UpdateJobStatus(ctx context.Context, jobStatus *model.JobStatus, status string) error {
	return nil
}

// GetJobStatusByNameAndID get the task execution status for a job status bu Name and ID
func (db *DataStoreMock) GetJobStatusByNameAndID(ctx context.Context, name string, ID string) (*model.JobStatus, error) {
	return nil, nil
}

// Shutdown shuts down the datastore GetJobs process
func (db *DataStoreMock) Shutdown() {

}
