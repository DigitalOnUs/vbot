package requests

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"github.com/sirupsen/logrus"
	df "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
)

// ErrorInvalidInput basic errors
var ErrorInvalidInput = errors.New("Invalid hook request from webhook")

// ConsumeIntents default for proxy
func ConsumeIntents(c *gin.Context) {
	wr := df.WebhookRequest{}
	if err := jsonpb.Unmarshal(c.Request.Body, &wr); err != nil {
		logrus.WithError(err).Error(ErrorInvalidInput.Error())
		c.Status(http.StatusBadRequest)
		return
	}
	fmt.Println(wr.GetQueryResult().GetOutputContexts())
}
