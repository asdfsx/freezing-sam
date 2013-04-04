package db

import (
	"database/sql"
	"fmt"
)

func makeConnstr(dbaddr string, user string, passwd string, dbname string, charset string) (connstr string) {
	connstr = fmt.Sprintf("%v:%v@%v/%v?charset=%v", user, passwd, dbaddr, dbname, charset)
	return
}

func Getlogfileconfig(dbaddr string, user string, passwd string, dbname string, charset string) [][]string {
	connstr := makeConnstr(dbaddr, user, passwd, dbname, charset)
	db, _ := sql.Open("mysql", connstr)
	defer db.Close()

	result := make([][]string, 5)
	sql := "select log_type, log_path, log_name, log_regex, load_priority from conf_ptengine_logfile order by load_priority"
	rows, _ := db.Query(sql)
	for rows.Next() {
		var log_type string
		var log_path string
		var log_name string
		var log_regex string
		var load_priority string
		rows.Scan(&log_type, &log_path, &log_name, &log_regex, &load_priority)
		result = append(result, []string{log_type, log_path, log_name, log_regex, load_priority})
	}
	return result
}

func Getexcludeconfig(dbaddr string, user string, passwd string, dbname string, charset string) [][]string {
	connstr := makeConnstr(dbaddr, user, passwd, dbname, charset)
	db, _ := sql.Open("mysql", connstr)
	defer db.Close()

	result := make([][]string, 5)
	sql := "select siteid, type, rule from exclude where status=1 order by siteid, type, rule"
	rows, _ := db.Query(sql)
	for rows.Next() {
		var siteid string
		var stype string
		var rule string
		rows.Scan(&siteid, &stype, &rule)
		result = append(result, []string{siteid, stype, rule})
	}
	return result
}

func Getosinfo(dbaddr string, user string, passwd string, dbname string, charset string) [][]string {
	connstr := makeConnstr(dbaddr, user, passwd, dbname, charset)
	db, _ := sql.Open("mysql", connstr)
	defer db.Close()

	result := make([][]string, 5)
	sql := "select os_type, os_version, os_keyword, os_pattern, os_patterntype, priority,keyword_priority from conf_os_info order by os_keyword"
	rows, _ := db.Query(sql)
	for rows.Next() {
		var os_type string
		var os_version string
		var os_keyword string
		var os_pattern string
		var os_patterntype string
		var priority string
		var keyword_priority string
		rows.Scan(&os_type, &os_version, &os_keyword, &os_pattern, &os_patterntype, &priority, &keyword_priority)
		result = append(result, []string{os_type, os_version, os_keyword, os_pattern, os_patterntype, priority, keyword_priority})
	}
	return result
}

func Getuainfo(dbaddr string, user string, passwd string, dbname string, charset string) [][]string {
	connstr := makeConnstr(dbaddr, user, passwd, dbname, charset)
	db, _ := sql.Open("mysql", connstr)
	defer db.Close()

	result := make([][]string, 5)
	sql := "select brand, brand_eng, ua_name, ua_pattern, ua_patterntype, priority from conf_ua_info order by brand"
	rows, _ := db.Query(sql)
	for rows.Next() {
		var brand string
		var brand_eng string
		var ua_name string
		var ua_pattern string
		var ua_patterntype string
		var priority string
		rows.Scan(&brand, &brand_eng, &ua_name, &ua_pattern, &ua_patterntype, &priority)
		result = append(result, []string{brand, brand_eng, ua_name, ua_pattern, ua_patterntype, priority})
	}
	return result
}

func Getbrowserinfo(dbaddr string, user string, passwd string, dbname string, charset string) [][]string {
	connstr := makeConnstr(dbaddr, user, passwd, dbname, charset)
	db, _ := sql.Open("mysql", connstr)
	defer db.Close()

	result := make([][]string, 5)
	sql := "select browser_type, browser_version, browser_pattern, browser_patterntype, priority, keyword_priority, browser_keyword from conf_browser_info order by browser_type"

	rows, _ := db.Query(sql)
	for rows.Next() {
		var browser_type string
		var browser_version string
		var browser_pattern string
		var browser_patterntype string
		var priority string
		var keyword_priority string
		var browser_keyword string
		rows.Scan(&browser_type, &browser_version, &browser_pattern, &browser_patterntype, &priority, &keyword_priority, &browser_keyword)
		result = append(result, []string{browser_type, browser_version, browser_pattern, browser_patterntype, priority, keyword_priority, browser_keyword})
	}
	return result
}

func Getadvertiseinfo(dbaddr string, user string, passwd string, dbname string, charset string) [][]string {
	connstr := makeConnstr(dbaddr, user, passwd, dbname, charset)
	db, _ := sql.Open("mysql", connstr)
	defer db.Close()

	result := make([][]string, 5)
	sql := "select ad_id, ad_name, sid, utm_name_tag, utm_source_tag, utm_medium_tag, utm_term_tag, utm_content_tag, utm_cust1_tag, utm_cust2_tag, utm_cust3_tag, utm_cust4_tag, utm_cust5_tag, is_valid, update_dt from ref_ad_conf where is_valid=1 order by ad_id"

	rows, _ := db.Query(sql)
	for rows.Next() {
		var ad_id string
		var ad_name string
		var sid string
		var utm_name_tag string
		var utm_source_tag string
		var utm_medium_tag string
		var utm_term_tag string
		var utm_content_tag string
		var utm_cust1_tag string
		var utm_cust2_tag string
		var utm_cust3_tag string
		var utm_cust4_tag string
		var utm_cust5_tag string
		var is_valid string
		var update_dt string
		rows.Scan(&ad_id, &ad_name, &sid, &utm_name_tag, &utm_source_tag, &utm_medium_tag, &utm_term_tag, &utm_content_tag, &utm_cust1_tag, &utm_cust2_tag, &utm_cust3_tag, &utm_cust4_tag, &utm_cust5_tag, &is_valid, &update_dt)
		result = append(result, []string{ad_id, ad_name, sid, utm_name_tag, utm_source_tag, utm_medium_tag, utm_term_tag, utm_content_tag, utm_cust1_tag, utm_cust2_tag, utm_cust3_tag, utm_cust4_tag, utm_cust5_tag, is_valid, update_dt})
	}
	return result
}

func Getseachengineinfo(dbaddr string, user string, passwd string, dbname string, charset string) [][]string {
	connstr := makeConnstr(dbaddr, user, passwd, dbname, charset)
	db, _ := sql.Open("mysql", connstr)
	defer db.Close()

	result := make([][]string, 5)
	sql := "select engine_name, engine_info, engine_domain, engine_keyword, engine_charset from conf_searchengine_info order by engine_name"

	rows, _ := db.Query(sql)
	for rows.Next() {
		var engine_name string
		var engine_info string
		var engine_domain string
		var engine_keyword string
		var engine_charset string
		rows.Scan(&engine_name, &engine_info, &engine_domain, &engine_keyword, &engine_charset)
		result = append(result, []string{engine_name, engine_info, engine_domain, engine_keyword, engine_charset})
	}
	return result
}

func Getsiteinfo(dbaddr string, user string, passwd string, dbname string, charset string) [][]string {
	connstr := makeConnstr(dbaddr, user, passwd, dbname, charset)
	db, _ := sql.Open("mysql", connstr)
	defer db.Close()

	result := make([][]string, 5)
	sql := "select sid, timezone, status from ref_site_status where status <> 2 order by sid"

	rows, _ := db.Query(sql)
	for rows.Next() {
		var sid string
		var timezone string
		var status string
		rows.Scan(&sid, &timezone, &status)
		result = append(result, []string{sid, timezone, status})
	}
	return result
}

func Getsnsinfo(dbaddr string, user string, passwd string, dbname string, charset string) [][]string {
	connstr := makeConnstr(dbaddr, user, passwd, dbname, charset)
	db, _ := sql.Open("mysql", connstr)
	defer db.Close()

	result := make([][]string, 5)
	sql := "select sns_name, sns_info, sns_domain from conf_sns_info order by sns_name"

	rows, _ := db.Query(sql)
	for rows.Next() {
		var sns_name string
		var sns_info string
		var sns_domain string
		rows.Scan(&sns_name, &sns_info, &sns_domain)
		result = append(result, []string{sns_name, sns_info, sns_domain})
	}
	return result
}

func Getpagegroupinfo(dbaddr string, user string, passwd string, dbname string, charset string) [][]string {
	connstr := makeConnstr(dbaddr, user, passwd, dbname, charset)
	db, _ := sql.Open("mysql", connstr)
	defer db.Close()

	result := make([][]string, 5)
	sql := "select page_group_id, sid, url_template, pattern from ref_page_group_conf where is_valid=1 order by sid"

	rows, _ := db.Query(sql)
	for rows.Next() {
		var page_group_id string
		var sid string
		var url_template string
		var pattern string
		rows.Scan(&page_group_id, &sid, &url_template, &pattern)
		result = append(result, []string{page_group_id, sid, url_template, pattern})
	}
	return result
}

func Getcvinfo(dbaddr string, user string, passwd string, dbname string, charset string) [][]string {
	connstr := makeConnstr(dbaddr, user, passwd, dbname, charset)
	db, _ := sql.Open("mysql", connstr)
	defer db.Close()

	result := make([][]string, 5)
	sql := "select conversion_id, sid, source_url, source_pattern, url_template, pattern, prev_url, prev_pattern, step, cv_values from ref_conversion_conf where is_valid=1 order by sid"

	rows, _ := db.Query(sql)
	for rows.Next() {
		var conversion_id string
		var sid string
		var source_url string
		var source_pattern string
		var url_template string
		var pattern string
		var prev_url string
		var prev_pattern string
		var step string
		var cv_values string
		rows.Scan(&conversion_id, &sid, &source_url, &source_pattern, &url_template, &pattern, &prev_url, &prev_pattern, &step, &cv_values)
		result = append(result, []string{conversion_id, sid, source_url, source_pattern, url_template, pattern, prev_url, prev_pattern, step, cv_values})
	}
	return result
}
