package api

import "net/http"

func (controller *MetricsController) handlePing(w http.ResponseWriter, r *http.Request) {
	success, error := controller.healthcheckService.PingDatabase()
	if !success {
		failedDatabasePingResponse(error)(w, r)
	}
}
