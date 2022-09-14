package main

import (
	awx "crucible/x/awx"
	"log"
)

func main() {
	var (
		jobTemplateId = 7
		inventoryId   = 1
	)

	a := awx.NewAWX("http://awx.localhost", "admin", "admin", nil)
	result, err := a.PingService.Ping()
	if err != nil {
		log.Fatalf("Ping a err: %s", err)
	}

	log.Println("Ping a: ", result)

	// Run job
	result2, err := a.JobTemplateService.Launch(jobTemplateId, map[string]interface{}{
		"inventory": inventoryId,
	}, map[string]string{})
	if err != nil {
		log.Fatalf("Lauch err: %s", err)
	}

	log.Println("Launch Job Template: ", result2)

	resultJob, _, err2 := a.JobService.GetJobEvents(result2.Job, map[string]string{
		"order_by":  "start_line",
		"page_size": "1000000",
	})
	if err2 != nil {
		log.Fatalf("Get Job Events err: %s", err)
	}

	log.Println("Get Job Events: ", resultJob)
}
