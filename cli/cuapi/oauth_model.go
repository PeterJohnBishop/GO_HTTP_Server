package cuapi

import tea "github.com/charmbracelet/bubbletea"

type OAuthModel struct {
	user     User
	teams    []Team
	message  string
	options  []string
	cursor   int
	selected map[int]struct{}
}

func InitOAuthModel() OAuthModel {
	return OAuthModel{
		user:     User{},
		teams:    []Team{},
		message:  "",
		options:  []string{"Start OAuth"},
		cursor:   0,
		selected: make(map[int]struct{}),
	}
}

func (m OAuthModel) Init() tea.Cmd {
	return nil
}

type Team struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Color   string `json:"color"`
	Avatar  string `json:"avatar"`
	Members []User `json:"members"`
}
type TeamResponse struct {
	Teams []Team `json:"teams"`
}

type User struct {
	ID                int     `json:"id"`
	Username          string  `json:"username"`
	Email             string  `json:"email"`
	Color             string  `json:"color"`
	ProfilePicture    *string `json:"profilePicture"`
	Initials          string  `json:"initials"`
	WeekStartDay      int     `json:"week_start_day"`
	GlobalFontSupport bool    `json:"global_font_support"`
	Timezone          string  `json:"timezone"`
}

type UserResponse struct {
	User User `json:"user"`
}
