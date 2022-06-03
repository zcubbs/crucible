package main

import (
	awxGo "github.com/zcubbs/crucible/awx"
	"log"
)

func main() {
	var (
		jobTemplateId = 7
		inventoryId   = 1
	)

	awx := awxGo.NewAWX("http://awx.localhost", "admin", "admin", nil)
	result, err := awx.PingService.Ping()
	if err != nil {
		log.Fatalf("Ping awx err: %s", err)
	}

	log.Println("Ping awx: ", result)

	// Run job
	result2, err := awx.JobTemplateService.Launch(jobTemplateId, map[string]interface{}{
		"inventory": inventoryId,
	}, map[string]string{})
	if err != nil {
		log.Fatalf("Lauch err: %s", err)
	}

	log.Println("Launch Job Template: ", result2)

	resultJob, _, err2 := awx.JobService.GetJobEvents(result2.Job, map[string]string{
		"order_by":  "start_line",
		"page_size": "1000000",
	})
	if err2 != nil {
		log.Fatalf("Get Job Events err: %s", err)
	}

	log.Println("Get Job Events: ", resultJob)
}
