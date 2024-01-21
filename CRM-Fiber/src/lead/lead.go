package lead

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/akhil/CRM-Fiber/database"
	"github.com/gofiber/fiber"

)
//Lead is the data type, and lead is the name of the variable
type Lead struct{
	// golang not understand JSON,
	// the json we receive will be in lower capital, so the golang will understan the JSON.
	// it means golang understand on status and can only function on opertion on that status
	// that he understand, and then turn into JSON, and then can send it to respond, which he will understand.
	gorm.Model
    Name    		string 		`json:"name"`
    Company 		string 		`json:"company"`
    Email   		string 		`json:"email"`
    Phone   		int    		`json:"phone"`
}

func GetLeads(c *fiber.Ctx){//Ctx = contets
	db := database.DBCon
	var leads Lead
	db.Find(&leads)
	c.JSON(leads)
}

func GetLead(c *fiber.Ctx){
	id := c.Params("id")
	db := database.DBCon
	var lead Lead
	db.Find(&lead, id)
	//so we could also respond using c, and we will send the lead
	//in JSON format
	c.JSON(lead)
}

func NewLead(c *fiber.Ctx){
	db := database.DBCon
	lead := new(Lead)
	if err := c.BodyParser(lead); err != nil{
		c.Status(503).Send(err)
		return
	}
	db.Create(&lead)
	c.JSON(lead)
}

func DeleteLead(c *fiber.Ctx){
	id := c.Params("id")
	db := database.DBCon
	var lead Lead
	db.First(&lead, id)
	if lead.Name == ""{//if the name is actually nothing then it's mean it ain't there, and id not found.
		c.Status(500).Send("Not finding any lead with Id")
		return
	}
	db.Delete(&lead)
	c.Send("Lead Were Deleted Seccefully")
}