package users

import (
	"fmt"
	"mjm/app/models"
	"mjm/app/render"
	"net/http"

	"github.com/gobuffalo/buffalo"

	"github.com/gobuffalo/pop/v6"
)

var (
	r = render.Engine
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

func View(c buffalo.Context) error {

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

	departments := models.Departments{}

	if err := tx.All(&departments); err != nil {
		return err
	}

	c.Set("departments", departments.Map())
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
	departments := models.Departments{}

	if err := tx.All(&departments); err != nil {
		return err
	}
	if verrs.HasAny() {
		c.Set("departments", departments.Map())
		c.Set("errors", verrs)
		c.Set("user", user)
		return c.Render(http.StatusUnprocessableEntity, r.HTML("/user/new.plush.html"))
	}

	c.Flash().Add("success", "User was created successfully")

	return c.Redirect(http.StatusSeeOther, "usersPath()")
}

func Edit(c buffalo.Context) error {

	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	user := &models.User{}
	departments := models.Departments{}

	if err := tx.Find(user, c.Param("user_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	if err := tx.All(&departments); err != nil {
		return err
	}
	c.Set("departments", departments.Map())
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

	c.Flash().Add("success", "User was updated successfully")

	return c.Redirect(http.StatusSeeOther, "usersPath()")
}

func Delete(c buffalo.Context) error {

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

	c.Flash().Add("success", "User was successfully destoyed")

	return c.Redirect(http.StatusSeeOther, "usersPath()")
}
