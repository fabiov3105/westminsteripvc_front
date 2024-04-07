package routes

import (
	"net/http"

	"github.com/fabiov3105/westminsteripvc_front/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/peoplesimple", controllers.Peoplesimple)
	http.HandleFunc("/peoplesimpleenabled", controllers.Peoplesimpleenabled)
	http.HandleFunc("/peoplesimpledisabled", controllers.Peoplesimpledisabled)
	http.HandleFunc("/newpeoplesimple", controllers.Newpeoplesimple)
	http.HandleFunc("/InsertNewPeopleSimple", controllers.InsertNewPeopleSimple)
	http.HandleFunc("/DeletePeopleSimple", controllers.Deletepeoplesimple)
	http.HandleFunc("/DisabledPeopleSimple", controllers.DisabledPeopleSimple)
	http.HandleFunc("/EnabledPeopleSimple", controllers.EnabledPeopleSimple)
	http.HandleFunc("/Updatepeoplesimple", controllers.Updatepeoplesimple)
	http.HandleFunc("/UpdatePeopleSimpleExec", controllers.Updatepeoplesimpleexec)
}
