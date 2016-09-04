package webserver

import (
	"net/http"

	"github.com/7sDream/rikka/common/util"
	"github.com/7sDream/rikka/plugins"
)

func viewHandleFunc(w http.ResponseWriter, r *http.Request) {
	taskID := util.GetTaskIDByRequest(r)

	l.Debug("Recieve a view request of task", taskID)
	l.Debug("Send a url request of task", taskID, "to plugin manager")

	var pURL *plugins.URLJSON
	var err error
	if pURL, err = plugins.GetURL(taskID, r, nil); err != nil {
		// state is not finished or error when get url, use view.html
		templateFilePath := "templates/view.html"
		l.Warn("Error happened when get url of task", taskID, ":", err)
		l.Warn("State of task", taskID, "is not finished(or error happened), render with", templateFilePath)
		context.TaskID = taskID
		err = util.RenderTemplate(templateFilePath, w, context)
		if util.ErrHandle(w, err) {
			// RenderTemplate error
			l.Error("Erro when render template", templateFilePath, ":", err)
		} else {
			// successfully
			l.Debug("Render template", templateFilePath, "successfully")
		}
		return
	}

	// state is finished, use viewFinish.html
	l.Debug("Recieve url of task", taskID, ":", pURL.URL)
	templateFilePath := "templates/viewFinish.html"
	context.URL = pURL.URL
	err = util.RenderTemplate(templateFilePath, w, context)
	if util.ErrHandle(w, err) {
		// RenderTemplate error
		l.Error("Error happened when render template", templateFilePath, ":", err)
	} else {
		// successfully
		l.Debug("Render template", templateFilePath, "successfully")
	}
}

// ViewHandler handle requset ask for photo view page(/view/TaskID), use templates/view.html
// Only accept GET Method
var viewHandler = util.RequestFilter(
	"", "GET", l,
	viewHandleFunc,
)