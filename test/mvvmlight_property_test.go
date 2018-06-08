package test
import "testing"
import "fmt"
import "lightstrategy/models"


func TestProcessData(t *testing.T){

	p1,err1:=models.ProcessData("       /// <summary> /// /// </summary>  public double OriginalAmount { get; set; }")
	if err1!=nil{
		t.Error(err1)
	}
	fmt.Println(p1)

	p2,err2:=models.ProcessData("       /// <summary> /// 1:未执行/// </summary>  public double OriginalAmount { get; set; }")
	if err2!=nil{
		t.Error(err2)
	}
	fmt.Println(p2)

	p3,err3:=models.ProcessData("       /// <summary> /// 1:未执行 2:执行中 3:执行完成 4:暂停 5:强制释放/// </summary>  public double OriginalAmount { get; set; }")
	if err3!=nil{
		t.Error(err3)
	}
	fmt.Println(p3)
}