package route

import (
	"time"

	"github.com/neJok/StonTactics/api/controller"
	"github.com/neJok/StonTactics/bootstrap"
	"github.com/neJok/StonTactics/domain"
	"github.com/neJok/StonTactics/mongo"
	"github.com/neJok/StonTactics/repository"
	"github.com/neJok/StonTactics/usecase"

	"github.com/gin-gonic/gin"
)

func NewFolderRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	fr := repository.NewFolderRepository(db, domain.CollectionFolders)
	str := repository.NewStrategyRepository(db, domain.CollectionStrategies)
	spr := repository.NewSpreadingRepository(db, domain.CollectionSpreadouts)
	sc := &controller.FolderController{
		FolderUsecase:    usecase.NewFolderUsecase(fr, timeout),
		StrategyUsecase:  usecase.NewStrategyUsecase(str, timeout),
		SpreadingUsecase: usecase.NewSpreadingUsecase(spr, timeout),
	}
	group.GET("/folder", sc.FetchAll)
	group.POST("/folder", sc.Create)
	group.PUT("/folder/strategy", sc.AddStrategy)
	group.PUT("/folder/spreading", sc.AddSpreading)
	group.DELETE("/folder/spreading", sc.RemoveSpreading)
	group.DELETE("/folder/strategy", sc.RemoveStrategy)
}
