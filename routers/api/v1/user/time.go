// Copyright 2020 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package user

import (
	"net/http"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/convert"
	"code.gitea.io/gitea/routers/api/v1/utils"
)

// listUserTrackedTimes lists all traced times of the user
func listUserTrackedTimes(ctx *context.APIContext, opts models.FindTrackedTimesOptions) {

	var err error
	if opts.CreatedBeforeUnix, opts.CreatedAfterUnix, err = utils.GetQueryBeforeSince(ctx); err != nil {
		ctx.InternalServerError(err)
		return
	}

	trackedTimes, err := models.GetTrackedTimes(opts)
	if err != nil {
		ctx.Error(http.StatusInternalServerError, "GetTrackedTimesByUser", err)
		return
	}

	if err = trackedTimes.LoadAttributes(); err != nil {
		ctx.Error(http.StatusInternalServerError, "LoadAttributes", err)
		return
	}

	ctx.JSON(http.StatusOK, convert.ToTrackedTimeList(trackedTimes))
}

// ListUserTrackedTimes lists all tracked times of the user
func ListUserTrackedTimes(ctx *context.APIContext) {
	// swagger:operation GET /users/{username}/times user userAllTrackedTimes
	// ---
	// summary: List the user's tracked times
	// parameters:
	// - name: username
	//   in: path
	//   description: username of user
	//   type: string
	//   required: true
	// - name: page
	//   in: query
	//   description: page number of results to return (1-based)
	//   type: integer
	// - name: limit
	//   in: query
	//   description: page size of results
	//   type: integer
	// - name: since
	//   in: query
	//   description: Only show times updated after the given time. This is a timestamp in RFC 3339 format
	//   type: string
	//   format: date-time
	// - name: before
	//   in: query
	//   description: Only show times updated before the given time. This is a timestamp in RFC 3339 format
	//   type: string
	//   format: date-time
	// produces:
	// - application/json
	// responses:
	//   "200":
	//     "$ref": "#/responses/TrackedTimeList"

	user := GetUserByParams(ctx)

	opts := models.FindTrackedTimesOptions{
		ListOptions: utils.GetListOptions(ctx),
		UserID:      user.ID,
	}

	listUserTrackedTimes(ctx, opts)
}

// ListMyTrackedTimes lists all tracked times of the current user
func ListMyTrackedTimes(ctx *context.APIContext) {
	// swagger:operation GET /user/times user userCurrentTrackedTimes
	// ---
	// summary: List the current user's tracked times
	// parameters:
	// - name: page
	//   in: query
	//   description: page number of results to return (1-based)
	//   type: integer
	// - name: limit
	//   in: query
	//   description: page size of results
	//   type: integer
	// produces:
	// - application/json
	// parameters:
	// - name: since
	//   in: query
	//   description: Only show times updated after the given time. This is a timestamp in RFC 3339 format
	//   type: string
	//   format: date-time
	// - name: before
	//   in: query
	//   description: Only show times updated before the given time. This is a timestamp in RFC 3339 format
	//   type: string
	//   format: date-time
	// responses:
	//   "200":
	//     "$ref": "#/responses/TrackedTimeList"

	opts := models.FindTrackedTimesOptions{
		ListOptions: utils.GetListOptions(ctx),
		UserID:      ctx.User.ID,
	}

	listUserTrackedTimes(ctx, opts)
}
