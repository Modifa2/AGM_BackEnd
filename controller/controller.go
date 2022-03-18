package controller

import (
	"fmt"
	"gopls-workspace/models"
	"net/http"
	"os"
	"strconv"

	// "strconv"

	s "gopls-workspace/services"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// Register New User
func Register(c *gin.Context) {
	db := s.DB{}
	//
	var rb models.Returnblock
	var u models.UserRegisted

	if err := c.ShouldBindBodyWith(&u, binding.JSON); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	fmt.Println("Binded Data",u)

	//use function
	rows, err := db.SaveOnDB("AGM.create_users", u)
	if err != nil {
		fmt.Println("QueryRow failed: ", err.Error())
		errormessage := fmt.Sprintf("%v\n", err)
		c.JSON(http.StatusBadRequest, errormessage)
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, rb.New(false, "", err))
	} else {
		c.JSON(http.StatusOK, rb.New(true, "User Registered", rows))
		fmt.Println("User Data",rows)

	}


		request  := models.GetUserDetails{}
		request.ID_number = u.ID_number

		rows1, err := db.GetUserByID("agm.get_users", request)
		if err != nil {
		fmt.Println("QueryRow failed: ", err.Error())
		errormessage := fmt.Sprintf("%v\n", err)
		c.JSON(http.StatusBadRequest, errormessage)
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, rb.New(false, "", err))
	} else {
		c.JSON(http.StatusOK, rb.New(true, "", rows1))
			fmt.Println(rows1)

	}

	assignementRequest := models.AssignUserToQECode{}

	ID :=  strconv.Itoa(rows1[0].ID)

	assignementRequest.Association = rows1[0].Organization
	assignementRequest.ID_number  =  ID 
	assignementRequest.QRCode = rows1[0].QRCODE

	 rows2, err := db.GetUserByID("agm.assign_to_QrCode", assignementRequest)
		if err != nil {
		fmt.Println("QueryRow failed: ", err.Error())
		errormessage := fmt.Sprintf("%v\n", err)
		c.JSON(http.StatusBadRequest, errormessage)
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, rb.New(false, "", err))
	} else {
		c.JSON(http.StatusOK, rb.New(true, "", rows2))
			fmt.Println(rows2)
	}



	//agm.assign_to_QrCode


	//Assign User Using User Details
}
//check In User check_in_user

func CheckInUser(c *gin.Context) {
	db := s.DB{}
	//
	var rb models.Returnblock
	var u models.CheckInUser
	// var m models.MarkAsDisplayed
	if err := c.ShouldBindBodyWith(&u, binding.JSON); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	//use function
	rows, err := db.CheckInUser("agm.check_in_user", u)
	if err != nil {
		fmt.Println("QueryRow failed: ", err.Error())
		errormessage := fmt.Sprintf("%v\n", err)
		c.JSON(http.StatusBadRequest, errormessage)
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, rb.New(false, "", err))
	} else {
		fmt.Println(rows)
		c.JSON(http.StatusOK, rb.New(true, "", rows))
	}


}
//get_all_users


func GetAllUsers(c *gin.Context) {
	db := s.DB{}
	//
	var rb models.Returnblock
	var u models.GetAllType
	if err := c.ShouldBindBodyWith(&u, binding.JSON); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	//use function
	rows, err := db.GetAllUsers("AGM.get_all_users", u)
	if err != nil {
		fmt.Println("QueryRow failed: ", err.Error())
		errormessage := fmt.Sprintf("%v\n", err)
		c.JSON(http.StatusBadRequest, errormessage)
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, rb.New(false, "", err))
	} else {
		c.JSON(http.StatusOK, rb.New(true, "", rows))
	}
}

func MarkAsDisplayed(c *gin.Context) {
	db := s.DB{}
	//
	var rb models.Returnblock
	var u models.MarkAsDisplayed
	if err := c.ShouldBindBodyWith(&u, binding.JSON); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	fmt.Println("Back End Input ",u.ID_Number)	
	rows1, err := db.MarkAsDisplayed("agm.display_user", u)
	if err != nil {
		fmt.Println("QueryRow failed: ", err.Error())
		errormessage := fmt.Sprintf("%v\n", err)
		c.JSON(http.StatusBadRequest, errormessage)
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, rb.New(false, "", err))
	} else {
		fmt.Println(rows1)
		c.JSON(http.StatusOK, rb.New(true, "", rows1))
	}

}
//
func ReturnUser(c *gin.Context) {
	var rb models.Returnblock
	var l models.Login
	if err := c.ShouldBindJSON(&l); err != nil {
		c.JSON(http.StatusUnprocessableEntity, rb.New(false, "Invalid json provided", err))
	}
	db := *new(s.DB)
	rows, err := db.ReturnUser("AGM.get_users_by_id", l)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		errormessage := fmt.Sprintf("%v\n", err)
		c.JSON(http.StatusBadRequest, rb.New(false, errormessage, err))
	}else{
			c.JSON(http.StatusOK, rb.New(true, "Success", rows))
	fmt.Fprintf(os.Stdout, "QueryRow Success: %v\n", rows)

	}

}

func CheckNewUser(c *gin.Context) {
	var rb models.Returnblock

	var l models.EmptyGetter
	db := *new(s.DB)
		if err := c.ShouldBindJSON(&l); err != nil {
		c.JSON(http.StatusUnprocessableEntity, rb.New(false, "Invalid json provided", err))
	}
	rows, err := db.ReturnUser("AGM.get_users_latest", l)
if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		errormessage := fmt.Sprintf("%v\n", err)
		c.JSON(http.StatusBadRequest, rb.New(false, errormessage, err))
	}else{
	c.JSON(http.StatusOK, rb.New(true, "Success", rows))
	fmt.Println(rows)
	fmt.Fprintf(os.Stdout, "QueryRow Success: %v\n", rows)
	}
}

func GetAssociations(c *gin.Context) {
	var rb models.Returnblock

	var l models.EmptyGetter
	db := *new(s.DB)
		if err := c.ShouldBindJSON(&l); err != nil {
		c.JSON(http.StatusUnprocessableEntity, rb.New(false, "Invalid json provided", err))
	}
	rows, err := db.ReturnAssociations("AGM.get_all_associations", l)
if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		errormessage := fmt.Sprintf("%v\n", err)
		c.JSON(http.StatusBadRequest, rb.New(false, errormessage, err))
	}else{
	c.JSON(http.StatusOK, rb.New(true, "Success", rows))
	fmt.Println(rows)
	fmt.Fprintf(os.Stdout, "QueryRow Success: %v\n", rows)
	}
}