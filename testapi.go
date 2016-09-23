package main

import(
  "encoding/json"
  "encoding/xml"
  "net/http"
  "fmt"
  "os"
	"encoding/csv"
  "io/ioutil"
  "database/sql"
  "log"
  _ "github.com/mattn/go-sqlite3"
//"strings"
	"time"
  "strconv"
)



/*type Object struct {
    Name string
    Description string
    Type string
    Level int
    Rarity string
    VendorValue int
    DefaultSkin uint32
    GameType []string
    Flags []string
    Restrictions []string
    Id int
    ChatLink []string
    Icon string
    Details details
}

type details struct{
  DType string
  DamageType string
  MinPower int
  MaxPower int
  Defense int
  InfusSlot []string
  InfixUpgrade infixUpgrade
}

type infixUpgrade struct{
  Id int
  Attribute []attribute
  SuffixId int
  SecondSuffixId int
}

type attribute struct{
  Attribute string
  Mdifier string
}*/
type item struct{
  XMLName xml.Name `xml:"items"`
  BanqueMatXml []banqueMatXml `xml:"mat"`
}
type banqueMatXml struct{
  XMLName xml.Name `xml:"mat"`
  Id int64    `xml:"id"`
  Category int64 `xml:"category"`
  Count int64  `xml:"count"`
}

func main() {

  //https://api.guildwars2.com/v2/tokeninfo?access_token=65D84368-DA6E-9D4A-8B6E-70C0395432961B8D9A2D-1F1E-4F28-B484-9D0DFE20DBFF
   //clef := "65D84368-DA6E-9D4A-8B6E-70C0395432961B8D9A2D-1F1E-4F28-B484-9D0DFE20DBFF"

   checkBank(getClef())
  //doEvery(10*time.Second)
  //mesItems:=getItems()

  //fmt.Println(mesItems[0])


/*
    foo2 := price{}
    getJson("https://api.guildwars2.com/v2/commerce/prices?id=19684", &foo2)
    fmt.Println(foo2.Buys.UnitePrice)*/
}

func checkBank(key string)  {
  //var objets items
  var foo1 []banqueMatXml
  //var tempo1 []banqueMatXml
  //var foo2 banqueMatXml
  fmt.Println("allo ?")
getJson("https://api.guildwars2.com/v2/account/materials?access_token="+key, &foo1)
fmt.Println("vous avez : ",len(foo1)," objects dans vos materiaux.")
  fmt.Println("allo 2 ?")

writer,_ :=os.OpenFile("./gwitem.xml", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
enc := xml.NewEncoder(writer)
fmt.Println("test")
//enc.Indent("  ", "    ")
fmt.Println(len(foo1))
/*for i := 0; i < len(foo1); i++ {
  fmt.Println("bla")

  tempo1[i].Id = foo1[i].Id
  tempo1[i].Category = foo1[i].Category
  tempo1[i].Count = foo1[i].Count
  }*/
  foo2 := &item{BanqueMatXml:foo1}
  //objets = append(objets,foo1[i].Id,foo1[i].Category,foo1[i].Count)

  /*foo2.Id = foo1[i].Id
  foo2.Category = foo1[i].Category
  foo2.Count = foo1[i].Count*/
  //fmt.Println(foo2)
  if err := enc.Encode(foo2); err != nil {
      fmt.Printf("error: %v\n", err)
    }



  var monItem item
  xmlContent, _ := ioutil.ReadFile("gwitem.xml")
  err := xml.Unmarshal(xmlContent, &monItem)
  //err = xml.Unmarshal(xmlContent, &R)
  //fmt.Println(monItem)
  if err != nil { panic(err) }
  itemlen := len(monItem.BanqueMatXml)

    db, err := sql.Open("sqlite3", "./itemgw.db")
    if err != nil {
      log.Fatal(err)
    }
    defer db.Close()

    for i := 0; i < itemlen; i++ {

      _,err =db.Exec("INSERT INTO bank VALUES (NULL,"+strconv.FormatInt(monItem.BanqueMatXml[i].Id,10)+","+strconv.FormatInt(monItem.BanqueMatXml[i].Category,10)+","+strconv.FormatInt(monItem.BanqueMatXml[i].Count,10)+")")


    }


}

func getClef()string{

  var clef maClef
  xmlContent, _ := ioutil.ReadFile("apikey.xml")
  err := xml.Unmarshal(xmlContent, &clef)
  //err = xml.Unmarshal(xmlContent, &R)
  if err != nil { panic(err) }
  //fmt.Println(primlen)
  return clef.Id
}


func getItems()  []items{
url := "https://api.guildwars2.com/v2/items"

var mesItems []items

getJson(url,&mesItems)
return mesItems

}


func doEvery(d time.Duration) {
	for x := range time.Tick(d) {
    p := pingApi(x)
    addCsv(p)
	}
}

func pingApi(t time.Time) price{
  url := "./prices.json"
  var foo1 price // or &Foo{}
  fmt.Println(t.Clock)
  getJson("https://api.guildwars2.com/v2/commerce/prices?id=19684", &foo1)
  //getJson(url,foo1)
  println(foo1.Buys.Unit_price)
  fmt.Println(foo1)


  file, err := ioutil.ReadFile(url)
    if err != nil {
      fmt.Printf("File error: %v\n", err)
      os.Exit(1)
  }
  //fmt.Printf("%s\n", string(file))

 jsontype := new(price)
json.Unmarshal(file, &jsontype)
fmt.Printf("Results: %v\n", jsontype.Id)

//addCsv(*foo1)
/*d := json.NewDecoder(strings.NewReader(jsontype))
d.UseNumber()
var x interface{}
if err := d.Decode(&x); err != nil {
    log.Fatal(err)
}
fmt.Printf("decoded to %#v\n", x)*/
return foo1
}

func addCsv(p price) {

	f, err := os.OpenFile("values.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	w := csv.NewWriter(f)
	//for i := 0; i < 10; i++ {
  		w.Write([]string{strconv.FormatInt(p.Id,10), strconv.FormatInt(p.Buys.Quantity,10), strconv.FormatInt(p.Buys.Unit_price,10), strconv.FormatInt(p.Sells.Quantity,10), strconv.FormatInt(p.Sells.Unit_price,10)})
	//}
	w.Flush()
}


func getJson(url string, target interface{}) error {
    r, err := http.Get(url)
    if err != nil {
        return err
    }
    defer r.Body.Close()

    return json.NewDecoder(r.Body).Decode(target)
}


func calcFees(buy int64, sell int64) float64{
  var profit float64
  profit = ((float64(sell)*0.85)-float64(buy))
  return profit
}
