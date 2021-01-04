package structure

type SearchStruct struct {
	Code int64      `json:"code"`
	Data SearchData `json:"data"`
}

type SearchData struct {
	TagList  []TagList  `json:"tag_list"`
	BookList []BookList `json:"book_list"`
}

type BookList struct {
	BookID          string          `json:"book_id"`
	BookName        string          `json:"book_name"`
	Description     string          `json:"description"`
	BookSrc         string          `json:"book_src"`
	CategoryIndex   string          `json:"category_index"`
	Tag             string          `json:"tag"`
	TotalWordCount  string          `json:"total_word_count"`
	UpStatus        string          `json:"up_status"`
	UpdateStatus    string          `json:"update_status"`
	IsPaid          string          `json:"is_paid"`
	Discount        int64           `json:"discount"`
	DiscountEndTime string          `json:"discount_end_time"`
	Cover           string          `json:"cover"`
	AuthorName      string          `json:"author_name"`
	Uptime          string          `json:"uptime"`
	Newtime         string          `json:"newtime"`
	ReviewAmount    string          `json:"review_amount"`
	RewardAmount    string          `json:"reward_amount"`
	ChapterAmount   string          `json:"chapter_amount"`
	IsOriginal      int64           `json:"is_original"`
	TotalClick      int64           `json:"total_click"`
	MonthClick      int64           `json:"month_click"`
	WeekClick       int64           `json:"week_click"`
	MonthNoVipClick int64           `json:"month_no_vip_click"`
	WeekNoVipClick  int64           `json:"week_no_vip_click"`
	TotalRecommend  string          `json:"total_recommend"`
	MonthRecommend  string          `json:"month_recommend"`
	WeekRecommend   string          `json:"week_recommend"`
	TotalFavor      string          `json:"total_favor"`
	MonthFavor      string          `json:"month_favor"`
	WeekFavor       string          `json:"week_favor"`
	CurrentYp       string          `json:"current_yp"`
	TotalYp         string          `json:"total_yp"`
	CurrentBlade    string          `json:"current_blade"`
	TotalBlade      string          `json:"total_blade"`
	WeekFansValue   string          `json:"week_fans_value"`
	MonthFansValue  string          `json:"month_fans_value"`
	TotalFansValue  string          `json:"total_fans_value"`
	LastChapterInfo LastChapterInfo `json:"last_chapter_info"`
	BookType        string          `json:"book_type"`
	TransverseCover string          `json:"transverse_cover"`
}

type LastChapterInfo struct {
	ChapterID         string `json:"chapter_id"`
	BookID            string `json:"book_id"`
	ChapterIndex      string `json:"chapter_index"`
	ChapterTitle      string `json:"chapter_title"`
	Uptime            string `json:"uptime"`
	Mtime             string `json:"mtime"`
	RecommendBookInfo string `json:"recommend_book_info"`
}

type TagList struct {
	TagName string `json:"tag_name"`
	Num     int64  `json:"num"`
}