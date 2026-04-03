package config

import (
	"fmt"
	"os"

	"github.com/pelletier/go-toml/v2"
)

// Config holds all runtime configuration for the blog.
type Config struct {
	BlogName               string            `toml:"blog_name"`
	AuthorName             string            `toml:"author_name"`
	Lang                   string            `toml:"lang"`
	HeaderContent          string            `toml:"header_content"`
	FooterContent          string            `toml:"footer_content"`
	PostsPerPage           int               `toml:"posts_per_page"`
	ExcerptLength          int               `toml:"excerpt_length"`
	ShowUncategorized      bool              `toml:"show_uncategorized"`
	UncategorizedLabel     string            `toml:"uncategorized_label"`
	ShowRenderTime         bool              `toml:"show_render_time"`
	PostsDir               string            `toml:"posts_dir"`
	PostIndexFile          string            `toml:"post_index_file"`
	DateFormat             string            `toml:"date_format"`
	DefaultMetaDescription string            `toml:"default_meta_description"`
	CSSTheme               string            `toml:"css_theme"`
	CSP                    CSPConfig         `toml:"csp"`
	Cache                  CacheConfig       `toml:"cache"`
	Feed                   FeedConfig        `toml:"feed"`
	Menu                   MenuConfig        `toml:"menu"`
	MenuLinks              []MenuLink        `toml:"menu_links"`
	Categories             map[string]Category `toml:"categories"`
	Labels                 Labels            `toml:"labels"`
}

// FeedConfig holds RSS feed generation settings.
type FeedConfig struct {
	Enabled     bool   `toml:"enabled"`
	Title       string `toml:"title"`
	Description string `toml:"description"`
	BaseURL     string `toml:"base_url"`
	MaxItems    int    `toml:"max_items"`
	OutputFile  string `toml:"output_file"`
}

// CSPConfig holds Content-Security-Policy settings.
type CSPConfig struct {
	Enabled bool   `toml:"enabled"`
	Header  string `toml:"header"`
}

// CacheConfig holds HTTP cache-control settings.
type CacheConfig struct {
	Enabled      bool `toml:"enabled"`
	MaxAgePages  int  `toml:"max_age_pages"`  // seconds
	MaxAgeAssets int  `toml:"max_age_assets"` // seconds
}

// MenuLink is a static navigation link.
type MenuLink struct {
	Label string `toml:"label"`
	URL   string `toml:"url"`
}

// MenuConfig holds all navigation placement configuration.
type MenuConfig struct {
	Pinned     []MenuCategoryRef `toml:"pinned"`     // direct inline nav links
	Categories MenuDropdown      `toml:"categories"` // dropdown section
}

// MenuDropdown describes the dropdown section of the nav.
type MenuDropdown struct {
	Label string            `toml:"label"` // dropdown button text, e.g. "Writings"
	Item  []MenuCategoryRef `toml:"item"`
}

// MenuCategoryRef references a category by slug with an ordering hint.
type MenuCategoryRef struct {
	Category string `toml:"category"` // must match a key in [categories.*]
	Order    int    `toml:"order"`    // lower = earlier in list
}

// Category defines configuration for a post category folder.
type Category struct {
	BlogName      string `toml:"blog_name"`
	HeaderContent string `toml:"header_content"`
	Folder        string `toml:"folder"`
	Index         bool   `toml:"index"`
	Menu          bool   `toml:"menu"`
}

// Labels holds all user-visible UI strings.
type Labels struct {
	ReadMore             string `toml:"read_more"`
	BackToAll            string `toml:"back_to_all"`
	BackToCategory       string `toml:"back_to_category"`
	NotFoundTitle        string `toml:"not_found_title"`
	NotFoundMessage      string `toml:"not_found_message"`
	NoPostsInCategory    string `toml:"no_posts_in_category"`
	PaginationPrev       string `toml:"pagination_prev"`
	PaginationNext       string `toml:"pagination_next"`
	PageIndicator        string `toml:"page_indicator"`
	AuthorBy             string `toml:"author_by"`
	PostedOn             string `toml:"posted_on"`
	SearchTitle          string `toml:"search_title"`
	SearchPlaceholder    string `toml:"search_placeholder"`
	SearchButton         string `toml:"search_button"`
	SearchShowingResults string `toml:"search_showing_results"`
	SearchEmptyQuery     string `toml:"search_empty_query"`
	SearchNoResults      string `toml:"search_no_results"`
	SearchResultsTitle   string `toml:"search_results_title"`
}

// Load reads and parses the TOML config file at the given path.
func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("config: cannot read %s: %w", path, err)
	}
	var cfg Config
	if err := toml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("config: cannot parse %s: %w", path, err)
	}

	// Apply defaults
	if cfg.Lang == "" {
		cfg.Lang = "en"
	}
	if cfg.PostsDir == "" {
		cfg.PostsDir = "posts"
	}
	if cfg.PostIndexFile == "" {
		cfg.PostIndexFile = "posts/posts.index.json"
	}
	if cfg.DateFormat == "" {
		cfg.DateFormat = "2006-01-02"
	}
	if cfg.PostsPerPage == 0 {
		cfg.PostsPerPage = 10
	}
	if cfg.ExcerptLength == 0 {
		cfg.ExcerptLength = 200
	}

	if cfg.Menu.Categories.Label == "" {
		cfg.Menu.Categories.Label = "Writings"
	}

	// Cache defaults — only apply when nothing was explicitly configured
	if !cfg.Cache.Enabled && cfg.Cache.MaxAgePages == 0 && cfg.Cache.MaxAgeAssets == 0 {
		cfg.Cache.Enabled = true
	}
	if cfg.Cache.MaxAgePages == 0 {
		cfg.Cache.MaxAgePages = 3600
	}
	if cfg.Cache.MaxAgeAssets == 0 {
		cfg.Cache.MaxAgeAssets = 86400
	}

	// Feed defaults
	if cfg.Feed.MaxItems == 0 {
		cfg.Feed.MaxItems = 50
	}
	if cfg.Feed.OutputFile == "" {
		cfg.Feed.OutputFile = "feed.xml"
	}
	if cfg.Feed.Enabled && cfg.Feed.BaseURL == "" {
		return nil, fmt.Errorf("config: feed.enabled is true but feed.base_url is not set")
	}

	// Backfill folder key from map key when not set explicitly
	for slug, cat := range cfg.Categories {
		if cat.Folder == "" {
			cat.Folder = slug
			cfg.Categories[slug] = cat
		}
	}

	return &cfg, nil
}

// MustLoad is like Load but panics on error. Suitable for program startup.
func MustLoad(path string) *Config {
	cfg, err := Load(path)
	if err != nil {
		panic(err)
	}
	return cfg
}
