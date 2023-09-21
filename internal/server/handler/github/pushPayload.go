package github

type Owner struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Login     string `json:"login"`
	Id        int    `json:"id"`
	NodeId    string `json:"node_id"`
	Url       string `json:"url"`
	HtmlUrl   string `json:"html_url"`
	Type      string `json:"type"`
	SiteAdmin bool   `json:"site_admin"`
}

type Repository struct {
	Id        int    `json:"id"`
	NodeId    string `json:"node_id"`
	Name      string `json:"name"`
	Fullname  string `json:"full_name"`
	IsPrivate bool   `json:"private"`
	Owner     Owner  `json:"owner"`
	HtmlUrl   string `json:"html_url"`
	Desc      string `json:"description"`
	Url       string `json:"url"`
	CreatedAt int    `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	PushedAt  int    `json:"pushed_at"`
	GitUrl    string `json:"git_url"`
	SshUrl    string `json:"ssh_url"`
}

type Pusher struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Sender struct {
	Login string `json:"login"`
	Id    int    `json:"id"`
	Type  string `json:"type"`
}

type CommitUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type Commit struct {
	Id         string     `json:"id"`
	TreeId     string     `json:"tree_id"`
	IsDistinct bool       `json:"distinct"`
	Message    string     `json:"message"`
	Timestamp  string     `json:"timestamp"`
	Url        string     `json:"url"`
	Author     CommitUser `json:"author"`
	Committer  CommitUser `json:"committer"`
	Added      []string   `json:"added"`
	Removed    []string   `json:"removed"`
	Modified   []string   `json:"modified"`
}

type HookRequestPush struct {
	Ref        string     `json:"ref"`
	Before     string     `json:"before"`
	After      string     `json:"after"`
	Repository Repository `json:"repository"`
	Pusher     Pusher     `json:"pusher"`
	Sender     Sender     `json:"sender"`
	IsCreated  bool       `json:"created"`
	IsDeleted  bool       `json:"deleted"`
	IsForced   bool       `json:"forced"`
	Commits    []Commit   `json:"commits"`
	HeadCommit Commit     `json:"head_commit"`
}
