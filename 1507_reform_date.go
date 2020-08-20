import "regexp"
func reformatDate(date string) string {
    //https://regex101.com/
    months := map[string]string{"Jan":"01", "Feb":"02", "Mar":"03", "Apr":"04", "May":"05", "Jun":"06", "Jul":"07", "Aug":"08", "Sep":"09", "Oct":"10", "Nov":"11", "Dec":"12"}
    rgx := regexp.MustCompile(`(\d+)\w{2}\s(\w{3})\s(\d{4})`)
    rs := rgx.FindStringSubmatch(date)
    dd := string(rs[1])
    if len(dd) < 2 {
        dd = "0"+dd
    }
    mm := months[rs[2]]
    return string(rs[3])+"-"+mm+"-"+dd
}