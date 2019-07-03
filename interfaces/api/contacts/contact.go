package contacts

import (
	"github.com/gin-gonic/gin"
	"portfolio-api/usecases/contacts"
	"portfolio-api/usecases/mailer"
	"portfolio-api/usecases/logging"
	"portfolio-api/usecases/ports"
	"net/http"
)

// json定義
type contactJSON struct {
	Name string `binding:"required" json:"name"`
	MailAddress string `binding:"required" json:"mail_address"`
	Content  string `binding:"required" json:"content"`
}


type ContactController struct {
	Usecase *contacts.ContactUsecase
}

func NewContactController(mailer mailer.Mailer, logging logging.Logging) *ContactController {
	return &ContactController {
		Usecase: &contacts.ContactUsecase{
			Mailer: mailer,
			Logging: logging,
		},
	}
}

func (controller *ContactController) SendMail(c *gin.Context) {
	var json contactJSON

	if err := c.Bind(&json); err != nil {
		c.JSON(http.StatusBadRequest, "うーん、期待するものが送られてきていないなあ")
		return
	}

	input := &ports.ContactInputPort{
		Name: json.Name,
		MailAddress: json.MailAddress,
		Content:  json.Content,
	}

	err := controller.Usecase.SendMail(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, "メール送信失敗・・すまぬ多分APIかメールサーバーが死んでる・・・")
		return
	}

	c.JSON(http.StatusOK, "メール送信おけ")
}