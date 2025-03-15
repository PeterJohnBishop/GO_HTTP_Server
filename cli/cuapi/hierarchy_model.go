package cuapi

import tea "github.com/charmbracelet/bubbletea"

type HierarchyModel struct {
	user     User
	team     Team
	space    Space
	folder   Folder
	list     List
	task     Task
	spaces   []Space
	folders  []Folder
	lists    []List
	tasks    []Task
	message  string
	options  []string
	cursor   int
	selected map[int]struct{}
	lvl      int
}

func InitHierarchyModel(u User, t Team) HierarchyModel {
	return HierarchyModel{
		user:     u,
		team:     t,
		space:    Space{},
		folder:   Folder{},
		list:     List{},
		task:     Task{},
		spaces:   []Space{},
		folders:  []Folder{},
		lists:    []List{},
		tasks:    []Task{},
		message:  "",
		options:  []string{},
		cursor:   0,
		selected: make(map[int]struct{}),
		lvl:      0,
	}
}

func (m HierarchyModel) Init() tea.Cmd {
	return getSpaces(m.team.ID)
}

// JSON-to-Go https://mholt.github.io/json-to-go/

type Space struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Color          string `json:"color"`
	Private        bool   `json:"private"`
	Avatar         any    `json:"avatar"`
	AdminCanManage bool   `json:"admin_can_manage"`
	Statuses       []struct {
		ID         string `json:"id"`
		Status     string `json:"status"`
		Type       string `json:"type"`
		Orderindex int    `json:"orderindex"`
		Color      string `json:"color"`
	} `json:"statuses"`
	MultipleAssignees bool `json:"multiple_assignees"`
	Features          struct {
		DueDates struct {
			Enabled            bool `json:"enabled"`
			StartDate          bool `json:"start_date"`
			RemapDueDates      bool `json:"remap_due_dates"`
			RemapClosedDueDate bool `json:"remap_closed_due_date"`
		} `json:"due_dates"`
		Sprints struct {
			Enabled bool `json:"enabled"`
		} `json:"sprints"`
		TimeTracking struct {
			Enabled           bool `json:"enabled"`
			Harvest           bool `json:"harvest"`
			Rollup            bool `json:"rollup"`
			DefaultToBillable int  `json:"default_to_billable"`
		} `json:"time_tracking"`
		Points struct {
			Enabled bool `json:"enabled"`
		} `json:"points"`
		CustomItems struct {
			Enabled bool `json:"enabled"`
		} `json:"custom_items"`
		Priorities struct {
			Enabled    bool `json:"enabled"`
			Priorities []struct {
				Color      string `json:"color"`
				ID         string `json:"id"`
				Orderindex string `json:"orderindex"`
				Priority   string `json:"priority"`
			} `json:"priorities"`
		} `json:"priorities"`
		Tags struct {
			Enabled bool `json:"enabled"`
		} `json:"tags"`
		TimeEstimates struct {
			Enabled     bool `json:"enabled"`
			Rollup      bool `json:"rollup"`
			PerAssignee bool `json:"per_assignee"`
		} `json:"time_estimates"`
		CheckUnresolved struct {
			Enabled    bool `json:"enabled"`
			Subtasks   any  `json:"subtasks"`
			Checklists any  `json:"checklists"`
			Comments   any  `json:"comments"`
		} `json:"check_unresolved"`
		Zoom struct {
			Enabled bool `json:"enabled"`
		} `json:"zoom"`
		Milestones struct {
			Enabled bool `json:"enabled"`
		} `json:"milestones"`
		CustomFields struct {
			Enabled bool `json:"enabled"`
		} `json:"custom_fields"`
		RemapDependencies struct {
			Enabled bool `json:"enabled"`
		} `json:"remap_dependencies"`
		DependencyWarning struct {
			Enabled bool `json:"enabled"`
		} `json:"dependency_warning"`
		StatusPies struct {
			Enabled bool `json:"enabled"`
		} `json:"status_pies"`
		MultipleAssignees struct {
			Enabled bool `json:"enabled"`
		} `json:"multiple_assignees"`
		Emails struct {
			Enabled bool `json:"enabled"`
		} `json:"emails"`
		SchedulerEnabled bool `json:"scheduler_enabled"`
	} `json:"features"`
	Archived bool `json:"archived"`
	Members  []struct {
		User struct {
			ID             int    `json:"id"`
			Username       string `json:"username"`
			Color          string `json:"color"`
			ProfilePicture any    `json:"profilePicture"`
			Initials       string `json:"initials"`
		} `json:"user"`
	} `json:"members"`
}

type SpaceResponse struct {
	Spaces []Space `json:"spaces"`
}

type Folder struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Orderindex       int    `json:"orderindex"`
	OverrideStatuses bool   `json:"override_statuses"`
	Hidden           bool   `json:"hidden"`
	Space            struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		Access bool   `json:"access"`
	} `json:"space"`
	TaskCount string `json:"task_count"`
	Archived  bool   `json:"archived"`
	Statuses  []struct {
		ID         string `json:"id"`
		Status     string `json:"status"`
		Orderindex int    `json:"orderindex"`
		Color      string `json:"color"`
		Type       string `json:"type"`
	} `json:"statuses"`
	Lists []struct {
		ID               string `json:"id"`
		Name             string `json:"name"`
		Orderindex       int    `json:"orderindex"`
		Content          string `json:"content,omitempty"`
		Status           any    `json:"status"`
		Priority         any    `json:"priority"`
		Assignee         any    `json:"assignee"`
		TaskCount        int    `json:"task_count"`
		DueDate          any    `json:"due_date"`
		StartDate        any    `json:"start_date"`
		Archived         bool   `json:"archived"`
		OverrideStatuses bool   `json:"override_statuses"`
		Statuses         []struct {
			ID          string `json:"id"`
			Status      string `json:"status"`
			Orderindex  int    `json:"orderindex"`
			Color       string `json:"color"`
			Type        string `json:"type"`
			StatusGroup string `json:"status_group"`
		} `json:"statuses"`
		PermissionLevel string `json:"permission_level"`
	} `json:"lists"`
	PermissionLevel string `json:"permission_level"`
}

type FolderResponse struct {
	Folders []Folder `json:"folders"`
}

type List struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Deleted    bool   `json:"deleted"`
	Orderindex int    `json:"orderindex"`
	Content    string `json:"content"`
	Priority   any    `json:"priority"`
	Assignee   any    `json:"assignee"`
	DueDate    any    `json:"due_date"`
	StartDate  any    `json:"start_date"`
	Folder     struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		Hidden bool   `json:"hidden"`
		Access bool   `json:"access"`
	} `json:"folder"`
	Space struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		Access bool   `json:"access"`
	} `json:"space"`
	InboundAddress   string `json:"inbound_address"`
	Archived         bool   `json:"archived"`
	OverrideStatuses bool   `json:"override_statuses"`
	Statuses         []struct {
		ID          string `json:"id"`
		Status      string `json:"status"`
		Orderindex  int    `json:"orderindex"`
		Color       string `json:"color"`
		Type        string `json:"type"`
		StatusGroup string `json:"status_group"`
	} `json:"statuses"`
	PermissionLevel string `json:"permission_level"`
}

type ListResponse struct {
	Lists []List `json:"lists"`
}

type Task struct {
	ID           string `json:"id"`
	CustomID     any    `json:"custom_id"`
	CustomItemID int    `json:"custom_item_id"`
	Name         string `json:"name"`
	TextContent  string `json:"text_content"`
	Description  string `json:"description"`
	Status       struct {
		ID         string `json:"id"`
		Status     string `json:"status"`
		Color      string `json:"color"`
		Orderindex int    `json:"orderindex"`
		Type       string `json:"type"`
	} `json:"status"`
	Orderindex  string `json:"orderindex"`
	DateCreated string `json:"date_created"`
	DateUpdated string `json:"date_updated"`
	DateClosed  any    `json:"date_closed"`
	DateDone    any    `json:"date_done"`
	Archived    bool   `json:"archived"`
	Creator     struct {
		ID             int    `json:"id"`
		Username       string `json:"username"`
		Color          string `json:"color"`
		Email          string `json:"email"`
		ProfilePicture any    `json:"profilePicture"`
	} `json:"creator"`
	Assignees      []any `json:"assignees"`
	GroupAssignees []any `json:"group_assignees"`
	Watchers       []struct {
		ID             int    `json:"id"`
		Username       string `json:"username"`
		Color          string `json:"color"`
		Initials       string `json:"initials"`
		Email          string `json:"email"`
		ProfilePicture any    `json:"profilePicture"`
	} `json:"watchers"`
	Checklists     []any  `json:"checklists"`
	Tags           []any  `json:"tags"`
	Parent         any    `json:"parent"`
	TopLevelParent any    `json:"top_level_parent"`
	Priority       any    `json:"priority"`
	DueDate        string `json:"due_date"`
	StartDate      string `json:"start_date"`
	Points         any    `json:"points"`
	TimeEstimate   any    `json:"time_estimate"`
	TimeSpent      int    `json:"time_spent"`
	CustomFields   []struct {
		ID         string `json:"id"`
		Name       string `json:"name"`
		Type       string `json:"type"`
		TypeConfig struct {
			Fields []any `json:"fields"`
		} `json:"type_config,omitempty"`
		DateCreated    string `json:"date_created"`
		HideFromGuests bool   `json:"hide_from_guests"`
		Required       bool   `json:"required"`
		TypeConfig0    struct {
		} `json:"type_config0,omitempty"`
		Value         string `json:"value,omitempty"`
		ValueRichtext any    `json:"value_richtext,omitempty"`
		TypeConfig1   struct {
		} `json:"type_config1,omitempty"`
		TypeConfig2 struct {
			Sorting string `json:"sorting"`
			Options []struct {
				ID    string `json:"id"`
				Label string `json:"label"`
				Color any    `json:"color"`
			} `json:"options"`
		} `json:"type_config2,omitempty"`
		TypeConfig3 struct {
		} `json:"type_config3,omitempty"`
		TypeConfig4 struct {
		} `json:"type_config4,omitempty"`
		TypeConfig5 struct {
		} `json:"type_config5,omitempty"`
		TypeConfig6 struct {
		} `json:"type_config66,omitempty"`
		TypeConfig7 struct {
		} `json:"type_config7,omitempty"`
		TypeConfig8 struct {
		} `json:"type_config8,omitempty"`
		TypeConfig9 struct {
		} `json:"type_config9,omitempty"`
		TypeConfig10 struct {
			Ai struct {
				Format string `json:"format"`
				Source string `json:"source"`
			} `json:"ai"`
		} `json:"type_config10,omitempty"`
		TypeConfig11 struct {
			Ai struct {
				Format     string `json:"format"`
				Source     string `json:"source"`
				UpdateFrom string `json:"updateFrom"`
			} `json:"ai"`
		} `json:"type_confi11,omitempty"`
		TypeConfig12 struct {
		} `json:"type_config12,omitempty"`
		TypeConfig13 struct {
		} `json:"type_config13,omitempty"`
		TypeConfig14 struct {
			Simple           bool     `json:"simple"`
			Formula          string   `json:"formula"`
			Version          string   `json:"version"`
			ResetAt          int64    `json:"reset_at"`
			IsDynamic        bool     `json:"is_dynamic"`
			ReturnTypes      []string `json:"return_types"`
			CalculationState string   `json:"calculation_state"`
		} `json:"type_config14,omitempty"`
		TypeConfig15 struct {
			Sorting string `json:"sorting"`
			Options []struct {
				ID    string `json:"id"`
				Label string `json:"label"`
				Color string `json:"color"`
			} `json:"options"`
		} `json:"type_config15,omitempty"`
		TypeConfig16 struct {
			IncludeGroups      bool `json:"include_groups"`
			IncludeGuests      bool `json:"include_guests"`
			IncludeTeamMembers bool `json:"include_team_members"`
			SingleUser         bool `json:"single_user"`
		} `json:"type_config16omitempty"`
		TypeConfig17 struct {
			Tracking struct {
				Subtasks         bool `json:"subtasks"`
				Checklists       bool `json:"checklists"`
				AssignedComments bool `json:"assigned_comments"`
			} `json:"tracking"`
			CompleteOn    int  `json:"complete_on"`
			SubtaskRollup bool `json:"subtask_rollup"`
		} `json:"type_config17,omitempty"`
		TypeConfig18 struct {
			End   int `json:"end"`
			Start int `json:"start"`
		} `json:"type_config18,omitempty"`
		TypeConfig19 struct {
			Count     int    `json:"count"`
			CodePoint string `json:"code_point"`
		} `json:"type_config19,omitempty"`
		TypeConfig20 struct {
		} `json:"type_config20,omitempty"`
		TypeConfig21 struct {
			CodePoint  string `json:"code_point"`
			HideVoters bool   `json:"hide_voters"`
		} `json:"type_config21,omitempty"`
	} `json:"custom_fields"`
	Dependencies []any  `json:"dependencies"`
	LinkedTasks  []any  `json:"linked_tasks"`
	Locations    []any  `json:"locations"`
	TeamID       string `json:"team_id"`
	URL          string `json:"url"`
	Sharing      struct {
		Public               bool     `json:"public"`
		PublicShareExpiresOn any      `json:"public_share_expires_on"`
		PublicFields         []string `json:"public_fields"`
		Token                any      `json:"token"`
		SeoOptimized         bool     `json:"seo_optimized"`
	} `json:"sharing"`
	PermissionLevel string `json:"permission_level"`
	List            struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		Access bool   `json:"access"`
	} `json:"list"`
	Project struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		Hidden bool   `json:"hidden"`
		Access bool   `json:"access"`
	} `json:"project"`
	Folder struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		Hidden bool   `json:"hidden"`
		Access bool   `json:"access"`
	} `json:"folder"`
	Space struct {
		ID string `json:"id"`
	} `json:"space"`
	Attachments []any `json:"attachments"`
}

type TaskResponse struct {
	Tasks []Task `json:"tasks"`
}
