package Time

import (
	"fmt"
	"strconv"
	"time"
)

func TimeFromTicks(ticks int64) time.Time {
	base := time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC).Unix()
	return time.Unix(ticks/10000000+base, ticks%10000000).UTC()
}

func main() {

	//caution : format string is `2006-01-02 15:04:05.000000000`
	current := time.Now()
	cur := strconv.Itoa(int(time.Now().Unix()))
	fmt.Println("origin.       : ", current)
	fmt.Println("origin string : ", current.String())
	fmt.Println("origin unix.  : ", current.Unix())
	fmt.Println("cur.          : ", cur)
	fmt.Println("sec.          : ", current.Second())
}
