package ground

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/model/autotransport"
	"github.com/ozonmp/omp-bot/internal/service/autotransport/ground"
)

type Service interface {
	Describe(groundID uint64) (*autotransport.Ground, error)
	List(cursor uint64, limit uint64) ([]autotransport.Ground, error)
	Create(ground autotransport.Ground) (uint64, error)
	Update(groundID uint64, ground autotransport.Ground) error
	Remove(groundID uint64) (bool, error)

	Count() uint64
}

type AutotransportGroundCommander struct {
	bot     *tgbotapi.BotAPI
	service Service
}

func NewGroundCommander(bot *tgbotapi.BotAPI) *AutotransportGroundCommander {
	groundService := ground.NewAutotransportGroundService()
	return &AutotransportGroundCommander{
		bot:     bot,
		service: groundService,
	}
}

func (c *AutotransportGroundCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("AutotransportGroundCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *AutotransportGroundCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(msg)
	case "get":
		c.Get(msg)
	case "delete":
		c.Delete(msg)
	case "new":
		c.New(msg)
	case "edit":
		c.Edit(msg)
	default:
		c.Default(msg)
	}
}
