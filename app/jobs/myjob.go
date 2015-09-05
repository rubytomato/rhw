package jobs

import (
	"fmt"
	"time"

	"github.com/revel/modules/jobs/app/jobs"
	"github.com/revel/revel"
)

type MyJob struct {
	Name string
}

func (j MyJob) Run() {
  fmt.Printf("MyJob: %s %s\n", j.Name, time.Now().String())
}

func reminder() {
  fmt.Printf("reminder: every 1m : %s \n", time.Now().String())
}

func init() {
	revel.OnAppStart(func() {
		jobs.Schedule("0 */5 * ? * ?", MyJob{Name: "job1"})
		jobs.Schedule("cron.every_1h", MyJob{Name: "job2"})
		jobs.Schedule("cron.every_10m", MyJob{Name: "job3"})
		jobs.Schedule("@every 1m", jobs.Func(reminder))
		jobs.Now(MyJob{Name: "job Now"})
	})
}
