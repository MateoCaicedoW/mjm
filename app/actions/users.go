package actions

import (
	"fmt"
	"mjm/app/models"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/pop/v6"
)

func List(c buffalo.Context) error {

	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	users := &models.Users{}

	q := tx.PaginateFromParams(c.Params())
	if err := q.All(users); err != nil {
		return err
	}

	c.Set("pagination", q.Paginator)
	c.Set("users", users)

	return c.Render(http.StatusOK, r.HTML("/user/index.plush.html"))
}

func Show(c buffalo.Context) error {

	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	user := &models.User{}

	if err := tx.Find(user, c.Param("user_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	c.Set("user", user)

	return c.Render(http.StatusOK, r.HTML("/user/show.plush.html"))
}

func New(c buffalo.Context) error {

	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	departments := []models.Departments{}

	if err := tx.All(&departments); err != nil {
		return err
	}

	c.Set("departments", departments)
	c.Set("user", &models.User{})

	return c.Render(http.StatusOK, r.HTML("/user/new.plush.html"))
}

func Create(c buffalo.Context) error {

	user := &models.User{}

	if err := c.Bind(user); err != nil {
		return err
	}

	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	verrs, err := tx.ValidateAndCreate(user)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		c.Set("errors", verrs)
		c.Set("user", user)
		return c.Render(http.StatusUnprocessableEntity, r.HTML("/user/new.plush.html"))
	}

	c.Flash().Add("success", "user.created.success")

	return c.Redirect(http.StatusSeeOther, "usersPath()", render.Data{"user_id": user.ID})
}

func Edit(c buffalo.Context) error {

	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	user := &models.User{}

	if err := tx.Find(user, c.Param("user_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	c.Set("user", user)

	return c.Render(http.StatusOK, r.HTML("/user/edit.plush.html"))
}

func Update(c buffalo.Context) error {

	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	user := &models.User{}

	if err := tx.Find(user, c.Param("user_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	if err := c.Bind(user); err != nil {
		return err
	}

	verrs, err := tx.ValidateAndUpdate(user)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		c.Set("errors", verrs)
		c.Set("user", user)
		return c.Render(http.StatusUnprocessableEntity, r.HTML("/user/edit.plush.html"))
	}

	c.Flash().Add("success", "user.updated.success")

	return c.Redirect(http.StatusSeeOther, "usersPath()", render.Data{"user_id": user.ID})
}

func Destroy(c buffalo.Context) error {

	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	user := &models.User{}

	if err := tx.Find(user, c.Param("user_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	if err := tx.Destroy(user); err != nil {
		return err
	}

	c.Flash().Add("success", "user.destroyed.success")

	return c.Redirect(http.StatusSeeOther, "usersPath()")
}
