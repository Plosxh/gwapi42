package main

import(
  "encoding/json"
  "net/http"
  "fmt"
  "os"
	"encoding/csv"
  "io/ioutil"
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

func main() {

  doEvery(5*time.Minute,pingApi)
/*
    foo2 := price{}
    getJson("https://api.guildwars2.com/v2/commerce/prices?id=19684", &foo2)
    fmt.Println(foo2.Buys.UnitePrice)*/
}


func doEvery(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}

func pingApi(t time.Time){
  url := "./prices.json"
  foo1 := new(price) // or &Foo{}
  getJson("https://api.guildwars2.com/v2/commerce/prices?id=19684", foo1)
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

addCsv(*foo1)
/*d := json.NewDecoder(strings.NewReader(jsontype))
d.UseNumber()
var x interface{}
if err := d.Decode(&x); err != nil {
    log.Fatal(err)
}
fmt.Printf("decoded to %#v\n", x)*/

}

func addCsv(p price) {

	f, err := os.OpenFile("values.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	w := csv.NewWriter(f)
	for i := 0; i < 10; i++ {
  		w.Write([]string{strconv.FormatInt(p.Id,10), strconv.FormatInt(p.Buys.Quantity,10), strconv.FormatInt(p.Buys.Unit_price,10), strconv.FormatInt(p.Sells.Quantity,10), strconv.FormatInt(p.Sells.Unit_price,10)})
	}
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


/*{
"name":"Corps de Koda",
"type":"Armor",
"level":80,
"rarity":"Exotic",
"vendor_value":370,
"default_skin":707,
"game_types":["Activity","Wvw","Dungeon","Pve"],
"flags":["AccountBound","NoSell","SoulBindOnUse"],
"restrictions":[],
"id":18154,
"chat_link":"[&AgHqRgAA]",
"icon":"https://render.guildwars2.com/file/5B3D97ACE0B564D69B7B020AE94016D49E01EEFC/218952.png",
"details":{
    "type":"Leggings",
    "weight_class":"Heavy",
    "defense":242,
    "infusion_slots":[],
    "infix_upgrade":{
        "id":154,
        "attributes":[{"attribute":"Precision","modifier":64},{"attribute":"Toughness","modifier":64},{"attribute":"ConditionDamage","modifier":90}]
      },
    "suffix_item_id":24857,
    "secondary_suffix_item_id":""
  }
}*/


/*{"name":"Arc long en bois tendre solide de feu",
"description":"",
"type":"Weapon",
"level":44,
"rarity":"Masterwork",
"vendor_value":120,
"default_skin":3942,
"game_types":["Activity","Wvw","Dungeon","Pve"],
"flags":["SoulBindOnUse"],
"restrictions":[],
"id":28445,
"chat_link":"[&AgEdbwAA]",
"icon":"https://render.guildwars2.com/file/C6110F52DF5AFE0F00A56F9E143E9732176DDDE9/65015.png",
"details":{
  "type":"LongBow",
  "damage_type":"Physical",
  "min_power":385,
  "max_power":452,
  "defense":0,
  "infusion_slots":[],
  "infix_upgrade":{
      "id":142,
      "attributes":[
      {"attribute":"Power","modifier":85},
      {"attribute":"Precision","modifier":61}
      ]},
  "suffix_item_id":24547,
  "secondary_suffix_item_id":""
  }
}*/


/*{
"name":"Barre aux baies d'Omnom",
"type":"Consumable",
"level":80,
"rarity":"Fine",
"vendor_value":33,
"game_types":["Wvw","Dungeon","Pve"],
"flags":["NoSell"],
"restrictions":[],
"id":12452,
"chat_link":"[&AgGkMAAA]",
"icon":"https://render.guildwars2.com/file/6BD5B65FBC6ED450219EC86DD570E59F4DA3791F/433643.png",
"details":{
    "type":"Food","
    duration_ms":1800000,
    "apply_count":1,
    "name":"Produit consommable",
    "icon":"https://render.guildwars2.com/file/779D3F0ABE5B46C09CFC57374DA8CC3A495F291C/436367.png",
    "description":"Découverte de magie +30%\nOr trouvé sur les monstres +40%\nExpérience à chaque ennemi tué +10%"
    }
}*/
