package echo

import (
	"errors"
	"net/http"

	mid "github.com/filipeandrade6/cooperagro/cmd/api/middleware/echo"
	"github.com/filipeandrade6/cooperagro/cmd/api/presenter"
	"github.com/filipeandrade6/cooperagro/domain/entity"
	"github.com/filipeandrade6/cooperagro/domain/usecase/unitofmeasure"
	"github.com/labstack/echo/v4"
)

func MakeUnitOfMeasureHandlers(e *echo.Group, service unitofmeasure.UseCase) {
	e.POST("/unitsofmeasure", createUnitOfMeasure(service), mid.AdminRequired)
	e.GET("/unitsofmeasure", readUnitOfMeasure(service))
	e.GET("/unitsofmeasure/:id", getUnitOfMeasure(service))
	e.PUT("/unitsofmeasure/:id", updateUnitOfMeasure(service), mid.AdminRequired)
	e.DELETE("/unitsofmeasure/:id", deleteUnitOfMeasure(service), mid.AdminRequired)
}

func createUnitOfMeasure(service unitofmeasure.UseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		var input presenter.EchoUnitOfMeasure
		if err := c.Bind(&input); err != nil {
			return echo.ErrBadRequest
		}

		id, err := service.CreateUnitOfMeasure(input.Name)
		if errors.Is(entity.ErrEntityAlreadyExists, err) {
			return c.NoContent(http.StatusConflict)
		}
		if errors.Is(entity.ErrInvalidEntity, err) {
			return echo.ErrBadRequest
		}
		if err != nil {
			return echo.ErrInternalServerError
		}

		return c.JSON(
			http.StatusCreated,
			echo.Map{"id": id.String()},
		)
	}
}

func getUnitOfMeasure(service unitofmeasure.UseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		idUUID, err := entity.StringToID(c.Param("id"))
		if err != nil {
			return echo.ErrBadRequest
		}

		data, err := service.GetUnitOfMeasureByID(idUUID)
		if errors.Is(err, entity.ErrNotFound) {
			return echo.ErrNotFound
		}
		if err != nil {
			return echo.ErrInternalServerError
		}

		return c.JSON(http.StatusOK, &presenter.UnitOfMeasure{
			ID:   data.ID,
			Name: data.Name,
		})
	}
}

func readUnitOfMeasure(service unitofmeasure.UseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		var data []*entity.UnitOfMeasure
		var err error

		name := c.QueryParam("name")
		if name != "" {
			data, err = service.SearchUnitOfMeasure(name)
		} else {
			data, err = service.ListUnitOfMeasure()
		}

		if errors.Is(err, entity.ErrNotFound) {
			return echo.ErrNotFound
		}
		if err != nil {
			return echo.ErrInternalServerError
		}

		var out []*presenter.UnitOfMeasure
		for _, d := range data {
			out = append(out, &presenter.UnitOfMeasure{
				ID:   d.ID,
				Name: d.Name,
			})
		}

		return c.JSON(http.StatusOK, out)
	}
}

func updateUnitOfMeasure(service unitofmeasure.UseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		idUUID, err := entity.StringToID(c.Param("id"))
		if err != nil {
			return echo.ErrBadRequest
		}

		var input presenter.UnitOfMeasure
		if err := c.Bind(&input); err != nil {
			return echo.ErrInternalServerError
		}

		err = service.UpdateUnitOfMeasure(&entity.UnitOfMeasure{
			ID:   idUUID,
			Name: input.Name,
		})
		switch {
		case errors.Is(entity.ErrInvalidEntity, err):
			return echo.ErrBadRequest

		case errors.Is(entity.ErrNotFound, err):
			return echo.ErrNotFound

		case errors.Is(entity.ErrEntityAlreadyExists, err):
			return c.NoContent(http.StatusConflict)

		case err != nil:
			return echo.ErrInternalServerError
		}

		return c.NoContent(http.StatusOK)
	}
}

func deleteUnitOfMeasure(service unitofmeasure.UseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		idUUID, err := entity.StringToID(c.Param("id"))
		if err != nil {
			return echo.ErrBadRequest
		}

		err = service.DeleteUnitOfMeasure(idUUID)
		if errors.Is(err, entity.ErrNotFound) {
			return echo.ErrNotFound
		}
		if err != nil {
			return echo.ErrInternalServerError
		}

		return c.NoContent(http.StatusOK)
	}
}
