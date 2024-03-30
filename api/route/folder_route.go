package route

import (
	"github.com/gin-gonic/gin"
	"stontactics/api/controller"
	"stontactics/bootstrap"
	"stontactics/domain"
	"stontactics/mongo"
	"stontactics/repository"
	"stontactics/usecase"
	"time"
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
}
