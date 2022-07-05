package types

import "time"

// MergeRequest contains all the GitLab merge request information
type MergeRequest struct {
	ID              int64      `json:"id"`
	TargetBranch    string     `json:"target_branch"`
	SourceBranch    string     `json:"source_branch"`
	SourceProjectID int64      `json:"source_project_id"`
	AssigneeID      int64      `json:"assignee_id"`
	AuthorID        int64      `json:"author_id"`
	Title           string     `json:"title"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	MilestoneID     int64      `json:"milestone_id"`
	State           string     `json:"state"`
	MergeStatus     string     `json:"merge_status"`
	TargetProjectID int64      `json:"target_project_id"`
	IID             int64      `json:"iid"`
	Description     string     `json:"description"`
	Position        int64      `json:"position"`
	LockedAt        time.Time  `json:"locked_at"`
	Source          Source     `json:"source"`
	Target          Target     `json:"target"`
	LastCommit      LastCommit `json:"last_commit"`
	WorkInProgress  bool       `json:"work_in_progress"`
	Assignee        User       `json:"assignee"`
	URL             string     `json:"url"`
}

// Source contains all the GitLab source information
type Source struct {
	Name              string `json:"name"`
	Description       string `json:"description"`
	WebURL            string `json:"web_url"`
	AvatarURL         string `json:"avatar_url"`
	GitSSHURL         string `json:"git_ssh_url"`
	GitHTTPURL        string `json:"git_http_url"`
	Namespace         string `json:"namespace"`
	VisibilityLevel   int64  `json:"visibility_level"`
	PathWithNamespace string `json:"path_with_namespace"`
	DefaultBranch     string `json:"default_branch"`
	Homepage          string `json:"homepage"`
	URL               string `json:"url"`
	SSHURL            string `json:"ssh_url"`
	HTTPURL           string `json:"http_url"`
}

// Target contains all the GitLab target information
type Target struct {
	Name              string `json:"name"`
	Description       string `json:"description"`
	WebURL            string `json:"web_url"`
	AvatarURL         string `json:"avatar_url"`
	GitSSHURL         string `json:"git_ssh_url"`
	GitHTTPURL        string `json:"git_http_url"`
	Namespace         string `json:"namespace"`
	VisibilityLevel   int64  `json:"visibility_level"`
	PathWithNamespace string `json:"path_with_namespace"`
	DefaultBranch     string `json:"default_branch"`
	Homepage          string `json:"homepage"`
	URL               string `json:"url"`
	SSHURL            string `json:"ssh_url"`
	HTTPURL           string `json:"http_url"`
}

// LastCommit contains all the GitLab last commit information
type LastCommit struct {
	ID        string    `json:"id"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
	URL       string    `json:"url"`
	Author    Author    `json:"author"`
}
