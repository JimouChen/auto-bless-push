# here put your blessing words
```go
func GetMeetingTime() time.Time {
	return time.Date(2022, 6, 1, 0, 0, 0, 0, time.Local)
}

func OurLoveTime(love bool) {
	meetTime := GetMeetingTime()
	loveTime := time.Since(meetTime)
	fmt.Println(loveTime)
	for love {
		loveTime++
	}
}
func main() {
	OurLoveTime(true)
}
```
