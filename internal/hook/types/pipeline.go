package types

import "time"

// PipelineEventPayload contains the information for GitLab's pipeline status change event
type PipelineEventPayload struct {
	ObjectKind       string                   `json:"object_kind"`
	User             User                     `json:"user"`
	Project          Project                  `json:"project"`
	Commit           Commit                   `json:"commit"`
	ObjectAttributes PipelineObjectAttributes `json:"object_attributes"`
	MergeRequest     MergeRequest             `json:"merge_request"`
	Builds           []Build                  `json:"builds"`
}

// PipelineObjectAttributes contains pipeline specific GitLab object attributes information
type PipelineObjectAttributes struct {
	ID         int64      `json:"id"`
	Ref        string     `json:"ref"`
	Tag        bool       `json:"tag"`
	SHA        string     `json:"sha"`
	BeforeSHA  string     `json:"before_sha"`
	Source     string     `json:"source"`
	Status     string     `json:"status"`
	Stages     []string   `json:"stages"`
	CreatedAt  time.Time  `json:"created_at"`
	FinishedAt time.Time  `json:"finished_at"`
	Duration   int64      `json:"duration"`
	Variables  []Variable `json:"variables"`
}

// Variable contains pipeline variables
type Variable struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Build contains all of the GitLab Build information
type Build struct {
	ID            int64         `json:"id"`
	Stage         string        `json:"stage"`
	Name          string        `json:"name"`
	Status        string        `json:"status"`
	CreatedAt     time.Time     `json:"created_at"`
	StartedAt     time.Time     `json:"started_at"`
	FinishedAt    time.Time     `json:"finished_at"`
	When          string        `json:"when"`
	Manual        bool          `json:"manual"`
	User          User          `json:"user"`
	Runner        Runner        `json:"runner"`
	ArtifactsFile ArtifactsFile `json:"artifactsfile"`
}

// ArtifactsFile contains all of the GitLab artifact information
type ArtifactsFile struct {
	Filename string `json:"filename"`
	Size     string `json:"size"`
}
